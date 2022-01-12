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

// CloudConnections cloud connections
//
// swagger:model CloudConnections
type CloudConnections struct {

	// Cloud Connections
	// Required: true
	CloudConnections []*CloudConnection `json:"cloudConnections"`
}

// Validate validates this cloud connections
func (m *CloudConnections) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnections) validateCloudConnections(formats strfmt.Registry) error {

	if err := validate.Required("cloudConnections", "body", m.CloudConnections); err != nil {
		return err
	}

	for i := 0; i < len(m.CloudConnections); i++ {
		if swag.IsZero(m.CloudConnections[i]) { // not required
			continue
		}

		if m.CloudConnections[i] != nil {
			if err := m.CloudConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudConnections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cloudConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cloud connections based on the context it is used
func (m *CloudConnections) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnections) contextValidateCloudConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CloudConnections); i++ {

		if m.CloudConnections[i] != nil {
			if err := m.CloudConnections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudConnections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cloudConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudConnections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudConnections) UnmarshalBinary(b []byte) error {
	var res CloudConnections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
