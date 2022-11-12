package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayRepository interface {
	SaveMatch(context.Context, *domain.Match) error
	SaveStats(context.Context, []dto.CreateMetricArgs, []dto.CreateWeaponMetricArgs) error
}
