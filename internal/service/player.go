package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type Player struct {
	repo playerRepository
}

func NewPlayer(r playerRepository) *Player {
	return &Player{
		repo: r,
	}
}

func (p *Player) Get(ctx context.Context, steamID uint64) (domain.Player, error) {
	// return p.repo.FindBySteamID(ctx, steamID)
	return domain.Player{}, nil
}

func (p *Player) GetStats(ctx context.Context, steamID uint64) (*domain.PlayerStats, error) {
	total, err := p.repo.GetTotalStats(ctx, steamID)
	if err != nil {
		return nil, err
	}

	return &domain.PlayerStats{
		Total: total,
		Calc: domain.NewPlayerCalcStats(domain.PlayerCalcStatsParams{
			MatchesPlayed: total.MatchesPlayed,
			Kills:         total.Kills,
			Deaths:        total.Deaths,
			HeadshotKills: total.HeadshotKills,
			Wins:          total.Wins,
			Loses:         total.Loses,
		}),
		Round: domain.NewPlayerRoundStats(domain.PlayerRoundStatsParams{
			Kills:              total.Kills,
			Deaths:             total.Deaths,
			DamageDealt:        total.DamageDealt,
			Assists:            total.Assists,
			RoundsPlayed:       total.RoundsPlayed,
			GrenadeDamageDealt: total.GrenadeDamageDealt,
			BlindedPlayers:     total.BlindedPlayers,
			BlindedTimes:       total.BlindedTimes,
		}),
	}, nil
}
