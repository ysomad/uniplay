package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

var _ paging.ListItem = &PlayerMatch{}

// PlayerMatch implements paging.ItemList interface so it can be used in paginated lists.
type PlayerMatch struct {
	ID         uuid.UUID
	Map        Map
	Score      MatchScore
	Stats      PlayerMatchStats
	State      MatchState
	UploadedAt time.Time
}

func (m *PlayerMatch) GetID() string      { return m.ID.String() }
func (m *PlayerMatch) GetTime() time.Time { return m.UploadedAt }

type PlayerMatchStats struct {
	Kills              int32
	Deaths             int32
	Assists            int32
	HeadshotPercentage float64
	ADR                float64
}
