package demoparser

import (
	"crypto/md5"
	"errors"
	"io"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
)

type demo struct {
	io.Reader
	id   uuid.UUID
	size int64
}

func newDemo(file io.ReadSeeker, header *multipart.FileHeader) (demo, error) {
	if file == nil {
		return demo{}, errors.New("nil demo file")
	}

	if header == nil {
		return demo{}, errors.New("nil file header")
	}

	if header.Size <= 0 {
		return demo{}, errors.New("file header size must be greater than 0")
	}

	ss := strings.Split(header.Filename, ".")
	if len(ss)-1 <= 0 || ss[len(ss)-1] != "dem" {
		return demo{}, errors.New("demo must have .dem file extension")
	}

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return demo{}, err
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return demo{}, err
	}

	return demo{
		Reader: file,
		id:     uuid.NewMD5(uuid.Nil, hash.Sum(nil)),
		size:   header.Size,
	}, nil
}
