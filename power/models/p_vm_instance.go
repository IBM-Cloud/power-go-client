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

// PVMInstance p VM instance
//
// swagger:model PVMInstance
type PVMInstance struct {

	// (deprecated - replaced by networks) The list of addresses and their network information
	Addresses []*PVMInstanceNetwork `json:"addresses"`

	// Console language and code
	ConsoleLanguage *ConsoleLanguage `json:"consoleLanguage,omitempty"`

	// Date/Time of PVM creation
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// crn
	Crn CRN `json:"crn,omitempty"`

	// The custom deployment type
	DeploymentType string `json:"deploymentType,omitempty"`

	// Size of allocated disk (in GB)
	// Required: true
	DiskSize *float64 `json:"diskSize"`

	// fault
	Fault *PVMInstanceFault `json:"fault,omitempty"`

	// health
	Health *PVMInstanceHealth `json:"health,omitempty"`

	// The PVM Instance Host ID (Internal Use Only)
	HostID int64 `json:"hostID,omitempty"`

	// The ImageID used by the server
	// Required: true
	ImageID *string `json:"imageID"`

	// The VTL license repository capacity TB value
	LicenseRepositoryCapacity int64 `json:"licenseRepositoryCapacity,omitempty"`

	// Maximum amount of memory that can be allocated (in GB, for resize)
	Maxmem float64 `json:"maxmem,omitempty"`

	// Maximum number of processors that can be allocated (for resize)
	Maxproc float64 `json:"maxproc,omitempty"`

	// Amount of memory allocated (in GB)
	// Required: true
	Memory *float64 `json:"memory"`

	// whether the instance can be migrated
	Migratable *bool `json:"migratable,omitempty"`

	// Minimum amount of memory that can be allocated (in GB, for resize)
	Minmem float64 `json:"minmem,omitempty"`

	// Minimum number of processors that can be allocated (for resize)
	Minproc float64 `json:"minproc,omitempty"`

	// (deprecated - replaced by networks) List of Network IDs
	// Required: true
	NetworkIDs []string `json:"networkIDs"`

	// The pvm instance networks information
	Networks []*PVMInstanceNetwork `json:"networks"`

	// OS system information (usually version and build)
	OperatingSystem string `json:"operatingSystem,omitempty"`

	// Type of the OS [aix, ibmi, rhel, sles, vtl, rhcos]
	// Required: true
	OsType *string `json:"osType"`

	// VM pinning policy to use [none, soft, hard]
	PinPolicy string `json:"pinPolicy,omitempty"`

	// The placement group of the server
	PlacementGroup *string `json:"placementGroup,omitempty"`

	// Processor type (dedicated, shared, capped)
	// Required: true
	// Enum: ["dedicated","shared","capped",""]
	ProcType *string `json:"procType"`

	// Number of processors allocated
	// Required: true
	Processors *float64 `json:"processors"`

	// The progress of an operation
	Progress float64 `json:"progress,omitempty"`

	// PCloud PVM Instance ID
	// Required: true
	PvmInstanceID *string `json:"pvmInstanceID"`

	// If this is an SAP pvm-instance the profile reference will link to the SAP profile
	SapProfile *SAPProfileReference `json:"sapProfile,omitempty"`

	// Name of the server
	// Required: true
	ServerName *string `json:"serverName"`

	// The shared processor pool of the server
	SharedProcessorPool string `json:"sharedProcessorPool,omitempty"`

	// The shared processor pool id
	SharedProcessorPoolID string `json:"sharedProcessorPoolID,omitempty"`

	// The pvm instance Software Licenses
	SoftwareLicenses *SoftwareLicenses `json:"softwareLicenses,omitempty"`

	// The pvm instance SRC lists
	Srcs [][]*SRC `json:"srcs"`

	// The status of the instance
	// Required: true
	Status *string `json:"status"`

	// The storage connection type
	StorageConnection string `json:"storageConnection,omitempty"`

	// Storage Pool where server is deployed
	StoragePool string `json:"storagePool,omitempty"`

	// Indicates if all volumes attached to the server must reside in the same storage pool; Defaults to true when initially deploying a PVMInstance
	StoragePoolAffinity *bool `json:"storagePoolAffinity,omitempty"`

	// Storage type where server is deployed
	// Required: true
	StorageType *string `json:"storageType"`

	// System type used to host the instance
	SysType string `json:"sysType,omitempty"`

	// Represents the task state of a virtual machine (VM).
	TaskState string `json:"taskState,omitempty"`

	// Date/Time of PVM last update
	// Format: date-time
	UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`

	// user tags
	UserTags Tags `json:"userTags,omitempty"`

	// The pvm instance virtual CPU information
	VirtualCores *VirtualCores `json:"virtualCores,omitempty"`

	// Information about Virtual Serial Number assigned to the PVM Instance
	VirtualSerialNumber *GetServerVirtualSerialNumber `json:"virtualSerialNumber,omitempty"`

	// List of volume IDs
	// Required: true
	VolumeIDs []string `json:"volumeIDs"`
}

// Validate validates this p VM instance
func (m *PVMInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsoleLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePvmInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSapProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualSerialNumber(formats); err != nil {
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

func (m *PVMInstance) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PVMInstance) validateConsoleLanguage(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsoleLanguage) { // not required
		return nil
	}

	if m.ConsoleLanguage != nil {
		if err := m.ConsoleLanguage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consoleLanguage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consoleLanguage")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateCrn(formats strfmt.Registry) error {
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

func (m *PVMInstance) validateDiskSize(formats strfmt.Registry) error {

	if err := validate.Required("diskSize", "body", m.DiskSize); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateFault(formats strfmt.Registry) error {
	if swag.IsZero(m.Fault) { // not required
		return nil
	}

	if m.Fault != nil {
		if err := m.Fault.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fault")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fault")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageID", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateNetworkIDs(formats strfmt.Registry) error {

	if err := validate.Required("networkIDs", "body", m.NetworkIDs); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateNetworks(formats strfmt.Registry) error {
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

func (m *PVMInstance) validateOsType(formats strfmt.Registry) error {

	if err := validate.Required("osType", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

var pVmInstanceTypeProcTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dedicated","shared","capped",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceTypeProcTypePropEnum = append(pVmInstanceTypeProcTypePropEnum, v)
	}
}

const (

	// PVMInstanceProcTypeDedicated captures enum value "dedicated"
	PVMInstanceProcTypeDedicated string = "dedicated"

	// PVMInstanceProcTypeShared captures enum value "shared"
	PVMInstanceProcTypeShared string = "shared"

	// PVMInstanceProcTypeCapped captures enum value "capped"
	PVMInstanceProcTypeCapped string = "capped"

	// PVMInstanceProcTypeEmpty captures enum value ""
	PVMInstanceProcTypeEmpty string = ""
)

// prop value enum
func (m *PVMInstance) validateProcTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceTypeProcTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstance) validateProcType(formats strfmt.Registry) error {

	if err := validate.Required("procType", "body", m.ProcType); err != nil {
		return err
	}

	// value enum
	if err := m.validateProcTypeEnum("procType", "body", *m.ProcType); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateProcessors(formats strfmt.Registry) error {

	if err := validate.Required("processors", "body", m.Processors); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validatePvmInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("pvmInstanceID", "body", m.PvmInstanceID); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateSapProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.SapProfile) { // not required
		return nil
	}

	if m.SapProfile != nil {
		if err := m.SapProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sapProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sapProfile")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) validateServerName(formats strfmt.Registry) error {

	if err := validate.Required("serverName", "body", m.ServerName); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateSoftwareLicenses(formats strfmt.Registry) error {
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

func (m *PVMInstance) validateSrcs(formats strfmt.Registry) error {
	if swag.IsZero(m.Srcs) { // not required
		return nil
	}

	for i := 0; i < len(m.Srcs); i++ {

		for ii := 0; ii < len(m.Srcs[i]); ii++ {
			if swag.IsZero(m.Srcs[i][ii]) { // not required
				continue
			}

			if m.Srcs[i][ii] != nil {
				if err := m.Srcs[i][ii].Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("srcs" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("srcs" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					}
					return err
				}
			}

		}

	}

	return nil
}

func (m *PVMInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateStorageType(formats strfmt.Registry) error {

	if err := validate.Required("storageType", "body", m.StorageType); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateUpdatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedDate", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstance) validateUserTags(formats strfmt.Registry) error {
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

func (m *PVMInstance) validateVirtualCores(formats strfmt.Registry) error {
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

func (m *PVMInstance) validateVirtualSerialNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualSerialNumber) { // not required
		return nil
	}

	if m.VirtualSerialNumber != nil {
		if err := m.VirtualSerialNumber.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualSerialNumber")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualSerialNumber")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) validateVolumeIDs(formats strfmt.Registry) error {

	if err := validate.Required("volumeIDs", "body", m.VolumeIDs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p VM instance based on the context it is used
func (m *PVMInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsoleLanguage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCrn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSapProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSrcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualCores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualSerialNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstance) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {

			if swag.IsZero(m.Addresses[i]) { // not required
				return nil
			}

			if err := m.Addresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PVMInstance) contextValidateConsoleLanguage(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsoleLanguage != nil {

		if swag.IsZero(m.ConsoleLanguage) { // not required
			return nil
		}

		if err := m.ConsoleLanguage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consoleLanguage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consoleLanguage")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) contextValidateCrn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PVMInstance) contextValidateFault(ctx context.Context, formats strfmt.Registry) error {

	if m.Fault != nil {

		if swag.IsZero(m.Fault) { // not required
			return nil
		}

		if err := m.Fault.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fault")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fault")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {

		if swag.IsZero(m.Health) { // not required
			return nil
		}

		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PVMInstance) contextValidateSapProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.SapProfile != nil {

		if swag.IsZero(m.SapProfile) { // not required
			return nil
		}

		if err := m.SapProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sapProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sapProfile")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstance) contextValidateSoftwareLicenses(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PVMInstance) contextValidateSrcs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Srcs); i++ {

		for ii := 0; ii < len(m.Srcs[i]); ii++ {

			if m.Srcs[i][ii] != nil {

				if swag.IsZero(m.Srcs[i][ii]) { // not required
					return nil
				}

				if err := m.Srcs[i][ii].ContextValidate(ctx, formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("srcs" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("srcs" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					}
					return err
				}
			}

		}

	}

	return nil
}

func (m *PVMInstance) contextValidateUserTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PVMInstance) contextValidateVirtualCores(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PVMInstance) contextValidateVirtualSerialNumber(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualSerialNumber != nil {

		if swag.IsZero(m.VirtualSerialNumber) { // not required
			return nil
		}

		if err := m.VirtualSerialNumber.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualSerialNumber")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualSerialNumber")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstance) UnmarshalBinary(b []byte) error {
	var res PVMInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
