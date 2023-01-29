package replay

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/ysomad/uniplay/internal/domain"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	replayGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/replay"
)

type replayService interface {
	CollectStats(context.Context, replay) (matchID uuid.UUID, err error)
}

type Controller struct {
	log    *zap.Logger
	replay replayService
}

func NewController(l *zap.Logger, r replayService) *Controller {
	return &Controller{
		log:    l,
		replay: r,
	}
}

func (c *Controller) UploadReplay(p replayGen.UploadReplayParams) replayGen.UploadReplayResponder {
	if err := p.HTTPRequest.ParseMultipartForm(150 << 20); err != nil {
		return replayGen.NewUploadReplayBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	file, header, err := p.HTTPRequest.FormFile("replay")
	if err != nil {
		return replayGen.NewUploadReplayBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	r, err := newReplay(file, header.Filename)
	if err != nil {
		return replayGen.NewUploadReplayBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	defer r.Close()

	matchID, err := c.replay.CollectStats(p.HTTPRequest.Context(), r)
	if err != nil {
		if errors.Is(err, domain.ErrMatchAlreadyExist) {
			return replayGen.NewUploadReplayConflict().WithPayload(&models.Error{
				Code:    domain.CodeMatchAlreadyExist,
				Message: err.Error(),
			})
		}

		return replayGen.NewUploadReplayInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := &models.UploadReplayResponse{MatchID: strfmt.UUID(matchID.String())}

	return replayGen.NewUploadReplayOK().WithPayload(payload)
}
