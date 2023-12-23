package httpapi

import (
	"net/http"

	"github.com/minio/minio-go/v7"
	kratos "github.com/ory/kratos-client-go"

	"github.com/ysomad/uniplay/server/internal/postgres"
)

type Deps struct {
	Minio             *minio.Client
	Kratos            *kratos.APIClient
	DemoStorage       postgres.DemoStorage
	DemoBucket        string
	OrganizerSchemaID string
}

func NewMux(d Deps) *http.ServeMux {
	demov1 := &demoV1{
		minio:  d.Minio,
		bucket: d.DemoBucket,
		demo:   d.DemoStorage,
	}

	kratosmw := newAuthMiddleware(d.Kratos, d.OrganizerSchemaID)

	mux := http.NewServeMux()
	mux.Handle("/v1/demos", kratosmw(http.HandlerFunc(demov1.upload)))

	return mux
}
