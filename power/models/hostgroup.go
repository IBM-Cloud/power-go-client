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

// Hostgroup Description of a hostgroup
//
// swagger:model Hostgroup
type Hostgroup struct {

	// Date/Time of hostgroup creation
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// List of hosts
	Hosts []HostHref `json:"hosts"`

	// Hostgroup ID
	ID string `json:"id,omitempty"`

	// Name of the hostgroup
	Name string `json:"name,omitempty"`

	// Name of the workspace owning the hostgroup
	Primary string `json:"primary,omitempty"`

	// Names of workspaces the hostgroup has been shared with
	Secondaries []string `json:"secondaries"`
}

// Validate validates this hostgroup
func (m *Hostgroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hostgroup) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Hostgroup) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	for i := 0; i < len(m.Hosts); i++ {

		if err := m.Hosts[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this hostgroup based on the context it is used
func (m *Hostgroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hostgroup) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hosts); i++ {

		if swag.IsZero(m.Hosts[i]) { // not required
			return nil
		}

		if err := m.Hosts[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Hostgroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hostgroup) UnmarshalBinary(b []byte) error {
	var res Hostgroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
