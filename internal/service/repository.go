package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type playerRepository interface {
	InsertReplayStats(context.Context, dto.CreateMetricArgs, dto.CreateWeaponMetricArgs, domain.Match) error
}
