package demoparser

import (
	"io"

	"github.com/ssssargsian/uniplay/internal/domain"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
)

type parser struct {
	demoinfocs.Parser
}

func New(r io.Reader) *parser {
	return &parser{demoinfocs.NewParser(r)}
}

func (p *parser) Parse() (pm *domain.PlayerMetrics, wm *domain.WeaponMetrics, match domain.Match, err error) {
	metrics := domain.NewPlayerMetrics()
	weaponMetrics := domain.NewWeaponMetrics()

	var (
		team1 domain.MatchTeam
		team2 domain.MatchTeam
	)

	// handle match end
	p.RegisterEventHandler(func(e events.AnnouncementWinPanelMatch) {
		match.Map = p.Header().MapName
		match.Duration = p.Header().PlaybackTime
	})

	// handle score update
	p.RegisterEventHandler(func(e events.ScoreUpdated) {
		switch e.TeamState.Team() {
		case common.TeamCounterTerrorists:
			team1.SetAll(e.TeamState.ClanName(), e.TeamState.Flag(), e.TeamState.Score())
		case common.TeamTerrorists:
			team2.SetAll(e.TeamState.ClanName(), e.TeamState.Flag(), e.TeamState.Score())
		}
	})

	// handle kills
	p.RegisterEventHandler(func(e events.Kill) {
		if !p.GameState().IsMatchStarted() {
			return
		}

		var (
			weapon string
		)

		if e.Weapon != nil {
			weapon = e.Weapon.String()
		}

		if e.Victim != nil {
			// количество смертей ОТ оружия
			metrics.Incr(e.Victim.SteamID64, domain.MetricDeath)
			weaponMetrics.Incr(e.Victim.SteamID64, domain.Weapon(weapon), domain.MetricDeath)
		}

		if e.Killer != nil {
			// количество убийств оружием
			metrics.Incr(e.Killer.SteamID64, domain.MetricKill)
			weaponMetrics.Incr(e.Killer.SteamID64, domain.Weapon(weapon), domain.MetricKill)

			switch {
			// количество хс оружием
			case e.IsHeadshot:
				metrics.Incr(e.Killer.SteamID64, domain.MetricHSKill)
				weaponMetrics.Incr(e.Killer.SteamID64, domain.Weapon(weapon), domain.MetricHSKill)

			// слепых убийств оружием
			case e.AttackerBlind:
				metrics.Incr(e.Killer.SteamID64, domain.MetricBlindKill)
				weaponMetrics.Incr(e.Killer.SteamID64, domain.Weapon(weapon), domain.MetricBlindKill)

			// вб убийств оружием
			case e.IsWallBang():
				metrics.Incr(e.Killer.SteamID64, domain.MetricWallbangKill)
				weaponMetrics.Incr(e.Killer.SteamID64, domain.Weapon(weapon), domain.MetricWallbangKill)

			// убийств без прицела оружием
			case e.NoScope:
				metrics.Incr(e.Killer.SteamID64, domain.MetricNoScopeKill)
				weaponMetrics.Incr(e.Killer.SteamID64, domain.Weapon(weapon), domain.MetricNoScopeKill)

			// убийств через смоук оружием
			case e.ThroughSmoke:
				metrics.Incr(e.Killer.SteamID64, domain.MetricThroughSmokeKill)
				weaponMetrics.Incr(e.Killer.SteamID64, domain.Weapon(weapon), domain.MetricThroughSmokeKill)
			}
		}

		if e.Assister != nil {
			// всего кол-во ассистов
			metrics.Incr(e.Assister.SteamID64, domain.MetricAssist)

			// кол-во ассистов флешкой
			if e.AssistedFlash {
				metrics.Incr(e.Assister.SteamID64, domain.MetricFlashbangAssist)
			}
		}
	})

	// handle player damage taken or dealt
	p.RegisterEventHandler(func(e events.PlayerHurt) {
		if !p.GameState().IsMatchStarted() {
			return
		}

		var weapon string

		if e.Weapon != nil {
			weapon = e.Weapon.String()
		}

		if e.Attacker != nil {
			// нанесено урона оружием
			metrics.Add(e.Attacker.SteamID64, domain.MetricDamageDealt, e.HealthDamage)
			weaponMetrics.Add(e.Attacker.SteamID64, domain.Weapon(weapon), domain.MetricDamageDealt, e.HealthDamage)
		}

		if e.Player != nil {
			// получено урона оружием
			metrics.Add(e.Player.SteamID64, domain.MetricDamageTaken, e.HealthDamage)
			weaponMetrics.Add(e.Player.SteamID64, domain.Weapon(weapon), domain.MetricDamageTaken, e.HealthDamage)
		}
	})

	// handle mvp of the round
	p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if !p.GameState().IsMatchStarted() {
			return
		}

		if e.Player != nil {
			metrics.Incr(e.Player.SteamID64, domain.MetricRoundMVPCount)
		}
	})

	// handle bomb defuse
	p.RegisterEventHandler(func(e events.BombDefused) {
		if !p.GameState().IsMatchStarted() {
			return
		}

		if e.BombEvent.Player != nil {
			metrics.Incr(e.BombEvent.Player.SteamID64, domain.MetricBombDefused)
		}
	})

	// handle bomb plant
	p.RegisterEventHandler(func(e events.BombPlanted) {
		if !p.GameState().IsMatchStarted() {
			return
		}

		if e.BombEvent.Player != nil {
			metrics.Incr(e.BombEvent.Player.SteamID64, domain.MetricBombPlanted)
		}
	})

	if err := p.ParseToEnd(); err != nil {
		return nil, nil, domain.Match{}, err
	}

	match.Team1 = team1
	match.Team2 = team2

	return metrics, weaponMetrics, match, nil
}
