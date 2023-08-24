package stat

// AVG returns average statistic with sum of terms and num as number of terms.
// Example based on calculating average damage per round(ADR):
// sum - damage per round, num - number of rounds.
func AVG(sum, num int32) float64 {
	if sum <= 0 || num <= 0 {
		return 0
	}

	return round(float64(sum) / float64(num))
}

// ADR returns average damage per round.
func ADR(dmg, rounds int32) float64 { return AVG(dmg, rounds) }

// WinRate returns win rate in percents based on total amount of wins and total matches played.
func WinRate(wins, matches int32) float64 {
	if wins <= 0 || matches <= 0 {
		return 0
	}

	return round(float64(wins) / float64(matches) * 100)
}

// KD returns kill death ratio based on overall kills and deaths.
func KD(kills, deaths int32) float64 {
	if kills <= 0 || deaths <= 0 {
		return 0
	}

	return round(float64(kills) / float64(deaths))
}

// HeadshotPercentage returns headshot percentage based on headshot and regular kills.
func HeadshotPercentage(hsKills, kills int32) float64 {
	if hsKills <= 0 || kills <= 0 {
		return 0
	}

	return round(float64(hsKills) / float64(kills) * 100)
}

// Accuracy returns accuracy in percents based on target and total hits.
func Accuracy(targetHits, totalHits int32) float64 {
	if targetHits <= 0 || totalHits <= 0 {
		return 0
	}

	return round(float64(targetHits) * 100 / float64(totalHits))
}
