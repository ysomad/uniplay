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
			stats.Incr(e.Victim.SteamID64, domain.MetricDeath)
		}

		if e.Killer != nil {
			stats.Incr(e.Killer.SteamID64, domain.MetricKill)

			switch {
			case e.IsHeadshot:
				stats.Incr(e.Killer.SteamID64, domain.MetricHSKill)
			case e.AttackerBlind:
				stats.Incr(e.Killer.SteamID64, domain.MetricBlindKill)
			case e.IsWallBang():
				stats.Incr(e.Killer.SteamID64, domain.MetricWallbangKill)
			case e.NoScope:
				stats.Incr(e.Killer.SteamID64, domain.MetricNoScopeKill)
			case e.ThroughSmoke:
				stats.Incr(e.Killer.SteamID64, domain.MetricThroughSmokeKill)
			}
		}

		if e.Assister != nil {
			stats.Incr(e.Assister.SteamID64, domain.MetricAssist)

			switch {
			case e.AssistedFlash:
				stats.Incr(e.Assister.SteamID64, domain.MetricFlashbangAssist)
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
			stats.Incr(e.Player.SteamID64, domain.MetricRountMVPCount)
		}
	})

	if err = p.ParseToEnd(); err != nil {
		log.Fatalf("failed to parse demo: %s", err.Error())
	}

	fmt.Println(stats)
	fmt.Println(score)
}
