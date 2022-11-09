package dto

import "github.com/google/uuid"

type CreateMetricArgs struct {
	MatchID       uuid.UUID
	PlayerSteamID uint64
	Metric        int16
	Value         int32
}
