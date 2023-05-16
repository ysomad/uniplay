package player

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type playerRepository interface {
	GetAll(context.Context, listParams) (paging.List[domain.Player], error)
	FindBySteamID(context.Context, domain.SteamID) (domain.Player, error)
	UpdateBySteamID(context.Context, domain.SteamID, updateParams) (domain.Player, error)

	GetBaseStats(context.Context, domain.SteamID) (*domain.PlayerBaseStats, error)
	GetWeaponBaseStats(context.Context, domain.SteamID, domain.WeaponStatsFilter) ([]*domain.WeaponBaseStats, error)

	GetMatchList(context.Context, matchListParams) (paging.TokenList[*domain.PlayerMatch], error)
	GetMostPlayedMaps(context.Context, domain.SteamID) ([]domain.MostPlayedMap, error)
	GetMostSuccessMaps(context.Context, domain.SteamID) ([]domain.MostSuccessMap, error)
}

type service struct {
	tracer trace.Tracer
	repo   playerRepository
}

func NewService(t trace.Tracer, r playerRepository) *service {
	return &service{
		tracer: t,
		repo:   r,
	}
}

func (s *service) GetStats(ctx context.Context, steamID domain.SteamID) (domain.PlayerStats, error) {
	ctx, span := s.tracer.Start(ctx, "player.Service.GetStats")
	defer span.End()

	ts, err := s.repo.GetBaseStats(ctx, steamID)
	if err != nil {
		return domain.PlayerStats{}, err
	}

	return domain.NewPlayerStats(ts), nil
}

func (s *service) GetWeaponStats(ctx context.Context, steamID domain.SteamID, f domain.WeaponStatsFilter) ([]domain.WeaponStats, error) {
	ctx, span := s.tracer.Start(ctx, "player.Service.GetWeaponStats")
	defer span.End()

	ts, err := s.repo.GetWeaponBaseStats(ctx, steamID, f)
	if err != nil {
		return nil, err
	}

	return domain.NewWeaponStats(ts), nil
}
