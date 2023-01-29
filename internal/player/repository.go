package player

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type playerRepository interface {
	GetTotalStats(ctx context.Context, steamID uint64) (*domain.PlayerTotalStats, error)
	GetTotalWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]*domain.WeaponTotalStat, error)
}
