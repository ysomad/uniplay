package replay

import (
	"testing"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/stretchr/testify/assert"
)

func Test_newStats(t *testing.T) {
	tests := []struct {
		name string
		want stats
	}{
		{
			name: "success",
			want: stats{
				playerStats: make(map[uint64]*playerStat),
				weaponStats: make(map[uint64]map[common.EquipmentType]*weaponStat),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newStats()
			assert.Equal(t, tt.want, got)
		})
	}
}

func _generateTestStats() *stats {
	s := &stats{
		playerStats: make(map[uint64]*playerStat),
		weaponStats: make(map[uint64]map[common.EquipmentType]*weaponStat),
	}

	equipment := []common.EquipmentType{
		common.EqP2000,
		common.EqGlock,
		common.EqP250,
		common.EqDeagle,
		common.EqFiveSeven,
		common.EqTec9,
		common.EqUSP,
		common.EqMP7,
		common.EqGalil,
		common.EqFamas,
		common.EqAK47,
		common.EqM4A4,
		common.EqScout,
		common.EqAWP,
		common.EqKnife,
		common.EqWorld,
		common.EqMolotov,
		common.EqIncendiary,
		common.EqHE,
	}

	for i := 1; i <= 10; i++ {
		steamID := uint64(1)
		s.playerStats[steamID] = &playerStat{
			steamID:            steamID,
			kills:              500,
			hsKills:            200,
			blindKills:         100,
			wallbangKills:      50,
			noScopeKills:       20,
			throughSmokeKills:  15,
			deaths:             400,
			assists:            250,
			flashbangAssists:   50,
			mvpCount:           100,
			damageTaken:        50000,
			damageDealt:        45000,
			grenadeDamageDealt: 3000,
			blindedPlayers:     50,
			blindedTimes:       50,
			bombsPlanted:       100,
			bombsDefused:       50,
		}
	}

	for i := 1; i <= 10; i++ {
		steamID := uint64(i)
		s.weaponStats[steamID] = make(map[common.EquipmentType]*weaponStat)

		for _, eq := range equipment {
			s.weaponStats[steamID][eq] = &weaponStat{
				steamID:           steamID,
				weaponID:          1,
				kills:             100,
				hsKills:           20,
				blindKills:        10,
				wallbangKills:     15,
				noScopeKills:      10,
				throughSmokeKills: 5,
				deaths:            100,
				assists:           50,
				damageTaken:       10000,
				damageDealt:       10000,
				shots:             1000,
				headHits:          500,
				chestHits:         200,
				stomachHits:       300,
				leftArmHits:       150,
				rightArmHits:      150,
				leftLegHits:       100,
				rightLegHits:      100,
			}
		}
	}

	return s
}

func Benchmark_stats_normalizeSync(b *testing.B) {
	b.StopTimer()
	s := _generateTestStats()

	b.StartTimer()
	_, _ = s.normalizeSync()
}

func Benchmark_stats_normalize(b *testing.B) {
	b.StopTimer()
	s := _generateTestStats()

	b.StartTimer()
	_, _ = s.normalize()
}

