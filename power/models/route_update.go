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

// RouteUpdate route update
//
// swagger:model RouteUpdate
type RouteUpdate struct {

	// Action
	// Enum: ["deliver"]
	Action *string `json:"action,omitempty"`

	// Indicates if the route is advertised externally
	AdvertiseExternally *bool `json:"advertiseExternally,omitempty"`

	// The route destination
	// Required: true
	Destination *string `json:"destination"`

	// The destination type
	// Enum: ["ipv4-address"]
	DestinationType *string `json:"destinationType,omitempty"`

	// Name of the route
	// Required: true
	Name interface{} `json:"name"`

	// The next hop
	// Required: true
	NextHop *string `json:"nextHop"`

	// The next hop type
	// Enum: ["ipv4-address"]
	NextHopType *string `json:"nextHopType,omitempty"`
}

// Validate validates this route update
func (m *RouteUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextHop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextHopType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var routeUpdateTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deliver"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeUpdateTypeActionPropEnum = append(routeUpdateTypeActionPropEnum, v)
	}
}

const (

	// RouteUpdateActionDeliver captures enum value "deliver"
	RouteUpdateActionDeliver string = "deliver"
)

// prop value enum
func (m *RouteUpdate) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routeUpdateTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RouteUpdate) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *RouteUpdate) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

var routeUpdateTypeDestinationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4-address"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeUpdateTypeDestinationTypePropEnum = append(routeUpdateTypeDestinationTypePropEnum, v)
	}
}

const (

	// RouteUpdateDestinationTypeIPV4DashAddress captures enum value "ipv4-address"
	RouteUpdateDestinationTypeIPV4DashAddress string = "ipv4-address"
)

// prop value enum
func (m *RouteUpdate) validateDestinationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routeUpdateTypeDestinationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RouteUpdate) validateDestinationType(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDestinationTypeEnum("destinationType", "body", *m.DestinationType); err != nil {
		return err
	}

	return nil
}

func (m *RouteUpdate) validateName(formats strfmt.Registry) error {

	if m.Name == nil {
		return errors.Required("name", "body", nil)
	}

	return nil
}

func (m *RouteUpdate) validateNextHop(formats strfmt.Registry) error {

	if err := validate.Required("nextHop", "body", m.NextHop); err != nil {
		return err
	}

	return nil
}

var routeUpdateTypeNextHopTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4-address"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeUpdateTypeNextHopTypePropEnum = append(routeUpdateTypeNextHopTypePropEnum, v)
	}
}

const (

	// RouteUpdateNextHopTypeIPV4DashAddress captures enum value "ipv4-address"
	RouteUpdateNextHopTypeIPV4DashAddress string = "ipv4-address"
)

// prop value enum
func (m *RouteUpdate) validateNextHopTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routeUpdateTypeNextHopTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RouteUpdate) validateNextHopType(formats strfmt.Registry) error {
	if swag.IsZero(m.NextHopType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNextHopTypeEnum("nextHopType", "body", *m.NextHopType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this route update based on context it is used
func (m *RouteUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouteUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteUpdate) UnmarshalBinary(b []byte) error {
	var res RouteUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
