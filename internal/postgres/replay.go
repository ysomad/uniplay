package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
)

type replayRepo struct {
	log    *zap.Logger
	client *pgclient.Client
}

func NewReplayRepo(l *zap.Logger, c *pgclient.Client) *replayRepo {
	return &replayRepo{
		log:    l,
		client: c,
	}
}

func (r *replayRepo) MatchExists(ctx context.Context, matchID uuid.UUID) (found bool, err error) {
	row := r.client.Pool.QueryRow(ctx, "select exists(select 1 from match where id = $1)", matchID)

	if err = row.Scan(&found); err != nil {
		return false, err
	}

	return found, nil
}

func (r *replayRepo) SaveStats(ctx context.Context, m *dto.ReplayMatch, ps []dto.PlayerStat, ws []dto.PlayerWeaponStat) (*domain.Match, error) {
	txFunc := func(tx pgx.Tx) error {
		steamIDs := append(m.Team1.PlayerSteamIDs, m.Team2.PlayerSteamIDs...)

		if err := r.savePlayers(ctx, tx, steamIDs); err != nil {
			return err
		}

		m, err := r.saveTeams(ctx, tx, m)
		if err != nil {
			return err
		}

		players := m.MatchTeamPlayerList()

		if err = r.addPlayersToTeams(ctx, tx, players); err != nil {
			return err
		}

		if err = r.saveMatch(ctx, tx, m); err != nil {
			return err
		}

		if err = r.saveMatchHistory(ctx, tx, players); err != nil {
			return err
		}

		if err = r.saveMatchStats(ctx, tx, m.ID, ps); err != nil {
			return err
		}

		if err = r.saveMatchWeaponStats(ctx, tx, m.ID, ws); err != nil {
			return err
		}

		return nil
	}

	if err := pgx.BeginTxFunc(ctx, r.client.Pool, pgx.TxOptions{}, txFunc); err != nil {
		return nil, err
	}

	return &domain.Match{
		ID:       m.ID,
		MapName:  m.MapName,
		Duration: m.Duration,
		Team1: domain.MatchTeam{
			ID:             m.Team1.ID,
			ClanName:       m.Team1.ClanName,
			FlagCode:       m.Team1.FlagCode,
			Score:          m.Team1.Score,
			PlayerSteamIDs: m.Team1.PlayerSteamIDs,
		},
		Team2: domain.MatchTeam{
			ID:             m.Team2.ID,
			ClanName:       m.Team2.ClanName,
			FlagCode:       m.Team2.FlagCode,
			Score:          m.Team2.Score,
			PlayerSteamIDs: m.Team2.PlayerSteamIDs,
		},
		UploadedAt: m.UploadedAt,
	}, nil
}

func (r *replayRepo) savePlayers(ctx context.Context, tx pgx.Tx, steamIDs []uint64) error {
	b := r.client.Builder.
		Insert("player").
		Columns("steam_id")

	for _, steamID := range steamIDs {
		b = b.Values(steamID)
	}

	sql, args, err := b.Suffix("ON CONFLICT(steam_id) DO NOTHING").ToSql()
	if err != nil {
		return err
	}

	r.log.Debug("replayRepo", zap.String("query", sql))

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) saveTeams(ctx context.Context, tx pgx.Tx, m *dto.ReplayMatch) (*dto.ReplayMatch, error) {
	sql, args, err := r.client.Builder.
		Insert("team").
		Columns("clan_name, flag_code").
		Values(m.Team1.ClanName, m.Team2.FlagCode).
		Values(m.Team2.ClanName, m.Team2.FlagCode).
		Suffix("ON CONFLICT(clan_name) DO UPDATE").
		Suffix("SET clan_name = EXCLUDED.clan_name").
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, err
	}

	r.log.Debug("replayRepo", zap.String("query", sql))

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
		// TODO: use app error
		return nil, errors.New("got no team ids")
	}

	m.Team1.ID = teamIDs[0].ID
	m.Team2.ID = teamIDs[1].ID
	return m, nil
}

func (r *replayRepo) addPlayersToTeams(ctx context.Context, tx pgx.Tx, players []dto.MatchTeamPlayer) error {
	b := r.client.Builder.
		Insert("team_player").
		Columns("team_id, player_steam_id")

	for _, p := range players {
		b = b.Values(p.TeamID, p.SteamID)
	}

	sql, args, err := b.
		Suffix("ON CONFLICT (team_id, player_steam_id) DO NOTHING").
		ToSql()
	if err != nil {
		return err
	}

	r.log.Debug("replayRepo", zap.String("query", sql))

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) saveMatch(ctx context.Context, tx pgx.Tx, m *dto.ReplayMatch) error {
	sql, args, err := r.client.Builder.
		Insert("match").
		Columns("id, map_name, team1_id, team1_score, team2_id, team2_score, duration, uploaded_at").
		Values(m.ID, m.MapName, m.Team1.ID, m.Team1.Score, m.Team2.ID, m.Team2.Score, m.Duration, m.UploadedAt).
		ToSql()
	if err != nil {
		return err
	}

	r.log.Debug("replayRepo", zap.String("query", sql))

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) saveMatchHistory(ctx context.Context, tx pgx.Tx, mp []dto.MatchTeamPlayer) error {
	b := r.client.Builder.
		Insert("player_match").
		Columns("player_steam_id, match_id, team_id, match_state")

	for _, p := range mp {
		b = b.Values(p.SteamID, p.MatchID, p.TeamID, p.MatchState)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	r.log.Debug("replayRepo", zap.String("query", sql))

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) saveMatchStats(ctx context.Context, tx pgx.Tx, matchID uuid.UUID, stats []dto.PlayerStat) error {
	b := r.client.Builder.
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
			s.SteamID,
			matchID,
			s.Kills,
			s.HSKills,
			s.BlindKills,
			s.WallbangKills,
			s.NoScopeKills,
			s.ThroughSmokeKills,
			s.Deaths,
			s.Assists,
			s.FlashbangAssists,
			s.MVPCount,
			s.DamageTaken,
			s.DamageDealt,
			s.GrenadeDamageDealt,
			s.BlindedPlayers,
			s.BlindedTimes,
			s.BombsPlanted,
			s.BombsDefused,
		)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	r.log.Debug("replayRepo", zap.String("query", sql))

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) saveMatchWeaponStats(ctx context.Context, tx pgx.Tx, matchID uuid.UUID, ws []dto.PlayerWeaponStat) error {
	b := r.client.Builder.
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
			"chest_hits",
			"stomach_hits",
			"left_arm_hits",
			"right_arm_hits",
			"left_leg_hits",
			"right_leg_hits",
		)

	for _, s := range ws {
		b = b.Values(
			s.SteamID,
			matchID,
			s.WeaponID,
			s.Kills,
			s.HSKills,
			s.BlindKills,
			s.WallbangKills,
			s.NoScopeKills,
			s.ThroughSmokeKills,
			s.Deaths,
			s.Assists,
			s.DamageTaken,
			s.DamageDealt,
			s.Shots,
			s.HeadHits,
			s.ChestHits,
			s.StomachHits,
			s.LeftArmHits,
			s.RightArmHits,
			s.LeftLegHits,
			s.RightLegHits,
		)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	r.log.Debug("replayRepo", zap.String("query", sql))

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
