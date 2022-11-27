package domain

import "github.com/google/uuid"

type PlayerMatchState int8

const (
	PlayerMatchLose = -1
	PlayerMatchDraw = 0
	PlayerMatchWin  = 1
)

type PlayerMatch struct {
	MatchID uuid.UUID
	SteamID uint64
	Team    string
	State   PlayerMatchState
}
