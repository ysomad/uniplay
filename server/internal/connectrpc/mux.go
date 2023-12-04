package connectrpc

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	kratos "github.com/ory/kratos-client-go"

	cabinv1 "github.com/ysomad/uniplay/server/internal/connectrpc/cabin/v1"
	"github.com/ysomad/uniplay/server/internal/gen/api/proto/cabin/v1/cabinv1connect"
	"github.com/ysomad/uniplay/server/internal/postgres"
)

type Deps struct {
	DemoStorage postgres.DemoStorage
	Kratos      *kratos.APIClient
	OrgSchemaID string
}

func NewMux(d Deps) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		return nil, fmt.Errorf("validate interceptor not created: %w", err)
	}

	authInterceptor := newAuthInterceptor(d.Kratos, d.OrgSchemaID)
	demosrv := cabinv1.NewDemoServer(d.DemoStorage)

	path, handler := cabinv1connect.NewDemoServiceHandler(demosrv, connect.WithInterceptors(
		validateInterceptor, authInterceptor))
	mux.Handle(path, handler)

	return mux, nil
}