func Test_stats_normalizeSync(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
		weaponStats map[uint64]map[common.EquipmentType]*weaponStat
	}

	tests := []struct {
		name   string
		fields fields
		want   []*playerStat
		want1  []*weaponStat
	}{
		{
			name: "empty stats",
			fields: fields{
				playerStats: map[uint64]*playerStat{},
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{},
			},
			want:  []*playerStat{},
			want1: nil,
		},
		{
			name: "2 players, 2 weapons",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					1: {
						steamID:            1,
						kills:              500,
						hsKills:            200,
						blindKills:         100,
						wallbangKills:      50,
						noScopeKills:       20,
						throughSmokeKills:  15,
						deaths:             400,
						assists:            250,
						flashbangAssists:   50,
						mvpCount:           100,
						damageTaken:        50000,
						damageDealt:        45000,
						grenadeDamageDealt: 3000,
						blindedPlayers:     50,
						blindedTimes:       50,
						bombsPlanted:       100,
						bombsDefused:       50,
					},
					2: {
						steamID:            2,
						kills:              500,
						hsKills:            200,
						blindKills:         100,
						wallbangKills:      50,
						noScopeKills:       20,
						throughSmokeKills:  15,
						deaths:             400,
						assists:            250,
						flashbangAssists:   50,
						mvpCount:           100,
						damageTaken:        50000,
						damageDealt:        45000,
						grenadeDamageDealt: 3000,
						blindedPlayers:     50,
						blindedTimes:       50,
						bombsPlanted:       100,
						bombsDefused:       50,
					},
				},
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					1: {
						common.EqAWP: {
							steamID:           1,
							weaponID:          int16(common.EqAWP),
							kills:             50,
							hsKills:           15,
							blindKills:        5,
							wallbangKills:     2,
							noScopeKills:      1,
							throughSmokeKills: 3,
							deaths:            25,
							assists:           10,
							damageTaken:       555,
							damageDealt:       1500,
							shots:             1000,
							headHits:          500,
							chestHits:         300,
							stomachHits:       200,
							leftArmHits:       50,
							rightArmHits:      50,
						},
						common.EqM4A1: {
							steamID:           1,
							weaponID:          int16(common.EqM4A1),
							kills:             23,
							hsKills:           10,
							blindKills:        2,
							wallbangKills:     1,
							noScopeKills:      1,
							throughSmokeKills: 5,
							deaths:            10,
							assists:           5,
							damageTaken:       1324,
							damageDealt:       32145,
							shots:             10000,
							headHits:          1235,
							chestHits:         231,
							stomachHits:       1235,
							leftArmHits:       21,
							rightArmHits:      24,
							leftLegHits:       55,
							rightLegHits:      68,
						},
					},
					2: {
						common.EqKnife: {
							steamID:     2,
							weaponID:    int16(common.EqKnife),
							kills:       3,
							deaths:      1,
							damageTaken: 100,
							damageDealt: 300,
							shots:       13,
							headHits:    1,
							chestHits:   5,
							stomachHits: 2,
						},
						common.EqHE: {
							steamID:     2,
							weaponID:    int16(common.EqHE),
							kills:       2,
							deaths:      5,
							assists:     2,
							damageTaken: 98,
							damageDealt: 156,
							shots:       10,
						},
					},
				},
			},
			want: []*playerStat{
				{
					steamID:            1,
					kills:              500,
					hsKills:            200,
					blindKills:         100,
					wallbangKills:      50,
					noScopeKills:       20,
					throughSmokeKills:  15,
					deaths:             400,
					assists:            250,
					flashbangAssists:   50,
					mvpCount:           100,
					damageTaken:        50000,
					damageDealt:        45000,
					grenadeDamageDealt: 3000,
					blindedPlayers:     50,
					blindedTimes:       50,
					bombsPlanted:       100,
					bombsDefused:       50,
				},
				{
					steamID:            2,
					kills:              500,
					hsKills:            200,
					blindKills:         100,
					wallbangKills:      50,
					noScopeKills:       20,
					throughSmokeKills:  15,
					deaths:             400,
					assists:            250,
					flashbangAssists:   50,
					mvpCount:           100,
					damageTaken:        50000,
					damageDealt:        45000,
					grenadeDamageDealt: 3000,
					blindedPlayers:     50,
					blindedTimes:       50,
					bombsPlanted:       100,
					bombsDefused:       50,
				},
			},
			want1: []*weaponStat{
				{
					steamID:           1,
					weaponID:          int16(common.EqAWP),
					kills:             50,
					hsKills:           15,
					blindKills:        5,
					wallbangKills:     2,
					noScopeKills:      1,
					throughSmokeKills: 3,
					deaths:            25,
					assists:           10,
					damageTaken:       555,
					damageDealt:       1500,
					shots:             1000,
					headHits:          500,
					chestHits:         300,
					stomachHits:       200,
					leftArmHits:       50,
					rightArmHits:      50,
				},
				{
					steamID:           1,
					weaponID:          int16(common.EqM4A1),
					kills:             23,
					hsKills:           10,
					blindKills:        2,
					wallbangKills:     1,
					noScopeKills:      1,
					throughSmokeKills: 5,
					deaths:            10,
					assists:           5,
					damageTaken:       1324,
					damageDealt:       32145,
					shots:             10000,
					headHits:          1235,
					chestHits:         231,
					stomachHits:       1235,
					leftArmHits:       21,
					rightArmHits:      24,
					leftLegHits:       55,
					rightLegHits:      68,
				},
				{
					steamID:     2,
					weaponID:    int16(common.EqKnife),
					kills:       3,
					deaths:      1,
					damageTaken: 100,
					damageDealt: 300,
					shots:       13,
					headHits:    1,
					chestHits:   5,
					stomachHits: 2,
				},
				{
					steamID:     2,
					weaponID:    int16(common.EqHE),
					kills:       2,
					deaths:      5,
					assists:     2,
					damageTaken: 98,
					damageDealt: 156,
					shots:       10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
			}
			got, got1 := s.normalizeSync()

			assert.ObjectsAreEqual(tt.want, got)
			assert.ObjectsAreEqual(tt.want1, got1)
		})
	}
}

