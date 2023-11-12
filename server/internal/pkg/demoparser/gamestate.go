package demoparser

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type gameStatus int8

const (
	gameStatusInProgress gameStatus = 0
	gameStatusDraw       gameStatus = 1
	gameStatusLose       gameStatus = -1
	gameStatusWin        gameStatus = 2
)

type round struct {
	StartedAt time.Time
	TeamA     *roundTeam
	TeamB     *roundTeam
	KillFeed  []*roundKill
	Reason    events.RoundEndReason
}

func newRound(ts *common.TeamState) *round {
	return &round{
		TeamA:     newRoundTeam(ts.Members(), ts.Team()),
		TeamB:     newRoundTeam(ts.Opponent.Members(), ts.Opponent.Team()),
		KillFeed:  make([]*roundKill, 0, 20),
		Reason:    events.RoundEndReasonStillInProgress,
		StartedAt: time.Now(),
	}
}

func (r *round) end(winner, loser *common.TeamState, reason events.RoundEndReason) {
	switch winner.Team() {
	case r.TeamA.Side:
		r.TeamA.roundEnd(winner)
		r.TeamB.roundEnd(loser)
	case r.TeamB.Side:
		r.TeamA.roundEnd(loser)
		r.TeamB.roundEnd(winner)
	}

	r.Reason = reason
}

type roundTeam struct {
	Survivors map[uint64]struct{}
	Cash      int // cash at start of round, must be set on round start
	CashSpend int // during round, must be set on round end
	EqValue   int // equipment value on round start, must be set on round end
	Side      common.Team
}

// newRoundTeam must be created at round start.
func newRoundTeam(members []*common.Player, side common.Team) *roundTeam {
	cash := 0
	survivors := make(map[uint64]struct{}, len(members))

	for _, m := range members {
		if !playerConnected(m) {
			slog.Error("player not added to round team", "player", m)
			continue
		}

		cash += m.Money()
		survivors[m.SteamID64] = struct{}{}
	}

	return &roundTeam{
		Cash:      cash,
		Side:      side,
		Survivors: survivors,
	}
}

func (rt *roundTeam) roundEnd(ts *common.TeamState) {
	rt.CashSpend = ts.MoneySpentThisRound()
	rt.EqValue = ts.RoundStartEquipmentValue()
}

type roundKill struct {
	Killer        uint64
	Victim        uint64
	Assister      uint64
	SinceStart    time.Duration
	Headshot      bool
	Wallbang      bool
	KillerBlind   bool
	ThroughSmoke  bool
	NoScope       bool
	AssistedFlash bool
	KillerSide    common.Team
	AssisterSide  common.Team
	Weapon        common.EquipmentType
}

func newRoundKill(kill events.Kill, roundStartedAt time.Time) *roundKill {
	k := &roundKill{
		Killer:       kill.Killer.SteamID64,
		KillerSide:   kill.Killer.Team,
		Victim:       kill.Victim.SteamID64,
		Headshot:     kill.IsHeadshot,
		Wallbang:     kill.IsWallBang(),
		KillerBlind:  kill.AttackerBlind,
		ThroughSmoke: kill.ThroughSmoke,
		NoScope:      kill.NoScope,
		SinceStart:   time.Duration(time.Since(roundStartedAt).Seconds()),
		Weapon:       kill.Weapon.Type,
	}

	if playerConnected(kill.Assister) {
		k.Assister = kill.Assister.SteamID64
		k.AssisterSide = kill.Assister.Team
		k.AssistedFlash = kill.AssistedFlash
	}

	return k
}

type gameState struct {
	Rounds     []*round
	teamA      team
	teamB      team
	knifeRound bool
	started    bool
}

func newGameState() *gameState {
	return &gameState{Rounds: make([]*round, 0, 50)}
}

func (gs *gameState) detectKnifeRound(pp []*common.Player) {
	gs.knifeRound = false

	for _, p := range pp {
		weapons := p.Weapons()
		if len(weapons) == 1 && weapons[0].Type == common.EqKnife {
			gs.knifeRound = true
			break
		}
	}

	slog.Info("knife round set to", "knife_round", gs.knifeRound)
}

func (gs *gameState) collectStats() bool {
	if gs.knifeRound || !gs.started {
		return false
	}

	return true
}

func (gs *gameState) startRound(ts *common.TeamState) {
	gs.Rounds = append(gs.Rounds, newRound(ts))
}

var (
	errNoRounds          = errors.New("no rounds found in game state")
	errInvalidVictimSide = errors.New("invalid victim side")
)

// killCount appends kill to last (current) round in game state.
func (gs *gameState) killCount(kill events.Kill) error {
	if len(gs.Rounds) < 1 {
		return fmt.Errorf("kill not counted: %w", errNoRounds)
	}

	// add kill to round kill feed
	currRound := gs.Rounds[len(gs.Rounds)-1]
	currRound.KillFeed = append(currRound.KillFeed, newRoundKill(kill, currRound.StartedAt))

	// remove victim from team survivors list
	switch kill.Victim.Team {
	case currRound.TeamA.Side:
		delete(currRound.TeamA.Survivors, kill.Victim.SteamID64)
	case currRound.TeamB.Side:
		delete(currRound.TeamB.Survivors, kill.Victim.SteamID64)
	default:
		return fmt.Errorf("kill not counted: %w (%d)", errInvalidVictimSide, kill.Victim.Team)
	}

	return nil
}

func (gs *gameState) endRound(r events.RoundEnd) error {
	if len(gs.Rounds) < 1 {
		return errNoRounds
	}

	gs.Rounds[len(gs.Rounds)-1].end(r.WinnerState, r.LoserState, r.Reason)

	return nil
}

type team struct {
	name    string
	flag    string
	players []uint64
	score   int
	side    common.Team
	status  gameStatus
}

func newTeam(name, flag string, side common.Team, pp []*common.Player) team {
	t := team{
		name:    name,
		flag:    flag,
		players: make([]uint64, 0, len(pp)),
		side:    side,
		status:  gameStatusInProgress,
	}

	for _, p := range pp {
		if !playerConnected(p) {
			slog.Info("player not added to team state",
				"steam_id", p.SteamID64,
				"name", p.Name)
			continue
		}
		t.players = append(t.players, p.SteamID64)
	}

	return t
}

func (t *team) swapSide() error {
	slog.Info("swapping team side", "team", t.name, "side", t.side)
	switch t.side {
	case common.TeamCounterTerrorists:
		t.side = common.TeamTerrorists
	case common.TeamTerrorists:
		t.side = common.TeamCounterTerrorists
	default:
		return fmt.Errorf("impossible to swap team with side: %d", t.side)
	}
	return nil
}
