package replay

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"

	"github.com/ssssargsian/uniplay/internal/domain"
)

// parser is a wrapper around demoinfocs.Parser.
// ONE parser must be used for ONE replay.
type parser struct {
	p demoinfocs.Parser

	log *zap.Logger

	isKnifeRound bool
	stats        stats
	match        *replayMatch
}

func newParser(r replay, l *zap.Logger) (*parser, error) {
	if (r == replay{}) {
		return nil, errors.New("parser: empty replay")
	}

	if l == nil {
		return nil, errors.New("parser: logger cannot be nil")
	}

	return &parser{
		p:            demoinfocs.NewParser(r),
		log:          l,
		isKnifeRound: false,
		stats:        newStats(),
		match:        new(replayMatch),
	}, nil
}

// parseReplayHeader parses replay header and generates match id from it.
// Must be called before collectStats().
func (p *parser) parseReplayHeader() (uuid.UUID, error) {
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

	p.match.uploadedAt = time.Now()
	return p.match.id, nil
}

// collectStats collects player stats from the replay.
func (p *parser) collectStats() (*replayMatch, []*playerStat, []*weaponStat, error) {
	if (p.match.id == uuid.UUID{}) {
		return nil, nil, nil, errors.New("parser: empty match id, call parseReplayHeader() first")
	}

	p.p.RegisterEventHandler(func(_ events.RoundFreezetimeEnd) {
		p.handleKnifeRound()
	})

	// https://github.com/markus-wa/demoinfocs-golang/discussions/366
	p.p.RegisterEventHandler(func(e events.MatchStartedChanged) {
		if e.OldIsStarted || !e.NewIsStarted {
			return
		}

		p.match.setTeams(p.p.GameState())
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
		if !p.matchStarted(p.p.GameState()) || !p.playerConnected(e.Shooter) {
			return
		}

		if e.Weapon == nil {
			return
		}

		p.stats.incrWeaponStat(e.Shooter.SteamID64, metricShot, e.Weapon.Type)
	})

	p.p.RegisterEventHandler(func(e events.PlayerHurt) {
		p.handlePlayerHurt(e)
	})

	p.p.RegisterEventHandler(func(e events.BombDefused) {
		if !p.matchStarted(p.p.GameState()) || !p.playerConnected(e.Player) {
			return
		}

		p.stats.incrPlayerStat(e.Player.SteamID64, metricBombDefused)
	})

	p.p.RegisterEventHandler(func(e events.BombPlanted) {
		if !p.matchStarted(p.p.GameState()) || !p.playerConnected(e.Player) {
			return
		}

		p.stats.incrPlayerStat(e.Player.SteamID64, metricBombPlanted)
	})

	p.p.RegisterEventHandler(func(e events.PlayerFlashed) {
		if !p.matchStarted(p.p.GameState()) || p.playerSpectator(e.Player) {
			return
		}

		if p.playerConnected(e.Player) {
			p.stats.incrPlayerStat(e.Player.SteamID64, metricBlinded)
		}

		if p.playerConnected(e.Attacker) {
			p.stats.incrPlayerStat(e.Attacker.SteamID64, metricBlind)
		}
	})

	p.p.RegisterEventHandler(func(e events.RoundMVPAnnouncement) {
		if !p.matchStarted(p.p.GameState()) || !p.playerConnected(e.Player) {
			return
		}

		if p.playerConnected(e.Player) {
			p.stats.incrPlayerStat(e.Player.SteamID64, metricRoundMVP)
		}
	})

	p.p.RegisterEventHandler(func(_ events.AnnouncementWinPanelMatch) {
		p.match.mapName = p.p.Header().MapName
		p.match.duration = p.p.Header().PlaybackTime
	})

	if err := p.p.ParseToEnd(); err != nil {
		return nil, nil, nil, err
	}

	playerStats, weaponStats := p.stats.normalizeSync()
	p.match.setTeamStates()

	return p.match, playerStats, weaponStats, nil
}

func (p *parser) close() error {
	if p == nil {
		return nil
	}
	return p.p.Close()
}

// playerSpectator checks whether player spectator or not.
func (p *parser) playerSpectator(player *common.Player) bool {
	if player == nil {
		return true
	}

	return player.Team == common.TeamSpectators || player.Team == common.TeamUnassigned
}

// matchStarted detects if stats can be collected to prevent collection of stats on knife or warmup rounds.
// return false if current round is knife round or match is not started.
func (p *parser) matchStarted(gs demoinfocs.GameState) bool {
	if p.isKnifeRound || !gs.IsMatchStarted() {
		return false
	}

	return true
}

