package replay

import (
	"context"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/ssssargsian/uniplay/internal/domain"
)

type service struct {
	log    *zap.Logger
	replay replayRepository
}

func NewService(l *zap.Logger, r replayRepository) *service {
	return &service{
		log:    l,
		replay: r,
	}
}

func (s *service) CollectStats(ctx context.Context, r replay) (matchID uuid.UUID, err error) {
	p, err := newParser(r, s.log)
	if err != nil {
		return uuid.UUID{}, err
	}

	defer p.close()

	matchID, err = p.parseReplayHeader()
	if err != nil {
		return uuid.UUID{}, err
	}

	matchExists, err := s.replay.MatchExists(ctx, matchID)
	if err != nil {
		return uuid.UUID{}, err
	}

	if matchExists {
		return uuid.UUID{}, domain.ErrMatchAlreadyExist
	}

	match, playerStats, weaponStats, err := p.collectStats()
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := s.replay.SaveStats(ctx, match, playerStats, weaponStats); err != nil {
		return uuid.UUID{}, err
	}

	return matchID, nil
}
