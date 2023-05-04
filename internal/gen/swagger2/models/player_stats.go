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

// PlayerStats player stats
//
// swagger:model PlayerStats
type PlayerStats struct {

	// base stats
	// Required: true
	BaseStats *PlayerStatsBaseStats `json:"base_stats"`

	// calculated stats
	// Required: true
	CalculatedStats *PlayerStatsCalculatedStats `json:"calculated_stats"`
}

// Validate validates this player stats
func (m *PlayerStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalculatedStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerStats) validateBaseStats(formats strfmt.Registry) error {

	if err := validate.Required("base_stats", "body", m.BaseStats); err != nil {
		return err
	}

	if m.BaseStats != nil {
		if err := m.BaseStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("base_stats")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerStats) validateCalculatedStats(formats strfmt.Registry) error {

	if err := validate.Required("calculated_stats", "body", m.CalculatedStats); err != nil {
		return err
	}

	if m.CalculatedStats != nil {
		if err := m.CalculatedStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calculated_stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("calculated_stats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this player stats based on the context it is used
func (m *PlayerStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCalculatedStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerStats) contextValidateBaseStats(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseStats != nil {
		if err := m.BaseStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("base_stats")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerStats) contextValidateCalculatedStats(ctx context.Context, formats strfmt.Registry) error {

	if m.CalculatedStats != nil {
		if err := m.CalculatedStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calculated_stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("calculated_stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayerStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerStats) UnmarshalBinary(b []byte) error {
	var res PlayerStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
