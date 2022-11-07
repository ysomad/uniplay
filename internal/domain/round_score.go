package domain

type RoundScore struct {
	t  uint8
	ct uint8
}

func NewRoundScore() RoundScore {
	return RoundScore{}
}

// Set sets sides score.
func (s *RoundScore) Set(TScore, CTScore int) {
	s.t, s.ct = uint8(TScore), uint8(CTScore)
}

// TotalRounds returns total number of rounds.
func (s RoundScore) TotalRounds() uint8 { return s.t + s.ct }
