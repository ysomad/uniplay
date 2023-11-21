package demoparser

import (
	"log/slog"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type gameState struct {
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

func (gs *gameState) gameStarted() bool {
	if gs.knifeRound || !gs.started {
		return false
	}

	return true
}
