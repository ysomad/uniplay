package service

import (
	"context"
	"io"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/replayparser"
)

type replay struct {
	repo replayRepository
}

func NewReplay(r replayRepository) *replay {
	return &replay{
		repo: r,
	}
}

func (r *replay) CollectStats(ctx context.Context, replay io.Reader) (*domain.Match, error) {
	p := replayparser.New(replay)
	defer p.Close()

	res, err := p.Parse()
	if err != nil {
		return nil, err
	}

	a := res.CreateMatchArgs()
	m, err := domain.NewMatch(a.MapName, a.Duration, a.Team1, a.Team2)
	if err != nil {
		return nil, err
	}

	if err = r.repo.SaveMatch(ctx, m); err != nil {
		return nil, err
	}

	if err = r.repo.SaveStats(ctx, res.CreateMetricsArgsList(m.ID), res.CreateWeaponMetricArgsList(m.ID)); err != nil {
		return nil, err
	}

	return m, nil
}
