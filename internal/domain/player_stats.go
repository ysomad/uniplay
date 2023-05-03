package domain

import (
	"time"

	"github.com/ysomad/uniplay/internal/pkg/floatrounder"
)

type PlayerStats struct {
	Base  *PlayerBaseStats
	Calc  PlayerCalcStats
	Round PlayerRoundStats
}

func NewPlayerStats(t *PlayerBaseStats) PlayerStats {
	return PlayerStats{
		Base:  t,
		Calc:  newPlayerCalcStats(t.Kills, t.Deaths, t.HeadshotKills, t.Wins, t.MatchesPlayed),
		Round: newPlayerRoundStats(t.Kills, t.Deaths, t.DamageDealt, t.Assists, t.GrenadeDamageDealt, t.BlindedPlayers, t.BlindedTimes, t.RoundsPlayed),
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
	Draws              int32
	TimePlayed         time.Duration
}

// PlayerCalcStats is a set of calculated stats from player total stats and match history.
type PlayerCalcStats struct {
	HeadshotPercentage float64
	KillDeathRatio     float64
	WinRate            float64
}

// calculateKD returns calculated and rounded kill death ratio.
func calculateKD(kills, deaths float64) float64 {
	if kills <= 0 || deaths <= 0 {
		return 0
	}

	return floatrounder.Round(kills / deaths)
}

// calculateHSPercentage returns calculated and rounded headshort percentage.
func calculateHSPercentage(hsKills, kills float64) float64 {
	if kills <= 0 {
		return 0
	}

	return floatrounder.Round(hsKills / kills * 100)
}

func newPlayerCalcStats(kills, deaths, hsKills, wins, matchesPlayed int32) PlayerCalcStats {
	s := PlayerCalcStats{}
	fKills := float64(kills)

	if fKills > 0 && deaths > 0 {
		s.KillDeathRatio = calculateKD(fKills, float64(deaths))
		s.HeadshotPercentage = calculateHSPercentage(float64(hsKills), fKills)
	}

	if matchesPlayed > 0 && wins >= 0 {
		s.WinRate = floatrounder.Round(float64(wins) / float64(matchesPlayed) * 100)
	}

	return s
}

// calculateADR returns calculated and rounded average damage per round.
func calculateADR(dmgDealt, roundsPlayed float64) float64 {
	if roundsPlayed <= 0 {
		return 0
	}

	return floatrounder.Round(dmgDealt / roundsPlayed)
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

func newPlayerRoundStats(kills, deaths, dmgDealt, assists, grenadeDmgDealt, blindedPlayers, blindedTimes, roundsPlayed int32) PlayerRoundStats {
	if roundsPlayed <= 0 {
		return PlayerRoundStats{}
	}

	floatRoundsPlayed := float64(roundsPlayed)

	return PlayerRoundStats{
		Kills:              floatrounder.Round(float64(kills) / floatRoundsPlayed),
		Assists:            floatrounder.Round(float64(assists) / floatRoundsPlayed),
		Deaths:             floatrounder.Round(float64(deaths) / floatRoundsPlayed),
		DamageDealt:        calculateADR(float64(dmgDealt), floatRoundsPlayed),
		GrenadeDamageDealt: floatrounder.Round(float64(grenadeDmgDealt) / floatRoundsPlayed),
		BlindedPlayers:     floatrounder.Round(float64(blindedPlayers) / floatRoundsPlayed),
		BlindedTimes:       floatrounder.Round(float64(blindedTimes) / floatRoundsPlayed),
	}
}
