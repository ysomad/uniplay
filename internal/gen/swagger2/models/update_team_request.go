// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateTeamRequest update team request
//
// swagger:model UpdateTeamRequest
type UpdateTeamRequest struct {

	// clan name
	// Required: true
	// Max Length: 16
	// Min Length: 2
	ClanName string `json:"clan_name"`

	// flag code
	// Required: true
	// Max Length: 2
	// Min Length: 2
	FlagCode string `json:"flag_code"`

	// institution id
	// Required: true
	// Minimum: 1
	InstitutionID int32 `json:"institution_id"`
}

// Validate validates this update team request
func (m *UpdateTeamRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClanName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlagCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitutionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTeamRequest) validateClanName(formats strfmt.Registry) error {

	if err := validate.RequiredString("clan_name", "body", m.ClanName); err != nil {
		return err
	}

	if err := validate.MinLength("clan_name", "body", m.ClanName, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("clan_name", "body", m.ClanName, 16); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTeamRequest) validateFlagCode(formats strfmt.Registry) error {

	if err := validate.RequiredString("flag_code", "body", m.FlagCode); err != nil {
		return err
	}

	if err := validate.MinLength("flag_code", "body", m.FlagCode, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("flag_code", "body", m.FlagCode, 2); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTeamRequest) validateInstitutionID(formats strfmt.Registry) error {

	if err := validate.Required("institution_id", "body", int32(m.InstitutionID)); err != nil {
		return err
	}

	if err := validate.MinimumInt("institution_id", "body", int64(m.InstitutionID), 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update team request based on context it is used
func (m *UpdateTeamRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTeamRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTeamRequest) UnmarshalBinary(b []byte) error {
	var res UpdateTeamRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
