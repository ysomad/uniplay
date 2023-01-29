package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMatchState(t *testing.T) {
	t.Parallel()

	type args struct {
		teamScore     int8
		opponentScore int8
	}
	tests := []struct {
		name string
		args args
		want MatchState
	}{
		{
			name: "win 16-15",
			args: args{
				teamScore:     16,
				opponentScore: 15,
			},
			want: MatchStateWin,
		},
		{
			name: "win 16-0",
			args: args{
				teamScore:     16,
				opponentScore: 0,
			},
			want: MatchStateWin,
		},
		{
			name: "win 16-4",
			args: args{
				teamScore:     16,
				opponentScore: 0,
			},
			want: MatchStateWin,
		},
		{
			name: "win 22-16",
			args: args{
				teamScore:     22,
				opponentScore: 16,
			},
			want: MatchStateWin,
		},

		{
			name: "lose 5-16",
			args: args{
				teamScore:     5,
				opponentScore: 16,
			},
			want: MatchStateLose,
		},
		{
			name: "lose 0-16",
			args: args{
				teamScore:     0,
				opponentScore: 16,
			},
			want: MatchStateLose,
		},
		{
			name: "lose 16-22",
			args: args{
				teamScore:     16,
				opponentScore: 22,
			},
			want: MatchStateLose,
		},
		{
			name: "lose 12-16",
			args: args{
				teamScore:     12,
				opponentScore: 16,
			},
			want: MatchStateLose,
		},
		{
			name: "draw 15-15",
			args: args{
				teamScore:     15,
				opponentScore: 15,
			},
			want: MatchStateDraw,
		},
		{
			name: "draw 0-0",
			args: args{
				teamScore:     0,
				opponentScore: 0,
			},
			want: MatchStateDraw,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMatchState(tt.args.teamScore, tt.args.opponentScore)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewMatchID(t *testing.T) {
	t.Parallel()

	type args struct {
		server        string
		client        string
		mapName       string
		matchDuration time.Duration
		ticks         int
		frames        int
		signonLen     int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				server:        "server",
				client:        "client",
				mapName:       "de_dust2",
				matchDuration: time.Minute * 25,
				ticks:         int(time.Minute * 25 * 128), // 128 server tickrate
				frames:        int(time.Minute * 25 * 128),
				signonLen:     1337,
			},
			wantErr: false,
		},
		{
			name: "invalid server",
			args: args{
				server:        "",
				client:        "client",
				mapName:       "de_dust2",
				matchDuration: time.Minute * 25,
				ticks:         int(time.Minute * 25 * 128),
				frames:        int(time.Minute * 25 * 128),
				signonLen:     1337,
			},
			wantErr: true,
		},
		{
			name: "invalid client",
			args: args{
				server:        "server",
				client:        "",
				mapName:       "de_dust2",
				matchDuration: time.Minute * 25,
				ticks:         int(time.Minute * 25 * 128),
				frames:        int(time.Minute * 25 * 128),
				signonLen:     1337,
			},
			wantErr: true,
		},
		{
			name: "match too short",
			args: args{
				server:        "server",
				client:        "client",
				mapName:       "de_dust2",
				matchDuration: time.Minute * 3,
				ticks:         int(time.Minute * 3 * 128),
				frames:        int(time.Minute * 3 * 128),
				signonLen:     1337,
			},
			wantErr: true,
		},
		{
			name: "invalid amount of ticks",
			args: args{
				server:        "server",
				client:        "client",
				mapName:       "de_dust2",
				matchDuration: time.Minute * 25,
				ticks:         0,
				frames:        int(time.Minute * 25 * 128),
				signonLen:     1337,
			},
			wantErr: true,
		},
		{
			name: "invalid amount of frames",
			args: args{
				server:        "server",
				client:        "client",
				mapName:       "de_dust2",
				matchDuration: time.Minute * 25,
				ticks:         int(time.Minute * 25 * 128),
				frames:        0,
				signonLen:     1337,
			},
			wantErr: true,
		},
		{
			name: "invalid amount of signon length",
			args: args{
				server:        "server",
				client:        "client",
				mapName:       "de_dust2",
				matchDuration: time.Minute * 25,
				ticks:         int(time.Minute * 25 * 128),
				frames:        int(time.Minute * 24),
				signonLen:     0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMatchID(tt.args.server, tt.args.client, tt.args.mapName, tt.args.matchDuration, tt.args.ticks, tt.args.frames, tt.args.signonLen)
			if err != nil {
				assert.Empty(t, got)
				assert.Equal(t, tt.wantErr, true)
				return
			}

			assert.NotEmpty(t, got)
		})
	}
}
