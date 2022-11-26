package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ysomad/pgxatomic"
	"go.uber.org/zap"
)

type compendiumRepo struct {
	log     *zap.Logger
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
}

func NewCompendiumRepo(l *zap.Logger, p pgxatomic.Pool, b sq.StatementBuilderType) *compendiumRepo {
	return &compendiumRepo{
		log:     l,
		pool:    p,
		builder: b,
	}
}

func (r *compendiumRepo) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	sql, args, err := r.builder.
		Select("w.id, w.name, wc.id, wc.name").
		From("weapon w").
		InnerJoin("weapon_class wc ON w.class_id = wc.id").
		OrderBy("w.name").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	w, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Weapon])
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (r *compendiumRepo) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
	sql, args, err := r.builder.
		Select("id, name").
		From("weapon_class").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	c, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.WeaponClass])
	if err != nil {
		return nil, err
	}

	return c, nil
}
