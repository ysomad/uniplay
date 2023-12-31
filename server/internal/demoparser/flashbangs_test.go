package demoparser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

func Test_newPlayerFlashbands(t *testing.T) {
	tests := map[string]struct {
		want playerFlashbangs
	}{
		"success": {
			want: map[uint64][]flashbang{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, newPlayerFlashbands())
		})
	}
}

func Test_playerFlashbangs_add(t *testing.T) {
	t.Parallel()
	type args struct {
		throwerSteamID uint64
		f              flashbang
	}
	tests := map[string]struct {
		pf   playerFlashbangs
		args args
		want playerFlashbangs
	}{
		"AddFlashbangToExistingPlayer": {
			pf: playerFlashbangs{
				123: []flashbang{
					{ThrowerSide: common.TeamCounterTerrorists, Victim: 456, RoundNum: 1, Duration: 5 * time.Second},
				},
			},
			args: args{
				throwerSteamID: 123,
				f: flashbang{
					ThrowerSide: common.TeamTerrorists,
					Victim:      789,
					RoundNum:    2,
					Duration:    10 * time.Second,
				},
			},
			want: playerFlashbangs{
				123: []flashbang{
					{ThrowerSide: common.TeamCounterTerrorists, Victim: 456, RoundNum: 1, Duration: 5 * time.Second},
					{ThrowerSide: common.TeamTerrorists, Victim: 789, RoundNum: 2, Duration: 10 * time.Second},
				},
			},
		},
		"AddFlashbangToNewPlayer": {
			pf: playerFlashbangs{},
			args: args{
				throwerSteamID: 999,
				f: flashbang{
					ThrowerSide: common.TeamCounterTerrorists,
					Victim:      111,
					RoundNum:    3,
					Duration:    8 * time.Second,
				},
			},
			want: playerFlashbangs{
				999: []flashbang{
					{ThrowerSide: common.TeamCounterTerrorists, Victim: 111, RoundNum: 3, Duration: 8 * time.Second},
				},
			},
		},
		"AddMultipleFlashbangsToSamePlayer": {
			pf: playerFlashbangs{},
			args: args{
				throwerSteamID: 777,
				f: flashbang{
					ThrowerSide: common.TeamTerrorists,
					Victim:      222,
					RoundNum:    4,
					Duration:    15 * time.Second,
				},
			},
			want: playerFlashbangs{
				777: []flashbang{
					{ThrowerSide: common.TeamTerrorists, Victim: 222, RoundNum: 4, Duration: 15 * time.Second},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.pf.add(tt.args.throwerSteamID, tt.args.f)
			assert.Equal(t, tt.want, tt.pf)
		})
	}
}
