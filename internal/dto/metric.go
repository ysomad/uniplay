package dto

import (
	"github.com/google/uuid"
	"github.com/ssssargsian/uniplay/internal/domain"
)

type CreateMetricArgs struct {
	MatchID       uuid.UUID
	PlayerSteamID uint64
	Metric        domain.Metric
	Value         int32
}
