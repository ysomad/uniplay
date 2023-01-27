package app

import (
	"log"
	"os"

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

	pg, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
	if err != nil {
		l.Fatal("pgclient.New", zap.Error(err))
	}

	// replay
	replayRepo := replay.NewPGStorage(l, pg)
	replayService := replay.NewService(l, replayRepo)
	replayController := replay.NewController(l, replayService)

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

	srv := newServer(conf, api)
	defer srv.Shutdown()

	if err = srv.Serve(); err != nil {
		l.Fatal("srv.Serve", zap.Error(err))
	}
}
