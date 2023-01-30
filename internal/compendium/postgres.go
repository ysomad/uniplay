package compendium

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

type Postgres struct {
	client *pgclient.Client
}

func NewPostgres(c *pgclient.Client) *Postgres {
	return &Postgres{
		client: c,
	}
}

func (s *Postgres) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	sql, args, err := s.client.Builder.
		Select("w.id as weapon_id, w.weapon, wc.id as class_id, wc.class").
		From("weapon w").
		InnerJoin("weapon_class wc ON w.class_id = wc.id").
		OrderBy("w.weapon").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	weapons, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Weapon])
	if err != nil {
		return nil, err
	}

	return weapons, nil
}

func (s *Postgres) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
	sql, args, err := s.client.Builder.
		Select("id, class").
		From("weapon_class").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	classes, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.WeaponClass])
	if err != nil {
		return nil, err
	}

	return classes, nil
}
