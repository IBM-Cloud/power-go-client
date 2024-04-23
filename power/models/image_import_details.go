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

// ImageImportDetails image import details
//
// swagger:model ImageImportDetails
type ImageImportDetails struct {

	// Origin of the license of the product
	// Required: true
	// Enum: [byol]
	LicenseType *string `json:"licenseType"`

	// Product within the image
	// Required: true
	// Enum: [Hana Netweaver]
	Product *string `json:"product"`

	// Vendor supporting the product
	// Required: true
	// Enum: [SAP]
	Vendor *string `json:"vendor"`
}

// Validate validates this image import details
func (m *ImageImportDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicenseType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var imageImportDetailsTypeLicenseTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["byol"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageImportDetailsTypeLicenseTypePropEnum = append(imageImportDetailsTypeLicenseTypePropEnum, v)
	}
}

const (

	// ImageImportDetailsLicenseTypeByol captures enum value "byol"
	ImageImportDetailsLicenseTypeByol string = "byol"
)

// prop value enum
func (m *ImageImportDetails) validateLicenseTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, imageImportDetailsTypeLicenseTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ImageImportDetails) validateLicenseType(formats strfmt.Registry) error {

	if err := validate.Required("licenseType", "body", m.LicenseType); err != nil {
		return err
	}

	// value enum
	if err := m.validateLicenseTypeEnum("licenseType", "body", *m.LicenseType); err != nil {
		return err
	}

	return nil
}

var imageImportDetailsTypeProductPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Hana","Netweaver"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageImportDetailsTypeProductPropEnum = append(imageImportDetailsTypeProductPropEnum, v)
	}
}

const (

	// ImageImportDetailsProductHana captures enum value "Hana"
	ImageImportDetailsProductHana string = "Hana"

	// ImageImportDetailsProductNetweaver captures enum value "Netweaver"
	ImageImportDetailsProductNetweaver string = "Netweaver"
)

// prop value enum
func (m *ImageImportDetails) validateProductEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, imageImportDetailsTypeProductPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ImageImportDetails) validateProduct(formats strfmt.Registry) error {

	if err := validate.Required("product", "body", m.Product); err != nil {
		return err
	}

	// value enum
	if err := m.validateProductEnum("product", "body", *m.Product); err != nil {
		return err
	}

	return nil
}

var imageImportDetailsTypeVendorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SAP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageImportDetailsTypeVendorPropEnum = append(imageImportDetailsTypeVendorPropEnum, v)
	}
}

const (

	// ImageImportDetailsVendorSAP captures enum value "SAP"
	ImageImportDetailsVendorSAP string = "SAP"
)

// prop value enum
func (m *ImageImportDetails) validateVendorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, imageImportDetailsTypeVendorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ImageImportDetails) validateVendor(formats strfmt.Registry) error {

	if err := validate.Required("vendor", "body", m.Vendor); err != nil {
		return err
	}

	// value enum
	if err := m.validateVendorEnum("vendor", "body", *m.Vendor); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this image import details based on context it is used
func (m *ImageImportDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageImportDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageImportDetails) UnmarshalBinary(b []byte) error {
	var res ImageImportDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
