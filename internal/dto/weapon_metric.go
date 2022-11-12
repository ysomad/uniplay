package dto

import (
	"github.com/ssssargsian/uniplay/internal/domain"
)

type CreateWeaponMetricArgs struct {
	MatchID       domain.MatchID
	PlayerSteamID uint64
	WeaponName    string
	WeaponClass   domain.EquipmentClass
	Metric        domain.Metric
	Value         int32
}
