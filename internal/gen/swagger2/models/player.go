// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Player player
//
// swagger:model Player
type Player struct {

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// steam id
	SteamID string `json:"steam_id,omitempty"`

	// team id
	TeamID int32 `json:"team_id,omitempty"`
}

// Validate validates this player
func (m *Player) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this player based on context it is used
func (m *Player) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Player) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Player) UnmarshalBinary(b []byte) error {
	var res Player
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}