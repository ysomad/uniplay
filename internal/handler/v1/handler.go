package v1

import (
	"context"

	"go.uber.org/zap"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

var _ v1.ServerInterface = &handler{}

type atomicRunner interface {
	Run(context.Context, func(ctx context.Context) error) error
}

type handler struct {
	log        *zap.Logger
	atomic     atomicRunner
	replay     replayService
	player     playerService
	statistic  statisticService
	compendium compendiumService
}

func NewHandler(l *zap.Logger, a atomicRunner, r replayService, p playerService, s statisticService, c compendiumService) *handler {
	return &handler{
		log:        l,
		atomic:     a,
		replay:     r,
		player:     p,
		statistic:  s,
		compendium: c,
	}
}
