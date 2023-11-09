package demoparser

import (
	"testing"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/stretchr/testify/assert"
)

func Test_parser_playerConnected(t *testing.T) {
	t.Parallel()
	testParser := demoinfocs.NewParser(newReadCloser("test parser"))
	type fields struct {
		Parser   demoinfocs.Parser
		demosize int64
	}
	type args struct {
		pl *common.Player
	}
	tests := []struct {
		args   args
		fields fields
		name   string
		want   bool
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
			fields: fields{
				Parser:   testParser,
				demosize: 1337,
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
			fields: fields{
				Parser:   testParser,
				demosize: 1337,
			},
			name: "Valid connected player",
			want: true,
		},
		{
			args: args{
				pl: &common.Player{
					SteamID64:   123456789,
					UserID:      133,
					Team:        common.TeamSpectators,
					IsBot:       false,
					IsConnected: true,
					IsUnknown:   false,
				},
			},
			fields: fields{
				Parser:   testParser,
				demosize: 1337,
			},
			name: "Player on Team Spectators",
			want: false,
		},
		{
			args: args{
				pl: &common.Player{
					SteamID64:   123456789,
					UserID:      133,
					Team:        common.TeamUnassigned,
					IsBot:       false,
					IsConnected: true,
					IsUnknown:   false,
				},
			},
			fields: fields{
				Parser:   testParser,
				demosize: 1337,
			},
			name: "Player on Team Unassigned",
			want: false,
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
			fields: fields{
				Parser:   testParser,
				demosize: 1337,
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
			fields: fields{
				Parser:   testParser,
				demosize: 1337,
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
			fields: fields{
				Parser:   testParser,
				demosize: 1337,
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
			fields: fields{
				Parser:   testParser,
				demosize: 1337,
			},
			name: "Unknown player",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				Parser:   tt.fields.Parser,
				demosize: tt.fields.demosize,
			}
			got := p.playerConnected(tt.args.pl)
			assert.Equal(t, tt.want, got)
		})
	}
}
