package v1

import (
	"context"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
	"github.com/ssssargsian/uniplay/internal/replay"
)

var _ v1.ServerInterface = &handler{}

type replayService interface {
	CollectStats(context.Context, replay.Replay) (*domain.Match, error)
}

type playerService interface {
	GetStats(ctx context.Context, steamID uint64) (domain.PlayerStats, error)
	GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStats, error)
}

type compendiumService interface {
	GetWeaponList(ctx context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error)
}

type handler struct {
	log        *zap.Logger
	replay     replayService
	player     playerService
	compendium compendiumService
}

func NewHandler(l *zap.Logger, r replayService, p playerService, c compendiumService) *handler {
	return &handler{
		log:        l,
		replay:     r,
		player:     p,
		compendium: c,
	}
}
