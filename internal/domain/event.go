package domain

type Event uint8

const (
	EventDeath = iota

	EventKill
	EventHSKill
	EventBlindKill
	EventWallbangKill
	EventNoScopeKill
	EventThroughSmokeKill

	EventAssist
	EventFlashAssist
)