// handleKnifeRound sets isKnifeRound to true if any player has no secondary weapon and first slot is a knife.
func (p *parser) handleKnifeRound() {
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
func (p *parser) playerConnected(player *common.Player) bool {
	return player != nil && player.SteamID64 != 0 && player.IsConnected && !player.IsBot && !player.IsUnknown
}

// hitgroupToMetric returns metric associated with the hitgroup.
func (p *parser) hitgroupToMetric(g events.HitGroup) metric {
	switch g {
	case events.HitGroupHead:
		return metricHitHead
	case events.HitGroupChest:
		return metricHitChest
	case events.HitGroupStomach:
		return metricHitStomach
	case events.HitGroupLeftArm:
		return metricHitLeftArm
	case events.HitGroupRightArm:
		return metricHitRightArm
	case events.HitGroupLeftLeg:
		return metricHitLeftLeg
	case events.HitGroupRightLeg:
		return metricHitRightLeg
	}
	return 0
}

// handleKills collects player and weapon stats on kill event.
func (p *parser) handleKills(e events.Kill) {
	if !p.matchStarted(p.p.GameState()) {
		return
	}

	if p.playerConnected(e.Victim) {
		// death amount FROM weapon
		p.stats.incrPlayerStat(e.Victim.SteamID64, metricDeath)

		if e.Weapon != nil {
			p.stats.incrWeaponStat(e.Victim.SteamID64, metricDeath, e.Weapon.Type)
		}
	}

	if p.playerConnected(e.Killer) {
		// kill amount
		p.stats.incrPlayerStat(e.Killer.SteamID64, metricKill)

		if e.Weapon != nil {
			p.stats.incrWeaponStat(e.Killer.SteamID64, metricKill, e.Weapon.Type)
		}

		// headshot kill amount
		if e.IsHeadshot {
			p.stats.incrPlayerStat(e.Killer.SteamID64, metricHSKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, metricHSKill, e.Weapon.Type)
			}
		}

		// blind kill amount
		if e.AttackerBlind {
			p.stats.incrPlayerStat(e.Killer.SteamID64, metricBlindKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, metricBlindKill, e.Weapon.Type)
			}
		}

		// wallbang kill amount
		if e.IsWallBang() {
			p.stats.incrPlayerStat(e.Killer.SteamID64, metricWallbangKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, metricWallbangKill, e.Weapon.Type)
			}
		}

		// noscope kill amount
		if e.NoScope {
			p.stats.incrPlayerStat(e.Killer.SteamID64, metricNoScopeKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, metricNoScopeKill, e.Weapon.Type)
			}
		}

		// through smoke kill amount
		if e.ThroughSmoke {
			p.stats.incrPlayerStat(e.Killer.SteamID64, metricThroughSmokeKill)

			if e.Weapon != nil {
				p.stats.incrWeaponStat(e.Killer.SteamID64, metricThroughSmokeKill, e.Weapon.Type)
			}
		}
	}

	if p.playerConnected(e.Assister) {
		// assist total amount
		p.stats.incrPlayerStat(e.Assister.SteamID64, metricAssist)

		// flashbang assist amount
		if e.AssistedFlash {
			p.stats.incrPlayerStat(e.Assister.SteamID64, metricFlashbangAssist)
		}

		if e.Weapon != nil {
			// assist with weapon
			p.stats.incrWeaponStat(e.Assister.SteamID64, metricAssist, e.Weapon.Type)
		}
	}
}

// handlePlayerHurt collects player and weapon stats on player hurt.
func (p *parser) handlePlayerHurt(e events.PlayerHurt) {
	if !p.matchStarted(p.p.GameState()) {
		return
	}

	if p.playerConnected(e.Attacker) {
		// dealt damage
		p.stats.addPlayerStat(e.Attacker.SteamID64, metricDamageDealt, e.HealthDamage)

		if e.Weapon != nil {
			// dealth damage with weapon
			p.stats.addWeaponStat(e.Attacker.SteamID64, metricDamageDealt, e.Weapon.Type, e.HealthDamage)

			// hitgroup shot
			p.stats.incrWeaponStat(e.Attacker.SteamID64, p.hitgroupToMetric(e.HitGroup), e.Weapon.Type)

			// collect grenade damage
			if e.Weapon.Class() == common.EqClassGrenade {
				p.stats.addPlayerStat(e.Attacker.SteamID64, metricGrenadeDamageDealt, e.HealthDamage)
			}
		}
	}

	if p.playerConnected(e.Player) {
		// taken damage
		p.stats.addPlayerStat(e.Player.SteamID64, metricDamageTaken, e.HealthDamage)

		if e.Weapon != nil {
			// taken damage from weapon
			p.stats.addWeaponStat(e.Player.SteamID64, metricDamageTaken, e.Weapon.Type, e.HealthDamage)
		}
	}
}
