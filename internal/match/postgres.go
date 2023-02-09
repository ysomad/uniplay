package match

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

type Postgres struct {
	tracer trace.Tracer
	client *pgclient.Client
}

func NewPostgres(t trace.Tracer, c *pgclient.Client) *Postgres {
	return &Postgres{
		tracer: t,
		client: c,
	}
}

func (p *Postgres) Exists(ctx context.Context, matchID uuid.UUID) (bool, error) {
	ctx, span := p.tracer.Start(ctx, "match.Postgres.Exists")
	defer span.End()

	row := p.client.Pool.QueryRow(ctx, "select exists(select 1 from match where id = $1)", matchID)

	var matchFound bool
	if err := row.Scan(&matchFound); err != nil {
		return false, err
	}

	return matchFound, nil
}

func (p *Postgres) GetScoreBoardRowsByID(ctx context.Context, matchID uuid.UUID) ([]*matchScoreBoardRow, error) {
	ctx, span := p.tracer.Start(ctx, "match.Postgres.GetScoreBoardRowsByID")
	defer span.End()

	sql, args, err := p.client.Builder.
		Select(
			"m.id as match_id",
			"mp.id as map_id",
			"mp.name as map_name",
			"mp.internal_name as map_internal_name",
			"mp.icon_url as map_icon_url",
			"pms.player_steam_id as steam_id",
			"p.display_name as player_name",
			"pm.team_id as team_id",
			"t.clan_name as team_name",
			"t.flag_code as team_flag_code",
			"tm.score as team_score",
			"tm.match_state as team_match_score",
			"pms.kills as kills",
			"pms.hs_kills as headshot_kills",
			"pms.deaths as deaths",
			"pms.assists as assists",
			"pms.damage_dealt as damage_dealt",
			"m.rounds as rounds_played",
			"pms.mvp_count as mvp_count",
			"m.duration as match_duration",
			"m.uploaded_at as match_uploaded_at",
		).
		From("match m").
		InnerJoin("map mp ON m.map_name = mp.internal_name").
		InnerJoin("player_match_stat pms ON m.id = pms.match_id").
		InnerJoin("player p ON p.steam_id = pms.player_steam_id").
		InnerJoin("player_match pm ON pms.player_steam_id = pm.player_steam_id AND m.id = pm.match_id").
		InnerJoin("team_match tm ON pm.team_id = tm.team_id AND m.id = tm.match_id").
		InnerJoin("team t ON tm.team_id = t.id").
		Where(sq.Eq{"m.id": matchID}).
		OrderBy("pms.kills DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	// TODO: отрефачить чтобы возвращался заполненный матч и строки для борда
	res, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[matchScoreBoardRow])
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *Postgres) CreateWithStats(ctx context.Context, match *replayMatch, ps []*playerStat, ws []*weaponStat) error { //nolint:gocognit // yes its T H I C C
	ctx, span := p.tracer.Start(ctx, "match.Postgres.CreateWithStats")
	defer span.End()

	txFunc := func(tx pgx.Tx) error {
		players := append(match.team1.players, match.team2.players...) //nolint:gocritic // why not ?

		if err := p.savePlayers(ctx, tx, players); err != nil {
			return err
		}

		savedMatch, err := p.saveTeams(ctx, tx, match)
		if err != nil {
			return err
		}

		teamPlayers := savedMatch.teamPlayers()

		err = p.saveTeamPlayers(ctx, tx, teamPlayers)
		if err != nil {
			return err
		}

		if err := p.saveMatch(ctx, tx, savedMatch); err != nil {
			return err
		}

		if err := p.savePlayersMatch(ctx, tx, teamPlayers); err != nil {
			return err
		}

		if err := p.savePlayerStats(ctx, tx, savedMatch.id, ps); err != nil {
			return err
		}

		if err := p.saveWeaponsStat(ctx, tx, savedMatch.id, ws); err != nil {
			return err
		}

		if err := p.saveTeamsMatch(ctx, tx, savedMatch); err != nil {
			return err
		}

		return nil
	}

	if err := pgx.BeginTxFunc(ctx, p.client.Pool, pgx.TxOptions{}, txFunc); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) saveTeamsMatch(ctx context.Context, tx pgx.Tx, m *replayMatch) error {
	sql, args, err := p.client.Builder.
		Insert("team_match").
		Columns("team_id, match_id, match_state, score").
		Values(m.team1.id, m.id, m.team1.matchState, m.team1.score).
		Values(m.team2.id, m.id, m.team2.matchState, m.team2.score).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) savePlayers(ctx context.Context, tx pgx.Tx, players []replayPlayer) error {
	b := p.client.Builder.
		Insert("player").
		Columns("id, steam_id, display_name")

	for _, p := range players {
		b = b.Values(p.id, p.steamID, p.displayName)
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
func (p *Postgres) saveTeams(ctx context.Context, tx pgx.Tx, m *replayMatch) (*replayMatch, error) {
	sql, args, err := p.client.Builder.
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
func (p *Postgres) saveTeamPlayers(ctx context.Context, tx pgx.Tx, players []teamPlayer) error {
	b := p.client.Builder.
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

func (p *Postgres) saveMatch(ctx context.Context, tx pgx.Tx, m *replayMatch) error {
	sql, args, err := p.client.Builder.
		Insert("match").
		Columns("id, map_name, rounds, duration, uploaded_at").
		Values(m.id, m.mapName, m.team1.score+m.team2.score, m.duration, m.uploadedAt).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// savePlayersMatch saves match and its state to player match history.
func (p *Postgres) savePlayersMatch(ctx context.Context, tx pgx.Tx, players []teamPlayer) error {
	b := p.client.Builder.
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
func (p *Postgres) savePlayerStats(ctx context.Context, tx pgx.Tx, matchID uuid.UUID, stats []*playerStat) error {
	b := p.client.Builder.
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
func (p *Postgres) saveWeaponsStat(ctx context.Context, tx pgx.Tx, matchID uuid.UUID, ws []*weaponStat) error {
	b := p.client.Builder.
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

func (p *Postgres) DeleteByID(ctx context.Context, matchID uuid.UUID) error {
	ctx, span := p.tracer.Start(ctx, "match.Postgres.DeleteByID")
	defer span.End()

	txFunc := func(tx pgx.Tx) error {
		if err := p.deletePlayerStats(ctx, tx, matchID); err != nil {
			return err
		}

		if err := p.deleteWeaponStats(ctx, tx, matchID); err != nil {
			return err
		}

		if err := p.deletePlayersMatch(ctx, tx, matchID); err != nil {
			return err
		}

		if err := p.deleteTeamsMatch(ctx, tx, matchID); err != nil {
			return err
		}

		if err := p.deleteMatch(ctx, tx, matchID); err != nil {
			return err
		}

		return nil
	}

	if err := pgx.BeginTxFunc(ctx, p.client.Pool, pgx.TxOptions{}, txFunc); err != nil {
		return err
	}

	return nil
}

// deletePlayerStats deletes all players stats associated with match.
func (p *Postgres) deletePlayerStats(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
	sql, args, err := p.client.Builder.
		Delete("player_match_stat").
		Where(sq.Eq{"match_id": matchID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// deleteWeaponStats deletes all players weapon stats associated with match.
func (p *Postgres) deleteWeaponStats(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
	sql, args, err := p.client.Builder.
		Delete("player_match_weapon_stat").
		Where(sq.Eq{"match_id": matchID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// deletePlayersMatch deletes match from player history of matches.
func (p *Postgres) deletePlayersMatch(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
	sql, args, err := p.client.Builder.
		Delete("player_match").
		Where(sq.Eq{"match_id": matchID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// deleteMatch deletes match by id.
func (p *Postgres) deleteMatch(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
	sql, args, err := p.client.Builder.
		Delete("match").
		Where(sq.Eq{"id": matchID}).
		ToSql()
	if err != nil {
		return err
	}

	ct, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() <= 0 {
		return domain.ErrMatchNotFound
	}

	return nil
}

// deleteTeamsMatch deletes match from team history.
func (p *Postgres) deleteTeamsMatch(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
	sql, args, err := p.client.Builder.
		Delete("team_match").
		Where(sq.Eq{"match_id": matchID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
