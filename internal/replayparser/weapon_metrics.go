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

// weaponClass returns domain weapon class separated by shotguns, machine guns and others.
func (p *weaponMetrics) weaponClass(t common.EquipmentType) domain.WeaponClassID {
	switch t {
	case common.EqP2000,
		common.EqGlock,
		common.EqP250,
		common.EqDeagle,
		common.EqFiveSeven,
		common.EqDualBerettas,
		common.EqTec9,
		common.EqCZ,
		common.EqUSP,
		common.EqRevolver:
		return domain.ClassPistol
	case common.EqMP7,
		common.EqMP9,
		common.EqBizon,
		common.EqMac10,
		common.EqUMP,
		common.EqP90,
		common.EqMP5:
		return domain.ClassSMG
	case common.EqSawedOff,
		common.EqNova,
		common.EqSwag7,
		common.EqXM1014:
		return domain.ClassShotgun
	case common.EqM249, common.EqNegev:
		return domain.ClassMachineGun
	case common.EqGalil,
		common.EqFamas,
		common.EqAK47,
		common.EqM4A4,
		common.EqM4A1,
		common.EqSG553,
		common.EqAUG:
		return domain.ClassAssaultRifle
	case common.EqAWP, common.EqScar20, common.EqG3SG1, common.EqSSG08:
		return domain.ClassSniperRifle
	case common.EqZeus, common.EqBomb, common.EqKnife:
		return domain.ClassEquipment
	case common.EqWorld:
		return domain.ClassOther
	case common.EqDecoy,
		common.EqMolotov,
		common.EqIncendiary,
		common.EqFlash,
		common.EqSmoke,
		common.EqHE:
		return domain.ClassGrenade
	}

	return 0
}

// TODO: refactor with goroutines
func (p *weaponMetrics) toDTO(matchID domain.MatchID) ([]dto.WeaponMetric, error) {
	if len(p.metrics) == 0 {
		return nil, errors.New("empty list of weapon metrics")
	}

	out := []dto.WeaponMetric{}

	for steamID, wmetrics := range p.metrics {
		for wm, metrics := range wmetrics {
			for m, v := range metrics {
				if wm.eqClass == common.EqClassUnknown {
					continue
				}

				out = append(out, dto.WeaponMetric{
					MatchID:       matchID,
					PlayerSteamID: uint64(steamID),
					WeaponID:      uint16(wm.eqType),
					Metric:        m,
					Value:         int32(v),
				})
			}
		}
	}

	return out, nil
}
