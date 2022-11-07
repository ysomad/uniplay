package domain

import (
	"sync"

	"github.com/ssssargsian/uniplay/internal/domain/metric"
)

type WeaponMetric struct {
	Metric        metric.Metric `json:"metric"`
	Weapon        string        `json:"weapon"`
	Value         int           `json:"value"`
	IsValueDamage bool          `json:"is_value_damage"`
}

type WeaponMetricOut struct {
	Metric        string `json:"metric"`
	Weapon        string `json:"weapon"`
	Value         int    `json:"value"`
	IsValueDamage bool   `json:"is_value_damage"`
}

type WeaponMetricsOut struct {
	Metrics map[SteamID][]WeaponMetricOut `json:"metrics,omitempty"`
}

type WeaponMetrics struct {
	mx      sync.RWMutex               `json:"-"`
	Metrics map[SteamID][]WeaponMetric `json:"metrics"`
}

func NewWeaponMetrics() *WeaponMetrics {
	return &WeaponMetrics{
		Metrics: make(map[SteamID][]WeaponMetric),
	}
}

func (p *WeaponMetrics) Add(steamID uint64, m WeaponMetric) {
	p.mx.Lock()
	defer p.mx.Unlock()

	if _, ok := p.Metrics[SteamID(steamID)]; !ok {
		p.Metrics[SteamID(steamID)] = []WeaponMetric{}
	}

	p.Metrics[SteamID(steamID)] = append(p.Metrics[SteamID(steamID)], m)
}

func (p *WeaponMetrics) Out() *WeaponMetricsOut {
	out := make(map[SteamID][]WeaponMetricOut)

	for steamID, metrics := range p.Metrics {
		for i, val := range metrics {
			if _, ok := out[steamID]; !ok {
				out[steamID] = make([]WeaponMetricOut, len(metrics))
			}

			out[steamID][i] = WeaponMetricOut{
				Metric:        val.Metric.String(),
				Weapon:        val.Weapon,
				Value:         val.Value,
				IsValueDamage: val.IsValueDamage,
			}
		}
	}

	return &WeaponMetricsOut{
		Metrics: out,
	}
}
