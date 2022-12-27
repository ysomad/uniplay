package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayRepository interface {
	SaveStats(context.Context, *dto.ReplayMatch, []dto.PlayerStat, []dto.PlayerWeaponStat) (*domain.Match, error)
	MatchExists(ctx context.Context, matchID uuid.UUID) (found bool, err error)
}

type playerRepository interface {
	GetTotalStats(ctx context.Context, steamID uint64) (domain.PlayerTotalStats, error)
	GetTotalWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponTotalStats, error)
}

type compendiumRepository interface {
	GetWeaponList(context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(context.Context) ([]domain.WeaponClass, error)
}
