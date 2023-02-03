package app

import (
	"log"

	"github.com/IBM/pgxpoolprometheus"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/opentelemetry"
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

	otel, err := opentelemetry.New(conf)
	if err != nil {
		l.Fatal("otel.New", zap.Error(err))
	}

	pgClient, err := pgclient.New(
		conf.PG.URL,
		pgclient.WithMaxConns(conf.PG.MaxConns),
		pgclient.WithQueryTracer(otel.PgxTracer),
	)
	if err != nil {
		l.Fatal("pgclient.New", zap.Error(err))
	}

	// pgx metrics
	pgxCollector := pgxpoolprometheus.NewCollector(pgClient.Pool, map[string]string{"db_name": conf.PG.DBName})

	if err = prometheus.Register(pgxCollector); err != nil {
		l.Fatal("prometheus.Register", zap.Error(err))
	}

	// replay
	replayRepo := replay.NewPostgres(otel.AppTracer, pgClient)
	replayService := replay.NewService(otel.AppTracer, replayRepo)
	replayController := replay.NewController(replayService)

	// compendium
	compendiumRepo := compendium.NewPostgres(pgClient)
	compendiumService := compendium.NewService(compendiumRepo)
	compendiumController := compendium.NewController(compendiumService)

	// player
	playerRepo := player.NewPostgres(otel.AppTracer, pgClient)
	playerService := player.NewService(otel.AppTracer, playerRepo)
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
