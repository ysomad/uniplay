package match

import (
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/stat"
)

type stats struct {
	playerStats map[uint64]*playerStat
	weaponStats map[uint64]map[common.EquipmentType]*weaponStat
}

func newStats() stats {
	return stats{
		playerStats: make(map[uint64]*playerStat),
		weaponStats: make(map[uint64]map[common.EquipmentType]*weaponStat),
	}
}

func (s *stats) normalize(roundsPlayed int8) ([]*playerStat, []*weaponStat) {
	playerStats := make([]*playerStat, 0, len(s.playerStats))

	for _, ps := range s.playerStats {
		ps.calculate(int32(roundsPlayed))
		playerStats = append(playerStats, ps)
	}

	var weaponStats []*weaponStat

	for _, weapons := range s.weaponStats {
		for _, ws := range weapons {
			ws.calculateTotalHits()
			ws.accuracy = domain.NewWeaponAccuracyStats(
				ws.shots,
				ws.totalHits,
				ws.headHits,
				ws.neckHits,
				ws.chestHits,
				ws.stomachHits,
				ws.armHits,
				ws.legHits,
			)
			weaponStats = append(weaponStats, ws)
		}
	}

	return playerStats, weaponStats
}

func (s *stats) addPlayerStat(steamID uint64, m metric, v int) {
	if !s.validMetric(m, v) {
		return
	}

	if _, ok := s.playerStats[steamID]; !ok {
		s.playerStats[steamID] = &playerStat{steamID: steamID}
	}

	s.playerStats[steamID].add(m, v)
}

func (s *stats) incrPlayerStat(steamID uint64, m metric) {
	s.addPlayerStat(steamID, m, 1)
}

// validWeapon checks whether weapon is valid or not.
func (s *stats) validWeapon(e common.EquipmentType) bool {
	if e.Class() == common.EqClassUnknown || e == common.EqUnknown {
		return false
	}

	return true
}

// validMetric checks whether metric and value is valid or not.
func (s *stats) validMetric(m metric, v int) bool {
	if v <= 0 || m <= 0 {
		return false
	}

	return true
}

func (s *stats) addWeaponStat(steamID uint64, m metric, e common.EquipmentType, v int) {
	if !s.validWeapon(e) || !s.validMetric(m, v) {
		return
	}

	if _, ok := s.weaponStats[steamID]; !ok {
		s.weaponStats[steamID] = make(map[common.EquipmentType]*weaponStat)
	}

	if _, ok := s.weaponStats[steamID][e]; !ok {
		s.weaponStats[steamID][e] = &weaponStat{
			steamID:  steamID,
			weaponID: int32(e),
		}
	}

	s.weaponStats[steamID][e].add(m, v)
}

func (s *stats) incrWeaponStat(steamID uint64, m metric, e common.EquipmentType) {
	s.addWeaponStat(steamID, m, e, 1)
}

type playerStat struct {
	steamID            uint64
	kills              int32
	hsKills            int32
	blindKills         int32
	wallbangKills      int32
	noScopeKills       int32
	throughSmokeKills  int32
	deaths             int32
	assists            int32
	flashbangAssists   int32
	mvpCount           int32
	damageTaken        int32
	damageDealt        int32
	grenadeDamageDealt int32
	blindedPlayers     int32
	blindedTimes       int32
	bombsPlanted       int32
	bombsDefused       int32

	hsPercentage           float64
	kd                     float64
	adr                    float64
	killsPerRound          float64
	assistsPerRound        float64
	deathsPerRound         float64
	blindedPlayersPerRound float64
}

func (s *playerStat) calculate(rounds int32) {
	s.hsPercentage = stat.HeadshotPercentage(s.hsKills, s.kills)
	s.kd = stat.KD(s.kills, s.deaths)
	s.adr = stat.ADR(s.damageDealt, rounds)
	s.killsPerRound = stat.AVG(s.kills, rounds)
	s.assistsPerRound = stat.AVG(s.assists, rounds)
	s.deathsPerRound = stat.AVG(s.deaths, rounds)
	s.blindedPlayersPerRound = stat.AVG(s.blindedPlayers, rounds)
}

