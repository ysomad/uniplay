package metric

type Metric uint8

const (
	Death Metric = iota + 1

	Kill
	HSKill
	BlindKill
	WallbangKill
	NoScopeKill
	ThroughSmokeKill

	Assist
	FlashbangAssist

	// damage without over-damage (dealt 200, metric will be 200)
	DamageTaken
	DamageDealt

	// damage with over-damage (dealt 200, metric will be 100 if player had 100 hp)
	DamageTakenWithOver
	DamageDealtWithOver

	RoundMVPCount
)

type WeaponMetric struct {
	Metric Metric
	Weapon string // HE Grenade for example or AWP
	Damage uint16
}

const (
	DamageKill   uint16 = 100
	DamageAssist uint16 = 0
)
