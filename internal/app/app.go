package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ssssargsian/uniplay/internal/domain"

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

	stats := domain.NewPlayerStats()
	score := domain.NewRoundScore()

	// handle score
	p.RegisterEventHandler(func(e events.RoundEnd) {
		gs := p.GameState()
		switch e.Winner {
		case common.TeamTerrorists:
			score.Update(gs.TeamTerrorists().Score()+1, gs.TeamCounterTerrorists().Score())
		case common.TeamCounterTerrorists:
			score.Update(gs.TeamTerrorists().Score(), gs.TeamCounterTerrorists().Score()+1)
		default:
			// tie
		}
	})

	// handle kills
	p.RegisterEventHandler(func(e events.Kill) {
		if e.Victim != nil {
			stats.Add(e.Victim.SteamID64, domain.MetricDeath, 1)
		}

		if e.Killer != nil {
			stats.Add(e.Killer.SteamID64, domain.MetricKill, 1)

			switch {
			case e.IsHeadshot:
				stats.Add(e.Killer.SteamID64, domain.MetricHSKill, 1)
			case e.AttackerBlind:
				stats.Add(e.Killer.SteamID64, domain.MetricBlindKill, 1)
			case e.IsWallBang():
				stats.Add(e.Killer.SteamID64, domain.MetricWallbangKill, 1)
			case e.NoScope:
				stats.Add(e.Killer.SteamID64, domain.MetricNoScopeKill, 1)
			case e.ThroughSmoke:
				stats.Add(e.Killer.SteamID64, domain.MetricThroughSmokeKill, 1)
			}
		}

		if e.Assister != nil {
			stats.Add(e.Assister.SteamID64, domain.MetricAssist, 1)

			switch {
			case e.AssistedFlash:
				stats.Add(e.Assister.SteamID64, domain.MetricFlashbangAssist, 1)
			}
		}
	})

	// handle player damage taken or dealt
	p.RegisterEventHandler(func(e events.PlayerHurt) {
		if e.Attacker != nil {
			stats.Add(e.Attacker.SteamID64, domain.MetricDamageDealt, uint16(e.HealthDamage)+uint16(e.ArmorDamage))
		}

		if e.Player != nil {
			stats.Add(e.Player.SteamID64, domain.MetricDamageTaken, uint16(e.HealthDamage)+uint16(e.ArmorDamage))
		}

	})

	// handle mvp of the round
	p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if e.Player != nil {
			stats.Add(e.Player.SteamID64, domain.MetricRoundMVP, 1)
		}
	})

	if err = p.ParseToEnd(); err != nil {
		log.Fatalf("failed to parse demo: %s", err.Error())
	}

	fmt.Println(stats)
	fmt.Println(score)
}
