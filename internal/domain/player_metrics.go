package domain

import "sync"

// SteamID represents steam uint64 id.
type SteamID uint64

// metrics is a map of metric, key is metric, value is amount of entries of the event.
type metrics map[metric]uint16

// playerMetrics is a map of player event entries.
type playerMetrics struct {
	mx      sync.RWMutex
	metrics map[SteamID]metrics
}

func NewPlayerMetrics() *playerMetrics {
	return &playerMetrics{
		metrics: make(map[SteamID]metrics),
	}
}

// Get returns stats of specific player with steamID.
func (p *playerMetrics) Get(steamID uint64) (metrics, bool) {
	p.mx.RLock()
	defer p.mx.RUnlock()

	v, ok := p.metrics[SteamID(steamID)]
	return v, ok
}

// Add n to amount of player metric entries in the stats map of specific player with steamID.
func (p *playerMetrics) Add(steamID uint64, m metric, n uint16) { p.add(steamID, m, n) }

// Incr increments metric entries count for playey with steamID.
func (p *playerMetrics) Incr(steamID uint64, m metric) { p.add(steamID, m, 1) }

func (p *playerMetrics) add(steamID uint64, m metric, n uint16) {
	p.mx.Lock()
	defer p.mx.Unlock()

	if _, ok := p.metrics[SteamID(steamID)]; !ok {
		p.metrics[SteamID(steamID)] = make(metrics)
	}

	p.metrics[SteamID(steamID)][m] += n
}
