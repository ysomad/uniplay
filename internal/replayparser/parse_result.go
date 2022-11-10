package replayparser

import (
	"github.com/google/uuid"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type parseResult struct {
	metrics       *playerMetrics
	weaponMetrics *weaponMetrics
	match         *dto.CreateMatchArgs
}

func (r *parseResult) CreateMetricArgsList(matchID uuid.UUID) []dto.CreateMetricArgs {
	return r.metrics.toDTO(matchID)
}

func (r *parseResult) CreateWeaponArgsList(matchID uuid.UUID) []dto.CreateWeaponMetricArgs {
	return r.weaponMetrics.toDTO(matchID)
}

func (r *parseResult) Match() *dto.CreateMatchArgs {
	return r.match
}
