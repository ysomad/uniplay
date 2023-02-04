package domain

import (
	"errors"
	"fmt"
	"time"
)

var (
	errInvalidServerName     = errors.New("invalid server name")
	errInvalidClientName     = errors.New("invalid client name")
	errInvalidMapName        = errors.New("invalid map name")
	errInvalidPlaybackTime   = fmt.Errorf("match must last more than %s", minMatchDuration.String())
	errInvalidPlaybackTicks  = errors.New("invalid amount of playback ticks")
	errInvalidPlaybackFrames = errors.New("invalid amount of playback frames")
	errInvalidSignonLength   = errors.New("invalid signon length")
)

// ReplayHeader is a replay header parsed from uploaded replay (aka demo).
type ReplayHeader struct {
	playbackTicks  int           // game duration in ticks
	playbackFrames int           // amount of frames aka demo-ticks recorded
	signonLength   int           // length of the Signon package in bytes
	server         string        // servers hostname config value
	client         string        // usually 'GOTV Demo'
	mapName        string        // de_cache, de_nuke, etc.
	playbackTime   time.Duration // replay duration in seconds
	filesize       int64         // size of replay file
}

func NewReplayHeader(
	ticks, frames, signonLen int,
	server, client, mapName string,
	playbackTime time.Duration,
	filesize int64,
) (*ReplayHeader, error) {
	if ticks <= 0 {
		return nil, errInvalidPlaybackTicks
	}

	if frames <= 0 {
		return nil, errInvalidPlaybackFrames
	}

	if signonLen <= 0 {
		return nil, errInvalidSignonLength
	}

	if server == "" {
		return nil, errInvalidServerName
	}

	if client == "" {
		return nil, errInvalidClientName
	}

	if mapName == "" {
		return nil, errInvalidMapName
	}

	if playbackTime < minMatchDuration {
		return nil, errInvalidPlaybackTime
	}

	return &ReplayHeader{
		playbackTicks:  ticks,
		playbackFrames: frames,
		signonLength:   signonLen,
		server:         server,
		client:         client,
		mapName:        mapName,
		playbackTime:   playbackTime,
		filesize:       filesize,
	}, nil
}
