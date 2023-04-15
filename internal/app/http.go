package app

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/institution"
	"github.com/ysomad/uniplay/internal/match"
	"github.com/ysomad/uniplay/internal/player"

	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations"
	compendiumGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/compendium"
	institutionGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/institution"
	matchGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/match"
	playerGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/player"
)

type apiDeps struct {
	match       *match.Controller
	compendium  *compendium.Controller
	player      *player.Controller
	institution *institution.Controller
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
	api.MatchCreateMatchHandler = matchGen.CreateMatchHandlerFunc(d.match.CreateMatch)
	api.MatchDeleteMatchHandler = matchGen.DeleteMatchHandlerFunc(d.match.DeleteMatch)
	api.MatchGetMatchHandler = matchGen.GetMatchHandlerFunc(d.match.GetMatch)

	api.CompendiumGetWeaponsHandler = compendiumGen.GetWeaponsHandlerFunc(d.compendium.GetWeapons)
	api.CompendiumGetWeaponClassesHandler = compendiumGen.GetWeaponClassesHandlerFunc(d.compendium.GetWeaponClasses)
	api.CompendiumGetMapsHandler = compendiumGen.GetMapsHandlerFunc(d.compendium.GetMaps)

	api.PlayerGetPlayerHandler = playerGen.GetPlayerHandlerFunc(d.player.GetPlayer)
	api.PlayerUpdatePlayerHandler = playerGen.UpdatePlayerHandlerFunc(d.player.UpdatePlayer)
	api.PlayerGetPlayerStatsHandler = playerGen.GetPlayerStatsHandlerFunc(d.player.GetPlayerStats)
	api.PlayerGetWeaponStatsHandler = playerGen.GetWeaponStatsHandlerFunc(d.player.GetWeaponStats)

	api.InstitutionGetInstitutionsHandler = institutionGen.GetInstitutionsHandlerFunc(d.institution.GetInstitutions)
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
