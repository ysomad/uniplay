package demoparser

import (
	"log/slog"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type gameState struct {
	knifeRound bool
}

func (gs *gameState) detectKnifeRound(pp []*common.Player) {
	gs.knifeRound = false

	playersWithKnifeOnly := 0

	for _, p := range pp {
		weapons := p.Weapons()
		if len(weapons) == 1 && weapons[0].Type == common.EqKnife {
			slog.Info("player has only knife", "player", p.Name)
			playersWithKnifeOnly++
		}
	}

	if playersWithKnifeOnly == len(pp) {
		gs.knifeRound = true
	}

	slog.Info("knife round set", "knife_round", gs.knifeRound)
}
