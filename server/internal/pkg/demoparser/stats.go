package demoparser

import (
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type dmgStats struct {
	Dealt int
	Taken int
}

type dmgGrenadeStats struct {
	dmgStats
	DealtWithGrenade int
}

type bombStats struct {
	Planted int
	Defused int
}

type hitStats struct {
	total   int
	head    int
	neck    int
	chest   int
	stomach int
	arms    int
	legs    int
}

type killStats struct {
	total   int
	hs      int
	blind   int
	wb      int
	smoke   int
	noscope int
}

// type accuracyStats struct {
// 	total   float64
// 	head    float64
// 	neck    float64
// 	chest   float64
// 	stomach float64
// 	arms    float64
// 	legs    float64
// }

// func newAccuracyStats(shots int, hits *hitStats) *accuracyStats {
// 	totalHits := int32(hits.total)
// 	return &accuracyStats{
// 		total:   stat.Accuracy(totalHits, int32(shots)),
// 		head:    stat.Accuracy(int32(hits.head), totalHits),
// 		neck:    stat.Accuracy(int32(hits.neck), totalHits),
// 		chest:   stat.Accuracy(int32(hits.chest), totalHits),
// 		stomach: stat.Accuracy(int32(hits.stomach), totalHits),
// 		arms:    stat.Accuracy(int32(hits.arms), totalHits),
// 		legs:    stat.Accuracy(int32(hits.legs), totalHits),
// 	}
// }

type playerStatsMap map[uint64]*playerStats

func (psm playerStatsMap) add(steamID uint64, ev event, val int) {
	if _, ok := psm[steamID]; !ok {
		psm[steamID] = &playerStats{}
	}

	psm[steamID].add(ev, val)
}

func (psm playerStatsMap) incr(steamID uint64, ev event) {
	psm.add(steamID, ev, 1)
}

type playerStats struct {
	kills          *killStats
	damage         *dmgGrenadeStats
	deaths         int
	assists        int
	fbAssists      int
	mvps           int
	blindedPlayers int
	blindedTimes   int
}

func (ps *playerStats) add(e event, v int) {
	switch e {
	case eventKill:
		ps.kills.total += v
	case eventHSKill:
		ps.kills.hs += v
	case eventBlindKill:
		ps.kills.blind += v
	case eventWBKill:
		ps.kills.wb += v
	case eventSmokeKill:
		ps.kills.smoke += v
	case eventNoScopeKill:
		ps.kills.noscope += v
	case eventDeath:
		ps.deaths += v
	case eventAssist:
		ps.assists += v
	case eventFBAssist:
		ps.fbAssists += v
	case eventRoundMVP:
		ps.mvps += v
	case eventDmgDealt:
		ps.damage.Dealt += v
	case eventDmgTaken:
		ps.damage.Taken += v
	case eventDmgGrenadeDealt:
		ps.damage.DealtWithGrenade += v
	case eventBlindedPlayer:
		ps.blindedPlayers += v
	case eventBecameBlind:
		ps.blindedTimes += v
	}
}

// equipValid checks whether equipment is valid for dealing damage to other players or yourself.
func equipValid(e common.EquipmentType) bool {
	if e.Class() == common.EqClassUnknown ||
		e == common.EqUnknown ||
		e == common.EqKevlar ||
		e == common.EqHelmet ||
		e == common.EqDefuseKit {
		return false
	}

	return true
}

type weaponStatsMap map[uint64]map[int]*weaponStats

func (ws weaponStatsMap) add(steamID uint64, ev event, et common.EquipmentType, val int) {
	if !equipValid(et) {
		return
	}

	_, ok := ws[steamID]
	if !ok {
		ws[steamID] = make(map[int]*weaponStats, 100)
	}

	weaponID := int(et)

	if _, ok = ws[steamID][weaponID]; !ok {
		ws[steamID][weaponID] = &weaponStats{}
	}

	ws[steamID][weaponID].add(ev, val)
}

func (ws weaponStatsMap) incr(steamID uint64, ev event, et common.EquipmentType) {
	ws.add(steamID, ev, et, 1)
}

type weaponStats struct {
	hits    *hitStats
	kills   *killStats
	damage  dmgStats
	deaths  int
	assists int
	shots   int
}

func (ws *weaponStats) add(e event, v int) {
	switch e {
	case eventKill:
		ws.kills.total += v
	case eventHSKill:
		ws.kills.hs += v
	case eventBlindKill:
		ws.kills.blind += v
	case eventWBKill:
		ws.kills.wb += v
	case eventSmokeKill:
		ws.kills.smoke += v
	case eventNoScopeKill:
		ws.kills.noscope += v
	case eventDeath:
		ws.deaths += v
	case eventAssist:
		ws.assists += v
	case eventDmgDealt:
		ws.damage.Dealt += v
	case eventDmgTaken:
		ws.damage.Taken += v
	case eventShot:
		ws.shots += v
	case eventHitHead:
		ws.hits.head += v
	case eventHitNeck:
		ws.hits.neck += v
	case eventHitStomach:
		ws.hits.stomach += v
	case eventHitArm:
		ws.hits.arms += v
	case eventHitLeg:
		ws.hits.legs += v
	}
}
