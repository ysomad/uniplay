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

type Demo struct {
	UploadedAt  time.Time
	ProcessedAt time.Time
	Status      DemoStatus
	Reason      string
	IdentityID  string // uploader of a demo
	ID          uuid.UUID
}
