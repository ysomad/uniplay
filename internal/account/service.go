package account

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type repository interface {
	Save(context.Context, *domain.Account) error
}

type service struct {
	account repository
}

func NewService(r repository) *service {
	return &service{
		account: r,
	}
}

func (s *service) Create(ctx context.Context, email, password string) (*domain.Account, error) {
	acc, err := domain.NewAccount(email, password)
	if err != nil {
		return nil, err
	}

	if err := s.account.Save(ctx, acc); err != nil {
		return nil, err
	}

	return acc, nil
}
