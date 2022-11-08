package domain

import (
	"sync"
)

type Weapon string

type WeaponMetrics struct {
	mx      sync.RWMutex
	Metrics map[SteamID]map[Weapon]map[Metric]int
}

func NewWeaponMetrics() *WeaponMetrics {
	return &WeaponMetrics{
		Metrics: make(map[SteamID]map[Weapon]map[Metric]int),
	}
}

func (p *WeaponMetrics) Add(sid SteamID, w Weapon, m Metric, v int) { p.add(sid, w, m, v) }
func (p *WeaponMetrics) Incr(sid SteamID, w Weapon, m Metric)       { p.add(sid, w, m, 1) }

func (p *WeaponMetrics) add(sid SteamID, w Weapon, m Metric, v int) {
	p.mx.Lock()
	defer p.mx.Unlock()

	if _, ok := p.Metrics[sid]; !ok {
		p.Metrics[sid] = make(map[Weapon]map[Metric]int)
	}

	if _, ok := p.Metrics[sid][w]; !ok {
		p.Metrics[sid][w] = make(map[Metric]int)
	}

	p.Metrics[sid][w][m] += v
}
