package replayparser

import (
	"fmt"
	"io"

	"github.com/ssssargsian/uniplay/internal/domain"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
)

// parser is a wrapper around demoinfocs.Parser.
// ONE parser must be used for ONE replay.
type parser struct {
	demoinfocs.Parser

	metrics       *playerMetrics
	weaponMetrics *weaponMetrics
	match         *match
}

func New(r io.Reader) *parser {
	return &parser{
		demoinfocs.NewParser(r),
		newPlayerMetrics(),
		newWeaponMetrics(),
		&match{},
	}
}

func (p *parser) Parse() (parseResult, error) {
	p.RegisterEventHandler(func(_ events.RoundFreezetimeEnd) {
		p.detectKnifeRound()
	})

	// TODO: https://github.com/markus-wa/demoinfocs-golang/discussions/366
	// p.RegisterEventHandler(func(e events.MatchStartedChanged) {
	// 	p.setTeams(p.GameState())
	// })

	p.RegisterEventHandler(func(_ events.MatchStart) {
		p.setTeams(p.GameState())
	})

	p.RegisterEventHandler(func(_ events.TeamSideSwitch) {
		fmt.Println("TEAM SWITCH")
		p.match.swapTeamSides()
	})

	p.RegisterEventHandler(func(e events.ScoreUpdated) {
		p.match.updateTeamsScore(e)
	})

	p.RegisterEventHandler(func(e events.Kill) {
		p.handleKills(e)
	})

	p.RegisterEventHandler(func(e events.PlayerHurt) {
		p.handlePlayerHurt(e)
	})

	p.RegisterEventHandler(func(e events.BombDefused) {
		if !p.collectStats(p.GameState()) || !p.playerConnected(e.Player) {
			return
		}
		p.metrics.incr(e.Player.SteamID64, domain.MetricBombDefused)
	})

	p.RegisterEventHandler(func(e events.BombPlanted) {
		if !p.collectStats(p.GameState()) || !p.playerConnected(e.Player) {
			return
		}
		p.metrics.incr(e.Player.SteamID64, domain.MetricBombPlanted)
	})

	p.RegisterEventHandler(func(e events.PlayerFlashed) {
		if !p.collectStats(p.GameState()) {
			return
		}

		if p.playerConnected(e.Player) {
			p.metrics.incr(e.Player.SteamID64, domain.MetricBlinded)
		}

		if p.playerConnected(e.Attacker) {
			p.metrics.incr(e.Attacker.SteamID64, domain.MetricBlind)
		}
	})

	p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if !p.collectStats(p.GameState()) || !p.playerConnected(e.Player) {
			return
		}
	})

	p.RegisterEventHandler(func(_ events.AnnouncementWinPanelMatch) {
		p.match.mapName = p.Header().MapName
		p.match.duration = p.Header().PlaybackTime
	})

	if err := p.ParseToEnd(); err != nil {
		return parseResult{}, err
	}

	return parseResult{
		metrics:       p.metrics,
		weaponMetrics: p.weaponMetrics,
		match:         p.match,
	}, nil
}

// collectStats detects if stats can be collected to prevent collection of stats on knife or warmup rounds.
// return false if current round is knife round or match is not started.
func (p *parser) collectStats(gs demoinfocs.GameState) bool {
	if p.match.isKnifeRound() || !gs.IsMatchStarted() {
		return false
	}

	return true
}

// detectKnifeRound sets isKnifeRound to true if any player has no secondary weapon and first slot is a knife.
func (p *parser) detectKnifeRound() {
	p.match.setIsKnifeRound(false)

	for _, player := range p.GameState().TeamCounterTerrorists().Members() {
		weapons := player.Weapons()
		if len(weapons) == 1 && weapons[0].Type == common.EqKnife {
			p.match.setIsKnifeRound(true)
			break
		}
	}
}

// playerConnected checks is player connected and steamID is not equal to 0.
func (p *parser) playerConnected(pl *common.Player) bool {
	return pl != nil && pl.SteamID64 != 0
}

