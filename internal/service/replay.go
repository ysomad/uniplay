package service

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/replayparser"
)

type Replay struct {
	log  *zap.Logger
	repo replayRepository
}

func NewReplay(l *zap.Logger, r replayRepository) *Replay {
	return &Replay{
		log:  l,
		repo: r,
	}
}

func (r *Replay) CollectStats(ctx context.Context, filename string) (*domain.Match, error) {
	replay, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	p, err := replayparser.New(replay, r.log)
	if err != nil {
		return nil, err
	}
	defer p.Close()

	matchID, err := p.ParseMatchID()
	if err != nil {
		return nil, err
	}

	found, err := r.repo.MatchExists(ctx, matchID)
	if err != nil {
		return nil, err
	}

	if found {
		return nil, domain.ErrMatchAlreadyExist
	}

	res, err := p.Parse()
	if err != nil {
		return nil, err
	}

	playerStats, weaponStats := res.Stats()

	match, err := r.repo.SaveStats(ctx, res.Match(), playerStats, weaponStats)
	if err != nil {
		return nil, err
	}

	return match, nil
}
