package domain

import (
	"sync"

	"github.com/ssssargsian/uniplay/internal/metric"
)

// SteamID represents steam uint64 id.
type SteamID uint64

// metrics is a map of metric, key is metric, value is amount of entries of the event.
type metrics map[metric.Metric]uint16

// PlayerMetrics is a map of player event entries.
type PlayerMetrics struct {
	mx      sync.RWMutex
	metrics map[SteamID]metrics
}

func NewPlayerMetrics() *PlayerMetrics {
	return &PlayerMetrics{
		metrics: make(map[SteamID]metrics),
	}
}

// Get returns stats of specific player with steamID.
func (p *PlayerMetrics) Get(steamID uint64) (metrics, bool) {
	p.mx.RLock()
	defer p.mx.RUnlock()

	v, ok := p.metrics[SteamID(steamID)]
	return v, ok
}

// Add n to amount of player metric entries in the stats map of specific player with steamID.
func (p *PlayerMetrics) Add(steamID uint64, m metric.Metric, n uint16) { p.add(steamID, m, n) }

// Incr increments metric entries count for player with steamID.
func (p *PlayerMetrics) Incr(steamID uint64, m metric.Metric) { p.add(steamID, m, 1) }

func (p *PlayerMetrics) add(steamID uint64, m metric.Metric, n uint16) {
	p.mx.Lock()
	defer p.mx.Unlock()

	if _, ok := p.metrics[SteamID(steamID)]; !ok {
		p.metrics[SteamID(steamID)] = make(metrics)
	}

	p.metrics[SteamID(steamID)][m] += n
}

type weaponEvents []metric.WeaponEvent

type PlayerWeaponEvents struct {
	mx      sync.RWMutex
	metrics map[SteamID]weaponEvents
}

func NewPlayerWeaponEvents() *PlayerWeaponEvents {
	return &PlayerWeaponEvents{
		metrics: make(map[SteamID]weaponEvents),
	}
}

func (w *PlayerWeaponEvents) Add(steamID uint64, m metric.WeaponEvent) {
	w.mx.Lock()
	defer w.mx.Unlock()

	if _, ok := w.metrics[SteamID(steamID)]; !ok {
		w.metrics[SteamID(steamID)] = []metric.WeaponEvent{}
	}

	w.metrics[SteamID(steamID)] = append(w.metrics[SteamID(steamID)], m)
}
