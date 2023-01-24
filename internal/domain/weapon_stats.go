package domain

import "math"

type WeaponStat struct {
	TotalStat    WeaponTotalStat    `json:"total_stat"`
	AccuracyStat WeaponAccuracyStat `json:"accuracy_stat"`
}

func NewWeaponStats(total []WeaponTotalStat) []WeaponStat {
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
func calcAccuracy(a, b float64) float64 {
	if a <= 0 || b <= 0 {
		return 0
	}
	return round(a * 100 / b)
}

func newWeaponAccuracyStat(shots, headHits, chestHits, stomachHits, lArmHits, rArmHits, lLegHits, rLegHits int32) WeaponAccuracyStat {
	hits := float64(headHits + chestHits + stomachHits + lArmHits + rArmHits + lLegHits + rLegHits)

	if hits <= 0 {
		return WeaponAccuracyStat{}
	}

	return WeaponAccuracyStat{
		Total:   calcAccuracy(hits, float64(shots)),
		Head:    calcAccuracy(float64(headHits), hits),
		Chest:   calcAccuracy(float64(chestHits), hits),
		Stomach: calcAccuracy(float64(stomachHits), hits),
		Arms:    calcAccuracy(float64(lArmHits+rArmHits), hits),
		Legs:    calcAccuracy(float64(lLegHits+rLegHits), hits),
	}
}

type WeaponStatsFilter struct {
	WeaponID int16
	ClassID  int8
}
