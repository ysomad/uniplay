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

// cleanup returns only ended rounds.
// there is a situations when game is restarting and round has no freeze time and round end events.
func (rh roundHistory) cleanup() []*round {
	rr := make([]*round, 0, len(rh.Rounds))

	for _, r := range rh.Rounds {
		// filter not ended rounds
		if r.Reason == events.RoundEndReasonStillInProgress {
			continue
		}

		rr = append(rr, r)
	}

	return rr
}

// start appends new round into rounds.
// currTime is time elapsed since demo start.
func (rh *roundHistory) start(t, ct []*common.Player, currTime time.Duration) {
	if len(t) <= 0 || len(ct) <= 0 {
		slog.Info("skipping round start for empty team", "t_players", t, "ct_players", ct)
		return
	}

	rh.Rounds = append(rh.Rounds, newRound(t, ct, currTime))
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
	if !isPlayerConnected(kill.Killer) || !isPlayerConnected(kill.Victim) {
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
	case currRound.T.Side:
		currRound.T.killCount(kill.Victim.SteamID64)
	case currRound.CT.Side:
		currRound.CT.killCount(kill.Victim.SteamID64)
	default:
		return errInvalidVictimSide
	}

	return nil
}

// currNum returns current round number.
func (rh roundHistory) currNum() int {
	return len(rh.Rounds)
}

func (rh roundHistory) current() (*round, error) {
	if len(rh.Rounds) < 1 {
		return nil, errors.New("no rounds found")
	}

	return rh.Rounds[len(rh.Rounds)-1], nil
}

type round struct {
	T        *roundTeam
	CT       *roundTeam
	KillFeed []*roundKill
	Time     time.Duration // time elapsed since demo start
	Reason   events.RoundEndReason
}

func newRound(t, ct []*common.Player, currTime time.Duration) *round {
	return &round{
		T:        newRoundTeam(t),
		CT:       newRoundTeam(ct),
		KillFeed: make([]*roundKill, 0, 20),
		Reason:   events.RoundEndReasonStillInProgress,
		Time:     currTime,
	}
}

// setWeapons saves players weapons to round team.
func (r *round) setWeapons(t, ct []*common.Player) {
	for _, m := range t {
		r.T.setPlayerWeapons(m)
	}
	for _, m := range ct {
		r.CT.setPlayerWeapons(m)
	}
}

func (r *round) end(winner, loser *common.TeamState, reason events.RoundEndReason) {
	switch winner.Team() {
	case r.T.Side:
		r.T.onRoundEnd(winner)
		r.CT.onRoundEnd(loser)
	case r.CT.Side:
		r.T.onRoundEnd(loser)
		r.CT.onRoundEnd(winner)
	default:
		slog.Error("round not ended",
			"winner", winner,
			"loser", loser,
			"t", r.T,
			"ct", r.CT)
		return
	}

	r.Reason = reason
}

type roundTeamPlayer struct {
	Weapons   []string
	CashSpend int
	Armor     bool
	Helmet    bool
	DefuseKit bool
	Bomb      bool
	Survived  bool
	Side      common.Team
}

func newRoundTeamPlayer(side common.Team) roundTeamPlayer {
	return roundTeamPlayer{
		CashSpend: 0,
		Survived:  true,
		Side:      side,
	}
}

type roundTeam struct {
	Players   map[uint64]roundTeamPlayer
	Cash      int // cash at start of round, must be set on round start
	CashSpend int // during round, must be set on round end
	EqValue   int // equipment value on round start, must be set on round end
	Score     int
	Side      common.Team
}

// newRoundTeam must be created at round start.
func newRoundTeam(pp []*common.Player) *roundTeam {
	cash := 0
	players := make(map[uint64]roundTeamPlayer, len(pp))

	for _, p := range pp {
		if !isPlayerConnected(p) {
			slog.Error("player not added to round team", "player", p)
			continue
		}

		players[p.SteamID64] = newRoundTeamPlayer(p.Team)
		cash += p.Money()
	}

	return &roundTeam{
		Cash:    cash,
		Side:    pp[0].Team,
		Players: players,
	}
}

// setPlayerWeapons saves player weapons into round team.
func (rt *roundTeam) setPlayerWeapons(p *common.Player) {
	pl := rt.Players[p.SteamID64]

	pl.Weapons = make([]string, 0, len(p.Inventory))
	pl.Helmet = p.HasHelmet()
	pl.Armor = p.Armor() > 0
	pl.DefuseKit = p.HasDefuseKit()

	for _, eq := range p.Inventory {
		if eq.Type == common.EqKnife {
			continue
		}

		if eq.Type == common.EqBomb {
			pl.Bomb = true
			continue
		}

		pl.Weapons = append(pl.Weapons, eq.String())
	}

	rt.Players[p.SteamID64] = pl
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

	if isPlayerConnected(kill.Assister) {
		k.Assister = kill.Assister.SteamID64
		k.AssisterSide = kill.Assister.Team
		k.AssistedFlash = kill.AssistedFlash
	}

	return k
}
