package domain

import "time"

type PlayerProfile struct {
	Player           Player
	Stats            PlayerStats
	TeamNames        []string
	WeaponStats      PlayerWeaponStats
	WeaponClassStats PlayerWeaponClassStats
}

type Player struct {
	SteamID      uint64
	MainTeamName string
	CreateTime   time.Time
	UpdateTime   time.Time
}

// TODO: update contract
type PlayerStats struct {
	MatchesPlayed         uint16
	RoundsPlayed          uint32
	KillDeathRatio        float64
	DamagePerRound        float64
	GrenadeDamagePerRound float64
	KillsPerRound         float64
	AssistsPerRound       float64
	DeathsPerRound        float64
	BlindPerRound         float64
	BlindedPerRound       float64
	HeadshotPercentage    float64
}

type MetricStats map[Metric]uint32

type PlayerWeaponStats []struct {
	WeaponName string
	TotalKills uint32
}

type PlayerWeaponClassStats []struct {
	WeaponClass string
	TotalKills  uint32
}
