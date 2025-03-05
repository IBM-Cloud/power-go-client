// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesOperationsPostReader is a Reader for the PcloudPvminstancesOperationsPost structure.
type PcloudPvminstancesOperationsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesOperationsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesOperationsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesOperationsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesOperationsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesOperationsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesOperationsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesOperationsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPvminstancesOperationsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesOperationsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations] pcloud.pvminstances.operations.post", response, response.Code())
	}
}

// NewPcloudPvminstancesOperationsPostOK creates a PcloudPvminstancesOperationsPostOK with default headers values
func NewPcloudPvminstancesOperationsPostOK() *PcloudPvminstancesOperationsPostOK {
	return &PcloudPvminstancesOperationsPostOK{}
}

/*
PcloudPvminstancesOperationsPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesOperationsPostOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud pvminstances operations post o k response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances operations post o k response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post o k response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances operations post o k response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post o k response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances operations post o k response
func (o *PcloudPvminstancesOperationsPostOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesOperationsPostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostOK %s", 200, payload)
}

func (o *PcloudPvminstancesOperationsPostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostOK %s", 200, payload)
}

func (o *PcloudPvminstancesOperationsPostOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostBadRequest creates a PcloudPvminstancesOperationsPostBadRequest with default headers values
func NewPcloudPvminstancesOperationsPostBadRequest() *PcloudPvminstancesOperationsPostBadRequest {
	return &PcloudPvminstancesOperationsPostBadRequest{}
}

/*
PcloudPvminstancesOperationsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesOperationsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post bad request response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post bad request response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post bad request response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post bad request response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post bad request response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances operations post bad request response
func (o *PcloudPvminstancesOperationsPostBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesOperationsPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesOperationsPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesOperationsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostUnauthorized creates a PcloudPvminstancesOperationsPostUnauthorized with default headers values
func NewPcloudPvminstancesOperationsPostUnauthorized() *PcloudPvminstancesOperationsPostUnauthorized {
	return &PcloudPvminstancesOperationsPostUnauthorized{}
}

/*
PcloudPvminstancesOperationsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesOperationsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post unauthorized response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post unauthorized response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post unauthorized response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post unauthorized response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post unauthorized response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances operations post unauthorized response
func (o *PcloudPvminstancesOperationsPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesOperationsPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesOperationsPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesOperationsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostForbidden creates a PcloudPvminstancesOperationsPostForbidden with default headers values
func NewPcloudPvminstancesOperationsPostForbidden() *PcloudPvminstancesOperationsPostForbidden {
	return &PcloudPvminstancesOperationsPostForbidden{}
}

/*
PcloudPvminstancesOperationsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesOperationsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post forbidden response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post forbidden response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post forbidden response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post forbidden response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post forbidden response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances operations post forbidden response
func (o *PcloudPvminstancesOperationsPostForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesOperationsPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesOperationsPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesOperationsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostNotFound creates a PcloudPvminstancesOperationsPostNotFound with default headers values
func NewPcloudPvminstancesOperationsPostNotFound() *PcloudPvminstancesOperationsPostNotFound {
	return &PcloudPvminstancesOperationsPostNotFound{}
}

/*
PcloudPvminstancesOperationsPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesOperationsPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post not found response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post not found response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post not found response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post not found response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post not found response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances operations post not found response
func (o *PcloudPvminstancesOperationsPostNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesOperationsPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesOperationsPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesOperationsPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostConflict creates a PcloudPvminstancesOperationsPostConflict with default headers values
func NewPcloudPvminstancesOperationsPostConflict() *PcloudPvminstancesOperationsPostConflict {
	return &PcloudPvminstancesOperationsPostConflict{}
}

/*
PcloudPvminstancesOperationsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesOperationsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post conflict response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post conflict response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post conflict response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post conflict response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post conflict response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud pvminstances operations post conflict response
func (o *PcloudPvminstancesOperationsPostConflict) Code() int {
	return 409
}

func (o *PcloudPvminstancesOperationsPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostConflict %s", 409, payload)
}

func (o *PcloudPvminstancesOperationsPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostConflict %s", 409, payload)
}

func (o *PcloudPvminstancesOperationsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostUnprocessableEntity creates a PcloudPvminstancesOperationsPostUnprocessableEntity with default headers values
func NewPcloudPvminstancesOperationsPostUnprocessableEntity() *PcloudPvminstancesOperationsPostUnprocessableEntity {
	return &PcloudPvminstancesOperationsPostUnprocessableEntity{}
}

/*
PcloudPvminstancesOperationsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesOperationsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post unprocessable entity response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post unprocessable entity response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post unprocessable entity response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post unprocessable entity response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post unprocessable entity response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud pvminstances operations post unprocessable entity response
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostInternalServerError creates a PcloudPvminstancesOperationsPostInternalServerError with default headers values
func NewPcloudPvminstancesOperationsPostInternalServerError() *PcloudPvminstancesOperationsPostInternalServerError {
	return &PcloudPvminstancesOperationsPostInternalServerError{}
}

/*
PcloudPvminstancesOperationsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesOperationsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post internal server error response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post internal server error response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post internal server error response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances operations post internal server error response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances operations post internal server error response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances operations post internal server error response
func (o *PcloudPvminstancesOperationsPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesOperationsPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesOperationsPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesOperationsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
