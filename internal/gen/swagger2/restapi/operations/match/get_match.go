// Code generated by go-swagger; DO NOT EDIT.

package match

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMatchHandlerFunc turns a function with the right signature into a get match handler
type GetMatchHandlerFunc func(GetMatchParams) GetMatchResponder

// Handle executing the request and returning a response
func (fn GetMatchHandlerFunc) Handle(params GetMatchParams) GetMatchResponder {
	return fn(params)
}

// GetMatchHandler interface for that can handle valid get match params
type GetMatchHandler interface {
	Handle(GetMatchParams) GetMatchResponder
}

// NewGetMatch creates a new http.Handler for the get match operation
func NewGetMatch(ctx *middleware.Context, handler GetMatchHandler) *GetMatch {
	return &GetMatch{Context: ctx, Handler: handler}
}

/*
	GetMatch swagger:route GET /matches/{match_id} match getMatch

Получение информации о матче
*/
type GetMatch struct {
	Context *middleware.Context
	Handler GetMatchHandler
}

func (o *GetMatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMatchParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
