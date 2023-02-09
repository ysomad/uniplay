// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetWeaponStatsParams creates a new GetWeaponStatsParams object
//
// There are no default values defined in the spec.
func NewGetWeaponStatsParams() GetWeaponStatsParams {

	return GetWeaponStatsParams{}
}

// GetWeaponStatsParams contains all the bound params for the get weapon stats operation
// typically these are obtained from a http.Request
//
// swagger:parameters getWeaponStats
type GetWeaponStatsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Фильтр по классу оружия
	  In: query
	*/
	ClassID *int16
	/*Фильтр по матчу
	  In: query
	*/
	MatchID *strfmt.UUID
	/*Steam ID игрока
	  Required: true
	  In: path
	*/
	SteamID string
	/*Фильтр по оружию
	  In: query
	*/
	WeaponID *int16
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetWeaponStatsParams() beforehand.
func (o *GetWeaponStatsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qClassID, qhkClassID, _ := qs.GetOK("class_id")
	if err := o.bindClassID(qClassID, qhkClassID, route.Formats); err != nil {
		res = append(res, err)
	}

	qMatchID, qhkMatchID, _ := qs.GetOK("match_id")
	if err := o.bindMatchID(qMatchID, qhkMatchID, route.Formats); err != nil {
		res = append(res, err)
	}

	rSteamID, rhkSteamID, _ := route.Params.GetOK("steam_id")
	if err := o.bindSteamID(rSteamID, rhkSteamID, route.Formats); err != nil {
		res = append(res, err)
	}

	qWeaponID, qhkWeaponID, _ := qs.GetOK("weapon_id")
	if err := o.bindWeaponID(qWeaponID, qhkWeaponID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClassID binds and validates parameter ClassID from query.
func (o *GetWeaponStatsParams) bindClassID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt16(raw)
	if err != nil {
		return errors.InvalidType("class_id", "query", "int16", raw)
	}
	o.ClassID = &value

	return nil
}

// bindMatchID binds and validates parameter MatchID from query.
func (o *GetWeaponStatsParams) bindMatchID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("match_id", "query", "strfmt.UUID", raw)
	}
	o.MatchID = (value.(*strfmt.UUID))

	if err := o.validateMatchID(formats); err != nil {
		return err
	}

	return nil
}

// validateMatchID carries on validations for parameter MatchID
func (o *GetWeaponStatsParams) validateMatchID(formats strfmt.Registry) error {

	if err := validate.FormatOf("match_id", "query", "uuid", o.MatchID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindSteamID binds and validates parameter SteamID from path.
func (o *GetWeaponStatsParams) bindSteamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.SteamID = raw

	return nil
}

// bindWeaponID binds and validates parameter WeaponID from query.
func (o *GetWeaponStatsParams) bindWeaponID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt16(raw)
	if err != nil {
		return errors.InvalidType("weapon_id", "query", "int16", raw)
	}
	o.WeaponID = &value

	return nil
}
