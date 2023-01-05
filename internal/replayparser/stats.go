package replayparser

import (
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type stats struct {
	playerStats map[uint64]*dto.PlayerStat
	weaponStats map[uint64]map[common.EquipmentType]*dto.PlayerWeaponStat
}

func newStats() stats {
	return stats{
		playerStats: make(map[uint64]*dto.PlayerStat),
		weaponStats: make(map[uint64]map[common.EquipmentType]*dto.PlayerWeaponStat),
	}
}

func (s *stats) addPlayerStat(steamID uint64, m domain.Metric, v int) {
	if m <= 0 || v <= 0 {
		return
	}

	if _, ok := s.playerStats[steamID]; !ok {
		s.playerStats[steamID] = &dto.PlayerStat{SteamID: steamID}
	}

	s.playerStats[steamID].Add(m, v)
}

func (s *stats) incrPlayerStat(steamID uint64, m domain.Metric) {
	s.addPlayerStat(steamID, m, 1)
}

func (s *stats) addWeaponStat(steamID uint64, m domain.Metric, e common.EquipmentType, v int) {
	if e.Class() == common.EqClassUnknown || e == common.EqUnknown || v <= 0 || m <= 0 {
		return
	}

	if _, ok := s.weaponStats[steamID]; !ok {
		s.weaponStats[steamID] = make(map[common.EquipmentType]*dto.PlayerWeaponStat)
	}

	if _, ok := s.weaponStats[steamID][e]; !ok {
		s.weaponStats[steamID][e] = &dto.PlayerWeaponStat{
			SteamID:  steamID,
			WeaponID: int16(e),
		}
	}

	s.weaponStats[steamID][e].AddStat(m, v)
}

func (s *stats) incrWeaponStat(steamID uint64, m domain.Metric, e common.EquipmentType) {
	s.addWeaponStat(steamID, m, e, 1)
}
