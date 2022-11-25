package dto

import "github.com/ssssargsian/uniplay/internal/domain"

type WeaponMetricSum struct {
	WeaponName string
	Metric     domain.Metric
	Value      uint32
}

type WeaponClassMetricSum struct {
	WeaponClass domain.WeaponClass
	Metric      domain.Metric
	Value       uint32
}
