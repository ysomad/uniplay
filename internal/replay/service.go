package replay

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/google/uuid"
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

type collectStatsRes struct {
	MatchID     uuid.UUID
	MatchNumber int32
}

func (s *Service) CollectStats(ctx context.Context, r replay) (collectStatsRes, error) {
	ctx, span := s.tracer.Start(ctx, "replay.Service.CollectStats")
	defer span.End()

	p := newParser(r)
	defer p.close()

	h, err := p.parseReplayHeader()
	if err != nil {
		return collectStatsRes{}, err
	}

	p.match.id, err = domain.NewMatchID(
		h.ServerName,
		h.ClientName,
		h.MapName,
		h.PlaybackTime,
		h.PlaybackTicks,
		h.PlaybackFrames,
		h.SignonLength,
	)
	if err != nil {
		return collectStatsRes{}, err
	}

	p.match.uploadedAt = time.Now()

	matchExists, err := s.replay.MatchExists(ctx, p.match.id)
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

	matchNum, err := s.replay.SaveStats(ctx, match, playerStats, weaponStats)
	if err != nil {
		return collectStatsRes{}, err
	}

	return collectStatsRes{
		MatchID:     p.match.id,
		MatchNumber: matchNum,
	}, nil
}
