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

type replay struct {
	io.ReadCloser
}

func newReplay(rc io.ReadCloser, filename string) (replay, error) {
	if rc == nil {
		return replay{}, errEmptyReplay
	}

	sub := strings.Split(filename, ".")
	lastIndex := len(sub) - 1

	if lastIndex <= 0 {
		return replay{}, errInvalidReplayFileExtension
	}

	if sub[lastIndex] != "dem" {
		return replay{}, errInvalidReplayFileExtension
	}

	return replay{rc}, nil
}