func Test_stats_normalize(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
		weaponStats map[uint64]map[common.EquipmentType]*weaponStat
	}

	tests := []struct {
		name   string
		fields fields
		want   []*playerStat
		want1  []*weaponStat
	}{
		{
			name: "empty stats",
			fields: fields{
				playerStats: map[uint64]*playerStat{},
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{},
			},
			want:  nil,
			want1: nil,
		},
		{
			name: "2 players, 2 weapons",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					1: {
						steamID:            1,
						kills:              500,
						hsKills:            200,
						blindKills:         100,
						wallbangKills:      50,
						noScopeKills:       20,
						throughSmokeKills:  15,
						deaths:             400,
						assists:            250,
						flashbangAssists:   50,
						mvpCount:           100,
						damageTaken:        50000,
						damageDealt:        45000,
						grenadeDamageDealt: 3000,
						blindedPlayers:     50,
						blindedTimes:       50,
						bombsPlanted:       100,
						bombsDefused:       50,
					},
					2: {
						steamID:            2,
						kills:              500,
						hsKills:            200,
						blindKills:         100,
						wallbangKills:      50,
						noScopeKills:       20,
						throughSmokeKills:  15,
						deaths:             400,
						assists:            250,
						flashbangAssists:   50,
						mvpCount:           100,
						damageTaken:        50000,
						damageDealt:        45000,
						grenadeDamageDealt: 3000,
						blindedPlayers:     50,
						blindedTimes:       50,
						bombsPlanted:       100,
						bombsDefused:       50,
					},
				},
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					1: {
						common.EqAWP: {
							steamID:           1,
							weaponID:          int16(common.EqAWP),
							kills:             50,
							hsKills:           15,
							blindKills:        5,
							wallbangKills:     2,
							noScopeKills:      1,
							throughSmokeKills: 3,
							deaths:            25,
							assists:           10,
							damageTaken:       555,
							damageDealt:       1500,
							shots:             1000,
							headHits:          500,
							chestHits:         300,
							stomachHits:       200,
							leftArmHits:       50,
							rightArmHits:      50,
						},
						common.EqM4A1: {
							steamID:           1,
							weaponID:          int16(common.EqM4A1),
							kills:             23,
							hsKills:           10,
							blindKills:        2,
							wallbangKills:     1,
							noScopeKills:      1,
							throughSmokeKills: 5,
							deaths:            10,
							assists:           5,
							damageTaken:       1324,
							damageDealt:       32145,
							shots:             10000,
							headHits:          1235,
							chestHits:         231,
							stomachHits:       1235,
							leftArmHits:       21,
							rightArmHits:      24,
							leftLegHits:       55,
							rightLegHits:      68,
						},
					},
					2: {
						common.EqKnife: {
							steamID:     2,
							weaponID:    int16(common.EqKnife),
							kills:       3,
							deaths:      1,
							damageTaken: 100,
							damageDealt: 300,
							shots:       13,
							headHits:    1,
							chestHits:   5,
							stomachHits: 2,
						},
						common.EqHE: {
							steamID:     2,
							weaponID:    int16(common.EqHE),
							kills:       2,
							deaths:      5,
							assists:     2,
							damageTaken: 98,
							damageDealt: 156,
							shots:       10,
						},
					},
				},
			},
			want: []*playerStat{
				{
					steamID:            1,
					kills:              500,
					hsKills:            200,
					blindKills:         100,
					wallbangKills:      50,
					noScopeKills:       20,
					throughSmokeKills:  15,
					deaths:             400,
					assists:            250,
					flashbangAssists:   50,
					mvpCount:           100,
					damageTaken:        50000,
					damageDealt:        45000,
					grenadeDamageDealt: 3000,
					blindedPlayers:     50,
					blindedTimes:       50,
					bombsPlanted:       100,
					bombsDefused:       50,
				},
				{
					steamID:            2,
					kills:              500,
					hsKills:            200,
					blindKills:         100,
					wallbangKills:      50,
					noScopeKills:       20,
					throughSmokeKills:  15,
					deaths:             400,
					assists:            250,
					flashbangAssists:   50,
					mvpCount:           100,
					damageTaken:        50000,
					damageDealt:        45000,
					grenadeDamageDealt: 3000,
					blindedPlayers:     50,
					blindedTimes:       50,
					bombsPlanted:       100,
					bombsDefused:       50,
				},
			},
			want1: []*weaponStat{
				{
					steamID:           1,
					weaponID:          int16(common.EqAWP),
					kills:             50,
					hsKills:           15,
					blindKills:        5,
					wallbangKills:     2,
					noScopeKills:      1,
					throughSmokeKills: 3,
					deaths:            25,
					assists:           10,
					damageTaken:       555,
					damageDealt:       1500,
					shots:             1000,
					headHits:          500,
					chestHits:         300,
					stomachHits:       200,
					leftArmHits:       50,
					rightArmHits:      50,
				},
				{
					steamID:           1,
					weaponID:          int16(common.EqM4A1),
					kills:             23,
					hsKills:           10,
					blindKills:        2,
					wallbangKills:     1,
					noScopeKills:      1,
					throughSmokeKills: 5,
					deaths:            10,
					assists:           5,
					damageTaken:       1324,
					damageDealt:       32145,
					shots:             10000,
					headHits:          1235,
					chestHits:         231,
					stomachHits:       1235,
					leftArmHits:       21,
					rightArmHits:      24,
					leftLegHits:       55,
					rightLegHits:      68,
				},
				{
					steamID:     2,
					weaponID:    int16(common.EqKnife),
					kills:       3,
					deaths:      1,
					damageTaken: 100,
					damageDealt: 300,
					shots:       13,
					headHits:    1,
					chestHits:   5,
					stomachHits: 2,
				},
				{
					steamID:     2,
					weaponID:    int16(common.EqHE),
					kills:       2,
					deaths:      5,
					assists:     2,
					damageTaken: 98,
					damageDealt: 156,
					shots:       10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
			}
			got, got1 := s.normalize()

			assert.EqualValues(t, tt.want, got)
			assert.EqualValues(t, tt.want1, got1)
		})
	}
}

