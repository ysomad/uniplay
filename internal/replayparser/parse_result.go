package replayparser

import (
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type parseResult struct {
	metrics       *playerMetrics
	weaponMetrics *weaponMetrics
	match         *dto.CreateMatchArgs
}

func (r *parseResult) CreateMetricsArgsList(matchID domain.MatchID) []dto.CreateMetricArgs {
	return r.metrics.toDTO(matchID)
}

func (r *parseResult) CreateWeaponMetricArgsList(matchID domain.MatchID) []dto.CreateWeaponMetricArgs {
	return r.weaponMetrics.toDTO(matchID)
}

func (r *parseResult) CreateMatchArgs() *dto.CreateMatchArgs {
	return r.match
}
