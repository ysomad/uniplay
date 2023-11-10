package demoparser

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_demoHeader_validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		server         string
		client         string
		mapName        string
		playbackTicks  int
		playbackFrames int
		signonLength   int
		playbackTime   time.Duration
		filesize       int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Valid Header",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: 16, server: "example.com", client: "player1", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: false,
		},
		{
			name:    "Valid Header (Future uploadedAt - 1m)",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: 16, server: "example.com", client: "player1", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: false,
		},

		{
			name:    "Invalid Header (Zero Values)",
			fields:  fields{},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Negative Playback Ticks)",
			fields:  fields{playbackTicks: -1, playbackFrames: 256, signonLength: 16, server: "example.com", client: "player1", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Negative Playback Frames)",
			fields:  fields{playbackTicks: 128, playbackFrames: -1, signonLength: 16, server: "example.com", client: "player1", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Negative Signon Length)",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: -1, server: "example.com", client: "player1", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Zero Signon Length)",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: 0, server: "example.com", client: "player1", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Empty Server)",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: 16, server: "", client: "player1", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Empty Client)",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: 16, server: "example.com", client: "", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Empty Map Name)",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: 16, server: "example.com", client: "player1", mapName: "", playbackTime: time.Second * 60, filesize: 1024},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Negative Filesize)",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: 16, server: "example.com", client: "player1", mapName: "de_dust2", playbackTime: time.Second * 60, filesize: -1024},
			wantErr: true,
		},
		{
			name:    "Invalid Header (Zero Playback Time)",
			fields:  fields{playbackTicks: 128, playbackFrames: 256, signonLength: 16, server: "example.com", client: "player1", mapName: "de_dust2", playbackTime: 0, filesize: 1024},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &demoHeader{
				playbackTicks:  tt.fields.playbackTicks,
				playbackFrames: tt.fields.playbackFrames,
				signonLength:   tt.fields.signonLength,
				server:         tt.fields.server,
				client:         tt.fields.client,
				mapName:        tt.fields.mapName,
				playbackTime:   tt.fields.playbackTime,
				filesize:       tt.fields.filesize,
			}
			err := h.validate()
			assert.Equal(t, tt.wantErr, (err != nil))
		})
	}
}

func Test_demoHeader_uuid(t *testing.T) {
	type fields struct {
		server         string
		client         string
		mapName        string
		playbackTicks  int
		playbackFrames int
		signonLength   int
		playbackTime   time.Duration
		filesize       int64
	}
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "All fields have non-zero values",
			fields: fields{
				server:         "example-server",
				client:         "example-client",
				mapName:        "example-map",
				playbackTicks:  100,
				playbackFrames: 200,
				signonLength:   30,
				playbackTime:   5 * time.Minute,
				filesize:       1024,
			},
			want: uuid.NewMD5(uuid.UUID{}, []byte("100,200,30,example-server,example-client,example-map,300000000000,1024")),
		},
		{
			name: "Some fields have zero values",
			fields: fields{
				server:         "example-server",
				client:         "example-client",
				mapName:        "example-map",
				playbackTicks:  0,
				playbackFrames: 0,
				signonLength:   0,
				playbackTime:   0,
				filesize:       0,
			},
			want: uuid.NewMD5(uuid.UUID{}, []byte("0,0,0,example-server,example-client,example-map,0,0")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &demoHeader{
				server:         tt.fields.server,
				client:         tt.fields.client,
				mapName:        tt.fields.mapName,
				playbackTicks:  tt.fields.playbackTicks,
				playbackFrames: tt.fields.playbackFrames,
				signonLength:   tt.fields.signonLength,
				playbackTime:   tt.fields.playbackTime,
				filesize:       tt.fields.filesize,
			}
			got := h.uuid()
			assert.Equal(t, tt.want, got)
		})
	}
}