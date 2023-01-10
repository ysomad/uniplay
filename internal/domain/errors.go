package domain

import "github.com/ssssargsian/uniplay/internal/pkg/apperror"

/*
`Match` - >= 600
`Metric` - >= 700
`Player` - >= 800
`Team` - >= 900
`WeaponStats` - >= 1000
*/

var (
	ErrMatchAlreadyExist = apperror.New(600, "match already exist")
)

var (
	ErrPlayerNotFound = apperror.New(800, "player not found")
)
