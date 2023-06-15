// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsPutReader is a Reader for the PcloudCloudconnectionsPut structure.
type PcloudCloudconnectionsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPcloudCloudconnectionsPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsPutRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudCloudconnectionsPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudCloudconnectionsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPcloudCloudconnectionsPutServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudconnectionsPutOK creates a PcloudCloudconnectionsPutOK with default headers values
func NewPcloudCloudconnectionsPutOK() *PcloudCloudconnectionsPutOK {
	return &PcloudCloudconnectionsPutOK{}
}

/*
PcloudCloudconnectionsPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsPutOK struct {
	Payload *models.CloudConnection
}

// IsSuccess returns true when this pcloud cloudconnections put o k response has a 2xx status code
func (o *PcloudCloudconnectionsPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections put o k response has a 3xx status code
func (o *PcloudCloudconnectionsPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put o k response has a 4xx status code
func (o *PcloudCloudconnectionsPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections put o k response has a 5xx status code
func (o *PcloudCloudconnectionsPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put o k response a status code equal to that given
func (o *PcloudCloudconnectionsPutOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudconnectionsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsPutOK) GetPayload() *models.CloudConnection {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutAccepted creates a PcloudCloudconnectionsPutAccepted with default headers values
func NewPcloudCloudconnectionsPutAccepted() *PcloudCloudconnectionsPutAccepted {
	return &PcloudCloudconnectionsPutAccepted{}
}

/*
PcloudCloudconnectionsPutAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsPutAccepted struct {
	Payload *models.JobReference
}

// IsSuccess returns true when this pcloud cloudconnections put accepted response has a 2xx status code
func (o *PcloudCloudconnectionsPutAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections put accepted response has a 3xx status code
func (o *PcloudCloudconnectionsPutAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put accepted response has a 4xx status code
func (o *PcloudCloudconnectionsPutAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections put accepted response has a 5xx status code
func (o *PcloudCloudconnectionsPutAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put accepted response a status code equal to that given
func (o *PcloudCloudconnectionsPutAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PcloudCloudconnectionsPutAccepted) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudconnectionsPutAccepted) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudconnectionsPutAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutBadRequest creates a PcloudCloudconnectionsPutBadRequest with default headers values
func NewPcloudCloudconnectionsPutBadRequest() *PcloudCloudconnectionsPutBadRequest {
	return &PcloudCloudconnectionsPutBadRequest{}
}

/*
PcloudCloudconnectionsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put bad request response has a 2xx status code
func (o *PcloudCloudconnectionsPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put bad request response has a 3xx status code
func (o *PcloudCloudconnectionsPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put bad request response has a 4xx status code
func (o *PcloudCloudconnectionsPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections put bad request response has a 5xx status code
func (o *PcloudCloudconnectionsPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put bad request response a status code equal to that given
func (o *PcloudCloudconnectionsPutBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudconnectionsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutUnauthorized creates a PcloudCloudconnectionsPutUnauthorized with default headers values
func NewPcloudCloudconnectionsPutUnauthorized() *PcloudCloudconnectionsPutUnauthorized {
	return &PcloudCloudconnectionsPutUnauthorized{}
}

/*
PcloudCloudconnectionsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put unauthorized response has a 2xx status code
func (o *PcloudCloudconnectionsPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put unauthorized response has a 3xx status code
func (o *PcloudCloudconnectionsPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put unauthorized response has a 4xx status code
func (o *PcloudCloudconnectionsPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections put unauthorized response has a 5xx status code
func (o *PcloudCloudconnectionsPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put unauthorized response a status code equal to that given
func (o *PcloudCloudconnectionsPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudconnectionsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutNotFound creates a PcloudCloudconnectionsPutNotFound with default headers values
func NewPcloudCloudconnectionsPutNotFound() *PcloudCloudconnectionsPutNotFound {
	return &PcloudCloudconnectionsPutNotFound{}
}

/*
PcloudCloudconnectionsPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put not found response has a 2xx status code
func (o *PcloudCloudconnectionsPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put not found response has a 3xx status code
func (o *PcloudCloudconnectionsPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put not found response has a 4xx status code
func (o *PcloudCloudconnectionsPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections put not found response has a 5xx status code
func (o *PcloudCloudconnectionsPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put not found response a status code equal to that given
func (o *PcloudCloudconnectionsPutNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudCloudconnectionsPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsPutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutMethodNotAllowed creates a PcloudCloudconnectionsPutMethodNotAllowed with default headers values
func NewPcloudCloudconnectionsPutMethodNotAllowed() *PcloudCloudconnectionsPutMethodNotAllowed {
	return &PcloudCloudconnectionsPutMethodNotAllowed{}
}

/*
PcloudCloudconnectionsPutMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type PcloudCloudconnectionsPutMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put method not allowed response has a 2xx status code
func (o *PcloudCloudconnectionsPutMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put method not allowed response has a 3xx status code
func (o *PcloudCloudconnectionsPutMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put method not allowed response has a 4xx status code
func (o *PcloudCloudconnectionsPutMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections put method not allowed response has a 5xx status code
func (o *PcloudCloudconnectionsPutMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put method not allowed response a status code equal to that given
func (o *PcloudCloudconnectionsPutMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PcloudCloudconnectionsPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PcloudCloudconnectionsPutMethodNotAllowed) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PcloudCloudconnectionsPutMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutRequestTimeout creates a PcloudCloudconnectionsPutRequestTimeout with default headers values
func NewPcloudCloudconnectionsPutRequestTimeout() *PcloudCloudconnectionsPutRequestTimeout {
	return &PcloudCloudconnectionsPutRequestTimeout{}
}

/*
PcloudCloudconnectionsPutRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsPutRequestTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put request timeout response has a 2xx status code
func (o *PcloudCloudconnectionsPutRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put request timeout response has a 3xx status code
func (o *PcloudCloudconnectionsPutRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put request timeout response has a 4xx status code
func (o *PcloudCloudconnectionsPutRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections put request timeout response has a 5xx status code
func (o *PcloudCloudconnectionsPutRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put request timeout response a status code equal to that given
func (o *PcloudCloudconnectionsPutRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PcloudCloudconnectionsPutRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsPutRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsPutRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutConflict creates a PcloudCloudconnectionsPutConflict with default headers values
func NewPcloudCloudconnectionsPutConflict() *PcloudCloudconnectionsPutConflict {
	return &PcloudCloudconnectionsPutConflict{}
}

/*
PcloudCloudconnectionsPutConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudCloudconnectionsPutConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put conflict response has a 2xx status code
func (o *PcloudCloudconnectionsPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put conflict response has a 3xx status code
func (o *PcloudCloudconnectionsPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put conflict response has a 4xx status code
func (o *PcloudCloudconnectionsPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections put conflict response has a 5xx status code
func (o *PcloudCloudconnectionsPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put conflict response a status code equal to that given
func (o *PcloudCloudconnectionsPutConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PcloudCloudconnectionsPutConflict) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutConflict  %+v", 409, o.Payload)
}

func (o *PcloudCloudconnectionsPutConflict) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutConflict  %+v", 409, o.Payload)
}

func (o *PcloudCloudconnectionsPutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutUnprocessableEntity creates a PcloudCloudconnectionsPutUnprocessableEntity with default headers values
func NewPcloudCloudconnectionsPutUnprocessableEntity() *PcloudCloudconnectionsPutUnprocessableEntity {
	return &PcloudCloudconnectionsPutUnprocessableEntity{}
}

/*
PcloudCloudconnectionsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudconnectionsPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put unprocessable entity response has a 2xx status code
func (o *PcloudCloudconnectionsPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put unprocessable entity response has a 3xx status code
func (o *PcloudCloudconnectionsPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put unprocessable entity response has a 4xx status code
func (o *PcloudCloudconnectionsPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections put unprocessable entity response has a 5xx status code
func (o *PcloudCloudconnectionsPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections put unprocessable entity response a status code equal to that given
func (o *PcloudCloudconnectionsPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudCloudconnectionsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudconnectionsPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudconnectionsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutInternalServerError creates a PcloudCloudconnectionsPutInternalServerError with default headers values
func NewPcloudCloudconnectionsPutInternalServerError() *PcloudCloudconnectionsPutInternalServerError {
	return &PcloudCloudconnectionsPutInternalServerError{}
}

/*
PcloudCloudconnectionsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put internal server error response has a 2xx status code
func (o *PcloudCloudconnectionsPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put internal server error response has a 3xx status code
func (o *PcloudCloudconnectionsPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put internal server error response has a 4xx status code
func (o *PcloudCloudconnectionsPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections put internal server error response has a 5xx status code
func (o *PcloudCloudconnectionsPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections put internal server error response a status code equal to that given
func (o *PcloudCloudconnectionsPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudconnectionsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPutServiceUnavailable creates a PcloudCloudconnectionsPutServiceUnavailable with default headers values
func NewPcloudCloudconnectionsPutServiceUnavailable() *PcloudCloudconnectionsPutServiceUnavailable {
	return &PcloudCloudconnectionsPutServiceUnavailable{}
}

/*
PcloudCloudconnectionsPutServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type PcloudCloudconnectionsPutServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections put service unavailable response has a 2xx status code
func (o *PcloudCloudconnectionsPutServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections put service unavailable response has a 3xx status code
func (o *PcloudCloudconnectionsPutServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections put service unavailable response has a 4xx status code
func (o *PcloudCloudconnectionsPutServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections put service unavailable response has a 5xx status code
func (o *PcloudCloudconnectionsPutServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections put service unavailable response a status code equal to that given
func (o *PcloudCloudconnectionsPutServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PcloudCloudconnectionsPutServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PcloudCloudconnectionsPutServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsPutServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PcloudCloudconnectionsPutServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPutServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
