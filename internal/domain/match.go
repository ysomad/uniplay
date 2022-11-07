package domain

import "time"

type Match struct {
	Map      string        `json:"map"`
	Duration time.Duration `json:"duration"`
	Team1    MatchTeam     `json:"team_1"`
	Team2    MatchTeam     `json:"team_2"`
}
