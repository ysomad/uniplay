package team

import (
	"errors"
	"net/http"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	gen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/team"
)

type Controller struct {
	team *service
}

func NewController(s *service) *Controller {
	return &Controller{
		team: s,
	}
}

func (c *Controller) GetTeamList(p gen.GetTeamListParams) gen.GetTeamListResponder {
	teamList, err := c.team.GetList(
		p.HTTPRequest.Context(),
		newListParams(p.Search, p.LastID, p.PageSize),
	)
	if err != nil {
		return gen.NewGetTeamListInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := &models.TeamList{
		Teams:   make([]models.TeamListItem, len(teamList.Items)),
		HasNext: teamList.HasNext,
	}

	for i, t := range teamList.Items {
		payload.Teams[i] = models.TeamListItem{
			ClanName: t.ClanName,
			FlagCode: t.FlagCode,
			ID:       t.ID,
		}

		// skip teams with no institution
		if t.InstID < 1 {
			continue
		}

		payload.Teams[i].Institution = &models.TeamListInstitution{
			City:      t.InstCity,
			ID:        t.InstID,
			LogoURL:   t.InstLogoURL,
			ShortName: t.InstShortName,
			Type:      int8(t.InstType),
		}
	}

	return gen.NewGetTeamListOK().WithPayload(payload)
}

func (c *Controller) GetTeamPlayers(p gen.GetTeamPlayersParams) gen.GetTeamPlayersResponder {
	teams, err := c.team.GetPlayers(p.HTTPRequest.Context(), p.TeamID)
	if err != nil {
		if errors.Is(err, domain.ErrTeamNotFound) {
			return gen.NewGetTeamPlayersNotFound().WithPayload(&models.Error{
				Code:    domain.CodeTeamNotFound,
				Message: err.Error(),
			})
		}

		return gen.NewGetTeamPlayersInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := make([]models.TeamPlayerListItem, len(teams))

	for i, t := range teams {
		payload[i] = models.TeamPlayerListItem{
			AvatarURL:   t.AvatarURL,
			DisplayName: t.DisplayName,
			FirstName:   t.FirstName,
			LastName:    t.LastName,
			SteamID:     t.SteamID.String(),
			IsCaptain:   t.IsCaptain,
		}
	}

	return gen.NewGetTeamPlayersOK().WithPayload(models.TeamPlayerList(payload))
}

func (c *Controller) UpdateTeam(p gen.UpdateTeamParams) gen.UpdateTeamResponder {
	t, err := c.team.Update(p.HTTPRequest.Context(), p.TeamID, updateParams{
		clanName:      p.Payload.ClanName,
		flagCode:      p.Payload.FlagCode,
		institutionID: p.Payload.InstitutionID,
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrTeamNotFound):
			return gen.NewUpdateTeamNotFound().WithPayload(&models.Error{
				Code:    domain.CodeTeamNotFound,
				Message: err.Error(),
			})
		case errors.Is(err, domain.ErrTeamClanNameTaken):
			return gen.NewUpdateTeamConflict().WithPayload(&models.Error{
				Code:    domain.CodeTeamClanNameTaken,
				Message: err.Error(),
			})
		}

		return gen.NewUpdateTeamInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return gen.NewUpdateTeamOK().WithPayload(&models.Team{
		ClanName:      t.ClanName,
		FlagCode:      t.FlagCode,
		ID:            t.ID,
		InstitutionID: t.InstitutionID,
	})
}
