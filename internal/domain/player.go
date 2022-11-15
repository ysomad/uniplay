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

type PlayerStats struct {
	TotalKills             uint32
	TotalDeaths            uint32
	KillDeathRatio         float32
	DamagePerRound         float32 // ADR
	GrenadeDamangePerRound float32
	KillsPerRound          float32 // KPR
	AssistsPerRound        float32
	DeathsPerRound         float32 // DPR
	BlindPerRound          float32
	BlindedPerRound        float32
	HeadshotPercentage     float32
	MatchesPlayed          uint16
	RoundsPlayed           uint32
}

type PlayerWeaponStats []struct {
	WeaponName string
	TotalKills uint32
}

type PlayerWeaponClassStats []struct {
	WeaponClass string
	TotalKills  uint32
}
