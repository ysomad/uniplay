package player

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type service struct {
	player playerRepository
}

func NewService(r playerRepository) *service {
	return &service{
		player: r,
	}
}

func (s *service) GetStats(ctx context.Context, steamID uint64) (domain.PlayerStats, error) {
	ts, err := s.player.GetTotalStats(ctx, steamID)
	if err != nil {
		return domain.PlayerStats{}, err
	}

	return domain.NewPlayerStats(ts), nil
}

func (s *service) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStat, error) {
	ts, err := s.player.GetTotalWeaponStats(ctx, steamID, f)
	if err != nil {
		return nil, err
	}

	return domain.NewWeaponStats(ts), nil
}
