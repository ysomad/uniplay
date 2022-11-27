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
	UpsertStats(context.Context, []dto.PlayerStat, []dto.WeaponStat) error
	SaveMatch(context.Context, *dto.Match) error
}

type playerRepository interface {
	FindBySteamID(ctx context.Context, steamID uint64) (domain.Player, error)
}

type statisticRepository interface {
	GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]dto.WeaponStatWithClass, error)
	GetWeaponClassStats(ctx context.Context, steamID uint64, classID uint8) ([]dto.WeaponClassStat, error)
}

type compendiumRepository interface {
	GetWeaponList(context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(context.Context) ([]domain.WeaponClass, error)
}
