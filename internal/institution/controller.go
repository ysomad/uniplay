package institution

import (
	"context"
	"net/http"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/institution"
)

type institutionService interface {
	GetInstitutionList(context.Context, domain.InstitutionFilter, paging.IntSeek[int32]) (paging.InfList[domain.Institution], error)
}

type Controller struct {
	institution institutionService
}

func NewController(s institutionService) *Controller {
	return &Controller{
		institution: s,
	}
}

func (c *Controller) GetInstitutions(p institution.GetInstitutionsParams) institution.GetInstitutionsResponder {
	var filter domain.InstitutionFilter

	if p.ShortName != nil {
		filter.ShortName = *p.ShortName
	}

	list, err := c.institution.GetInstitutionList(p.HTTPRequest.Context(), filter, paging.NewIntSeek(p.LastID, p.PageSize))
	if err != nil {
		return institution.NewGetInstitutionsInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := models.InstitutionList{
		Institutions: make([]models.InstitutionListItem, len(list.Items)),
		HasNext:      list.HasNext,
	}

	for i, inst := range list.Items {
		payload.Institutions[i] = models.InstitutionListItem{
			ID:        inst.ID,
			Name:      inst.Name,
			ShortName: inst.ShortName,
			LogoURL:   inst.LogoURL,
		}
	}

	return institution.NewGetInstitutionsOK().WithPayload(&payload)
}
