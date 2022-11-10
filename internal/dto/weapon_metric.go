package dto

import (
	"github.com/google/uuid"
	"github.com/ssssargsian/uniplay/internal/domain"
)

type CreateWeaponMetricArgs struct {
	MatchID       uuid.UUID
	PlayerSteamID uint64
	WeaponName    string
	WeaponClass   domain.EquipmentClass
	Metric        domain.Metric
	Value         int32
}
