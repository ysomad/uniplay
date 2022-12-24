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

func (p *Player) Get(ctx context.Context, steamID uint64) (domain.Player, error) {
	// return p.repo.FindBySteamID(ctx, steamID)
	return domain.Player{}, nil
}
