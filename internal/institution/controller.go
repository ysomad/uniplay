package institution

import (
	"net/http"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/institution"
)

type Controller struct {
	institution *service
}

func NewController(s *service) *Controller {
	return &Controller{
		institution: s,
	}
}

func (c *Controller) GetInstitutions(p institution.GetInstitutionsParams) institution.GetInstitutionsResponder {
	list, err := c.institution.GetList(
		p.HTTPRequest.Context(),
		newGetListParams(
			p.Search,
			domain.NewInstitutionFilter(p.City, p.Type),
			paging.NewIntSeek(p.LastID, p.PageSize),
		),
	)
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
			City:      inst.City,
			Type:      int32(inst.Type),
		}
	}

	return institution.NewGetInstitutionsOK().WithPayload(&payload)
}