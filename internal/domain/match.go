package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const defaultTeamSize = 5

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

func NewMatchTeam(id, score int32, state MatchState, name, flag string) *MatchTeam {
	return &MatchTeam{
		ID:         id,
		Score:      score,
		State:      state,
		Name:       name,
		FlagCode:   flag,
		ScoreBoard: make([]*MatchScoreBoardRow, 0, defaultTeamSize),
	}
}

type MatchScoreBoardRow struct {
	SteamID            SteamID
	PlayerName         string
	PlayerAvatarURL    string
	PlayerCaptain      bool
	Kills              int32
	Deaths             int32
	Assists            int32
	MVPCount           int32
	KillDeathRatio     float64
	DamagePerRound     float64
	HeadshotPercentage float64
}

func NewMatchScoreBoardRow(
	steamID uint64,
	playerName string,
	playerAvatarURL string,
	playerCaptain bool,
	kills int32,
	hsKills int32,
	deaths int32,
	assists int32,
	mvps int32,
	dmgDealt int32,
	roundsPlayed int32,
) MatchScoreBoardRow {
	return MatchScoreBoardRow{
		SteamID:            SteamID(steamID),
		PlayerName:         playerName,
		PlayerAvatarURL:    playerAvatarURL,
		PlayerCaptain:      playerCaptain,
		Kills:              kills,
		Deaths:             deaths,
		Assists:            assists,
		MVPCount:           mvps,
		KillDeathRatio:     calculateKD(float64(kills), float64(deaths)),
		DamagePerRound:     calculateADR(float64(dmgDealt), float64(roundsPlayed)),
		HeadshotPercentage: calculateHSPercentage(float64(hsKills), float64(kills)),
	}
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
