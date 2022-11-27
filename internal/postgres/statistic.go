package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
	"github.com/ysomad/pgxatomic"
	"go.uber.org/zap"
)

type statisticRepo struct {
	log     *zap.Logger
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
}

func NewStatisticRepo(l *zap.Logger, p pgxatomic.Pool, b sq.StatementBuilderType) *statisticRepo {
	return &statisticRepo{
		log:     l,
		pool:    p,
		builder: b,
	}
}

func (r *statisticRepo) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]dto.WeaponStatWithClass, error) {
	sb := r.builder.
		Select("ws.weapon_id, w.weapon, wc.id, wc.class, ws.metric, ws.value").
		From("weapon_statistic ws").
		InnerJoin("weapon w ON w.id = ws.weapon_id").
		InnerJoin("weapon_class wc ON w.class_id = wc.id").
		Where(sq.Eq{"player_steam_id": steamID})

	switch {
	case f.WeaponID != 0:
		sb = sb.Where(sq.Eq{"ws.weapon_id": f.WeaponID})
	case f.WeaponClassID != 0:
		sb = sb.Where(sq.Eq{"wc.id": f.WeaponClassID})
	}

	sql, args, err := sb.OrderBy("w.weapon").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	m, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dto.WeaponStatWithClass])
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *statisticRepo) GetWeaponClassStats(ctx context.Context, steamID uint64, classID uint8) ([]dto.WeaponClassStat, error) {
	sb := r.builder.
		Select("wc.id, wc.class, ws.metric, SUM(ws.value)").
		From("weapon_statistic ws").
		InnerJoin("weapon w ON w.id = ws.weapon_id").
		InnerJoin("weapon_class wc ON w.class_id = wc.id").
		Where(sq.Eq{"player_steam_id": steamID})

	if classID != 0 {
		sb = sb.Where(sq.Eq{"weapon_class": classID})
	}

	sql, args, err := sb.
		GroupBy("wc.id, wc.class, ws.metric").
		OrderBy("wc.id").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	m, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dto.WeaponClassStat])
	if err != nil {
		return nil, err
	}

	return m, nil
}
