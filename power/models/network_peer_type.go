// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NetworkPeerType Type of the network peer
//
// swagger:model NetworkPeerType
type NetworkPeerType string

func NewNetworkPeerType(value NetworkPeerType) *NetworkPeerType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NetworkPeerType.
func (m NetworkPeerType) Pointer() *NetworkPeerType {
	return &m
}

const (

	// NetworkPeerTypeL2 captures enum value "L2"
	NetworkPeerTypeL2 NetworkPeerType = "L2"

	// NetworkPeerTypeL3BGP captures enum value "L3BGP"
	NetworkPeerTypeL3BGP NetworkPeerType = "L3BGP"

	// NetworkPeerTypeL3Static captures enum value "L3Static"
	NetworkPeerTypeL3Static NetworkPeerType = "L3Static"
)

// for schema
var networkPeerTypeEnum []interface{}

func init() {
	var res []NetworkPeerType
	if err := json.Unmarshal([]byte(`["L2","L3BGP","L3Static"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkPeerTypeEnum = append(networkPeerTypeEnum, v)
	}
}

func (m NetworkPeerType) validateNetworkPeerTypeEnum(path, location string, value NetworkPeerType) error {
	if err := validate.EnumCase(path, location, value, networkPeerTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this network peer type
func (m NetworkPeerType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNetworkPeerTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this network peer type based on context it is used
func (m NetworkPeerType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}