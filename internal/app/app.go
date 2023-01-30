package app

import (
	"context"
	"log"

	"github.com/exaring/otelpgx"
	"go.uber.org/zap"

	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/player"
	"github.com/ysomad/uniplay/internal/replay"

	"github.com/ysomad/uniplay/internal/pkg/logger"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

func Run(conf *config.Config) {
	l, err := logger.New(conf.Log.Level)
	if err != nil {
		log.Fatalf("logger.New: %s", err.Error())
	}

	// tracing
	jaegerExp, err := newJaegerExporter(conf.Jaeger)
	if err != nil {
		l.Fatal("newJaegerExporter", zap.Error(err))
	}

	shutdownTraceProvider := newTraceProvider(conf.App, jaegerExp)

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
	replayRepo := replay.NewPostgres(pgClient)
	replayService := replay.NewService(replayRepo)
	replayController := replay.NewController(replayService)

	// compendium
	compendiumRepo := compendium.NewPostgres(pgClient)
	compendiumService := compendium.NewService(compendiumRepo)
	compendiumController := compendium.NewController(compendiumService)

	// player
	playerRepo := player.NewPostgres(pgClient)
	playerService := player.NewService(playerRepo)
	playerController := player.NewController(playerService)

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
