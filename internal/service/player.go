package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type Player struct {
	repo playerRepository
}

func NewPlayer(r playerRepository) *Player {
	return &Player{
		repo: r,
	}
}

func (p *Player) GetStats(ctx context.Context, steamID uint64) (domain.PlayerStats, error) {
	total, err := p.repo.GetTotalStats(ctx, steamID)
	if err != nil {
		return domain.PlayerStats{}, err
	}

	return domain.NewPlayerStats(total), nil
}

func (p *Player) GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStats, error) {
	total, err := p.repo.GetTotalWeaponStats(ctx, steamID, f)
	if err != nil {
		return nil, err
	}

	return domain.NewWeaponStatsList(total), nil
}
