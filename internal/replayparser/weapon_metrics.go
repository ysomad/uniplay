package replayparser

import (
	"sync"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type weapon string

type weaponMetrics struct {
	mu      sync.RWMutex
	Metrics map[steamID]map[weapon]map[domain.Metric]int
}

func newWeaponMetrics() *weaponMetrics {
	return &weaponMetrics{
		Metrics: make(map[steamID]map[weapon]map[domain.Metric]int),
	}
}

// add adds v into weapon metrics of specific player.
func (p *weaponMetrics) add(steamID64 uint64, weaponName string, m domain.Metric, v int) {
	p.addv(steamID(steamID64), weapon(weaponName), m, v)
}

// incr increments weapon metric of specific player.
func (p *weaponMetrics) incr(steamID64 uint64, weaponName string, m domain.Metric) {
	p.addv(steamID(steamID64), weapon(weaponName), m, 1)
}

func (p *weaponMetrics) addv(sid steamID, w weapon, m domain.Metric, v int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.Metrics[sid]; !ok {
		p.Metrics[sid] = make(map[weapon]map[domain.Metric]int)
	}

	if _, ok := p.Metrics[sid][w]; !ok {
		p.Metrics[sid][w] = make(map[domain.Metric]int)
	}

	p.Metrics[sid][w][m] += v
}
