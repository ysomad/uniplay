package player

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/otel"
)

type Service struct {
	player playerRepository
}

func NewService(r playerRepository) *Service {
	return &Service{
		player: r,
	}
}

func (s *Service) GetStats(ctx context.Context, steamID uint64) (domain.PlayerStats, error) {
	_, span := otel.StartTrace(ctx, libraryName, "player.Service.GetStats")
	defer span.End()

	ts, err := s.player.GetTotalStats(ctx, steamID)
	if err != nil {
		return domain.PlayerStats{}, err
	}

	return domain.NewPlayerStats(ts), nil
}

func (s *Service) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStat, error) {
	_, span := otel.StartTrace(ctx, libraryName, "player.Service.GetWeaponStats")
	defer span.End()

	ts, err := s.player.GetTotalWeaponStats(ctx, steamID, f)
	if err != nil {
		return nil, err
	}

	return domain.NewWeaponStats(ts), nil
}
