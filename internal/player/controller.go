package player

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/ysomad/uniplay/internal/domain"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/player"
)

type playerService interface {
	GetStats(ctx context.Context, steamID uint64, f domain.PlayerStatsFilter) (domain.PlayerStats, error)
	GetWeaponStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]domain.WeaponStats, error)
}

type Controller struct {
	player playerService
}

func NewController(p playerService) *Controller {
	return &Controller{
		player: p,
	}
}

func (c *Controller) GetPlayerStats(p player.GetPlayerStatsParams) player.GetPlayerStatsResponder {
	var err error

	steamID, err := strconv.ParseUint(p.SteamID, 10, 64)
	if err != nil {
		return player.NewGetPlayerStatsBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	filter := domain.PlayerStatsFilter{}

	if p.MatchID != nil {
		filter.MatchID, err = uuid.Parse(p.MatchID.String())
		if err != nil {
			return player.NewGetPlayerStatsBadRequest().WithPayload(&models.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
	}

	s, err := c.player.GetStats(p.HTTPRequest.Context(), steamID, filter)
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

	// copilot is a LEGEND
	payload := &models.PlayerStats{
		BaseStats: &models.PlayerStatsBaseStats{
			Assists:            s.Base.Assists,
			BlindKills:         s.Base.BlindKills,
			BlindedPlayers:     s.Base.BlindedPlayers,
			BlindedTimes:       s.Base.BlindedTimes,
			BombsDefused:       s.Base.BombsDefused,
			BombsPlanted:       s.Base.BombsPlanted,
			DamageDealt:        s.Base.DamageDealt,
			DamageTaken:        s.Base.DamageTaken,
			Deaths:             s.Base.Deaths,
			Draws:              s.Base.Draws,
			FlashbangAssists:   s.Base.FlashbangAssists,
			GrenadeDamageDealt: s.Base.GrenadeDamageDealt,
			HeadshotKills:      s.Base.HeadshotKills,
			Kills:              s.Base.Kills,
			Loses:              s.Base.Loses,
			MatchesPlayed:      s.Base.MatchesPlayed,
			Mvps:               s.Base.MVPCount,
			NoscopeKills:       s.Base.NoScopeKills,
			RoundsPlayed:       s.Base.RoundsPlayed,
			ThroughSmokeKills:  s.Base.ThroughSmokeKills,
			TimePlayed:         int64(s.Base.TimePlayed),
			WallbangKills:      s.Base.WallbangKills,
			Wins:               s.Base.Wins,
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
	var err error

	steamID, err := strconv.ParseUint(p.SteamID, 10, 64)
	if err != nil {
		return player.NewGetWeaponStatsBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	filter := domain.WeaponStatsFilter{
		WeaponID: p.WeaponID,
		ClassID:  p.ClassID,
	}

	if p.MatchID != nil {
		filter.MatchID, err = uuid.Parse(p.MatchID.String())
		if err != nil {
			return player.NewGetWeaponStatsBadRequest().WithPayload(&models.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
	}

	weaponStats, err := c.player.GetWeaponStats(p.HTTPRequest.Context(), steamID, filter)
	if err != nil {
		if errors.Is(err, domain.ErrPlayerNotFound) {
			return player.NewGetWeaponStatsNotFound().WithPayload(&models.Error{
				Code:    domain.CodePlayerNotFound,
				Message: err.Error(),
			})
		}

		return player.NewGetWeaponStatsInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := make(models.PlayerWeaponStats, len(weaponStats))

	for i, s := range weaponStats {
		payload[i] = models.PlayerWeaponStatsInner{
			BaseStats: &models.PlayerWeaponStatsInnerBaseStats{
				Assists:           s.Base.Assists,
				BlindKills:        s.Base.BlindKills,
				ChestHits:         s.Base.ChestHits,
				DamageDealt:       s.Base.DamageDealt,
				DamageTaken:       s.Base.DamageTaken,
				Deaths:            s.Base.Deaths,
				HeadHits:          s.Base.HeadHits,
				NeckHits:          s.Base.NeckHits,
				HeadshotKills:     s.Base.HeadshotKills,
				Kills:             s.Base.Kills,
				LeftArmHits:       s.Base.LeftArmHits,
				LeftLegHits:       s.Base.LeftLegHits,
				NoscopeKills:      s.Base.NoScopeKills,
				RightArmHits:      s.Base.RightArmHits,
				RightLegHits:      s.Base.RightLegHits,
				Shots:             s.Base.Shots,
				StomachHits:       s.Base.StomachHits,
				ThroughSmokeKills: s.Base.ThroughSmokeKills,
				WallbangKills:     s.Base.WallbangKills,
				Weapon:            s.Base.Weapon,
				WeaponID:          s.Base.WeaponID,
			},
			AccuracyStats: models.PlayerWeaponStatsInnerAccuracyStats{
				Arms:    s.Accuracy.Arms,
				Chest:   s.Accuracy.Chest,
				Head:    s.Accuracy.Head,
				Neck:    s.Accuracy.Neck,
				Legs:    s.Accuracy.Legs,
				Stomach: s.Accuracy.Stomach,
				Total:   s.Accuracy.Total,
			},
		}
	}

	return player.NewGetWeaponStatsOK().WithPayload(payload)
}
