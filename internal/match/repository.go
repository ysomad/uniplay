package match

import (
	"context"

	"github.com/google/uuid"
)

type matchRepository interface {
	DeleteByID(ctx context.Context, matchID uuid.UUID) error
}
