package domain

import (
	"time"

	"github.com/google/uuid"
)

type DemoStatus string

const (
	DemoStatusAwaiting  DemoStatus = "AWAITING"
	DemoStatusProcessed DemoStatus = "PROCESSED"
	DemoStatusError     DemoStatus = "ERROR"
)

func (s DemoStatus) Valid() bool {
	switch s {
	case DemoStatusAwaiting, DemoStatusProcessed, DemoStatusError:
		return true
	default:
		return false
	}
}

type Demo struct {
	UploadedAt  time.Time
	ProcessedAt time.Time
	Status      DemoStatus
	Reason      string
	IdentityID  string // uploader of a demo
	ID          uuid.UUID
}
