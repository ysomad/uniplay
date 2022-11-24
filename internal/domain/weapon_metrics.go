package domain

type WeaponMetrics struct {
	Deaths            uint32
	Kills             uint32
	HSKills           uint32
	BlindKills        uint32
	WallbangKills     uint32
	NoScopeKills      uint32
	ThroughSmokeKills uint32
	Assists           uint32
	DamageTaken       uint32
	DamageDealt       uint32
}

type PlayerWeaponMetrics map[string]WeaponMetrics
