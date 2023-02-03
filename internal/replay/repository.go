package replay

import (
	"context"

	"github.com/google/uuid"
)

type replayRepository interface {
	SaveStats(context.Context, *replayMatch, []*playerStat, []*weaponStat) (matchNumber int32, err error)
	MatchExists(ctx context.Context, matchID uuid.UUID) (found bool, err error)
}
