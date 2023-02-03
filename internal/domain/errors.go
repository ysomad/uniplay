package domain

import "errors"

/*
`Match` - >= 600
`Metric` - >= 700
`Player` - >= 800
`Team` - >= 900
`WeaponStats` - >= 1000
*/

const (
	CodeMatchAlreadyExist = 600
	CodeMatchNotFound     = 601
	CodePlayerNotFound    = 800
)

var (
	ErrMatchAlreadyExist = errors.New("match already exist")
	ErrMatchNotFound     = errors.New("match not found")
	ErrPlayerNotFound    = errors.New("player not found")
)
