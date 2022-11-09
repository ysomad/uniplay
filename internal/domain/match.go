package domain

import "time"

type Match struct {
	MapName  string
	Duration time.Duration
	Team1    MatchTeam
	Team2    MatchTeam
}

type MatchTeam struct {
	Name     string
	FlagCode string
	Score    int
}

func (t *MatchTeam) SetAll(name, flag string, score int) {
	t.Name = name
	t.FlagCode = flag
	t.Score = score
}
