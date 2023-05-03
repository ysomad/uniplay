package player

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ysomad/uniplay/internal/domain"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	gen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/player"
)

type Controller struct {
	player *service
}

func NewController(s *service) *Controller {
	return &Controller{
		player: s,
	}
}

func (c *Controller) GetPlayer(p gen.GetPlayerParams) gen.GetPlayerResponder {
	steamID, err := domain.NewSteamID(p.SteamID)
	if err != nil {
		return gen.NewGetPlayerInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	pl, err := c.player.GetBySteamID(p.HTTPRequest.Context(), steamID)
	if err != nil {
		if errors.Is(err, domain.ErrPlayerNotFound) {
			return gen.NewGetPlayerNotFound().WithPayload(&models.Error{
				Code:    domain.CodePlayerNotFound,
				Message: err.Error(),
			})
		}

		return gen.NewGetPlayerInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return gen.NewGetPlayerOK().WithPayload(&models.Player{
		SteamID:     pl.SteamID.String(),
		TeamID:      pl.TeamID,
		DisplayName: pl.DisplayName,
		FirstName:   pl.FirstName,
		LastName:    pl.LastName,
		AvatarURL:   pl.AvatarURL,
	})
}

func (c *Controller) UpdatePlayer(p gen.UpdatePlayerParams) gen.UpdatePlayerResponder {
	steamID, err := domain.NewSteamID(p.SteamID)
	if err != nil {
		return gen.NewUpdatePlayerInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	pl, err := c.player.UpdateBySteamID(p.HTTPRequest.Context(), steamID, updateParams{
		teamID:    p.Payload.TeamID,
		firstName: p.Payload.FirstName,
		lastName:  p.Payload.LastName,
		avatarURL: p.Payload.AvatarURL.String(),
	})
	if err != nil {
		if errors.Is(err, domain.ErrPlayerNotFound) {
			return gen.NewUpdatePlayerNotFound().WithPayload(&models.Error{
				Code:    domain.CodePlayerNotFound,
				Message: err.Error(),
			})
		}

		return gen.NewUpdatePlayerInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return gen.NewUpdatePlayerOK().WithPayload(&models.Player{
		SteamID:     pl.SteamID.String(),
		DisplayName: pl.DisplayName,
		FirstName:   pl.FirstName,
		LastName:    pl.LastName,
		TeamID:      pl.TeamID,
		AvatarURL:   pl.AvatarURL,
	})
}

func (c *Controller) GetPlayerStats(p gen.GetPlayerStatsParams) gen.GetPlayerStatsResponder {
	var err error

	steamID, err := strconv.ParseUint(p.SteamID, 10, 64)
	if err != nil {
		return gen.NewGetPlayerStatsBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	s, err := c.player.GetStats(p.HTTPRequest.Context(), steamID)
	if err != nil {
		if errors.Is(err, domain.ErrPlayerNotFound) {
			return gen.NewGetPlayerStatsNotFound().WithPayload(&models.Error{
				Code:    domain.CodePlayerNotFound,
				Message: err.Error(),
			})
		}

		return gen.NewGetPlayerStatsInternalServerError().WithPayload(&models.Error{
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

	return gen.NewGetPlayerStatsOK().WithPayload(payload)
}

func (c *Controller) GetWeaponStats(p gen.GetWeaponStatsParams) gen.GetWeaponStatsResponder {
	var err error

	steamID, err := strconv.ParseUint(p.SteamID, 10, 64)
	if err != nil {
		return gen.NewGetWeaponStatsBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	weaponStats, err := c.player.GetWeaponStats(
		p.HTTPRequest.Context(),
		steamID,
		domain.NewWeaponStatsFilter(p.WeaponID, p.ClassID),
	)
	if err != nil {
		if errors.Is(err, domain.ErrPlayerNotFound) {
			return gen.NewGetWeaponStatsNotFound().WithPayload(&models.Error{
				Code:    domain.CodePlayerNotFound,
				Message: err.Error(),
			})
		}

		return gen.NewGetWeaponStatsInternalServerError().WithPayload(&models.Error{
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

	return gen.NewGetWeaponStatsOK().WithPayload(payload)
}
