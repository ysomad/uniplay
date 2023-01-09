package replay

import (
	"context"
	"os"

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

func (s *service) CollectStats(ctx context.Context, filename string) (*domain.Match, error) {
	replay, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	parser, err := newParser(replay, s.log)
	if err != nil {
		return nil, err
	}
	defer parser.close()

	matchID, err := parser.parseReplayHeader()
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

	match, playerStats, weaponStats, err := parser.collectStats()
	if err != nil {
		return nil, err
	}

	return s.replay.SaveStats(ctx, match, playerStats, weaponStats)
}
