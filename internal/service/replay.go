package service

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
	"github.com/ssssargsian/uniplay/internal/replayparser"
)

type replay struct {
	log  *zap.Logger
	repo replayRepository
}

func NewReplay(l *zap.Logger, r replayRepository) *replay {
	return &replay{
		log:  l,
		repo: r,
	}
}

// TODO: refactor with goroutines
func (r *replay) CollectStats(ctx context.Context, filename string) (*dto.Match, error) {
	replay, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	p := replayparser.New(replay, r.log)
	defer p.Close()

	res, err := p.Parse()
	if err != nil {
		return nil, err
	}

	match := res.Match()
	match.ID, err = domain.NewMatchID(&domain.MatchIDArgs{
		MapName:       match.MapName,
		Team1Name:     match.Team1.ClanName,
		Team1Score:    match.Team1.Score,
		Team2Name:     match.Team2.ClanName,
		Team2Score:    match.Team2.Score,
		MatchDuration: match.Duration,
	})
	if err != nil {
		return nil, err
	}

	err = r.repo.SaveTeams(ctx, dto.Teams{
		Team1Name:  match.Team1.ClanName,
		Team1Flag:  match.Team1.FlagCode,
		Team2Name:  match.Team2.ClanName,
		Team2Flag:  match.Team2.FlagCode,
		CreateTime: match.UploadTime,
	})
	if err != nil {
		return nil, err
	}

	if err = r.repo.SaveMatch(ctx, match); err != nil {
		return nil, err
	}

	teamPlayers := res.TeamPlayers()

	err = r.repo.SavePlayers(ctx, dto.MatchPlayers{
		MatchID:    match.ID,
		Players:    teamPlayers,
		CreateTime: match.UploadTime,
	})
	if err != nil {
		return nil, err
	}

	if err = r.repo.AddPlayersToTeams(ctx, teamPlayers); err != nil {
		return nil, err
	}

	playerStats, err := res.PlayerStats()
	if err != nil {
		return nil, err
	}

	weaponStats, err := res.WeaponStats()
	if err != nil {
		return nil, err
	}

	if err = r.repo.UpsertStats(ctx, playerStats, weaponStats); err != nil {
		return nil, err
	}

	return match, nil
}
