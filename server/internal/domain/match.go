package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// MatchScore represents score of match in format 16:1.
type MatchScore string

func NewMatchScore(score1, score2 int8) MatchScore {
	return MatchScore(fmt.Sprintf("%d:%d", score1, score2))
}

const DefaultTeamSize = 5

type Match struct {
	ID           uuid.UUID
	Map          Map
	Team1        *MatchTeam
	Team2        *MatchTeam
	RoundsPlayed int32
	Duration     time.Duration
	UploadedAt   time.Time
}

type MatchTeam struct {
	ID         int32
	Score      int32
	State      MatchState
	Name       string
	FlagCode   string
	ScoreBoard []*MatchScoreBoardRow
}

type MatchScoreBoardRow struct {
	SteamID            SteamID
	PlayerName         string
	PlayerAvatarURL    string
	IsPlayerCaptain    bool
	Kills              int32
	Deaths             int32
	Assists            int32
	MVPCount           int32
	KD                 float64
	ADR                float64
	HeadshotPercentage float64
}

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
