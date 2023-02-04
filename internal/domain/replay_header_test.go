package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewReplayHeader(t *testing.T) {
	type args struct {
		ticks        int
		frames       int
		signonLen    int
		server       string
		client       string
		mapName      string
		playbackTime time.Duration
		filesize     int64
	}
	tests := []struct {
		name    string
		args    args
		want    *ReplayHeader
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ticks:        1488,
				frames:       1337,
				signonLen:    228,
				server:       "test",
				client:       "test",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want: &ReplayHeader{
				playbackTicks:  1488,
				playbackFrames: 1337,
				signonLength:   228,
				server:         "test",
				client:         "test",
				mapName:        "de_dust",
				playbackTime:   time.Minute * 25,
				filesize:       425632453253245323,
			},
			wantErr: false,
		},
		{
			name: "0 ticks",
			args: args{
				ticks:        0,
				frames:       1337,
				signonLen:    228,
				server:       "test",
				client:       "test",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative ticks",
			args: args{
				ticks:        -5,
				frames:       1337,
				signonLen:    228,
				server:       "test",
				client:       "test",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "0 frames",
			args: args{
				ticks:        1488,
				frames:       0,
				signonLen:    228,
				server:       "test",
				client:       "test",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative frames",
			args: args{
				ticks:        1488,
				frames:       -5,
				signonLen:    228,
				server:       "test",
				client:       "test",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "0 signon",
			args: args{
				ticks:        1488,
				frames:       133777,
				signonLen:    0,
				server:       "test",
				client:       "test",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative signon",
			args: args{
				ticks:        1488,
				frames:       133777,
				signonLen:    -5,
				server:       "test",
				client:       "test",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "empty server name",
			args: args{
				ticks:        1488,
				frames:       133777,
				signonLen:    123,
				server:       "",
				client:       "test",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty client name",
			args: args{
				ticks:        1488,
				frames:       133777,
				signonLen:    123,
				server:       "test",
				client:       "",
				mapName:      "de_dust2",
				playbackTime: time.Minute * 25,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid playback time",
			args: args{
				ticks:        1488,
				frames:       133777,
				signonLen:    123,
				server:       "test",
				client:       "",
				mapName:      "de_dust2",
				playbackTime: minMatchDuration - time.Minute,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative playback time",
			args: args{
				ticks:        1488,
				frames:       133777,
				signonLen:    123,
				server:       "test",
				client:       "",
				mapName:      "de_dust2",
				playbackTime: -9999,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "0 playback time",
			args: args{
				ticks:        1488,
				frames:       133777,
				signonLen:    123,
				server:       "test",
				client:       "",
				mapName:      "de_dust2",
				playbackTime: 0,
				filesize:     425632453253245323,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReplayHeader(tt.args.ticks, tt.args.frames, tt.args.signonLen, tt.args.server, tt.args.client, tt.args.mapName, tt.args.playbackTime, tt.args.filesize)

			assert.Equal(t, tt.wantErr, (err != nil))
			assert.ObjectsAreEqual(tt.want, got)
		})
	}
}
