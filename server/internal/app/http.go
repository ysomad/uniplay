package app

import (
	"io"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	jsoniter "github.com/json-iterator/go"

	"github.com/ysomad/uniplay/internal/account"
	"github.com/ysomad/uniplay/internal/compendium"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/institution"
	"github.com/ysomad/uniplay/internal/match"
	"github.com/ysomad/uniplay/internal/player"
	"github.com/ysomad/uniplay/internal/team"

	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations"
	accountGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/account"
	compendiumGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/compendium"
	institutionGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/institution"
	matchGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/match"
	playerGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/player"
	teamGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/team"
)

type apiDeps struct {
	match       *match.Controller
	compendium  *compendium.Controller
	account     *account.Controller
	player      *player.Controller
	institution *institution.Controller
	team        *team.Controller
}

func jsonConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(r io.Reader, data any) error {
		dec := jsoniter.NewDecoder(r)
		dec.UseNumber()

		return dec.Decode(data)
	})
}

func jsonProducer() runtime.Producer {
	return runtime.ProducerFunc(func(w io.Writer, data any) error {
		enc := jsoniter.NewEncoder(w)
		enc.SetEscapeHTML(false)

		return enc.Encode(data)
	})
}

func newAPI(d apiDeps) (*operations.UniplayAPI, error) {
	spec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations.NewUniplayAPI(spec)
	api.UseSwaggerUI()

	// use jsoniter insted of encoding/json
	api.JSONConsumer = jsonConsumer()
	api.JSONProducer = jsonProducer()

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
	api.CompendiumGetCitiesHandler = compendiumGen.GetCitiesHandlerFunc(d.compendium.GetCities)

	api.AccountCreateAccountHandler = accountGen.CreateAccountHandlerFunc(d.account.CreateAccount)

	api.PlayerGetPlayerHandler = playerGen.GetPlayerHandlerFunc(d.player.GetPlayer)
	api.PlayerGetPlayerListHandler = playerGen.GetPlayerListHandlerFunc(d.player.GetPlayerList)
	api.PlayerUpdatePlayerHandler = playerGen.UpdatePlayerHandlerFunc(d.player.UpdatePlayer)
	api.PlayerGetPlayerStatsHandler = playerGen.GetPlayerStatsHandlerFunc(d.player.GetPlayerStats)
	api.PlayerGetWeaponStatsHandler = playerGen.GetWeaponStatsHandlerFunc(d.player.GetWeaponStats)
	api.PlayerGetPlayerMatchesHandler = playerGen.GetPlayerMatchesHandlerFunc(d.player.GetPlayerMatches)
	api.PlayerGetMostPlayedMapsHandler = playerGen.GetMostPlayedMapsHandlerFunc(d.player.GetMostPlayedMaps)
	api.PlayerGetMostSuccessMapsHandler = playerGen.GetMostSuccessMapsHandlerFunc(d.player.GetMostSuccessMaps)

	api.InstitutionGetInstitutionsHandler = institutionGen.GetInstitutionsHandlerFunc(d.institution.GetInstitutions)

	api.TeamGetTeamListHandler = teamGen.GetTeamListHandlerFunc(d.team.GetTeamList)
	api.TeamGetTeamPlayersHandler = teamGen.GetTeamPlayersHandlerFunc(d.team.GetTeamPlayers)
	api.TeamUpdateTeamHandler = teamGen.UpdateTeamHandlerFunc(d.team.UpdateTeam)
	api.TeamSetTeamCaptainHandler = teamGen.SetTeamCaptainHandlerFunc(d.team.SetTeamCaptain)
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

	mux.Handle("/", api.Serve(nil))

	return mux
}
