package replay

import (
	"errors"
	"io"
	"strings"
)

var (
	errInvalidReplayFileExtension = errors.New("replay must have .dem file extension")
	errEmptyReplay                = errors.New("replay is empty")
)

type Replay struct {
	io.ReadCloser
}

func New(rc io.ReadCloser, filename string) (Replay, error) {
	if rc == nil {
		return Replay{}, errEmptyReplay
	}

	sub := strings.Split(filename, ".")
	lastIndex := len(sub) - 1

	if lastIndex <= 0 {
		return Replay{}, errInvalidReplayFileExtension
	}

	if sub[lastIndex] != "dem" {
		return Replay{}, errInvalidReplayFileExtension
	}

	return Replay{rc}, nil
}
