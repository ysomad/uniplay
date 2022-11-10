package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayRepository interface {
	InsertStats(context.Context, []dto.CreateMetricArgs, []dto.CreateWeaponMetricArgs) error
}

type matchRepository interface {
	Create(context.Context, *dto.CreateMatchArgs) (*domain.Match, error)
}
