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

// PcloudPvminstancesDeleteReader is a Reader for the PcloudPvminstancesDelete structure.
type PcloudPvminstancesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudPvminstancesDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesDeleteOK creates a PcloudPvminstancesDeleteOK with default headers values
func NewPcloudPvminstancesDeleteOK() *PcloudPvminstancesDeleteOK {
	return &PcloudPvminstancesDeleteOK{}
}

/*
PcloudPvminstancesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud pvminstances delete o k response has a 2xx status code
func (o *PcloudPvminstancesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances delete o k response has a 3xx status code
func (o *PcloudPvminstancesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances delete o k response has a 4xx status code
func (o *PcloudPvminstancesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances delete o k response has a 5xx status code
func (o *PcloudPvminstancesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances delete o k response a status code equal to that given
func (o *PcloudPvminstancesDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudPvminstancesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteBadRequest creates a PcloudPvminstancesDeleteBadRequest with default headers values
func NewPcloudPvminstancesDeleteBadRequest() *PcloudPvminstancesDeleteBadRequest {
	return &PcloudPvminstancesDeleteBadRequest{}
}

/*
PcloudPvminstancesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances delete bad request response has a 2xx status code
func (o *PcloudPvminstancesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances delete bad request response has a 3xx status code
func (o *PcloudPvminstancesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances delete bad request response has a 4xx status code
func (o *PcloudPvminstancesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances delete bad request response has a 5xx status code
func (o *PcloudPvminstancesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances delete bad request response a status code equal to that given
func (o *PcloudPvminstancesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPvminstancesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteUnauthorized creates a PcloudPvminstancesDeleteUnauthorized with default headers values
func NewPcloudPvminstancesDeleteUnauthorized() *PcloudPvminstancesDeleteUnauthorized {
	return &PcloudPvminstancesDeleteUnauthorized{}
}

/*
PcloudPvminstancesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances delete unauthorized response has a 2xx status code
func (o *PcloudPvminstancesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances delete unauthorized response has a 3xx status code
func (o *PcloudPvminstancesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances delete unauthorized response has a 4xx status code
func (o *PcloudPvminstancesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances delete unauthorized response has a 5xx status code
func (o *PcloudPvminstancesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances delete unauthorized response a status code equal to that given
func (o *PcloudPvminstancesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudPvminstancesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteForbidden creates a PcloudPvminstancesDeleteForbidden with default headers values
func NewPcloudPvminstancesDeleteForbidden() *PcloudPvminstancesDeleteForbidden {
	return &PcloudPvminstancesDeleteForbidden{}
}

/*
PcloudPvminstancesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances delete forbidden response has a 2xx status code
func (o *PcloudPvminstancesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances delete forbidden response has a 3xx status code
func (o *PcloudPvminstancesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances delete forbidden response has a 4xx status code
func (o *PcloudPvminstancesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances delete forbidden response has a 5xx status code
func (o *PcloudPvminstancesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances delete forbidden response a status code equal to that given
func (o *PcloudPvminstancesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudPvminstancesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteNotFound creates a PcloudPvminstancesDeleteNotFound with default headers values
func NewPcloudPvminstancesDeleteNotFound() *PcloudPvminstancesDeleteNotFound {
	return &PcloudPvminstancesDeleteNotFound{}
}

/*
PcloudPvminstancesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances delete not found response has a 2xx status code
func (o *PcloudPvminstancesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances delete not found response has a 3xx status code
func (o *PcloudPvminstancesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances delete not found response has a 4xx status code
func (o *PcloudPvminstancesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances delete not found response has a 5xx status code
func (o *PcloudPvminstancesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances delete not found response a status code equal to that given
func (o *PcloudPvminstancesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudPvminstancesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteGone creates a PcloudPvminstancesDeleteGone with default headers values
func NewPcloudPvminstancesDeleteGone() *PcloudPvminstancesDeleteGone {
	return &PcloudPvminstancesDeleteGone{}
}

/*
PcloudPvminstancesDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudPvminstancesDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances delete gone response has a 2xx status code
func (o *PcloudPvminstancesDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances delete gone response has a 3xx status code
func (o *PcloudPvminstancesDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances delete gone response has a 4xx status code
func (o *PcloudPvminstancesDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances delete gone response has a 5xx status code
func (o *PcloudPvminstancesDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances delete gone response a status code equal to that given
func (o *PcloudPvminstancesDeleteGone) IsCode(code int) bool {
	return code == 410
}

func (o *PcloudPvminstancesDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudPvminstancesDeleteGone) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudPvminstancesDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteInternalServerError creates a PcloudPvminstancesDeleteInternalServerError with default headers values
func NewPcloudPvminstancesDeleteInternalServerError() *PcloudPvminstancesDeleteInternalServerError {
	return &PcloudPvminstancesDeleteInternalServerError{}
}

/*
PcloudPvminstancesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances delete internal server error response has a 2xx status code
func (o *PcloudPvminstancesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances delete internal server error response has a 3xx status code
func (o *PcloudPvminstancesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances delete internal server error response has a 4xx status code
func (o *PcloudPvminstancesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances delete internal server error response has a 5xx status code
func (o *PcloudPvminstancesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances delete internal server error response a status code equal to that given
func (o *PcloudPvminstancesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPvminstancesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
