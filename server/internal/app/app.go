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
	"github.com/ory/client-go"

	"github.com/ysomad/uniplay/server/internal/config"
	"github.com/ysomad/uniplay/server/internal/httpapi"
	"github.com/ysomad/uniplay/server/internal/httpapi/middleware"
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

	slog.Debug("starting app", "config", conf)

	kratosClient := client.NewAPIClient(&client.Configuration{
		UserAgent:  fmt.Sprintf("%s/%s/%s/go", conf.App.Name, conf.App.Ver, conf.App.Environment),
		Debug:      conf.Kratos.Debug,
		Servers:    []client.ServerConfiguration{{URL: conf.Kratos.URL}},
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

	pgClient, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
	if err != nil {
		slog.Error("postgres client not created", "error", err.Error())
		os.Exit(1)
	}

	demoStorage := postgres.DemoStorage{Client: pgClient}

	demoV1 := httpapi.NewDemoV1(minioClient, conf.ObjectStorage.DemoBucket, demoStorage)
	kratosMW := middleware.NewKratos(kratosClient, conf.Kratos)
	mux := http.NewServeMux()

	// TODO: replace with identity.Organizer in production
	mux.Handle("/v1/demos", kratosMW.SessionAuth(http.HandlerFunc(demoV1.Upload)))

	srv := httpserver.New(mux, httpserver.WithHostPort(conf.HTTP.Host, conf.HTTP.Port))

	slog.Info("http server started", "host", conf.HTTP.Host, "port", conf.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("received signal from httpserver", "signal", s.String())
	case err := <-srv.Notify():
		slog.Error("got error from http server notify", "error", err.Error())
	}

	if err := srv.Shutdown(); err != nil {
		slog.Error("got error on http server shutdown", "error", err.Error())
	}
}
