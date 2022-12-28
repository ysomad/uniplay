package postgres

import (
	"context"

	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
)

type compendiumRepo struct {
	log    *zap.Logger
	client *pgclient.Client
}

func NewCompendiumRepo(l *zap.Logger, c *pgclient.Client) *compendiumRepo {
	return &compendiumRepo{
		log:    l,
		client: c,
	}
}

func (r *compendiumRepo) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	sql, args, err := r.client.Builder.
		Select("w.id, w.weapon, wc.id, wc.class").
		From("weapon w").
		InnerJoin("weapon_class wc ON w.class_id = wc.id").
		OrderBy("w.weapon").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.client.Pool.Query(ctx, sql, args...)
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
	sql, args, err := r.client.Builder.
		Select("id, class").
		From("weapon_class").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	c, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.WeaponClass])
	if err != nil {
		return nil, err
	}

	return c, nil
}
