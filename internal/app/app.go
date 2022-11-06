package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ssssargsian/uniplay/internal/domain"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
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

	p.RegisterEventHandler(func(e events.Kill) {
		if e.Victim != nil {
			stats.Incr(e.Victim.SteamID64, domain.EventDeath)
		}

		if e.Killer != nil {
			stats.Incr(e.Killer.SteamID64, domain.EventKill)

			switch {
			case e.IsHeadshot:
				stats.Incr(e.Killer.SteamID64, domain.EventHSKill)
			case e.AttackerBlind:
				stats.Incr(e.Killer.SteamID64, domain.EventBlindKill)
			case e.IsWallBang():
				stats.Incr(e.Killer.SteamID64, domain.EventWallbangKill)
			case e.NoScope:
				stats.Incr(e.Killer.SteamID64, domain.EventNoScopeKill)
			case e.ThroughSmoke:
				stats.Incr(e.Killer.SteamID64, domain.EventThroughSmokeKill)
			}
		}

		if e.Assister != nil {
			stats.Incr(e.Assister.SteamID64, domain.EventAssist)

			switch {
			case e.AssistedFlash:
				stats.Incr(e.Assister.SteamID64, domain.EventFlashbangAssist)
			}
		}
	})

	if err = p.ParseToEnd(); err != nil {
		log.Fatalf("failed to parse demo: %s", err.Error())
	}

	fmt.Println(stats.Slug())
}
