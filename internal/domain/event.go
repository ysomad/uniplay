package domain

type Event uint8

const (
	eventUndefined = iota

	EventDeath

	EventKill
	EventHSKill
	EventBlindKill
	EventWallbangKill
	EventNoScopeKill
	EventThroughSmokeKill

	EventAssist
	EventFlashbangAssist // assist with flashbang
)

type EventSlug string

func (e Event) Slug() EventSlug {
	switch e {
	case EventDeath:
		return "deaths"
	case EventKill:
		return "kills"
	case EventHSKill:
		return "headshot_kills"
	case EventBlindKill:
		return "blind_kills"
	case EventWallbangKill:
		return "wallbang_kills"
	case EventNoScopeKill:
		return "noscope_kills"
	case EventThroughSmokeKill:
		return "through_smoke_kills"
	case EventAssist:
		return "assists"
	case EventFlashbangAssist:
		return "flashbang_assists"
	}

	return "undefined_events"
}
