package v1

import (
	"context"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
	"github.com/ssssargsian/uniplay/internal/service"
)

var _ v1.ServerInterface = &handler{}

type replayService interface {
	SaveReplay(ctx context.Context) error // TODO: IMPLEMENT WITH FILE STORAGE
	CollectStats(ctx context.Context, filename string) (*domain.Match, error)
}

type handler struct {
	log        *zap.Logger
	replay     replayService
	player     *service.Player
	compendium *service.Compendium
}

func NewHandler(l *zap.Logger, r replayService, p *service.Player, c *service.Compendium) *handler {
	return &handler{
		log:        l,
		replay:     r,
		player:     p,
		compendium: c,
	}
}
