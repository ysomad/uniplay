package app

import (
	"context"
	"log"
	"os"

	"go.uber.org/zap"

	"github.com/exaring/otelpgx"
	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/player"
	"github.com/ysomad/uniplay/internal/replay"

	"github.com/ysomad/uniplay/internal/pkg/logger"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

func Run(conf *config.Config) {
	l, err := logger.New(os.Stderr, conf.Log.Level)
	if err != nil {
		log.Fatalf("logger.New: %s", err.Error())
	}

	// tracing
	jaegerExp, err := newJaegerExporter(conf.Jaeger)
	if err != nil {
		l.Fatal("newJaegerExporter", zap.Error(err))
	}

	shutdownTraceProvider, err := newTraceProvider(conf.App, jaegerExp)
	if err != nil {
		l.Fatal("newTraceProvider", zap.Error(err))
	}

	defer func() {
		if err = shutdownTraceProvider(context.Background()); err != nil {
			l.Fatal("shutdownTraceProvider", zap.Error(err))
		}
	}()

	// postgres
	pgTracer := otelpgx.NewTracer(otelpgx.WithTrimSQLInSpanName())

	pgClient, err := pgclient.New(
		conf.PG.URL,
		pgclient.WithMaxConns(conf.PG.MaxConns),
		pgclient.WithQueryTracer(pgTracer),
	)
	if err != nil {
		l.Fatal("pgclient.New", zap.Error(err))
	}

	// replay
	replayRepo := replay.NewPGStorage(l, pgClient)
	replayService := replay.NewService(l, replayRepo)
	replayController := replay.NewController(l, replayService)

	// compendium
	compendiumRepo := compendium.NewPGStorage(l, pgClient)
	compendiumService := compendium.NewService(compendiumRepo)
	compendiumController := compendium.NewController(l, compendiumService)

	// player
	playerRepo := player.NewPGStorage(l, pgClient)
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

	srv := newServer(conf.HTTP, api)

	defer func() {
		if err = srv.Shutdown(); err != nil {
			l.Fatal("srv.Shutdown", zap.Error(err))
		}
	}()

	h := newHandler(api)
	srv.SetHandler(h)

	if err = srv.Serve(); err != nil {
		l.Fatal("srv.Serve", zap.Error(err))
	}
}
