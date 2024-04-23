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

// PowerEdgeRouterAction power edge router action
//
// swagger:model PowerEdgeRouterAction
type PowerEdgeRouterAction struct {

	// Name of the action to take; can be migrate-start, migrate-validate
	// Required: true
	// Enum: [migrate-start migrate-validate]
	Action *string `json:"action"`
}

// Validate validates this power edge router action
func (m *PowerEdgeRouterAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerEdgeRouterActionTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["migrate-start","migrate-validate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerEdgeRouterActionTypeActionPropEnum = append(powerEdgeRouterActionTypeActionPropEnum, v)
	}
}

const (

	// PowerEdgeRouterActionActionMigrateDashStart captures enum value "migrate-start"
	PowerEdgeRouterActionActionMigrateDashStart string = "migrate-start"

	// PowerEdgeRouterActionActionMigrateDashValidate captures enum value "migrate-validate"
	PowerEdgeRouterActionActionMigrateDashValidate string = "migrate-validate"
)

// prop value enum
func (m *PowerEdgeRouterAction) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerEdgeRouterActionTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerEdgeRouterAction) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power edge router action based on context it is used
func (m *PowerEdgeRouterAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerEdgeRouterAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerEdgeRouterAction) UnmarshalBinary(b []byte) error {
	var res PowerEdgeRouterAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
