package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWeaponStats(t *testing.T) {
	type args struct {
		total []*WeaponTotalStat
	}
	tests := []struct {
		name string
		args args
		want []WeaponStat
	}{
		{
			name: "success",
			args: args{
				total: []*WeaponTotalStat{
					{
						WeaponID:          1,
						Weapon:            "ak-47",
						Kills:             1337,
						HeadshotKills:     228,
						BlindKills:        99,
						WallbangKills:     123,
						NoScopeKills:      543,
						ThroughSmokeKills: 111,
						Deaths:            546,
						Assists:           123,
						DamageTaken:       12345123,
						DamageDealt:       23452353,
						Shots:             1589,
						HeadHits:          321,
						ChestHits:         123,
						StomachHits:       256,
						LeftArmHits:       95,
						RightArmHits:      64,
						LeftLegHits:       45,
						RightLegHits:      88,
					},
					{
						WeaponID:          2,
						Weapon:            "mp4a4",
						Kills:             1337,
						HeadshotKills:     228,
						BlindKills:        99,
						WallbangKills:     123,
						NoScopeKills:      543,
						ThroughSmokeKills: 111,
						Deaths:            546,
						Assists:           123,
						DamageTaken:       12345123,
						DamageDealt:       23452353,
						Shots:             12345,
						HeadHits:          5431,
						ChestHits:         312,
						StomachHits:       123,
						LeftArmHits:       435,
						RightArmHits:      23,
						LeftLegHits:       65,
						RightLegHits:      54,
					},
					{
						WeaponID:          3,
						Weapon:            "awp",
						Kills:             1337,
						HeadshotKills:     228,
						BlindKills:        99,
						WallbangKills:     123,
						NoScopeKills:      543,
						ThroughSmokeKills: 111,
						Deaths:            546,
						Assists:           123,
						DamageTaken:       12345123,
						DamageDealt:       23452353,
						Shots:             12356,
						HeadHits:          653,
						ChestHits:         341,
						StomachHits:       312,
						LeftArmHits:       234,
						RightArmHits:      54,
						LeftLegHits:       45,
						RightLegHits:      11,
					},
				},
			},
			want: []WeaponStat{
				{
					TotalStat: &WeaponTotalStat{
						WeaponID:          1,
						Weapon:            "ak-47",
						Kills:             1337,
						HeadshotKills:     228,
						BlindKills:        99,
						WallbangKills:     123,
						NoScopeKills:      543,
						ThroughSmokeKills: 111,
						Deaths:            546,
						Assists:           123,
						DamageTaken:       12345123,
						DamageDealt:       23452353,
						Shots:             1589,
						HeadHits:          321,
						ChestHits:         123,
						StomachHits:       256,
						LeftArmHits:       95,
						RightArmHits:      64,
						LeftLegHits:       45,
						RightLegHits:      88,
					},
					AccuracyStat: WeaponAccuracyStat{
						Total:   62.43,
						Head:    32.36,
						Chest:   12.4,
						Stomach: 25.81,
						Arms:    16.03,
						Legs:    13.41,
					},
				},
				{
					TotalStat: &WeaponTotalStat{
						WeaponID:          2,
						Weapon:            "mp4a4",
						Kills:             1337,
						HeadshotKills:     228,
						BlindKills:        99,
						WallbangKills:     123,
						NoScopeKills:      543,
						ThroughSmokeKills: 111,
						Deaths:            546,
						Assists:           123,
						DamageTaken:       12345123,
						DamageDealt:       23452353,
						Shots:             12345,
						HeadHits:          5431,
						ChestHits:         312,
						StomachHits:       123,
						LeftArmHits:       435,
						RightArmHits:      23,
						LeftLegHits:       65,
						RightLegHits:      54,
					},
					AccuracyStat: WeaponAccuracyStat{
						Total:   52.19,
						Head:    84.29,
						Chest:   4.84,
						Stomach: 1.91,
						Arms:    7.11,
						Legs:    1.85,
					},
				},
				{
					TotalStat: &WeaponTotalStat{
						WeaponID:          3,
						Weapon:            "awp",
						Kills:             1337,
						HeadshotKills:     228,
						BlindKills:        99,
						WallbangKills:     123,
						NoScopeKills:      543,
						ThroughSmokeKills: 111,
						Deaths:            546,
						Assists:           123,
						DamageTaken:       12345123,
						DamageDealt:       23452353,
						Shots:             12356,
						HeadHits:          653,
						ChestHits:         341,
						StomachHits:       312,
						LeftArmHits:       234,
						RightArmHits:      54,
						LeftLegHits:       45,
						RightLegHits:      11,
					},
					AccuracyStat: WeaponAccuracyStat{
						Total:   13.35,
						Head:    39.58,
						Chest:   20.67,
						Stomach: 18.91,
						Arms:    17.45,
						Legs:    3.39,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewWeaponStats(tt.args.total)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_round(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{63.312312123},
			want: 63.31,
		},
		{
			name: "2",
			args: args{0},
			want: 00.00,
		},
		{
			name: "3",
			args: args{999.9999999},
			want: 1000,
		},
		{
			name: "4",
			args: args{123.567},
			want: 123.57,
		},
		{
			name: "5",
			args: args{0.99999},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := round(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calcAccuracy(t *testing.T) {
	type args struct {
		sum int32
		num int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "100%",
			args: args{
				sum: 100,
				num: 100,
			},
			want: 100,
		},
		{
			name: "50%",
			args: args{
				sum: 50,
				num: 100,
			},
			want: 50,
		},
		{
			name: "33.33%",
			args: args{
				sum: 33,
				num: 99,
			},
			want: 33.33,
		},
		{
			name: "100% 2",
			args: args{
				sum: 9999,
				num: 9999,
			},
			want: 100,
		},
		{
			name: "0%",
			args: args{
				sum: 0,
				num: 0,
			},
			want: 0,
		},
		{
			name: "0%",
			args: args{
				sum: -5,
				num: -5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcAccuracy(tt.args.sum, tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newWeaponAccuracyStat(t *testing.T) {
	type args struct {
		shots       int32
		headHits    int32
		chestHits   int32
		stomachHits int32
		lArmHits    int32
		rArmHits    int32
		lLegHits    int32
		rLegHits    int32
	}
	tests := []struct {
		name string
		args args
		want WeaponAccuracyStat
	}{
		{
			name: "success",
			args: args{
				shots:       1589,
				headHits:    321,
				chestHits:   123,
				stomachHits: 256,
				lArmHits:    95,
				rArmHits:    64,
				lLegHits:    45,
				rLegHits:    88,
			},
			want: WeaponAccuracyStat{
				Total:   62.43,
				Head:    32.36,
				Chest:   12.4,
				Stomach: 25.81,
				Arms:    16.03,
				Legs:    13.41,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newWeaponAccuracyStat(tt.args.shots, tt.args.headHits, tt.args.chestHits, tt.args.stomachHits, tt.args.lArmHits, tt.args.rArmHits, tt.args.lLegHits, tt.args.rLegHits)
			assert.Equal(t, tt.want, got)
		})
	}
}