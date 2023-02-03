package match

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/match"
)

type matchService interface {
	Delete(ctx context.Context, matchID uuid.UUID) error
}

type Controller struct {
	match matchService
}

func NewController(m matchService) *Controller {
	return &Controller{
		match: m,
	}
}

func (c *Controller) DeleteMatch(p match.DeleteMatchParams) match.DeleteMatchResponder {
	matchID, err := uuid.Parse(p.MatchID.String())
	if err != nil {
		return match.NewDeleteMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if err := c.match.Delete(p.HTTPRequest.Context(), matchID); err != nil {
		if errors.Is(err, domain.ErrMatchNotFound) {
			return match.NewDeleteMatchNotFound().WithPayload(&models.Error{
				Code:    domain.CodeMatchNotFound,
				Message: err.Error(),
			})
		}

		return match.NewDeleteMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return match.NewDeleteMatchNoContent()
}
