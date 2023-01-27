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

	// calculated stats
	// Required: true
	CalculatedStats PlayerStatsCalculatedStats `json:"calculated_stats"`

	// round stats
	// Required: true
	RoundStats PlayerStatsRoundStats `json:"round_stats"`

	// total stats
	// Required: true
	TotalStats *PlayerStatsTotalStats `json:"total_stats"`
}

// Validate validates this player stats
func (m *PlayerStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalculatedStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoundStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerStats) validateCalculatedStats(formats strfmt.Registry) error {

	if err := m.CalculatedStats.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("calculated_stats")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("calculated_stats")
		}
		return err
	}

	return nil
}

func (m *PlayerStats) validateRoundStats(formats strfmt.Registry) error {

	if err := m.RoundStats.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("round_stats")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("round_stats")
		}
		return err
	}

	return nil
}

func (m *PlayerStats) validateTotalStats(formats strfmt.Registry) error {

	if err := validate.Required("total_stats", "body", m.TotalStats); err != nil {
		return err
	}

	if m.TotalStats != nil {
		if err := m.TotalStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total_stats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this player stats based on the context it is used
func (m *PlayerStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCalculatedStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoundStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerStats) contextValidateCalculatedStats(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CalculatedStats.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("calculated_stats")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("calculated_stats")
		}
		return err
	}

	return nil
}

func (m *PlayerStats) contextValidateRoundStats(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RoundStats.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("round_stats")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("round_stats")
		}
		return err
	}

	return nil
}

func (m *PlayerStats) contextValidateTotalStats(ctx context.Context, formats strfmt.Registry) error {

	if m.TotalStats != nil {
		if err := m.TotalStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total_stats")
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
