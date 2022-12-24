package replayparser

import (
	"errors"
	"io"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/ssssargsian/uniplay/internal/domain"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
)

// parser is a wrapper around demoinfocs.Parser.
// ONE parser must be used for ONE replay.
type parser struct {
	p demoinfocs.Parser

	log *zap.Logger

	isKnifeRound bool
	stats        stats
	match        match
}

func New(r io.Reader, l *zap.Logger) (*parser, error) {
	if r == nil {
		return nil, errors.New("parser: reader cannot be nil")
	}

	return &parser{
		p:            demoinfocs.NewParser(r),
		log:          l,
		isKnifeRound: false,
		stats:        newStats(),
		match:        match{},
	}, nil
}

// ParseMatchID returns match id generated from demo header and sets it to a local match struct.
// Must be used before demo parse to prevent parsing demo which was parsed before to save a little bit of time.
func (p *parser) ParseMatchID() (uuid.UUID, error) {
	h, err := p.p.ParseHeader()
	if err != nil {
		return uuid.UUID{}, err
	}

	p.match.id, err = domain.NewMatchID(
		h.ServerName,
		h.ClientName,
		h.MapName,
		h.PlaybackTime,
		h.PlaybackTicks,
		h.PlaybackFrames,
		h.SignonLength,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return p.match.id, nil
}

func (p *parser) Close() error {
	if p != nil {
		return p.p.Close()
	}
	return nil
}

// Parse collect player stats from demo.
func (p *parser) Parse() (parseResult, error) {
	if (p.match.id == uuid.UUID{}) {
		return parseResult{}, errors.New("parser: empty match id")
	}

	p.p.RegisterEventHandler(func(_ events.RoundFreezetimeEnd) {
		p.detectKnifeRound()
	})

	// https://github.com/markus-wa/demoinfocs-golang/discussions/366
	p.p.RegisterEventHandler(func(e events.MatchStartedChanged) {
		if e.OldIsStarted || !e.NewIsStarted {
			return
		}

		p.setTeams(p.p.GameState())
	})

	p.p.RegisterEventHandler(func(_ events.TeamSideSwitch) {
		p.match.swapTeamSides()
	})

	p.p.RegisterEventHandler(func(e events.ScoreUpdated) {
		p.match.updateTeamsScore(e)
	})

	p.p.RegisterEventHandler(func(e events.Kill) {
		p.handleKills(e)
	})

	p.p.RegisterEventHandler(func(e events.WeaponFire) {
		if !p.collectStats(p.p.GameState()) || !p.playerConnected(e.Shooter) {
			return
		}

		if e.Weapon == nil {
			return
		}

		p.stats.incrWeaponStat(e.Shooter.SteamID64, domain.MetricShot, e.Weapon.Type)
	})

	p.p.RegisterEventHandler(func(e events.PlayerHurt) {
		p.handlePlayerHurt(e)
	})

	p.p.RegisterEventHandler(func(e events.BombDefused) {
		if !p.collectStats(p.p.GameState()) || !p.playerConnected(e.Player) {
			return
		}

		p.stats.incrPlayerStat(e.Player.SteamID64, domain.MetricBombDefused)
	})

	p.p.RegisterEventHandler(func(e events.BombPlanted) {
		if !p.collectStats(p.p.GameState()) || !p.playerConnected(e.Player) {
			return
		}

		p.stats.incrPlayerStat(e.Player.SteamID64, domain.MetricBombPlanted)
	})

	p.p.RegisterEventHandler(func(e events.PlayerFlashed) {
		if !p.collectStats(p.p.GameState()) || p.playerSpectator(e.Player) {
			return
		}

		if p.playerConnected(e.Player) {
			p.stats.incrPlayerStat(e.Player.SteamID64, domain.MetricBlinded)
		}

		if p.playerConnected(e.Attacker) {
			p.stats.incrPlayerStat(e.Attacker.SteamID64, domain.MetricBlind)
		}
	})

	p.p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if !p.collectStats(p.p.GameState()) || !p.playerConnected(e.Player) {
			return
		}

		if p.playerConnected(e.Player) {
			p.stats.incrPlayerStat(e.Player.SteamID64, domain.MetricRoundMVP)
		}
	})

	p.p.RegisterEventHandler(func(_ events.AnnouncementWinPanelMatch) {
		p.match.mapName = p.p.Header().MapName
		p.match.duration = p.p.Header().PlaybackTime
	})

	if err := p.p.ParseToEnd(); err != nil {
		return parseResult{}, err
	}

	return parseResult{
		match: p.match,
		stats: p.stats,
	}, nil
}

func (p *parser) playerSpectator(player *common.Player) bool {
	return player != nil && (player.Team == common.TeamSpectators || player.Team == common.TeamUnassigned)
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

	for _, player := range p.p.GameState().TeamCounterTerrorists().Members() {
		weapons := player.Weapons()
		if len(weapons) == 1 && weapons[0].Type == common.EqKnife {
			p.isKnifeRound = true
			break
		}
	}
}

