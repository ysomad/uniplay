package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/ysomad/uniplay/internal/pkg/argon2"
)

type Account struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Verified  bool
	CreatedAt time.Time
}

func NewAccount(email, password string) (a *Account, err error) {
	a = &Account{}

	a.Password, err = argon2.GenerateFromPassword(password, argon2.DefaultParams)
	if err != nil {
		return nil, err
	}

	a.ID, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	a.Email = email
	a.CreatedAt = time.Now()

	return a, nil
}
