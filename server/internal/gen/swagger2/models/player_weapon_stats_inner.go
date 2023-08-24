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

// PlayerWeaponStatsInner player weapon stats inner
//
// swagger:model PlayerWeaponStats_inner
type PlayerWeaponStatsInner struct {

	// accuracy stats
	// Required: true
	AccuracyStats *PlayerWeaponStatsInnerAccuracyStats `json:"accuracy_stats"`

	// base stats
	// Required: true
	BaseStats *PlayerWeaponStatsInnerBaseStats `json:"base_stats"`

	// weapon
	Weapon string `json:"weapon,omitempty"`

	// weapon id
	WeaponID int16 `json:"weapon_id,omitempty"`
}

// Validate validates this player weapon stats inner
func (m *PlayerWeaponStatsInner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccuracyStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerWeaponStatsInner) validateAccuracyStats(formats strfmt.Registry) error {

	if err := validate.Required("accuracy_stats", "body", m.AccuracyStats); err != nil {
		return err
	}

	if m.AccuracyStats != nil {
		if err := m.AccuracyStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accuracy_stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accuracy_stats")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerWeaponStatsInner) validateBaseStats(formats strfmt.Registry) error {

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

// ContextValidate validate this player weapon stats inner based on the context it is used
func (m *PlayerWeaponStatsInner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccuracyStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaseStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerWeaponStatsInner) contextValidateAccuracyStats(ctx context.Context, formats strfmt.Registry) error {

	if m.AccuracyStats != nil {
		if err := m.AccuracyStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accuracy_stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accuracy_stats")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerWeaponStatsInner) contextValidateBaseStats(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PlayerWeaponStatsInner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerWeaponStatsInner) UnmarshalBinary(b []byte) error {
	var res PlayerWeaponStatsInner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}