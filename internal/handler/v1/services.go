package v1

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayService interface {
	CollectStats(ctx context.Context, filename string) (*dto.Match, error)
}

type playerService interface {
	Get(ctx context.Context, steamID uint64) (domain.Player, error)
}

type statisticService interface {
	GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStats, error)
	GetWeaponClassStats(ctx context.Context, steamID uint64, classID uint8) ([]domain.WeaponClassStats, error)
}

type compendiumService interface {
	GetWeaponList(context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(context.Context) ([]domain.WeaponClass, error)
}
