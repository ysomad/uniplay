package demoparser

import (
	"testing"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/stretchr/testify/assert"
)

func Test_playerConnected(t *testing.T) {
	t.Parallel()
	type args struct {
		pl *common.Player
	}
	tests := []struct {
		args args
		name string
		want bool
	}{
		{
			args: args{
				pl: &common.Player{
					SteamID64:   0,
					Team:        0,
					UserID:      0,
					IsBot:       false,
					IsConnected: false,
					IsUnknown:   false,
				},
			},
			name: "Empty player",
			want: false,
		},
		{
			args: args{
				pl: &common.Player{
					SteamID64:   123456789,
					UserID:      133,
					Team:        common.TeamTerrorists,
					IsBot:       false,
					IsConnected: true,
					IsUnknown:   false,
				},
			},
			name: "Valid connected player",
			want: true,
		},
		{
			args: args{
				pl: &common.Player{
					SteamID64:   0,
					UserID:      133,
					Team:        common.TeamTerrorists,
					IsBot:       false,
					IsConnected: true,
					IsUnknown:   false,
				},
			},
			name: "Player with zero SteamID64",
			want: false,
		},
		{
			args: args{
				pl: &common.Player{
					SteamID64:   133,
					UserID:      0,
					Team:        common.TeamTerrorists,
					IsBot:       false,
					IsConnected: true,
					IsUnknown:   false,
				},
			},
			name: "Player with zero UserID",
			want: false,
		},
		{
			args: args{
				pl: &common.Player{
					SteamID64:   123456789,
					UserID:      133,
					Team:        common.TeamCounterTerrorists,
					IsBot:       true,
					IsConnected: true,
					IsUnknown:   false,
				},
			},
			name: "Bot player",
			want: false,
		},
		{
			args: args{
				pl: &common.Player{
					SteamID64:   123456789,
					UserID:      133,
					Team:        common.TeamTerrorists,
					IsBot:       false,
					IsConnected: true,
					IsUnknown:   true,
				},
			},
			name: "Unknown player",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := playerConnected(tt.args.pl)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_playerSpectator(t *testing.T) {
	t.Parallel()
	type args struct {
		pl *common.Player
	}
	tests := []struct {
		args args
		name string
		want bool
	}{
		{
			args: args{
				pl: &common.Player{
					Team: common.TeamSpectators,
				},
			},
			name: "Player in Team Spectators",
			want: true,
		},
		{
			args: args{
				pl: &common.Player{
					Team: common.TeamUnassigned,
				},
			},
			name: "Player in Team Unassigned",
			want: true,
		},
		{
			args: args{
				pl: &common.Player{
					Team: common.TeamTerrorists,
				},
			},
			name: "Player in Team Terrorists",
			want: false,
		},
		{
			args: args{
				pl: &common.Player{
					Team: common.TeamCounterTerrorists,
				},
			},
			name: "Player in Team Counter-Terrorists",
			want: false,
		},
		{
			args: args{
				pl: nil,
			},
			name: "Nil Player",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := playerSpectator(tt.args.pl)
			assert.Equal(t, tt.want, got)
		})
	}
}
