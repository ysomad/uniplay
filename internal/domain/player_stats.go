package domain

import (
	"time"
)

type PlayerStats struct {
	Total *PlayerTotalStats
	Calc  PlayerCalcStats
	Round PlayerRoundStats
}

func NewPlayerStats(t *PlayerTotalStats) PlayerStats {
	return PlayerStats{
		Total: t,
		Calc:  newPlayerCalcStats(t.Kills, t.Deaths, t.HeadshotKills, t.Wins, t.MatchesPlayed),
		Round: newPlayerRoundStats(t.Kills, t.Deaths, t.DamageDealt, t.Assists, t.GrenadeDamageDealt, t.BlindedPlayers, t.BlindedTimes, t.RoundsPlayed),
	}
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

func newPlayerCalcStats(kills, deaths int32, hsKills, wins, matchesPlayed int16) PlayerCalcStats {
	s := PlayerCalcStats{}
	fKills := float64(kills)

	if fKills > 0 && deaths > 0 {
		s.KillDeathRatio = round(fKills / float64(deaths))
		s.HeadshotPercentage = round(float64(hsKills) / fKills * 100)
	}

	if matchesPlayed > 0 {
		s.WinRate = round(float64(wins) / float64(matchesPlayed) * 100)
	}

	return s
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

func newPlayerRoundStats(kills, deaths, dmgDealt int32, assists, grenadeDmgDealt, blindedPlayers, blindedTimes, roundsPlayed int16) PlayerRoundStats {
	if roundsPlayed <= 0 {
		return PlayerRoundStats{}
	}

	floatRoundsPlayed := float64(roundsPlayed)
	return PlayerRoundStats{
		Kills:              round(float64(kills) / floatRoundsPlayed),
		Assists:            round(float64(assists) / floatRoundsPlayed),
		Deaths:             round(float64(deaths) / floatRoundsPlayed),
		DamageDealt:        round(float64(dmgDealt) / floatRoundsPlayed),
		GrenadeDamageDealt: round(float64(grenadeDmgDealt) / floatRoundsPlayed),
		BlindedPlayers:     round(float64(blindedPlayers) / floatRoundsPlayed),
		BlindedTimes:       round(float64(blindedTimes) / floatRoundsPlayed),
	}
}
