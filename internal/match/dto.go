package match

import (
	"time"

	"github.com/google/uuid"

	"github.com/ysomad/uniplay/internal/domain"
)

type matchScoreBoardRow struct {
	MatchID         uuid.UUID
	MapName         string
	MapIconURL      string
	SteamID         uint64
	PlayerName      string
	PlayerAvatarURL string
	PlayerCaptain   bool
	TeamID          int32
	TeamName        string
	TeamFlagCode    string
	TeamScore       int32
	TeamMatchState  domain.MatchState
	Kills           int32
	HeadshotKills   int32
	Deaths          int32
	Assists         int32
	DamageDealt     int32
	RoundsPlayed    int32
	MVPCount        int32
	MatchDuration   time.Duration
	MatchUploadedAt time.Time
}
