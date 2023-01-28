package app

import (
	"github.com/go-openapi/loads"

	"github.com/ssssargsian/uniplay/internal/compendium"
	"github.com/ssssargsian/uniplay/internal/config"
	"github.com/ssssargsian/uniplay/internal/player"
	"github.com/ssssargsian/uniplay/internal/replay"

	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi"
	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations"
	compendiumGen "github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations/compendium"
	playerGen "github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations/player"
	replayGen "github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations/replay"
)

type apiDeps struct {
	replay     *replay.Controller
	compendium *compendium.Controller
	player     *player.Controller
}

func newAPI(d apiDeps) (*operations.UniplayAPI, error) {
	spec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations.NewUniplayAPI(spec)
	api.UseSwaggerUI()

	attachHandlers(api, d)

	return api, nil
}

func attachHandlers(api *operations.UniplayAPI, d apiDeps) {
	api.CompendiumGetWeaponsHandler = compendiumGen.GetWeaponsHandlerFunc(d.compendium.GetWeapons)
	api.CompendiumGetWeaponClassesHandler = compendiumGen.GetWeaponClassesHandlerFunc(d.compendium.GetWeaponClasses)

	api.ReplayUploadReplayHandler = replayGen.UploadReplayHandlerFunc(d.replay.UploadReplay)

	api.PlayerGetPlayerStatsHandler = playerGen.GetPlayerStatsHandlerFunc(d.player.GetPlayerStats)
	api.PlayerGetWeaponStatsHandler = playerGen.GetWeaponStatsHandlerFunc(d.player.GetWeaponStats)
}

func newServer(conf config.HTTP, api *operations.UniplayAPI) *restapi.Server {
	srv := restapi.NewServer(api)
	srv.Host = conf.Host
	srv.Port = conf.Port
	srv.EnabledListeners = []string{"http"}

	return srv
}
