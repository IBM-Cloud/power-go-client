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

// Volume volume
//
// swagger:model Volume
type Volume struct {

	// Auxiliary volume name at storage host level
	AuxVolumeName string `json:"auxVolumeName,omitempty"`

	// true if volume is auxiliary otherwise false
	Auxiliary *bool `json:"auxiliary,omitempty"`

	// Indicates if the volume is the server's boot volume
	BootVolume *bool `json:"bootVolume,omitempty"`

	// Indicates if the volume is boot capable
	Bootable *bool `json:"bootable,omitempty"`

	// Consistency Group Name if volume is a part of volume group
	ConsistencyGroupName string `json:"consistencyGroupName,omitempty"`

	// Creation Date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// crn
	Crn CRN `json:"crn,omitempty"`

	// Indicates if the volume should be deleted when the server terminates
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty"`

	// Type of Disk
	DiskType string `json:"diskType,omitempty"`

	// Freeze time of remote copy relationship
	// Format: date-time
	FreezeTime *strfmt.DateTime `json:"freezeTime,omitempty"`

	// Volume Group ID
	GroupID string `json:"groupID,omitempty"`

	// Amount of iops assigned to the volume
	IoThrottleRate string `json:"ioThrottleRate,omitempty"`

	// Last Update Date
	// Required: true
	// Format: date-time
	LastUpdateDate *strfmt.DateTime `json:"lastUpdateDate"`

	// Master volume name at storage host level
	MasterVolumeName string `json:"masterVolumeName,omitempty"`

	// Mirroring state for replication enabled volume
	MirroringState string `json:"mirroringState,omitempty"`

	// Volume Name
	// Required: true
	Name *string `json:"name"`

	// true if volume does not exist on storage controller, as volume has been deleted by deleting its paired volume from the mapped replication site.
	OutOfBandDeleted bool `json:"outOfBandDeleted,omitempty"`

	// indicates whether master/aux volume is playing the primary role
	// Enum: ["master","aux"]
	PrimaryRole string `json:"primaryRole,omitempty"`

	// List of PCloud PVM Instance attached to the volume
	PvmInstanceIDs []string `json:"pvmInstanceIDs"`

	// True if volume is replication enabled otherwise false
	ReplicationEnabled *bool `json:"replicationEnabled,omitempty"`

	// List of replication site for volume replication
	ReplicationSites []string `json:"replicationSites,omitempty"`

	// Replication status of a volume
	ReplicationStatus string `json:"replicationStatus,omitempty"`

	// type of replication(metro,global)
	ReplicationType string `json:"replicationType,omitempty"`

	// Indicates if the volume is shareable between VMs
	Shareable *bool `json:"shareable,omitempty"`

	// Volume Size
	// Required: true
	Size *float64 `json:"size"`

	// Volume State
	State string `json:"state,omitempty"`

	// user tags
	UserTags Tags `json:"userTags,omitempty"`

	// Volume ID
	// Required: true
	VolumeID *string `json:"volumeID"`

	// Volume pool, name of storage pool where the volume is located
	VolumePool string `json:"volumePool,omitempty"`

	// Volume type, name of storage template used to create the volume
	VolumeType string `json:"volumeType,omitempty"`

	// Volume world wide name
	Wwn string `json:"wwn,omitempty"`
}

// Validate validates this volume
func (m *Volume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreezeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volume) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateCrn(formats strfmt.Registry) error {
	if swag.IsZero(m.Crn) { // not required
		return nil
	}

	if err := m.Crn.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("crn")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("crn")
		}
		return err
	}

	return nil
}

func (m *Volume) validateFreezeTime(formats strfmt.Registry) error {
	if swag.IsZero(m.FreezeTime) { // not required
		return nil
	}

	if err := validate.FormatOf("freezeTime", "body", "date-time", m.FreezeTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateLastUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdateDate", "body", m.LastUpdateDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdateDate", "body", "date-time", m.LastUpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var volumeTypePrimaryRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["master","aux"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeTypePrimaryRolePropEnum = append(volumeTypePrimaryRolePropEnum, v)
	}
}

const (

	// VolumePrimaryRoleMaster captures enum value "master"
	VolumePrimaryRoleMaster string = "master"

	// VolumePrimaryRoleAux captures enum value "aux"
	VolumePrimaryRoleAux string = "aux"
)

// prop value enum
func (m *Volume) validatePrimaryRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, volumeTypePrimaryRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Volume) validatePrimaryRole(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryRole) { // not required
		return nil
	}

	// value enum
	if err := m.validatePrimaryRoleEnum("primaryRole", "body", m.PrimaryRole); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateUserTags(formats strfmt.Registry) error {
	if swag.IsZero(m.UserTags) { // not required
		return nil
	}

	if err := m.UserTags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userTags")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userTags")
		}
		return err
	}

	return nil
}

func (m *Volume) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volumeID", "body", m.VolumeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this volume based on the context it is used
func (m *Volume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCrn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volume) contextValidateCrn(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Crn) { // not required
		return nil
	}

	if err := m.Crn.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("crn")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("crn")
		}
		return err
	}

	return nil
}

func (m *Volume) contextValidateUserTags(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UserTags.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userTags")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userTags")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Volume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Volume) UnmarshalBinary(b []byte) error {
	var res Volume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
