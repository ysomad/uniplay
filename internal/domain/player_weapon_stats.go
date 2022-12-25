package domain

type PlayerWeaponStats struct {
	TotalStats    *PlayerWeaponTotalStats    `json:"total_stats"`
	AccuracyStats *PlayerWeaponAccuracyStats `json:"accuracy_stats"`
}

type PlayerWeaponTotalStats struct {
	WeaponID          int16  `json:"weapon_id"`
	Weapon            string `json:"weapon"`
	Kills             int32  `json:"kills"`
	HeadshotKills     int16  `json:"headshot_kills"`
	BlindKills        int16  `json:"blind_kills"`
	WallbangKills     int16  `json:"wallbang_kills"`
	NoScopeKills      int16  `json:"no_scope_kills"`
	ThroughSmokeKills int16  `json:"through_smoke_kills"`
	Deaths            int32  `json:"deaths"`
	Assists           int16  `json:"assists"`
	DamageTaken       int32  `json:"damage_taken"`
	DamageDealt       int32  `json:"damage_dealt"`
	Shots             int32  `json:"shots"`
	Hits              int32  `json:"hits"`
	HeadHits          int16  `json:"head_hits"`
	ChestHits         int16  `json:"chest_hits"`
	StomachHits       int16  `json:"stomach_hits"`
	LeftArmHits       int16  `json:"left_arm_hits"`
	RightArmHits      int16  `json:"right_arm_hits"`
	LeftLegHits       int16  `json:"left_leg_hits"`
	RightLegHits      int16  `json:"right_leg_hits"`
}

type PlayerWeaponAccuracyStats struct {
	Total   float64 `json:"total"`
	Head    float64 `json:"head"`
	Chest   float64 `json:"chest"`
	Stomach float64 `json:"stomach"`
	Arms    float64 `json:"arms"`
	Legs    float64 `json:"legs"`
}

func NewPlayerWeaponAccuracyStats(
	shots,
	hits,
	headHits,
	chestHits,
	stomachHits,
	lArmHits,
	rArmHits,
	lLegHits,
	rLegHits float64,
) PlayerWeaponAccuracyStats {
	return PlayerWeaponAccuracyStats{
		Total:   shots / hits,
		Head:    shots / headHits,
		Chest:   shots / chestHits,
		Stomach: shots / stomachHits,
		Arms:    shots / (lArmHits + rArmHits),
		Legs:    shots / (lLegHits + rLegHits),
	}
}
