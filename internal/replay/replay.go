package replay

import (
	"errors"
	"io"
	"strings"
)

var (
	ErrInvalidReplayFileExtension = errors.New("replay must have .dem file extension")
	ErrEmptyReplay                = errors.New("replay is empty")
)

type Replay struct {
	io.ReadCloser
}

func New(rc io.ReadCloser, filename string) (Replay, error) {
	if rc == nil {
		return Replay{}, ErrEmptyReplay
	}

	sub := strings.Split(filename, ".")
	lastIndex := len(sub) - 1

	if lastIndex <= 0 {
		return Replay{}, ErrInvalidReplayFileExtension
	}

	if sub[lastIndex] != "dem" {
		return Replay{}, ErrInvalidReplayFileExtension
	}

	return Replay{rc}, nil
}
