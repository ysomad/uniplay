package compendium

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

type postgres struct {
	client *pgclient.Client
}

func NewPostgres(c *pgclient.Client) *postgres {
	return &postgres{
		client: c,
	}
}

func (p *postgres) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	sql, args, err := p.client.Builder.
		Select("w.id as weapon_id, w.weapon, wc.id as class_id, wc.class").
		From("weapon w").
		InnerJoin("weapon_class wc ON w.class_id = wc.id").
		OrderBy("w.weapon").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Weapon])
}

func (p *postgres) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
	sql, args, err := p.client.Builder.
		Select("id, class").
		From("weapon_class").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[domain.WeaponClass])
}

func (p *postgres) GetMapList(ctx context.Context) ([]domain.Map, error) {
	sql, args, err := p.client.Builder.
		Select("name, icon_url").
		From("map").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Map])
}
