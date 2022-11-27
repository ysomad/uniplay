package dto

import "github.com/ssssargsian/uniplay/internal/domain"

type WeaponStatWithClass struct {
	WeaponID uint16
	Weapon   string
	ClassID  uint8
	Class    string
	Metric   domain.Metric
	Value    uint32
}

type WeaponClassStat struct {
	ClassID uint8
	Class   string
	Metric  domain.Metric
	Value   uint32
}
