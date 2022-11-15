package replayparser

import (
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

// TODO: refactor
func (p *parser) Parse() (parseResult, error) {
	p.RegisterEventHandler(func(e events.RoundFreezetimeEnd) {
		p.detectKnifeRound()
	})

	p.RegisterEventHandler(func(e events.MatchStart) {
		gs := p.GameState()

		// set teams and players after match start
		t := gs.TeamTerrorists()
		p.match.setTeam1(newMatchTeam(t.ClanName(), t.Flag(), t.Team(), t.Members()))

		ct := gs.TeamCounterTerrorists()
		p.match.setTeam2(newMatchTeam(ct.ClanName(), ct.Flag(), ct.Team(), ct.Members()))
	})

	p.RegisterEventHandler(func(e events.TeamSideSwitch) {
		p.match.swapTeamSides()
	})

	p.RegisterEventHandler(func(e events.ScoreUpdated) {
		p.match.updateTeamsScore(e)
	})

	p.handleKills()
	p.handlePlayerHurt()
	p.handleBombEvents()

	p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if !p.collectStats(p.GameState()) {
			return
		}

		if e.Player != nil && e.Player.SteamID64 != 0 {
			p.metrics.incr(e.Player.SteamID64, domain.MetricRoundMVPCount)
		}
	})

	p.RegisterEventHandler(func(e events.AnnouncementWinPanelMatch) {
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

// handleKills counts all kills and weapon kills.
func (p *parser) handleKills() {
	p.RegisterEventHandler(func(e events.Kill) {
		if !p.collectStats(p.GameState()) {
			return
		}

		var (
			weapon      string
			weaponClass domain.EquipmentClass
		)

		if e.Weapon != nil {
			weapon = e.Weapon.String()
			weaponClass = domain.EquipmentClass(e.Weapon.Class())
		}

		if e.Victim != nil && e.Victim.SteamID64 != 0 {
			// death amount FROM weapon
			p.metrics.incr(e.Victim.SteamID64, domain.MetricDeath)
			p.weaponMetrics.incr(e.Victim.SteamID64, weaponMetric{
				weaponName:  weapon,
				weaponClass: weaponClass,
			}, domain.MetricDeath)
		}

		if e.Killer != nil && e.Killer.SteamID64 != 0 {
			// kill amount
			p.metrics.incr(e.Killer.SteamID64, domain.MetricKill)
			p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
				weaponName:  weapon,
				weaponClass: weaponClass,
			}, domain.MetricKill)

			switch {
			// headshot kill amount
			case e.IsHeadshot:
				p.metrics.incr(e.Killer.SteamID64, domain.MetricHSKill)
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  weapon,
					weaponClass: weaponClass,
				}, domain.MetricHSKill)

			// blind kill amount
			case e.AttackerBlind:
				p.metrics.incr(e.Killer.SteamID64, domain.MetricBlindKill)
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  weapon,
					weaponClass: weaponClass,
				}, domain.MetricBlindKill)

			// wallbang kill amount
			case e.IsWallBang():
				p.metrics.incr(e.Killer.SteamID64, domain.MetricWallbangKill)
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  weapon,
					weaponClass: weaponClass,
				}, domain.MetricWallbangKill)

			// noscope kill amount
			case e.NoScope:
				p.metrics.incr(e.Killer.SteamID64, domain.MetricNoScopeKill)
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  weapon,
					weaponClass: weaponClass,
				}, domain.MetricNoScopeKill)

			// through smoke kill amount
			case e.ThroughSmoke:
				p.metrics.incr(e.Killer.SteamID64, domain.MetricThroughSmokeKill)
				p.weaponMetrics.incr(e.Killer.SteamID64, weaponMetric{
					weaponName:  weapon,
					weaponClass: weaponClass,
				}, domain.MetricThroughSmokeKill)
			}
		}

		if e.Assister != nil && e.Assister.SteamID64 != 0 {
			// assist total amount
			p.metrics.incr(e.Assister.SteamID64, domain.MetricAssist)

			// flashbang assist amount
			if e.AssistedFlash {
				p.metrics.incr(e.Assister.SteamID64, domain.MetricFlashbangAssist)
			}
		}
	})
}

// handlePlayerHurt calculates metrics for taken, dealt damage and for taken and dealt damage by a weapon.
func (p *parser) handlePlayerHurt() {
	p.RegisterEventHandler(func(e events.PlayerHurt) {
		if !p.collectStats(p.GameState()) {
			return
		}

		var (
			weaponName  string
			weaponClass domain.EquipmentClass
		)
		if e.Weapon != nil {
			weaponName = e.Weapon.String()
			weaponClass = domain.EquipmentClass(e.Weapon.Class())
		}

		if e.Attacker != nil && e.Attacker.SteamID64 != 0 {
			// dealt damage
			p.metrics.add(e.Attacker.SteamID64, domain.MetricDamageDealt, e.HealthDamage)
			p.weaponMetrics.add(e.Attacker.SteamID64, weaponMetric{
				weaponName:  weaponName,
				weaponClass: weaponClass,
			}, domain.MetricDamageDealt, e.HealthDamage)
		}

		if e.Player != nil && e.Player.SteamID64 != 0 {
			// taken damage
			p.metrics.add(e.Player.SteamID64, domain.MetricDamageTaken, e.HealthDamage)
			p.weaponMetrics.add(e.Player.SteamID64, weaponMetric{
				weaponName:  weaponName,
				weaponClass: weaponClass,
			}, domain.MetricDamageTaken, e.HealthDamage)
		}
	})
}

// handleBomdEvent counts number of planted and defused bombs.
func (p *parser) handleBombEvents() {
	p.RegisterEventHandler(func(e events.BombDefused) {
		if !p.collectStats(p.GameState()) {
			return
		}

		if e.BombEvent.Player != nil && e.BombEvent.Player.SteamID64 != 0 {
			p.metrics.incr(e.BombEvent.Player.SteamID64, domain.MetricBombDefused)
		}
	})

	p.RegisterEventHandler(func(e events.BombPlanted) {
		if !p.collectStats(p.GameState()) {
			return
		}

		if e.BombEvent.Player != nil && e.BombEvent.Player.SteamID64 != 0 {
			p.metrics.incr(e.BombEvent.Player.SteamID64, domain.MetricBombPlanted)
		}
	})
}
