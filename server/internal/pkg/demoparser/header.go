package demoparser

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type demoHeader struct {
	server         string
	client         string
	mapName        string
	playbackTicks  int
	playbackFrames int
	signonLength   int
	playbackTime   time.Duration
	filesize       int64
}

func (h *demoHeader) validate() error {
	if h.playbackTicks <= 0 {
		return errors.New("playback ticks must be greater than 0")
	}

	if h.playbackFrames <= 0 {
		return errors.New("playback frames must be greater than 0")
	}

	if h.signonLength <= 0 {
		return errors.New("signon length must be greater than 0")
	}

	if h.server == "" {
		return errors.New("server must not be empty")
	}

	if h.client == "" {
		return errors.New("client must not be empty")
	}

	if h.mapName == "" {
		return errors.New("map name must not be empty")
	}

	if h.playbackTime <= 0 {
		return errors.New("playback time must be greater than 0")
	}

	if h.filesize <= 0 {
		return errors.New("demo file size must be greater than 0")
	}

	return nil
}

// uuid generates md5 uuid from all fields of demo header.
func (h *demoHeader) uuid() uuid.UUID {
	s := fmt.Sprintf(
		"%d,%d,%d,%s,%s,%s,%d,%d",
		h.playbackTicks,
		h.playbackFrames,
		h.signonLength,
		h.server,
		h.client,
		h.mapName,
		h.playbackTime,
		h.filesize)
	return uuid.NewMD5(uuid.UUID{}, []byte(s))
}
