package domain

import (
	"time"

	"github.com/ysomad/uniplay/internal/pkg/stat"
)

type PlayerStats struct {
	*PlayerBaseStats
	*PlayerCalcStats
}

func NewPlayerStats(s *PlayerBaseStats) PlayerStats {
	return PlayerStats{
		s,
		&PlayerCalcStats{
			HeadshotPercentage:     stat.HeadshotPercentage(s.HeadshotKills, s.Kills),
			WinRate:                stat.WinRate(s.Wins, s.MatchesPlayed),
			KD:                     stat.KD(s.Kills, s.Deaths),
			ADR:                    stat.ADR(s.DamageDealt, s.RoundsPlayed),
			KillsPerRound:          stat.AVG(s.Kills, s.RoundsPlayed),
			AssistsPerRound:        stat.AVG(s.Assists, s.RoundsPlayed),
			DeathsPerRound:         stat.AVG(s.Deaths, s.RoundsPlayed),
			GrenadeDmgPerRound:     stat.AVG(s.GrenadeDamageDealt, s.RoundsPlayed),
			BlindedPlayersPerRound: stat.AVG(s.BlindedPlayers, s.RoundsPlayed),
		},
	}
}

// PlayerBaseStats is a set of base statistics of a player.
type PlayerBaseStats struct {
	Kills              int32
	HeadshotKills      int32
	BlindKills         int32
	WallbangKills      int32
	NoScopeKills       int32
	ThroughSmokeKills  int32
	Deaths             int32
	Assists            int32
	FlashbangAssists   int32
	MVPCount           int32
	DamageTaken        int32
	DamageDealt        int32
	GrenadeDamageDealt int32
	BlindedPlayers     int32
	BlindedTimes       int32
	BombsPlanted       int32
	BombsDefused       int32
	RoundsPlayed       int32
	MatchesPlayed      int32
	Wins               int32
	Loses              int32
	TimePlayed         time.Duration
}

// PlayerCalcStats is a set of calculated stats from player total stats and match history.
type PlayerCalcStats struct {
	HeadshotPercentage     float64
	WinRate                float64
	KD                     float64
	ADR                    float64
	KillsPerRound          float64
	AssistsPerRound        float64
	DeathsPerRound         float64
	GrenadeDmgPerRound     float64
	BlindedPlayersPerRound float64
}
