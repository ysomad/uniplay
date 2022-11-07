package domain

type metric uint8

const (
	MetricDeath metric = iota + 1

	MetricKill
	MetricHSKill
	MetricBlindKill
	MetricWallbangKill
	MetricNoScopeKill
	MetricThroughSmokeKill

	MetricAssist
	MetricFlashbangAssist

	MetricRountMVPCount

	MetricDamageTaken
	MetricDamageDealt
)
