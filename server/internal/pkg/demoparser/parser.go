package demoparser

import (
	"fmt"
	"time"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type parser struct {
	demoinfocs.Parser
	playerStats playerStatsMap
	weaponStats weaponStatsMap
	demosize    int64
}

func New(d Demo) *parser {
	return &parser{
		Parser:      demoinfocs.NewParser(d),
		demosize:    d.size,
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

	// demoid := newDemoID(&dh)

	return nil
}

func (p *parser) playerConnected(pl *common.Player) bool {
	if pl == nil || pl.SteamID64 == 0 || pl.UserID == 0 || !pl.IsConnected || pl.IsBot || pl.IsUnknown {
		return false
	}

	if pl.Team == common.TeamSpectators || pl.Team == common.TeamUnassigned {
		return false
	}

	return true
}

func (p *parser) hurtHandler(e events.PlayerHurt) {
	if p.playerConnected(e.Player) {
		p.playerStats.add(e.Player.SteamID64, eventDmgTaken, e.HealthDamage)
		p.weaponStats.add(e.Player.SteamID64, eventDmgTaken, e.Weapon.Type, e.HealthDamage)
	}

	if p.playerConnected(e.Attacker) {
		p.playerStats.add(e.Attacker.SteamID64, eventDmgDealt, e.HealthDamage)
		p.weaponStats.add(e.Attacker.SteamID64, eventDmgDealt, e.Weapon.Type, e.HealthDamage)
		p.weaponStats.incr(e.Attacker.SteamID64, p.hitgroupToEvent(e.HitGroup), e.Weapon.Type)

		if e.Weapon.Class() == common.EqClassGrenade {
			p.playerStats.add(e.Attacker.SteamID64, eventDmgGrenadeDealt, e.HealthDamage)
		}
	}
}
