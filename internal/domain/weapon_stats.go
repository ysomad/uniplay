package domain

import (
	"github.com/google/uuid"
	"github.com/ysomad/uniplay/internal/pkg/floatrounder"
)

type WeaponStats struct {
	Base     *WeaponBaseStats
	Accuracy WeaponAccuracyStats
}

func NewWeaponStats(total []*WeaponBaseStats) []WeaponStats {
	res := make([]WeaponStats, len(total))

	for i, s := range total {
		res[i] = WeaponStats{
			Base: s,
			Accuracy: newWeaponAccuracyStats(
				s.Shots,
				s.HeadHits,
				s.NeckHits,
				s.ChestHits,
				s.StomachHits,
				s.LeftArmHits,
				s.RightArmHits,
				s.LeftLegHits,
				s.RightLegHits,
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
	HeadHits          int32
	NeckHits          int32
	ChestHits         int32
	StomachHits       int32
	LeftArmHits       int32
	RightArmHits      int32
	LeftLegHits       int32
	RightLegHits      int32
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

// calcAccuracy returns accuracy in percentage.
func calcAccuracy(sum, num int32) float64 {
	if sum <= 0 || num <= 0 {
		return 0
	}

	return floatrounder.Round(float64(sum) * 100 / float64(num))
}

func newWeaponAccuracyStats(shots, headHits, neckHits, chestHits, stomachHits, lArmHits, rArmHits, lLegHits, rLegHits int32) WeaponAccuracyStats {
	hits := headHits + neckHits + chestHits + stomachHits + lArmHits + rArmHits + lLegHits + rLegHits

	if hits <= 0 {
		return WeaponAccuracyStats{}
	}

	return WeaponAccuracyStats{
		Total:   calcAccuracy(hits, shots),
		Head:    calcAccuracy(headHits, hits),
		Neck:    calcAccuracy(neckHits, hits),
		Chest:   calcAccuracy(chestHits, hits),
		Stomach: calcAccuracy(stomachHits, hits),
		Arms:    calcAccuracy(lArmHits+rArmHits, hits),
		Legs:    calcAccuracy(lLegHits+rLegHits, hits),
	}
}

type WeaponStatsFilter struct {
	WeaponID int16
	ClassID  int16
	MatchID  uuid.UUID
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