// playerConnected checks is player connected and steamID is not equal to 0.
func (p *parser) playerConnected(pl *common.Player) bool {
	return pl != nil && pl.SteamID64 != 0
}

// setTeams sets teams clan names, flags, sides (ct/t) and their members to p.match.
func (p *parser) setTeams(gs demoinfocs.GameState) {
	t := gs.TeamTerrorists()
	p.match.team1 = newMatchTeam(t.ClanName(), t.Flag(), t.Team(), t.Members())

	ct := gs.TeamCounterTerrorists()
	p.match.team2 = newMatchTeam(ct.ClanName(), ct.Flag(), ct.Team(), ct.Members())
}

func (p *parser) hitgroupToMetric(g events.HitGroup) domain.Metric {
	switch g {
	case events.HitGroupHead:
		return domain.MetricHitHead
	case events.HitGroupChest:
		return domain.MetricHitChest
	case events.HitGroupStomach:
		return domain.MetricHitStomach
	case events.HitGroupLeftArm:
		return domain.MetricHitLeftArm
	case events.HitGroupRightArm:
		return domain.MetricHitRightArm
	case events.HitGroupLeftLeg:
		return domain.MetricHitLeftLeg
	case events.HitGroupRightLeg:
		return domain.MetricHitRightLeg
	}
	return 0
}

// handleKills collects metrics and weapon metrics on kill event.
func (p *parser) handleKills(e events.Kill) {
	if !p.collectStats(p.p.GameState()) {
		return
	}

	if p.playerConnected(e.Victim) {
		// death amount FROM weapon
		p.stats.incrPlayerStat(e.Victim.SteamID64, domain.MetricDeath)

		if e.Weapon != nil {
			p.stats.incrWeaponStat(e.Victim.SteamID64, domain.MetricDeath, e.Weapon.Type)
		}
	}

	if p.playerConnected(e.Killer) {
		// kill amount
		p.stats.incrPlayerStat(e.Killer.SteamID64, domain.MetricKill)

		if e.Weapon != nil {
			p.stats.incrWeaponStat(e.Killer.SteamID64, domain.MetricKill, e.Weapon.Type)
		}

		// headshot kill amount
		if e.IsHeadshot {
			p.stats.incrPlayerStat(e.Killer.SteamID64, domain.MetricHSKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, domain.MetricHSKill, e.Weapon.Type)
			}
		}

		// blind kill amount
		if e.AttackerBlind {
			p.stats.incrPlayerStat(e.Killer.SteamID64, domain.MetricBlindKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, domain.MetricBlindKill, e.Weapon.Type)
			}
		}

		// wallbang kill amount
		if e.IsWallBang() {
			p.stats.incrPlayerStat(e.Killer.SteamID64, domain.MetricWallbangKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, domain.MetricWallbangKill, e.Weapon.Type)
			}
		}

		// noscope kill amount
		if e.NoScope {
			p.stats.incrPlayerStat(e.Killer.SteamID64, domain.MetricNoScopeKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, domain.MetricNoScopeKill, e.Weapon.Type)
			}
		}

		// through smoke kill amount
		if e.ThroughSmoke {
			p.stats.incrPlayerStat(e.Killer.SteamID64, domain.MetricThroughSmokeKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, domain.MetricThroughSmokeKill, e.Weapon.Type)
			}
		}
	}

	if p.playerConnected(e.Assister) {
		// assist total amount
		p.stats.incrPlayerStat(e.Assister.SteamID64, domain.MetricAssist)

		// flashbang assist amount
		if e.AssistedFlash {
			p.stats.incrPlayerStat(e.Assister.SteamID64, domain.MetricFlashbangAssist)
		}

		if e.Weapon != nil {
			// assist with weapon
			p.stats.incrWeaponStat(e.Assister.SteamID64, domain.MetricAssist, e.Weapon.Type)
		}
	}
}

// handlePlayerHurt collects metrics and weapon metrics on player hurt.
func (p *parser) handlePlayerHurt(e events.PlayerHurt) {
	if !p.collectStats(p.p.GameState()) {
		return
	}

	if p.playerConnected(e.Attacker) {
		// dealt damage
		p.stats.addPlayerStat(e.Attacker.SteamID64, domain.MetricDamageDealt, e.HealthDamage)

		if e.Weapon != nil {
			// dealth damage with weapon
			p.stats.addWeaponStat(e.Attacker.SteamID64, domain.MetricDamageDealt, e.Weapon.Type, e.HealthDamage)

			// hitgroup shot
			metric := p.hitgroupToMetric(e.HitGroup)
			p.stats.incrWeaponStat(e.Attacker.SteamID64, metric, e.Weapon.Type)
		}
	}

	if p.playerConnected(e.Player) {
		// taken damage
		p.stats.addPlayerStat(e.Player.SteamID64, domain.MetricDamageTaken, e.HealthDamage)

		if e.Weapon != nil {
			// taken damage from weapon
			p.stats.addWeaponStat(e.Player.SteamID64, domain.MetricDamageTaken, e.Weapon.Type, e.HealthDamage)
		}
	}
}
