package domain

import "time"

type PlayerStats struct {
	Total    *PlayerTotalStats
	Basic    PlayerBasicStats
	PerRound *PlayerPerRoundStats
}

// PlayerTotalStats is a set of total statistics of a player.
type PlayerTotalStats struct {
	Deaths            uint32
	Kills             uint32
	HSKills           uint32
	BlindKills        uint32
	WallbangKills     uint32
	NoScopeKills      uint32
	ThroughSmokeKills uint32
	Assists           uint32
	FlashbangAssists  uint32
	DamageTaken       uint32
	DamageDealt       uint32
	BombsPlanted      uint32
	BombsDefused      uint32
	MVPCount          uint32
	BlindedPlayers    uint32
	BlindedTimes      uint32
}

// PlayerBasicStats is a set of calculated stats from metrics.
type PlayerBasicStats struct {
	KillDeathRatio float64
	HSPercentage   float64
	RoundsPlayed   uint32
	TimePlayed     time.Duration
	MatchesPlayed  uint16
	Wins           uint16
	Loses          uint16
	Draws          uint16
}

// PlayerPerRoundStats is a set of AVG player stats per round.
type PlayerPerRoundStats struct {
	Kills              float64
	Assists            float64
	Deaths             float64
	DamageDealt        float64
	GrenadeDamageDealt float64
	BlindedPlayers     float64
	BlindedTimes       float64
}
