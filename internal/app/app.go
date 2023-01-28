package app

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/compendium"
	"github.com/ssssargsian/uniplay/internal/config"
	"github.com/ssssargsian/uniplay/internal/player"
	"github.com/ssssargsian/uniplay/internal/replay"

	"github.com/ssssargsian/uniplay/internal/pkg/logger"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
)

func Run(conf *config.Config) {
	l, err := logger.New(os.Stderr, conf.Log.Level)
	if err != nil {
		log.Fatalf("logger.New: %s", err.Error())
	}

	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal("os.Create", zap.Error(err))
	}
	defer f.Close()

	exp, err := newStdoutTracerExporter(f)
	if err != nil {
		l.Fatal("newStdoutTracerExporter", zap.Error(err))
	}

	tp, err := newTraceProvider(conf.App, exp)
	if err != nil {
		l.Fatal("newTraceProvider", zap.Error(err))
	}
	defer func() {
		if err = tp.Shutdown(context.Background()); err != nil {
			l.Fatal("tp.Shutdown", zap.Error(err))
		}
	}()

	otel.SetTracerProvider(tp)

	tracer := tp.Tracer(conf.App.Name)

	pg, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
	if err != nil {
		l.Fatal("pgclient.New", zap.Error(err))
	}

	// replay
	replayRepo := replay.NewPGStorage(l, tracer, pg)
	replayService := replay.NewService(l, tracer, replayRepo)
	replayController := replay.NewController(l, tracer, replayService)

	// compendium
	compendiumRepo := compendium.NewPGStorage(l, pg)
	compendiumService := compendium.NewService(compendiumRepo)
	compendiumController := compendium.NewController(l, compendiumService)

	// player
	playerRepo := player.NewPGStorage(l, pg)
	playerService := player.NewService(playerRepo)
	playerController := player.NewController(l, playerService)

	// go-swagger
	api, err := newAPI(apiDeps{
		replay:     replayController,
		compendium: compendiumController,
		player:     playerController,
	})
	if err != nil {
		l.Fatal("newAPI", zap.Error(err))
	}

	// параметр внутри - миддлвари, которые выполнятся после раутинга и валидации
	_ = api.Serve(nil)

	srv := newServer(conf.HTTP, api)
	defer func() {
		if err = srv.Shutdown(); err != nil {
			l.Fatal("srv.Shutdown", zap.Error(err))
		}
	}()

	if err = srv.Serve(); err != nil {
		l.Fatal("srv.Serve", zap.Error(err))
	}
}
