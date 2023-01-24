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
	minMatchDuration  = time.Minute * 5
	minServerTickrate = 64
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

// NewMatchID returns match id generated from meta data received from replay header.
func NewMatchID(server, client, mapName string, matchDuration time.Duration, ticks, frames, signonLen int) (uuid.UUID, error) {
	if server == "" {
		return uuid.UUID{}, errors.New("server name cannot be empty")
	}

	if client == "" {
		return uuid.UUID{}, errors.New("client name cannot be empty")
	}

	if matchDuration < minMatchDuration {
		return uuid.UUID{}, fmt.Errorf("match must last more than %s", minMatchDuration.String())
	}

	minPlaybackTicks := int(matchDuration) * minServerTickrate

	if ticks <= 0 || ticks < minPlaybackTicks {
		return uuid.UUID{}, errors.New(
			"invalid amount of playback ticks, must be more or equal to playback time * server tickrate")
	}

	if frames <= 0 || time.Duration(frames) <= matchDuration {
		return uuid.UUID{}, errors.New("invalid amount of frames, must be more than playback time")
	}

	if signonLen <= 0 {
		return uuid.UUID{}, errors.New("invalid signon length")
	}

	s := fmt.Sprintf("%s,%s,%s,%d,%d,%d,%d", server, client, mapName, matchDuration, ticks, frames, signonLen)
	return uuid.NewMD5(uuid.UUID{}, []byte(s)), nil
}
