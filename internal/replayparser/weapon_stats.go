package replayparser

import (
	"errors"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type weaponMetric struct {
	eqType  common.EquipmentType
	eqClass common.EquipmentClass
}

type weaponStats struct {
	stats map[uint64]map[weaponMetric]map[domain.Metric]int
}

func newWeaponStats() *weaponStats {
	return &weaponStats{
		stats: make(map[uint64]map[weaponMetric]map[domain.Metric]int),
	}
}

// add adds v into weapon metrics of specific player.
func (s *weaponStats) add(steamID uint64, wm weaponMetric, m domain.Metric, v int) {
	s.addv(steamID, wm, m, v)
}

// incr increments weapon metric of specific player.
func (s *weaponStats) incr(steamID uint64, wm weaponMetric, m domain.Metric) {
	s.addv(steamID, wm, m, 1)
}

func (s *weaponStats) addv(steamID uint64, wm weaponMetric, m domain.Metric, v int) {
	if _, ok := s.stats[steamID]; !ok {
		s.stats[steamID] = make(map[weaponMetric]map[domain.Metric]int)
	}

	if _, ok := s.stats[steamID][wm]; !ok {
		s.stats[steamID][wm] = make(map[domain.Metric]int)
	}

	s.stats[steamID][wm][m] += v
}

func (s *weaponStats) toDTO() ([]dto.WeaponStat, error) {
	if len(s.stats) == 0 {
		return nil, errors.New("empty list of weapon metrics")
	}

	var out []dto.WeaponStat

	for steamID, stats := range s.stats {
		for wm, weaponMetrics := range stats {
			for m, v := range weaponMetrics {
				// skip unknown weapons or empty metric values
				if wm.eqClass == common.EqClassUnknown || wm.eqType == common.EqUnknown || v <= 0 {
					continue
				}

				// skip weapon fire for grenades and equipment
				if m == domain.MetricShot && (wm.eqClass == common.EqClassGrenade || wm.eqClass == common.EqClassEquipment) {
					continue
				}

				out = append(out, dto.NewWeaponStat(steamID, uint16(wm.eqType), m, uint32(v)))
			}
		}
	}

	return out, nil
}
