package demoparser

import (
	"io"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/metric"

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

func (p *parser) Parse() (*domain.PlayerMetrics, *domain.PlayerWeaponEvents, domain.RoundScore, error) {
	metrics := domain.NewPlayerMetrics()
	weaponEvents := domain.NewPlayerWeaponEvents()
	roundScore := domain.NewRoundScore()

	// handle score
	p.RegisterEventHandler(func(e events.RoundEnd) {
		gs := p.GameState()
		switch e.Winner {
		case common.TeamTerrorists:
			roundScore.Set(gs.TeamTerrorists().Score()+1, gs.TeamCounterTerrorists().Score())
		case common.TeamCounterTerrorists:
			roundScore.Set(gs.TeamTerrorists().Score(), gs.TeamCounterTerrorists().Score()+1)
		default:
			// tie
		}
	})

	// handle kills
	p.RegisterEventHandler(func(e events.Kill) {
		var (
			victimHealth uint16
			victimArmor  uint16
			weapon       string
		)

		if e.Weapon != nil {
			weapon = e.Weapon.String()
		}

		if e.Victim != nil {
			victimHealth, victimArmor = uint16(e.Victim.Health()), uint16(e.Victim.Armor())

			// total deaths
			metrics.Incr(e.Victim.SteamID64, metric.Death)
			weaponEvents.Add(e.Victim.SteamID64, metric.WeaponEvent{
				Event:        metric.Death,
				Weapon:       weapon,
				HealthDamage: victimHealth,
				ArmorDamage:  victimArmor,
			})
		}

		if e.Killer != nil {
			// kills total
			metrics.Incr(e.Killer.SteamID64, metric.Kill)
			weaponEvents.Add(e.Killer.SteamID64, metric.WeaponEvent{
				Event:        metric.Kill,
				Weapon:       weapon,
				HealthDamage: victimHealth,
				ArmorDamage:  victimArmor,
			})

			switch {
			// headshot kills
			case e.IsHeadshot:
				metrics.Incr(e.Killer.SteamID64, metric.HSKill)
				weaponEvents.Add(e.Killer.SteamID64, metric.WeaponEvent{
					Event:        metric.HSKill,
					Weapon:       weapon,
					HealthDamage: victimHealth,
					ArmorDamage:  victimArmor,
				})

			// blind kills
			case e.AttackerBlind:
				metrics.Incr(e.Killer.SteamID64, metric.BlindKill)
				weaponEvents.Add(e.Killer.SteamID64, metric.WeaponEvent{
					Event:        metric.BlindKill,
					Weapon:       weapon,
					HealthDamage: victimHealth,
					ArmorDamage:  victimArmor,
				})

			// wallbang kills
			case e.IsWallBang():
				metrics.Incr(e.Killer.SteamID64, metric.WallbangKill)
				weaponEvents.Add(e.Killer.SteamID64, metric.WeaponEvent{
					Event:        metric.WallbangKill,
					Weapon:       weapon,
					HealthDamage: victimHealth,
					ArmorDamage:  victimArmor,
				})

			// noscope kills
			case e.NoScope:
				metrics.Incr(e.Killer.SteamID64, metric.NoScopeKill)
				weaponEvents.Add(e.Killer.SteamID64, metric.WeaponEvent{
					Event:        metric.NoScopeKill,
					Weapon:       weapon,
					HealthDamage: victimHealth,
					ArmorDamage:  victimArmor,
				})

			// kills through smoke
			case e.ThroughSmoke:
				metrics.Incr(e.Killer.SteamID64, metric.ThroughSmokeKill)
				weaponEvents.Add(e.Killer.SteamID64, metric.WeaponEvent{
					Event:        metric.ThroughSmokeKill,
					Weapon:       weapon,
					HealthDamage: victimHealth,
					ArmorDamage:  victimArmor,
				})
			}
		}

		if e.Assister != nil {
			// total assists
			metrics.Incr(e.Assister.SteamID64, metric.Assist)

			// total assists with flash
			if e.AssistedFlash {
				metrics.Incr(e.Assister.SteamID64, metric.FlashbangAssist)
			}
		}
	})

	// handle player damage taken or dealt
	p.RegisterEventHandler(func(e events.PlayerHurt) {
		totalDamage := uint16(e.HealthDamage) + uint16(e.ArmorDamage)      // total damage
		damage := uint16(e.HealthDamageTaken) + uint16(e.ArmorDamageTaken) // damage excluding over damage

		var weapon string

		if e.Weapon != nil {
			weapon = e.Weapon.String()
		}

		if e.Attacker != nil {
			// dealt damage without over damage
			metrics.Add(e.Attacker.SteamID64, metric.DamageDealt, damage)
			weaponEvents.Add(e.Attacker.SteamID64, metric.WeaponEvent{
				Event:        metric.DamageDealt,
				Weapon:       weapon,
				HealthDamage: uint16(e.HealthDamageTaken),
				ArmorDamage:  uint16(e.ArmorDamageTaken),
			})

			// dealt total damage with over damage
			metrics.Add(e.Attacker.SteamID64, metric.DamageDealtWithOver, totalDamage)
			weaponEvents.Add(e.Attacker.SteamID64, metric.WeaponEvent{
				Event:        metric.DamageDealtWithOver,
				Weapon:       weapon,
				HealthDamage: uint16(e.HealthDamage),
				ArmorDamage:  uint16(e.ArmorDamage),
			})
		}

		if e.Player != nil {
			// damage taken without over damage
			metrics.Add(e.Player.SteamID64, metric.DamageTaken, damage)
			weaponEvents.Add(e.Player.SteamID64, metric.WeaponEvent{
				Event:        metric.DamageTaken,
				Weapon:       weapon,
				HealthDamage: uint16(e.HealthDamageTaken),
				ArmorDamage:  uint16(e.ArmorDamageTaken),
			})

			// total damage taken with over damage
			metrics.Add(e.Player.SteamID64, metric.DamageTakenWithOver, totalDamage)
			weaponEvents.Add(e.Player.SteamID64, metric.WeaponEvent{
				Event:        metric.DamageTakenWithOver,
				Weapon:       weapon,
				HealthDamage: uint16(e.HealthDamage),
				ArmorDamage:  uint16(e.ArmorDamage),
			})
		}
	})

	// handle mvp of the round
	p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if e.Player != nil {
			metrics.Incr(e.Player.SteamID64, metric.RoundMVPCount)
		}
	})

	if err := p.ParseToEnd(); err != nil {
		return nil, nil, domain.RoundScore{}, err
	}

	return metrics, weaponEvents, roundScore, nil
}
