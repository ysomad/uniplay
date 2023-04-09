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
)

// InstitutionList institution list
//
// swagger:model InstitutionList
type InstitutionList struct {

	// has next
	HasNext bool `json:"has_next"`

	// institutions
	Institutions []InstitutionListItem `json:"institutions"`
}

// Validate validates this institution list
func (m *InstitutionList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstitutions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstitutionList) validateInstitutions(formats strfmt.Registry) error {
	if swag.IsZero(m.Institutions) { // not required
		return nil
	}

	for i := 0; i < len(m.Institutions); i++ {

		if err := m.Institutions[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("institutions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("institutions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this institution list based on the context it is used
func (m *InstitutionList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstitutions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstitutionList) contextValidateInstitutions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Institutions); i++ {

		if err := m.Institutions[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("institutions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("institutions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstitutionList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstitutionList) UnmarshalBinary(b []byte) error {
	var res InstitutionList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}