package replay

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
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
	tracer trace.Tracer
	replay replayService
}

func NewController(l *zap.Logger, t trace.Tracer, r replayService) *Controller {
	return &Controller{
		log:    l,
		tracer: t,
		replay: r,
	}
}

func (c *Controller) UploadReplay(p replayGen.UploadReplayParams) replayGen.UploadReplayResponder {
	ctx, span := c.tracer.Start(p.HTTPRequest.Context(), "replay.Controller.UploadReplay")
	defer span.End()

	if err := p.HTTPRequest.ParseMultipartForm(50 << 20); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return replayGen.NewUploadReplayBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	file, header, err := p.HTTPRequest.FormFile("replay")
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return replayGen.NewUploadReplayBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	r, err := newReplay(file, header.Filename)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return replayGen.NewUploadReplayBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	defer r.Close()

	matchID, err := c.replay.CollectStats(ctx, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

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
