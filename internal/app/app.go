package app

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/compendium"
	"github.com/ssssargsian/uniplay/internal/config"
	"github.com/ssssargsian/uniplay/internal/replay"

	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi"
	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations"
	compendiumGen "github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations/compendium"
	replayGen "github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations/replay"

	"github.com/ssssargsian/uniplay/internal/pkg/logger"
	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
)

func Run(conf *config.Config) {
	l, err := logger.New(os.Stderr, conf.Log.Level)
	if err != nil {
		log.Fatalf("logger.New: %s", err.Error())
	}

	pgClient, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
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

	// go-swagger
	api, err := newAPI(apiDeps{
		replay:     replayController,
		compendium: compendiumController,
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

type apiDeps struct {
	replay     *replay.Controller
	compendium *compendium.Controller
}

func newAPI(d apiDeps) (*operations.UniplayAPI, error) {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "2.0")
	if err != nil {
		return nil, err
	}

	api := operations.NewUniplayAPI(spec)
	api.UseSwaggerUI()

	api.CompendiumGetWeaponsHandler = compendiumGen.GetWeaponsHandlerFunc(d.compendium.GetWeapons)
	api.CompendiumGetWeaponClassesHandler = compendiumGen.GetWeaponClassesHandlerFunc(d.compendium.GetWeaponClasses)

	api.ReplayUploadReplayHandler = replayGen.UploadReplayHandlerFunc(d.replay.UploadReplay)

	return api, nil
}

func newServer(conf *config.Config, api *operations.UniplayAPI) *restapi.Server {
	srv := restapi.NewServer(api)
	srv.Host = conf.HTTP.Host
	srv.Port = conf.HTTP.Port
	srv.EnabledListeners = []string{"http"}
	return srv
}
