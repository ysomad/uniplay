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

func (p *WeaponMetrics) Add(steamID64 uint64, w Weapon, m Metric, v int) { p.add(steamID64, w, m, v) }
func (p *WeaponMetrics) Incr(steamID64 uint64, w Weapon, m Metric)       { p.add(steamID64, w, m, 1) }

func (p *WeaponMetrics) add(steamID64 uint64, w Weapon, m Metric, v int) {
	p.mx.Lock()
	defer p.mx.Unlock()

	steamID := SteamID(steamID64)

	if _, ok := p.Metrics[steamID]; !ok {
		p.Metrics[steamID] = make(map[Weapon]map[Metric]int)
	}

	if _, ok := p.Metrics[steamID][w]; !ok {
		p.Metrics[steamID][w] = make(map[Metric]int)
	}

	p.Metrics[steamID][w][m] += v
}
