package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrMatchAlreadyExist = errors.New("match from the replay already exist")
)

type Match struct {
	ID         MatchID
	MapName    string
	Duration   time.Duration
	Team1      MatchTeam
	Team2      MatchTeam
	UploadTime time.Time
}

type MatchTeam struct {
	ClanName       string
	FlagCode       string
	Score          uint8
	PlayerSteamIDs []uint64
}

type MatchID struct {
	uuid.UUID
}
type MatchIDArgs struct {
	MapName       string
	Team1Name     string
	Team1Score    uint8
	Team2Name     string
	Team2Score    uint8
	MatchDuration time.Duration
}

func NewMatchID(a *MatchIDArgs) (MatchID, error) {
	if err := a.validate(); err != nil {
		return MatchID{}, err
	}

	s := fmt.Sprintf(
		"%s,%s,%d,%s,%d,%d",
		a.MapName,
		a.Team1Name,
		a.Team1Score,
		a.Team2Name,
		a.Team2Score,
		a.MatchDuration)
	return MatchID{uuid.NewMD5(uuid.UUID{}, []byte(s))}, nil
}

func (a *MatchIDArgs) validate() error {
	if a.MapName == "" {
		return errors.New("empty map name")
	}
	if a.Team1Name == "" {
		return errors.New("empty team 1 name")
	}
	if a.Team2Name == "" {
		return errors.New("empty team 2 name")
	}
	if a.MatchDuration <= time.Minute {
		return errors.New("match cannot last less than a minute")
	}
	return nil
}
