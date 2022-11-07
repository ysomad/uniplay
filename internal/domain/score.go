package domain

type roundScore struct {
	t  uint8
	ct uint8
}

func NewRoundScore() roundScore {
	return roundScore{}
}

// Set sets sides score.
func (s *roundScore) Set(TScore, CTScore int) {
	s.t, s.ct = uint8(TScore), uint8(CTScore)
}

// TotalRounds returns total number of rounds.
func (s roundScore) TotalRounds() uint8 { return s.t + s.ct }
