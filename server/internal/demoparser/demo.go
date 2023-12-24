package demoparser

import (
	"crypto/md5" //nolint:gosec // its weak but ok for the project
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
)

type Demo struct {
	io.Reader
	ID   uuid.UUID
	Size int64
}

func NewDemo(file io.ReadSeeker, header *multipart.FileHeader) (Demo, error) {
	if file == nil {
		return Demo{}, errors.New("nil demo file")
	}

	if header == nil {
		return Demo{}, errors.New("nil file header")
	}

	if header.Size <= 0 {
		return Demo{}, errors.New("file header size must be greater than 0")
	}

	ss := strings.Split(header.Filename, ".")
	if len(ss)-1 <= 0 || ss[len(ss)-1] != "dem" {
		return Demo{}, errors.New("demo must have .dem file extension")
	}

	hash := md5.New() //nolint:gosec // using weak primitive lmao

	if _, err := io.Copy(hash, file); err != nil {
		return Demo{}, err
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return Demo{}, err
	}

	return Demo{
		Reader: file,
		ID:     uuid.NewMD5(uuid.Nil, hash.Sum(nil)),
		Size:   header.Size,
	}, nil
}

func (d Demo) Filename() string {
	return fmt.Sprintf("%s.dem", d.ID)
}
