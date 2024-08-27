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

// NetworkSecurityGroupRule network security group rule
//
// swagger:model NetworkSecurityGroupRule
type NetworkSecurityGroupRule struct {

	// The action to take if the rule matches network traffic
	// Required: true
	// Enum: ["allow","deny"]
	Action *string `json:"action"`

	// destination port
	DestinationPort *NetworkSecurityGroupRulePort `json:"destinationPort,omitempty"`

	// The ID of the rule in a Network Security Group
	// Required: true
	ID *string `json:"id"`

	// protocol
	// Required: true
	Protocol *NetworkSecurityGroupRuleProtocol `json:"protocol"`

	// remote
	// Required: true
	Remote *NetworkSecurityGroupRuleRemote `json:"remote"`

	// source port
	SourcePort *NetworkSecurityGroupRulePort `json:"sourcePort,omitempty"`
}

// Validate validates this network security group rule
func (m *NetworkSecurityGroupRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkSecurityGroupRuleTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkSecurityGroupRuleTypeActionPropEnum = append(networkSecurityGroupRuleTypeActionPropEnum, v)
	}
}

const (

	// NetworkSecurityGroupRuleActionAllow captures enum value "allow"
	NetworkSecurityGroupRuleActionAllow string = "allow"

	// NetworkSecurityGroupRuleActionDeny captures enum value "deny"
	NetworkSecurityGroupRuleActionDeny string = "deny"
)

// prop value enum
func (m *NetworkSecurityGroupRule) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkSecurityGroupRuleTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkSecurityGroupRule) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSecurityGroupRule) validateDestinationPort(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationPort) { // not required
		return nil
	}

	if m.DestinationPort != nil {
		if err := m.DestinationPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationPort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationPort")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupRule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSecurityGroupRule) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if m.Protocol != nil {
		if err := m.Protocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupRule) validateRemote(formats strfmt.Registry) error {

	if err := validate.Required("remote", "body", m.Remote); err != nil {
		return err
	}

	if m.Remote != nil {
		if err := m.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remote")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupRule) validateSourcePort(formats strfmt.Registry) error {
	if swag.IsZero(m.SourcePort) { // not required
		return nil
	}

	if m.SourcePort != nil {
		if err := m.SourcePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourcePort")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network security group rule based on the context it is used
func (m *NetworkSecurityGroupRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationPort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourcePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSecurityGroupRule) contextValidateDestinationPort(ctx context.Context, formats strfmt.Registry) error {

	if m.DestinationPort != nil {

		if swag.IsZero(m.DestinationPort) { // not required
			return nil
		}

		if err := m.DestinationPort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationPort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationPort")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupRule) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.Protocol != nil {

		if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupRule) contextValidateRemote(ctx context.Context, formats strfmt.Registry) error {

	if m.Remote != nil {

		if err := m.Remote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remote")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupRule) contextValidateSourcePort(ctx context.Context, formats strfmt.Registry) error {

	if m.SourcePort != nil {

		if swag.IsZero(m.SourcePort) { // not required
			return nil
		}

		if err := m.SourcePort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourcePort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSecurityGroupRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSecurityGroupRule) UnmarshalBinary(b []byte) error {
	var res NetworkSecurityGroupRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
