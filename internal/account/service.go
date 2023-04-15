package account

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type repository interface {
	Save(context.Context, domain.Account) error
}

type hashVerifier interface {
	Hash(plain string) (string, error)
	Verify(plain, hash string) (bool, error)
}

type service struct {
	account  repository
	password hashVerifier
}

func NewService(r repository, h hashVerifier) *service {
	return &service{
		account:  r,
		password: h,
	}
}

func (s *service) Create(ctx context.Context, email, password string) (domain.Account, error) {
	hash, err := s.password.Hash(password)
	if err != nil {
		return domain.Account{}, err
	}

	acc, err := domain.NewAccount(email, hash)
	if err != nil {
		return domain.Account{}, err
	}

	if err = s.account.Save(ctx, acc); err != nil {
		return domain.Account{}, err
	}

	return acc, nil
}
