package match

import (
	"errors"
	"io"
	"mime/multipart"
	"strings"
)

var (
	errInvalidReplayFileExtension = errors.New("replay must have .dem file extension")
	errInvalidReplayFileSize      = errors.New("invalid replay file size")
	errEmptyReplay                = errors.New("replay is empty")
	errEmptyFileHeader            = errors.New("empty file header")
)

type replay struct {
	io.ReadCloser
	size int64
}

func newReplay(rc io.ReadCloser, h *multipart.FileHeader) (replay, error) {
	if rc == nil {
		return replay{}, errEmptyReplay
	}

	if h == nil {
		return replay{}, errEmptyFileHeader
	}

	if h.Size <= 0 {
		return replay{}, errInvalidReplayFileSize
	}

	sub := strings.Split(h.Filename, ".")
	lastIndex := len(sub) - 1

	if lastIndex <= 0 {
		return replay{}, errInvalidReplayFileExtension
	}

	if sub[lastIndex] != "dem" {
		return replay{}, errInvalidReplayFileExtension
	}

	return replay{rc, h.Size}, nil
}
