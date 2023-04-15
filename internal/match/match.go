package match

import (
	"time"

	"github.com/google/uuid"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"

	"github.com/ysomad/uniplay/internal/domain"
)

type replayMatch struct {
	id         uuid.UUID
	team1      replayTeam
	team2      replayTeam
	mapName    string
	duration   time.Duration
	uploadedAt time.Time
}

// setTeams set new team to a match from GameState.
func (m *replayMatch) setTeams(t, ct *common.TeamState) {
	m.team1 = newReplayTeam(t.ClanName(), t.Flag(), t.Team(), t.Members())
	m.team2 = newReplayTeam(ct.ClanName(), ct.Flag(), ct.Team(), ct.Members())
}

func (m *replayMatch) swapTeamSides() {
	m.team1.swapSide()
	m.team2.swapSide()
}

func (m *replayMatch) updateTeamsScore(e events.ScoreUpdated) {
	switch e.TeamState.Team() { //nolint:exhaustive // teams can be only ct or t
	case m.team1._side:
		m.team1.score = int8(e.NewScore)
	case m.team2._side:
		m.team2.score = int8(e.NewScore)
	}
}

// setTeamStates sets MatchState to teams depending on their score.
func (m *replayMatch) setTeamStates() {
	m.team1.matchState = domain.NewMatchState(m.team1.score, m.team2.score)
	m.team2.matchState = domain.NewMatchState(m.team2.score, m.team1.score)
}

type teamPlayer struct {
	steamID    uint64
	teamID     int16
	matchID    uuid.UUID
	matchState domain.MatchState
}

// teamPlayers returns list of all steam ids of players participated in the match.
func (m *replayMatch) teamPlayers() []teamPlayer {
	res := make([]teamPlayer, len(m.team1.players)+len(m.team2.players))

	for i, p := range m.team1.players {
		res[i] = teamPlayer{
			steamID:    p.steamID,
			teamID:     m.team1.id,
			matchID:    m.id,
			matchState: m.team1.matchState,
		}
	}

	for i, p := range m.team2.players {
		res[i+len(m.team2.players)] = teamPlayer{
			steamID:    p.steamID,
			teamID:     m.team2.id,
			matchID:    m.id,
			matchState: m.team2.matchState,
		}
	}

	return res
}

type replayPlayer struct {
	steamID     uint64
	displayName string
}

type replayTeam struct {
	id         int16
	clanName   string
	flagCode   string
	score      int8
	matchState domain.MatchState
	players    []replayPlayer

	_side common.Team
}

func newReplayTeam(name, flag string, side common.Team, players []*common.Player) replayTeam {
	rt := replayTeam{
		clanName: name,
		flagCode: flag,
		players:  make([]replayPlayer, 0, len(players)),
		_side:    side,
	}

	for _, p := range players {
		if p == nil || p.SteamID64 == 0 {
			continue
		}

		rt.players = append(rt.players, replayPlayer{
			steamID:     p.SteamID64,
			displayName: p.Name,
		})
	}

	return rt
}

func (t *replayTeam) swapSide() {
	switch t._side { //nolint:exhaustive // player side cannot be spectator or unassigned
	case common.TeamCounterTerrorists:
		t._side = common.TeamTerrorists
	case common.TeamTerrorists:
		t._side = common.TeamCounterTerrorists
	}
}
