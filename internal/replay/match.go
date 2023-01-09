package replay

import (
	"time"

	"github.com/google/uuid"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type replayMatch struct {
	id         uuid.UUID
	team1      replayTeam
	team2      replayTeam
	mapName    string
	duration   time.Duration
	uploadedAt time.Time
}

func (m *replayMatch) toDomain() *domain.Match {
	return &domain.Match{
		ID:       m.id,
		MapName:  m.mapName,
		Duration: m.duration,
		Team1: domain.MatchTeam{
			ID:       m.team1.id,
			ClanName: m.team1.clanName,
			FlagCode: m.team1.flagCode,
			Score:    m.team1.score,
			Players:  m.team1.players,
		},
		Team2: domain.MatchTeam{
			ID:       m.team2.id,
			ClanName: m.team2.clanName,
			FlagCode: m.team2.flagCode,
			Score:    m.team2.score,
			Players:  m.team2.players,
		},
		UploadedAt: m.uploadedAt,
	}
}

// setTeams set new team to a match from GameState.
func (m *replayMatch) setTeams(gs demoinfocs.GameState) {
	t := gs.TeamTerrorists()
	m.team1 = newTeam(t.ClanName(), t.Flag(), t.Team(), t.Members())

	ct := gs.TeamCounterTerrorists()
	m.team2 = newTeam(ct.ClanName(), ct.Flag(), ct.Team(), ct.Members())
}

func (m *replayMatch) swapTeamSides() {
	m.team1.swapSide()
	m.team2.swapSide()
}

func (m *replayMatch) updateTeamsScore(e events.ScoreUpdated) {
	switch e.TeamState.Team() {
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

	for i, steamID := range m.team1.players {
		res[i] = teamPlayer{
			steamID:    steamID,
			teamID:     m.team1.id,
			matchID:    m.id,
			matchState: m.team1.matchState,
		}
	}

	for i, steamID := range m.team2.players {
		res[i+len(m.team2.players)] = teamPlayer{
			steamID:    steamID,
			teamID:     m.team2.id,
			matchID:    m.id,
			matchState: m.team2.matchState,
		}
	}

	return res
}

type replayTeam struct {
	id         int16
	clanName   string
	flagCode   string
	score      int8
	matchState domain.MatchState
	players    []uint64

	_side common.Team
}

func newTeam(name, flag string, side common.Team, players []*common.Player) replayTeam {
	var steamIDs []uint64

	for _, p := range players {
		if p != nil && p.SteamID64 != 0 {
			steamIDs = append(steamIDs, p.SteamID64)
		}
	}

	return replayTeam{
		clanName: name,
		flagCode: flag,
		players:  steamIDs,
		_side:    side,
	}
}

func (t *replayTeam) swapSide() {
	switch t._side {
	case common.TeamCounterTerrorists:
		t._side = common.TeamTerrorists
	case common.TeamTerrorists:
		t._side = common.TeamCounterTerrorists
	}
}
