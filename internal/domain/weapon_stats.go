package domain

// WeaponStat is a set of weapon statistics calculated from sum of metrics.
// Each field corresponds to specific weapon metric.
type WeaponStat struct {
	Assists           uint32 `json:"assists"`
	BlindKills        uint32 `json:"blind_kills"`
	DamageDealt       uint32 `json:"damage_dealt"`
	DamageTaken       uint32 `json:"damage_taken"`
	Deaths            uint32 `json:"deaths"`
	HSKills           uint32 `json:"headshot_kills"`
	Kills             uint32 `json:"kills"`
	NoscopeKills      uint32 `json:"noscope_kills"`
	ThroughSmokeKills uint32 `json:"through_smoke_kills"`
	WallbangKills     uint32 `json:"wallbang_kills"`
}

// SetStat sets v into specific field depends on metric.
func (s *WeaponStat) SetStat(m Metric, v uint32) {
	switch m {
	case MetricDeath:
		s.Deaths = v
	case MetricKill:
		s.Kills = v
	case MetricHSKill:
		s.HSKills = v
	case MetricBlindKill:
		s.BlindKills = v
	case MetricWallbangKill:
		s.WallbangKills = v
	case MetricNoScopeKill:
		s.NoscopeKills = v
	case MetricThroughSmokeKill:
		s.ThroughSmokeKills = v
	case MetricAssist:
		s.Assists = v
	case MetricDamageTaken:
		s.DamageTaken = v
	case MetricDamageDealt:
		s.DamageDealt = v
	}
}

type WeaponStats map[string]*WeaponStat
type WeaponClassStats map[WeaponClass]WeaponStat

type WeaponStatsFilter struct {
	WeaponName  string
	WeaponClass WeaponClass
}
