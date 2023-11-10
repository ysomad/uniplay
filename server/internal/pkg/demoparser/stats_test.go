package demoparser

import (
	"testing"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/stretchr/testify/assert"
)

func Test_playerStatsMap_add(t *testing.T) {
	t.Parallel()
	type args struct {
		steamID uint64
		ev      event
		val     int
	}
	tests := []struct {
		name string
		psm  playerStatsMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.psm.add(tt.args.steamID, tt.args.ev, tt.args.val)
		})
	}
}

func Test_playerStats_add(t *testing.T) {
	t.Parallel()
	type fields struct {
		kills          *killStats
		damage         *dmgGrenadeStats
		deaths         int
		assists        int
		fbAssists      int
		mvps           int
		blindedPlayers int
		blindedTimes   int
	}
	type args struct {
		e event
		v int
	}
	tests := []struct {
		want   *playerStats
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventKill,
				v: 5,
			},
			want: &playerStats{
				Kills: &killStats{
					Total: 5,
				},
			},
		},
		{
			name: "Add HS kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventHSKill,
				v: 3,
			},
			want: &playerStats{
				Kills: &killStats{
					HS: 3,
				},
			},
		},
		{
			name: "Add Blind kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventBlindKill,
				v: 2,
			},
			want: &playerStats{
				Kills: &killStats{
					Blind: 2,
				},
			},
		},
		{
			name: "Add WB kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventWBKill,
				v: 1,
			},
			want: &playerStats{
				Kills: &killStats{
					Wallbang: 1,
				},
			},
		},
		{
			name: "Add NoScope kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventNoScopeKill,
				v: 33,
			},
			want: &playerStats{
				Kills: &killStats{
					NoScope: 33,
				},
			},
		},
		{
			name: "Add Smoke kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventSmokeKill,
				v: 4,
			},
			want: &playerStats{
				Kills: &killStats{
					Smoke: 4,
				},
			},
		},
		{
			name: "Add Death",
			fields: fields{
				deaths: 0,
			},
			args: args{
				e: eventDeath,
				v: 1,
			},
			want: &playerStats{
				Deaths: 1,
			},
		},
		{
			name: "Add Assist",
			fields: fields{
				assists: 0,
			},
			args: args{
				e: eventAssist,
				v: 2,
			},
			want: &playerStats{
				Assists: 2,
			},
		},
		{
			name: "Add FB Assist",
			fields: fields{
				fbAssists: 0,
			},
			args: args{
				e: eventFBAssist,
				v: 1,
			},
			want: &playerStats{
				FlashbangAssists: 1,
			},
		},
		{
			name: "Add MVP",
			fields: fields{
				mvps: 0,
			},
			args: args{
				e: eventRoundMVP,
				v: 3,
			},
			want: &playerStats{
				MVPs: 3,
			},
		},
		{
			name: "Add Blinded Player",
			fields: fields{
				blindedPlayers: 0,
			},
			args: args{
				e: eventBlindedPlayer,
				v: 2,
			},
			want: &playerStats{
				BlindedPlayers: 2,
			},
		},
		{
			name: "Add Became Blind",
			fields: fields{
				blindedTimes: 0,
			},
			args: args{
				e: eventBecameBlind,
				v: 5,
			},
			want: &playerStats{
				BlindedTimes: 5,
			},
		},
		{
			name: "Add Dealt damage",
			fields: fields{
				damage: &dmgGrenadeStats{},
			},
			args: args{
				e: eventDmgDealt,
				v: 50,
			},
			want: &playerStats{
				Damage: &dmgGrenadeStats{
					dmgStats: dmgStats{
						Dealt: 50,
						Taken: 0,
					},
					DealtWithGrenade: 0,
				},
			},
		},
		{
			name: "Add Taken damage",
			fields: fields{
				damage: &dmgGrenadeStats{},
			},
			args: args{
				e: eventDmgTaken,
				v: 30,
			},
			want: &playerStats{
				Damage: &dmgGrenadeStats{
					dmgStats: dmgStats{
						Dealt: 0,
						Taken: 30,
					},
					DealtWithGrenade: 0,
				},
			},
		},
		{
			name: "Add Dealt with Grenade damage",
			fields: fields{
				damage: &dmgGrenadeStats{},
			},
			args: args{
				e: eventDmgGrenadeDealt,
				v: 15,
			},
			want: &playerStats{
				Damage: &dmgGrenadeStats{
					dmgStats: dmgStats{
						Dealt: 0,
						Taken: 0,
					},
					DealtWithGrenade: 15,
				},
			},
		},
		{
			name:   "Default case",
			fields: fields{},
			args: args{
				e: event(0), // unknown event
				v: 10,
			},
			want: &playerStats{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &playerStats{
				Kills:            tt.fields.kills,
				Damage:           tt.fields.damage,
				Deaths:           tt.fields.deaths,
				Assists:          tt.fields.assists,
				FlashbangAssists: tt.fields.fbAssists,
				MVPs:             tt.fields.mvps,
				BlindedPlayers:   tt.fields.blindedPlayers,
				BlindedTimes:     tt.fields.blindedTimes,
			}
			ps.add(tt.args.e, tt.args.v)
			assert.Equal(t, tt.want, ps)
		})
	}
}

