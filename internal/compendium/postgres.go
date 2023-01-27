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

	weapons, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Weapon])
	if err != nil {
		return nil, err
	}

	return weapons, nil
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

	classes, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.WeaponClass])
	if err != nil {
		return nil, err
	}

	return classes, nil
}
