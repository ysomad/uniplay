package v1

import (
	"context"

	"go.uber.org/zap"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
	"github.com/ssssargsian/uniplay/internal/service"
)

var _ v1.ServerInterface = &handler{}

type atomicRunner interface {
	Run(context.Context, func(ctx context.Context) error) error
}

type handler struct {
	log        *zap.Logger
	atomic     atomicRunner
	replay     *service.Replay
	player     *service.Player
	statistic  *service.Statistic
	compendium *service.Compendium
}

func NewHandler(l *zap.Logger, a atomicRunner, r *service.Replay, p *service.Player, s *service.Statistic, c *service.Compendium) *handler {
	return &handler{
		log:        l,
		atomic:     a,
		replay:     r,
		player:     p,
		statistic:  s,
		compendium: c,
	}
}
