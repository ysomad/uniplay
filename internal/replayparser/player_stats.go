package replayparser

import (
	"errors"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

// playerStats is a map of player metrics.
type playerStats struct {
	stats map[uint64]map[domain.Metric]int
}

func newPlayerStats() *playerStats {
	return &playerStats{
		stats: make(map[uint64]map[domain.Metric]int),
	}
}

// add adds n to player metric.
func (s *playerStats) add(steamID uint64, m domain.Metric, n int) {
	s.addn(steamID, m, n)
}

// incr increments player metric.
func (s *playerStats) incr(steamID uint64, m domain.Metric) {
	s.addn(steamID, m, 1)
}

func (s *playerStats) addn(steamID uint64, m domain.Metric, n int) {
	if _, ok := s.stats[steamID]; !ok {
		s.stats[steamID] = make(map[domain.Metric]int)
	}

	s.stats[steamID][m] += n
}

func (s *playerStats) toDTO() ([]dto.PlayerStat, error) {
	if len(s.stats) == 0 {
		return nil, errors.New("empty list of metrics")
	}

	var ps []dto.PlayerStat
	for steamID, metrics := range s.stats {
		for m, v := range metrics {
			if v <= 0 {
				continue
			}

			ps = append(ps, dto.NewPlayerStat(steamID, m, uint32(v)))
		}
	}

	return ps, nil
}