func Test_stats_addPlayerStat(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
	}
	type args struct {
		steamID uint64
		m       metric
		v       int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[uint64]*playerStat
	}{
		{
			name: "death",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					1: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 1,
				m:       metricDeath,
				v:       4,
			},
			want: map[uint64]*playerStat{
				1: {
					kills:  5,
					deaths: 7,
				},
				2: {
					damageTaken: 567,
					damageDealt: 456,
				},
			},
		},
		{
			name: "kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					1: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 2,
				m:       metricKill,
				v:       10,
			},
			want: map[uint64]*playerStat{
				1: {
					kills:  5,
					deaths: 3,
				},
				2: {
					kills:       10,
					damageTaken: 567,
					damageDealt: 456,
				},
			},
		},
		{
			name: "hs kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 4,
				m:       metricHSKill,
				v:       3,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				2: {
					damageTaken: 567,
					damageDealt: 456,
				},
				4: {hsKills: 3},
			},
		},

		{
			name: "blind kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 5,
				m:       metricBlindKill,
				v:       2,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				2: {
					damageTaken: 567,
					damageDealt: 456,
				},
				5: {blindKills: 2},
			},
		},
		{
			name: "wallbang kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken:   567,
						damageDealt:   456,
						wallbangKills: 33,
					},
				},
			},
			args: args{
				steamID: 2,
				m:       metricWallbangKill,
				v:       4,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				2: {
					damageTaken:   567,
					damageDealt:   456,
					wallbangKills: 37,
				},
			},
		},
		{
			name: "noscope kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:        5,
						deaths:       3,
						noScopeKills: 7,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 3,
				m:       metricNoScopeKill,
				v:       6,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:        5,
					deaths:       3,
					noScopeKills: 13,
				},
				2: {
					damageTaken: 567,
					damageDealt: 456,
				},
			},
		},
		{
			name: "through smoke kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					7: {},
				},
			},
			args: args{
				steamID: 7,
				m:       metricThroughSmokeKill,
				v:       1,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				7: {throughSmokeKills: 1},
			},
		},
		{
			name: "assist",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					13: {
						throughSmokeKills: 5,
						wallbangKills:     7,
						assists:           3,
					},
				},
			},
			args: args{
				steamID: 13,
				m:       metricAssist,
				v:       9,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				13: {
					throughSmokeKills: 5,
					wallbangKills:     7,
					assists:           12,
				},
			},
		},
		{
			name: "flashbang assist",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
					},
					13: {
						assists: 3,
					},
				},
			},
			args: args{
				steamID: 3,
				m:       metricFlashbangAssist,
				v:       2,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 7,
				},
				13: {
					assists: 3,
				},
			},
		},
		{
			name: "damage taken",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
					},
					13: {
						assists: 3,
					},
				},
			},
			args: args{
				steamID: 1,
				m:       metricDamageTaken,
				v:       99,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
				},
				13: {
					assists: 3,
				},
				1: {
					damageTaken: 99,
				},
			},
		},
		{
			name: "damage dealt",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
					},
					13: {
						assists:     3,
						damageDealt: 133,
					},
				},
			},
			args: args{
				steamID: 13,
				m:       metricDamageDealt,
				v:       7,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
				},
				13: {
					assists:     3,
					damageDealt: 140,
				},
			},
		},
		{
			name: "bomb plant",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					13: {
						assists:     3,
						damageDealt: 133,
					},
				},
			},
			args: args{
				steamID: 3,
				m:       metricBombPlanted,
				v:       1,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     78,
				},
				13: {
					assists:     3,
					damageDealt: 133,
				},
			},
		},
		{
			name: "bomb defuse",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					13: {
						assists:     3,
						damageDealt: 133,
					},
					5: {
						kills: 5,
					},
				},
			},
			args: args{
				steamID: 5,
				m:       metricBombDefused,
				v:       1,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     77,
				},
				13: {
					assists:     3,
					damageDealt: 133,
				},
				5: {
					kills:        5,
					bombsDefused: 1,
				},
			},
		},
		{
			name: "round mvp",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					4: {
						mvpCount: 3,
						kills:    133,
					},
				},
			},

			args: args{
				steamID: 4,
				m:       metricRoundMVP,
				v:       1,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     77,
				},
				4: {
					mvpCount: 4,
					kills:    133,
				},
			},
		},
		{
			name: "blind",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					4: {
						mvpCount: 3,
						kills:    133,
					},
				},
			},
			args: args{
				steamID: 5,
				m:       metricBlinded,
				v:       1,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     77,
				},
				4: {
					mvpCount: 3,
					kills:    133,
				},
				5: {blindedTimes: 1},
			},
		},
		{
			name: "blinded player",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					4: {
						mvpCount: 3,
						kills:    133,
					},
					5: {blindedTimes: 1},
				},
			},
			args: args{
				steamID: 5,
				m:       metricBlind,
				v:       3,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     77,
				},
				4: {
					mvpCount: 3,
					kills:    133,
				},
				5: {
					blindedTimes:   1,
					blindedPlayers: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
			}
			s.addPlayerStat(tt.args.steamID, tt.args.m, tt.args.v)

			assert.ObjectsAreEqual(tt.want, s.playerStats)
		})
	}
}

func Test_stats_incrPlayerStat(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
	}
	type args struct {
		steamID uint64
		m       metric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[uint64]*playerStat
	}{
		{
			name: "death",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					1: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 1,
				m:       metricDeath,
			},
			want: map[uint64]*playerStat{
				1: {
					kills:  5,
					deaths: 4,
				},
				2: {
					damageTaken: 567,
					damageDealt: 456,
				},
			},
		},
		{
			name: "kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					1: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 2,
				m:       metricKill,
			},
			want: map[uint64]*playerStat{
				1: {
					kills:  5,
					deaths: 3,
				},
				2: {
					kills:       6,
					damageTaken: 567,
					damageDealt: 456,
				},
			},
		},
		{
			name: "hs kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 4,
				m:       metricHSKill,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				2: {
					damageTaken: 567,
					damageDealt: 456,
				},
				4: {hsKills: 1},
			},
		},

		{
			name: "blind kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 5,
				m:       metricBlindKill,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				2: {
					damageTaken: 567,
					damageDealt: 456,
				},
				5: {blindKills: 1},
			},
		},
		{
			name: "wallbang kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					2: {
						damageTaken:   567,
						damageDealt:   456,
						wallbangKills: 33,
					},
				},
			},
			args: args{
				steamID: 2,
				m:       metricWallbangKill,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				2: {
					damageTaken:   567,
					damageDealt:   456,
					wallbangKills: 34,
				},
			},
		},
		{
			name: "noscope kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:        5,
						deaths:       3,
						noScopeKills: 7,
					},
					2: {
						damageTaken: 567,
						damageDealt: 456,
					},
				},
			},
			args: args{
				steamID: 3,
				m:       metricNoScopeKill,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:        5,
					deaths:       3,
					noScopeKills: 8,
				},
				2: {
					damageTaken: 567,
					damageDealt: 456,
				},
			},
		},
		{
			name: "through smoke kill",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					7: {},
				},
			},
			args: args{
				steamID: 7,
				m:       metricThroughSmokeKill,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				7: {throughSmokeKills: 1},
			},
		},
		{
			name: "assist",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:  5,
						deaths: 3,
					},
					13: {
						throughSmokeKills: 5,
						wallbangKills:     7,
						assists:           3,
					},
				},
			},
			args: args{
				steamID: 13,
				m:       metricAssist,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:  5,
					deaths: 3,
				},
				13: {
					throughSmokeKills: 5,
					wallbangKills:     7,
					assists:           4,
				},
			},
		},
		{
			name: "flashbang assist",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
					},
					13: {
						assists: 3,
					},
				},
			},
			args: args{
				steamID: 3,
				m:       metricFlashbangAssist,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 6,
				},
				13: {
					assists: 3,
				},
			},
		},
		{
			name: "damage taken",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
					},
					13: {
						assists: 3,
					},
				},
			},
			args: args{
				steamID: 1,
				m:       metricDamageTaken,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
				},
				13: {
					assists: 3,
				},
				1: {
					damageTaken: 1,
				},
			},
		},
		{
			name: "damage dealt",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
					},
					13: {
						assists:     3,
						damageDealt: 133,
					},
				},
			},
			args: args{
				steamID: 13,
				m:       metricDamageDealt,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
				},
				13: {
					assists:     3,
					damageDealt: 134,
				},
			},
		},
		{
			name: "bomb plant",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					13: {
						assists:     3,
						damageDealt: 133,
					},
				},
			},
			args: args{
				steamID: 3,
				m:       metricBombPlanted,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     78,
				},
				13: {
					assists:     3,
					damageDealt: 133,
				},
			},
		},
		{
			name: "bomb defuse",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					13: {
						assists:     3,
						damageDealt: 133,
					},
					5: {
						kills: 5,
					},
				},
			},
			args: args{
				steamID: 5,
				m:       metricBombDefused,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     77,
				},
				13: {
					assists:     3,
					damageDealt: 133,
				},
				5: {
					kills:        5,
					bombsDefused: 1,
				},
			},
		},
		{
			name: "round mvp",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					4: {
						mvpCount: 3,
						kills:    133,
					},
				},
			},

			args: args{
				steamID: 4,
				m:       metricRoundMVP,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     77,
				},
				4: {
					mvpCount: 4,
					kills:    133,
				},
			},
		},
		{
			name: "blind",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					4: {
						mvpCount: 3,
						kills:    133,
					},
					5: {blindedTimes: 5},
				},
			},
			args: args{
				steamID: 5,
				m:       metricBlinded,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     77,
				},
				4: {
					mvpCount: 3,
					kills:    133,
				},
				5: {blindedTimes: 6},
			},
		},
		{
			name: "blinded player",
			fields: fields{
				playerStats: map[uint64]*playerStat{
					3: {
						kills:            5,
						deaths:           3,
						flashbangAssists: 5,
						bombsPlanted:     77,
					},
					4: {
						mvpCount: 3,
						kills:    133,
					},
					5: {
						blindedTimes:   1,
						blindedPlayers: 3,
					},
				},
			},
			args: args{
				steamID: 5,
				m:       metricBlind,
			},
			want: map[uint64]*playerStat{
				3: {
					kills:            5,
					deaths:           3,
					flashbangAssists: 5,
					bombsPlanted:     77,
				},
				4: {
					mvpCount: 3,
					kills:    133,
				},
				5: {
					blindedTimes:   1,
					blindedPlayers: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
			}
			s.incrPlayerStat(tt.args.steamID, tt.args.m)

			assert.ObjectsAreEqual(tt.want, s.playerStats)
		})
	}
}

