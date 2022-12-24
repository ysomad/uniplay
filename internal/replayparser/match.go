package replayparser

import (
	"time"

	"github.com/google/uuid"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
)

type match struct {
	id       uuid.UUID
	team1    *matchTeam
	team2    *matchTeam
	mapName  string
	duration time.Duration
}

func (m *match) swapTeamSides() {
	m.team1.swapSide()
	m.team2.swapSide()
}

func (m *match) updateTeamsScore(e events.ScoreUpdated) {
	switch e.TeamState.Team() {
	case m.team1.side:
		m.team1.score = e.NewScore
	case m.team2.side:
		m.team2.score = e.NewScore
	}
}

type matchTeam struct {
	clanName       string
	flagCode       string
	side           common.Team
	score          int
	playerSteamIDs []uint64
}

func newMatchTeam(name, flag string, side common.Team, players []*common.Player) *matchTeam {
	var steamIDs []uint64

	for _, p := range players {
		if p != nil && p.SteamID64 != 0 {
			steamIDs = append(steamIDs, p.SteamID64)
		}
	}

	return &matchTeam{
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
