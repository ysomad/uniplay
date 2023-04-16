// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/account"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/compendium"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/institution"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/match"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/player"
)

// NewUniplayAPI creates a new Uniplay instance
func NewUniplayAPI(spec *loads.Document) *UniplayAPI {
	return &UniplayAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		AccountCreateAccountHandler: account.CreateAccountHandlerFunc(func(params account.CreateAccountParams) account.CreateAccountResponder {
			return account.CreateAccountNotImplemented()
		}),
		MatchCreateMatchHandler: match.CreateMatchHandlerFunc(func(params match.CreateMatchParams) match.CreateMatchResponder {
			return match.CreateMatchNotImplemented()
		}),
		MatchDeleteMatchHandler: match.DeleteMatchHandlerFunc(func(params match.DeleteMatchParams) match.DeleteMatchResponder {
			return match.DeleteMatchNotImplemented()
		}),
		CompendiumGetCitiesHandler: compendium.GetCitiesHandlerFunc(func(params compendium.GetCitiesParams) compendium.GetCitiesResponder {
			return compendium.GetCitiesNotImplemented()
		}),
		InstitutionGetInstitutionsHandler: institution.GetInstitutionsHandlerFunc(func(params institution.GetInstitutionsParams) institution.GetInstitutionsResponder {
			return institution.GetInstitutionsNotImplemented()
		}),
		CompendiumGetMapsHandler: compendium.GetMapsHandlerFunc(func(params compendium.GetMapsParams) compendium.GetMapsResponder {
			return compendium.GetMapsNotImplemented()
		}),
		MatchGetMatchHandler: match.GetMatchHandlerFunc(func(params match.GetMatchParams) match.GetMatchResponder {
			return match.GetMatchNotImplemented()
		}),
		PlayerGetPlayerHandler: player.GetPlayerHandlerFunc(func(params player.GetPlayerParams) player.GetPlayerResponder {
			return player.GetPlayerNotImplemented()
		}),
		PlayerGetPlayerMatchesHandler: player.GetPlayerMatchesHandlerFunc(func(params player.GetPlayerMatchesParams) player.GetPlayerMatchesResponder {
			return player.GetPlayerMatchesNotImplemented()
		}),
		PlayerGetPlayerStatsHandler: player.GetPlayerStatsHandlerFunc(func(params player.GetPlayerStatsParams) player.GetPlayerStatsResponder {
			return player.GetPlayerStatsNotImplemented()
		}),
		CompendiumGetWeaponClassesHandler: compendium.GetWeaponClassesHandlerFunc(func(params compendium.GetWeaponClassesParams) compendium.GetWeaponClassesResponder {
			return compendium.GetWeaponClassesNotImplemented()
		}),
		PlayerGetWeaponStatsHandler: player.GetWeaponStatsHandlerFunc(func(params player.GetWeaponStatsParams) player.GetWeaponStatsResponder {
			return player.GetWeaponStatsNotImplemented()
		}),
		CompendiumGetWeaponsHandler: compendium.GetWeaponsHandlerFunc(func(params compendium.GetWeaponsParams) compendium.GetWeaponsResponder {
			return compendium.GetWeaponsNotImplemented()
		}),
		PlayerUpdatePlayerHandler: player.UpdatePlayerHandlerFunc(func(params player.UpdatePlayerParams) player.UpdatePlayerResponder {
			return player.UpdatePlayerNotImplemented()
		}),
	}
}

/*UniplayAPI the uniplay API */
type UniplayAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AccountCreateAccountHandler sets the operation handler for the create account operation
	AccountCreateAccountHandler account.CreateAccountHandler
	// MatchCreateMatchHandler sets the operation handler for the create match operation
	MatchCreateMatchHandler match.CreateMatchHandler
	// MatchDeleteMatchHandler sets the operation handler for the delete match operation
	MatchDeleteMatchHandler match.DeleteMatchHandler
	// CompendiumGetCitiesHandler sets the operation handler for the get cities operation
	CompendiumGetCitiesHandler compendium.GetCitiesHandler
	// InstitutionGetInstitutionsHandler sets the operation handler for the get institutions operation
	InstitutionGetInstitutionsHandler institution.GetInstitutionsHandler
	// CompendiumGetMapsHandler sets the operation handler for the get maps operation
	CompendiumGetMapsHandler compendium.GetMapsHandler
	// MatchGetMatchHandler sets the operation handler for the get match operation
	MatchGetMatchHandler match.GetMatchHandler
	// PlayerGetPlayerHandler sets the operation handler for the get player operation
	PlayerGetPlayerHandler player.GetPlayerHandler
	// PlayerGetPlayerMatchesHandler sets the operation handler for the get player matches operation
	PlayerGetPlayerMatchesHandler player.GetPlayerMatchesHandler
	// PlayerGetPlayerStatsHandler sets the operation handler for the get player stats operation
	PlayerGetPlayerStatsHandler player.GetPlayerStatsHandler
	// CompendiumGetWeaponClassesHandler sets the operation handler for the get weapon classes operation
	CompendiumGetWeaponClassesHandler compendium.GetWeaponClassesHandler
	// PlayerGetWeaponStatsHandler sets the operation handler for the get weapon stats operation
	PlayerGetWeaponStatsHandler player.GetWeaponStatsHandler
	// CompendiumGetWeaponsHandler sets the operation handler for the get weapons operation
	CompendiumGetWeaponsHandler compendium.GetWeaponsHandler
	// PlayerUpdatePlayerHandler sets the operation handler for the update player operation
	PlayerUpdatePlayerHandler player.UpdatePlayerHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *UniplayAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *UniplayAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *UniplayAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *UniplayAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *UniplayAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *UniplayAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *UniplayAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *UniplayAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *UniplayAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the UniplayAPI
