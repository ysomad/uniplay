package replayparser

import (
	"sync"

	"github.com/ssssargsian/uniplay/internal/domain"
)

// steamID represents steam uint64 id.
type steamID uint64

// playerMetrics is a map of player metrics.
type playerMetrics struct {
	mu      sync.RWMutex
	Metrics map[steamID]map[domain.Metric]int
}

func newPlayerMetrics() *playerMetrics {
	return &playerMetrics{
		Metrics: make(map[steamID]map[domain.Metric]int),
	}
}

// add adds n to player metric.
func (p *playerMetrics) add(steamID64 uint64, m domain.Metric, n int) {
	p.addn(steamID(steamID64), m, n)
}

// incr increments player metric.
func (p *playerMetrics) incr(steamID64 uint64, m domain.Metric) {
	p.addn(steamID(steamID64), m, 1)
}

func (p *playerMetrics) addn(sid steamID, m domain.Metric, n int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.Metrics[sid]; !ok {
		p.Metrics[sid] = make(map[domain.Metric]int)
	}

	p.Metrics[sid][m] += n
}
