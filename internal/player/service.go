package player

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/uniplay/internal/domain"
)

type Service struct {
	tracer trace.Tracer
	player playerRepository
}

func NewService(t trace.Tracer, r playerRepository) *Service {
	return &Service{
		tracer: t,
		player: r,
	}
}

func (s *Service) GetStats(ctx context.Context, steamID uint64) (domain.PlayerStats, error) {
	ctx, span := s.tracer.Start(ctx, "player.Service.GetStats")
	defer span.End()

	ts, err := s.player.GetTotalStats(ctx, steamID)
	if err != nil {
		return domain.PlayerStats{}, err
	}

	return domain.NewPlayerStats(ts), nil
}

func (s *Service) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStat, error) {
	ctx, span := s.tracer.Start(ctx, "player.Service.GetWeaponStats")
	defer span.End()

	ts, err := s.player.GetTotalWeaponStats(ctx, steamID, f)
	if err != nil {
		return nil, err
	}

	return domain.NewWeaponStats(ts), nil
}
