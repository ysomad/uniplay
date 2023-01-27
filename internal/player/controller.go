package player

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/models"
	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations/player"
)

type playerService interface {
	GetStats(ctx context.Context, steamID uint64) (domain.PlayerStats, error)
	GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStat, error)
}

type Controller struct {
	log    *zap.Logger
	player playerService
}

func NewController(l *zap.Logger, p playerService) *Controller {
	return &Controller{
		log:    l,
		player: p,
	}
}

func (c *Controller) GetPlayerStats(p player.GetPlayerStatsParams) player.GetPlayerStatsResponder {
	steamID, err := strconv.ParseUint(p.SteamID, 10, 64)
	if err != nil {
		return player.NewGetPlayerStatsNotFound().WithPayload(&models.Error{
			Code:    domain.CodePlayerNotFound,
			Message: err.Error(),
		})
	}

	s, err := c.player.GetStats(p.HTTPRequest.Context(), steamID)
	if err != nil {

		if errors.Is(err, domain.ErrPlayerNotFound) {
			return player.NewGetPlayerStatsNotFound().WithPayload(&models.Error{
				Code:    domain.CodePlayerNotFound,
				Message: err.Error(),
			})
		}

		return player.NewGetPlayerStatsInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := &models.PlayerStats{
		TotalStats: &models.PlayerStatsTotalStats{
			Assists:            s.Total.Assists,
			BlindKills:         s.Total.BlindKills,
			BlindedPlayers:     s.Total.BlindedPlayers,
			BlindedTimes:       s.Total.BlindedTimes,
			BombsDefused:       s.Total.BombsDefused,
			BombsPlanted:       s.Total.BombsPlanted,
			DamageDealt:        s.Total.DamageDealt,
			DamageTaken:        s.Total.DamageTaken,
			Deaths:             s.Total.Deaths,
			Draws:              s.Total.Draws,
			FlashbangAssists:   s.Total.FlashbangAssists,
			GrenadeDamageDealt: s.Total.GrenadeDamageDealt,
			HeadshotKills:      s.Total.HeadshotKills,
			Kills:              s.Total.Kills,
			Loses:              s.Total.Loses,
			MatchesPlayed:      s.Total.MatchesPlayed,
			MvpCount:           s.Total.MVPCount,
			NoscopeKills:       s.Total.NoScopeKills,
			RoundsPlayed:       s.Total.RoundsPlayed,
			ThroughSmokeKills:  s.Total.ThroughSmokeKills,
			TimePlayed:         int64(s.Total.TimePlayed),
			WallbangKills:      s.Total.WallbangKills,
			Wins:               s.Total.Wins,
		},
		RoundStats: models.PlayerStatsRoundStats{
			Assists:            s.Round.Assists,
			BlindedPlayers:     s.Round.BlindedPlayers,
			BlindedTimes:       s.Round.BlindedTimes,
			DamageDealt:        s.Round.DamageDealt,
			Deaths:             s.Round.Deaths,
			GrenadeDamageDealt: s.Round.GrenadeDamageDealt,
			Kills:              s.Round.Kills,
		},
		CalculatedStats: models.PlayerStatsCalculatedStats{
			HeadshotPercentage: s.Calc.HeadshotPercentage,
			KillDeathRatio:     s.Calc.KillDeathRatio,
			WinRate:            s.Calc.WinRate,
		},
	}

	return player.NewGetPlayerStatsOK().WithPayload(payload)
}

func (c *Controller) GetWeaponStats(p player.GetWeaponStatsParams) player.GetWeaponStatsResponder {

	return player.NewGetWeaponStatsOK()
}
