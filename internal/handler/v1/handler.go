package v1

import (
	"context"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

var _ v1.ServerInterface = &handler{}

type atomicRunner interface {
	Run(context.Context, func(ctx context.Context) error) error
}

type replayService interface {
	CollectStats(ctx context.Context, filename string) (*dto.Match, error)
}

type playerService interface {
	Get(ctx context.Context, steamID uint64) (domain.Player, error)
}

// TODO: IMPLEMENT
type statisticService interface {
	GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) (domain.WeaponStats, error)
	GetWeaponClassStats(ctx context.Context, steamID uint64) (domain.WeaponClassStats, error)
}

type handler struct {
	log       *zap.Logger
	atomic    atomicRunner
	replay    replayService
	player    playerService
	statistic statisticService
}

func NewHandler(l *zap.Logger, a atomicRunner, r replayService, p playerService, s statisticService) *handler {
	return &handler{
		log:       l,
		atomic:    a,
		replay:    r,
		player:    p,
		statistic: s,
	}
}
