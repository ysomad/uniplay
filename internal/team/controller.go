package team

import (
	"net/http"

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
