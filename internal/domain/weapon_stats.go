package domain

import "github.com/ysomad/uniplay/internal/pkg/stat"

type WeaponStats struct {
	Base     *WeaponBaseStats
	Accuracy *WeaponAccuracyStats
}

func NewWeaponStats(b []*WeaponBaseStats) []WeaponStats {
	res := make([]WeaponStats, len(b))

	for i, s := range b {
		res[i] = WeaponStats{
			Base: s,
			Accuracy: NewWeaponAccuracyStats(
				s.Shots,
				s.TotalHits,
				s.HeadHits,
				s.NeckHits,
				s.ChestHits,
				s.StomachHits,
				s.ArmHits,
				s.LegHits,
			),
		}
	}

	return res
}

type WeaponBaseStats struct {
	WeaponID          int16
	Weapon            string
	Kills             int32
	HeadshotKills     int32
	BlindKills        int32
	WallbangKills     int32
	NoScopeKills      int32
	ThroughSmokeKills int32
	Deaths            int32
	Assists           int32
	DamageTaken       int32
	DamageDealt       int32
	Shots             int32
	TotalHits         int32
	HeadHits          int32
	NeckHits          int32
	ChestHits         int32
	StomachHits       int32
	ArmHits           int32
	LegHits           int32
}

type WeaponAccuracyStats struct {
	Total   float64
	Head    float64
	Neck    float64
	Chest   float64
	Stomach float64
	Arms    float64
	Legs    float64
}

func NewWeaponAccuracyStats(shots, hits, headHits, neckHits, chestHits, stomachHits, armHits, legHits int32) *WeaponAccuracyStats {
	return &WeaponAccuracyStats{
		Total:   stat.Accuracy(hits, shots),
		Head:    stat.Accuracy(headHits, hits),
		Neck:    stat.Accuracy(neckHits, hits),
		Chest:   stat.Accuracy(chestHits, hits),
		Stomach: stat.Accuracy(stomachHits, hits),
		Arms:    stat.Accuracy(armHits, hits),
		Legs:    stat.Accuracy(legHits, hits),
	}
}

type WeaponStatsFilter struct {
	WeaponID int16
	ClassID  int16
}

func NewWeaponStatsFilter(weaponID, classID *int16) WeaponStatsFilter {
	f := WeaponStatsFilter{}

	if weaponID != nil {
		f.WeaponID = *weaponID
	}

	if classID != nil {
		f.ClassID = *classID
	}

	return f
}
