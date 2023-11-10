package demoparser

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type parser struct {
	demoinfocs.Parser
	playerStats playerStatsMap
	weaponStats weaponStatsMap
	gameState   *gameState
	demosize    int64
}

func New(d Demo) *parser {
	return &parser{
		Parser:      demoinfocs.NewParser(d),
		demosize:    d.size,
		gameState:   &gameState{},
		playerStats: make(playerStatsMap, 20),
		weaponStats: make(weaponStatsMap, 20),
	}
}

func (p *parser) Parse() error {
	h, err := p.ParseHeader()
	if err != nil {
		return fmt.Errorf("demo header not parsed: %w", err)
	}

	dh := &demoHeader{
		server:         h.ServerName,
		client:         h.ClientName,
		mapName:        h.MapName,
		playbackTicks:  h.PlaybackTicks,
		playbackFrames: h.PlaybackFrames,
		signonLength:   h.SignonLength,
		playbackTime:   h.PlaybackTime,
		filesize:       p.demosize,
		uploadedAt:     time.Now(),
	}

	if err := dh.validate(); err != nil {
		return fmt.Errorf("parsed invalid demo header: %w", err)
	}

	p.RegisterEventHandler(p.killHandler)
	p.RegisterEventHandler(p.hurtHandler)
	p.RegisterEventHandler(p.weaponFireHandler)
	p.RegisterEventHandler(p.playerFlashedHandler)

	p.RegisterEventHandler(p.roundFreezetimeEndHandler)
	p.RegisterEventHandler(p.teamSideSwitchHandler)
	p.RegisterEventHandler(p.roundMVPAnnouncementHandler)

	p.RegisterEventHandler(p.bombDefusedHandler)
	p.RegisterEventHandler(p.bombPlantedHandler)

	return nil
}

func (p *parser) roundFreezetimeEndHandler(_ events.RoundFreezetimeEnd) {
	p.gameState.detectKnifeRound(p.GameState().TeamTerrorists().Members())
}

