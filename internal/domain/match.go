package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type matchIDArgs struct {
	mapName       string
	team1name     string
	team1score    uint8
	team2name     string
	team2score    uint8
	matchDuration time.Duration
}

type MatchID struct {
	uuid.UUID
}

func newMatchID(a *matchIDArgs) MatchID {
	s := fmt.Sprintf(
		"%s,%s,%d,%s,%d,%d",
		a.mapName,
		a.team1name,
		a.team1score,
		a.team2name,
		a.team2score,
		a.matchDuration,
	)
	return MatchID{uuid.NewMD5(uuid.UUID{}, []byte(s))}
}

type Match struct {
	ID         MatchID
	MapName    string
	Duration   time.Duration
	Team1      MatchTeam
	Team2      MatchTeam
	UploadTime time.Time
}

func NewMatch(mapName string, matchDuration time.Duration, t1, t2 MatchTeam) (*Match, error) {
	if mapName == "" {
		return nil, errors.New("empty mapName")
	}

	if matchDuration.Minutes() < 5 {
		return nil, errors.New("too short match duration, minimum is 5 minutes")
	}

	if err := t1.validate(); err != nil {
		return nil, err
	}

	if err := t2.validate(); err != nil {
		return nil, err
	}

	return &Match{
		ID: newMatchID(&matchIDArgs{
			mapName:       mapName,
			team1name:     t1.Name,
			team1score:    t1.Score,
			team2name:     t2.Name,
			team2score:    t2.Score,
			matchDuration: matchDuration,
		}),
		MapName:    mapName,
		Duration:   matchDuration,
		Team1:      t1,
		Team2:      t2,
		UploadTime: time.Now(),
	}, nil
}

type MatchTeam struct {
	Name           string
	FlagCode       string
	Score          uint8
	PlayerSteamIDs []uint64
}

func (t *MatchTeam) SetAll(name, flag string, score uint8, steamIDs []uint64) {
	t.Name = name
	t.FlagCode = flag
	t.Score = score
	t.PlayerSteamIDs = steamIDs
}

func (t *MatchTeam) validate() error {
	if t.Name == "" {
		return errors.New("empty team name")
	}
	if len(t.PlayerSteamIDs) < 5 {
		return errors.New("amount of player steam ids in team must be atleast 5")
	}
	return nil
}
