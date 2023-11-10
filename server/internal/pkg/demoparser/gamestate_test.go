package demoparser

import (
	"testing"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/stretchr/testify/assert"
)

func Test_gameState_detectKnifeRound(t *testing.T) {
	t.Parallel()
	type args struct {
		players []*common.Player
	}
	tests := []struct {
		gameState *gameState
		want      *gameState
		name      string
		args      args
	}{
		{
			name:      "empty list of players",
			args:      args{players: []*common.Player{}},
			gameState: &gameState{knifeRound: false},
			want:      &gameState{knifeRound: false},
		},
		{
			name:      "empty list of players",
			args:      args{players: []*common.Player{}},
			gameState: &gameState{knifeRound: true},
			want:      &gameState{knifeRound: false},
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
			gameState: &gameState{knifeRound: true},
			want:      &gameState{knifeRound: false},
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
			gameState: &gameState{knifeRound: false},
			want:      &gameState{knifeRound: true},
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
			gameState: &gameState{knifeRound: false},
			want:      &gameState{knifeRound: true},
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
			gameState: &gameState{knifeRound: true},
			want:      &gameState{knifeRound: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.gameState.detectKnifeRound(tt.args.players)
			assert.Equal(t, tt.want, tt.gameState)
		})
	}
}

func Test_gameState_collectStats(t *testing.T) {
	type fields struct {
		rounds     []round
		teamA      team
		teamB      team
		knifeRound bool
		started    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Regular Round, Game Started",
			fields: fields{
				knifeRound: false,
				started:    true,
			},
			want: true,
		},
		{
			name: "Knife Round",
			fields: fields{
				knifeRound: true,
				started:    true,
			},
			want: false,
		},
		{
			name: "Game Not Started",
			fields: fields{
				knifeRound: false,
				started:    false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &gameState{
				rounds:     tt.fields.rounds,
				teamA:      tt.fields.teamA,
				teamB:      tt.fields.teamB,
				knifeRound: tt.fields.knifeRound,
				started:    tt.fields.started,
			}
			got := gs.collectStats()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newTeam(t *testing.T) {
	type args struct {
		name string
		flag string
		pp   []*common.Player
		side common.Team
	}
	tests := []struct {
		name string
		args args
		want team
	}{
		{
			name: "Valid Team",
			args: args{
				name: "TeamA",
				flag: "FLAG_A",
				side: common.TeamTerrorists,
				pp: []*common.Player{
					{
						SteamID64:   123456789,
						UserID:      1,
						IsConnected: true,
						IsBot:       false,
						IsUnknown:   false,
					},
					{
						SteamID64:   987654321,
						UserID:      2,
						IsConnected: true,
						IsBot:       false,
						IsUnknown:   false,
					},
				},
			},
			want: team{
				name: "TeamA",
				flag: "FLAG_A",
				players: []uint64{
					123456789,
					987654321,
				},
				side:   common.TeamTerrorists,
				status: gameStatusUnknown,
			},
		},
		{
			name: "Team with Disconnected Player",
			args: args{
				name: "TeamB",
				flag: "FLAG_B",
				side: common.TeamCounterTerrorists,
				pp: []*common.Player{
					{
						SteamID64:   111111111,
						UserID:      3,
						IsConnected: false,
						IsBot:       false,
						IsUnknown:   false,
					},
				},
			},
			want: team{
				name:    "TeamB",
				flag:    "FLAG_B",
				players: []uint64{},
				side:    common.TeamCounterTerrorists,
				status:  gameStatusUnknown,
			},
		},
		{
			name: "Team with Bot Player",
			args: args{
				name: "TeamC",
				flag: "FLAG_C",
				side: common.TeamTerrorists,
				pp: []*common.Player{
					{
						SteamID64:   222222222,
						UserID:      4,
						IsConnected: true,
						IsBot:       true,
						IsUnknown:   false,
					},
				},
			},
			want: team{
				name:    "TeamC",
				flag:    "FLAG_C",
				players: []uint64{},
				side:    common.TeamTerrorists,
				status:  gameStatusUnknown,
			},
		},
		{
			name: "Team with Unknown Player",
			args: args{
				name: "TeamD",
				flag: "FLAG_D",
				side: common.TeamCounterTerrorists,
				pp: []*common.Player{
					{
						SteamID64:   333333333,
						UserID:      5,
						IsConnected: true,
						IsBot:       false,
						IsUnknown:   true,
					},
				},
			},
			want: team{
				name:    "TeamD",
				flag:    "FLAG_D",
				players: []uint64{},
				side:    common.TeamCounterTerrorists,
				status:  gameStatusUnknown,
			},
		},
		{
			name: "Nil Team",
			args: args{
				name: "TeamE",
				flag: "FLAG_E",
				side: common.TeamTerrorists,
				pp:   nil,
			},
			want: team{
				name:    "TeamE",
				flag:    "FLAG_E",
				players: []uint64{},
				side:    common.TeamTerrorists,
				status:  gameStatusUnknown,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newTeam(tt.args.name, tt.args.flag, tt.args.side, tt.args.pp)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_team_swapSide(t *testing.T) {
	type fields struct {
		name    string
		flag    string
		players []uint64
		score   int
		side    common.Team
		status  gameStatus
	}
	tests := []struct {
		name     string
		fields   fields
		wantSide common.Team
		wantErr  bool
	}{
		{
			name: "Swap from CT to T",
			fields: fields{
				name:    "TeamA",
				flag:    "FLAG_A",
				players: []uint64{1, 2, 3},
				score:   5,
				side:    common.TeamCounterTerrorists,
				status:  gameStatusUnknown,
			},
			wantSide: common.TeamTerrorists,
			wantErr:  false,
		},
		{
			name: "Swap from T to CT",
			fields: fields{
				name:    "TeamB",
				flag:    "FLAG_B",
				players: []uint64{4, 5, 6},
				score:   3,
				side:    common.TeamTerrorists,
				status:  gameStatusUnknown,
			},
			wantSide: common.TeamCounterTerrorists,
			wantErr:  false,
		},
		{
			name: "Spectators Side",
			fields: fields{
				name:    "TeamC",
				flag:    "FLAG_C",
				players: []uint64{7, 8, 9},
				score:   7,
				side:    common.TeamSpectators,
				status:  gameStatusUnknown,
			},
			wantSide: common.TeamSpectators,
			wantErr:  true,
		},
		{
			name: "Unassigned Side",
			fields: fields{
				name:    "TeamE",
				flag:    "FLAG_E",
				players: []uint64{13, 14, 15},
				score:   8,
				side:    common.TeamUnassigned,
				status:  gameStatusUnknown,
			},
			wantSide: common.TeamUnassigned,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &team{
				name:    tt.fields.name,
				flag:    tt.fields.flag,
				players: tt.fields.players,
				score:   tt.fields.score,
				side:    tt.fields.side,
				status:  tt.fields.status,
			}
			err := tr.swapSide()
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.wantSide, tr.side)
		})
	}
}
