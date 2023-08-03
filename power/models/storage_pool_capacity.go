// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StoragePoolCapacity Storage pool capacity
//
// swagger:model StoragePoolCapacity
type StoragePoolCapacity struct {

	// Available/Unused pool capacity (GB)
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`

	// Maximum allocation storage size (GB)
	// Required: true
	MaxAllocationSize *int64 `json:"maxAllocationSize"`

	// Pool name
	PoolName string `json:"poolName,omitempty"`

	// true if storage-pool is replication enabled and can be used to manage replication enabled volumes
	ReplicationEnabled bool `json:"replicationEnabled,omitempty"`

	// Storage host/controller for this storage pool
	StorageHost string `json:"storageHost,omitempty"`

	// Storage type of the storage pool
	StorageType string `json:"storageType,omitempty"`

	// Total pool capacity (GB)
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
}

// Validate validates this storage pool capacity
func (m *StoragePoolCapacity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxAllocationSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePoolCapacity) validateMaxAllocationSize(formats strfmt.Registry) error {

	if err := validate.Required("maxAllocationSize", "body", m.MaxAllocationSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this storage pool capacity based on context it is used
func (m *StoragePoolCapacity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StoragePoolCapacity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePoolCapacity) UnmarshalBinary(b []byte) error {
	var res StoragePoolCapacity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
