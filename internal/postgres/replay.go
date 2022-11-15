package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ysomad/pgxatomic"

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

func (r *replayRepo) SavePlayers(ctx context.Context, players dto.PlayerSteamIDs) error {
	sb := r.builder.
		Insert("player").
		Columns("steam_id, create_time, update_time")

	for _, steamID := range players.SteamIDs {
		sb = sb.Values(steamID, players.CreateTime, players.CreateTime)
	}

	sql, args, err := sb.Suffix("ON CONFLICT(steam_id) DO NOTHING").ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
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
	// TODO: предовратить добавление игрока в одну и ту же команду

	sb := r.builder.
		Insert("team_player").
		Columns("team_name, player_steam_id")

	for _, p := range players {
		sb = sb.Values(p.TeamName, p.PlayerSteamID)
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
		return err
	}

	return nil
}

func (r *replayRepo) SaveMetrics(ctx context.Context, metrics []dto.Metric, wmetrics []dto.WeaponMetric) error {
	// metrics
	sb := r.builder.
		Insert("metric").
		Columns("match_id, player_steam_id, metric, value")

	for _, m := range metrics {
		sb = sb.Values(m.MatchID, m.PlayerSteamID, m.Metric, m.Value)
	}

	sql, args, err := sb.ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	// weapon metrics
	sb = r.builder.
		Insert("weapon_metric").
		Columns("match_id, player_steam_id, weapon_name, weapon_class, metric, value")

	for _, wm := range wmetrics {
		sb = sb.Values(wm.MatchID, wm.PlayerSteamID, wm.WeaponName, wm.WeaponClass, wm.Metric, wm.Value)
	}

	sql, args, err = sb.ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