func (o *UniplayAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AccountCreateAccountHandler == nil {
		unregistered = append(unregistered, "account.CreateAccountHandler")
	}
	if o.MatchCreateMatchHandler == nil {
		unregistered = append(unregistered, "match.CreateMatchHandler")
	}
	if o.MatchDeleteMatchHandler == nil {
		unregistered = append(unregistered, "match.DeleteMatchHandler")
	}
	if o.CompendiumGetCitiesHandler == nil {
		unregistered = append(unregistered, "compendium.GetCitiesHandler")
	}
	if o.InstitutionGetInstitutionsHandler == nil {
		unregistered = append(unregistered, "institution.GetInstitutionsHandler")
	}
	if o.CompendiumGetMapsHandler == nil {
		unregistered = append(unregistered, "compendium.GetMapsHandler")
	}
	if o.MatchGetMatchHandler == nil {
		unregistered = append(unregistered, "match.GetMatchHandler")
	}
	if o.PlayerGetPlayerHandler == nil {
		unregistered = append(unregistered, "player.GetPlayerHandler")
	}
	if o.PlayerGetPlayerMatchesHandler == nil {
		unregistered = append(unregistered, "player.GetPlayerMatchesHandler")
	}
	if o.PlayerGetPlayerStatsHandler == nil {
		unregistered = append(unregistered, "player.GetPlayerStatsHandler")
	}
	if o.CompendiumGetWeaponClassesHandler == nil {
		unregistered = append(unregistered, "compendium.GetWeaponClassesHandler")
	}
	if o.PlayerGetWeaponStatsHandler == nil {
		unregistered = append(unregistered, "player.GetWeaponStatsHandler")
	}
	if o.CompendiumGetWeaponsHandler == nil {
		unregistered = append(unregistered, "compendium.GetWeaponsHandler")
	}
	if o.PlayerUpdatePlayerHandler == nil {
		unregistered = append(unregistered, "player.UpdatePlayerHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *UniplayAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *UniplayAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *UniplayAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *UniplayAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *UniplayAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *UniplayAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the uniplay API
func (o *UniplayAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *UniplayAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/accounts"] = account.NewCreateAccount(o.context, o.AccountCreateAccountHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/matches"] = match.NewCreateMatch(o.context, o.MatchCreateMatchHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/matches/{match_id}"] = match.NewDeleteMatch(o.context, o.MatchDeleteMatchHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/compendiums/cities"] = compendium.NewGetCities(o.context, o.CompendiumGetCitiesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/institutions"] = institution.NewGetInstitutions(o.context, o.InstitutionGetInstitutionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/compendiums/maps"] = compendium.NewGetMaps(o.context, o.CompendiumGetMapsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/matches/{match_id}"] = match.NewGetMatch(o.context, o.MatchGetMatchHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/players/{steam_id}"] = player.NewGetPlayer(o.context, o.PlayerGetPlayerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/players/{steam_id}/matches"] = player.NewGetPlayerMatches(o.context, o.PlayerGetPlayerMatchesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/players/{steam_id}/stats"] = player.NewGetPlayerStats(o.context, o.PlayerGetPlayerStatsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/compendiums/weapon-classes"] = compendium.NewGetWeaponClasses(o.context, o.CompendiumGetWeaponClassesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/players/{steam_id}/weapons"] = player.NewGetWeaponStats(o.context, o.PlayerGetWeaponStatsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/compendiums/weapons"] = compendium.NewGetWeapons(o.context, o.CompendiumGetWeaponsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/players/{steam_id}"] = player.NewUpdatePlayer(o.context, o.PlayerUpdatePlayerHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *UniplayAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *UniplayAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *UniplayAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *UniplayAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *UniplayAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
