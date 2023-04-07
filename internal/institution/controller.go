package institution

import (
	"context"
	"net/http"

	"github.com/ysomad/uniplay/internal/domain"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/institution"
)

type institutionService interface {
	GetInstitutionList(context.Context, domain.InstitutionFilter) ([]domain.Institution, error)
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
	filter := domain.InstitutionFilter{}
	if p.ShortName != nil {
		filter.ShortName = *p.ShortName
	}

	institutions, err := c.institution.GetInstitutionList(p.HTTPRequest.Context(), filter)
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
