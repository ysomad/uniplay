package replay

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/otel"
)

type Service struct {
	log    *zap.Logger
	replay replayRepository
}

func NewService(l *zap.Logger, r replayRepository) *Service {
	return &Service{
		log:    l,
		replay: r,
	}
}

func (s *Service) CollectStats(ctx context.Context, r replay) (matchID uuid.UUID, err error) {
	_, span := otel.StartTrace(ctx, libraryName, "replay.Service.CollectStats")
	defer span.End()

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

	match, playerStats, weaponStats, err := p.collectStats(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := s.replay.SaveStats(ctx, match, playerStats, weaponStats); err != nil {
		return uuid.UUID{}, err
	}

	return matchID, nil
}
