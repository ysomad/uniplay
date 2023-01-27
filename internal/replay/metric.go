package replay

type metric int8

const (
	metricDeath metric = iota + 1
	metricKill
	metricHSKill
	metricBlindKill
	metricWallbangKill
	metricNoScopeKill
	metricThroughSmokeKill
	metricAssist
	metricFlashbangAssist
	metricDamageTaken
	metricDamageDealt
	metricBombPlanted
	metricBombDefused
	metricRoundMVP
	metricBlind
	metricBlinded
	metricShot
	metricHitHead
	metricHitNeck
	metricHitChest
	metricHitStomach
	metricHitLeftArm
	metricHitRightArm
	metricHitLeftLeg
	metricHitRightLeg
	metricGrenadeDamageDealt
)
