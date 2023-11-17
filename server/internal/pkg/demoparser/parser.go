package demoparser

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"os"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msgs2"
)

type parser struct {
	demoinfocs.Parser
	playerStats playerStatsMap
	weaponStats weaponStatsMap
	gameState   *gameState
	rounds      roundHistory
	demosize    int64
}

func New(demofile io.Reader, demoheader *multipart.FileHeader) (*parser, error) {
	d, err := newDemo(demofile, demoheader)
	if err != nil {
		return nil, fmt.Errorf("demo not created: %w", err)
	}
	return &parser{
		Parser:      demoinfocs.NewParser(d),
		demosize:    d.size,
		gameState:   &gameState{},
		playerStats: make(playerStatsMap, 20),
		weaponStats: make(weaponStatsMap, 20),
		rounds:      newRoundHistory(),
	}, nil
}

func (p *parser) Parse() error {
	h, err := p.ParseHeader()
	if err != nil {
		return fmt.Errorf("demo header not parsed: %w", err)
	}

	slog.Info("parsed demo header", "header", h)

	dh := &demoHeader{
		server:         h.ServerName,
		client:         h.ClientName,
		mapName:        h.MapName,
		playbackTicks:  h.PlaybackTicks,
		playbackFrames: h.PlaybackFrames,
		signonLength:   h.SignonLength,
		playbackTime:   h.PlaybackTime,
		filesize:       p.demosize,
	}

	// TODO: skip validation because file header comes empty from demoinfocs
	// if err := dh.validate(); err != nil {
	// 	return fmt.Errorf("parsed invalid demo header: %w", err)
	// }

	slog.Info("generated demo id", "id", dh.uuid())

	p.RegisterEventHandler(p.killHandler)
	p.RegisterEventHandler(p.hurtHandler)
	p.RegisterEventHandler(p.weaponFireHandler)
	p.RegisterEventHandler(p.playerFlashedHandler)

	p.RegisterEventHandler(p.matchStartHandler)
	// p.RegisterEventHandler(p.matchStartedChangedHandler)
	p.RegisterEventHandler(p.roundFreezetimeEndHandler)
	p.RegisterEventHandler(p.teamSideSwitchHandler)
	p.RegisterEventHandler(p.roundMVPAnnouncementHandler)

	p.RegisterEventHandler(p.roundStartHandler)
	p.RegisterEventHandler(p.roundEndHandler)

	p.RegisterEventHandler(p.bombDefusedHandler)
	p.RegisterEventHandler(p.bombPlantedHandler)

	p.RegisterNetMessageHandler(func(msg *msgs2.CSVCMsg_ServerInfo) {
		slog.Info("server ip", "val", msg.GameSessionConfig.GetServerIpAddress())
		slog.Info("map", "val", msg.GameSessionConfig.GetS1Mapname())
	})

	if err = p.ParseToEnd(); err != nil {
		return fmt.Errorf("demo not parsed: %w", err)
	}

	p.weaponStats.calculateUnobtainableStats()

	playerStatsBytes, err := json.MarshalIndent(p.playerStats, "", "  ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("playerstats.json", playerStatsBytes, 0o644); err != nil {
		return err
	}

	weaponStatsBytes, err := json.MarshalIndent(p.weaponStats, "", "  ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("weaponstats.json", weaponStatsBytes, 0o644); err != nil {
		return err
	}

	roundsBytes, err := json.MarshalIndent(p.rounds, "", "  ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("rounds.json", roundsBytes, 0o644); err != nil {
		return err
	}

	return nil
}

