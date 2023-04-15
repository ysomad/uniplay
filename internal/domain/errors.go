package domain

import "errors"

/*
`Match` - >= 800
`Player` - >= 700
`Team` - >= 900
`WeaponStats` - >= 1000
*/

const (
	CodePlayerNotFound    = 700
	CodeMatchAlreadyExist = 800
	CodeMatchNotFound     = 801
)

var (
	ErrMatchAlreadyExist = errors.New("match already exist")
	ErrMatchNotFound     = errors.New("match not found")
	ErrPlayerNotFound    = errors.New("player not found")
)
