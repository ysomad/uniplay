package domain

import (
	"reflect"
	"testing"
)

func TestNewPlayerStats(t *testing.T) {
	type args struct {
		t *PlayerTotalStats
	}
	tests := []struct {
		name string
		args args
		want PlayerStats
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlayerStats(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newPlayerCalcStats(t *testing.T) {
	type args struct {
		kills         int32
		deaths        int32
		hsKills       int32
		wins          int32
		matchesPlayed int32
	}
	tests := []struct {
		name string
		args args
		want PlayerCalcStats
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPlayerCalcStats(tt.args.kills, tt.args.deaths, tt.args.hsKills, tt.args.wins, tt.args.matchesPlayed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPlayerCalcStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newPlayerRoundStats(t *testing.T) {
	type args struct {
		kills           int32
		deaths          int32
		dmgDealt        int32
		assists         int32
		grenadeDmgDealt int32
		blindedPlayers  int32
		blindedTimes    int32
		roundsPlayed    int32
	}
	tests := []struct {
		name string
		args args
		want PlayerRoundStats
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPlayerRoundStats(tt.args.kills, tt.args.deaths, tt.args.dmgDealt, tt.args.assists, tt.args.grenadeDmgDealt, tt.args.blindedPlayers, tt.args.blindedTimes, tt.args.roundsPlayed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPlayerRoundStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
