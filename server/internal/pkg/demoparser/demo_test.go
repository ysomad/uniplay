package demoparser

import (
	"io"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newReadCloser(data string) io.ReadCloser {
	reader := strings.NewReader(data)
	readCloser := io.NopCloser(reader)
	return readCloser
}

func TestNewDemo(t *testing.T) {
	t.Parallel()
	type args struct {
		rc io.ReadCloser
		h  *multipart.FileHeader
	}
	tests := []struct {
		name    string
		args    args
		want    Demo
		wantErr bool
	}{
		{
			name: "empty demo file",
			args: args{
				rc: nil,
				h: &multipart.FileHeader{
					Filename: "test.dem",
					Size:     1337,
				},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "empty file header",
			args: args{
				rc: newReadCloser("urmom"),
				h:  &multipart.FileHeader{},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "0 file header size",
			args: args{
				rc: newReadCloser("deez nuts"),
				h: &multipart.FileHeader{
					Filename: "test.dem",
					Size:     0,
				},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "negative file header size",
			args: args{
				rc: newReadCloser("deez nuts"),
				h: &multipart.FileHeader{
					Filename: "test.dem",
					Size:     -1488,
				},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "filename without .",
			args: args{
				rc: newReadCloser("deez nuts"),
				h: &multipart.FileHeader{
					Filename: "urmomspaghetti",
					Size:     322,
				},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "filename with multiple .",
			args: args{
				rc: newReadCloser("yoink"),
				h: &multipart.FileHeader{
					Filename: "ur.mom.ru",
					Size:     322,
				},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "filename with a lot of dots",
			args: args{
				rc: newReadCloser("yoink"),
				h: &multipart.FileHeader{
					Filename: "...ur...mom.ru...",
					Size:     322,
				},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "filename with invalid file extension",
			args: args{
				rc: newReadCloser("yoink"),
				h: &multipart.FileHeader{
					Filename: "dick.head",
					Size:     666,
				},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				rc: newReadCloser("success"),
				h: &multipart.FileHeader{
					Filename: "test.dem",
					Size:     10000,
				},
			},
			want:    Demo{newReadCloser("success"), 10000},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDemo(tt.args.rc, tt.args.h)
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.want, got)
		})
	}
}
