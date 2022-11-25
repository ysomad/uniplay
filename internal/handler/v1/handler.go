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

type statisticService interface {
	GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) (domain.WeaponStats, error)
	GetWeaponClassStats(ctx context.Context, steamID uint64, c domain.WeaponClassID) (domain.WeaponClassStats, error)
}

type compendiumService interface {
	GetWeaponList(context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(context.Context) ([]domain.WeaponClass, error)
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
