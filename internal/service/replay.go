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

	/*
		1. сделать match id из мета информации из демки, например "карта-длительность-команда1-счет1-команда2-счет2" в base64
		2. создать матч, если создался, то создавать стату, если нет - вернуть ошибку что эта демка уже была загружена ранее.
	*/

	a := res.CreateMatchArgs()
	m, err := domain.NewMatch(a.MapName, a.Duration, a.Team1, a.Team2)
	if err != nil {
		return nil, err
	}

	if err = r.repo.SaveMatch(ctx, m); err != nil {
		return nil, err
	}

	// if err = r.replayRepo.SaveStats(ctx, res.CreateMetricArgsList(m.ID), res.CreateWeaponArgsList(m.ID)); err != nil {
	// 	return nil, err
	// }

	return m, nil
}
