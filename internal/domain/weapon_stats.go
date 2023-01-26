package domain

import "math"

type WeaponStat struct {
	TotalStat    *WeaponTotalStat   `json:"total_stat"`
	AccuracyStat WeaponAccuracyStat `json:"accuracy_stat"`
}

func NewWeaponStats(total []*WeaponTotalStat) []WeaponStat {
	res := make([]WeaponStat, len(total))

	for i, s := range total {
		res[i] = WeaponStat{
			TotalStat: s,
			AccuracyStat: newWeaponAccuracyStat(
				s.Shots,
				s.HeadHits,
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

type WeaponTotalStat struct {
	WeaponID          int16  `json:"weapon_id"`
	Weapon            string `json:"weapon"`
	Kills             int32  `json:"kills"`
	HeadshotKills     int32  `json:"headshot_kills"`
	BlindKills        int32  `json:"blind_kills"`
	WallbangKills     int32  `json:"wallbang_kills"`
	NoScopeKills      int32  `json:"no_scope_kills"`
	ThroughSmokeKills int32  `json:"through_smoke_kills"`
	Deaths            int32  `json:"deaths"`
	Assists           int32  `json:"assists"`
	DamageTaken       int32  `json:"damage_taken"`
	DamageDealt       int32  `json:"damage_dealt"`
	Shots             int32  `json:"shots"`
	HeadHits          int32  `json:"head_hits"`
	ChestHits         int32  `json:"chest_hits"`
	StomachHits       int32  `json:"stomach_hits"`
	LeftArmHits       int32  `json:"left_arm_hits"`
	RightArmHits      int32  `json:"right_arm_hits"`
	LeftLegHits       int32  `json:"left_leg_hits"`
	RightLegHits      int32  `json:"right_leg_hits"`
}

type WeaponAccuracyStat struct {
	Total   float64 `json:"total"`
	Head    float64 `json:"head"`
	Chest   float64 `json:"chest"`
	Stomach float64 `json:"stomach"`
	Arms    float64 `json:"arms"`
	Legs    float64 `json:"legs"`
}

// round rounds float64 to 2 decimal places.
func round(n float64) float64 { return math.Round(n*100) / 100 }

// calcAccuracy returns accuracy in percentage.
func calcAccuracy(sum, num int32) float64 {
	if sum <= 0 || num <= 0 {
		return 0
	}
	return round(float64(sum) * 100 / float64(num))
}

func newWeaponAccuracyStat(shots, headHits, chestHits, stomachHits, lArmHits, rArmHits, lLegHits, rLegHits int32) WeaponAccuracyStat {
	hits := headHits + chestHits + stomachHits + lArmHits + rArmHits + lLegHits + rLegHits

	if hits <= 0 {
		return WeaponAccuracyStat{}
	}

	return WeaponAccuracyStat{
		Total:   calcAccuracy(hits, shots),
		Head:    calcAccuracy(headHits, hits),
		Chest:   calcAccuracy(chestHits, hits),
		Stomach: calcAccuracy(stomachHits, hits),
		Arms:    calcAccuracy(lArmHits+rArmHits, hits),
		Legs:    calcAccuracy(lLegHits+rLegHits, hits),
	}
}

type WeaponStatsFilter struct {
	WeaponID int16
	ClassID  int8
}
