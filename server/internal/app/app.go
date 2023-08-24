package app

import (
	"log"

	"github.com/rs/cors"
	"github.com/ysomad/uniplay/internal/account"
	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/institution"
	"github.com/ysomad/uniplay/internal/match"
	"github.com/ysomad/uniplay/internal/pkg/logger"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
	"github.com/ysomad/uniplay/internal/player"
	"github.com/ysomad/uniplay/internal/team"
	"go.uber.org/zap"
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

	pgClient, err := pgclient.New(
		conf.PG.URL,
		pgclient.WithMaxConns(conf.PG.MaxConns),
	)
	if err != nil {
		l.Fatal("pgclient.New", zap.Error(err))
	}

	// match
	matchPostgres := match.NewPostgres(pgClient)
	matchService := match.NewService(matchPostgres)
	matchController := match.NewController(matchService)

	// compendium
	compendiumPostgres := compendium.NewPostgres(pgClient)
	compendiumService := compendium.NewService(compendiumPostgres)
	compendiumController := compendium.NewController(compendiumService)

	// account
	accountPostgres := account.NewPostgres(pgClient)
	accountService := account.NewService(accountPostgres)
	accountController := account.NewController(accountService)

	// player
	playerPostgres := player.NewPostgres(pgClient)
	playerService := player.NewService(playerPostgres)
	playerController := player.NewController(playerService)

	// institution
	institutionPostgres := institution.NewPostgres(pgClient)
	institutionService := institution.NewService(institutionPostgres)
	institutionController := institution.NewController(institutionService)

	// team
	teamPostgres := team.NewPostgres(pgClient)
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
