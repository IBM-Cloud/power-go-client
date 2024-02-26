// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkReference network reference
//
// swagger:model NetworkReference
type NetworkReference struct {

	// access config
	AccessConfig AccessConfig `json:"accessConfig,omitempty"`

	// DHCP Managed Network
	DhcpManaged bool `json:"dhcpManaged,omitempty"`

	// Link to Network resource
	// Required: true
	Href *string `json:"href"`

	// (deprecated - replaced by mtu) Enable MTU Jumbo Network (for multi-zone locations only)
	Jumbo bool `json:"jumbo,omitempty"`

	// Maximum transmission unit
	// Maximum: 9000
	// Minimum: 1450
	Mtu *int64 `json:"mtu,omitempty"`

	// Network Name
	// Required: true
	Name *string `json:"name"`

	// Unique Network ID
	// Required: true
	NetworkID *string `json:"networkID"`

	// Type of Network - 'vlan' (private network) 'pub-vlan' (public network) 'dhcp-vlan' (for satellite locations only)
	// Required: true
	// Enum: [vlan pub-vlan dhcp-vlan]
	Type *string `json:"type"`

	// VLAN ID
	// Required: true
	VlanID *float64 `json:"vlanID"`
}

// Validate validates this network reference
func (m *NetworkReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkReference) validateAccessConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessConfig) { // not required
		return nil
	}

	if err := m.AccessConfig.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accessConfig")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accessConfig")
		}
		return err
	}

	return nil
}

func (m *NetworkReference) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *NetworkReference) validateMtu(formats strfmt.Registry) error {
	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", *m.Mtu, 1450, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", *m.Mtu, 9000, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkReference) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NetworkReference) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("networkID", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

var networkReferenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vlan","pub-vlan","dhcp-vlan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkReferenceTypeTypePropEnum = append(networkReferenceTypeTypePropEnum, v)
	}
}

const (

	// NetworkReferenceTypeVlan captures enum value "vlan"
	NetworkReferenceTypeVlan string = "vlan"

	// NetworkReferenceTypePubDashVlan captures enum value "pub-vlan"
	NetworkReferenceTypePubDashVlan string = "pub-vlan"

	// NetworkReferenceTypeDhcpDashVlan captures enum value "dhcp-vlan"
	NetworkReferenceTypeDhcpDashVlan string = "dhcp-vlan"
)

// prop value enum
func (m *NetworkReference) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkReferenceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkReference) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *NetworkReference) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlanID", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this network reference based on the context it is used
func (m *NetworkReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkReference) contextValidateAccessConfig(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AccessConfig) { // not required
		return nil
	}

	if err := m.AccessConfig.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accessConfig")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accessConfig")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkReference) UnmarshalBinary(b []byte) error {
	var res NetworkReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
