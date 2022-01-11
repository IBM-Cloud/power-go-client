// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesCapturePostReader is a Reader for the PcloudPvminstancesCapturePost structure.
type PcloudPvminstancesCapturePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesCapturePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesCapturePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudPvminstancesCapturePostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesCapturePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesCapturePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPvminstancesCapturePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesCapturePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesCapturePostOK creates a PcloudPvminstancesCapturePostOK with default headers values
func NewPcloudPvminstancesCapturePostOK() *PcloudPvminstancesCapturePostOK {
	return &PcloudPvminstancesCapturePostOK{}
}

/* PcloudPvminstancesCapturePostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesCapturePostOK struct {
	Payload models.Object
}

func (o *PcloudPvminstancesCapturePostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudPvminstancesCapturePostOK  %+v", 200, o.Payload)
}
func (o *PcloudPvminstancesCapturePostOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesCapturePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesCapturePostAccepted creates a PcloudPvminstancesCapturePostAccepted with default headers values
func NewPcloudPvminstancesCapturePostAccepted() *PcloudPvminstancesCapturePostAccepted {
	return &PcloudPvminstancesCapturePostAccepted{}
}

/* PcloudPvminstancesCapturePostAccepted describes a response with status code 202, with default header values.

Accepted, upload to cloud storage in progress
*/
type PcloudPvminstancesCapturePostAccepted struct {
	Payload models.Object
}

func (o *PcloudPvminstancesCapturePostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudPvminstancesCapturePostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudPvminstancesCapturePostAccepted) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesCapturePostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesCapturePostBadRequest creates a PcloudPvminstancesCapturePostBadRequest with default headers values
func NewPcloudPvminstancesCapturePostBadRequest() *PcloudPvminstancesCapturePostBadRequest {
	return &PcloudPvminstancesCapturePostBadRequest{}
}

/* PcloudPvminstancesCapturePostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesCapturePostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesCapturePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudPvminstancesCapturePostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudPvminstancesCapturePostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesCapturePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesCapturePostUnauthorized creates a PcloudPvminstancesCapturePostUnauthorized with default headers values
func NewPcloudPvminstancesCapturePostUnauthorized() *PcloudPvminstancesCapturePostUnauthorized {
	return &PcloudPvminstancesCapturePostUnauthorized{}
}

/* PcloudPvminstancesCapturePostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesCapturePostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesCapturePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudPvminstancesCapturePostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudPvminstancesCapturePostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesCapturePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesCapturePostUnprocessableEntity creates a PcloudPvminstancesCapturePostUnprocessableEntity with default headers values
func NewPcloudPvminstancesCapturePostUnprocessableEntity() *PcloudPvminstancesCapturePostUnprocessableEntity {
	return &PcloudPvminstancesCapturePostUnprocessableEntity{}
}

/* PcloudPvminstancesCapturePostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesCapturePostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesCapturePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudPvminstancesCapturePostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudPvminstancesCapturePostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesCapturePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesCapturePostInternalServerError creates a PcloudPvminstancesCapturePostInternalServerError with default headers values
func NewPcloudPvminstancesCapturePostInternalServerError() *PcloudPvminstancesCapturePostInternalServerError {
	return &PcloudPvminstancesCapturePostInternalServerError{}
}

/* PcloudPvminstancesCapturePostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesCapturePostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesCapturePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudPvminstancesCapturePostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudPvminstancesCapturePostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesCapturePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
