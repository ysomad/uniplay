// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlayerList player list
//
// swagger:model PlayerList
type PlayerList struct {

	// has next
	// Required: true
	HasNext bool `json:"has_next"`

	// players
	// Required: true
	Players []PlayerListItem `json:"players"`
}

// Validate validates this player list
func (m *PlayerList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerList) validateHasNext(formats strfmt.Registry) error {

	if err := validate.Required("has_next", "body", bool(m.HasNext)); err != nil {
		return err
	}

	return nil
}

func (m *PlayerList) validatePlayers(formats strfmt.Registry) error {

	if err := validate.Required("players", "body", m.Players); err != nil {
		return err
	}

	for i := 0; i < len(m.Players); i++ {

		if err := m.Players[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("players" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("players" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this player list based on the context it is used
func (m *PlayerList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlayers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerList) contextValidatePlayers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Players); i++ {

		if err := m.Players[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("players" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("players" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayerList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerList) UnmarshalBinary(b []byte) error {
	var res PlayerList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}