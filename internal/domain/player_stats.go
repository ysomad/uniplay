package domain

import (
	"time"
)

type PlayerStats struct {
	Total *PlayerTotalStats
	Calc  *PlayerCalcStats
	Round *PlayerRoundStats
}

// PlayerTotalStats is a set of total statistics of a player.
type PlayerTotalStats struct {
	Kills             int32
	HSKills           int16
	BlindKills        int16
	WallbangKills     int16
	NoScopeKills      int16
	ThroughSmokeKills int16
	Deaths            int32
	Assists           int16
	FlashbangAssists  int16
	MVPCount          int16
	DamageTaken       int32
	DamageDealt       int32
	BlindedPlayers    int16
	BlindedTimes      int16
	BombsPlanted      int16
	BombsDefused      int16
}

func (ts *PlayerTotalStats) Add(m Metric, v int) {
	switch m {
	case MetricKill:
		ts.Kills += int32(v)
	case MetricHSKill:
		ts.HSKills += int16(v)
	case MetricBlindKill:
		ts.BlindKills += int16(v)
	case MetricWallbangKill:
		ts.WallbangKills += int16(v)
	case MetricNoScopeKill:
		ts.NoScopeKills += int16(v)
	case MetricThroughSmokeKill:
		ts.ThroughSmokeKills += int16(v)
	case MetricDeath:
		ts.Deaths += int32(v)
	case MetricAssist:
		ts.Assists += int16(v)
	case MetricFlashbangAssist:
		ts.FlashbangAssists += int16(v)
	case MetricRoundMVP:
		ts.MVPCount += int16(v)
	case MetricDamageTaken:
		ts.DamageTaken += int32(v)
	case MetricDamageDealt:
		ts.DamageDealt += int32(v)
	case MetricBlind:
		ts.BlindedPlayers += int16(v)
	case MetricBlinded:
		ts.BlindedTimes += int16(v)
	case MetricBombPlanted:
		ts.BombsPlanted += int16(v)
	case MetricBombDefused:
		ts.BombsDefused += int16(v)
	}
}

// PlayerCalcStats is a set of calculated stats from player total stats and match history.
type PlayerCalcStats struct {
	KillDeathRatio float64
	HSPercentage   float64
	WinRate        float64
	TimePlayed     time.Duration
	Wins           int16
	Loses          int16
	Draws          int16
	RoundsPlayed   int16
	MatchesPlayed  int16
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
