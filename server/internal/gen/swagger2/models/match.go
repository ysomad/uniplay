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

// Match match
//
// swagger:model Match
type Match struct {

	// длительность матча в минутах
	Duration int32 `json:"duration,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// map
	Map Map `json:"map,omitempty"`

	// rounds played
	RoundsPlayed int32 `json:"rounds_played,omitempty"`

	// team1
	Team1 MatchTeam `json:"team1,omitempty"`

	// team2
	Team2 MatchTeam `json:"team2,omitempty"`

	// uploaded at
	// Format: date-time
	UploadedAt strfmt.DateTime `json:"uploaded_at,omitempty"`
}

// Validate validates this match
func (m *Match) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Match) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Match) validateMap(formats strfmt.Registry) error {
	if swag.IsZero(m.Map) { // not required
		return nil
	}

	if err := m.Map.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("map")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("map")
		}
		return err
	}

	return nil
}

func (m *Match) validateTeam1(formats strfmt.Registry) error {
	if swag.IsZero(m.Team1) { // not required
		return nil
	}

	if err := m.Team1.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("team1")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("team1")
		}
		return err
	}

	return nil
}

func (m *Match) validateTeam2(formats strfmt.Registry) error {
	if swag.IsZero(m.Team2) { // not required
		return nil
	}

	if err := m.Team2.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("team2")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("team2")
		}
		return err
	}

	return nil
}

func (m *Match) validateUploadedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("uploaded_at", "body", "date-time", m.UploadedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this match based on the context it is used
func (m *Match) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Match) contextValidateMap(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Map.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("map")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("map")
		}
		return err
	}

	return nil
}

func (m *Match) contextValidateTeam1(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Team1.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("team1")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("team1")
		}
		return err
	}

	return nil
}

func (m *Match) contextValidateTeam2(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Team2.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("team2")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("team2")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Match) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Match) UnmarshalBinary(b []byte) error {
	var res Match
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}