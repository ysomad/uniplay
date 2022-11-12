package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayRepository interface {
	SaveStats(context.Context, []dto.CreateMetricArgs, []dto.CreateWeaponMetricArgs) error
}

type matchRepository interface {
	Save(context.Context, *domain.Match) error
}
