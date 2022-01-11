// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesVolumesGetReader is a Reader for the PcloudPvminstancesVolumesGet structure.
type PcloudPvminstancesVolumesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVolumesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesVolumesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesVolumesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesVolumesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesVolumesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesVolumesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesVolumesGetOK creates a PcloudPvminstancesVolumesGetOK with default headers values
func NewPcloudPvminstancesVolumesGetOK() *PcloudPvminstancesVolumesGetOK {
	return &PcloudPvminstancesVolumesGetOK{}
}

/* PcloudPvminstancesVolumesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesVolumesGetOK struct {
	Payload *models.Volume
}

func (o *PcloudPvminstancesVolumesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetOK  %+v", 200, o.Payload)
}
func (o *PcloudPvminstancesVolumesGetOK) GetPayload() *models.Volume {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetBadRequest creates a PcloudPvminstancesVolumesGetBadRequest with default headers values
func NewPcloudPvminstancesVolumesGetBadRequest() *PcloudPvminstancesVolumesGetBadRequest {
	return &PcloudPvminstancesVolumesGetBadRequest{}
}

/* PcloudPvminstancesVolumesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesVolumesGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudPvminstancesVolumesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetUnauthorized creates a PcloudPvminstancesVolumesGetUnauthorized with default headers values
func NewPcloudPvminstancesVolumesGetUnauthorized() *PcloudPvminstancesVolumesGetUnauthorized {
	return &PcloudPvminstancesVolumesGetUnauthorized{}
}

/* PcloudPvminstancesVolumesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesVolumesGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudPvminstancesVolumesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetNotFound creates a PcloudPvminstancesVolumesGetNotFound with default headers values
func NewPcloudPvminstancesVolumesGetNotFound() *PcloudPvminstancesVolumesGetNotFound {
	return &PcloudPvminstancesVolumesGetNotFound{}
}

/* PcloudPvminstancesVolumesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesVolumesGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudPvminstancesVolumesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetInternalServerError creates a PcloudPvminstancesVolumesGetInternalServerError with default headers values
func NewPcloudPvminstancesVolumesGetInternalServerError() *PcloudPvminstancesVolumesGetInternalServerError {
	return &PcloudPvminstancesVolumesGetInternalServerError{}
}

/* PcloudPvminstancesVolumesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesVolumesGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudPvminstancesVolumesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
