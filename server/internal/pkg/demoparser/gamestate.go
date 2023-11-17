package demoparser

import (
	"log/slog"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type gameState struct {
	// Rounds     []*round
	// teamA      team
	// teamB      team
	knifeRound bool
	started    bool
}

// func newGameState() *gameState {
// 	return &gameState{Rounds: make([]*round, 0, 50)}
// }

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

// func (gs *gameState) startRound(ts *common.TeamState) {
// 	gs.Rounds = append(gs.Rounds, newRound(ts))
// }

// killCount appends kill to last (current) round in game state.
// func (gs *gameState) killCount(kill events.Kill) error {
// 	if len(gs.Rounds) < 1 {
// 		return fmt.Errorf("kill not counted: %w", errNoRounds)
// 	}

// 	// add kill to round kill feed
// 	currRound := gs.Rounds[len(gs.Rounds)-1]
// 	currRound.KillFeed = append(currRound.KillFeed, newRoundKill(kill, currRound.StartedAt))

// 	// remove victim from team survivors list
// 	switch kill.Victim.Team {
// 	case currRound.TeamA.Side:
// 		delete(currRound.TeamA.Survivors, kill.Victim.SteamID64)
// 	case currRound.TeamB.Side:
// 		delete(currRound.TeamB.Survivors, kill.Victim.SteamID64)
// 	default:
// 		return fmt.Errorf("kill not counted: %w (%d)", errInvalidVictimSide, kill.Victim.Team)
// 	}

// 	return nil
// }

// func (gs *gameState) endRound(r events.RoundEnd) error {
// 	if len(gs.Rounds) < 1 {
// 		return errNoRounds
// 	}

// 	gs.Rounds[len(gs.Rounds)-1].end(r.WinnerState, r.LoserState, r.Reason)

// 	return nil
// }

// type team struct {
// 	name    string
// 	flag    string
// 	players []uint64
// 	score   int
// 	side    common.Team
// 	status  gameStatus
// }

// func newTeam(name, flag string, side common.Team, pp []*common.Player) team {
// 	t := team{
// 		name:    name,
// 		flag:    flag,
// 		players: make([]uint64, 0, len(pp)),
// 		side:    side,
// 		status:  gameStatusInProgress,
// 	}

// 	for _, p := range pp {
// 		if !playerConnected(p) {
// 			slog.Info("player not added to team state",
// 				"steam_id", p.SteamID64,
// 				"name", p.Name)
// 			continue
// 		}
// 		t.players = append(t.players, p.SteamID64)
// 	}

// 	return t
// }

// func (t *team) swapSide() error {
// 	slog.Info("swapping team side", "team", t.name, "side", t.side)
// 	switch t.side {
// 	case common.TeamCounterTerrorists:
// 		t.side = common.TeamTerrorists
// 	case common.TeamTerrorists:
// 		t.side = common.TeamCounterTerrorists
// 	default:
// 		return fmt.Errorf("impossible to swap team with side: %d", t.side)
// 	}
// 	return nil
// }
