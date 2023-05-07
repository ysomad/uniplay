package domain

import "errors"

/*
`Match` - >= 800
`Player` - >= 700
`Team` - >= 900
Account - >= 1000
*/

const (
	CodePlayerNotFound = 700

	CodeMatchNotFound     = 800
	CodeMatchAlreadyExist = 801

	CodeTeamNotFound      = 900
	CodeTeamClanNameTaken = 901

	CodeAccountEmailTaken = 1000
)

var (
	ErrPlayerNotFound = errors.New("player not found")

	ErrMatchAlreadyExist = errors.New("match already exist")
	ErrMatchNotFound     = errors.New("match not found")

	ErrAccountEmailTaken = errors.New("account with given email already exist")

	ErrTeamNotFound      = errors.New("team not found")
	ErrTeamClanNameTaken = errors.New("given clan name already taken")
)
