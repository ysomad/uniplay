package replayparser

import (
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type weaponMetric struct {
	weaponName  string
	weaponClass common.EquipmentClass
}

type weaponMetrics struct {
	metrics map[steamID]map[weaponMetric]map[domain.Metric]int
}

func newWeaponMetrics() *weaponMetrics {
	return &weaponMetrics{
		metrics: make(map[steamID]map[weaponMetric]map[domain.Metric]int),
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
	if _, ok := p.metrics[sid]; !ok {
		p.metrics[sid] = make(map[weaponMetric]map[domain.Metric]int)
	}

	if _, ok := p.metrics[sid][wm]; !ok {
		p.metrics[sid][wm] = make(map[domain.Metric]int)
	}

	p.metrics[sid][wm][m] += v
}

// TODO: refactor with goroutines
func (p *weaponMetrics) toDTO(matchID domain.MatchID) []dto.WeaponMetric {
	args := []dto.WeaponMetric{}

	for steamID, wmetrics := range p.metrics {
		for wm, metrics := range wmetrics {
			for m, v := range metrics {
				args = append(args, dto.WeaponMetric{
					MatchID:       matchID,
					PlayerSteamID: uint64(steamID),
					WeaponName:    wm.weaponName,
					WeaponClass:   domain.EquipmentClass(wm.weaponClass),
					Metric:        m,
					Value:         int32(v),
				})
			}
		}
	}

	return args
}
