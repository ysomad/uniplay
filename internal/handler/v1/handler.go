package v1

import (
	"context"
	"io"

	"github.com/ssssargsian/uniplay/internal/dto"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
	"go.uber.org/zap"
)

var _ v1.ServerInterface = &handler{}

type atomicRunner interface {
	Run(context.Context, func(ctx context.Context) error) error
}

type replayService interface {
	CollectStats(context.Context, io.Reader) (*dto.Match, error)
}

type handler struct {
	log    *zap.Logger
	atomic atomicRunner
	replay replayService
}

func NewHandler(l *zap.Logger, a atomicRunner, r replayService) *handler {
	return &handler{
		log:    l,
		atomic: a,
		replay: r,
	}
}
