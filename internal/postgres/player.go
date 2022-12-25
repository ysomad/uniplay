package postgres

import (
	"context"
	"errors"

	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
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
			"sum(ps.kills) as kills",
			"sum(ps.hs_kills) as hs_kills",
			"sum(ps.blind_kills) as blind_kills",
			"sum(ps.wallbang_kills) as wallbang_kills",
			"sum(ps.noscope_kills) as noscope_kills",
			"sum(ps.through_smoke_kills) as through_smoke_kills",
			"sum(ps.deaths) as deaths",
			"sum(ps.assists) as assists",
			"sum(ps.flashbang_assists) as flashband_assists",
			"sum(ps.mvp_count) as mvp_count",
			"sum(ps.damage_taken) as damage_taken",
			"sum(ps.damage_dealt) as damage_dealt",
			"sum(ps.grenade_damage_dealt) as grenade_damage_dealt",
			"sum(ps.blinded_players) as blinded_players",
			"sum(ps.blinded_times) as blinded_times",
			"sum(ps.bombs_planted) as bombs_planted",
			"sum(ps.bombs_defused) as bombs_defused",
			"sum(m.team1_score) + sum(m.team2_score) as rounds_played",
			"count(m.id) as matches_played",
			"coalesce((case when pm.match_state = 1 then count(pm.*) end), 0) as wins",
			"coalesce((case when pm.match_state = -1 then count(pm.*) end), 0) as loses",
			"coalesce((case when pm.match_state = 0 then count(pm.*) end), 0) as draws",
			"sum(m.duration) as time_played").
		From("player_match_stat ps").
		LeftJoin("player_match pm ON ps.player_steam_id = pm.player_steam_id").
		LeftJoin("match m ON pm.match_id = m.id").
		Where(sq.Eq{"ps.player_steam_id": steamID}).
		GroupBy("pm.match_state").
		ToSql()
	if err != nil {
		return nil, err
	}

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
