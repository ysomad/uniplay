package domain

import (
	"time"
)

type Player struct {
	SteamID      uint64
	TeamName     string
	TeamFlagCode string
	CreateTime   time.Time
	UpdateTime   time.Time
}