func Test_stats_validWeapon(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
		weaponStats map[uint64]map[common.EquipmentType]*weaponStat
	}
	type args struct {
		e common.EquipmentType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid weapon",
			args: args{
				e: common.EqGlock,
			},
			want: true,
		},
		{
			name: "invalid weapon",
			args: args{
				e: common.EqUnknown,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(stats)
			got := s.validWeapon(tt.args.e)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_stats_validMetric(t *testing.T) {
	type args struct {
		m metric
		v int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid metric with valid value",
			args: args{
				m: metricHSKill,
				v: 1,
			},
			want: true,
		},
		{
			name: "valid metric with invalid value",
			args: args{
				m: metricHSKill,
				v: 0,
			},
			want: false,
		},
		{
			name: "invalid metric and invalid value",
			args: args{
				m: 0,
				v: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(stats)
			got := s.validMetric(tt.args.m, tt.args.v)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_stats_addWeaponStat(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
		weaponStats map[uint64]map[common.EquipmentType]*weaponStat
	}
	type args struct {
		steamID uint64
		m       metric
		e       common.EquipmentType
		v       int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[uint64]map[common.EquipmentType]*weaponStat
	}{
		{
			name: "add 1 kill to existed weapon",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					2: {
						common.EqUSP: &weaponStat{kills: 4},
					},
				},
			},
			args: args{
				steamID: 2,
				m:       metricKill,
				e:       common.EqUSP,
				v:       1,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				2: {
					common.EqUSP: &weaponStat{kills: 5},
				},
			},
		},
		{
			name: "add 5 kills to not existing weapon",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					11: {
						common.EqUSP: &weaponStat{kills: 4},
					},
				},
			},
			args: args{
				steamID: 11,
				m:       metricKill,
				e:       common.EqAWP,
				v:       5,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				11: {
					common.EqUSP: &weaponStat{kills: 4},
					common.EqAWP: &weaponStat{kills: 5},
				},
			},
		},
		{
			name: "invalid metric, valid value",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					7: {
						common.EqM4A1: &weaponStat{
							kills:       11,
							damageDealt: 1342,
							damageTaken: 566,
							deaths:      5,
						},
					},
					11: {
						common.EqAWP: &weaponStat{
							kills:       5,
							damageDealt: 400,
							damageTaken: 500,
							deaths:      5,
						},
					},
				},
			},
			args: args{
				steamID: 7,
				m:       -1, // invalid metric
				e:       common.EqUSP,
				v:       5,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				7: {
					common.EqM4A1: &weaponStat{
						kills:       11,
						damageDealt: 1342,
						damageTaken: 566,
						deaths:      5,
					},
				},
				11: {
					common.EqAWP: &weaponStat{
						kills:       5,
						damageDealt: 400,
						damageTaken: 500,
						deaths:      5,
					},
				},
			},
		},
		{
			name: "invalid metric and invalid value",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					7: {
						common.EqM4A1: &weaponStat{
							kills:       11,
							damageDealt: 1342,
							damageTaken: 566,
							deaths:      5,
						},
					},
					11: {
						common.EqAWP: &weaponStat{
							kills:       5,
							damageDealt: 400,
							damageTaken: 500,
							deaths:      5,
						},
					},
				},
			},
			args: args{
				steamID: 7,
				m:       -1,           // invalid metric
				e:       common.EqUSP, // invalid value
				v:       0,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				7: {
					common.EqM4A1: &weaponStat{
						kills:       11,
						damageDealt: 1342,
						damageTaken: 566,
						deaths:      5,
					},
				},
				11: {
					common.EqAWP: &weaponStat{
						kills:       5,
						damageDealt: 400,
						damageTaken: 500,
						deaths:      5,
					},
				},
			},
		},
		{
			name: "invalid weapon",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					7: {
						common.EqM4A1: &weaponStat{
							kills:       11,
							damageDealt: 1342,
							damageTaken: 566,
							deaths:      5,
						},
					},
					11: {
						common.EqAWP: &weaponStat{
							kills:       5,
							damageDealt: 400,
							damageTaken: 500,
							deaths:      5,
						},
					},
				},
			},
			args: args{
				steamID: 7,
				m:       metricDeath,
				e:       common.EqUnknown, // invalid weapon
				v:       1,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				7: {
					common.EqM4A1: &weaponStat{
						kills:       11,
						damageDealt: 1342,
						damageTaken: 566,
						deaths:      5,
					},
				},
				11: {
					common.EqAWP: &weaponStat{
						kills:       5,
						damageDealt: 400,
						damageTaken: 500,
						deaths:      5,
					},
				},
			},
		},
		{
			name: "invalid steam id",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					7: {
						common.EqM4A1: &weaponStat{
							kills:       11,
							damageDealt: 1342,
							damageTaken: 566,
							deaths:      5,
						},
					},
					11: {
						common.EqAWP: &weaponStat{
							kills:       5,
							damageDealt: 400,
							damageTaken: 500,
							deaths:      5,
						},
					},
				},
			},
			args: args{
				steamID: 0, // invalid steam id
				m:       metricDamageDealt,
				e:       common.EqHE,
				v:       56,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				7: {
					common.EqM4A1: &weaponStat{
						kills:       11,
						damageDealt: 1342,
						damageTaken: 566,
						deaths:      5,
					},
				},
				11: {
					common.EqAWP: &weaponStat{
						kills:       5,
						damageDealt: 400,
						damageTaken: 500,
						deaths:      5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{weaponStats: tt.fields.weaponStats}
			s.addWeaponStat(tt.args.steamID, tt.args.m, tt.args.e, tt.args.v)

			assert.ObjectsAreEqual(tt.want, s.weaponStats)
		})
	}
}

