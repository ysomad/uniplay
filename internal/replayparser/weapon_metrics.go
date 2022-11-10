package replayparser

import (
	"sync"

	"github.com/google/uuid"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type weaponMetric struct {
	weaponName  string
	weaponClass domain.EquipmentClass
}

type weaponMetrics struct {
	mu      sync.RWMutex
	Metrics map[steamID]map[weaponMetric]map[domain.Metric]int
}

func newWeaponMetrics() *weaponMetrics {
	return &weaponMetrics{
		Metrics: make(map[steamID]map[weaponMetric]map[domain.Metric]int),
	}
}

// add adds v into weapon metrics of specific player.
func (p *weaponMetrics) add(steamID64 uint64, wm weaponMetric, m domain.Metric, v int) {
	p.addv(steamID(steamID64), wm, m, v)
}

// incr increments weapon metric of specific player.
func (p *weaponMetrics) incr(steamID64 uint64, wm weaponMetric, m domain.Metric) {
	p.addv(steamID(steamID64), wm, m, 1)
}

func (p *weaponMetrics) addv(sid steamID, wm weaponMetric, m domain.Metric, v int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.Metrics[sid]; !ok {
		p.Metrics[sid] = make(map[weaponMetric]map[domain.Metric]int)
	}

	if _, ok := p.Metrics[sid][wm]; !ok {
		p.Metrics[sid][wm] = make(map[domain.Metric]int)
	}

	p.Metrics[sid][wm][m] += v
}

func (p *weaponMetrics) toDTO(matchID uuid.UUID) []dto.CreateWeaponMetricArgs {
	args := []dto.CreateWeaponMetricArgs{}

	for steamID, wmetrics := range p.Metrics {
		for wm, metrics := range wmetrics {
			for m, v := range metrics {
				args = append(args, dto.CreateWeaponMetricArgs{
					MatchID:       matchID,
					PlayerSteamID: uint64(steamID),
					WeaponName:    wm.weaponName,
					WeaponClass:   wm.weaponClass,
					Metric:        m,
					Value:         int32(v),
				})
			}
		}
	}

	return args
}
