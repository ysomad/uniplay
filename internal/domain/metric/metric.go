package metric

type Metric uint8

const (
	Death Metric = iota + 1

	Kill
	HSKill
	BlindKill
	WallbangKill
	NoScopeKill
	ThroughSmokeKill

	Assist
	FlashbangAssist

	DamageTaken
	DamageDealt

	BombPlanted
	BombDefused

	RoundMVPCount
)

var strMetrics = map[Metric]string{
	Death:            "deaths",
	Kill:             "kills",
	HSKill:           "headshot_kills",
	BlindKill:        "blind_kills",
	WallbangKill:     "wallbang_kills",
	NoScopeKill:      "noscope_kills",
	ThroughSmokeKill: "through_smoke_kills",
	Assist:           "assists",
	FlashbangAssist:  "flashbang_assists",
	DamageTaken:      "damage_taken",
	DamageDealt:      "damage_dealt",
	BombPlanted:      "bomb_planted",
	BombDefused:      "bomb_defused",
	RoundMVPCount:    "round_mvps",
}

func (m Metric) String() string {
	if _, ok := strMetrics[m]; !ok {
		return "undefined"
	}
	return strMetrics[m]
}
