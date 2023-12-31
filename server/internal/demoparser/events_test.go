package demoparser

import (
	"testing"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	"github.com/stretchr/testify/assert"
)

func TestParser_hitgroupToEvent(t *testing.T) {
	type fields struct {
		Parser      demoinfocs.Parser
		playerStats playerStatsMap
		weaponStats weaponStatsMap
		flashbangs  playerFlashbangs
		gameState   *gameState
		rounds      roundHistory
		demosize    int64
		demoID      string
	}
	type args struct {
		hg events.HitGroup
	}
	tests := map[string]struct {
		fields fields
		args   args
		want   event
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := &Parser{
				Parser:      tt.fields.Parser,
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
				flashbangs:  tt.fields.flashbangs,
				gameState:   tt.fields.gameState,
				rounds:      tt.fields.rounds,
				demosize:    tt.fields.demosize,
				demoID:      tt.fields.demoID,
			}
			assert.Equal(t, tt.want, p.hitgroupToEvent(tt.args.hg))
		})
	}
}
