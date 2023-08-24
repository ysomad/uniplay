package match

import (
	"context"

	"github.com/google/uuid"
	"github.com/ysomad/uniplay/internal/domain"
)

type repository interface {
	CreateWithStats(context.Context, *replayMatch, []*playerStat, []*weaponStat) error
	Exists(ctx context.Context, matchID uuid.UUID) (found bool, err error)
	DeleteByID(ctx context.Context, matchID uuid.UUID) error
	FindByID(ctx context.Context, matchID uuid.UUID) (domain.Match, error)
}

type service struct {
	match repository
}

func NewService(m repository) *service {
	return &service{
		match: m,
	}
}

func (s *service) CreateFromReplay(ctx context.Context, r replay) (uuid.UUID, error) {
	p := newParser(r)
	defer p.close()

	if err := p.parseReplayHeader(); err != nil {
		return uuid.Nil, err
	}

	matchExists, err := s.match.Exists(ctx, p.match.id)
	if err != nil {
		return uuid.Nil, err
	}

	if matchExists {
		return uuid.Nil, domain.ErrMatchAlreadyExist
	}

	match, playerStats, weaponStats, err := p.collectStats(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	match.score = domain.NewMatchScore(match.team1.score, match.team2.score)
	match.roundsTotal = match.team1.score + match.team2.score

	if err := s.match.CreateWithStats(ctx, match, playerStats, weaponStats); err != nil {
		return uuid.Nil, err
	}

	return p.match.id, nil
}

// DeleteByID deletes match and all stats associated with it, including player match history.
func (s *service) DeleteByID(ctx context.Context, matchID uuid.UUID) error {
	return s.match.DeleteByID(ctx, matchID)
}

func (s *service) GetByID(ctx context.Context, matchID uuid.UUID) (domain.Match, error) {
	m, err := s.match.FindByID(ctx, matchID)
	if err != nil {
		return domain.Match{}, err
	}

	return m, nil
}