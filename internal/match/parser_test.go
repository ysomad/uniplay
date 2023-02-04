package match

import (
	"io"
	"strings"
	"testing"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
	"github.com/stretchr/testify/assert"
)

func Test_newParser(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		r    replay
		want *parser
	}{
		{
			name: "success",
			r:    replay{io.NopCloser(strings.NewReader("TEST")), 555},
			want: &parser{
				p:          demoinfocs.NewParser(io.NopCloser(strings.NewReader("TEST"))),
				knifeRound: false,
				stats: stats{
					playerStats: map[uint64]*playerStat{},
					weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{},
				},
				match:          new(replayMatch),
				replayFilesize: 555,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newParser(tt.r)

			assert.ObjectsAreEqual(tt.want, got)
		})
	}
}

func Test_parser_collectAllowed(t *testing.T) {
	t.Parallel()

	type args struct {
		matchStarted bool
	}

	tests := []struct {
		name   string
		args   args
		parser *parser
		want   bool
	}{
		{
			name:   "knife round, match not started",
			args:   args{matchStarted: false},
			parser: &parser{knifeRound: true},
			want:   false,
		},
		{
			name:   "not knife round, match not started",
			args:   args{matchStarted: false},
			parser: &parser{knifeRound: false},
			want:   false,
		},
		{
			name:   "knife round, match started",
			args:   args{matchStarted: true},
			parser: &parser{knifeRound: true},
			want:   false,
		},
		{
			name:   "not knife round, match started",
			args:   args{matchStarted: true},
			parser: &parser{knifeRound: false},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.parser.collectAllowed(tt.args.matchStarted)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parser_detectKnifeRound(t *testing.T) {
	t.Parallel()

	type args struct {
		players []*common.Player
	}

	tests := []struct {
		name   string
		args   args
		parser *parser
		want   *parser
	}{
		{
			name:   "empty list of players",
			args:   args{players: []*common.Player{}},
			parser: &parser{knifeRound: false},
			want:   &parser{knifeRound: false},
		},
		{
			name:   "empty list of players",
			args:   args{players: []*common.Player{}},
			parser: &parser{knifeRound: true},
			want:   &parser{knifeRound: false},
		},
		{
			name: "couple players with couple weapons",
			args: args{players: []*common.Player{
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqAK47},
						1: {Type: common.EqKnife},
						2: {Type: common.EqKevlar},
						3: {Type: common.EqUSP},
					},
				},
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqAK47},
						1: {Type: common.EqKnife},
						2: {Type: common.EqKevlar},
						3: {Type: common.EqUSP},
					},
				},
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqAK47},
						1: {Type: common.EqKnife},
						2: {Type: common.EqKevlar},
						3: {Type: common.EqUSP},
					},
				},
			}},
			parser: &parser{knifeRound: true},
			want:   &parser{knifeRound: false},
		},
		{
			name: "couple players all with knifes only",
			args: args{players: []*common.Player{
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqKnife},
					},
				},
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqKnife},
					},
				},
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqKnife},
					},
				},
			}},
			parser: &parser{knifeRound: false},
			want:   &parser{knifeRound: true},
		},
		{
			name: "couple players all with knifes only (with true knife round by default)",
			args: args{players: []*common.Player{
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqKnife},
					},
				},
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqKnife},
					},
				},
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqKnife},
					},
				},
			}},
			parser: &parser{knifeRound: false},
			want:   &parser{knifeRound: true},
		},
		{
			name: "one player with knife only",
			args: args{players: []*common.Player{
				{
					Inventory: map[int]*common.Equipment{
						0: {Type: common.EqKnife},
					},
				},
			}},
			parser: &parser{knifeRound: true},
			want:   &parser{knifeRound: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.parser.detectKnifeRound(tt.args.players)

			assert.Equal(t, tt.want, tt.parser)
		})
	}
}

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

func Test_parser_playerValid(t *testing.T) {
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
		{
			name:     "neck hitgroup",
			hitgroup: events.HitGroupNeck,
			want:     metricHitNeck,
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
