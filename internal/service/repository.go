package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type playerRepository interface {
	GetTotalStats(ctx context.Context, steamID uint64) (*domain.PlayerTotalStats, error)
	GetTotalWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponTotalStats, error)
}

type compendiumRepository interface {
	GetWeaponList(context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(context.Context) ([]domain.WeaponClass, error)
}
