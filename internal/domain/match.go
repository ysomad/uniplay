package domain

import "time"

type Match struct {
	Map      string
	Duration time.Duration
	Team1    MatchTeam
	Team2    MatchTeam
}
