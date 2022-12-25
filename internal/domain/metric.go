package domain

type Metric uint8

const (
	MetricDeath Metric = 1

	MetricKill             Metric = 2
	MetricHSKill           Metric = 3
	MetricBlindKill        Metric = 4
	MetricWallbangKill     Metric = 5
	MetricNoScopeKill      Metric = 6
	MetricThroughSmokeKill Metric = 7

	MetricAssist          Metric = 8
	MetricFlashbangAssist Metric = 9

	MetricDamageTaken Metric = 10
	MetricDamageDealt Metric = 11

	MetricBombPlanted Metric = 12
	MetricBombDefused Metric = 13

	MetricRoundMVP Metric = 14

	MetricBlind   Metric = 15 // сколько раз ослепил
	MetricBlinded Metric = 16 // был ослеплен

	MetricShot        Metric = 17
	MetricHitHead     Metric = 18
	MetricHitChest    Metric = 19
	MetricHitStomach  Metric = 20
	MetricHitLeftArm  Metric = 21
	MetricHitRightArm Metric = 22
	MetricHitLeftLeg  Metric = 23
	MetricHitRightLeg Metric = 24

	MetricGrenadeDamageDealt Metric = 25
)
