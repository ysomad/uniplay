package domain

import (
	"time"

	"github.com/ysomad/uniplay/internal/pkg/argon2"

	"github.com/google/uuid"
)

type Account struct {
	ID         uuid.UUID
	Email      string
	Password   string
	IsVerified bool
	IsAdmin    bool
	CreatedAt  time.Time
}

func NewAccount(email, password string, created time.Time) (a *Account, err error) {
	a = new(Account)

	a.Password, err = argon2.GenerateFromPassword(password, argon2.DefaultParams)
	if err != nil {
		return nil, err
	}

	a.ID = uuid.New()
	a.Email = email
	a.CreatedAt = created

	return a, nil
}
