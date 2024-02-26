// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstanceCreate p VM instance create
//
// swagger:model PVMInstanceCreate
type PVMInstanceCreate struct {

	// The name of the host or hostgroup where to deploy the VM
	DeployTarget string `json:"deployTarget,omitempty"`

	// The custom deployment type
	DeploymentType string `json:"deploymentType,omitempty"`

	// Image ID of the image to use for the server
	// Required: true
	ImageID *string `json:"imageID"`

	// The name of the SSH key pair provided to the server for authenticating users (looked up in the tenant's list of keys)
	KeyPairName string `json:"keyPairName,omitempty"`

	// The VTL license repository capacity TB value
	LicenseRepositoryCapacity int64 `json:"licenseRepositoryCapacity,omitempty"`

	// Amount of memory allocated (in GB)
	// Required: true
	Memory *float64 `json:"memory"`

	// (deprecated - replaced by pinPolicy) Indicates if the server is allowed to migrate between hosts
	Migratable *bool `json:"migratable,omitempty"`

	// (deprecated - replaced by networks) List of Network IDs
	NetworkIDs []string `json:"networkIDs"`

	// The pvm instance networks information
	Networks []*PVMInstanceAddNetwork `json:"networks"`

	// pin policy
	PinPolicy PinPolicy `json:"pinPolicy,omitempty"`

	// The placement group for the server
	PlacementGroup string `json:"placementGroup,omitempty"`

	// Processor type (dedicated, shared, capped)
	// Required: true
	// Enum: [dedicated shared capped]
	ProcType *string `json:"procType"`

	// Number of processors allocated
	// Required: true
	Processors *float64 `json:"processors"`

	// Affinity policy for replicants being created; affinity for the same host, anti-affinity for different hosts, none for no preference
	// Enum: [affinity anti-affinity none]
	ReplicantAffinityPolicy *string `json:"replicantAffinityPolicy,omitempty"`

	// How to name the created vms
	// Enum: [prefix suffix]
	ReplicantNamingScheme *string `json:"replicantNamingScheme,omitempty"`

	// Number of duplicate instances to create in this request
	Replicants float64 `json:"replicants,omitempty"`

	// Name of the server to create
	// Required: true
	ServerName *string `json:"serverName"`

	// The shared processor pool for server deployment
	SharedProcessorPool string `json:"sharedProcessorPool,omitempty"`

	// The pvm instance Software Licenses
	SoftwareLicenses *SoftwareLicenses `json:"softwareLicenses,omitempty"`

	// The storage affinity data; ignored if storagePool is provided; Only valid when you deploy one of the IBM supplied stock images. Storage pool for a custom image (an imported image or an image that is created from a PVMInstance capture) defaults to the storage pool the image was created in
	StorageAffinity *StorageAffinity `json:"storageAffinity,omitempty"`

	// The storage connection type
	// Enum: [vSCSI]
	StorageConnection string `json:"storageConnection,omitempty"`

	// The storage connection type
	// Enum: [vSCSI maxVolumeSupport]
	StorageConnectionV2 string `json:"storageConnectionV2,omitempty"`

	// Storage Pool for server deployment; if provided then storageAffinity will be ignored; Only valid when you deploy one of the IBM supplied stock images. Storage pool for a custom image (an imported image or an image that is created from a PVMInstance capture) defaults to the storage pool the image was created in
	StoragePool string `json:"storagePool,omitempty"`

	// Storage type for server deployment; if storageType is not provided the storage type will default to 'tier3'.
	StorageType string `json:"storageType,omitempty"`

	// System type used to host the instance
	SysType string `json:"sysType,omitempty"`

	// Cloud init user defined data
	UserData string `json:"userData,omitempty"`

	// The pvm instance virtual CPU information
	VirtualCores *VirtualCores `json:"virtualCores,omitempty"`

	// List of volume IDs
	VolumeIDs []string `json:"volumeIDs"`
}

// Validate validates this p VM instance create
func (m *PVMInstanceCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePinPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicantAffinityPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicantNamingScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageConnectionV2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualCores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceCreate) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageID", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceCreate) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceCreate) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PVMInstanceCreate) validatePinPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.PinPolicy) { // not required
		return nil
	}

	if err := m.PinPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pinPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("pinPolicy")
		}
		return err
	}

	return nil
}

var pVmInstanceCreateTypeProcTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dedicated","shared","capped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceCreateTypeProcTypePropEnum = append(pVmInstanceCreateTypeProcTypePropEnum, v)
	}
}

const (

	// PVMInstanceCreateProcTypeDedicated captures enum value "dedicated"
	PVMInstanceCreateProcTypeDedicated string = "dedicated"

	// PVMInstanceCreateProcTypeShared captures enum value "shared"
	PVMInstanceCreateProcTypeShared string = "shared"

	// PVMInstanceCreateProcTypeCapped captures enum value "capped"
	PVMInstanceCreateProcTypeCapped string = "capped"
)

// prop value enum
func (m *PVMInstanceCreate) validateProcTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceCreateTypeProcTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceCreate) validateProcType(formats strfmt.Registry) error {

	if err := validate.Required("procType", "body", m.ProcType); err != nil {
		return err
	}

	// value enum
	if err := m.validateProcTypeEnum("procType", "body", *m.ProcType); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceCreate) validateProcessors(formats strfmt.Registry) error {

	if err := validate.Required("processors", "body", m.Processors); err != nil {
		return err
	}

	return nil
}

var pVmInstanceCreateTypeReplicantAffinityPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["affinity","anti-affinity","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceCreateTypeReplicantAffinityPolicyPropEnum = append(pVmInstanceCreateTypeReplicantAffinityPolicyPropEnum, v)
	}
}

