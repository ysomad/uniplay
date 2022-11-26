package dto

import (
	"time"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type Match struct {
	ID         domain.MatchID
	MapName    string
	Duration   time.Duration
	Team1      MatchTeam
	Team2      MatchTeam
	UploadTime time.Time
}

type MatchTeam struct {
	ClanName       string
	FlagCode       string
	Score          uint8
	PlayerSteamIDs []uint64
}

type MatchPlayers struct {
	MatchID    domain.MatchID
	Players    []TeamPlayer
	CreateTime time.Time
}

type Teams struct {
	Team1Name  string
	Team1Flag  string
	Team2Name  string
	Team2Flag  string
	CreateTime time.Time
}

type TeamPlayer struct {
	TeamName string
	SteamID  uint64
}

type Metric struct {
	MatchID       domain.MatchID
	PlayerSteamID uint64
	Metric        domain.Metric
	Value         int32
}

type WeaponMetric struct {
	MatchID       domain.MatchID
	PlayerSteamID uint64
	WeaponID      uint16
	Metric        domain.Metric
	Value         int32
}
