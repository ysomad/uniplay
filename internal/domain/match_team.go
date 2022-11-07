package domain

type MatchTeam struct {
	Name     string `json:"name"`
	FlagCode string `json:"flag_code"`
	Score    int    `json:"score"`
}

func (t *MatchTeam) SetAll(name, flag string, score int) {
	t.Name, t.FlagCode, t.Score = name, flag, score
}

func (t *MatchTeam) IsWinner(opponent MatchTeam) bool {
	return t.Score > opponent.Score
}

type MatchTeams struct {
	Team1 MatchTeam `json:"team1"`
	Team2 MatchTeam `json:"team2"`
}
