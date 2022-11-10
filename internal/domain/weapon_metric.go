package domain

import "github.com/google/uuid"

type WeaponMetric struct {
	MatchID       uuid.UUID
	PlayerSteamID uint64
	Metric        Metric
	WeaponName    string
	WeaponClass   EquipmentClass
	Value         int32
}
