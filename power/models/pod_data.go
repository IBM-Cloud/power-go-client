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

// PodData Description of a PPCaaS Pod
//
// swagger:model PodData
type PodData struct {

	// Number of available cores in the Pod
	// Required: true
	Cores *float64 `json:"cores"`

	// Amount of available memory in the Pod (GB)
	// Required: true
	Memory *int64 `json:"memory"`

	// ID of the Satellite Location
	// Required: true
	SatLocationID *string `json:"satLocationID"`

	// Amount of available storage in the Pod (GB)
	// Required: true
	Storage *int64 `json:"storage"`

	// Total number of usable cores in the Pod
	// Required: true
	TotalCores *float64 `json:"totalCores"`

	// Total amount of usable memory in the Pod (GB)
	// Required: true
	TotalMemory *int64 `json:"totalMemory"`

	// Total number of physical cores in the Pod
	// Required: true
	TotalPhysCores *float64 `json:"totalPhysCores"`

	// Total amount of physical memory in the Pod (GB)
	// Required: true
	TotalPhysMemory *int64 `json:"totalPhysMemory"`

	// Total amount of physical storage in the Pod (GB)
	// Required: true
	TotalPhysStorage *float64 `json:"totalPhysStorage"`

	// Total amount of usable storage in the Pod (GB)
	// Required: true
	TotalStorage *int64 `json:"totalStorage"`
}

// Validate validates this pod data
func (m *PodData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSatLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalPhysCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalPhysMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalPhysStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodData) validateCores(formats strfmt.Registry) error {

	if err := validate.Required("cores", "body", m.Cores); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateSatLocationID(formats strfmt.Registry) error {

	if err := validate.Required("satLocationID", "body", m.SatLocationID); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateTotalCores(formats strfmt.Registry) error {

	if err := validate.Required("totalCores", "body", m.TotalCores); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateTotalMemory(formats strfmt.Registry) error {

	if err := validate.Required("totalMemory", "body", m.TotalMemory); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateTotalPhysCores(formats strfmt.Registry) error {

	if err := validate.Required("totalPhysCores", "body", m.TotalPhysCores); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateTotalPhysMemory(formats strfmt.Registry) error {

	if err := validate.Required("totalPhysMemory", "body", m.TotalPhysMemory); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateTotalPhysStorage(formats strfmt.Registry) error {

	if err := validate.Required("totalPhysStorage", "body", m.TotalPhysStorage); err != nil {
		return err
	}

	return nil
}

func (m *PodData) validateTotalStorage(formats strfmt.Registry) error {

	if err := validate.Required("totalStorage", "body", m.TotalStorage); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pod data based on context it is used
func (m *PodData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodData) UnmarshalBinary(b []byte) error {
	var res PodData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
