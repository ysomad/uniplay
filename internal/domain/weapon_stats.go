package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type WeaponStats struct {
	WeaponID      uint16              `json:"weapon_id"`
	Weapon        string              `json:"weapon"`
	Stats         *WeaponStat         `json:"stats"`
	AccuracyStats *WeaponAccuracyStat `json:"accuracy_stats"`
}

func (s *WeaponStats) SetStats(m Metric, v uint32) {
	s.Stats.SetStat(m, v)
	s.AccuracyStats.SetStat(m, v)
}

// WeaponStat is a set of weapon statistics calculated from sum of metrics.
// Each field corresponds to specific weapon metric.
type WeaponStat struct {
	Assists           uint32 `json:"assists"`
	BlindKills        uint32 `json:"blind_kills"`
	DamageDealt       uint32 `json:"damage_dealt"`
	DamageTaken       uint32 `json:"damage_taken"`
	Deaths            uint32 `json:"deaths"`
	HSKills           uint32 `json:"headshot_kills"`
	Kills             uint32 `json:"kills"`
	NoscopeKills      uint32 `json:"noscope_kills"`
	ThroughSmokeKills uint32 `json:"through_smoke_kills"`
	WallbangKills     uint32 `json:"wallbang_kills"`
}

// SetStat sets v into specific field depends on metric.
func (s *WeaponStat) SetStat(m Metric, v uint32) {
	switch m {
	case MetricDeath:
		s.Deaths = v
	case MetricKill:
		s.Kills = v
	case MetricHSKill:
		s.HSKills = v
	case MetricBlindKill:
		s.BlindKills = v
	case MetricWallbangKill:
		s.WallbangKills = v
	case MetricNoScopeKill:
		s.NoscopeKills = v
	case MetricThroughSmokeKill:
		s.ThroughSmokeKills = v
	case MetricAssist:
		s.Assists = v
	case MetricDamageTaken:
		s.DamageTaken = v
	case MetricDamageDealt:
		s.DamageDealt = v
	}
}

type WeaponAccuracyStat struct {
	Shots    uint32  `json:"shots"`
	Accuracy float64 `json:"accuracy"`

	Head     float64 `json:"head"`
	HeadHits uint32  `json:"head_hits"`

	Chest     float64 `json:"chest"`
	ChestHits uint32  `json:"chest_hits"`

	Stomach     float64 `json:"stomach"`
	StomachHits uint32  `json:"stomach_hits"`

	LeftArm     float64 `json:"left_arm"`
	LeftArmHits uint32  `json:"left_arm_hits"`

	RightArm      float64 `json:"right_arm"`
	RightArmsHits uint32  `json:"right_arms_hits"`

	LeftLeg     float64 `json:"left_leg"`
	LeftLegHits uint32  `json:"left_leg_hits"`

	RightLeg     float64 `json:"right_leg"`
	RightLegHits uint32  `json:"right_leg_hits"`
}

func (s *WeaponAccuracyStat) SetStat(m Metric, v uint32) {
	switch m {
	case MetricShot:
		s.Shots = v
	case MetricHitHead:
		s.HeadHits = v
	case MetricHitChest:
		s.ChestHits = v
	case MetricHitStomach:
		s.StomachHits = v
	case MetricHitLeftArm:
		s.LeftArmHits = v
	case MetricHitRightArm:
		s.RightArmsHits = v
	case MetricHitLeftLeg:
		s.LeftLegHits = v
	case MetricHitRightLeg:
		s.RightLegHits = v
	}
}

type WeaponStatsFilter struct {
	WeaponID      uint16
	WeaponClassID uint8
}

type WeaponClassStats struct {
	ClassID uint8       `json:"class_id"`
	Class   string      `json:"class"`
	Stats   *WeaponStat `json:"stats"`
}

type WeaponStatID struct {
	uuid.UUID
}

func NewWeaponStatID(steamID uint64, weaponID uint16, m Metric) WeaponStatID {
	return WeaponStatID{uuid.NewMD5(uuid.UUID{}, []byte(fmt.Sprintf("%d,%d,%d", steamID, weaponID, m)))}
}
