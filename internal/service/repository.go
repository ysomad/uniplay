package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayRepository interface {
	SavePlayers(context.Context, dto.MatchPlayers) error
	SaveTeams(context.Context, dto.Teams) error
	AddPlayersToTeams(context.Context, []dto.TeamPlayer) error
	SaveMetrics(context.Context, []dto.Metric, []dto.WeaponMetric) error
	SaveMatch(context.Context, *dto.Match) error
}

type playerRepository interface {
	FindBySteamID(ctx context.Context, steamID uint64) (domain.Player, error)
}
