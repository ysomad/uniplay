// Code generated by go-swagger; DO NOT EDIT.

package compendium

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetWeaponsHandlerFunc turns a function with the right signature into a get weapons handler
type GetWeaponsHandlerFunc func(GetWeaponsParams) GetWeaponsResponder

// Handle executing the request and returning a response
func (fn GetWeaponsHandlerFunc) Handle(params GetWeaponsParams) GetWeaponsResponder {
	return fn(params)
}

// GetWeaponsHandler interface for that can handle valid get weapons params
type GetWeaponsHandler interface {
	Handle(GetWeaponsParams) GetWeaponsResponder
}

// NewGetWeapons creates a new http.Handler for the get weapons operation
func NewGetWeapons(ctx *middleware.Context, handler GetWeaponsHandler) *GetWeapons {
	return &GetWeapons{Context: ctx, Handler: handler}
}

/*
	GetWeapons swagger:route GET /compendiums/weapons compendium getWeapons

Получение списка оружий
*/
type GetWeapons struct {
	Context *middleware.Context
	Handler GetWeaponsHandler
}

func (o *GetWeapons) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWeaponsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
