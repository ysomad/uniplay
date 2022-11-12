package replayparser

import (
	"io"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"

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
	match         *dto.CreateMatchArgs

	isKnifeRound bool
}

func New(r io.Reader) *parser {
	return &parser{
		demoinfocs.NewParser(r),
		newPlayerMetrics(),
		newWeaponMetrics(),
		&dto.CreateMatchArgs{},
		false,
	}
}

func (p *parser) Parse() (parseResult, error) {
	p.RegisterEventHandler(func(e events.RoundFreezetimeEnd) {
		p.detectKnifeRound()
	})

	p.handleKills()
	p.handlePlayerHurt()
	p.handleScoreUpdate()
	p.handleMVPAnnouncement()
	p.handleBombEvents()

	p.RegisterEventHandler(func(e events.AnnouncementWinPanelMatch) {
		p.match.MapName = p.Header().MapName
		p.match.Duration = p.Header().PlaybackTime
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
	if p.isKnifeRound || !gs.IsMatchStarted() {
		return false
	}

	return true
}

// detectKnifeRound sets isKnifeRound to true if any player has no secondary weapon and first slot is a knife.
func (p *parser) detectKnifeRound() {
	p.isKnifeRound = false

	for _, player := range p.GameState().TeamCounterTerrorists().Members() {
		weapons := player.Weapons()
		if len(weapons) == 1 && weapons[0].Type == common.EqKnife {
			p.isKnifeRound = true
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

		if e.Victim != nil {
			// death amount FROM weapon
			p.metrics.incr(e.Victim.SteamID64, domain.MetricDeath)
			p.weaponMetrics.incr(e.Victim.SteamID64, weaponMetric{
				weaponName:  weapon,
				weaponClass: weaponClass,
			}, domain.MetricDeath)
		}

		if e.Killer != nil {
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

		if e.Assister != nil {
			// assist total amount
			p.metrics.incr(e.Assister.SteamID64, domain.MetricAssist)

			// flashbang assist amount
			if e.AssistedFlash {
				p.metrics.incr(e.Assister.SteamID64, domain.MetricFlashbangAssist)
			}
		}
	})
}

// handleScoreUpdate updates match teams score on ScoreUpdated event.
func (p *parser) handleScoreUpdate() {
	p.RegisterEventHandler(func(e events.ScoreUpdated) {
		teamMembers := e.TeamState.Members()
		playerSteamIDs := make([]uint64, len(teamMembers))

		for i, player := range teamMembers {
			playerSteamIDs[i] = player.SteamID64
		}

		switch e.TeamState.Team() {
		case common.TeamTerrorists:
			p.match.Team1.SetAll(e.TeamState.ClanName(), e.TeamState.Flag(), uint8(e.TeamState.Score()), playerSteamIDs)
		case common.TeamCounterTerrorists:
			p.match.Team2.SetAll(e.TeamState.ClanName(), e.TeamState.Flag(), uint8(e.TeamState.Score()), playerSteamIDs)
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

		if e.Attacker != nil {
			// dealt damage
			p.metrics.add(e.Attacker.SteamID64, domain.MetricDamageDealt, e.HealthDamage)
			p.weaponMetrics.add(e.Attacker.SteamID64, weaponMetric{
				weaponName:  weaponName,
				weaponClass: weaponClass,
			}, domain.MetricDamageDealt, e.HealthDamage)
		}

		if e.Player != nil {
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

		if e.BombEvent.Player != nil {
			p.metrics.incr(e.BombEvent.Player.SteamID64, domain.MetricBombDefused)
		}
	})

	p.RegisterEventHandler(func(e events.BombPlanted) {
		if !p.collectStats(p.GameState()) {
			return
		}

		if e.BombEvent.Player != nil {
			p.metrics.incr(e.BombEvent.Player.SteamID64, domain.MetricBombPlanted)
		}
	})
}

// handleMVPAnnouncement increments player mvp amount metric on RoundMVPAnnouncement event.
func (p *parser) handleMVPAnnouncement() {
	p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if !p.collectStats(p.GameState()) {
			return
		}

		if e.Player != nil {
			p.metrics.incr(e.Player.SteamID64, domain.MetricRoundMVPCount)
		}
	})
}
