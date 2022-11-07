package metric

type WeaponEvent struct {
	Event        Metric
	Weapon       string // HE Grenade for example or AWP
	HealthDamage uint16
	ArmorDamage  uint16
}
