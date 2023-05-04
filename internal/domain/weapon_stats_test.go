package domain

// func TestNewWeaponStats(t *testing.T) {
// 	t.Parallel()

// 	type args struct {
// 		total []*WeaponBaseStats
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []WeaponStats
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				total: []*WeaponBaseStats{
// 					{
// 						WeaponID:          1,
// 						Weapon:            "ak-47",
// 						Kills:             1337,
// 						HeadshotKills:     228,
// 						BlindKills:        99,
// 						WallbangKills:     123,
// 						NoScopeKills:      543,
// 						ThroughSmokeKills: 111,
// 						Deaths:            546,
// 						Assists:           123,
// 						DamageTaken:       12345123,
// 						DamageDealt:       23452353,
// 						Shots:             1589,
// 						HeadHits:          321,
// 						ChestHits:         123,
// 						StomachHits:       256,
// 						LeftArmHits:       95,
// 						RightArmHits:      64,
// 						LeftLegHits:       45,
// 						RightLegHits:      88,
// 					},
// 					{
// 						WeaponID:          2,
// 						Weapon:            "mp4a4",
// 						Kills:             1337,
// 						HeadshotKills:     228,
// 						BlindKills:        99,
// 						WallbangKills:     123,
// 						NoScopeKills:      543,
// 						ThroughSmokeKills: 111,
// 						Deaths:            546,
// 						Assists:           123,
// 						DamageTaken:       12345123,
// 						DamageDealt:       23452353,
// 						Shots:             12345,
// 						HeadHits:          5431,
// 						ChestHits:         312,
// 						StomachHits:       123,
// 						LeftArmHits:       435,
// 						RightArmHits:      23,
// 						LeftLegHits:       65,
// 						RightLegHits:      54,
// 					},
// 					{
// 						WeaponID:          3,
// 						Weapon:            "awp",
// 						Kills:             1337,
// 						HeadshotKills:     228,
// 						BlindKills:        99,
// 						WallbangKills:     123,
// 						NoScopeKills:      543,
// 						ThroughSmokeKills: 111,
// 						Deaths:            546,
// 						Assists:           123,
// 						DamageTaken:       12345123,
// 						DamageDealt:       23452353,
// 						Shots:             12356,
// 						HeadHits:          653,
// 						ChestHits:         341,
// 						StomachHits:       312,
// 						LeftArmHits:       234,
// 						RightArmHits:      54,
// 						LeftLegHits:       45,
// 						RightLegHits:      11,
// 					},
// 				},
// 			},
// 			want: []WeaponStats{
// 				{
// 					Base: &WeaponBaseStats{
// 						WeaponID:          1,
// 						Weapon:            "ak-47",
// 						Kills:             1337,
// 						HeadshotKills:     228,
// 						BlindKills:        99,
// 						WallbangKills:     123,
// 						NoScopeKills:      543,
// 						ThroughSmokeKills: 111,
// 						Deaths:            546,
// 						Assists:           123,
// 						DamageTaken:       12345123,
// 						DamageDealt:       23452353,
// 						Shots:             1589,
// 						HeadHits:          321,
// 						ChestHits:         123,
// 						StomachHits:       256,
// 						LeftArmHits:       95,
// 						RightArmHits:      64,
// 						LeftLegHits:       45,
// 						RightLegHits:      88,
// 					},
// 					Accuracy: WeaponAccuracyStats{
// 						Total:   62.43,
// 						Head:    32.36,
// 						Chest:   12.4,
// 						Stomach: 25.81,
// 						Arms:    16.03,
// 						Legs:    13.41,
// 					},
// 				},
// 				{
// 					Base: &WeaponBaseStats{
// 						WeaponID:          2,
// 						Weapon:            "mp4a4",
// 						Kills:             1337,
// 						HeadshotKills:     228,
// 						BlindKills:        99,
// 						WallbangKills:     123,
// 						NoScopeKills:      543,
// 						ThroughSmokeKills: 111,
// 						Deaths:            546,
// 						Assists:           123,
// 						DamageTaken:       12345123,
// 						DamageDealt:       23452353,
// 						Shots:             12345,
// 						HeadHits:          5431,
// 						ChestHits:         312,
// 						StomachHits:       123,
// 						LeftArmHits:       435,
// 						RightArmHits:      23,
// 						LeftLegHits:       65,
// 						RightLegHits:      54,
// 					},
// 					Accuracy: WeaponAccuracyStats{
// 						Total:   52.19,
// 						Head:    84.29,
// 						Chest:   4.84,
// 						Stomach: 1.91,
// 						Arms:    7.11,
// 						Legs:    1.85,
// 					},
// 				},
// 				{
// 					Base: &WeaponBaseStats{
// 						WeaponID:          3,
// 						Weapon:            "awp",
// 						Kills:             1337,
// 						HeadshotKills:     228,
// 						BlindKills:        99,
// 						WallbangKills:     123,
// 						NoScopeKills:      543,
// 						ThroughSmokeKills: 111,
// 						Deaths:            546,
// 						Assists:           123,
// 						DamageTaken:       12345123,
// 						DamageDealt:       23452353,
// 						Shots:             12356,
// 						HeadHits:          653,
// 						ChestHits:         341,
// 						StomachHits:       312,
// 						LeftArmHits:       234,
// 						RightArmHits:      54,
// 						LeftLegHits:       45,
// 						RightLegHits:      11,
// 					},
// 					Accuracy: WeaponAccuracyStats{
// 						Total:   13.35,
// 						Head:    39.58,
// 						Chest:   20.67,
// 						Stomach: 18.91,
// 						Arms:    17.45,
// 						Legs:    3.39,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := NewWeaponStats(tt.args.total)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }

// func Test_calcAccuracy(t *testing.T) {
// 	t.Parallel()

// 	type args struct {
// 		sum int32
// 		num int32
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want float64
// 	}{
// 		{
// 			name: "100%",
// 			args: args{
// 				sum: 100,
// 				num: 100,
// 			},
// 			want: 100,
// 		},
// 		{
// 			name: "50%",
// 			args: args{
// 				sum: 50,
// 				num: 100,
// 			},
// 			want: 50,
// 		},
// 		{
// 			name: "33.33%",
// 			args: args{
// 				sum: 33,
// 				num: 99,
// 			},
// 			want: 33.33,
// 		},
// 		{
// 			name: "100% 2",
// 			args: args{
// 				sum: 9999,
// 				num: 9999,
// 			},
// 			want: 100,
// 		},
// 		{
// 			name: "0%",
// 			args: args{
// 				sum: 0,
// 				num: 0,
// 			},
// 			want: 0,
// 		},
// 		{
// 			name: "0%",
// 			args: args{
// 				sum: -5,
// 				num: -5,
// 			},
// 			want: 0,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := calcAccuracy(tt.args.sum, tt.args.num)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }

// func Test_newWeaponAccuracyStats(t *testing.T) {
// 	t.Parallel()

// 	type args struct {
// 		shots       int32
// 		headHits    int32
// 		neckHits    int32
// 		chestHits   int32
// 		stomachHits int32
// 		lArmHits    int32
// 		rArmHits    int32
// 		lLegHits    int32
// 		rLegHits    int32
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want WeaponAccuracyStats
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				shots:       1589,
// 				headHits:    321,
// 				neckHits:    131,
// 				chestHits:   123,
// 				stomachHits: 256,
// 				lArmHits:    95,
// 				rArmHits:    64,
// 				lLegHits:    45,
// 				rLegHits:    88,
// 			},
// 			want: WeaponAccuracyStats{
// 				Total:   70.67,
// 				Head:    28.58,
// 				Neck:    11.67,
// 				Chest:   10.95,
// 				Stomach: 22.80,
// 				Arms:    14.16,
// 				Legs:    11.84,
// 			},
// 		},
// 		{
// 			name: "0 hits",
// 			args: args{
// 				shots:       1589,
// 				headHits:    0,
// 				neckHits:    0,
// 				chestHits:   0,
// 				stomachHits: 0,
// 				lArmHits:    0,
// 				rArmHits:    0,
// 				lLegHits:    0,
// 				rLegHits:    0,
// 			},
// 			want: WeaponAccuracyStats{},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := newWeaponAccuracyStats(
// 				tt.args.shots,
// 				tt.args.headHits,
// 				tt.args.neckHits,
// 				tt.args.chestHits,
// 				tt.args.stomachHits,
// 				tt.args.lArmHits,
// 				tt.args.rArmHits,
// 				tt.args.lLegHits,
// 				tt.args.rLegHits)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }
