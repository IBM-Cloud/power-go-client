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

// VolumesCloneAsyncRequest volumes clone async request
//
// swagger:model VolumesCloneAsyncRequest
type VolumesCloneAsyncRequest struct {

	// Base name of the new cloned volume(s).
	// Cloned Volume names will be prefixed with 'clone-'
	//     and suffixed with '-#####' (where ##### is a 5 digit random number)
	// If multiple volumes cloned they will be further suffixed with an incremental number starting with 1.
	//   Example volume names using name="volume-abcdef"
	//     single volume clone will be named "clone-volume-abcdef-83081"
	//     multi volume clone will be named "clone-volume-abcdef-73721-1", "clone-volume-abcdef-73721-2", ...
	//
	// Required: true
	Name *string `json:"name"`

	// Target storage tier for the cloned volumes. Use to clone a set of volumes from one storage tier
	// to a different storage tier. Cloned volumes must remain in the same storage pool as
	// the source volumes.
	//
	TargetStorageTier string `json:"targetStorageTier,omitempty"`

	// List of volumes to be cloned
	// Required: true
	VolumeIDs []string `json:"volumeIDs"`
}

// Validate validates this volumes clone async request
func (m *VolumesCloneAsyncRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeIDs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumesCloneAsyncRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VolumesCloneAsyncRequest) validateVolumeIDs(formats strfmt.Registry) error {

	if err := validate.Required("volumeIDs", "body", m.VolumeIDs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this volumes clone async request based on context it is used
func (m *VolumesCloneAsyncRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumesCloneAsyncRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumesCloneAsyncRequest) UnmarshalBinary(b []byte) error {
	var res VolumesCloneAsyncRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
