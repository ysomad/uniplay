package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ory "github.com/ory/client-go"

	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/httpapi"
	"github.com/ysomad/uniplay/internal/httpapi/middleware"
	"github.com/ysomad/uniplay/internal/pkg/httpserver"
)

func runHTTPServer(conf *config.Config) {
	mux := newHTTPMux(conf)
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

func newHTTPMux(conf *config.Config) *http.ServeMux {
	oryClient := ory.NewAPIClient(&ory.Configuration{
		UserAgent:  fmt.Sprintf("%s/%s/%s/go", conf.App.Name, conf.App.Ver, conf.App.Environment),
		Debug:      conf.Kratos.Debug,
		Servers:    []ory.ServerConfiguration{{URL: conf.Kratos.URL}},
		HTTPClient: &http.Client{Timeout: conf.Kratos.ClientTimeout},
	})

	demoV1 := httpapi.NewDemoV1()
	mux := http.NewServeMux()

	// TODO: replace with identity.Organizer in production
	mux.Handle("/v1/demos", middleware.NewSessionAuth(oryClient, conf.Kratos.OrganizerSchemaID)(http.HandlerFunc(demoV1.Upload)))

	return mux
}
