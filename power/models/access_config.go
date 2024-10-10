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

// AccessConfig (deprecated - replaced by network peer id)
// Network communication configuration (for satellite locations only)
//   - `internal-only` - network is only used for internal host communication
//   - `outbound-only` - network will be capable of egress traffic
//   - `bidirectional-static-route` - network will be capable of ingress and egress traffic via static routes
//   - `bidirectional-bgp` - network will be capable of ingress and egress traffic via bgp configuration
//   - `bidirectional-l2out` - network will be capable of ingress and egress traffic via l2out ACI configuration
//
// swagger:model AccessConfig
type AccessConfig string

func NewAccessConfig(value AccessConfig) *AccessConfig {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AccessConfig.
func (m AccessConfig) Pointer() *AccessConfig {
	return &m
}

const (

	// AccessConfigInternalDashOnly captures enum value "internal-only"
	AccessConfigInternalDashOnly AccessConfig = "internal-only"

	// AccessConfigOutboundDashOnly captures enum value "outbound-only"
	AccessConfigOutboundDashOnly AccessConfig = "outbound-only"

	// AccessConfigBidirectionalDashStaticDashRoute captures enum value "bidirectional-static-route"
	AccessConfigBidirectionalDashStaticDashRoute AccessConfig = "bidirectional-static-route"

	// AccessConfigBidirectionalDashBgp captures enum value "bidirectional-bgp"
	AccessConfigBidirectionalDashBgp AccessConfig = "bidirectional-bgp"

	// AccessConfigBidirectionalDashL2out captures enum value "bidirectional-l2out"
	AccessConfigBidirectionalDashL2out AccessConfig = "bidirectional-l2out"
)

// for schema
var accessConfigEnum []interface{}

func init() {
	var res []AccessConfig
	if err := json.Unmarshal([]byte(`["internal-only","outbound-only","bidirectional-static-route","bidirectional-bgp","bidirectional-l2out"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessConfigEnum = append(accessConfigEnum, v)
	}
}

func (m AccessConfig) validateAccessConfigEnum(path, location string, value AccessConfig) error {
	if err := validate.EnumCase(path, location, value, accessConfigEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this access config
func (m AccessConfig) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccessConfigEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this access config based on context it is used
func (m AccessConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
