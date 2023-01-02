package postgres

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
)

type playerRepo struct {
	log    *zap.Logger
	client *pgclient.Client
}

func NewPlayerRepo(l *zap.Logger, c *pgclient.Client) *playerRepo {
	return &playerRepo{
		log:    l,
		client: c,
	}
}

func (r *playerRepo) GetTotalStats(ctx context.Context, steamID uint64) (*domain.PlayerTotalStats, error) {
	sql, args, err := r.client.Builder.
		Select(
			"sum(ps.kills) as total_kills",
			"sum(ps.hs_kills) as total_hs_kills",
			"sum(ps.blind_kills) as total_blind_kills",
			"sum(ps.wallbang_kills) as total_wb_kills",
			"sum(ps.noscope_kills) as total_noscope_kills",
			"sum(ps.through_smoke_kills) as total_smoke_kills",
			"sum(ps.deaths) as total_deaths",
			"sum(ps.assists) as total_assists",
			"sum(ps.flashbang_assists) as total_fb_assists",
			"sum(ps.mvp_count) as total_mvp_count",
			"sum(ps.damage_taken) as total_dmg_taken",
			"sum(ps.damage_dealt) as total_dmg_dealt",
			"sum(ps.grenade_damage_dealt) as total_grenade_dmg_dealt",
			"sum(ps.blinded_players) as total_blinded_players",
			"sum(ps.blinded_times) as total_blinded_times",
			"sum(ps.bombs_planted) as total_bombs_planted",
			"sum(ps.bombs_defused) as total_bombs_defused",
			"sum(m.team1_score) + sum(m.team2_score) as total_rounds_played",
			"count(m.id) as total_matches_played",
			"coalesce((case when pm.match_state = 1 then count(pm.*) end), 0) as total_wins",
			"coalesce((case when pm.match_state = -1 then count(pm.*) end), 0) as total_loses",
			"coalesce((case when pm.match_state = 0 then count(pm.*) end), 0) as total_draws",
			"sum(m.duration) as total_time_played").
		From("player_match_stat ps").
		InnerJoin("player_match pm ON ps.player_steam_id = pm.player_steam_id").
		InnerJoin("match m ON pm.match_id = m.id").
		Where(sq.Eq{"ps.player_steam_id": steamID}).
		GroupBy("pm.match_state").
		ToSql()
	if err != nil {
		return nil, err
	}

	r.log.Debug("playerRepo", zap.String("query", sql))

	rows, err := r.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	stats, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[domain.PlayerTotalStats])
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrPlayerNotFound
		}

		return nil, err
	}

	return stats, nil
}

func (r *playerRepo) GetTotalWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponTotalStats, error) {
	b := r.client.Builder.
		Select(
			"ws.weapon_id",
			"w.weapon",
			"sum(ws.kills) as total_kills",
			"sum(ws.hs_kills) as total_hs_kills",
			"sum(ws.blind_kills) as total_blind_kills",
			"sum(ws.wallbang_kills) as total_wb_kills",
			"sum(ws.noscope_kills) as total_noscope_kills",
			"sum(ws.through_smoke_kills) as total_smoke_kills",
			"sum(ws.deaths) as total_deaths",
			"sum(ws.assists) as total_assists",
			"sum(ws.damage_taken) as total_dmg_taken",
			"sum(ws.damage_dealt) as total_dmg_dealt",
			"sum(ws.shots) as total_shots",
			"sum(ws.head_hits) as total_head_hits",
			"sum(ws.chest_hits) as total_chest_hits",
			"sum(ws.stomach_hits) as total_stomach_hits",
			"sum(ws.left_arm_hits) as total_l_arm_hits",
			"sum(ws.right_arm_hits) as total_r_arm_hits",
			"sum(ws.left_leg_hits) as total_l_leg_hits",
			"sum(ws.right_leg_hits) as total_r_leg_hits").
		From("player_match_weapon_stat ws").
		InnerJoin("weapon w ON ws.weapon_id = w.id").
		Where(sq.Eq{"ws.player_steam_id": steamID})

	switch {
	case f.WeaponID > 0:
		b = b.Where(sq.Eq{"ws.weapon_id": f.WeaponID})
	case f.ClassID > 0:
		b = b.Where(sq.Eq{"w.class_id": f.ClassID})
	}

	sql, args, err := b.
		GroupBy("ws.weapon_id", "w.weapon").
		OrderBy("total_kills DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	r.log.Debug("playerRepo", zap.String("query", sql))

	rows, err := r.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	weaponStats, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.WeaponTotalStats])
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrPlayerNotFound
		}

		return nil, err
	}

	if len(weaponStats) <= 0 {
		return nil, domain.ErrPlayerNotFound
	}

	return weaponStats, nil
}
