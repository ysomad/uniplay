package demoparser

import (
	"time"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type flashbang struct {
	ThrowerSide common.Team
	Victim      uint64
	RoundNum    int
	Duration    time.Duration
}

type playerFlashbangs map[uint64][]flashbang

func newPlayerFlashbands() playerFlashbangs {
	return make(playerFlashbangs, 20)
}

func (pf playerFlashbangs) add(throwerSteamID uint64, f flashbang) {
	if _, ok := pf[throwerSteamID]; !ok {
		pf[throwerSteamID] = make([]flashbang, 0, 100)
	}

	pf[throwerSteamID] = append(pf[throwerSteamID], f)
}
