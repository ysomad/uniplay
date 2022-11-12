package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
	"github.com/ysomad/pgxatomic"
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

func (r *replayRepo) SaveMatch(ctx context.Context, m *domain.Match) error {
	// insert teams
	sql, args, err := r.builder.
		Insert("team").
		Columns("name, flag_code, create_time, update_time").
		Values(m.Team1.Name, m.Team1.FlagCode, m.UploadTime, m.UploadTime).
		Values(m.Team2.Name, m.Team2.FlagCode, m.UploadTime, m.UploadTime).
		Suffix("ON CONFLICT DO NOTHING").
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	// insert players
	sb := r.builder.
		Insert("player").
		Columns("steam_id", "create_time", "update_time")

	type teamPlayer struct {
		teamName string
		steamID  uint64
	}

	team1len := len(m.Team1.PlayerSteamIDs)
	teamPlayers := make([]teamPlayer, team1len+len(m.Team2.PlayerSteamIDs))
	for i, steamID := range m.Team1.PlayerSteamIDs {
		teamPlayers[i] = teamPlayer{
			teamName: m.Team1.Name,
			steamID:  steamID,
		}
	}
	for i, steamID := range m.Team2.PlayerSteamIDs {
		teamPlayers[i+team1len] = teamPlayer{
			teamName: m.Team2.Name,
			steamID:  steamID,
		}
	}

	for _, p := range teamPlayers {
		sb = sb.Values(p.steamID, m.UploadTime, m.UploadTime)
	}

	sql, args, err = sb.Suffix("ON CONFLICT DO NOTHING").ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	// insert team players
	sb = r.builder.Insert("team_player").Columns("team_name, player_steam_id, is_active")

	for _, p := range teamPlayers {
		sb = sb.Values(p.teamName, p.steamID, true)
	}

	sql, args, err = sb.Suffix("ON CONFLICT DO NOTHING").ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	// insert match
	sql, args, err = r.builder.
		Insert("match").
		Columns("id, map_name, team1_name, team1_score, team2_name, team2_score, duration, upload_time").
		Values(m.ID, m.MapName, m.Team1.Name, m.Team1.Score, m.Team2.Name, m.Team2.Score, m.Duration, m.UploadTime).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *replayRepo) SaveStats(ctx context.Context, metrics []dto.CreateMetricArgs, wmetrics []dto.CreateWeaponMetricArgs) error {
	return nil
}
