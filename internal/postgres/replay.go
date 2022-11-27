package postgres

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ysomad/pgxatomic"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayRepo struct {
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
}

func NewReplayRepo(p pgxatomic.Pool, b sq.StatementBuilderType) *replayRepo {
	return &replayRepo{
		pool:    p,
		builder: b,
	}
}

func (r *replayRepo) SavePlayers(ctx context.Context, mp dto.MatchPlayers) error {
	// build queries
	playerSB := r.builder.
		Insert("player").
		Columns("steam_id, create_time, update_time")

	matchPlayerSB := r.builder.
		Insert("match_player").
		Columns("match_id, player_steam_id, team_name, match_state")

	for _, player := range mp.Players {
		playerSB = playerSB.Values(player.SteamID, mp.CreateTime, mp.CreateTime)
		matchPlayerSB = matchPlayerSB.Values(mp.MatchID, player.SteamID, player.TeamName, player.MatchState)
	}

	// save player
	playerSQL, playerArgs, err := playerSB.Suffix("ON CONFLICT(steam_id) DO NOTHING").ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, playerSQL, playerArgs...); err != nil {
		return err
	}

	// save match_player
	matchPlayerSQL, matchPlayerArgs, err := matchPlayerSB.ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, matchPlayerSQL, matchPlayerArgs...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) SaveTeams(ctx context.Context, t dto.Teams) error {
	sql, args, err := r.builder.
		Insert("team").
		Columns("name, flag_code, create_time, update_time").
		Values(t.Team1Name, t.Team1Flag, t.CreateTime, t.CreateTime).
		Values(t.Team2Name, t.Team2Flag, t.CreateTime, t.CreateTime).
		Suffix("ON CONFLICT (name) DO NOTHING").
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) AddPlayersToTeams(ctx context.Context, players []dto.TeamPlayer) error {
	sb := r.builder.
		Insert("team_player").
		Columns("team_name, player_steam_id")

	for _, p := range players {
		sb = sb.Values(p.TeamName, p.SteamID)
	}

	sql, args, err := sb.Suffix("ON CONFLICT (team_name, player_steam_id) DO NOTHING").ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) SaveMatch(ctx context.Context, m *dto.Match) error {
	sql, args, err := r.builder.
		Insert("match").
		Columns("id, map_name, team1_name, team1_score, team2_name, team2_score, duration, upload_time").
		Values(m.ID, m.MapName, m.Team1.ClanName, m.Team1.Score, m.Team2.ClanName, m.Team2.Score, m.Duration, m.UploadTime).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return domain.ErrMatchAlreadyExist
			}
		}

		return err
	}

	return nil
}

func (r *replayRepo) UpsertStats(ctx context.Context, ps []dto.PlayerStat, ws []dto.WeaponStat) error {
	// player stats
	sb := r.builder.
		Insert("player_statistic AS s").
		Columns("id, player_steam_id, metric, value")

	for _, s := range ps {
		sb = sb.Values(s.ID, s.SteamID, s.Metric, s.Value)
	}

	sql, args, err := sb.
		Suffix("ON CONFLICT (id) DO UPDATE").
		Suffix("SET value = s.value + EXCLUDED.value").
		Suffix("WHERE s.id = EXCLUDED.id").
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	// weapon stats
	sb = r.builder.
		Insert("weapon_statistic AS s").
		Columns("id, player_steam_id, weapon_id, metric, value")

	for _, s := range ws {
		sb = sb.Values(s.ID, s.SteamID, s.WeaponID, s.Metric, s.Value)
	}

	sql, args, err = sb.
		Suffix("ON CONFLICT (id) DO UPDATE").
		Suffix("SET value = s.value + EXCLUDED.value").
		Suffix("WHERE s.id = EXCLUDED.id").
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
