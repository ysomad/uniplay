// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/account"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/compendium"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/institution"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/match"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/player"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/team"
)

//go:generate swagger generate server --target ../../swagger2 --name Uniplay --spec ../../../../swagger2.yaml --principal interface{} --exclude-main --strict-responders

func configureFlags(api *operations.UniplayAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UniplayAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// match.CreateMatchMaxParseMemory = 32 << 20

	if api.AccountCreateAccountHandler == nil {
		api.AccountCreateAccountHandler = account.CreateAccountHandlerFunc(func(params account.CreateAccountParams) account.CreateAccountResponder {
			return account.CreateAccountNotImplemented()
		})
	}
	if api.MatchCreateMatchHandler == nil {
		api.MatchCreateMatchHandler = match.CreateMatchHandlerFunc(func(params match.CreateMatchParams) match.CreateMatchResponder {
			return match.CreateMatchNotImplemented()
		})
	}
	if api.MatchDeleteMatchHandler == nil {
		api.MatchDeleteMatchHandler = match.DeleteMatchHandlerFunc(func(params match.DeleteMatchParams) match.DeleteMatchResponder {
			return match.DeleteMatchNotImplemented()
		})
	}
	if api.CompendiumGetCitiesHandler == nil {
		api.CompendiumGetCitiesHandler = compendium.GetCitiesHandlerFunc(func(params compendium.GetCitiesParams) compendium.GetCitiesResponder {
			return compendium.GetCitiesNotImplemented()
		})
	}
	if api.InstitutionGetInstitutionsHandler == nil {
		api.InstitutionGetInstitutionsHandler = institution.GetInstitutionsHandlerFunc(func(params institution.GetInstitutionsParams) institution.GetInstitutionsResponder {
			return institution.GetInstitutionsNotImplemented()
		})
	}
	if api.CompendiumGetMapsHandler == nil {
		api.CompendiumGetMapsHandler = compendium.GetMapsHandlerFunc(func(params compendium.GetMapsParams) compendium.GetMapsResponder {
			return compendium.GetMapsNotImplemented()
		})
	}
	if api.MatchGetMatchHandler == nil {
		api.MatchGetMatchHandler = match.GetMatchHandlerFunc(func(params match.GetMatchParams) match.GetMatchResponder {
			return match.GetMatchNotImplemented()
		})
	}
	if api.PlayerGetMostPlayedMapsHandler == nil {
		api.PlayerGetMostPlayedMapsHandler = player.GetMostPlayedMapsHandlerFunc(func(params player.GetMostPlayedMapsParams) player.GetMostPlayedMapsResponder {
			return player.GetMostPlayedMapsNotImplemented()
		})
	}
	if api.PlayerGetMostSuccessMapsHandler == nil {
		api.PlayerGetMostSuccessMapsHandler = player.GetMostSuccessMapsHandlerFunc(func(params player.GetMostSuccessMapsParams) player.GetMostSuccessMapsResponder {
			return player.GetMostSuccessMapsNotImplemented()
		})
	}
	if api.PlayerGetPlayerHandler == nil {
		api.PlayerGetPlayerHandler = player.GetPlayerHandlerFunc(func(params player.GetPlayerParams) player.GetPlayerResponder {
			return player.GetPlayerNotImplemented()
		})
	}
	if api.PlayerGetPlayerListHandler == nil {
		api.PlayerGetPlayerListHandler = player.GetPlayerListHandlerFunc(func(params player.GetPlayerListParams) player.GetPlayerListResponder {
			return player.GetPlayerListNotImplemented()
		})
	}
	if api.PlayerGetPlayerMatchesHandler == nil {
		api.PlayerGetPlayerMatchesHandler = player.GetPlayerMatchesHandlerFunc(func(params player.GetPlayerMatchesParams) player.GetPlayerMatchesResponder {
			return player.GetPlayerMatchesNotImplemented()
		})
	}
	if api.PlayerGetPlayerStatsHandler == nil {
		api.PlayerGetPlayerStatsHandler = player.GetPlayerStatsHandlerFunc(func(params player.GetPlayerStatsParams) player.GetPlayerStatsResponder {
			return player.GetPlayerStatsNotImplemented()
		})
	}
	if api.TeamGetTeamListHandler == nil {
		api.TeamGetTeamListHandler = team.GetTeamListHandlerFunc(func(params team.GetTeamListParams) team.GetTeamListResponder {
			return team.GetTeamListNotImplemented()
		})
	}
	if api.TeamGetTeamPlayersHandler == nil {
		api.TeamGetTeamPlayersHandler = team.GetTeamPlayersHandlerFunc(func(params team.GetTeamPlayersParams) team.GetTeamPlayersResponder {
			return team.GetTeamPlayersNotImplemented()
		})
	}
	if api.CompendiumGetWeaponClassesHandler == nil {
		api.CompendiumGetWeaponClassesHandler = compendium.GetWeaponClassesHandlerFunc(func(params compendium.GetWeaponClassesParams) compendium.GetWeaponClassesResponder {
			return compendium.GetWeaponClassesNotImplemented()
		})
	}
	if api.PlayerGetWeaponStatsHandler == nil {
		api.PlayerGetWeaponStatsHandler = player.GetWeaponStatsHandlerFunc(func(params player.GetWeaponStatsParams) player.GetWeaponStatsResponder {
			return player.GetWeaponStatsNotImplemented()
		})
	}
	if api.CompendiumGetWeaponsHandler == nil {
		api.CompendiumGetWeaponsHandler = compendium.GetWeaponsHandlerFunc(func(params compendium.GetWeaponsParams) compendium.GetWeaponsResponder {
			return compendium.GetWeaponsNotImplemented()
		})
	}
	if api.TeamSetTeamCaptainHandler == nil {
		api.TeamSetTeamCaptainHandler = team.SetTeamCaptainHandlerFunc(func(params team.SetTeamCaptainParams) team.SetTeamCaptainResponder {
			return team.SetTeamCaptainNotImplemented()
		})
	}
	if api.PlayerUpdatePlayerHandler == nil {
		api.PlayerUpdatePlayerHandler = player.UpdatePlayerHandlerFunc(func(params player.UpdatePlayerParams) player.UpdatePlayerResponder {
			return player.UpdatePlayerNotImplemented()
		})
	}
	if api.TeamUpdateTeamHandler == nil {
		api.TeamUpdateTeamHandler = team.UpdateTeamHandlerFunc(func(params team.UpdateTeamParams) team.UpdateTeamResponder {
			return team.UpdateTeamNotImplemented()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
