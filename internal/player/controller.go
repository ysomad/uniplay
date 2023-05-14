package player

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-openapi/strfmt"
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

func (c *Controller) GetPlayerList(p gen.GetPlayerListParams) gen.GetPlayerListResponder {
	lp, err := newListParams(p.Search, p.LastSteamID, p.PageSize)
	if err != nil {
		return gen.NewGetPlayerListInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	playerList, err := c.player.GetList(p.HTTPRequest.Context(), lp)
	if err != nil {
		return gen.NewGetPlayerListInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := models.PlayerList{
		Players: make([]models.PlayerListItem, len(playerList.Items)),
		HasNext: playerList.HasNext,
	}

	for i, p := range playerList.Items {
		payload.Players[i] = models.PlayerListItem{
			SteamID:     p.SteamID.String(),
			TeamID:      p.TeamID,
			DisplayName: p.DisplayName,
			FirstName:   p.FirstName,
			LastName:    p.LastName,
			AvatarURL:   p.AvatarURL,
		}
	}

	return gen.NewGetPlayerListOK().WithPayload(&payload)
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

	return gen.NewGetPlayerStatsOK().WithPayload(&models.PlayerStats{
		BaseStats: &models.PlayerStatsBaseStats{
			Assists:            s.Assists,
			BlindKills:         s.BlindKills,
			BlindedPlayers:     s.BlindedPlayers,
			BlindedTimes:       s.BlindedTimes,
			BombsDefused:       s.BombsDefused,
			BombsPlanted:       s.BombsPlanted,
			DamageDealt:        s.DamageDealt,
			DamageTaken:        s.DamageTaken,
			Deaths:             s.Deaths,
			FlashbangAssists:   s.FlashbangAssists,
			GrenadeDamageDealt: s.GrenadeDamageDealt,
			HeadshotKills:      s.HeadshotKills,
			Kills:              s.Kills,
			Loses:              s.Loses,
			MatchesPlayed:      s.MatchesPlayed,
			Mvps:               s.MVPCount,
			NoscopeKills:       s.NoScopeKills,
			RoundsPlayed:       s.RoundsPlayed,
			ThroughSmokeKills:  s.ThroughSmokeKills,
			TimePlayed:         int64(s.TimePlayed),
			WallbangKills:      s.WallbangKills,
			Wins:               s.Wins,
		},
		CalculatedStats: &models.PlayerStatsCalculatedStats{
			ADR:                    s.ADR,
			KD:                     s.KD,
			AssistsPerRound:        s.AssistsPerRound,
			BlindedPlayersPerRound: s.BlindedPlayersPerRound,
			DeathsPerRound:         s.DeathsPerRound,
			GrenadeDamagePerRound:  s.GrenadeDmgPerRound,
			HeadshotPercentage:     s.HeadshotPercentage,
			KillsPerRound:          s.KillsPerRound,
			WinRate:                s.WinRate,
		},
	})
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
			Weapon:   s.Base.Weapon,
			WeaponID: s.Base.WeaponID,
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
				ArmHits:           s.Base.ArmHits,
				LegHits:           s.Base.LegHits,
				TotalHits:         s.Base.TotalHits,
				NoscopeKills:      s.Base.NoScopeKills,
				Shots:             s.Base.Shots,
				StomachHits:       s.Base.StomachHits,
				ThroughSmokeKills: s.Base.ThroughSmokeKills,
				WallbangKills:     s.Base.WallbangKills,
			},
			AccuracyStats: &models.PlayerWeaponStatsInnerAccuracyStats{
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

func (c *Controller) GetPlayerMatches(p gen.GetPlayerMatchesParams) gen.GetPlayerMatchesResponder {
	lp, err := newMatchListParams(p.SteamID, p.PageToken, p.PageSize)
	if err != nil {
		return gen.NewGetPlayerMatchesInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	list, err := c.player.GetMatchList(p.HTTPRequest.Context(), lp)
	if err != nil {
		return gen.NewGetPlayerMatchesInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := models.PlayerMatchList{
		Matches:       make([]models.PlayerMatch, len(list.Items)),
		NextPageToken: string(list.NextPageToken),
	}

	for i, m := range list.Items {
		payload.Matches[i] = models.PlayerMatch{
			ID: strfmt.UUID(m.ID.String()),
			Map: models.Map{
				IconURL: m.Map.IconURL,
				Name:    m.Map.Name,
			},
			MatchState: int8(m.State),
			MatchStats: &models.PlayerMatchMatchStats{
				ADR:                m.Stats.ADR,
				Assists:            m.Stats.Assists,
				Deaths:             m.Stats.Deaths,
				HeadshotPercentage: m.Stats.HeadshotPercentage,
				Kills:              m.Stats.Kills,
			},
			Score:      string(m.Score),
			UploadedAt: strfmt.DateTime(m.UploadedAt),
		}
	}

	return gen.NewGetPlayerMatchesOK().WithPayload(&payload)
}
