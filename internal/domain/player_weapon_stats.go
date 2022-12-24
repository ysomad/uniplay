package domain

type PlayerWeaponStats struct {
	WeaponID          int16
	Weapon            string
	Kills             int32
	HSKills           int16
	BlindKills        int16
	WallbangKills     int16
	NoScopeKills      int16
	ThroughSmokeKills int16
	Deaths            int32
	Assists           int16
	DamageTaken       int32
	DamageDealt       int32
	Accuracy          float64
	Shots             int32
	Hits              int32
	HeadHits          int32
	ChestHits         int32
	StomachHits       int32
	LeftArmHits       int32
	RightArmHits      int32
	LeftLegHits       int32
	RightLegHits      int32
}
