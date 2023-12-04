package v1

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ysomad/uniplay/server/internal/domain"
	pb "github.com/ysomad/uniplay/server/internal/gen/api/proto/cabin/v1"
	connectpb "github.com/ysomad/uniplay/server/internal/gen/api/proto/cabin/v1/cabinv1connect"
	appctx "github.com/ysomad/uniplay/server/internal/kratosctx"
	"github.com/ysomad/uniplay/server/internal/postgres"
)

var _ connectpb.DemoServiceHandler = &DemoServer{}

type DemoServer struct {
	demo postgres.DemoStorage
}

func NewDemoServer(s postgres.DemoStorage) *DemoServer {
	return &DemoServer{demo: s}
}

func (s *DemoServer) GetDemo(ctx context.Context, r *connect.Request[pb.GetDemoRequest]) (*connect.Response[pb.GetDemoResponse], error) {
	demo, err := s.demo.GetOne(ctx, r.Msg.DemoId, appctx.IdentityID(ctx))
	if err != nil {
		if errors.Is(err, postgres.ErrDemoNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.GetDemoResponse{
		Demo: &pb.Demo{
			Id:          demo.ID.String(),
			Status:      pb.DemoStatus(pb.DemoStatus_value[string(demo.Status)]),
			Reason:      demo.Reason,
			UploadedAt:  timestamppb.New(demo.UploadedAt),
			ProcessedAt: timestamppb.New(demo.ProcessedAt),
		},
	}), nil
}

func (s *DemoServer) ListDemos(ctx context.Context, r *connect.Request[pb.ListDemosRequest]) (*connect.Response[pb.ListDemosResponse], error) {
	demos, err := s.demo.GetAll(ctx, appctx.IdentityID(ctx), domain.DemoStatus(r.Msg.Status.String()))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&pb.ListDemosResponse{Demos: make([]*pb.Demo, len(demos))})

	for i, d := range demos {
		res.Msg.Demos[i] = &pb.Demo{
			Id:          d.ID.String(),
			Status:      pb.DemoStatus(pb.DemoStatus_value[string(d.Status)]),
			Reason:      d.Reason,
			UploadedAt:  timestamppb.New(d.UploadedAt),
			ProcessedAt: timestamppb.New(d.ProcessedAt),
		}
	}

	return res, nil
}
