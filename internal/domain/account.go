package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Verified  bool
	CreatedAt time.Time
}

func NewAccount(email, hashedPassword string) (a Account, err error) {
	a.ID, err = uuid.NewRandom()
	if err != nil {
		return Account{}, err
	}

	a.Email = email
	a.Password = hashedPassword
	a.CreatedAt = time.Now()

	return a, nil
}