func Test_stats_incrWeaponStat(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
		weaponStats map[uint64]map[common.EquipmentType]*weaponStat
	}
	type args struct {
		steamID uint64
		m       metric
		e       common.EquipmentType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[uint64]map[common.EquipmentType]*weaponStat
	}{

		{
			name: "add 1 kill to existed weapon",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					2: {
						common.EqUSP: &weaponStat{kills: 4},
					},
				},
			},
			args: args{
				steamID: 2,
				m:       metricKill,
				e:       common.EqUSP,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				2: {
					common.EqUSP: &weaponStat{kills: 5},
				},
			},
		},
		{
			name: "add 5 kills to not existing weapon",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					11: {
						common.EqUSP: &weaponStat{kills: 4},
					},
				},
			},
			args: args{
				steamID: 11,
				m:       metricKill,
				e:       common.EqAWP,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				11: {
					common.EqUSP: &weaponStat{kills: 4},
					common.EqAWP: &weaponStat{kills: 5},
				},
			},
		},
		{
			name: "invalid metric, valid value",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					7: {
						common.EqM4A1: &weaponStat{
							kills:       11,
							damageDealt: 1342,
							damageTaken: 566,
							deaths:      5,
						},
					},
					11: {
						common.EqAWP: &weaponStat{
							kills:       5,
							damageDealt: 400,
							damageTaken: 500,
							deaths:      5,
						},
					},
				},
			},
			args: args{
				steamID: 7,
				m:       -1, // invalid metric
				e:       common.EqUSP,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				7: {
					common.EqM4A1: &weaponStat{
						kills:       11,
						damageDealt: 1342,
						damageTaken: 566,
						deaths:      5,
					},
				},
				11: {
					common.EqAWP: &weaponStat{
						kills:       5,
						damageDealt: 400,
						damageTaken: 500,
						deaths:      5,
					},
				},
			},
		},
		{
			name: "invalid metric and invalid value",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					7: {
						common.EqM4A1: &weaponStat{
							kills:       11,
							damageDealt: 1342,
							damageTaken: 566,
							deaths:      5,
						},
					},
					11: {
						common.EqAWP: &weaponStat{
							kills:       5,
							damageDealt: 400,
							damageTaken: 500,
							deaths:      5,
						},
					},
				},
			},
			args: args{
				steamID: 7,
				m:       -1,           // invalid metric
				e:       common.EqUSP, // invalid value
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				7: {
					common.EqM4A1: &weaponStat{
						kills:       11,
						damageDealt: 1342,
						damageTaken: 566,
						deaths:      5,
					},
				},
				11: {
					common.EqAWP: &weaponStat{
						kills:       5,
						damageDealt: 400,
						damageTaken: 500,
						deaths:      5,
					},
				},
			},
		},
		{
			name: "invalid weapon",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					7: {
						common.EqM4A1: &weaponStat{
							kills:       11,
							damageDealt: 1342,
							damageTaken: 566,
							deaths:      5,
						},
					},
					11: {
						common.EqAWP: &weaponStat{
							kills:       5,
							damageDealt: 400,
							damageTaken: 500,
							deaths:      5,
						},
					},
				},
			},
			args: args{
				steamID: 7,
				m:       metricDeath,
				e:       common.EqUnknown, // invalid weapon
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				7: {
					common.EqM4A1: &weaponStat{
						kills:       11,
						damageDealt: 1342,
						damageTaken: 566,
						deaths:      5,
					},
				},
				11: {
					common.EqAWP: &weaponStat{
						kills:       5,
						damageDealt: 400,
						damageTaken: 500,
						deaths:      5,
					},
				},
			},
		},
		{
			name: "invalid steam id",
			fields: fields{
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{
					7: {
						common.EqM4A1: &weaponStat{
							kills:       11,
							damageDealt: 1342,
							damageTaken: 566,
							deaths:      5,
						},
					},
					11: {
						common.EqAWP: &weaponStat{
							kills:       5,
							damageDealt: 400,
							damageTaken: 500,
							deaths:      5,
						},
					},
				},
			},
			args: args{
				steamID: 0, // invalid steam id
				m:       metricDamageDealt,
				e:       common.EqHE,
			},
			want: map[uint64]map[common.EquipmentType]*weaponStat{
				7: {
					common.EqM4A1: &weaponStat{
						kills:       11,
						damageDealt: 1342,
						damageTaken: 566,
						deaths:      5,
					},
				},
				11: {
					common.EqAWP: &weaponStat{
						kills:       5,
						damageDealt: 400,
						damageTaken: 500,
						deaths:      5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
			}
			s.incrWeaponStat(tt.args.steamID, tt.args.m, tt.args.e)

			assert.ObjectsAreEqual(tt.want, s.weaponStats)
		})
	}
}

