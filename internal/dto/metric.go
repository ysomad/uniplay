package dto

import "github.com/ssssargsian/uniplay/internal/domain"

type WeaponMetricSum struct {
	WeaponID uint16
	Weapon   string
	ClassID  uint8
	Class    string
	Metric   domain.Metric
	Value    uint32
}

type WeaponClassMetricSum struct {
	WeaponClassID uint8
	Metric        domain.Metric
	Value         uint32
}
