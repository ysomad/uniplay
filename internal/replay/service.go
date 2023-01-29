package replay

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/ysomad/uniplay/internal/domain"
)

type service struct {
	log    *zap.Logger
	tracer trace.Tracer
	replay replayRepository
}

func NewService(l *zap.Logger, t trace.Tracer, r replayRepository) *service {
	return &service{
		log:    l,
		tracer: t,
		replay: r,
	}
}

func (s *service) CollectStats(ctx context.Context, r replay) (matchID uuid.UUID, err error) {
	ctx, span := s.tracer.Start(ctx, "replay.service.CollectStats")
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

	match, playerStats, weaponStats, err := p.collectStats()
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := s.replay.SaveStats(ctx, match, playerStats, weaponStats); err != nil {
		return uuid.UUID{}, err
	}

	return matchID, nil
}
