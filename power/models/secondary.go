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

// Secondary Information to create a secondary host group
//
// swagger:model Secondary
type Secondary struct {

	// Name of the host group to create in the secondary workspace
	Name string `json:"name,omitempty"`

	// ID of the workspace to share the host group with
	// Required: true
	Workspace *string `json:"workspace"`
}

// Validate validates this secondary
func (m *Secondary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secondary) validateWorkspace(formats strfmt.Registry) error {

	if err := validate.Required("workspace", "body", m.Workspace); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this secondary based on context it is used
func (m *Secondary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secondary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secondary) UnmarshalBinary(b []byte) error {
	var res Secondary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