const (

	// PVMInstanceCreateReplicantAffinityPolicyAffinity captures enum value "affinity"
	PVMInstanceCreateReplicantAffinityPolicyAffinity string = "affinity"

	// PVMInstanceCreateReplicantAffinityPolicyAntiDashAffinity captures enum value "anti-affinity"
	PVMInstanceCreateReplicantAffinityPolicyAntiDashAffinity string = "anti-affinity"

	// PVMInstanceCreateReplicantAffinityPolicyNone captures enum value "none"
	PVMInstanceCreateReplicantAffinityPolicyNone string = "none"
)

// prop value enum
func (m *PVMInstanceCreate) validateReplicantAffinityPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceCreateTypeReplicantAffinityPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceCreate) validateReplicantAffinityPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicantAffinityPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateReplicantAffinityPolicyEnum("replicantAffinityPolicy", "body", *m.ReplicantAffinityPolicy); err != nil {
		return err
	}

	return nil
}

var pVmInstanceCreateTypeReplicantNamingSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["prefix","suffix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceCreateTypeReplicantNamingSchemePropEnum = append(pVmInstanceCreateTypeReplicantNamingSchemePropEnum, v)
	}
}

const (

	// PVMInstanceCreateReplicantNamingSchemePrefix captures enum value "prefix"
	PVMInstanceCreateReplicantNamingSchemePrefix string = "prefix"

	// PVMInstanceCreateReplicantNamingSchemeSuffix captures enum value "suffix"
	PVMInstanceCreateReplicantNamingSchemeSuffix string = "suffix"
)

// prop value enum
func (m *PVMInstanceCreate) validateReplicantNamingSchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceCreateTypeReplicantNamingSchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceCreate) validateReplicantNamingScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicantNamingScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateReplicantNamingSchemeEnum("replicantNamingScheme", "body", *m.ReplicantNamingScheme); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceCreate) validateServerName(formats strfmt.Registry) error {

	if err := validate.Required("serverName", "body", m.ServerName); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceCreate) validateSoftwareLicenses(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareLicenses) { // not required
		return nil
	}

	if m.SoftwareLicenses != nil {
		if err := m.SoftwareLicenses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceCreate) validateStorageAffinity(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageAffinity) { // not required
		return nil
	}

	if m.StorageAffinity != nil {
		if err := m.StorageAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageAffinity")
			}
			return err
		}
	}

	return nil
}

var pVmInstanceCreateTypeStorageConnectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vSCSI"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceCreateTypeStorageConnectionPropEnum = append(pVmInstanceCreateTypeStorageConnectionPropEnum, v)
	}
}

const (

	// PVMInstanceCreateStorageConnectionVSCSI captures enum value "vSCSI"
	PVMInstanceCreateStorageConnectionVSCSI string = "vSCSI"
)

// prop value enum
func (m *PVMInstanceCreate) validateStorageConnectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceCreateTypeStorageConnectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceCreate) validateStorageConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageConnection) { // not required
		return nil
	}

	// value enum
	if err := m.validateStorageConnectionEnum("storageConnection", "body", m.StorageConnection); err != nil {
		return err
	}

	return nil
}

var pVmInstanceCreateTypeStorageConnectionV2PropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vSCSI","maxVolumeSupport"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceCreateTypeStorageConnectionV2PropEnum = append(pVmInstanceCreateTypeStorageConnectionV2PropEnum, v)
	}
}

const (

	// PVMInstanceCreateStorageConnectionV2VSCSI captures enum value "vSCSI"
	PVMInstanceCreateStorageConnectionV2VSCSI string = "vSCSI"

	// PVMInstanceCreateStorageConnectionV2MaxVolumeSupport captures enum value "maxVolumeSupport"
	PVMInstanceCreateStorageConnectionV2MaxVolumeSupport string = "maxVolumeSupport"
)

// prop value enum
func (m *PVMInstanceCreate) validateStorageConnectionV2Enum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceCreateTypeStorageConnectionV2PropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceCreate) validateStorageConnectionV2(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageConnectionV2) { // not required
		return nil
	}

	// value enum
	if err := m.validateStorageConnectionV2Enum("storageConnectionV2", "body", m.StorageConnectionV2); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceCreate) validateVirtualCores(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualCores) { // not required
		return nil
	}

	if m.VirtualCores != nil {
		if err := m.VirtualCores.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualCores")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualCores")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this p VM instance create based on the context it is used
func (m *PVMInstanceCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePinPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualCores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceCreate) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {

			if swag.IsZero(m.Networks[i]) { // not required
				return nil
			}

			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PVMInstanceCreate) contextValidatePinPolicy(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.PinPolicy) { // not required
		return nil
	}

	if err := m.PinPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pinPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("pinPolicy")
		}
		return err
	}

	return nil
}

func (m *PVMInstanceCreate) contextValidateSoftwareLicenses(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareLicenses != nil {

		if swag.IsZero(m.SoftwareLicenses) { // not required
			return nil
		}

		if err := m.SoftwareLicenses.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceCreate) contextValidateStorageAffinity(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageAffinity != nil {

		if swag.IsZero(m.StorageAffinity) { // not required
			return nil
		}

		if err := m.StorageAffinity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageAffinity")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceCreate) contextValidateVirtualCores(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualCores != nil {

		if swag.IsZero(m.VirtualCores) { // not required
			return nil
		}

		if err := m.VirtualCores.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualCores")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualCores")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceCreate) UnmarshalBinary(b []byte) error {
	var res PVMInstanceCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
