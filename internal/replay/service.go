package replay

import (
	"context"
	"time"

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

func (s *Service) CollectStats(ctx context.Context, r replay) (matchNumber int32, err error) {
	ctx, span := s.tracer.Start(ctx, "replay.Service.CollectStats")
	defer span.End()

	p := newParser(r)
	defer p.close()

	h, err := p.parseReplayHeader()
	if err != nil {
		return 0, err
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
		return 0, err
	}

	p.match.uploadedAt = time.Now()

	matchExists, err := s.replay.MatchExists(ctx, p.match.id)
	if err != nil {
		return 0, err
	}

	if matchExists {
		return 0, domain.ErrMatchAlreadyExist
	}

	match, playerStats, weaponStats, err := p.collectStats(ctx)
	if err != nil {
		return 0, err
	}

	matchNumber, err = s.replay.SaveStats(ctx, match, playerStats, weaponStats)
	if err != nil {
		return 0, err
	}

	return matchNumber, nil
}
