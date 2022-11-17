package replayparser

import (
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

// steamID represents steam uint64 id.
type steamID uint64

// playerMetrics is a map of player metrics.
type playerMetrics struct {
	metrics map[steamID]map[domain.Metric]int
}

func newPlayerMetrics() *playerMetrics {
	return &playerMetrics{
		metrics: make(map[steamID]map[domain.Metric]int),
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
	if _, ok := p.metrics[sid]; !ok {
		p.metrics[sid] = make(map[domain.Metric]int)
	}

	p.metrics[sid][m] += n
}

// TODO: refactor with goroutines
func (p *playerMetrics) toDTO(matchID domain.MatchID) []dto.Metric {
	args := []dto.Metric{}

	for steamID, metrics := range p.metrics {
		for m, v := range metrics {
			args = append(args, dto.Metric{
				MatchID:       matchID,
				PlayerSteamID: uint64(steamID),
				Metric:        m,
				Value:         int32(v),
			})
		}
	}

	return args
}