func Test_playerStat_add(t *testing.T) {
	type args struct {
		m metric
		v int
	}
	tests := []struct {
		name string
		args args
		want *playerStat
	}{
		{
			name: "death",
			args: args{
				m: metricDeath,
				v: 4,
			},
			want: &playerStat{
				deaths: 4,
			},
		},
		{
			name: "kill",
			args: args{
				m: metricKill,
				v: 10,
			},
			want: &playerStat{
				kills: 10,
			},
		},
		{
			name: "hs kill",
			args: args{
				m: metricHSKill,
				v: 3,
			},
			want: &playerStat{
				hsKills: 3,
			},
		},

		{
			name: "blind kill",
			args: args{
				m: metricBlindKill,
				v: 2,
			},
			want: &playerStat{
				blindKills: 2,
			},
		},
		{
			name: "wallbang kill",
			args: args{
				m: metricWallbangKill,
				v: 4,
			},
			want: &playerStat{
				wallbangKills: 4,
			},
		},
		{
			name: "noscope kill",
			args: args{
				m: metricNoScopeKill,
				v: 6,
			},
			want: &playerStat{
				noScopeKills: 6,
			},
		},
		{
			name: "through smoke kill",
			args: args{
				m: metricThroughSmokeKill,
				v: 1,
			},
			want: &playerStat{
				throughSmokeKills: 1,
			},
		},
		{
			name: "assist",
			args: args{
				m: metricAssist,
				v: 9,
			},
			want: &playerStat{
				assists: 9,
			},
		},
		{
			name: "flashbang assist",
			args: args{
				m: metricFlashbangAssist,
				v: 2,
			},
			want: &playerStat{
				flashbangAssists: 2,
			},
		},
		{
			name: "damage taken",
			args: args{
				m: metricDamageTaken,
				v: 99,
			},
			want: &playerStat{
				damageTaken: 99,
			},
		},
		{
			name: "damage dealt",
			args: args{
				m: metricDamageDealt,
				v: 35,
			},
			want: &playerStat{
				damageDealt: 35,
			},
		},
		{
			name: "bomb plant",
			args: args{
				m: metricBombPlanted,
				v: 1,
			},
			want: &playerStat{
				bombsPlanted: 1,
			},
		},
		{
			name: "bomb defuse",
			args: args{
				m: metricBombDefused,
				v: 1,
			},
			want: &playerStat{
				bombsDefused: 1,
			},
		},
		{
			name: "round mvp",
			args: args{
				m: metricRoundMVP,
				v: 1,
			},
			want: &playerStat{
				mvpCount: 1,
			},
		},
		{
			name: "blind",
			args: args{
				m: metricBlinded,
				v: 1,
			},
			want: &playerStat{
				blindedTimes: 1,
			},
		},
		{
			name: "blinded player",
			args: args{
				m: metricBlind,
				v: 1,
			},
			want: &playerStat{
				blindedPlayers: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(playerStat)
			s.add(tt.args.m, tt.args.v)

			assert.Equal(t, tt.want, s)
		})
	}
}

