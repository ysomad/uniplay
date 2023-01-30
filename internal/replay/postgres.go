package replay

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/ysomad/uniplay/internal/pkg/otel"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

type Postgres struct {
	client *pgclient.Client
}

func NewPostgres(c *pgclient.Client) *Postgres {
	return &Postgres{
		client: c,
	}
}

func (s *Postgres) MatchExists(ctx context.Context, matchID uuid.UUID) (bool, error) {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.MatchExists")
	defer span.End()

	row := s.client.Pool.QueryRow(ctx, "select exists(select 1 from match where id = $1)", matchID)

	var matchFound bool
	if err := row.Scan(&matchFound); err != nil {
		return false, err
	}

	return matchFound, nil
}

func (s *Postgres) SaveStats(ctx context.Context, match *replayMatch, ps []*playerStat, ws []*weaponStat) error {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.SaveStats")
	defer span.End()

	txFunc := func(tx pgx.Tx) error {
		steamIDs := append(match.team1.players, match.team2.players...) //nolint:gocritic // why not ?

		if err := s.savePlayers(ctx, tx, steamIDs); err != nil {
			return err
		}

		savedMatch, err := s.saveTeams(ctx, tx, match)
		if err != nil {
			return err
		}

		teamPlayers := savedMatch.teamPlayers()

		if err := s.saveTeamPlayers(ctx, tx, teamPlayers); err != nil {
			return err
		}

		if err := s.saveMatch(ctx, tx, savedMatch); err != nil {
			return err
		}

		if err := s.savePlayersMatch(ctx, tx, teamPlayers); err != nil {
			return err
		}

		if err := s.savePlayerStats(ctx, tx, savedMatch.id, ps); err != nil {
			return err
		}

		if err := s.saveWeaponsStat(ctx, tx, savedMatch.id, ws); err != nil {
			return err
		}

		return nil
	}

	if err := pgx.BeginTxFunc(ctx, s.client.Pool, pgx.TxOptions{}, txFunc); err != nil {
		return err
	}

	return nil
}

