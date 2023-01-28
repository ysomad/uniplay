package replay

import (
	"testing"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
	"github.com/stretchr/testify/assert"
)

func Test_parser_playerSpectator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		player *common.Player
		want   bool
	}{
		{
			name:   "nil player",
			player: nil,
			want:   true,
		},
		{
			name: "player spectator",
			player: &common.Player{
				Team: common.TeamSpectators,
			},
			want: true,
		},
		{
			name: "player not spectator",
			player: &common.Player{
				Team: common.TeamTerrorists,
			},
			want: false,
		},
		{
			name: "player not spectator 2",
			player: &common.Player{
				Team: common.TeamCounterTerrorists,
			},
			want: false,
		},
		{
			name: "player team not assigned",
			player: &common.Player{
				Team: common.TeamUnassigned,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(parser)
			got := p.playerSpectator(tt.player)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parser_playerConnected(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		player *common.Player
		want   bool
	}{
		{
			name:   "nil player",
			player: nil,
			want:   false,
		},
		{
			name: "invalid steam id",
			player: &common.Player{
				SteamID64:   0,
				IsBot:       false,
				IsConnected: true,
				IsUnknown:   false,
			},
			want: false,
		},
		{
			name: "player not connected",
			player: &common.Player{
				SteamID64:   1,
				IsBot:       false,
				IsConnected: false,
				IsUnknown:   false,
			},
			want: false,
		},
		{
			name: "player is bot",
			player: &common.Player{
				SteamID64:   1,
				IsBot:       true,
				IsConnected: true,
				IsUnknown:   false,
			},
			want: false,
		},
		{
			name: "player is unknown",
			player: &common.Player{
				SteamID64:   1,
				IsBot:       false,
				IsConnected: true,
				IsUnknown:   true,
			},
			want: false,
		},
		{
			name: "player connected",
			player: &common.Player{
				SteamID64:   1,
				IsBot:       false,
				IsConnected: true,
				IsUnknown:   false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(parser)
			got := p.playerValid(tt.player)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parser_hitgroupToMetric(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hitgroup events.HitGroup
		want     metric
	}{
		{
			name:     "generic hitgroup",
			hitgroup: events.HitGroupGeneric,
			want:     0,
		},
		{
			name:     "gear hitgroup",
			hitgroup: events.HitGroupGear,
			want:     0,
		},
		{
			name:     "head hitgroup",
			hitgroup: events.HitGroupHead,
			want:     metricHitHead,
		},
		{
			name:     "chest hitgroup",
			hitgroup: events.HitGroupChest,
			want:     metricHitChest,
		},
		{
			name:     "stomach hitgroup",
			hitgroup: events.HitGroupStomach,
			want:     metricHitStomach,
		},
		{
			name:     "left arm hitgroup",
			hitgroup: events.HitGroupLeftArm,
			want:     metricHitLeftArm,
		},
		{
			name:     "right arm hitgroup",
			hitgroup: events.HitGroupRightArm,
			want:     metricHitRightArm,
		},
		{
			name:     "left leg hitgroup",
			hitgroup: events.HitGroupLeftLeg,
			want:     metricHitLeftLeg,
		},
		{
			name:     "right leg hitgroup",
			hitgroup: events.HitGroupRightLeg,
			want:     metricHitRightLeg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(parser)
			got := p.hitgroupToMetric(tt.hitgroup)

			assert.Equal(t, tt.want, got)
		})
	}
}
