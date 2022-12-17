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
	TeamName   string
	SteamID    uint64
	MatchState domain.PlayerMatchState
}

type PlayerStat struct {
	ID      domain.PlayerStatID
	SteamID uint64
	Metric  domain.Metric
	Value   uint32
}

func NewPlayerStat(steamID uint64, m domain.Metric, v uint32) PlayerStat {
	return PlayerStat{
		ID:      domain.NewPlayerStatID(steamID, m),
		SteamID: steamID,
		Metric:  m,
		Value:   v,
	}
}

type WeaponStat struct {
	ID       domain.WeaponStatID
	SteamID  uint64
	WeaponID uint16
	Metric   domain.Metric
	Value    uint32
}

func NewWeaponStat(steamID uint64, weaponID uint16, m domain.Metric, v uint32) WeaponStat {
	return WeaponStat{
		ID:       domain.NewWeaponStatID(steamID, weaponID, m),
		SteamID:  steamID,
		WeaponID: weaponID,
		Metric:   m,
		Value:    v,
	}
}
