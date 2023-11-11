package demoparser

import (
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

func playerConnected(p *common.Player) bool {
	if p == nil {
		return false
	}
	if p.SteamID64 == 0 ||
		p.UserID == 0 ||
		!p.IsConnected ||
		p.IsBot ||
		p.IsUnknown {
		return false
	}
	return true
}

func playerSpectator(p *common.Player) bool {
	if p == nil {
		return true
	}
	return p.Team == common.TeamSpectators || p.Team == common.TeamUnassigned
}
