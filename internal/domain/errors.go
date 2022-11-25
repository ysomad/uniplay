package domain

import "github.com/ssssargsian/uniplay/internal/pkg/apperror"

var (
	ErrMatchAlreadyExist = apperror.New(600, "match already exist")
)

var (
	ErrWeaponStatsNotFound      = apperror.New(1000, "player has no weapon stats")
	ErrWeaponClassStatsNotFound = apperror.New(1001, "player has no weapon class stats")
)
