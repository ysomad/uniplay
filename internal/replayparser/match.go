package replayparser

import (
	"fmt"
	"sync"
	"time"

	"github.com/ssssargsian/uniplay/internal/dto"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
)

type match struct {
	mu            sync.RWMutex
	_team1        matchTeam
	_team2        matchTeam
	_isKnifeRound bool

	mapName  string
	duration time.Duration
}

func (m *match) swapTeamSides() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m._team1.swapSide()
	m._team2.swapSide()
}

func (m *match) updateTeamsScore(e events.ScoreUpdated) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if e.NewScore <= 0 {
		return
	}

	switch e.TeamState.Team() {
	case m._team1.side:
		m._team1.score = e.NewScore
	case m._team2.side:
		m._team2.score = e.NewScore
	}

	fmt.Println(m._team1.score)
	fmt.Println(m._team2.score)
}

// isKnifeRound returns true if current round is a knife round.
func (m *match) isKnifeRound() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m._isKnifeRound
}

func (m *match) toDTO() *dto.Match {
	return &dto.Match{
		MapName:  m.mapName,
		Duration: m.duration,
		Team1:    m._team1.toDTO(),
		Team2:    m._team2.toDTO(),
	}
}

// knifeRound returns true if current round is a knife round.
func (m *match) setIsKnifeRound(v bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m._isKnifeRound = v
}

func (m *match) setTeam1(t matchTeam) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m._team1 = t
}

func (m *match) setTeam2(t matchTeam) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m._team2 = t
}

type matchTeam struct {
	clanName       string
	flagCode       string
	side           common.Team
	score          int
	playerSteamIDs []uint64
}

func newMatchTeam(name, flag string, side common.Team, players []*common.Player) matchTeam {
	var steamIDs []uint64

	for _, p := range players {
		if p != nil && p.SteamID64 != 0 {
			steamIDs = append(steamIDs, p.SteamID64)
		}
	}

	return matchTeam{
		clanName:       name,
		flagCode:       flag,
		side:           side,
		playerSteamIDs: steamIDs,
	}
}

func (t *matchTeam) swapSide() {
	switch t.side {
	case common.TeamCounterTerrorists:
		t.side = common.TeamTerrorists
	case common.TeamTerrorists:
		t.side = common.TeamCounterTerrorists
	}
}

func (t *matchTeam) toDTO() dto.MatchTeam {
	return dto.MatchTeam{
		ClanName:       t.clanName,
		FlagCode:       t.flagCode,
		Score:          uint8(t.score),
		PlayerSteamIDs: t.playerSteamIDs,
	}
}
