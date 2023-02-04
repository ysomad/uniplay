package match

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/google/uuid"
	"github.com/ysomad/uniplay/internal/domain"
)

type Service struct {
	tracer trace.Tracer
	match  matchRepository
}

func NewService(t trace.Tracer, m matchRepository) *Service {
	return &Service{
		tracer: t,
		match:  m,
	}
}

type collectStatsRes struct {
	MatchID     uuid.UUID
	MatchNumber int32
}

func (s *Service) CreateFromReplay(ctx context.Context, r replay) (collectStatsRes, error) {
	ctx, span := s.tracer.Start(ctx, "match.Service.CreateFromReplay")
	defer span.End()

	p := newParser(r)
	defer p.close()

	if _, err := p.parseReplayHeader(); err != nil {
		return collectStatsRes{}, err
	}

	matchExists, err := s.match.Exists(ctx, p.match.id)
	if err != nil {
		return collectStatsRes{}, err
	}

	if matchExists {
		return collectStatsRes{}, domain.ErrMatchAlreadyExist
	}

	match, playerStats, weaponStats, err := p.collectStats(ctx)
	if err != nil {
		return collectStatsRes{}, err
	}

	matchNum, err := s.match.CreateWithStats(ctx, match, playerStats, weaponStats)
	if err != nil {
		return collectStatsRes{}, err
	}

	return collectStatsRes{
		MatchID:     p.match.id,
		MatchNumber: matchNum,
	}, nil
}

// DeleteByID deletes match and all stats associated with it, including player match history.
func (s *Service) DeleteByID(ctx context.Context, matchID uuid.UUID) error {
	return s.match.DeleteByID(ctx, matchID)
}
