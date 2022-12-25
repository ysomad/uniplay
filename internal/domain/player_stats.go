package domain

import (
	"time"
)

type PlayerStats struct {
	Total *PlayerTotalStats
	Calc  PlayerCalcStats
	Round PlayerRoundStats
}

// PlayerTotalStats is a set of total statistics of a player.
type PlayerTotalStats struct {
	Kills              int32
	HeadshotKills      int16
	BlindKills         int16
	WallbangKills      int16
	NoScopeKills       int16
	ThroughSmokeKills  int16
	Deaths             int32
	Assists            int16
	FlashbangAssists   int16
	MVPCount           int16
	DamageTaken        int32
	DamageDealt        int32
	GrenadeDamageDealt int16
	BlindedPlayers     int16
	BlindedTimes       int16
	BombsPlanted       int16
	BombsDefused       int16
	RoundsPlayed       int16
	MatchesPlayed      int16
	Wins               int16
	Loses              int16
	Draws              int16
	TimePlayed         time.Duration
}

// PlayerCalcStats is a set of calculated stats from player total stats and match history.
type PlayerCalcStats struct {
	HeadshotPercentage float64
	KillDeathRatio     float64
	WinRate            float64
}

type PlayerCalcStatsParams struct {
	MatchesPlayed int16
	Kills         int32
	Deaths        int32
	HeadshotKills int16
	Wins          int16
	Loses         int16
}

func NewPlayerCalcStats(p PlayerCalcStatsParams) PlayerCalcStats {
	return PlayerCalcStats{
		KillDeathRatio:     float64(p.Deaths) / float64(p.Kills) * 100,
		HeadshotPercentage: float64(p.HeadshotKills) / float64(p.Kills) * 100,
		WinRate:            float64(p.Wins) / float64(p.MatchesPlayed) * 100,
	}
}

// PlayerRoundStats is a set of AVG player stats per round.
type PlayerRoundStats struct {
	Kills              float64
	Assists            float64
	Deaths             float64
	DamageDealt        float64
	GrenadeDamageDealt float64
	BlindedPlayers     float64
	BlindedTimes       float64
}

type PlayerRoundStatsParams struct {
	Kills              int32
	Deaths             int32
	DamageDealt        int32
	Assists            int16
	RoundsPlayed       int16
	GrenadeDamageDealt int16
	BlindedPlayers     int16
	BlindedTimes       int16
}

func NewPlayerRoundStats(p PlayerRoundStatsParams) PlayerRoundStats {
	floatRoundsPlayed := float64(p.RoundsPlayed)
	return PlayerRoundStats{
		Kills:              float64(p.Kills) / floatRoundsPlayed,
		Assists:            float64(p.Assists) / floatRoundsPlayed,
		Deaths:             float64(p.Deaths) / floatRoundsPlayed,
		DamageDealt:        float64(p.DamageDealt) / floatRoundsPlayed,
		GrenadeDamageDealt: float64(p.GrenadeDamageDealt) / floatRoundsPlayed,
		BlindedPlayers:     float64(p.BlindedPlayers) / floatRoundsPlayed,
		BlindedTimes:       float64(p.BlindedTimes) / floatRoundsPlayed,
	}
}
