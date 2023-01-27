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

const minMatchDuration = time.Minute * 5

func NewMatchState(teamScore, opponentScore int8) MatchState {
	if teamScore > opponentScore {
		return MatchStateWin
	}
	if teamScore < opponentScore {
		return MatchStateLose
	}
	return MatchStateDraw
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

	if ticks <= 0 {
		return uuid.UUID{}, errors.New("invalid amount if playback ticks")
	}

	if frames <= 0 {
		return uuid.UUID{}, errors.New("invalid amount of playback frames")
	}

	if signonLen <= 0 {
		return uuid.UUID{}, errors.New("invalid signon length")
	}

	s := fmt.Sprintf("%s,%s,%s,%d,%d,%d,%d", server, client, mapName, matchDuration, ticks, frames, signonLen)
	return uuid.NewMD5(uuid.UUID{}, []byte(s)), nil
}
