package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

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

type weapon struct {
	WeaponID int16  `db:"weapon_id"`
	Weapon   string `db:"weapon"`
	ClassID  int8   `db:"class_id"`
	Class    string `db:"class"`
}

func (r *compendiumRepo) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	sql, args, err := r.client.Builder.
		Select("w.id as weapon_id, w.weapon, wc.id as class_id, wc.class").
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

	w, err := pgx.CollectRows(rows, pgx.RowToStructByPos[weapon])
	if err != nil {
		return nil, err
	}

	res := make([]domain.Weapon, len(w))
	for i, v := range w {
		res[i] = domain.Weapon(v)
	}

	return res, nil
}

type weaponClass struct {
	ID    int8   `db:"id"`
	Class string `db:"class"`
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

	c, err := pgx.CollectRows(rows, pgx.RowToStructByName[weaponClass])
	if err != nil {
		return nil, err
	}

	res := make([]domain.WeaponClass, len(c))
	for i, v := range c {
		res[i] = domain.WeaponClass(v)
	}

	return res, nil
}
