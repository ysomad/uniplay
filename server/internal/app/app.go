package app

import (
	"context"
	"log"

	"github.com/IBM/pgxpoolprometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	"go.uber.org/zap"

	"github.com/ysomad/uniplay/internal/account"
	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/institution"
	"github.com/ysomad/uniplay/internal/match"
	"github.com/ysomad/uniplay/internal/player"
	"github.com/ysomad/uniplay/internal/team"

	"github.com/ysomad/uniplay/internal/pkg/logger"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

func Run(conf *config.Config, flags Flags) { //nolint:funlen // main func
	if flags.InDocker {
		initTempDir()
	}

	l, err := logger.New(conf.Log.Level)
	if err != nil {
		log.Fatalf("logger.New: %s", err.Error())
	}

	l.Info("starting app", zap.Any("flags", flags))

	otel, err := newOpenTelemetry(conf)
	if err != nil {
		l.Fatal("otel.New", zap.Error(err))
	}

	defer otel.meterProvider.Shutdown(context.Background())
	defer otel.tracerProvider.Shutdown(context.Background())

	pgClient, err := pgclient.New(
		conf.PG.URL,
		pgclient.WithMaxConns(conf.PG.MaxConns),
		pgclient.WithQueryTracer(otel.pgxTracer),
	)
	if err != nil {
		l.Fatal("pgclient.New", zap.Error(err))
	}

	// pgx metrics
	pgxCollector := pgxpoolprometheus.NewCollector(pgClient.Pool, map[string]string{"db_name": conf.PG.DBName})

	if err = prometheus.Register(pgxCollector); err != nil {
		l.Fatal("prometheus.Register", zap.Error(err))
	}

	// match
	matchPostgres := match.NewPostgres(otel.appTracer, pgClient)
	matchService := match.NewService(otel.appTracer, matchPostgres)
	matchController := match.NewController(matchService)

	// compendium
	compendiumPostgres := compendium.NewPostgres(pgClient)
	compendiumService := compendium.NewService(compendiumPostgres)
	compendiumController := compendium.NewController(compendiumService)

	// account
	accountPostgres := account.NewPostgres(otel.appTracer, pgClient)
	accountService := account.NewService(accountPostgres)
	accountController := account.NewController(accountService)

	// player
	playerPostgres := player.NewPostgres(otel.appTracer, pgClient)
	playerService := player.NewService(otel.appTracer, playerPostgres)
	playerController := player.NewController(playerService)

	// institution
	institutionPostgres := institution.NewPostgres(otel.appTracer, pgClient)
	institutionService := institution.NewService(institutionPostgres)
	institutionController := institution.NewController(institutionService)

	// team
	teamPostgres := team.NewPostgres(otel.appTracer, pgClient)
	teamService := team.NewService(teamPostgres)
	teamController := team.NewController(teamService)

	// go-swagger
	api, err := newAPI(apiDeps{
		compendium:  compendiumController,
		account:     accountController,
		player:      playerController,
		match:       matchController,
		institution: institutionController,
		team:        teamController,
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
	corsHandler := cors.Default().Handler(h)
	srv.SetHandler(corsHandler)

	if err = srv.Serve(); err != nil {
		l.Fatal("srv.Serve", zap.Error(err))
	}
}
