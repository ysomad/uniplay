package demoparser

import (
	"errors"
	"io"
	"mime/multipart"
	"strings"
)

type demo struct {
	io.Reader
	size int64
}

func newDemo(rc io.Reader, h *multipart.FileHeader) (demo, error) {
	if rc == nil {
		return demo{}, errors.New("nil demo file")
	}

	if h == nil {
		return demo{}, errors.New("nil file header")
	}

	if h.Size <= 0 {
		return demo{}, errors.New("file header size must be greater than 0")
	}

	ss := strings.Split(h.Filename, ".")
	if len(ss)-1 <= 0 || ss[len(ss)-1] != "dem" {
		return demo{}, errors.New("demo must have .dem file extension")
	}

	return demo{rc, h.Size}, nil
}
