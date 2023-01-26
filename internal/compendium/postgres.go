package compendium

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
)

type pgStorage struct {
	log    *zap.Logger
	client *pgclient.Client
}

func NewPGStorage(l *zap.Logger, c *pgclient.Client) *pgStorage {
	return &pgStorage{
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

func (s *pgStorage) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
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

	w, err := pgx.CollectRows(rows, pgx.RowToStructByName[weapon])
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

func (s *pgStorage) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
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
