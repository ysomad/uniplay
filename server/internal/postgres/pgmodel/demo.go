package pgmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"

	"github.com/ysomad/uniplay/server/internal/domain"
)

type Demo struct {
	UploadedAt  time.Time            `db:"uploaded_at"`
	ProcessedAt zeronull.Timestamptz `db:"processed_at"`
	Status      domain.DemoStatus    `db:"status"`
	Reason      zeronull.Text        `db:"reason"`
	IdentityID  string               `db:"identity_id"` // uploader of a demo
	ID          uuid.UUID            `db:"id"`
}
