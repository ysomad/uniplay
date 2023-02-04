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
		h *ReplayHeader
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				h: &ReplayHeader{
					server:         "server",
					client:         "client",
					mapName:        "de_dust2",
					playbackTime:   time.Minute * 25,
					playbackTicks:  int(time.Minute * 25 * 128), // 128 server tickrate
					playbackFrames: int(time.Minute * 25 * 128),
					signonLength:   1337,
					filesize:       5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMatchID(tt.args.h)
			assert.NotEmpty(t, got)
		})
	}
}
