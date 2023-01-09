package replay

import (
	"context"

	"github.com/google/uuid"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type replayRepository interface {
	SaveStats(context.Context, *replayMatch, []playerStat, []weaponStat) (*domain.Match, error)
	MatchExists(ctx context.Context, matchID uuid.UUID) (found bool, err error)
}
