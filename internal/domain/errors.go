package domain

import "github.com/ssssargsian/uniplay/internal/pkg/apperror"

var (
	ErrReplayTeamsNotSaved = apperror.New(500, "parsed teams from replay not saved")
)

var (
	ErrMatchAlreadyExist = apperror.New(600, "match already exist")
)

var (
	ErrPlayerNotFound = apperror.New(800, "player not found")
)

var (
	ErrWeaponStatsNotFound      = apperror.New(1000, "player has no weapon stats")
	ErrWeaponClassStatsNotFound = apperror.New(1001, "player has no weapon class stats")
)
