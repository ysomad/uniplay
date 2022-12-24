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
	if teamScore == opponentScore {
		return MatchStateDraw
	}
	return 0
}

type Match struct {
	ID         uuid.UUID
	MapName    string
	Duration   time.Duration
	Team1      MatchTeam
	Team2      MatchTeam
	UploadedAt time.Time
}

// NewMatchID returns match id generated from meta data which was parsed from replay.
func NewMatchID(server, client, mapName string, duration time.Duration, ticks, frames, signonLen int) (uuid.UUID, error) {
	if server == "" {
		return uuid.UUID{}, errors.New("server name cannot be empty string")
	}
	if client == "" {
		return uuid.UUID{}, errors.New("client name cannot be empty string")
	}
	if duration < minMatchDuration {
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

	format := fmt.Sprintf("%s,%s,%s,%d,%d,%d,%d", server, client, mapName, duration, ticks, frames, signonLen)
	return uuid.NewMD5(uuid.UUID{}, []byte(format)), nil
}

type MatchTeam struct {
	ID             int16
	ClanName       string
	FlagCode       string
	Score          int8
	PlayerSteamIDs []uint64
}
