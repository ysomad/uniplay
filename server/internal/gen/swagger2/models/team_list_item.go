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

// TeamListItem team list item
//
// swagger:model TeamListItem
type TeamListItem struct {

	// clan name
	// Required: true
	ClanName string `json:"clan_name"`

	// flag code
	FlagCode string `json:"flag_code,omitempty"`

	// id
	// Required: true
	ID int32 `json:"id"`

	// institution
	Institution *TeamListInstitution `json:"institution,omitempty"`
}

// Validate validates this team list item
func (m *TeamListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClanName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamListItem) validateClanName(formats strfmt.Registry) error {

	if err := validate.RequiredString("clan_name", "body", m.ClanName); err != nil {
		return err
	}

	return nil
}

func (m *TeamListItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *TeamListItem) validateInstitution(formats strfmt.Registry) error {
	if swag.IsZero(m.Institution) { // not required
		return nil
	}

	if m.Institution != nil {
		if err := m.Institution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("institution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("institution")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this team list item based on the context it is used
func (m *TeamListItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstitution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamListItem) contextValidateInstitution(ctx context.Context, formats strfmt.Registry) error {

	if m.Institution != nil {
		if err := m.Institution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("institution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("institution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamListItem) UnmarshalBinary(b []byte) error {
	var res TeamListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}