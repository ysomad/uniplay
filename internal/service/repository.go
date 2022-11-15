package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayRepository interface {
	SavePlayers(context.Context, dto.PlayerSteamIDs) error
	SaveTeams(context.Context, dto.Teams) error
	AddPlayersToTeams(context.Context, []dto.TeamPlayer) error
	SaveMetrics(context.Context, []dto.Metric, []dto.WeaponMetric) error
	SaveMatch(context.Context, *dto.Match) error
}
