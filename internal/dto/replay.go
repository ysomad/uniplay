package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssssargsian/uniplay/internal/domain"
)

type ReplayMatch struct {
	ID         uuid.UUID
	Team1      ReplayTeam
	Team2      ReplayTeam
	MapName    string
	Duration   time.Duration
	UploadedAt time.Time
}

func (m *ReplayMatch) MatchTeamPlayerList() []MatchTeamPlayer {
	res := make([]MatchTeamPlayer, len(m.Team1.PlayerSteamIDs)+len(m.Team2.PlayerSteamIDs))

	// team 1 players
	for i, steamID := range m.Team1.PlayerSteamIDs {
		res[i] = MatchTeamPlayer{
			SteamID:    steamID,
			TeamID:     m.Team1.ID,
			MatchID:    m.ID,
			MatchState: m.Team1.State,
		}
	}

	// team 2 players
	for i, steamID := range m.Team2.PlayerSteamIDs {
		res[i+len(m.Team2.PlayerSteamIDs)] = MatchTeamPlayer{
			SteamID:    steamID,
			TeamID:     m.Team2.ID,
			MatchID:    m.ID,
			MatchState: m.Team2.State,
		}
	}

	return res
}

type ReplayTeam struct {
	ID             int16
	ClanName       string
	FlagCode       string
	Score          int8
	PlayerSteamIDs []uint64
	State          domain.MatchState
}

type MatchTeamPlayer struct {
	SteamID    uint64
	TeamID     int16
	MatchID    uuid.UUID
	MatchState domain.MatchState
}

type PlayerStat struct {
	SteamID           uint64
	Kills             int
	HSKills           int
	BlindKills        int
	WallbangKills     int
	NoScopeKills      int
	ThroughSmokeKills int
	Deaths            int
	Assists           int
	FlashbangAssists  int
	MVPCount          int
	DamageTaken       int
	DamageDealt       int
	BlindedPlayers    int
	BlindedTimes      int
	BombsPlanted      int
	BombsDefused      int
}

func (ts *PlayerStat) Add(m domain.Metric, v int) {
	switch m {
	case domain.MetricKill:
		ts.Kills += v
	case domain.MetricHSKill:
		ts.HSKills += v
	case domain.MetricBlindKill:
		ts.BlindKills += v
	case domain.MetricWallbangKill:
		ts.WallbangKills += v
	case domain.MetricNoScopeKill:
		ts.NoScopeKills += v
	case domain.MetricThroughSmokeKill:
		ts.ThroughSmokeKills += v
	case domain.MetricDeath:
		ts.Deaths += v
	case domain.MetricAssist:
		ts.Assists += v
	case domain.MetricFlashbangAssist:
		ts.FlashbangAssists += v
	case domain.MetricRoundMVP:
		ts.MVPCount += v
	case domain.MetricDamageTaken:
		ts.DamageTaken += v
	case domain.MetricDamageDealt:
		ts.DamageDealt += v
	case domain.MetricBlind:
		ts.BlindedPlayers += v
	case domain.MetricBlinded:
		ts.BlindedTimes += v
	case domain.MetricBombPlanted:
		ts.BombsPlanted += v
	case domain.MetricBombDefused:
		ts.BombsDefused += v
	}
}

type PlayerWeaponStat struct {
	SteamID           uint64
	WeaponID          int16
	Kills             int
	HSKills           int
	BlindKills        int
	WallbangKills     int
	NoScopeKills      int
	ThroughSmokeKills int
	Deaths            int
	Assists           int
	DamageTaken       int
	DamageDealt       int
	Shots             int
	Hits              int
	HeadHits          int
	ChestHits         int
	StomachHits       int
	LeftArmHits       int
	RightArmHits      int
	LeftLegHits       int
	RightLegHits      int
}

func (ws *PlayerWeaponStat) AddStat(m domain.Metric, v int) {
	switch m {
	case domain.MetricKill:
		ws.Kills += v
	case domain.MetricHSKill:
		ws.HSKills += v
	case domain.MetricBlindKill:
		ws.BlindKills += v
	case domain.MetricWallbangKill:
		ws.WallbangKills += v
	case domain.MetricNoScopeKill:
		ws.NoScopeKills += v
	case domain.MetricThroughSmokeKill:
		ws.ThroughSmokeKills += v
	case domain.MetricDeath:
		ws.Deaths += v
	case domain.MetricAssist:
		ws.Assists += v
	case domain.MetricDamageTaken:
		ws.DamageTaken += v
	case domain.MetricDamageDealt:
		ws.DamageDealt += v
	case domain.MetricShot:
		ws.Shots += v
	case domain.MetricHitHead:
		ws.HeadHits += v
	case domain.MetricHitChest:
		ws.ChestHits += v
	case domain.MetricHitStomach:
		ws.StomachHits += v
	case domain.MetricHitLeftArm:
		ws.LeftArmHits += v
	case domain.MetricHitRightArm:
		ws.RightArmHits += v
	case domain.MetricHitLeftLeg:
		ws.LeftLegHits += v
	case domain.MetricHitRightLeg:
		ws.RightLegHits += v
	}
}
