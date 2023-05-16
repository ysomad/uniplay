// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMostPlayedMapsHandlerFunc turns a function with the right signature into a get most played maps handler
type GetMostPlayedMapsHandlerFunc func(GetMostPlayedMapsParams) GetMostPlayedMapsResponder

// Handle executing the request and returning a response
func (fn GetMostPlayedMapsHandlerFunc) Handle(params GetMostPlayedMapsParams) GetMostPlayedMapsResponder {
	return fn(params)
}

// GetMostPlayedMapsHandler interface for that can handle valid get most played maps params
type GetMostPlayedMapsHandler interface {
	Handle(GetMostPlayedMapsParams) GetMostPlayedMapsResponder
}

// NewGetMostPlayedMaps creates a new http.Handler for the get most played maps operation
func NewGetMostPlayedMaps(ctx *middleware.Context, handler GetMostPlayedMapsHandler) *GetMostPlayedMaps {
	return &GetMostPlayedMaps{Context: ctx, Handler: handler}
}

/*
	GetMostPlayedMaps swagger:route GET /players/{steam_id}/most-played-maps player getMostPlayedMaps

Получение списка карт сыгранных наибольшее кол-во раз
*/
type GetMostPlayedMaps struct {
	Context *middleware.Context
	Handler GetMostPlayedMapsHandler
}

func (o *GetMostPlayedMaps) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMostPlayedMapsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
