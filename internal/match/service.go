package match

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	tracer trace.Tracer
	match  matchRepository
}

func NewService(t trace.Tracer, mr matchRepository) *Service {
	return &Service{
		tracer: t,
		match:  mr,
	}
}

// Delete deletes match and all stats associated with it, including player match history.
func (s *Service) Delete(ctx context.Context, matchID uuid.UUID) error {
	return s.match.DeleteByID(ctx, matchID)
}
