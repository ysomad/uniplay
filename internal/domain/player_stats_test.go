package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
		{
			name: "success",
			args: args{
				t: &PlayerTotalStats{
					Kills:              1337,
					HeadshotKills:      567,
					BlindKills:         59,
					WallbangKills:      98,
					NoScopeKills:       156,
					ThroughSmokeKills:  23,
					Deaths:             999,
					Assists:            333,
					FlashbangAssists:   19,
					MVPCount:           67,
					DamageTaken:        456356,
					DamageDealt:        345154,
					GrenadeDamageDealt: 12345,
					BlindedPlayers:     234,
					BlindedTimes:       345,
					BombsPlanted:       25,
					BombsDefused:       23,
					RoundsPlayed:       634,
					MatchesPlayed:      22,
					Wins:               18,
					Loses:              4,
					Draws:              0,
					TimePlayed:         time.Minute * 20 * 20,
				},
			},
			want: PlayerStats{
				Total: &PlayerTotalStats{
					Kills:              1337,
					HeadshotKills:      567,
					BlindKills:         59,
					WallbangKills:      98,
					NoScopeKills:       156,
					ThroughSmokeKills:  23,
					Deaths:             999,
					Assists:            333,
					FlashbangAssists:   19,
					MVPCount:           67,
					DamageTaken:        456356,
					DamageDealt:        345154,
					GrenadeDamageDealt: 12345,
					BlindedPlayers:     234,
					BlindedTimes:       345,
					BombsPlanted:       25,
					BombsDefused:       23,
					RoundsPlayed:       634,
					MatchesPlayed:      22,
					Wins:               18,
					Loses:              4,
					Draws:              0,
					TimePlayed:         time.Minute * 20 * 20,
				},
				Calc: PlayerCalcStats{
					HeadshotPercentage: 42.41,
					KillDeathRatio:     1.34,
					WinRate:            81.82,
				},
				Round: PlayerRoundStats{
					Kills:              2.11,
					Assists:            0.53,
					Deaths:             1.58,
					DamageDealt:        544.41,
					GrenadeDamageDealt: 19.47,
					BlindedPlayers:     0.37,
					BlindedTimes:       0.54,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPlayerStats(tt.args.t)
			assert.Equal(t, tt.want, got)
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
		{
			name: "0 kills",
			args: args{
				kills:         0,
				deaths:        1337,
				hsKills:       567,
				wins:          55,
				matchesPlayed: 213,
			},
			want: PlayerCalcStats{
				HeadshotPercentage: 0,
				KillDeathRatio:     0,
				WinRate:            25.82,
			},
		},
		{
			name: "negative number of kills",
			args: args{
				kills:         -13,
				deaths:        1337,
				hsKills:       567,
				wins:          55,
				matchesPlayed: 213,
			},
			want: PlayerCalcStats{
				HeadshotPercentage: 0,
				KillDeathRatio:     0,
				WinRate:            25.82,
			},
		},
		{
			name: "negative number of deaths",
			args: args{
				kills:         1337,
				deaths:        -56,
				hsKills:       567,
				wins:          55,
				matchesPlayed: 213,
			},
			want: PlayerCalcStats{
				HeadshotPercentage: 0,
				KillDeathRatio:     0,
				WinRate:            25.82,
			},
		},
		{
			name: "0 deaths",
			args: args{
				kills:         1337,
				deaths:        0,
				hsKills:       567,
				wins:          55,
				matchesPlayed: 213,
			},
			want: PlayerCalcStats{
				HeadshotPercentage: 0,
				KillDeathRatio:     0,
				WinRate:            25.82,
			},
		},
		{
			name: "0 matches played",
			args: args{
				kills:         1337,
				deaths:        228,
				hsKills:       567,
				wins:          55,
				matchesPlayed: 0,
			},
			want: PlayerCalcStats{
				HeadshotPercentage: 42.41,
				KillDeathRatio:     5.86,
				WinRate:            0,
			},
		},
		{
			name: "negative number of wins",
			args: args{
				kills:         1337,
				deaths:        228,
				hsKills:       567,
				wins:          -5,
				matchesPlayed: 213,
			},
			want: PlayerCalcStats{
				HeadshotPercentage: 42.41,
				KillDeathRatio:     5.86,
				WinRate:            0,
			},
		},
		{
			name: "0 wins",
			args: args{
				kills:         1337,
				deaths:        228,
				hsKills:       567,
				wins:          0,
				matchesPlayed: 213,
			},
			want: PlayerCalcStats{
				HeadshotPercentage: 42.41,
				KillDeathRatio:     5.86,
				WinRate:            0,
			},
		},
		{
			name: "negative number of matches played",
			args: args{
				kills:         1337,
				deaths:        228,
				hsKills:       567,
				wins:          55,
				matchesPlayed: -54,
			},
			want: PlayerCalcStats{
				HeadshotPercentage: 42.41,
				KillDeathRatio:     5.86,
				WinRate:            0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newPlayerCalcStats(tt.args.kills, tt.args.deaths, tt.args.hsKills, tt.args.wins, tt.args.matchesPlayed)
			assert.Equal(t, tt.want, got)
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
		{
			name: "success",
			args: args{
				kills:           1337,
				deaths:          1067,
				dmgDealt:        324566,
				assists:         324,
				grenadeDmgDealt: 23004,
				blindedPlayers:  56,
				blindedTimes:    99,
				roundsPlayed:    613,
			},
			want: PlayerRoundStats{
				Kills:              2.18,
				Assists:            0.53,
				Deaths:             1.74,
				DamageDealt:        529.47,
				GrenadeDamageDealt: 37.53,
				BlindedPlayers:     0.09,
				BlindedTimes:       0.16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newPlayerRoundStats(tt.args.kills, tt.args.deaths, tt.args.dmgDealt, tt.args.assists, tt.args.grenadeDmgDealt, tt.args.blindedPlayers, tt.args.blindedTimes, tt.args.roundsPlayed)
			assert.Equal(t, tt.want, got)
		})
	}
}
