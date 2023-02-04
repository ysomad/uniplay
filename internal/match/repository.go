package match

import (
	"context"

	"github.com/google/uuid"
)

type matchRepository interface {
	CreateWithStats(context.Context, *replayMatch, []*playerStat, []*weaponStat) (matchNumber int32, err error)
	Exists(ctx context.Context, matchID uuid.UUID) (found bool, err error)
	DeleteByID(ctx context.Context, matchID uuid.UUID) error
}
