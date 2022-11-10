package service

import (
	"context"
	"io"
	"os"

	"github.com/google/uuid"

	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayParser interface {
	Parse() (interface {
		CreateMetricArgsList(matchID uuid.UUID) []dto.CreateMetricArgs
		CreateWeaponArgsList(matchID uuid.UUID) []dto.CreateWeaponMetricArgs
		CreateMatchArgs() *dto.CreateMatchArgs
	}, error)
	Close() error
}

type replayParserFactory func(io.Reader) replayParser

type replay struct {
	newReplayParser replayParserFactory
	replayRepo      replayRepository
	matchRepo       matchRepository
}

func NewReplay(rpf replayParserFactory, rp replayRepository, mr matchRepository) *replay {
	return &replay{
		newReplayParser: rpf,
		replayRepo:      rp,
		matchRepo:       mr,
	}
}

func (r *replay) CollectStats(ctx context.Context, replay io.Reader) error {
	demo, err := os.Open("./test-data/1.dem")
	if err != nil {
		return err
	}
	defer demo.Close()

	rp := r.newReplayParser(demo)
	defer rp.Close()

	res, err := rp.Parse()
	if err != nil {
		return err
	}

	match, err := r.matchRepo.Create(ctx, res.CreateMatchArgs())
	if err != nil {
		return err
	}

	err = r.replayRepo.InsertStats(ctx, res.CreateMetricArgsList(match.ID), res.CreateWeaponArgsList(match.ID))
	if err != nil {
		return err
	}

	return nil
}
