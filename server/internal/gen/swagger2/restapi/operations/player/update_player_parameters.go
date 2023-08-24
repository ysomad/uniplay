// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
)

// NewUpdatePlayerParams creates a new UpdatePlayerParams object
//
// There are no default values defined in the spec.
func NewUpdatePlayerParams() UpdatePlayerParams {

	return UpdatePlayerParams{}
}

// UpdatePlayerParams contains all the bound params for the update player operation
// typically these are obtained from a http.Request
//
// swagger:parameters updatePlayer
type UpdatePlayerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Payload *models.UpdatePlayerRequest
	/*Steam ID игрока
	  Required: true
	  In: path
	*/
	SteamID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdatePlayerParams() beforehand.
func (o *UpdatePlayerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UpdatePlayerRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("payload", "body", ""))
			} else {
				res = append(res, errors.NewParseError("payload", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Payload = &body
			}
		}
	} else {
		res = append(res, errors.Required("payload", "body", ""))
	}

	rSteamID, rhkSteamID, _ := route.Params.GetOK("steam_id")
	if err := o.bindSteamID(rSteamID, rhkSteamID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSteamID binds and validates parameter SteamID from path.
func (o *UpdatePlayerParams) bindSteamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.SteamID = raw

	return nil
}