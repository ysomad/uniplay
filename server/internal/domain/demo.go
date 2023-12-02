package domain

import (
	"time"

	"github.com/google/uuid"
)

type DemoStatus string

const (
	DemoStatusUploaded   DemoStatus = "uploaded"
	DemoStatusProcessing DemoStatus = "processing"
	DemoStatusDone       DemoStatus = "done"
	DemoStatusError      DemoStatus = "error"
)

type Demo struct {
	UploadedAt  time.Time
	ProcessedAt time.Time
	Status      DemoStatus
	Reason      string
	Uploader    string
	ID          uuid.UUID
}
