package domain

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID       uuid.UUID
	MapName  string
	Duration time.Duration
	Team1    MatchTeam
	Team2    MatchTeam
}

type MatchTeam struct {
	Name           string
	FlagCode       string
	Score          int8
	PlayerSteamIDs []uint64
}

func (t *MatchTeam) SetAll(name, flag string, score int8, steamIDs []uint64) {
	t.Name = name
	t.FlagCode = flag
	t.Score = score
	t.PlayerSteamIDs = steamIDs
}
