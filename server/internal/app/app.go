package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	kratos "github.com/ory/kratos-client-go"

	"github.com/ysomad/uniplay/server/internal/config"
	"github.com/ysomad/uniplay/server/internal/connectrpc"
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

	kratosClient := kratos.NewAPIClient(&kratos.Configuration{
		UserAgent:  fmt.Sprintf("%s/%s/%s/go", conf.App.Name, conf.App.Ver, conf.App.Environment),
		Debug:      conf.Kratos.Debug,
		Servers:    []kratos.ServerConfiguration{{URL: conf.Kratos.URL}},
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
	connectmux, err := connectrpc.NewMux(connectrpc.Deps{
		DemoStorage: demoStorage,
		Kratos:      kratosClient,
		OrgSchemaID: conf.Kratos.OrganizerSchemaID,
	})
	if err != nil {
		slog.Error("connect rpc mux not created", "error", err.Error())
		os.Exit(1)
	}

	connectsrv := httpserver.New(connectmux, httpserver.WithAddr(conf.Connect.Host, conf.Connect.Port))

	// http
	stdmux := httpapi.NewMux(httpapi.Deps{
		Minio:             minioClient,
		Kratos:            kratosClient,
		DemoStorage:       demoStorage,
		DemoBucket:        conf.ObjectStorage.DemoBucket,
		OrganizerSchemaID: conf.Kratos.OrganizerSchemaID,
	})
	stdsrv := httpserver.New(stdmux, httpserver.WithAddr(conf.HTTP.Host, conf.HTTP.Port))

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
