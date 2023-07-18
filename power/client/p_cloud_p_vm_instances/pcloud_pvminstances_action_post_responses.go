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

// PcloudPvminstancesActionPostReader is a Reader for the PcloudPvminstancesActionPost structure.
type PcloudPvminstancesActionPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesActionPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesActionPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesActionPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesActionPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesActionPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesActionPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesActionPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesActionPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action] pcloud.pvminstances.action.post", response, response.Code())
	}
}

// NewPcloudPvminstancesActionPostOK creates a PcloudPvminstancesActionPostOK with default headers values
func NewPcloudPvminstancesActionPostOK() *PcloudPvminstancesActionPostOK {
	return &PcloudPvminstancesActionPostOK{}
}

/*
PcloudPvminstancesActionPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesActionPostOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud pvminstances action post o k response has a 2xx status code
func (o *PcloudPvminstancesActionPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances action post o k response has a 3xx status code
func (o *PcloudPvminstancesActionPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances action post o k response has a 4xx status code
func (o *PcloudPvminstancesActionPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances action post o k response has a 5xx status code
func (o *PcloudPvminstancesActionPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances action post o k response a status code equal to that given
func (o *PcloudPvminstancesActionPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances action post o k response
func (o *PcloudPvminstancesActionPostOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesActionPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesActionPostOK) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesActionPostOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesActionPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesActionPostBadRequest creates a PcloudPvminstancesActionPostBadRequest with default headers values
func NewPcloudPvminstancesActionPostBadRequest() *PcloudPvminstancesActionPostBadRequest {
	return &PcloudPvminstancesActionPostBadRequest{}
}

/*
PcloudPvminstancesActionPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesActionPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances action post bad request response has a 2xx status code
func (o *PcloudPvminstancesActionPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances action post bad request response has a 3xx status code
func (o *PcloudPvminstancesActionPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances action post bad request response has a 4xx status code
func (o *PcloudPvminstancesActionPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances action post bad request response has a 5xx status code
func (o *PcloudPvminstancesActionPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances action post bad request response a status code equal to that given
func (o *PcloudPvminstancesActionPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances action post bad request response
func (o *PcloudPvminstancesActionPostBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesActionPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesActionPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesActionPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesActionPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesActionPostUnauthorized creates a PcloudPvminstancesActionPostUnauthorized with default headers values
func NewPcloudPvminstancesActionPostUnauthorized() *PcloudPvminstancesActionPostUnauthorized {
	return &PcloudPvminstancesActionPostUnauthorized{}
}

/*
PcloudPvminstancesActionPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesActionPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances action post unauthorized response has a 2xx status code
func (o *PcloudPvminstancesActionPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances action post unauthorized response has a 3xx status code
func (o *PcloudPvminstancesActionPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances action post unauthorized response has a 4xx status code
func (o *PcloudPvminstancesActionPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances action post unauthorized response has a 5xx status code
func (o *PcloudPvminstancesActionPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances action post unauthorized response a status code equal to that given
func (o *PcloudPvminstancesActionPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances action post unauthorized response
func (o *PcloudPvminstancesActionPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesActionPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesActionPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesActionPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesActionPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesActionPostForbidden creates a PcloudPvminstancesActionPostForbidden with default headers values
func NewPcloudPvminstancesActionPostForbidden() *PcloudPvminstancesActionPostForbidden {
	return &PcloudPvminstancesActionPostForbidden{}
}

/*
PcloudPvminstancesActionPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesActionPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances action post forbidden response has a 2xx status code
func (o *PcloudPvminstancesActionPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances action post forbidden response has a 3xx status code
func (o *PcloudPvminstancesActionPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances action post forbidden response has a 4xx status code
func (o *PcloudPvminstancesActionPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances action post forbidden response has a 5xx status code
func (o *PcloudPvminstancesActionPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances action post forbidden response a status code equal to that given
func (o *PcloudPvminstancesActionPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances action post forbidden response
func (o *PcloudPvminstancesActionPostForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesActionPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesActionPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesActionPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesActionPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesActionPostNotFound creates a PcloudPvminstancesActionPostNotFound with default headers values
func NewPcloudPvminstancesActionPostNotFound() *PcloudPvminstancesActionPostNotFound {
	return &PcloudPvminstancesActionPostNotFound{}
}

/*
PcloudPvminstancesActionPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesActionPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances action post not found response has a 2xx status code
func (o *PcloudPvminstancesActionPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances action post not found response has a 3xx status code
func (o *PcloudPvminstancesActionPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances action post not found response has a 4xx status code
func (o *PcloudPvminstancesActionPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances action post not found response has a 5xx status code
func (o *PcloudPvminstancesActionPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances action post not found response a status code equal to that given
func (o *PcloudPvminstancesActionPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances action post not found response
func (o *PcloudPvminstancesActionPostNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesActionPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesActionPostNotFound) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesActionPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesActionPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesActionPostConflict creates a PcloudPvminstancesActionPostConflict with default headers values
func NewPcloudPvminstancesActionPostConflict() *PcloudPvminstancesActionPostConflict {
	return &PcloudPvminstancesActionPostConflict{}
}

/*
PcloudPvminstancesActionPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesActionPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances action post conflict response has a 2xx status code
func (o *PcloudPvminstancesActionPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances action post conflict response has a 3xx status code
func (o *PcloudPvminstancesActionPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances action post conflict response has a 4xx status code
func (o *PcloudPvminstancesActionPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances action post conflict response has a 5xx status code
func (o *PcloudPvminstancesActionPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances action post conflict response a status code equal to that given
func (o *PcloudPvminstancesActionPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud pvminstances action post conflict response
func (o *PcloudPvminstancesActionPostConflict) Code() int {
	return 409
}

func (o *PcloudPvminstancesActionPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPvminstancesActionPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPvminstancesActionPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesActionPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesActionPostInternalServerError creates a PcloudPvminstancesActionPostInternalServerError with default headers values
func NewPcloudPvminstancesActionPostInternalServerError() *PcloudPvminstancesActionPostInternalServerError {
	return &PcloudPvminstancesActionPostInternalServerError{}
}

/*
PcloudPvminstancesActionPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesActionPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances action post internal server error response has a 2xx status code
func (o *PcloudPvminstancesActionPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances action post internal server error response has a 3xx status code
func (o *PcloudPvminstancesActionPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances action post internal server error response has a 4xx status code
func (o *PcloudPvminstancesActionPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances action post internal server error response has a 5xx status code
func (o *PcloudPvminstancesActionPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances action post internal server error response a status code equal to that given
func (o *PcloudPvminstancesActionPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances action post internal server error response
func (o *PcloudPvminstancesActionPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesActionPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesActionPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action][%d] pcloudPvminstancesActionPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesActionPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesActionPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
