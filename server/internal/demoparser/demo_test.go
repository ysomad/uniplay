package demoparser

import (
	"errors"
	"io"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testReadSeeker struct {
	data   []byte
	offset int64
}

func newTestReadSeeker(data []byte) *testReadSeeker {
	return &testReadSeeker{data: data}
}

func (d *testReadSeeker) Read(p []byte) (n int, err error) {
	if d.offset >= int64(len(d.data)) {
		return 0, io.EOF
	}

	n = copy(p, d.data[d.offset:])
	d.offset += int64(n)

	return n, nil
}

func (d *testReadSeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		d.offset = offset
	case io.SeekCurrent:
		d.offset += offset
	case io.SeekEnd:
		d.offset = int64(len(d.data)) + offset
	default:
		return 0, errors.New("invalid whence")
	}

	if d.offset < 0 {
		d.offset = 0
	} else if d.offset > int64(len(d.data)) {
		d.offset = int64(len(d.data))
	}

	return d.offset, nil
}

func TestNewDemo(t *testing.T) {
	t.Parallel()
	type args struct {
		rc io.ReadSeeker
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
				rc: newTestReadSeeker([]byte("urmom")),
				h:  &multipart.FileHeader{},
			},
			want:    Demo{},
			wantErr: true,
		},
		{
			name: "0 file header size",
			args: args{
				rc: newTestReadSeeker([]byte("deez nuts")),
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
				rc: newTestReadSeeker([]byte("deez nuts")),
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
				rc: newTestReadSeeker([]byte("deez nuts")),
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
				rc: newTestReadSeeker([]byte("yoink")),
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
				rc: newTestReadSeeker([]byte("yoink")),
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
				rc: newTestReadSeeker([]byte("yoink")),
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
				rc: newTestReadSeeker([]byte("success")),
				h: &multipart.FileHeader{
					Filename: "test.dem",
					Size:     10000,
				},
			},
			want: Demo{
				Reader: newTestReadSeeker([]byte("success")),
				ID:     "ec5b2df0-efa4-339f-bebb-f05273ccbf3a",
				Size:   10000,
			},
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