// add adds v to a particular field depending on a received metric.
func (s *playerStat) add(m metric, v int) { //nolint:gocyclo // no other options to implement this
	//nolint:exhaustive // each metric corresponds to specific playerStat field, no need to check all
	switch m {
	case metricKill:
		s.kills += int32(v)
	case metricHSKill:
		s.hsKills += int32(v)
	case metricBlindKill:
		s.blindKills += int32(v)
	case metricWallbangKill:
		s.wallbangKills += int32(v)
	case metricNoScopeKill:
		s.noScopeKills += int32(v)
	case metricThroughSmokeKill:
		s.throughSmokeKills += int32(v)
	case metricDeath:
		s.deaths += int32(v)
	case metricAssist:
		s.assists += int32(v)
	case metricFlashbangAssist:
		s.flashbangAssists += int32(v)
	case metricRoundMVP:
		s.mvpCount += int32(v)
	case metricDamageTaken:
		s.damageTaken += int32(v)
	case metricDamageDealt:
		s.damageDealt += int32(v)
	case metricBlind:
		s.blindedPlayers += int32(v)
	case metricBlinded:
		s.blindedTimes += int32(v)
	case metricBombPlanted:
		s.bombsPlanted += int32(v)
	case metricBombDefused:
		s.bombsDefused += int32(v)
	case metricGrenadeDamageDealt:
		s.grenadeDamageDealt += int32(v)
	}
}

type weaponStat struct {
	steamID           uint64
	weaponID          int32
	kills             int32
	hsKills           int32
	blindKills        int32
	wallbangKills     int32
	noScopeKills      int32
	throughSmokeKills int32
	deaths            int32
	assists           int32
	damageTaken       int32
	damageDealt       int32
	shots             int32
	totalHits         int32
	headHits          int32
	neckHits          int32
	chestHits         int32
	stomachHits       int32
	armHits           int32
	legHits           int32

	accuracy *domain.WeaponAccuracyStats
}

func (ws *weaponStat) calculateTotalHits() {
	ws.totalHits = ws.headHits +
		ws.neckHits +
		ws.chestHits +
		ws.stomachHits +
		ws.armHits +
		ws.legHits
}

// add adds v to a particular field depending on a received metric.
func (ws *weaponStat) add(m metric, v int) { //nolint:gocyclo // no other options to implement this
	//nolint:exhaustive // each metric corresponds to specific weaponStat field, no need to check all
	switch m {
	case metricKill:
		ws.kills += int32(v)
	case metricHSKill:
		ws.hsKills += int32(v)
	case metricBlindKill:
		ws.blindKills += int32(v)
	case metricWallbangKill:
		ws.wallbangKills += int32(v)
	case metricNoScopeKill:
		ws.noScopeKills += int32(v)
	case metricThroughSmokeKill:
		ws.throughSmokeKills += int32(v)
	case metricDeath:
		ws.deaths += int32(v)
	case metricAssist:
		ws.assists += int32(v)
	case metricDamageTaken:
		ws.damageTaken += int32(v)
	case metricDamageDealt:
		ws.damageDealt += int32(v)
	case metricShot:
		ws.shots += int32(v)
	case metricHitHead:
		ws.headHits += int32(v)
	case metricHitNeck:
		ws.neckHits += int32(v)
	case metricHitChest:
		ws.chestHits += int32(v)
	case metricHitStomach:
		ws.stomachHits += int32(v)
	case metricHitLeftArm, metricHitRightArm:
		ws.armHits += int32(v)
	case metricHitLeftLeg, metricHitRightLeg:
		ws.legHits += int32(v)
	}
}
