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
	t.Parallel()
	type fields struct {
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
				knifeRound: tt.fields.knifeRound,
				started:    tt.fields.started,
			}
			got := gs.collectStats()
			assert.Equal(t, tt.want, got)
		})
	}
}
