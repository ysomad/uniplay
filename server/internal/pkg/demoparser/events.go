package demoparser

import "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"

type event uint8

const (
	eventKill event = iota + 1
	eventHSKill
	eventBlindKill
	eventWBKill
	eventSmokeKill

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

func (p *parser) hitgroupToEvent(g events.HitGroup) event {
	switch g {
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
		return 0
	}
}
