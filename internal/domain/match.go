package domain

import (
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
func NewMatchID(h *ReplayHeader) uuid.UUID {
	return uuid.NewMD5(uuid.UUID{}, []byte(fmt.Sprintf(
		"%d,%d,%d,%s,%s,%s,%d,%d",
		h.playbackTicks,
		h.playbackFrames,
		h.signonLength,
		h.server,
		h.client,
		h.mapName,
		h.playbackTime,
		h.filesize,
	)))
}
