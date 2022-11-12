package dto

import (
	"github.com/ssssargsian/uniplay/internal/domain"
)

type CreateMetricArgs struct {
	MatchID       domain.MatchID
	PlayerSteamID uint64
	Metric        domain.Metric
	Value         int32
}
