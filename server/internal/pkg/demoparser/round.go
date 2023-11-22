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

// killCount appends kill to last round kill feed and removes player from survivors.
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
		delete(currRound.TeamA.Survivors, kill.Victim.SteamID64)
	case currRound.TeamB.Side:
		delete(currRound.TeamB.Survivors, kill.Victim.SteamID64)
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

type roundTeam struct {
	Survivors map[uint64]struct{}
	Cash      int // cash at start of round, must be set on round start
	CashSpend int // during round, must be set on round end
	EqValue   int // equipment value on round start, must be set on round end
	Score     int
	Side      common.Team
}

// newRoundTeam must be created at round start.
func newRoundTeam(members []*common.Player, side common.Team) *roundTeam {
	cash := 0
	survivors := make(map[uint64]struct{}, len(members))

	for _, m := range members {
		if !playerConnected(m) {
			slog.Error("player not added to round team", "player", m)
			continue
		}

		cash += m.Money()
		survivors[m.SteamID64] = struct{}{}
	}

	return &roundTeam{
		Cash:      cash,
		Side:      side,
		Survivors: survivors,
	}
}

// onRoundEnd calculates money spent this round and freezetime end equipment value.
func (rt *roundTeam) onRoundEnd(ts *common.TeamState) {
	if ts.Team() == rt.Side {
		rt.CashSpend = ts.MoneySpentThisRound()
		rt.EqValue = ts.FreezeTimeEndEquipmentValue()
		rt.Score = ts.Score()
	} else {
		slog.Error("got invalid team state on round end",
			"team", rt,
			"state", ts)
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
