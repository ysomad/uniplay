package app

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/player"
	"github.com/ysomad/uniplay/internal/replay"

	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations"
	compendiumGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/compendium"
	playerGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/player"
	replayGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/replay"
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

	_attachHandlers(api, d)

	return api, nil
}

func _attachHandlers(api *operations.UniplayAPI, d apiDeps) {
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

func newHandler(api *operations.UniplayAPI) http.Handler {
	mux := http.DefaultServeMux

	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", otelhttp.NewHandler(api.Serve(nil), ""))

	return mux
}
