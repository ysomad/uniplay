package demoparser

import (
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

func playerConnected(pl *common.Player) bool {
	if pl == nil {
		return false
	}
	if pl.SteamID64 == 0 ||
		pl.UserID == 0 ||
		!pl.IsConnected ||
		pl.IsBot ||
		pl.IsUnknown {
		return false
	}
	return true
}

func playerSpectator(pl *common.Player) bool {
	if pl == nil {
		return true
	}
	return pl.Team == common.TeamSpectators || pl.Team == common.TeamUnassigned
}
