package domain

type Metric int8

const (
	MetricDeath Metric = iota + 1
	MetricKill
	MetricHSKill
	MetricBlindKill
	MetricWallbangKill
	MetricNoScopeKill
	MetricThroughSmokeKill
	MetricAssist
	MetricFlashbangAssist
	MetricDamageTaken
	MetricDamageDealt
	MetricBombPlanted
	MetricBombDefused
	MetricRoundMVP
	MetricBlind
	MetricBlinded
	MetricShot
	MetricHitHead
	MetricHitChest
	MetricHitStomach
	MetricHitLeftArm
	MetricHitRightArm
	MetricHitLeftLeg
	MetricHitRightLeg
	MetricGrenadeDamageDealt
)
