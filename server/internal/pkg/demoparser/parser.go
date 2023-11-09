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
	demosize int64
}

func New(d Demo) *parser {
	return &parser{
		Parser:   demoinfocs.NewParser(d),
		demosize: d.size,
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
	}
}
