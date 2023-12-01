package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/ilyakaznacheev/cleanenv"
	ory "github.com/ory/client-go"

	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/domain/identity"
	"github.com/ysomad/uniplay/internal/errorx"
	"github.com/ysomad/uniplay/internal/middleware"
	"github.com/ysomad/uniplay/internal/pkg/demoparser"
	"github.com/ysomad/uniplay/internal/pkg/httpserver"
)

type uploadDemoRes struct {
	DemoJobID uuid.UUID `json:"demo_job_id"`
}

type demoHandlerV1 struct{}

func (h *demoHandlerV1) uploadDemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		errorx.WriteStatus(w, http.StatusMethodNotAllowed)
		return
	}

	// TODO: return specific error
	if err := r.ParseMultipartForm(10 >> 20); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// TODO: save demo file to some shared storage some-uuid.dem
	// save job with filename in job db
	// return job id
	// parse demo in background worker
	f, fh, err := r.FormFile("demo")
	if err != nil {
		return
	}
	defer f.Close()

	slog.Info("got demo file", "header", fh)

	parser, err := demoparser.New(f, fh)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer parser.Close()

	if err := parser.Parse(); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(uploadDemoRes{
		DemoJobID: uuid.New(), // TODO: return REAL demo job id
	}); err != nil {
		errorx.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func main() {
	var conf config.Config

	if err := cleanenv.ReadConfig("./configs/local.yml", &conf); err != nil {
		log.Fatalf("config parse error: %s", err)
	}

	Run(&conf)
}

func Run(conf *config.Config) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	})))

	oryClient := ory.NewAPIClient(&ory.Configuration{
		UserAgent:  fmt.Sprintf("%s/%s/%s/go", conf.App.Name, conf.App.Ver, conf.App.Environment),
		Debug:      conf.Kratos.Debug,
		Servers:    []ory.ServerConfiguration{{URL: conf.Kratos.URL}},
		HTTPClient: &http.Client{Timeout: conf.Kratos.ClientTimeout},
	})

	demoHandlerV1 := &demoHandlerV1{}

	mux := http.NewServeMux()

	// TODO: replace with identity.Organizer in production
	mux.Handle("/v1/demos", middleware.NewSessionAuth(oryClient, identity.User)(http.HandlerFunc(demoHandlerV1.uploadDemo)))

	srv := httpserver.New(mux, httpserver.WithHostPort(conf.HTTP.Host, conf.HTTP.Port))

	slog.Info("http server started", "host", conf.HTTP.Host, "port", conf.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("received signal from httpserver", "signal", s.String())
	case err := <-srv.Notify():
		slog.Info("got error from http server notify", "error", err.Error())
	}

	if err := srv.Shutdown(); err != nil {
		slog.Info("got error on http server shutdown", "error", err.Error())
	}
}
