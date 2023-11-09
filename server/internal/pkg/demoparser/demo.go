package demoparser

import (
	"errors"
	"io"
	"mime/multipart"
	"strings"
)

type Demo struct {
	io.ReadCloser
	size int64
}

func NewDemo(rc io.ReadCloser, h *multipart.FileHeader) (Demo, error) {
	if rc == nil {
		return Demo{}, errors.New("nil demo file")
	}

	if h == nil {
		return Demo{}, errors.New("nil file header")
	}

	if h.Size <= 0 {
		return Demo{}, errors.New("file header size must be greater than 0")
	}

	ss := strings.Split(h.Filename, ".")
	if len(ss)-1 <= 0 || ss[len(ss)-1] != "dem" {
		return Demo{}, errors.New("demo must have .dem file extension")
	}

	return Demo{rc, h.Size}, nil
}
