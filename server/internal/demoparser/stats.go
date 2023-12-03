package demoparser

import (
	"log/slog"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"

	"github.com/ysomad/uniplay/server/internal/pkg/stat"
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
	Total   int
	Head    int
	Neck    int
	Chest   int
	Stomach int
	Arms    int
	Legs    int
}

// calculateTotal calculate total amount of hits.
func (hs *hitStats) calculateTotal() {
	hs.Total = hs.Head + hs.Neck + hs.Chest + hs.Stomach + hs.Arms + hs.Legs
}

type killStats struct {
	Total    int
	HS       int
	Blind    int
	Wallbang int
	Smoke    int
	NoScope  int
}

type accuracyStats struct {
	Total   float64
	Head    float64
	Neck    float64
	Chest   float64
	Stomach float64
	Arms    float64
	Legs    float64
}

func newAccuracyStats(shots int, hits *hitStats) *accuracyStats {
	totalHits := int32(hits.Total)
	return &accuracyStats{
		Total:   stat.Accuracy(totalHits, int32(shots)),
		Head:    stat.Accuracy(int32(hits.Head), totalHits),
		Neck:    stat.Accuracy(int32(hits.Neck), totalHits),
		Chest:   stat.Accuracy(int32(hits.Chest), totalHits),
		Stomach: stat.Accuracy(int32(hits.Stomach), totalHits),
		Arms:    stat.Accuracy(int32(hits.Arms), totalHits),
		Legs:    stat.Accuracy(int32(hits.Legs), totalHits),
	}
}

type playerStatsMap map[uint64]*playerStats

func (psm playerStatsMap) add(steamID uint64, ev event, val int) {
	if ev == eventUnknown || val <= 0 {
		slog.Error("player stat not added",
			"steam_id", steamID,
			"event", ev,
			"value", val)
		return
	}

	if _, ok := psm[steamID]; !ok {
		psm[steamID] = newPlayerStats()
	}

	psm[steamID].add(ev, val)
}

func (psm playerStatsMap) incr(steamID uint64, ev event) {
	psm.add(steamID, ev, 1)
}

type playerStats struct {
	Kills            *killStats
	Damage           *dmgGrenadeStats
	Bomb             bombStats
	Deaths           int
	Assists          int
	FlashbangAssists int
	MVPs             int
	BlindedPlayers   int
	BlindedTimes     int
}

func newPlayerStats() *playerStats {
	return &playerStats{
		Kills:  &killStats{},
		Damage: &dmgGrenadeStats{},
	}
}

func (ps *playerStats) add(e event, v int) {
	switch e {
	case eventKill:
		ps.Kills.Total += v
	case eventHSKill:
		ps.Kills.HS += v
	case eventBlindKill:
		ps.Kills.Blind += v
	case eventWBKill:
		ps.Kills.Wallbang += v
	case eventSmokeKill:
		ps.Kills.Smoke += v
	case eventNoScopeKill:
		ps.Kills.NoScope += v
	case eventDeath:
		ps.Deaths += v
	case eventAssist:
		ps.Assists += v
	case eventFBAssist:
		ps.FlashbangAssists += v
	case eventRoundMVP:
		ps.MVPs += v
	case eventDmgDealt:
		ps.Damage.Dealt += v
	case eventDmgTaken:
		ps.Damage.Taken += v
	case eventDmgGrenadeDealt:
		ps.Damage.DealtWithGrenade += v
	case eventBlindedPlayer:
		ps.BlindedPlayers += v
	case eventBecameBlind:
		ps.BlindedTimes += v
	case eventBombDefused:
		ps.Bomb.Defused += v
	case eventBombPlanted:
		ps.Bomb.Planted += v
	default:
		slog.Error("player stats event not added",
			"event", e,
			"value", v)
	}
}

// isWeapon checks whether equipment is valid for dealing damage to other players or yourself.
func isWeapon(e common.EquipmentType) bool {
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
	if ev == eventUnknown || val <= 0 {
		slog.Error("weapon stat not added",
			"steam_id", steamID,
			"event", ev,
			"value", val,
			"weapon", et.String())
		return
	}

	if !isWeapon(et) {
		return
	}

	_, ok := ws[steamID]
	if !ok {
		ws[steamID] = make(map[int]*weaponStats, 100)
	}

	weaponID := int(et)

	if _, ok = ws[steamID][weaponID]; !ok {
		ws[steamID][weaponID] = newWeaponStats()
	}

	ws[steamID][weaponID].add(ev, val)
}

func (ws weaponStatsMap) incr(steamID uint64, ev event, et common.EquipmentType) {
	ws.add(steamID, ev, et, 1)
}

// calculateUnobtainableStats calculates unobtainable stats from demo for every player weapon stats.
// (because its not obtainable from demo and must be calculated manually).
func (wsm weaponStatsMap) calculateUnobtainableStats() {
	for _, wss := range wsm {
		for _, ws := range wss {
			ws.Hits.calculateTotal()
			ws.Accuracy = newAccuracyStats(ws.Shots, ws.Hits)
		}
	}
}

type weaponStats struct {
	Hits     *hitStats
	Kills    *killStats
	Accuracy *accuracyStats
	Damage   dmgStats
	Assists  int
	Shots    int
}

func newWeaponStats() *weaponStats {
	return &weaponStats{
		Hits:  &hitStats{},
		Kills: &killStats{},
	}
}

func (ws *weaponStats) add(e event, v int) {
	switch e {
	case eventKill:
		ws.Kills.Total += v
	case eventHSKill:
		ws.Kills.HS += v
	case eventBlindKill:
		ws.Kills.Blind += v
	case eventWBKill:
		ws.Kills.Wallbang += v
	case eventSmokeKill:
		ws.Kills.Smoke += v
	case eventNoScopeKill:
		ws.Kills.NoScope += v
	case eventAssist:
		ws.Assists += v
	case eventDmgDealt:
		ws.Damage.Dealt += v
	case eventDmgTaken:
		ws.Damage.Taken += v
	case eventShot:
		ws.Shots += v
	case eventHitHead:
		ws.Hits.Head += v
	case eventHitNeck:
		ws.Hits.Neck += v
	case eventHitStomach:
		ws.Hits.Stomach += v
	case eventHitChest:
		ws.Hits.Chest += v
	case eventHitArm:
		ws.Hits.Arms += v
	case eventHitLeg:
		ws.Hits.Legs += v
	default:
		slog.Error("weapon stats event not added",
			"event", e,
			"value", v)
	}
}
