package replay

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/otel"
)

var (
	errEmptyMatchID = errors.New("parser: empty match id, call parseReplayHeader() first")
	errEmptyLogger  = errors.New("parser: logger cannot be nil")
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
	if l == nil {
		return nil, errEmptyLogger
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
func (p *parser) collectStats(ctx context.Context) (*replayMatch, []*playerStat, []*weaponStat, error) {
	_, span := otel.StartTrace(ctx, libraryName, "replay.parser.collectStats")
	defer span.End()

	span.SetAttributes(attribute.String("match_id", p.match.id.String()))

	if (p.match.id == uuid.UUID{}) {
		return nil, nil, nil, errEmptyMatchID
	}

	p.p.RegisterEventHandler(p.roundFreezetimeEndHandler)
	p.p.RegisterEventHandler(p.matchStartedChangedHandler)
	p.p.RegisterEventHandler(p.teamSideSwitchHandler)
	p.p.RegisterEventHandler(p.scoreUpdatedHandler)
	p.p.RegisterEventHandler(p.killHandler)
	p.p.RegisterEventHandler(p.weaponFireHandler)
	p.p.RegisterEventHandler(p.playerHurtHandler)
	p.p.RegisterEventHandler(p.bombDefusedHandler)
	p.p.RegisterEventHandler(p.bombPlantedHandler)
	p.p.RegisterEventHandler(p.playerFlashedHandler)
	p.p.RegisterEventHandler(p.roundMVPAnnouncementHandler)
	p.p.RegisterEventHandler(p.announcementWinPanelMatchHandler)

	if err := p.p.ParseToEnd(); err != nil {
		return nil, nil, nil, err
	}

	p.match.setTeamStates()
	playerStats, weaponStats := p.stats.normalizeSync()

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

// collectAllowed detects if stats can be collected to prevent collection of stats on knife or warmup rounds.
// return false if current round is knife round or match is not started.
func (p *parser) collectAllowed(matchStarted bool) bool {
	if p.isKnifeRound || !matchStarted {
		return false
	}

	return true
}

// playerValid checks is player connected, steam id not equal to 0 and player is not a bot or unknown.
func (p *parser) playerValid(player *common.Player) bool {
	return player != nil && player.SteamID64 != 0 && player.IsConnected && !player.IsBot && !player.IsUnknown
}

// detectKnifeRound detects is current round is a knife round,
// sets isKnifeRound to true if any player has no secondary weapon and first slot is a knife.
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

// hitgroupToMetric returns metric associated with the hitgroup.
func (p *parser) hitgroupToMetric(g events.HitGroup) metric {
	switch g {
	case events.HitGroupHead:
		return metricHitHead
	case events.HitGroupNeck:
		return metricHitNeck
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
	case events.HitGroupGeneric, events.HitGroupGear:
		return 0
	}

	return 0
}

func (p *parser) killHandler(e events.Kill) {
	if !p.collectAllowed(p.p.GameState().IsMatchStarted()) {
		return
	}

	p.handleKillerKill(e)

	if p.playerValid(e.Victim) {
		p.stats.incrPlayerStat(e.Victim.SteamID64, metricDeath)
		p.stats.incrWeaponStat(e.Victim.SteamID64, metricDeath, e.Weapon.Type)
	}

	if p.playerValid(e.Assister) {
		p.stats.incrPlayerStat(e.Assister.SteamID64, metricAssist)
		p.stats.incrWeaponStat(e.Assister.SteamID64, metricAssist, e.Weapon.Type)

		if e.AssistedFlash {
			p.stats.incrPlayerStat(e.Assister.SteamID64, metricFlashbangAssist)
		}
	}
}

func (p *parser) handleKillerKill(e events.Kill) {
	if !p.playerValid(e.Killer) {
		return
	}

	p.stats.incrPlayerStat(e.Killer.SteamID64, metricKill)
	p.stats.incrWeaponStat(e.Killer.SteamID64, metricKill, e.Weapon.Type)

	switch {
	case e.IsHeadshot:
		p.stats.incrPlayerStat(e.Killer.SteamID64, metricHSKill)
		p.stats.incrWeaponStat(e.Killer.SteamID64, metricHSKill, e.Weapon.Type)
	case e.AttackerBlind:
		p.stats.incrPlayerStat(e.Killer.SteamID64, metricBlindKill)
		p.stats.incrWeaponStat(e.Killer.SteamID64, metricBlindKill, e.Weapon.Type)
	case e.IsWallBang():
		p.stats.incrPlayerStat(e.Killer.SteamID64, metricWallbangKill)
		p.stats.incrWeaponStat(e.Killer.SteamID64, metricWallbangKill, e.Weapon.Type)
	case e.NoScope:
		p.stats.incrPlayerStat(e.Killer.SteamID64, metricNoScopeKill)
		p.stats.incrWeaponStat(e.Killer.SteamID64, metricNoScopeKill, e.Weapon.Type)
	case e.ThroughSmoke:
		p.stats.incrPlayerStat(e.Killer.SteamID64, metricThroughSmokeKill)
		p.stats.incrWeaponStat(e.Killer.SteamID64, metricThroughSmokeKill, e.Weapon.Type)
	}
}

func (p *parser) playerHurtHandler(e events.PlayerHurt) { //nolint:gocritic // demoinfocs
	if !p.collectAllowed(p.p.GameState().IsMatchStarted()) {
		return
	}

	if p.playerValid(e.Attacker) {
		p.stats.addPlayerStat(e.Attacker.SteamID64, metricDamageDealt, e.HealthDamage)
		p.stats.addWeaponStat(e.Attacker.SteamID64, metricDamageDealt, e.Weapon.Type, e.HealthDamage)
		p.stats.incrWeaponStat(e.Attacker.SteamID64, p.hitgroupToMetric(e.HitGroup), e.Weapon.Type)

		if e.Weapon.Class() == common.EqClassGrenade {
			p.stats.addPlayerStat(e.Attacker.SteamID64, metricGrenadeDamageDealt, e.HealthDamage)
		}
	}

	if p.playerValid(e.Player) {
		p.stats.addPlayerStat(e.Player.SteamID64, metricDamageTaken, e.HealthDamage)
		p.stats.addWeaponStat(e.Player.SteamID64, metricDamageTaken, e.Weapon.Type, e.HealthDamage)
	}
}

func (p *parser) roundFreezetimeEndHandler(_ events.RoundFreezetimeEnd) {
	p.detectKnifeRound()
}

func (p *parser) matchStartedChangedHandler(e events.MatchStartedChanged) {
	// https://github.com/markus-wa/demoinfocs-golang/discussions/366
	if e.OldIsStarted || !e.NewIsStarted {
		return
	}

	p.match.setTeams(p.p.GameState())
}

func (p *parser) teamSideSwitchHandler(_ events.TeamSideSwitch) {
	p.match.swapTeamSides()
}

func (p *parser) scoreUpdatedHandler(e events.ScoreUpdated) {
	p.match.updateTeamsScore(e)
}

func (p *parser) weaponFireHandler(e events.WeaponFire) {
	if e.Weapon == nil {
		return
	}

	if !p.collectAllowed(p.p.GameState().IsMatchStarted()) || !p.playerValid(e.Shooter) {
		return
	}

	p.stats.incrWeaponStat(e.Shooter.SteamID64, metricShot, e.Weapon.Type)
}

func (p *parser) bombDefusedHandler(e events.BombDefused) {
	if !p.collectAllowed(p.p.GameState().IsMatchStarted()) || !p.playerValid(e.Player) {
		return
	}

	p.stats.incrPlayerStat(e.Player.SteamID64, metricBombDefused)
}

func (p *parser) bombPlantedHandler(e events.BombPlanted) {
	if !p.collectAllowed(p.p.GameState().IsMatchStarted()) || !p.playerValid(e.Player) {
		return
	}

	p.stats.incrPlayerStat(e.Player.SteamID64, metricBombPlanted)
}

func (p *parser) playerFlashedHandler(e events.PlayerFlashed) {
	// spectator can be flashed
	if !p.collectAllowed(p.p.GameState().IsMatchStarted()) || p.playerSpectator(e.Player) {
		return
	}

	if p.playerValid(e.Player) {
		p.stats.incrPlayerStat(e.Player.SteamID64, metricBlinded)
	}

	if p.playerValid(e.Attacker) {
		p.stats.incrPlayerStat(e.Attacker.SteamID64, metricBlind)
	}
}

func (p *parser) roundMVPAnnouncementHandler(e events.RoundMVPAnnouncement) {
	if !p.collectAllowed(p.p.GameState().IsMatchStarted()) || !p.playerValid(e.Player) {
		return
	}

	p.stats.incrPlayerStat(e.Player.SteamID64, metricRoundMVP)
}

func (p *parser) announcementWinPanelMatchHandler(_ events.AnnouncementWinPanelMatch) {
	p.match.mapName = p.p.Header().MapName
	p.match.duration = p.p.Header().PlaybackTime
}
