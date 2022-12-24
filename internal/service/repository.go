package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayRepository interface {
	SaveStats(context.Context, *dto.ReplayMatch, []*dto.PlayerStat, []*dto.PlayerWeaponStat) (*domain.Match, error)
	MatchExists(ctx context.Context, matchID uuid.UUID) (found bool, err error)
}

type playerRepository interface {
	// FindBySteamID(ctx context.Context, steamID uint64) (domain.Player, error)
}

type statisticRepository interface {
	// GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]dto.StorageWeaponStat, error)
	// GetWeaponClassStats(ctx context.Context, steamID uint64, classID uint8) ([]dto.WeaponClassStat, error)
}

type compendiumRepository interface {
	// GetWeaponList(context.Context) ([]domain.Weapon, error)
	// GetWeaponClassList(context.Context) ([]domain.WeaponClass, error)
}
