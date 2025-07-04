// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateServerVirtualSerialNumber update server virtual serial number
//
// swagger:model UpdateServerVirtualSerialNumber
type UpdateServerVirtualSerialNumber struct {

	// Description of the Virtual Serial Number
	Description *string `json:"description,omitempty"`

	// software tier
	SoftwareTier SoftwareTier `json:"softwareTier,omitempty"`
}

// Validate validates this update server virtual serial number
func (m *UpdateServerVirtualSerialNumber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSoftwareTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateServerVirtualSerialNumber) validateSoftwareTier(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareTier) { // not required
		return nil
	}

	if err := m.SoftwareTier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("softwareTier")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("softwareTier")
		}
		return err
	}

	return nil
}

// ContextValidate validate this update server virtual serial number based on the context it is used
func (m *UpdateServerVirtualSerialNumber) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSoftwareTier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateServerVirtualSerialNumber) contextValidateSoftwareTier(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SoftwareTier) { // not required
		return nil
	}

	if err := m.SoftwareTier.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("softwareTier")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("softwareTier")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateServerVirtualSerialNumber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateServerVirtualSerialNumber) UnmarshalBinary(b []byte) error {
	var res UpdateServerVirtualSerialNumber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
