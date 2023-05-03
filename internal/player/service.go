package player

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/uniplay/internal/domain"
)

type repository interface {
	FindBySteamID(context.Context, domain.SteamID) (domain.Player, error)
	UpdateBySteamID(context.Context, domain.SteamID, updateParams) (domain.Player, error)
	GetBaseStats(ctx context.Context, steamID uint64) (*domain.PlayerBaseStats, error)
	GetWeaponBaseStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]*domain.WeaponBaseStats, error)
}

type service struct {
	tracer trace.Tracer
	player repository
}

func NewService(t trace.Tracer, r repository) *service {
	return &service{
		tracer: t,
		player: r,
	}
}

func (s *service) GetBySteamID(ctx context.Context, steamID domain.SteamID) (domain.Player, error) {
	return s.player.FindBySteamID(ctx, steamID)
}

func (s *service) UpdateBySteamID(ctx context.Context, steamID domain.SteamID, p updateParams) (domain.Player, error) {
	return s.player.UpdateBySteamID(ctx, steamID, p)
}

func (s *service) GetStats(ctx context.Context, steamID uint64) (domain.PlayerStats, error) {
	ctx, span := s.tracer.Start(ctx, "player.Service.GetStats")
	defer span.End()

	ts, err := s.player.GetBaseStats(ctx, steamID)
	if err != nil {
		return domain.PlayerStats{}, err
	}

	return domain.NewPlayerStats(ts), nil
}

func (s *service) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStats, error) {
	ctx, span := s.tracer.Start(ctx, "player.Service.GetWeaponStats")
	defer span.End()

	ts, err := s.player.GetWeaponBaseStats(ctx, steamID, f)
	if err != nil {
		return nil, err
	}

	return domain.NewWeaponStats(ts), nil
}
