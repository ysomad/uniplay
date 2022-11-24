package domain

import (
	"time"

	"github.com/ssssargsian/uniplay/internal/pkg/apperror"
)

var (
	ErrPlayerNotFound = apperror.New(800, "player not found")
)

type Player struct {
	SteamID      uint64
	TeamName     string
	TeamFlagCode string
	CreateTime   time.Time
	UpdateTime   time.Time
}
