package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	ory "github.com/ory/client-go"

	"github.com/ysomad/uniplay/server/internal/config"
	v1 "github.com/ysomad/uniplay/server/internal/connect/cabin/v1"
	"github.com/ysomad/uniplay/server/internal/connect/interceptor"
	"github.com/ysomad/uniplay/server/internal/gen/api/proto/cabin/v1/cabinv1connect"
	"github.com/ysomad/uniplay/server/internal/httpapi"
	"github.com/ysomad/uniplay/server/internal/pkg/httpserver"
	"github.com/ysomad/uniplay/server/internal/postgres"
	"github.com/ysomad/uniplay/server/internal/postgres/pgclient"
)

func Run(conf *config.Config, f Flags) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: conf.Log.SlogLevel(),
	})))

	if f.Migrate {
		mustMigrate(conf.PG.URL, f.MigrationsDir)
	}

	pgClient, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
	if err != nil {
		slog.Error("postgres client not created", "error", err.Error())
		os.Exit(1)
	}

	oryClient := ory.NewAPIClient(&ory.Configuration{
		UserAgent:  fmt.Sprintf("%s/%s/%s/go", conf.App.Name, conf.App.Ver, conf.App.Environment),
		Debug:      conf.Kratos.Debug,
		Servers:    []ory.ServerConfiguration{{URL: conf.Kratos.URL}},
		HTTPClient: &http.Client{Timeout: conf.Kratos.ClientTimeout},
	})

	minioClient, err := minio.New(conf.ObjectStorage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.ObjectStorage.AccessKey, conf.ObjectStorage.SecretKey, ""),
		Secure: conf.ObjectStorage.SSL,
	})
	if err != nil {
		slog.Error("minio client not created", "error", err)
		os.Exit(1)
	}

	demoStorage := postgres.NewDemoStorage(pgClient)

	slog.Debug("starting app", "config", conf)

	// connect
	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		slog.Error("protovalidate interceptor not created", "error", err)
		os.Exit(1)
	}

	demoServer := v1.NewDemoServer(demoStorage)

	connectsrv := newConnectSrv(connectSrvDeps{
		demosrv:             demoServer,
		kratos:              oryClient,
		conf:                conf.Connect,
		kratosConf:          conf.Kratos,
		validateInterceptor: validateInterceptor,
	})

	// http
	stdmux := httpapi.NewMux(httpapi.Deps{
		Minio:             minioClient,
		Ory:               oryClient,
		DemoStorage:       demoStorage,
		DemoBucket:        conf.ObjectStorage.DemoBucket,
		OrganizerSchemaID: conf.Kratos.OrganizerSchemaID,
	})
	stdsrv := httpserver.New(stdmux, httpserver.WithHostPort(conf.HTTP.Host, conf.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	select {
	case s := <-interrupt:
		slog.Info("received interrupt signal", "signal", s.String())
	case err := <-stdsrv.Notify():
		slog.Error("got error from http server", "error", err.Error())
	case err := <-connectsrv.Notify():
		slog.Error("got error from connect server", "error", err.Error())
	}

	if err := stdsrv.Shutdown(); err != nil {
		slog.Error("got error on http server shutdown", "error", err.Error())
	}

	if err := connectsrv.Shutdown(); err != nil {
		slog.Error("go error on connect server shutdown", "error", err.Error())
	}
}

type connectSrvDeps struct {
	demosrv             *v1.DemoServer
	kratos              *ory.APIClient
	conf                config.Connect
	kratosConf          config.Kratos
	validateInterceptor *validate.Interceptor
}

func newConnectSrv(deps connectSrvDeps) *httpserver.Server {
	defer slog.Info("connect server started", "host", deps.conf.Host, "port", deps.conf.Port)
	mux := http.NewServeMux()
	path, handler := cabinv1connect.NewDemoServiceHandler(
		deps.demosrv, connect.WithInterceptors(
			deps.validateInterceptor,
			interceptor.NewAuth(deps.kratos, deps.kratosConf),
		))
	mux.Handle(path, handler)
	return httpserver.New(mux, httpserver.WithHostPort(deps.conf.Host, deps.conf.Port))
}