func (p *parser) killHandler(e events.Kill) {
	if !p.gameState.collectStats() {
		slog.Info("skipped kill event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if playerConnected(e.Killer) {
		p.playerStats.incr(e.Killer.SteamID64, eventKill)
		p.weaponStats.incr(e.Killer.SteamID64, eventKill, e.Weapon.Type)

		if e.IsHeadshot {
			p.playerStats.incr(e.Killer.SteamID64, eventHSKill)
			p.weaponStats.incr(e.Killer.SteamID64, eventHSKill, e.Weapon.Type)
		}

		if e.AttackerBlind {
			p.playerStats.incr(e.Killer.SteamID64, eventBlindKill)
			p.weaponStats.incr(e.Killer.SteamID64, eventBlindKill, e.Weapon.Type)
		}

		if e.IsWallBang() {
			p.playerStats.incr(e.Killer.SteamID64, eventWBKill)
			p.weaponStats.incr(e.Killer.SteamID64, eventWBKill, e.Weapon.Type)
		}

		if e.NoScope {
			p.playerStats.incr(e.Killer.SteamID64, eventNoScopeKill)
			p.weaponStats.incr(e.Killer.SteamID64, eventNoScopeKill, e.Weapon.Type)
		}

		if e.ThroughSmoke {
			p.playerStats.incr(e.Killer.SteamID64, eventSmokeKill)
			p.weaponStats.incr(e.Killer.SteamID64, eventSmokeKill, e.Weapon.Type)
		}
	} else {
		slog.Error("kill by unconnected player")
	}

	if playerConnected(e.Victim) {
		p.playerStats.incr(e.Victim.SteamID64, eventDeath)
		p.weaponStats.incr(e.Victim.SteamID64, eventDeath, e.Weapon.Type)
	} else {
		slog.Error("killed unconnected player")
	}

	if playerConnected(e.Assister) {
		p.playerStats.incr(e.Assister.SteamID64, eventAssist)
		p.weaponStats.incr(e.Assister.SteamID64, eventAssist, e.Weapon.Type)

		if e.AssistedFlash {
			p.playerStats.incr(e.Assister.SteamID64, eventFBAssist)
		}
	} else {
		slog.Error("kill assist by unconnected player")
	}
}

func (p *parser) hurtHandler(e events.PlayerHurt) {
	if !p.gameState.collectStats() {
		slog.Info("skipped player hurt event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if playerConnected(e.Player) {
		p.playerStats.add(e.Player.SteamID64, eventDmgTaken, e.HealthDamage)
		p.weaponStats.add(e.Player.SteamID64, eventDmgTaken, e.Weapon.Type, e.HealthDamage)
	} else {
		slog.Error("unconnected player got hurt")
	}

	if playerConnected(e.Attacker) {
		p.playerStats.add(e.Attacker.SteamID64, eventDmgDealt, e.HealthDamage)
		p.weaponStats.add(e.Attacker.SteamID64, eventDmgDealt, e.Weapon.Type, e.HealthDamage)
		p.weaponStats.incr(e.Attacker.SteamID64, p.hitgroupToEvent(e.HitGroup), e.Weapon.Type)

		if e.Weapon.Class() == common.EqClassGrenade {
			p.playerStats.add(e.Attacker.SteamID64, eventDmgGrenadeDealt, e.HealthDamage)
		}
	} else {
		slog.Error("unconnected attacker hurt other player")
	}
}

func (p *parser) teamSideSwitchHandler(_ events.TeamSideSwitch) {
	if err := p.gameState.teamA.swapSide(); err != nil {
		slog.Error("team A side not swapped", "err", err.Error())
	}
	if err := p.gameState.teamB.swapSide(); err != nil {
		slog.Error("team B side not swapped", "err", err.Error())
	}
}

func (p *parser) weaponFireHandler(e events.WeaponFire) {
	if !p.gameState.collectStats() {
		slog.Info("skipped weapon fire event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if !playerConnected(e.Shooter) {
		slog.Error("fire from unconnected player")
		return
	}

	if e.Weapon == nil {
		slog.Error("fire from nil weapon",
			"steam_id", e.Shooter.SteamID64,
			"name", e.Shooter.Name)
		return
	}

	if !equipValid(e.Weapon.Type) {
		slog.Error("fire from invalid weapon",
			"steam_id", e.Shooter.SteamID64,
			"name", e.Shooter.Name,
			"weapon", e.Weapon.String())
		return
	}

	p.weaponStats.incr(e.Shooter.SteamID64, eventShot, e.Weapon.Type)
}

func (p *parser) bombPlantedHandler(e events.BombPlanted) {
	if !p.gameState.collectStats() {
		slog.Info("skipped bomd planted event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if !playerConnected(e.Player) {
		slog.Error("bomb planted by unconnected player", "site", e.Site)
		return
	}

	p.playerStats.incr(e.Player.SteamID64, eventBombPlanted)
}

func (p *parser) bombDefusedHandler(e events.BombDefused) {
	if !p.gameState.collectStats() {
		slog.Info("skipped bomd defused event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if !playerConnected(e.Player) {
		slog.Error("bomb defused by unconnected player", "site", e.Site)
		return
	}

	p.playerStats.incr(e.Player.SteamID64, eventBombDefused)
}

func (p *parser) playerFlashedHandler(e events.PlayerFlashed) {
	if !p.gameState.collectStats() {
		slog.Info("skipped player flashed event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if playerSpectator(e.Player) {
		slog.Info("flashed spectator")
		return
	}

	if playerConnected(e.Player) {
		p.playerStats.incr(e.Player.SteamID64, eventBecameBlind)
	} else {
		slog.Error("flashed by unconnected player")
	}

	if playerConnected(e.Attacker) {
		p.playerStats.incr(e.Attacker.SteamID64, eventBlindedPlayer)
	} else {
		slog.Error("unconnected player flashed other player")
	}
}

func (p *parser) roundMVPAnnouncementHandler(e events.RoundMVPAnnouncement) {
	if !p.gameState.collectStats() {
		slog.Info("skipped mvp announcement event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if !playerConnected(e.Player) {
		slog.Error("announced mvp for unconnected player")
		return
	}

	p.playerStats.incr(e.Player.SteamID64, eventRoundMVP)
}
