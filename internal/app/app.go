package app

import (
	"context"
	"log"

	"github.com/IBM/pgxpoolprometheus"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/ysomad/uniplay/internal/account"
	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/institution"
	"github.com/ysomad/uniplay/internal/match"
	"github.com/ysomad/uniplay/internal/player"

	"github.com/ysomad/uniplay/internal/pkg/argon2"
	"github.com/ysomad/uniplay/internal/pkg/logger"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

func Run(conf *config.Config) {
	l, err := logger.New(conf.Log.Level)
	if err != nil {
		log.Fatalf("logger.New: %s", err.Error())
	}

	otel, err := newOpenTelemetry(conf)
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

	argon2ID := argon2.New()

	// match
	matchPostgres := match.NewPostgres(otel.AppTracer, pgClient)
	matchService := match.NewService(otel.AppTracer, matchPostgres)
	matchController := match.NewController(matchService)

	// compendium
	compendiumPostgres := compendium.NewPostgres(pgClient)
	compendiumService := compendium.NewService(compendiumPostgres)
	compendiumController := compendium.NewController(compendiumService)

	// account
	accountPostgres := account.NewPostgres(otel.AppTracer, pgClient)
	accountService := account.NewService(accountPostgres, argon2ID)
	accountController := account.NewController(accountService)

	// player
	playerPostgres := player.NewPostgres(otel.AppTracer, pgClient)
	playerService := player.NewService(otel.AppTracer, playerPostgres)
	playerController := player.NewController(playerService)

	// institution
	institutionPostgres := institution.NewPostgres(otel.AppTracer, pgClient)
	institutionService := institution.NewService(institutionPostgres)
	institutionController := institution.NewController(institutionService)

	// go-swagger
	api, err := newAPI(apiDeps{
		compendium:  compendiumController,
		account:     accountController,
		player:      playerController,
		match:       matchController,
		institution: institutionController,
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

	for _, f := range otel.CleanupFuncs {
		if err = f(context.Background()); err != nil {
			l.Fatal("cleanupFuncs", zap.Error(err))
		}
	}
}
