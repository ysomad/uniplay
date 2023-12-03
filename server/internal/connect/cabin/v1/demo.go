package v1

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ysomad/uniplay/server/internal/gen/connect/cabin/v1"
	connectpb "github.com/ysomad/uniplay/server/internal/gen/connect/cabin/v1/demov1connect"
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
	demo, err := s.demo.GetOne(ctx, r.Msg.DemoId)
	if err != nil {
		if errors.Is(err, postgres.ErrDemoNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.GetDemoResponse{
		Demo: &pb.Demo{
			Id:          demo.ID.String(),
			IdentityId:  demo.IdentityID,
			Status:      pb.DemoStatus(pb.DemoStatus_value[string(demo.Status)]),
			Reason:      demo.Reason,
			UploadedAt:  timestamppb.New(demo.UploadedAt),
			ProcessedAt: timestamppb.New(demo.ProcessedAt),
		},
	}), nil
}

func (s *DemoServer) ListDemos(context.Context, *connect.Request[pb.ListDemosRequest]) (*connect.Response[pb.ListDemosResponse], error) {
	return nil, nil
}
