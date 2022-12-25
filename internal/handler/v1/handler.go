package v1

import (
	"go.uber.org/zap"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
	"github.com/ssssargsian/uniplay/internal/service"
)

var _ v1.ServerInterface = &handler{}

type handler struct {
	log        *zap.Logger
	replay     *service.Replay
	player     *service.Player
	statistic  *service.Statistic
	compendium *service.Compendium
}

func NewHandler(l *zap.Logger, r *service.Replay, p *service.Player, s *service.Statistic, c *service.Compendium) *handler {
	return &handler{
		log:        l,
		replay:     r,
		player:     p,
		statistic:  s,
		compendium: c,
	}
}
