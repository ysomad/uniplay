package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MatchState int8

const (
	MatchStateLose MatchState = -1
	MatchStateDraw MatchState = 0
	MatchStateWin  MatchState = 1
)

const (
	minMatchDuration = time.Minute * 5
)

func NewMatchState(teamScore, opponentScore int8) MatchState {
	if teamScore > opponentScore {
		return MatchStateWin
	}
	if teamScore < opponentScore {
		return MatchStateLose
	}
	return MatchStateDraw
}

type Match struct {
	ID         uuid.UUID
	MapName    string
	Duration   time.Duration
	Team1      MatchTeam
	Team2      MatchTeam
	UploadedAt time.Time
}

type MatchTeam struct {
	ID       int16
	ClanName string
	FlagCode string
	Score    int8
	Players  []uint64
}

// NewMatchID returns match id generated from meta data from replay header.
func NewMatchID(server, client, mapName string, matchDuration time.Duration, ticks, frames, signonLen int) (uuid.UUID, error) {
	if server == "" {
		return uuid.UUID{}, errors.New("server name cannot be empty string")
	}
	if client == "" {
		return uuid.UUID{}, errors.New("client name cannot be empty string")
	}
	if matchDuration < minMatchDuration {
		return uuid.UUID{}, errors.New("match duration cannot last less than 5 minutes")
	}
	if ticks <= 0 {
		return uuid.UUID{}, errors.New("got invalid amount of playback ticks")
	}
	if frames <= 0 {
		return uuid.UUID{}, errors.New("got invalid amount of playback frames")
	}
	if signonLen <= 0 {
		return uuid.UUID{}, errors.New("got invalid amount of signon length")
	}

	s := fmt.Sprintf("%s,%s,%s,%d,%d,%d,%d", server, client, mapName, matchDuration, ticks, frames, signonLen)
	return uuid.NewMD5(uuid.UUID{}, []byte(s)), nil
}
