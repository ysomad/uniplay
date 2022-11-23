package domain

import "github.com/google/uuid"

type Metric uint8

const (
	// TODO: refactor to iota
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

	MetricRoundMVPCount Metric = 14

	MetricBlind   Metric = 15 // сколько раз ослепил
	MetricBlinded Metric = 16 // был ослеплен
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
	MetricBlind:            "blinded_players",
	MetricBlinded:          "was_blinded",
}

func (m Metric) String() string {
	if _, ok := strMetrics[m]; !ok {
		return "undefined_metric"
	}
	return strMetrics[m]
}

type MetricModel struct {
	MatchID       uuid.UUID
	PlayerSteamID uint64
	Metric        Metric
	Value         int32
}
