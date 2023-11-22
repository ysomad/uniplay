package demoparser

import (
	"errors"
	"log/slog"
	"time"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

var (
	errNoRounds                  = errors.New("no rounds found in game state")
	errInvalidVictimSide         = errors.New("invalid victim side")
	errUnconnectedKillerOrVictim = errors.New("unconnected killer or victim")
)

type roundHistory struct {
	Rounds []*round
}

func newRoundHistory() roundHistory {
	return roundHistory{Rounds: make([]*round, 0, 50)}
}

// start appends new round into rounds.
// currTime is time elapsed since demo start.
func (rh *roundHistory) start(ts *common.TeamState, currTime time.Duration) {
	if len(ts.Members()) <= 0 {
		slog.Info("skipping round start for empty team")
		return
	}

	if ts.Members()[0].Money() <= 0 {
		slog.Info("skipping round start for no cash member")
		return
	}

	rh.Rounds = append(rh.Rounds, newRound(ts, currTime))
}

// endCurrent ends latest rounds in rounds history.
func (rh *roundHistory) endCurrent(e events.RoundEnd) error {
	if len(rh.Rounds) < 1 {
		return errNoRounds
	}

	rh.Rounds[len(rh.Rounds)-1].end(e.WinnerState, e.LoserState, e.Reason)

	return nil
}

// killCount appends kill to last round kill feed.
func (rh roundHistory) killCount(kill events.Kill, killTime time.Duration) error {
	if !playerConnected(kill.Killer) || !playerConnected(kill.Victim) {
		return errUnconnectedKillerOrVictim
	}

	if len(rh.Rounds) < 1 {
		return errNoRounds
	}

	// add kill to round kill feed
	currRound := rh.Rounds[len(rh.Rounds)-1]
	currRound.KillFeed = append(currRound.KillFeed, newRoundKill(kill, currRound.Time, killTime))

	// remove victim from team survivors list
	switch kill.Victim.Team {
	case currRound.TeamA.Side:
		// delete(currRound.TeamA.Survivors, kill.Victim.SteamID64)
		currRound.TeamA.killCount(kill.Victim.SteamID64)
	case currRound.TeamB.Side:
		// delete(currRound.TeamB.Survivors, kill.Victim.SteamID64)
		currRound.TeamB.killCount(kill.Victim.SteamID64)
	default:
		return errInvalidVictimSide
	}

	return nil
}

type round struct {
	TeamA    *roundTeam
	TeamB    *roundTeam
	KillFeed []*roundKill
	Time     time.Duration // time elapsed since demo start
	Reason   events.RoundEndReason
}

func newRound(ts *common.TeamState, currTime time.Duration) *round {
	return &round{
		TeamA:    newRoundTeam(ts.Members(), ts.Team()),
		TeamB:    newRoundTeam(ts.Opponent.Members(), ts.Opponent.Team()),
		KillFeed: make([]*roundKill, 0, 20),
		Reason:   events.RoundEndReasonStillInProgress,
		Time:     currTime,
	}
}

func (r *round) end(winner, loser *common.TeamState, reason events.RoundEndReason) {
	switch winner.Team() {
	case r.TeamA.Side:
		r.TeamA.onRoundEnd(winner)
		r.TeamB.onRoundEnd(loser)
	case r.TeamB.Side:
		r.TeamA.onRoundEnd(loser)
		r.TeamB.onRoundEnd(winner)
	default:
		slog.Error("round not ended",
			"winner", winner,
			"loser", loser,
			"team_a", r.TeamA,
			"team_b", r.TeamB)
		return
	}

	r.Reason = reason
}

type roundPlayer struct {
	Inventory []common.EquipmentType
	CashSpend int
	Survived  bool
}

func newRoundPlayer(inventory map[int]*common.Equipment) roundPlayer {
	p := roundPlayer{
		Inventory: make([]common.EquipmentType, 0, len(inventory)),
		CashSpend: 0,
		Survived:  true,
	}

	for _, eq := range inventory {
		if eq == nil {
			slog.Error("nil equipment when creating round player")
			continue
		}

		p.Inventory = append(p.Inventory, eq.Type)
	}

	return p
}

type roundTeam struct {
	Players   map[uint64]roundPlayer
	Cash      int // cash at start of round, must be set on round start
	CashSpend int // during round, must be set on round end
	EqValue   int // equipment value on round start, must be set on round end
	Score     int
	Side      common.Team
}

// newRoundTeam must be created at round start.
func newRoundTeam(members []*common.Player, side common.Team) *roundTeam {
	cash := 0
	players := make(map[uint64]roundPlayer, len(members))

	for _, m := range members {
		if !playerConnected(m) {
			slog.Error("player not added to round team", "player", m)
			continue
		}

		if m.Inventory == nil {
			slog.Error("player has empty inventory when creating round team", "player", m)
			continue
		}

		players[m.SteamID64] = newRoundPlayer(m.Inventory)
		cash += m.Money()
	}

	return &roundTeam{
		Cash:    cash,
		Side:    side,
		Players: players,
	}
}

// killCount sets survived to false for specified player.
func (rt *roundTeam) killCount(steamID uint64) {
	player, ok := rt.Players[steamID]
	if !ok {
		slog.Error("kill not counted, player not found in round team",
			"steam_id", steamID,
			"round_team", rt)
		return
	}

	player.Survived = false
	rt.Players[steamID] = player
}

func (rt *roundTeam) setPlayerCashSpend(steamID uint64, spend int) error {
	pl, ok := rt.Players[steamID]
	if !ok {
		return errors.New("player not found in round team")
	}
	pl.CashSpend = spend
	rt.Players[steamID] = pl
	return nil
}

// onRoundEnd calculates money spent this round and freezetime end equipment value.
func (rt *roundTeam) onRoundEnd(ts *common.TeamState) {
	if ts.Team() == rt.Side {
		rt.CashSpend = ts.MoneySpentThisRound()
		rt.EqValue = ts.FreezeTimeEndEquipmentValue()
		rt.Score = ts.Score()

		for _, m := range ts.Members() {
			err := rt.setPlayerCashSpend(m.SteamID64, m.MoneySpentThisRound())
			if err != nil {
				slog.Error("player from team state not found in round team", "round_team", rt)
				continue
			}
		}
	} else {
		slog.Error("got invalid team state on round end",
			"rount_team", rt,
			"team_state", ts)
	}
}

type roundKill struct {
	Killer        uint64
	Victim        uint64
	Assister      uint64
	SinceStart    uint16
	Headshot      bool
	Wallbang      bool
	KillerBlind   bool
	ThroughSmoke  bool
	NoScope       bool
	AssistedFlash bool
	KillerSide    common.Team
	AssisterSide  common.Team
	Weapon        common.EquipmentType
}

func newRoundKill(kill events.Kill, roundTime, killTime time.Duration) *roundKill {
	k := &roundKill{
		Killer:       kill.Killer.SteamID64,
		KillerSide:   kill.Killer.Team,
		Victim:       kill.Victim.SteamID64,
		Headshot:     kill.IsHeadshot,
		Wallbang:     kill.IsWallBang(),
		KillerBlind:  kill.AttackerBlind,
		ThroughSmoke: kill.ThroughSmoke,
		NoScope:      kill.NoScope,
		SinceStart:   uint16((killTime - roundTime).Seconds()),
		Weapon:       kill.Weapon.Type,
	}

	if playerConnected(kill.Assister) {
		k.Assister = kill.Assister.SteamID64
		k.AssisterSide = kill.Assister.Team
		k.AssistedFlash = kill.AssistedFlash
	}

	return k
}
