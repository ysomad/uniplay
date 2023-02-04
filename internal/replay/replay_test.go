package replay

import (
	"io"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newReplay(t *testing.T) {
	type args struct {
		rc io.ReadCloser
		fh *multipart.FileHeader
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
				rc: nil,
			},
			want: replay{
				ReadCloser: nil,
				size:       0,
			},
			wantErr: true,
		},
		{
			name: "empty file header",
			args: args{
				rc: io.NopCloser(strings.NewReader("TEST")),
				fh: nil,
			},
			want: replay{
				ReadCloser: nil,
			},
			wantErr: true,
		},
		{
			name: "invalid replay filename",
			args: args{
				rc: io.NopCloser(strings.NewReader("TEST")),
				fh: &multipart.FileHeader{Filename: "invalid", Size: 555},
			},
			want: replay{
				ReadCloser: nil,
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				rc: io.NopCloser(strings.NewReader("TEST")),
				fh: &multipart.FileHeader{Filename: "test.dem", Size: 5555},
			},
			want: replay{
				ReadCloser: io.NopCloser(strings.NewReader("TEST")),
				size:       5555,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newReplay(tt.args.rc, tt.args.fh)

			assert.Equal(t, tt.wantErr, (err != nil))
			assert.ObjectsAreEqual(tt.want, got)
		})
	}
}
