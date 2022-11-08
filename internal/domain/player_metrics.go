package domain

import (
	"sync"
)

// SteamID represents steam uint64 id.
type SteamID uint64

// PlayerMetrics is a map of player event entries.
type PlayerMetrics struct {
	mx      sync.RWMutex
	Metrics map[SteamID]map[Metric]int
}

func NewPlayerMetrics() *PlayerMetrics {
	return &PlayerMetrics{
		Metrics: make(map[SteamID]map[Metric]int),
	}
}

type PlayerMetricsOut struct {
	Metrics map[SteamID]map[string]int
}

// Get returns stats of specific player with steamID.
func (p *PlayerMetrics) Get(steamID uint64) (map[Metric]int, bool) {
	p.mx.RLock()
	defer p.mx.RUnlock()

	v, ok := p.Metrics[SteamID(steamID)]
	return v, ok
}

// Add n to amount of player metric entries in the stats map of specific player with steamID.
func (p *PlayerMetrics) Add(steamID uint64, m Metric, n int) { p.add(steamID, m, n) }

// Incr increments metric entries count for player with steamID.
func (p *PlayerMetrics) Incr(steamID uint64, m Metric) { p.add(steamID, m, 1) }

func (p *PlayerMetrics) add(steamID uint64, m Metric, n int) {
	p.mx.Lock()
	defer p.mx.Unlock()

	if _, ok := p.Metrics[SteamID(steamID)]; !ok {
		p.Metrics[SteamID(steamID)] = make(map[Metric]int)
	}

	p.Metrics[SteamID(steamID)][m] += n
}

func (p *PlayerMetrics) Out() *PlayerMetricsOut {
	out := make(map[SteamID]map[string]int)

	for steamID, metrics := range p.Metrics {
		for metric, val := range metrics {
			if _, ok := out[steamID]; !ok {
				out[steamID] = make(map[string]int)
			}

			out[steamID][metric.String()] = val
		}
	}

	return &PlayerMetricsOut{
		Metrics: out,
	}
}
