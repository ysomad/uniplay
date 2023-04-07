package institution

import (
	"context"
	"net/http"

	"github.com/ysomad/uniplay/internal/domain"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/institution"
)

type institutionService interface {
	GetInstitutionList(context.Context, domain.InstitutionFilter, domain.InstitutionPagination) ([]domain.Institution, error)
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
	var (
		filter   domain.InstitutionFilter
		pageSize int32
		offset   int32
	)

	switch {
	case p.ShortName != nil:
		filter.ShortName = *p.ShortName
	case p.PageSize != nil:
		pageSize = *p.PageSize
	case p.Offset != nil:
		offset = *p.Offset
	}

	institutions, err := c.institution.GetInstitutionList(
		p.HTTPRequest.Context(),
		filter,
		domain.NewInstitutionPagination(pageSize, offset),
	)
	if err != nil {
		return institution.NewGetInstitutionsInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := make(models.InstitutionList, len(institutions))

	for i, institution := range institutions {
		payload[i] = models.InstitutionListInner{
			ID:        int32(institution.ID),
			Name:      institution.Name,
			ShortName: institution.ShortName,
			LogoURL:   institution.LogoURL,
		}
	}

	return institution.NewGetInstitutionsOK().WithPayload(payload)
}
