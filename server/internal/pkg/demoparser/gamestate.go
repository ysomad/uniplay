package demoparser

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type roundReason uint8

const (
	roundReasonBomb roundReason = iota + 1
	roundReasonDefused
	roundReasonElimination
)

type gameStatus int8

const (
	gameStatusUnknown gameStatus = 0
	gameStatusDraw    gameStatus = 1
	gameStatusLose    gameStatus = -1
	gameStatusWin     gameStatus = 2
)

type round struct {
	killFeed   []*roundKill
	cash       int
	cashSpend  int
	equipValue int
	survivorsA int
	survivorsB int
	reason     roundReason
}

type roundKill struct {
	killer        uint64
	victim        uint64
	assister      uint64
	headshot      bool
	wallbang      bool
	blind         bool
	throughSmoke  bool
	noScope       bool
	assistedFlash bool
	sinceStart    time.Duration // time passed since round start
}

type gameState struct {
	rounds     []round
	teamA      team
	teamB      team
	knifeRound bool
	started    bool
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
		status:  gameStatusUnknown,
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
