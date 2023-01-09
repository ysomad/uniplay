package replay

import (
	"context"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type service struct {
	log    *zap.Logger
	replay repository
}

func NewService(l *zap.Logger, r repository) *service {
	return &service{
		log:    l,
		replay: r,
	}
}

func (s *service) CollectStats(ctx context.Context, r Replay) (*domain.Match, error) {
	p, err := newParser(r, s.log)
	if err != nil {
		return nil, err
	}
	defer p.close()

	matchID, err := p.parseReplayHeader()
	if err != nil {
		return nil, err
	}

	matchExists, err := s.replay.MatchExists(ctx, matchID)
	if err != nil {
		return nil, err
	}

	if matchExists {
		return nil, domain.ErrMatchAlreadyExist
	}

	match, playerStats, weaponStats, err := p.collectStats()
	if err != nil {
		return nil, err
	}

	return s.replay.SaveStats(ctx, match, playerStats, weaponStats)
}
