// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MatchTeamScoreboard match team scoreboard
//
// swagger:model MatchTeam_scoreboard
type MatchTeamScoreboard struct {

	// assists
	Assists int32 `json:"assists,omitempty"`

	// damage per round
	DamagePerRound float64 `json:"damage_per_round,omitempty"`

	// deaths
	Deaths int32 `json:"deaths,omitempty"`

	// headshot percentage
	HeadshotPercentage float64 `json:"headshot_percentage,omitempty"`

	// kill death ratio
	KillDeathRatio float64 `json:"kill_death_ratio,omitempty"`

	// kills
	Kills int32 `json:"kills,omitempty"`

	// mvps
	Mvps int32 `json:"mvps,omitempty"`

	// player name
	PlayerName string `json:"player_name,omitempty"`

	// steam id
	SteamID string `json:"steam_id,omitempty"`
}

// Validate validates this match team scoreboard
func (m *MatchTeamScoreboard) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this match team scoreboard based on context it is used
func (m *MatchTeamScoreboard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MatchTeamScoreboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchTeamScoreboard) UnmarshalBinary(b []byte) error {
	var res MatchTeamScoreboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}