func (p *parser) roundStartHandler(e events.RoundStart) {
	if !p.gameState.collectStats() {
		slog.Info("round start event skip",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	p.rounds.start(p.GameState().TeamTerrorists())
}

func (p *parser) roundEndHandler(e events.RoundEnd) {
	if !p.gameState.collectStats() {
		slog.Info("round end event skip")
		return
	}

	if err := p.rounds.endCurrent(e); err != nil {
		slog.Error("current round not ended: %w", err)
	}
}

func (p *parser) matchStartedChangedHandler(e events.MatchStartedChanged) {
	// https://github.com/markus-wa/demoinfocs-golang/discussions/366
	// if e.OldIsStarted || !e.NewIsStarted {
	// 	return
	// }
	//
	// p.gameState.started = true
	// team := p.GameState().TeamCounterTerrorists()
	// p.gameState.teamA = newTeam(team.ClanName(), team.Flag(), team.Team(), team.Members())
	// team = p.GameState().TeamTerrorists()
	// p.gameState.teamB = newTeam(team.ClanName(), team.Flag(), team.Team(), team.Members())
}

func (p *parser) matchStartHandler(e events.MatchStart) {
	slog.Info("MATCH START")
	p.gameState.started = true
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

	// collect stats
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
		slog.Error("kill by unconnected player", "event", e)
	}

	if playerConnected(e.Victim) {
		p.playerStats.incr(e.Victim.SteamID64, eventDeath)
		p.weaponStats.incr(e.Victim.SteamID64, eventDeath, e.Weapon.Type)
	} else {
		slog.Error("killed unconnected player", "event", e)
	}

	if playerConnected(e.Assister) {
		p.playerStats.incr(e.Assister.SteamID64, eventAssist)
		p.weaponStats.incr(e.Assister.SteamID64, eventAssist, e.Weapon.Type)

		if e.AssistedFlash {
			p.playerStats.incr(e.Assister.SteamID64, eventFBAssist)
		}
	} else {
		slog.Info("kill assist by unconnected player", "event", e)
	}

	if err := p.rounds.killCount(e); err != nil {
		slog.Error("kill not counted", "err", err.Error(), "kill", e)
	}
}

func (p *parser) hurtHandler(e events.PlayerHurt) {
	if !p.gameState.collectStats() {
		slog.Info("skipped player hurt event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if e.HealthDamage == 0 {
		slog.Error("skipping 0 damage hurt event", "event", e)
		return
	}

	if playerConnected(e.Player) {
		p.playerStats.add(e.Player.SteamID64, eventDmgTaken, e.HealthDamage)
		p.weaponStats.add(e.Player.SteamID64, eventDmgTaken, e.Weapon.Type, e.HealthDamage)
	} else {
		slog.Error("unconnected player got hurt", "event", e)
	}

	if playerConnected(e.Attacker) {
		p.playerStats.add(e.Attacker.SteamID64, eventDmgDealt, e.HealthDamage)
		p.weaponStats.add(e.Attacker.SteamID64, eventDmgDealt, e.Weapon.Type, e.HealthDamage)

		if e.HitGroup != events.HitGroupGeneric && e.HitGroup != events.HitGroupGear {
			p.weaponStats.incr(e.Attacker.SteamID64, p.hitgroupToEvent(e.HitGroup), e.Weapon.Type)
		}

		if e.Weapon.Class() == common.EqClassGrenade {
			p.playerStats.add(e.Attacker.SteamID64, eventDmgGrenadeDealt, e.HealthDamage)
		}
	} else {
		slog.Error("unconnected attacker hurt other player", "event", e)
	}
}

func (p *parser) teamSideSwitchHandler(_ events.TeamSideSwitch) {
	slog.Info("team side switch")

	// if err := p.gameState.teamA.swapSide(); err != nil {
	// 	slog.Error("team A side not swapped", "err", err.Error())
	// }
	//
	// if err := p.gameState.teamB.swapSide(); err != nil {
	// 	slog.Error("team B side not swapped", "err", err.Error())
	// }
}

func (p *parser) weaponFireHandler(e events.WeaponFire) {
	if !p.gameState.collectStats() {
		slog.Info("skipped weapon fire event",
			"knife_round", p.gameState.knifeRound,
			"game_started", p.gameState.started)
		return
	}

	if !playerConnected(e.Shooter) {
		slog.Error("fire from unconnected player", "shooter", e.Shooter)
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
		slog.Error("bomb planted by unconnected player",
			"site", e.Site,
			"player", e.Player)
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
		slog.Error("bomb defused by unconnected player",
			"site", e.Site,
			"player", e.Player)
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
		slog.Info("flashed spectator", "player", e.Player)
		return
	}

	if playerConnected(e.Player) {
		p.playerStats.incr(e.Player.SteamID64, eventBecameBlind)
	} else {
		slog.Error("flashed by unconnected player", "player", e.Player)
	}

	if playerConnected(e.Attacker) {
		p.playerStats.incr(e.Attacker.SteamID64, eventBlindedPlayer)
	} else {
		slog.Error("unconnected player flashed other player", "attacker", e.Attacker)
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
		slog.Error("announced mvp for unconnected player", "player", e.Player)
		return
	}

	p.playerStats.incr(e.Player.SteamID64, eventRoundMVP)
}