func Test_weaponStat_add(t *testing.T) {
	type args struct {
		m metric
		v int
	}
	tests := []struct {
		name string
		in   *weaponStat
		args args
		want *weaponStat
	}{
		{
			name: "invalid metric, valid value",
			in:   &weaponStat{kills: 35, damageDealt: 1356, shots: 13456},
			args: args{
				m: 0,
				v: 1,
			},
			want: &weaponStat{kills: 35, damageDealt: 1356, shots: 13456},
		},
		{
			name: "invalid metric and invalid value",
			in:   &weaponStat{kills: 35, damageDealt: 1356, shots: 13456},
			args: args{
				m: 0,
				v: 0,
			},
			want: &weaponStat{kills: 35, damageDealt: 1356, shots: 13456},
		},
		{
			name: "valid metric and invalid value",
			in:   &weaponStat{kills: 35, damageDealt: 1356, shots: 13456},
			args: args{
				m: metricHitHead,
				v: 0,
			},
			want: &weaponStat{kills: 35, damageDealt: 1356, shots: 13456},
		},
		{
			name: "kill",
			in:   &weaponStat{kills: 35, damageDealt: 1356, shots: 13456},
			args: args{
				m: metricKill,
				v: 3,
			},
			want: &weaponStat{kills: 38, damageDealt: 1356, shots: 13456},
		},
		{
			name: "hs kill",
			in:   &weaponStat{kills: 35, hsKills: 13, damageDealt: 1356, shots: 13456},
			args: args{
				m: metricHSKill,
				v: 5,
			},
			want: &weaponStat{kills: 35, hsKills: 18, damageDealt: 1356, shots: 13456},
		},
		{
			name: "blind kill",
			in:   &weaponStat{kills: 35, hsKills: 13, damageDealt: 1356, shots: 13456, blindKills: 3},
			args: args{
				m: metricBlindKill,
				v: 1,
			},
			want: &weaponStat{kills: 35, hsKills: 13, damageDealt: 1356, shots: 13456, blindKills: 4},
		},
		{
			name: "wb kill",
			in:   &weaponStat{kills: 35, damageDealt: 1356, blindKills: 3, wallbangKills: 7},
			args: args{
				m: metricWallbangKill,
				v: 7,
			},
			want: &weaponStat{kills: 35, damageDealt: 1356, blindKills: 3, wallbangKills: 14},
		},
		{
			name: "noscope kill",
			in:   &weaponStat{kills: 35, damageDealt: 1356, blindKills: 3, wallbangKills: 7, noScopeKills: 5},
			args: args{
				m: metricNoScopeKill,
				v: 3,
			},
			want: &weaponStat{kills: 35, damageDealt: 1356, blindKills: 3, wallbangKills: 7, noScopeKills: 8},
		},
		{
			name: "through smoke kill",
			in:   &weaponStat{kills: 35, blindKills: 3, wallbangKills: 7, noScopeKills: 8, throughSmokeKills: 13},
			args: args{
				m: metricThroughSmokeKill,
				v: 3,
			},
			want: &weaponStat{kills: 35, blindKills: 3, wallbangKills: 7, noScopeKills: 8, throughSmokeKills: 16},
		},
		{
			name: "death",
			in:   &weaponStat{kills: 35, deaths: 28},
			args: args{
				m: metricDeath,
				v: 1,
			},
			want: &weaponStat{kills: 35, deaths: 29},
		},
		{
			name: "assist",
			in:   &weaponStat{kills: 35, deaths: 28, assists: 15},
			args: args{
				m: metricAssist,
				v: 10,
			},
			want: &weaponStat{kills: 35, deaths: 28, assists: 25},
		},
		{
			name: "damage taken",
			in:   &weaponStat{kills: 35, deaths: 28, assists: 15, damageTaken: 1346},
			args: args{
				m: metricDamageTaken,
				v: 136,
			},
			want: &weaponStat{kills: 35, deaths: 28, assists: 15, damageTaken: 1482},
		},
		{
			name: "damage taken",
			in:   &weaponStat{kills: 35, deaths: 28, assists: 15, damageTaken: 1346},
			args: args{
				m: metricDamageTaken,
				v: 136,
			},
			want: &weaponStat{kills: 35, deaths: 28, assists: 15, damageTaken: 1482},
		},
		{
			name: "damage dealt",
			in:   &weaponStat{kills: 35, deaths: 28, assists: 15, damageTaken: 1346, damageDealt: 567},
			args: args{
				m: metricDamageDealt,
				v: 76,
			},
			want: &weaponStat{kills: 35, deaths: 28, assists: 15, damageTaken: 1346, damageDealt: 643},
		},
		{
			name: "shot",
			in:   &weaponStat{kills: 35, deaths: 28, shots: 79},
			args: args{
				m: metricShot,
				v: 1,
			},
			want: &weaponStat{kills: 35, deaths: 28, shots: 80},
		},
		{
			name: "head hit",
			in:   &weaponStat{kills: 35, deaths: 28, shots: 79, headHits: 5},
			args: args{
				m: metricHitHead,
				v: 5,
			},
			want: &weaponStat{kills: 35, deaths: 28, shots: 79, headHits: 10},
		},
		{
			name: "chest hit",
			in:   &weaponStat{kills: 35, deaths: 28, shots: 79, chestHits: 3},
			args: args{
				m: metricHitChest,
				v: 7,
			},
			want: &weaponStat{kills: 35, deaths: 28, shots: 79, chestHits: 10},
		},
		{
			name: "stomach hit",
			in:   &weaponStat{kills: 35, deaths: 28, shots: 79, stomachHits: 3},
			args: args{
				m: metricHitStomach,
				v: 7,
			},
			want: &weaponStat{kills: 35, deaths: 28, shots: 79, stomachHits: 10},
		},
		{
			name: "left arm hit",
			in:   &weaponStat{kills: 35, deaths: 28, shots: 79, leftArmHits: 45},
			args: args{
				m: metricHitLeftArm,
				v: 1,
			},
			want: &weaponStat{kills: 35, deaths: 28, shots: 79, leftArmHits: 46},
		},
		{
			name: "right arm hit",
			in:   &weaponStat{kills: 35, deaths: 28, shots: 79, rightArmHits: 13},
			args: args{
				m: metricHitRightArm,
				v: 11,
			},
			want: &weaponStat{kills: 35, deaths: 28, shots: 79, rightArmHits: 24},
		},
		{
			name: "left leg hit",
			in:   &weaponStat{kills: 35, deaths: 28, shots: 79, leftLegHits: 5},
			args: args{
				m: metricHitLeftLeg,
				v: 3,
			},
			want: &weaponStat{kills: 35, deaths: 28, shots: 79, leftLegHits: 8},
		},
		{
			name: "right leg hit",
			in:   &weaponStat{kills: 35, deaths: 28, shots: 79, rightLegHits: 9},
			args: args{
				m: metricHitRightLeg,
				v: 1,
			},
			want: &weaponStat{kills: 35, deaths: 28, shots: 79, rightLegHits: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.add(tt.args.m, tt.args.v)
			assert.Equal(t, tt.want, tt.in)
		})
	}
}