// setTeams sets teams clan names, flags, sides (ct/t) and their members to p.match.
func (p *parser) setTeams(gs demoinfocs.GameState) {
	t := gs.TeamTerrorists()
	p.match.setTeam1(newMatchTeam(t.ClanName(), t.Flag(), t.Team(), t.Members()))

	ct := gs.TeamCounterTerrorists()
	p.match.setTeam2(newMatchTeam(ct.ClanName(), ct.Flag(), ct.Team(), ct.Members()))

}

// handleKills collects metrics and weapon metrics on kill event.
func (p *parser) handleKills(e events.Kill) {
	if !p.collectStats(p.GameState()) {
		return
	}

	if p.playerConnected(e.Victim) {
		// death amount FROM weapon
		p.metrics.incr(e.Victim.SteamID64, domain.MetricDeath)

		if e.Weapon != nil {
			p.weaponMetrics.incr(e.Victim.SteamID64, weaponMetric{
				weaponName:  e.Weapon.String(),
				weaponClass: e.Weapon.Class(),
			}, domain.MetricDeath)
		}
	}

	if p.playerConnected(e.Killer) {
		// kill amount
		p.metrics.incr(e.Killer.SteamID64, domain.MetricKill)

		if e.Weapon != nil {
			p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
				weaponName:  e.Weapon.String(),
				weaponClass: e.Weapon.Class(),
			}, domain.MetricKill)
		}

		switch {
		// headshot kill amount
		case e.IsHeadshot:
			p.metrics.incr(e.Killer.SteamID64, domain.MetricHSKill)

			if e.Weapon != nil {
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  e.Weapon.String(),
					weaponClass: e.Weapon.Class(),
				}, domain.MetricHSKill)
			}

		// blind kill amount
		case e.AttackerBlind:
			p.metrics.incr(e.Killer.SteamID64, domain.MetricBlindKill)

			if e.Weapon != nil {
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  e.Weapon.String(),
					weaponClass: e.Weapon.Class(),
				}, domain.MetricBlindKill)
			}

		// wallbang kill amount
		case e.IsWallBang():
			p.metrics.incr(e.Killer.SteamID64, domain.MetricWallbangKill)

			if e.Weapon != nil {
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  e.Weapon.String(),
					weaponClass: e.Weapon.Class(),
				}, domain.MetricWallbangKill)
			}

		// noscope kill amount
		case e.NoScope:
			p.metrics.incr(e.Killer.SteamID64, domain.MetricNoScopeKill)

			if e.Weapon != nil {
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  e.Weapon.String(),
					weaponClass: e.Weapon.Class(),
				}, domain.MetricNoScopeKill)
			}

		// through smoke kill amount
		case e.ThroughSmoke:
			p.metrics.incr(e.Killer.SteamID64, domain.MetricThroughSmokeKill)

			if e.Weapon != nil {
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  e.Weapon.String(),
					weaponClass: e.Weapon.Class(),
				}, domain.MetricThroughSmokeKill)
			}
		}
	}

	if p.playerConnected(e.Assister) {
		// assist total amount
		p.metrics.incr(e.Assister.SteamID64, domain.MetricAssist)

		// flashbang assist amount
		if e.AssistedFlash {
			p.metrics.incr(e.Assister.SteamID64, domain.MetricFlashbangAssist)
		}
	}
}

// handlePlayerHurt collects metrics and weapon metrics on player hurt.
func (p *parser) handlePlayerHurt(e events.PlayerHurt) {
	if !p.collectStats(p.GameState()) {
		return
	}

	if p.playerConnected(e.Attacker) {
		// dealt damage
		p.metrics.add(e.Attacker.SteamID64, domain.MetricDamageDealt, e.HealthDamage)

		if e.Weapon != nil {
			p.weaponMetrics.add(e.Attacker.SteamID64, weaponMetric{
				weaponName:  e.Weapon.String(),
				weaponClass: e.Weapon.Class(),
			}, domain.MetricDamageDealt, e.HealthDamage)
		}
	}

	if p.playerConnected(e.Player) {
		// taken damage
		p.metrics.add(e.Player.SteamID64, domain.MetricDamageTaken, e.HealthDamage)

		if e.Weapon != nil {
			p.weaponMetrics.add(e.Player.SteamID64, weaponMetric{
				weaponName:  e.Weapon.String(),
				weaponClass: e.Weapon.Class(),
			}, domain.MetricDamageTaken, e.HealthDamage)
		}
	}
}
