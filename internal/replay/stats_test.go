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

func Test_stats_normalize(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
		weaponStats map[uint64]map[common.EquipmentType]*weaponStat
	}
	tests := []struct {
		name   string
		fields fields
		want   []playerStat
		want1  []weaponStat
	}{
		{
			name: "empty stats",
			fields: fields{
				playerStats: map[uint64]*playerStat{},
				weaponStats: map[uint64]map[common.EquipmentType]*weaponStat{},
			},
			want:  []playerStat{},
			want1: []weaponStat{},
		},
		// TODO: write success case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
			}
			got, got1 := s.normalize()

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_stats_addPlayerStat(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
		weaponStats map[uint64]map[common.EquipmentType]*weaponStat
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
			}
			s.addPlayerStat(tt.args.steamID, tt.args.m, tt.args.v)
		})
	}
}

func Test_stats_incrPlayerStat(t *testing.T) {
	type fields struct {
		playerStats map[uint64]*playerStat
		weaponStats map[uint64]map[common.EquipmentType]*weaponStat
	}
	type args struct {
		steamID uint64
		m       metric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
			}
			s.incrPlayerStat(tt.args.steamID, tt.args.m)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
			}
			s.addWeaponStat(tt.args.steamID, tt.args.m, tt.args.e, tt.args.v)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stats{
				playerStats: tt.fields.playerStats,
				weaponStats: tt.fields.weaponStats,
			}
			s.incrWeaponStat(tt.args.steamID, tt.args.m, tt.args.e)
		})
	}
}

func Test_playerStat_add(t *testing.T) {
	t.Parallel()

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
	type fields struct {
		steamID           uint64
		weaponID          int16
		kills             int
		hsKills           int
		blindKills        int
		wallbangKills     int
		noScopeKills      int
		throughSmokeKills int
		deaths            int
		assists           int
		damageTaken       int
		damageDealt       int
		shots             int
		headHits          int
		chestHits         int
		stomachHits       int
		leftArmHits       int
		rightArmHits      int
		leftLegHits       int
		rightLegHits      int
	}
	type args struct {
		m metric
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &weaponStat{
				steamID:           tt.fields.steamID,
				weaponID:          tt.fields.weaponID,
				kills:             tt.fields.kills,
				hsKills:           tt.fields.hsKills,
				blindKills:        tt.fields.blindKills,
				wallbangKills:     tt.fields.wallbangKills,
				noScopeKills:      tt.fields.noScopeKills,
				throughSmokeKills: tt.fields.throughSmokeKills,
				deaths:            tt.fields.deaths,
				assists:           tt.fields.assists,
				damageTaken:       tt.fields.damageTaken,
				damageDealt:       tt.fields.damageDealt,
				shots:             tt.fields.shots,
				headHits:          tt.fields.headHits,
				chestHits:         tt.fields.chestHits,
				stomachHits:       tt.fields.stomachHits,
				leftArmHits:       tt.fields.leftArmHits,
				rightArmHits:      tt.fields.rightArmHits,
				leftLegHits:       tt.fields.leftLegHits,
				rightLegHits:      tt.fields.rightLegHits,
			}
			ws.add(tt.args.m, tt.args.v)
		})
	}
}
