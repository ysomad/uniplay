package match

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	matchGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/match"

	"github.com/ysomad/uniplay/internal/domain"
)

type matchService interface {
	CreateFromReplay(context.Context, replay) (collectStatsRes, error)
	DeleteByID(ctx context.Context, matchID uuid.UUID) error
}

type Controller struct {
	match matchService
}

func NewController(m matchService) *Controller {
	return &Controller{
		match: m,
	}
}

const msgReplayFileNotFound = "replay file not found in request"

func (c *Controller) CreateMatch(p matchGen.CreateMatchParams) matchGen.CreateMatchResponder {
	formFile, ok := p.Replay.(*runtime.File)
	if !ok {
		return matchGen.NewCreateMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: msgReplayFileNotFound,
		})
	}

	r, err := newReplay(formFile.Data, formFile.Header)
	if err != nil {
		return matchGen.NewCreateMatchBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	defer r.Close()

	res, err := c.match.CreateFromReplay(p.HTTPRequest.Context(), r)
	if err != nil {
		if errors.Is(err, domain.ErrMatchAlreadyExist) {
			return matchGen.NewCreateMatchConflict().WithPayload(&models.Error{
				Code:    domain.CodeMatchAlreadyExist,
				Message: err.Error(),
			})
		}

		return matchGen.NewCreateMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := &models.CreateMatchResponse{
		MatchID:     strfmt.UUID(res.MatchID.String()),
		MatchNumber: res.MatchNumber,
	}

	return matchGen.NewCreateMatchOK().WithPayload(payload)
}

func (c *Controller) DeleteMatch(p matchGen.DeleteMatchParams) matchGen.DeleteMatchResponder {
	matchID, err := uuid.Parse(p.MatchID.String())
	if err != nil {
		return matchGen.NewDeleteMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if err := c.match.DeleteByID(p.HTTPRequest.Context(), matchID); err != nil {
		if errors.Is(err, domain.ErrMatchNotFound) {
			return matchGen.NewDeleteMatchNotFound().WithPayload(&models.Error{
				Code:    domain.CodeMatchNotFound,
				Message: err.Error(),
			})
		}

		return matchGen.NewDeleteMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return matchGen.NewDeleteMatchNoContent()
}
