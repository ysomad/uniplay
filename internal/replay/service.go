package replay

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/uniplay/internal/domain"
)

type Service struct {
	tracer trace.Tracer
	replay replayRepository
}

func NewService(t trace.Tracer, r replayRepository) *Service {
	return &Service{
		tracer: t,
		replay: r,
	}
}

func (s *Service) CollectStats(ctx context.Context, r replay) (matchID uuid.UUID, err error) {
	ctx, span := s.tracer.Start(ctx, "replay.Service.CollectStats")
	defer span.End()

	p, err := newParser(r)
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
