package replay

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newReplay(t *testing.T) {
	type args struct {
		rc       io.ReadCloser
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    replay
		wantErr bool
	}{
		{
			name: "empty read closer",
			args: args{
				rc:       nil,
				filename: "",
			},
			want: replay{
				ReadCloser: nil,
			},
			wantErr: true,
		},
		{
			name: "empty filename",
			args: args{
				rc:       io.NopCloser(strings.NewReader("TEST")),
				filename: "",
			},
			want: replay{
				ReadCloser: nil,
			},
			wantErr: true,
		},
		{
			name: "invalid filename",
			args: args{
				rc:       io.NopCloser(strings.NewReader("TEST")),
				filename: "invalid.filename",
			},
			want: replay{
				ReadCloser: nil,
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				rc:       io.NopCloser(strings.NewReader("TEST")),
				filename: "test.dem",
			},
			want: replay{
				ReadCloser: io.NopCloser(strings.NewReader("TEST")),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newReplay(tt.args.rc, tt.args.filename)

			assert.Equal(t, tt.wantErr, (err != nil))
			assert.ObjectsAreEqual(tt.want, got)
		})
	}
}
