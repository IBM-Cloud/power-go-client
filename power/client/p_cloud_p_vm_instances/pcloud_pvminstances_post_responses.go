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

// PcloudPvminstancesPostReader is a Reader for the PcloudPvminstancesPost structure.
type PcloudPvminstancesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPcloudPvminstancesPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudPvminstancesPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPvminstancesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPcloudPvminstancesPostGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesPostOK creates a PcloudPvminstancesPostOK with default headers values
func NewPcloudPvminstancesPostOK() *PcloudPvminstancesPostOK {
	return &PcloudPvminstancesPostOK{}
}

/* PcloudPvminstancesPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesPostOK struct {
	Payload models.PVMInstanceList
}

func (o *PcloudPvminstancesPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostOK  %+v", 200, o.Payload)
}
func (o *PcloudPvminstancesPostOK) GetPayload() models.PVMInstanceList {
	return o.Payload
}

func (o *PcloudPvminstancesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostCreated creates a PcloudPvminstancesPostCreated with default headers values
func NewPcloudPvminstancesPostCreated() *PcloudPvminstancesPostCreated {
	return &PcloudPvminstancesPostCreated{}
}

/* PcloudPvminstancesPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudPvminstancesPostCreated struct {
	Payload models.PVMInstanceList
}

func (o *PcloudPvminstancesPostCreated) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostCreated  %+v", 201, o.Payload)
}
func (o *PcloudPvminstancesPostCreated) GetPayload() models.PVMInstanceList {
	return o.Payload
}

func (o *PcloudPvminstancesPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostAccepted creates a PcloudPvminstancesPostAccepted with default headers values
func NewPcloudPvminstancesPostAccepted() *PcloudPvminstancesPostAccepted {
	return &PcloudPvminstancesPostAccepted{}
}

/* PcloudPvminstancesPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudPvminstancesPostAccepted struct {
	Payload models.PVMInstanceList
}

func (o *PcloudPvminstancesPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudPvminstancesPostAccepted) GetPayload() models.PVMInstanceList {
	return o.Payload
}

func (o *PcloudPvminstancesPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostBadRequest creates a PcloudPvminstancesPostBadRequest with default headers values
func NewPcloudPvminstancesPostBadRequest() *PcloudPvminstancesPostBadRequest {
	return &PcloudPvminstancesPostBadRequest{}
}

/* PcloudPvminstancesPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudPvminstancesPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostUnauthorized creates a PcloudPvminstancesPostUnauthorized with default headers values
func NewPcloudPvminstancesPostUnauthorized() *PcloudPvminstancesPostUnauthorized {
	return &PcloudPvminstancesPostUnauthorized{}
}

/* PcloudPvminstancesPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudPvminstancesPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostConflict creates a PcloudPvminstancesPostConflict with default headers values
func NewPcloudPvminstancesPostConflict() *PcloudPvminstancesPostConflict {
	return &PcloudPvminstancesPostConflict{}
}

/* PcloudPvminstancesPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesPostConflict struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudPvminstancesPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostUnprocessableEntity creates a PcloudPvminstancesPostUnprocessableEntity with default headers values
func NewPcloudPvminstancesPostUnprocessableEntity() *PcloudPvminstancesPostUnprocessableEntity {
	return &PcloudPvminstancesPostUnprocessableEntity{}
}

/* PcloudPvminstancesPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudPvminstancesPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostInternalServerError creates a PcloudPvminstancesPostInternalServerError with default headers values
func NewPcloudPvminstancesPostInternalServerError() *PcloudPvminstancesPostInternalServerError {
	return &PcloudPvminstancesPostInternalServerError{}
}

/* PcloudPvminstancesPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudPvminstancesPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostGatewayTimeout creates a PcloudPvminstancesPostGatewayTimeout with default headers values
func NewPcloudPvminstancesPostGatewayTimeout() *PcloudPvminstancesPostGatewayTimeout {
	return &PcloudPvminstancesPostGatewayTimeout{}
}

/* PcloudPvminstancesPostGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. Request is still processing and taking longer than expected.
*/
type PcloudPvminstancesPostGatewayTimeout struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostGatewayTimeout  %+v", 504, o.Payload)
}
func (o *PcloudPvminstancesPostGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesPostGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
