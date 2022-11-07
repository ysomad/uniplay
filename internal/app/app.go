package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/metric"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	common "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
)

func Run() {
	demo, err := os.Open("./test-data/1.dem")
	if err != nil {
		log.Fatalf("failed to open demo file: %s", err.Error())
	}
	defer demo.Close()

	p := demoinfocs.NewParser(demo)
	defer p.Close()

	metrics := domain.NewPlayerMetrics()
	weaponMetrics := domain.NewPlayerWeaponMetrics()
	score := domain.NewRoundScore()

	// handle score
	p.RegisterEventHandler(func(e events.RoundEnd) {
		gs := p.GameState()
		switch e.Winner {
		case common.TeamTerrorists:
			score.Set(gs.TeamTerrorists().Score()+1, gs.TeamCounterTerrorists().Score())
		case common.TeamCounterTerrorists:
			score.Set(gs.TeamTerrorists().Score(), gs.TeamCounterTerrorists().Score()+1)
		default:
			// tie
		}
	})

	// handle kills
	p.RegisterEventHandler(func(e events.Kill) {
		if e.Victim != nil {
			metrics.Incr(e.Victim.SteamID64, metric.Death)
			weaponMetrics.Add(e.Victim.SteamID64, metric.WeaponMetric{
				Metric: metric.Death,
				Weapon: e.Weapon.String(),
				Damage: metric.DamageKill,
			})
		}

		if e.Killer != nil {
			// kills total
			metrics.Incr(e.Killer.SteamID64, metric.Kill)
			weaponMetrics.Add(e.Killer.SteamID64, metric.WeaponMetric{
				Metric: metric.Kill,
				Weapon: e.Weapon.String(),
				Damage: metric.DamageKill,
			})

			switch {

			// headshot kills
			case e.IsHeadshot:
				metrics.Incr(e.Killer.SteamID64, metric.HSKill)
				weaponMetrics.Add(e.Killer.SteamID64, metric.WeaponMetric{
					Metric: metric.HSKill,
					Weapon: e.Weapon.String(),
					Damage: metric.DamageKill,
				})

			// blind kills
			case e.AttackerBlind:
				metrics.Incr(e.Killer.SteamID64, metric.BlindKill)
				weaponMetrics.Add(e.Killer.SteamID64, metric.WeaponMetric{
					Metric: metric.BlindKill,
					Weapon: e.Weapon.String(),
					Damage: metric.DamageKill,
				})

			// wallbang kills
			case e.IsWallBang():
				metrics.Incr(e.Killer.SteamID64, metric.WallbangKill)
				weaponMetrics.Add(e.Killer.SteamID64, metric.WeaponMetric{
					Metric: metric.WallbangKill,
					Weapon: e.Weapon.String(),
					Damage: metric.DamageKill,
				})

			// noscope kills
			case e.NoScope:
				metrics.Incr(e.Killer.SteamID64, metric.NoScopeKill)
				weaponMetrics.Add(e.Killer.SteamID64, metric.WeaponMetric{
					Metric: metric.NoScopeKill,
					Weapon: e.Weapon.String(),
					Damage: metric.DamageKill,
				})

			// kills through smoke
			case e.ThroughSmoke:
				metrics.Incr(e.Killer.SteamID64, metric.ThroughSmokeKill)
				weaponMetrics.Add(e.Killer.SteamID64, metric.WeaponMetric{
					Metric: metric.ThroughSmokeKill,
					Weapon: e.Weapon.String(),
					Damage: metric.DamageKill,
				})
			}
		}

		if e.Assister != nil {
			// assist with weapon
			metrics.Incr(e.Assister.SteamID64, metric.Assist)
			weaponMetrics.Add(e.Assister.SteamID64, metric.WeaponMetric{
				Metric: metric.Assist,
				Weapon: e.Weapon.String(),
				Damage: metric.DamageAssist,
			})

			// assist with flash
			if e.AssistedFlash {
				metrics.Incr(e.Assister.SteamID64, metric.FlashbangAssist)
				weaponMetrics.Add(e.Assister.SteamID64, metric.WeaponMetric{
					Metric: metric.FlashbangAssist,
					Weapon: e.Weapon.String(),
					Damage: metric.DamageAssist,
				})
			}
		}
	})

	// handle player damage taken or dealt
	p.RegisterEventHandler(func(e events.PlayerHurt) {
		totalDamage := uint16(e.HealthDamage) + uint16(e.ArmorDamage)      // total damage
		damage := uint16(e.HealthDamageTaken) + uint16(e.ArmorDamageTaken) // damage excluding over damage

		if e.Attacker != nil {
			// dealt damage
			metrics.Add(e.Attacker.SteamID64, metric.DamageDealt, damage)
			weaponMetrics.Add(e.Attacker.SteamID64, metric.WeaponMetric{
				Metric: metric.DamageDealt,
				Weapon: e.Weapon.String(),
				Damage: damage,
			})

			// dealt total damage
			metrics.Add(e.Attacker.SteamID64, metric.DamageDealtWithOver, totalDamage)
			weaponMetrics.Add(e.Attacker.SteamID64, metric.WeaponMetric{
				Metric: metric.DamageDealtWithOver,
				Weapon: e.Weapon.String(),
				Damage: totalDamage,
			})
		}

		if e.Player != nil {
			// damage taken
			metrics.Add(e.Player.SteamID64, metric.DamageTaken, damage)
			weaponMetrics.Add(e.Player.SteamID64, metric.WeaponMetric{
				Metric: metric.DamageTaken,
				Weapon: e.Weapon.String(),
				Damage: damage,
			})

			// total damage taken
			metrics.Add(e.Player.SteamID64, metric.DamageTakenWithOver, totalDamage)
			weaponMetrics.Add(e.Player.SteamID64, metric.WeaponMetric{
				Metric: metric.DamageTakenWithOver,
				Weapon: e.Weapon.String(),
				Damage: totalDamage,
			})
		}
	})

	// handle mvp of the round
	p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if e.Player != nil {
			metrics.Incr(e.Player.SteamID64, metric.RoundMVPCount)
		}
	})

	if err = p.ParseToEnd(); err != nil {
		log.Fatalf("failed to parse demo: %s", err.Error())
	}

	fmt.Println(metrics)
	fmt.Println(score)
	fmt.Println(weaponMetrics)
}
