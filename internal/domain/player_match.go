package domain

import "github.com/google/uuid"

type PlayerMatchState string

const (
	PlayerMatchWin  = "WIN"
	PlayerMatchLose = "LOSE"
	PlayerMatchDraw = "DRAW"
)

func NewPlayerMatchState(teamScore, opponentScore uint8) PlayerMatchState {
	if teamScore > opponentScore {
		return PlayerMatchWin
	}

	if teamScore < opponentScore {
		return PlayerMatchLose
	}

	return PlayerMatchDraw
}

type PlayerMatch struct {
	MatchID uuid.UUID
	SteamID uint64
	Team    string
	State   PlayerMatchState
}
