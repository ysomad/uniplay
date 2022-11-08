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

// Get returns stats of specific player with steamID.
func (p *PlayerMetrics) Get(steamID uint64) (map[Metric]int, bool) {
	p.mx.RLock()
	defer p.mx.RUnlock()

	v, ok := p.Metrics[SteamID(steamID)]
	return v, ok
}

// Add n to amount of player metric entries in the stats map of specific player with steamID.
func (p *PlayerMetrics) Add(sid SteamID, m Metric, n int) { p.add(sid, m, n) }

// Incr increments metric entries count for player with steamID.
func (p *PlayerMetrics) Incr(sid SteamID, m Metric) { p.add(sid, m, 1) }

func (p *PlayerMetrics) add(sid SteamID, m Metric, n int) {
	p.mx.Lock()
	defer p.mx.Unlock()

	if _, ok := p.Metrics[sid]; !ok {
		p.Metrics[sid] = make(map[Metric]int)
	}

	p.Metrics[sid][m] += n
}
