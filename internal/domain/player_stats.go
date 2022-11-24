package domain

import "time"

// PlayerStats is a calculated statistics from PlayerMetrics.
type PlayerStats struct {
	KillDeathRatio        float64
	DamagePerRound        float64
	GrenadeDamagePerRound float64
	KillsPerRound         float64
	AssistsPerRound       float64
	DeathsPerRound        float64
	BlindPerRound         float64
	BlindedPerRound       float64
	HeadshotPercentage    float64
	TimeStats             PlayerTimeStats
}

type PlayerTimeStats struct {
	RoundsPlayed  uint32
	TimePlayed    time.Duration
	MatchesPlayed uint16
}
