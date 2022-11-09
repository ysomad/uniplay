package domain

import "github.com/google/uuid"

type Metric uint8

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

	MetricRoundMVPCount
)

var strMetrics = map[Metric]string{
	MetricDeath:            "deaths",
	MetricKill:             "kills",
	MetricHSKill:           "headshot_kills",
	MetricBlindKill:        "blind_kills",
	MetricWallbangKill:     "wallbang_kills",
	MetricNoScopeKill:      "noscope_kills",
	MetricThroughSmokeKill: "through_smoke_kills",
	MetricAssist:           "assists",
	MetricFlashbangAssist:  "flashbang_assists",
	MetricDamageTaken:      "damage_taken",
	MetricDamageDealt:      "damage_dealt",
	MetricBombPlanted:      "bomb_planted",
	MetricBombDefused:      "bomb_defused",
	MetricRoundMVPCount:    "round_mvps",
}

func (m Metric) String() string {
	if _, ok := strMetrics[m]; !ok {
		return "undefined"
	}
	return strMetrics[m]
}

type MetricModel struct {
	MatchID       uuid.UUID
	PlayerSteamID uint64
	Metric        Metric
	Value         int32
}
