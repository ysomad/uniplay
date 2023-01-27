package replay

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"

	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/models"
	replayGen "github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations/replay"
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
	err := p.HTTPRequest.ParseMultipartForm(50 << 20)
	if err != nil {
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

	replay, err := newReplay(file, header.Filename)
	if err != nil {
		return replayGen.NewUploadReplayBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	defer replay.Close()

	matchID, err := c.replay.CollectStats(p.HTTPRequest.Context(), replay)
	if err != nil {

		if errors.Is(err, domain.ErrMatchAlreadyExist) {
			return replayGen.NewUploadReplayInternalServerError().WithPayload(&models.Error{
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