func Test_equipValid(t *testing.T) {
	t.Parallel()
	type args struct {
		e common.EquipmentType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Unknown equipment type",
			args: args{
				e: common.EqUnknown,
			},
			want: false,
		},
		{
			name: "Kevlar",
			args: args{
				e: common.EqKevlar,
			},
			want: false,
		},
		{
			name: "Helmet",
			args: args{
				e: common.EqHelmet,
			},
			want: false,
		},
		{
			name: "Defuse Kit",
			args: args{
				e: common.EqDefuseKit,
			},
			want: false,
		},
		{
			name: "Valid equipment type",
			args: args{
				e: common.EqP2000,
			},
			want: true,
		},
		{
			name: "Another valid equipment type",
			args: args{
				e: common.EqM4A4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equipValid(tt.args.e)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_weaponStatsMap_add(t *testing.T) {
	t.Parallel()
	type args struct {
		steamID uint64
		ev      event
		et      common.EquipmentType
		val     int
	}
	tests := []struct {
		ws   weaponStatsMap
		want weaponStatsMap
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ws.add(tt.args.steamID, tt.args.ev, tt.args.et, tt.args.val)
			assert.Equal(t, tt.want, tt.ws)
		})
	}
}

func Test_weaponStats_add(t *testing.T) {
	t.Parallel()
	type fields struct {
		hits    *hitStats
		kills   *killStats
		damage  dmgStats
		deaths  int
		assists int
		shots   int
	}
	type args struct {
		e event
		v int
	}
	tests := []struct {
		want   *weaponStats
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventKill,
				v: 5,
			},
			want: &weaponStats{
				Kills: &killStats{
					Total: 5,
				},
			},
		},
		{
			name: "Add HS kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventHSKill,
				v: 3,
			},
			want: &weaponStats{
				Kills: &killStats{
					HS: 3,
				},
			},
		},
		{
			name: "Add Blind kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventBlindKill,
				v: 2,
			},
			want: &weaponStats{
				Kills: &killStats{
					Blind: 2,
				},
			},
		},
		{
			name: "Add NoScope kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventNoScopeKill,
				v: 5,
			},
			want: &weaponStats{
				Kills: &killStats{
					NoScope: 5,
				},
			},
		},

		{
			name: "Add WB kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventWBKill,
				v: 1,
			},
			want: &weaponStats{
				Kills: &killStats{
					Wallbang: 1,
				},
			},
		},
		{
			name: "Add Smoke kills",
			fields: fields{
				kills: &killStats{},
			},
			args: args{
				e: eventSmokeKill,
				v: 4,
			},
			want: &weaponStats{
				Kills: &killStats{
					Smoke: 4,
				},
			},
		},
		{
			name: "Add Death",
			fields: fields{
				deaths: 0,
			},
			args: args{
				e: eventDeath,
				v: 1,
			},
			want: &weaponStats{
				Deaths: 1,
			},
		},
		{
			name: "Add Assist",
			fields: fields{
				assists: 0,
			},
			args: args{
				e: eventAssist,
				v: 2,
			},
			want: &weaponStats{
				Assists: 2,
			},
		},
		{
			name: "Add Dealt damage",
			fields: fields{
				damage: dmgStats{},
			},
			args: args{
				e: eventDmgDealt,
				v: 50,
			},
			want: &weaponStats{
				Damage: dmgStats{
					Dealt: 50,
					Taken: 0,
				},
			},
		},
		{
			name: "Add Taken damage",
			fields: fields{
				damage: dmgStats{},
			},
			args: args{
				e: eventDmgTaken,
				v: 30,
			},
			want: &weaponStats{
				Damage: dmgStats{
					Dealt: 0,
					Taken: 30,
				},
			},
		},
		{
			name: "Add Shot",
			fields: fields{
				shots: 0,
			},
			args: args{
				e: eventShot,
				v: 5,
			},
			want: &weaponStats{
				Shots: 5,
			},
		},
		{
			name: "Add Hit Head",
			fields: fields{
				hits: &hitStats{},
			},
			args: args{
				e: eventHitHead,
				v: 2,
			},
			want: &weaponStats{
				Hits: &hitStats{
					Head: 2,
				},
			},
		},
		{
			name: "Add Hit Neck",
			fields: fields{
				hits: &hitStats{},
			},
			args: args{
				e: eventHitNeck,
				v: 1,
			},
			want: &weaponStats{
				Hits: &hitStats{
					Neck: 1,
				},
			},
		},
		{
			name: "Add Hit Stomach",
			fields: fields{
				hits: &hitStats{},
			},
			args: args{
				e: eventHitStomach,
				v: 3,
			},
			want: &weaponStats{
				Hits: &hitStats{
					Stomach: 3,
				},
			},
		},
		{
			name: "Add Hit Arm",
			fields: fields{
				hits: &hitStats{},
			},
			args: args{
				e: eventHitArm,
				v: 1,
			},
			want: &weaponStats{
				Hits: &hitStats{
					Arms: 1,
				},
			},
		},
		{
			name: "Add Hit Leg",
			fields: fields{
				hits: &hitStats{},
			},
			args: args{
				e: eventHitLeg,
				v: 2,
			},
			want: &weaponStats{
				Hits: &hitStats{
					Legs: 2,
				},
			},
		},
		{
			name:   "Default case",
			fields: fields{},
			args: args{
				e: event(0),
				v: 10,
			},
			want: &weaponStats{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &weaponStats{
				Hits:    tt.fields.hits,
				Kills:   tt.fields.kills,
				Damage:  tt.fields.damage,
				Deaths:  tt.fields.deaths,
				Assists: tt.fields.assists,
				Shots:   tt.fields.shots,
			}
			ws.add(tt.args.e, tt.args.v)
			assert.Equal(t, tt.want, ws)
		})
	}
}