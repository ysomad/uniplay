package domain

import "github.com/google/uuid"

func NewPlayerID() (uuid.UUID, error) { return uuid.NewRandom() }
