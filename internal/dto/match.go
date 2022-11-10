package dto

import (
	"time"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type CreateMatchArgs struct {
	MapName  string
	Duration time.Duration
	Team1    domain.MatchTeam
	Team2    domain.MatchTeam
}
