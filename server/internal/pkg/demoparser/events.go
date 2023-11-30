package demoparser

import (
	"log/slog"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type event uint8

const (
	eventUnknown event = iota
	eventKill
	eventHSKill
	eventBlindKill
	eventWBKill
	eventSmokeKill
	eventNoScopeKill

	eventDeath
	eventAssist
	eventFBAssist
	eventRoundMVP

	eventBecameBlind
	eventBlindedPlayer

	eventBombPlanted
	eventBombDefused

	eventDmgDealt
	eventDmgTaken
	eventDmgGrenadeDealt

	eventShot
	eventHitHead
	eventHitNeck
	eventHitChest
	eventHitStomach
	eventHitArm
	eventHitLeg
)

func (p *parser) hitgroupToEvent(hg events.HitGroup) event {
	switch hg {
	case events.HitGroupHead:
		return eventHitHead
	case events.HitGroupNeck:
		return eventHitNeck
	case events.HitGroupChest:
		return eventHitChest
	case events.HitGroupStomach:
		return eventHitStomach
	case events.HitGroupLeftArm, events.HitGroupRightArm:
		return eventHitArm
	case events.HitGroupLeftLeg, events.HitGroupRightLeg:
		return eventHitLeg
	default:
		slog.Error("unsupported hitgroup", "hitgroup", hg)
		return eventUnknown
	}
}
