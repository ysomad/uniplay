package service

import (
	"context"
	"io"
	"time"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
	"github.com/ssssargsian/uniplay/internal/replayparser"
)

type replay struct {
	repo replayRepository
}

func NewReplay(r replayRepository) *replay {
	return &replay{
		repo: r,
	}
}

func (r *replay) CollectStats(ctx context.Context, replay io.Reader) (*dto.Match, error) {
	p := replayparser.New(replay)
	defer p.Close()

	res, err := p.Parse()
	if err != nil {
		return nil, err
	}

	m := res.Match()
	now := time.Now()

	err = r.repo.SaveTeams(ctx, dto.Teams{
		Team1Name:  m.Team1.ClanName,
		Team1Flag:  m.Team1.FlagCode,
		Team2Name:  m.Team2.ClanName,
		Team2Flag:  m.Team2.FlagCode,
		CreateTime: now,
	})
	if err != nil {
		return nil, err
	}

	err = r.repo.SavePlayers(ctx, dto.PlayerSteamIDs{
		SteamIDs:   res.PlayerSteamIDs(),
		CreateTime: now,
	})
	if err != nil {
		return nil, err
	}

	if err = r.repo.AddPlayersToTeams(ctx, res.TeamPlayers()); err != nil {
		return nil, err
	}

	matchID, err := domain.NewMatchID(&domain.MatchIDArgs{
		MapName:       m.MapName,
		MatchDuration: m.Duration,
		Team1Name:     m.Team1.ClanName,
		Team1Score:    m.Team1.Score,
		Team2Name:     m.Team2.ClanName,
		Team2Score:    m.Team2.Score,
	})
	if err != nil {
		return nil, err
	}

	m.ID = matchID
	m.UploadTime = now

	if err = r.repo.SaveMatch(ctx, m); err != nil {
		return nil, err
	}

	return m, nil
}
