// Code generated by go-swagger; DO NOT EDIT.

package institution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetInstitutionsHandlerFunc turns a function with the right signature into a get institutions handler
type GetInstitutionsHandlerFunc func(GetInstitutionsParams) GetInstitutionsResponder

// Handle executing the request and returning a response
func (fn GetInstitutionsHandlerFunc) Handle(params GetInstitutionsParams) GetInstitutionsResponder {
	return fn(params)
}

// GetInstitutionsHandler interface for that can handle valid get institutions params
type GetInstitutionsHandler interface {
	Handle(GetInstitutionsParams) GetInstitutionsResponder
}

// NewGetInstitutions creates a new http.Handler for the get institutions operation
func NewGetInstitutions(ctx *middleware.Context, handler GetInstitutionsHandler) *GetInstitutions {
	return &GetInstitutions{Context: ctx, Handler: handler}
}

/*
	GetInstitutions swagger:route GET /institutions institution getInstitutions

Получение списка учебных заведений
*/
type GetInstitutions struct {
	Context *middleware.Context
	Handler GetInstitutionsHandler
}

func (o *GetInstitutions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetInstitutionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}