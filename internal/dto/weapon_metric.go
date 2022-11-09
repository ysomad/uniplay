package dto

import "github.com/google/uuid"

type CreateWeaponMetricArgs struct {
	MatchID       uuid.UUID
	PlayerSteamID uint64
	WeaponName    string
	WeaponClass   string
	Metric        int16
	Value         int32
}
