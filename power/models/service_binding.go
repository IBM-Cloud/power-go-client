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

// ServiceBinding service binding
//
// swagger:model ServiceBinding
type ServiceBinding struct {

	// credentials
	Credentials Object `json:"credentials,omitempty"`

	// route service url
	RouteServiceURL string `json:"route_service_url,omitempty"`

	// syslog drain url
	SyslogDrainURL string `json:"syslog_drain_url,omitempty"`

	// volume mounts
	VolumeMounts []*ServiceBindingVolumeMount `json:"volume_mounts"`
}

// Validate validates this service binding
func (m *ServiceBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumeMounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceBinding) validateVolumeMounts(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeMounts) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeMounts); i++ {
		if swag.IsZero(m.VolumeMounts[i]) { // not required
			continue
		}

		if m.VolumeMounts[i] != nil {
			if err := m.VolumeMounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volume_mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volume_mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service binding based on the context it is used
func (m *ServiceBinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVolumeMounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceBinding) contextValidateVolumeMounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeMounts); i++ {

		if m.VolumeMounts[i] != nil {

			if swag.IsZero(m.VolumeMounts[i]) { // not required
				return nil
			}

			if err := m.VolumeMounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volume_mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volume_mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceBinding) UnmarshalBinary(b []byte) error {
	var res ServiceBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