func (s *Postgres) savePlayers(ctx context.Context, tx pgx.Tx, steamIDs []uint64) error {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.savePlayers")
	defer span.End()

	b := s.client.Builder.
		Insert("player").
		Columns("steam_id")

	for _, steamID := range steamIDs {
		b = b.Values(steamID)
	}

	sql, args, err := b.Suffix("ON CONFLICT(steam_id) DO NOTHING").ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

var errNoTeamIDsFound = errors.New("no team ids found")

// saveTeams saves match teams, if team with given clan name already exist, returns its id in match.
func (s *Postgres) saveTeams(ctx context.Context, tx pgx.Tx, m *replayMatch) (*replayMatch, error) {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.saveTeams")
	defer span.End()

	sql, args, err := s.client.Builder.
		Insert("team").
		Columns("clan_name, flag_code").
		Values(m.team1.clanName, m.team2.flagCode).
		Values(m.team2.clanName, m.team2.flagCode).
		Suffix("ON CONFLICT(clan_name) DO UPDATE").
		Suffix("SET clan_name = EXCLUDED.clan_name").
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	type teamID struct {
		ID int16
	}

	teamIDs, err := pgx.CollectRows(rows, pgx.RowToStructByPos[teamID])
	if err != nil {
		return nil, err
	}

	if len(teamIDs) != 2 {
		return nil, errNoTeamIDsFound
	}

	m.team1.id = teamIDs[0].ID
	m.team2.id = teamIDs[1].ID

	return m, nil
}

// saveTeamPlayers saves players to teams in which they was playing last game.
func (s *Postgres) saveTeamPlayers(ctx context.Context, tx pgx.Tx, players []teamPlayer) error {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.saveTeamPlayers")
	defer span.End()

	b := s.client.Builder.
		Insert("team_player").
		Columns("team_id, player_steam_id")

	for _, p := range players {
		b = b.Values(p.teamID, p.steamID)
	}

	sql, args, err := b.
		Suffix("ON CONFLICT (team_id, player_steam_id) DO NOTHING").
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (s *Postgres) saveMatch(ctx context.Context, tx pgx.Tx, m *replayMatch) error {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.saveMatch")
	defer span.End()

	sql, args, err := s.client.Builder.
		Insert("match").
		Columns("id, map_name, team1_id, team1_score, team2_id, team2_score, duration, uploaded_at").
		Values(m.id, m.mapName, m.team1.id, m.team1.score, m.team2.id, m.team2.score, m.duration, m.uploadedAt).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// savePlayersMatch saves match and its state to player match history.
func (s *Postgres) savePlayersMatch(ctx context.Context, tx pgx.Tx, players []teamPlayer) error {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.savePlayersMatch")
	defer span.End()

	b := s.client.Builder.
		Insert("player_match").
		Columns("player_steam_id, match_id, team_id, match_state")

	for _, p := range players {
		b = b.Values(p.steamID, p.matchID, p.teamID, p.matchState)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// savePlayerStats saves players statistic from specific match.
func (s *Postgres) savePlayerStats(ctx context.Context, tx pgx.Tx, matchID uuid.UUID, stats []*playerStat) error {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.savePlayerStats")
	defer span.End()

	b := s.client.Builder.
		Insert("player_match_stat").
		Columns(
			"player_steam_id",
			"match_id",
			"kills",
			"hs_kills",
			"blind_kills",
			"wallbang_kills",
			"noscope_kills",
			"through_smoke_kills",
			"deaths",
			"assists",
			"flashbang_assists",
			"mvp_count",
			"damage_taken",
			"damage_dealt",
			"grenade_damage_dealt",
			"blinded_players",
			"blinded_times",
			"bombs_planted",
			"bombs_defused",
		)

	for _, s := range stats {
		b = b.Values(
			s.steamID,
			matchID,
			s.kills,
			s.hsKills,
			s.blindKills,
			s.wallbangKills,
			s.noScopeKills,
			s.throughSmokeKills,
			s.deaths,
			s.assists,
			s.flashbangAssists,
			s.mvpCount,
			s.damageTaken,
			s.damageDealt,
			s.grenadeDamageDealt,
			s.blindedPlayers,
			s.blindedTimes,
			s.bombsPlanted,
			s.bombsDefused,
		)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// saveWeaponsStat saves players weapon statistic of specific match.
func (s *Postgres) saveWeaponsStat(ctx context.Context, tx pgx.Tx, matchID uuid.UUID, ws []*weaponStat) error {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Postgres.saveWeaponStat")
	defer span.End()

	b := s.client.Builder.
		Insert("player_match_weapon_stat").
		Columns(
			"player_steam_id",
			"match_id",
			"weapon_id",
			"kills",
			"hs_kills",
			"blind_kills",
			"wallbang_kills",
			"noscope_kills",
			"through_smoke_kills",
			"deaths",
			"assists",
			"damage_taken",
			"damage_dealt",
			"shots",
			"head_hits",
			"neck_hits",
			"chest_hits",
			"stomach_hits",
			"left_arm_hits",
			"right_arm_hits",
			"left_leg_hits",
			"right_leg_hits",
		)

	for _, s := range ws {
		b = b.Values(
			s.steamID,
			matchID,
			s.weaponID,
			s.kills,
			s.hsKills,
			s.blindKills,
			s.wallbangKills,
			s.noScopeKills,
			s.throughSmokeKills,
			s.deaths,
			s.assists,
			s.damageTaken,
			s.damageDealt,
			s.shots,
			s.headHits,
			s.neckHits,
			s.chestHits,
			s.stomachHits,
			s.leftArmHits,
			s.rightArmHits,
			s.leftLegHits,
			s.rightLegHits,
		)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
