// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSetTeamCaptainParams creates a new SetTeamCaptainParams object
//
// There are no default values defined in the spec.
func NewSetTeamCaptainParams() SetTeamCaptainParams {

	return SetTeamCaptainParams{}
}

// SetTeamCaptainParams contains all the bound params for the set team captain operation
// typically these are obtained from a http.Request
//
// swagger:parameters setTeamCaptain
type SetTeamCaptainParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Steam ID игрока
	  Required: true
	  In: path
	*/
	SteamID string
	/*ID команды
	  Required: true
	  In: path
	*/
	TeamID int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSetTeamCaptainParams() beforehand.
func (o *SetTeamCaptainParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSteamID, rhkSteamID, _ := route.Params.GetOK("steam_id")
	if err := o.bindSteamID(rSteamID, rhkSteamID, route.Formats); err != nil {
		res = append(res, err)
	}

	rTeamID, rhkTeamID, _ := route.Params.GetOK("team_id")
	if err := o.bindTeamID(rTeamID, rhkTeamID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSteamID binds and validates parameter SteamID from path.
func (o *SetTeamCaptainParams) bindSteamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.SteamID = raw

	return nil
}

// bindTeamID binds and validates parameter TeamID from path.
func (o *SetTeamCaptainParams) bindTeamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("team_id", "path", "int32", raw)
	}
	o.TeamID = value

	return nil
}