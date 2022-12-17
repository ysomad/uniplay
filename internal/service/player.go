package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type player struct {
	repo playerRepository
}

func NewPlayer(r playerRepository) *player {
	return &player{
		repo: r,
	}
}

func (p *player) Get(ctx context.Context, steamID uint64) (domain.Player, error) {
	return p.repo.FindBySteamID(ctx, steamID)
}
