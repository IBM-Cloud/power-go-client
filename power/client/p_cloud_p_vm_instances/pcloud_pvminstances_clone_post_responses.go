// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesClonePostReader is a Reader for the PcloudPvminstancesClonePost structure.
type PcloudPvminstancesClonePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesClonePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudPvminstancesClonePostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesClonePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesClonePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesClonePostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPvminstancesClonePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesClonePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesClonePostAccepted creates a PcloudPvminstancesClonePostAccepted with default headers values
func NewPcloudPvminstancesClonePostAccepted() *PcloudPvminstancesClonePostAccepted {
	return &PcloudPvminstancesClonePostAccepted{}
}

/*
PcloudPvminstancesClonePostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudPvminstancesClonePostAccepted struct {
	Payload *models.PVMInstance
}

// IsSuccess returns true when this pcloud pvminstances clone post accepted response has a 2xx status code
func (o *PcloudPvminstancesClonePostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances clone post accepted response has a 3xx status code
func (o *PcloudPvminstancesClonePostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances clone post accepted response has a 4xx status code
func (o *PcloudPvminstancesClonePostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances clone post accepted response has a 5xx status code
func (o *PcloudPvminstancesClonePostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances clone post accepted response a status code equal to that given
func (o *PcloudPvminstancesClonePostAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PcloudPvminstancesClonePostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudPvminstancesClonePostAccepted) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudPvminstancesClonePostAccepted) GetPayload() *models.PVMInstance {
	return o.Payload
}

func (o *PcloudPvminstancesClonePostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostBadRequest creates a PcloudPvminstancesClonePostBadRequest with default headers values
func NewPcloudPvminstancesClonePostBadRequest() *PcloudPvminstancesClonePostBadRequest {
	return &PcloudPvminstancesClonePostBadRequest{}
}

/*
PcloudPvminstancesClonePostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesClonePostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances clone post bad request response has a 2xx status code
func (o *PcloudPvminstancesClonePostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances clone post bad request response has a 3xx status code
func (o *PcloudPvminstancesClonePostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances clone post bad request response has a 4xx status code
func (o *PcloudPvminstancesClonePostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances clone post bad request response has a 5xx status code
func (o *PcloudPvminstancesClonePostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances clone post bad request response a status code equal to that given
func (o *PcloudPvminstancesClonePostBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPvminstancesClonePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesClonePostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesClonePostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesClonePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostUnauthorized creates a PcloudPvminstancesClonePostUnauthorized with default headers values
func NewPcloudPvminstancesClonePostUnauthorized() *PcloudPvminstancesClonePostUnauthorized {
	return &PcloudPvminstancesClonePostUnauthorized{}
}

/*
PcloudPvminstancesClonePostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesClonePostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances clone post unauthorized response has a 2xx status code
func (o *PcloudPvminstancesClonePostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances clone post unauthorized response has a 3xx status code
func (o *PcloudPvminstancesClonePostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances clone post unauthorized response has a 4xx status code
func (o *PcloudPvminstancesClonePostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances clone post unauthorized response has a 5xx status code
func (o *PcloudPvminstancesClonePostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances clone post unauthorized response a status code equal to that given
func (o *PcloudPvminstancesClonePostUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudPvminstancesClonePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesClonePostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesClonePostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesClonePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostConflict creates a PcloudPvminstancesClonePostConflict with default headers values
func NewPcloudPvminstancesClonePostConflict() *PcloudPvminstancesClonePostConflict {
	return &PcloudPvminstancesClonePostConflict{}
}

/*
PcloudPvminstancesClonePostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesClonePostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances clone post conflict response has a 2xx status code
func (o *PcloudPvminstancesClonePostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances clone post conflict response has a 3xx status code
func (o *PcloudPvminstancesClonePostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances clone post conflict response has a 4xx status code
func (o *PcloudPvminstancesClonePostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances clone post conflict response has a 5xx status code
func (o *PcloudPvminstancesClonePostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances clone post conflict response a status code equal to that given
func (o *PcloudPvminstancesClonePostConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PcloudPvminstancesClonePostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPvminstancesClonePostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPvminstancesClonePostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesClonePostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostUnprocessableEntity creates a PcloudPvminstancesClonePostUnprocessableEntity with default headers values
func NewPcloudPvminstancesClonePostUnprocessableEntity() *PcloudPvminstancesClonePostUnprocessableEntity {
	return &PcloudPvminstancesClonePostUnprocessableEntity{}
}

/*
PcloudPvminstancesClonePostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesClonePostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances clone post unprocessable entity response has a 2xx status code
func (o *PcloudPvminstancesClonePostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances clone post unprocessable entity response has a 3xx status code
func (o *PcloudPvminstancesClonePostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances clone post unprocessable entity response has a 4xx status code
func (o *PcloudPvminstancesClonePostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances clone post unprocessable entity response has a 5xx status code
func (o *PcloudPvminstancesClonePostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances clone post unprocessable entity response a status code equal to that given
func (o *PcloudPvminstancesClonePostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudPvminstancesClonePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPvminstancesClonePostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPvminstancesClonePostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesClonePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostInternalServerError creates a PcloudPvminstancesClonePostInternalServerError with default headers values
func NewPcloudPvminstancesClonePostInternalServerError() *PcloudPvminstancesClonePostInternalServerError {
	return &PcloudPvminstancesClonePostInternalServerError{}
}

/*
PcloudPvminstancesClonePostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesClonePostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances clone post internal server error response has a 2xx status code
func (o *PcloudPvminstancesClonePostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances clone post internal server error response has a 3xx status code
func (o *PcloudPvminstancesClonePostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances clone post internal server error response has a 4xx status code
func (o *PcloudPvminstancesClonePostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances clone post internal server error response has a 5xx status code
func (o *PcloudPvminstancesClonePostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances clone post internal server error response a status code equal to that given
func (o *PcloudPvminstancesClonePostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPvminstancesClonePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesClonePostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesClonePostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesClonePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
