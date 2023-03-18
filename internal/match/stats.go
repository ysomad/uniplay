package match

import (
	"sync"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
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

// normalize returns restructured player and weapon stats.
func (s *stats) normalize() ([]*playerStat, []*weaponStat) { //nolint:unused // for benchmarks
	var (
		wg          sync.WaitGroup
		playerStats []*playerStat
		weaponStats []*weaponStat
	)

	wg.Add(2)

	go func() {
		for _, ps := range s.playerStats {
			playerStats = append(playerStats, ps)
		}

		wg.Done()
	}()

	go func() {
		for _, weapons := range s.weaponStats {
			for _, ws := range weapons {
				weaponStats = append(weaponStats, ws)
			}
		}

		wg.Done()
	}()

	wg.Wait()

	return playerStats, weaponStats
}

func (s *stats) normalizeSync() ([]*playerStat, []*weaponStat) {
	playerStats := make([]*playerStat, 0, len(s.playerStats))

	for _, ps := range s.playerStats {
		playerStats = append(playerStats, ps)
	}

	var weaponStats []*weaponStat

	for _, weapons := range s.weaponStats {
		for _, ws := range weapons {
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
	kills              int
	hsKills            int
	blindKills         int
	wallbangKills      int
	noScopeKills       int
	throughSmokeKills  int
	deaths             int
	assists            int
	flashbangAssists   int
	mvpCount           int
	damageTaken        int
	damageDealt        int
	grenadeDamageDealt int
	blindedPlayers     int
	blindedTimes       int
	bombsPlanted       int
	bombsDefused       int
}

// add adds v to a particular field depending on a received metric.
func (s *playerStat) add(m metric, v int) { //nolint:gocyclo // no other options to implement this
	//nolint:exhaustive // each metric corresponds to specific playerStat field, no need to check all
	switch m {
	case metricKill:
		s.kills += v
	case metricHSKill:
		s.hsKills += v
	case metricBlindKill:
		s.blindKills += v
	case metricWallbangKill:
		s.wallbangKills += v
	case metricNoScopeKill:
		s.noScopeKills += v
	case metricThroughSmokeKill:
		s.throughSmokeKills += v
	case metricDeath:
		s.deaths += v
	case metricAssist:
		s.assists += v
	case metricFlashbangAssist:
		s.flashbangAssists += v
	case metricRoundMVP:
		s.mvpCount += v
	case metricDamageTaken:
		s.damageTaken += v
	case metricDamageDealt:
		s.damageDealt += v
	case metricBlind:
		s.blindedPlayers += v
	case metricBlinded:
		s.blindedTimes += v
	case metricBombPlanted:
		s.bombsPlanted += v
	case metricBombDefused:
		s.bombsDefused += v
	case metricGrenadeDamageDealt:
		s.grenadeDamageDealt += v
	}
}

type weaponStat struct {
	steamID           uint64
	weaponID          int32
	kills             int
	hsKills           int
	blindKills        int
	wallbangKills     int
	noScopeKills      int
	throughSmokeKills int
	deaths            int
	assists           int
	damageTaken       int
	damageDealt       int
	shots             int
	headHits          int
	neckHits          int
	chestHits         int
	stomachHits       int
	leftArmHits       int
	rightArmHits      int
	leftLegHits       int
	rightLegHits      int
}

// add adds v to a particular field depending on a received metric.
func (ws *weaponStat) add(m metric, v int) { //nolint:gocyclo // no other options to implement this
	//nolint:exhaustive // each metric corresponds to specific weaponStat field, no need to check all
	switch m {
	case metricKill:
		ws.kills += v
	case metricHSKill:
		ws.hsKills += v
	case metricBlindKill:
		ws.blindKills += v
	case metricWallbangKill:
		ws.wallbangKills += v
	case metricNoScopeKill:
		ws.noScopeKills += v
	case metricThroughSmokeKill:
		ws.throughSmokeKills += v
	case metricDeath:
		ws.deaths += v
	case metricAssist:
		ws.assists += v
	case metricDamageTaken:
		ws.damageTaken += v
	case metricDamageDealt:
		ws.damageDealt += v
	case metricShot:
		ws.shots += v
	case metricHitHead:
		ws.headHits += v
	case metricHitNeck:
		ws.neckHits += v
	case metricHitChest:
		ws.chestHits += v
	case metricHitStomach:
		ws.stomachHits += v
	case metricHitLeftArm:
		ws.leftArmHits += v
	case metricHitRightArm:
		ws.rightArmHits += v
	case metricHitLeftLeg:
		ws.leftLegHits += v
	case metricHitRightLeg:
		ws.rightLegHits += v
	}
}
