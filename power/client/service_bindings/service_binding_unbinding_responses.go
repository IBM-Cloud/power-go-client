// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBindingUnbindingReader is a Reader for the ServiceBindingUnbinding structure.
type ServiceBindingUnbindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingUnbindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBindingUnbindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewServiceBindingUnbindingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBindingUnbindingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBindingUnbindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBindingUnbindingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBindingUnbindingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewServiceBindingUnbindingGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}] serviceBinding.unbinding", response, response.Code())
	}
}

// NewServiceBindingUnbindingOK creates a ServiceBindingUnbindingOK with default headers values
func NewServiceBindingUnbindingOK() *ServiceBindingUnbindingOK {
	return &ServiceBindingUnbindingOK{}
}

/*
ServiceBindingUnbindingOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBindingUnbindingOK struct {
	Payload models.Object
}

// IsSuccess returns true when this service binding unbinding o k response has a 2xx status code
func (o *ServiceBindingUnbindingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service binding unbinding o k response has a 3xx status code
func (o *ServiceBindingUnbindingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding unbinding o k response has a 4xx status code
func (o *ServiceBindingUnbindingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service binding unbinding o k response has a 5xx status code
func (o *ServiceBindingUnbindingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding unbinding o k response a status code equal to that given
func (o *ServiceBindingUnbindingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service binding unbinding o k response
func (o *ServiceBindingUnbindingOK) Code() int {
	return 200
}

func (o *ServiceBindingUnbindingOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingUnbindingOK) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingUnbindingOK) GetPayload() models.Object {
	return o.Payload
}

func (o *ServiceBindingUnbindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingAccepted creates a ServiceBindingUnbindingAccepted with default headers values
func NewServiceBindingUnbindingAccepted() *ServiceBindingUnbindingAccepted {
	return &ServiceBindingUnbindingAccepted{}
}

/*
ServiceBindingUnbindingAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ServiceBindingUnbindingAccepted struct {
	Payload *models.AsyncOperation
}

// IsSuccess returns true when this service binding unbinding accepted response has a 2xx status code
func (o *ServiceBindingUnbindingAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service binding unbinding accepted response has a 3xx status code
func (o *ServiceBindingUnbindingAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding unbinding accepted response has a 4xx status code
func (o *ServiceBindingUnbindingAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this service binding unbinding accepted response has a 5xx status code
func (o *ServiceBindingUnbindingAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding unbinding accepted response a status code equal to that given
func (o *ServiceBindingUnbindingAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the service binding unbinding accepted response
func (o *ServiceBindingUnbindingAccepted) Code() int {
	return 202
}

func (o *ServiceBindingUnbindingAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingAccepted  %+v", 202, o.Payload)
}

func (o *ServiceBindingUnbindingAccepted) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingAccepted  %+v", 202, o.Payload)
}

func (o *ServiceBindingUnbindingAccepted) GetPayload() *models.AsyncOperation {
	return o.Payload
}

func (o *ServiceBindingUnbindingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingBadRequest creates a ServiceBindingUnbindingBadRequest with default headers values
func NewServiceBindingUnbindingBadRequest() *ServiceBindingUnbindingBadRequest {
	return &ServiceBindingUnbindingBadRequest{}
}

/*
ServiceBindingUnbindingBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBindingUnbindingBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding unbinding bad request response has a 2xx status code
func (o *ServiceBindingUnbindingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding unbinding bad request response has a 3xx status code
func (o *ServiceBindingUnbindingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding unbinding bad request response has a 4xx status code
func (o *ServiceBindingUnbindingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding unbinding bad request response has a 5xx status code
func (o *ServiceBindingUnbindingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding unbinding bad request response a status code equal to that given
func (o *ServiceBindingUnbindingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service binding unbinding bad request response
func (o *ServiceBindingUnbindingBadRequest) Code() int {
	return 400
}

func (o *ServiceBindingUnbindingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBindingUnbindingBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBindingUnbindingBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingUnbindingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingUnauthorized creates a ServiceBindingUnbindingUnauthorized with default headers values
func NewServiceBindingUnbindingUnauthorized() *ServiceBindingUnbindingUnauthorized {
	return &ServiceBindingUnbindingUnauthorized{}
}

/*
ServiceBindingUnbindingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBindingUnbindingUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding unbinding unauthorized response has a 2xx status code
func (o *ServiceBindingUnbindingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding unbinding unauthorized response has a 3xx status code
func (o *ServiceBindingUnbindingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding unbinding unauthorized response has a 4xx status code
func (o *ServiceBindingUnbindingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding unbinding unauthorized response has a 5xx status code
func (o *ServiceBindingUnbindingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding unbinding unauthorized response a status code equal to that given
func (o *ServiceBindingUnbindingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service binding unbinding unauthorized response
func (o *ServiceBindingUnbindingUnauthorized) Code() int {
	return 401
}

func (o *ServiceBindingUnbindingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBindingUnbindingUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBindingUnbindingUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingUnbindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingForbidden creates a ServiceBindingUnbindingForbidden with default headers values
func NewServiceBindingUnbindingForbidden() *ServiceBindingUnbindingForbidden {
	return &ServiceBindingUnbindingForbidden{}
}

/*
ServiceBindingUnbindingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBindingUnbindingForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding unbinding forbidden response has a 2xx status code
func (o *ServiceBindingUnbindingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding unbinding forbidden response has a 3xx status code
func (o *ServiceBindingUnbindingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding unbinding forbidden response has a 4xx status code
func (o *ServiceBindingUnbindingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding unbinding forbidden response has a 5xx status code
func (o *ServiceBindingUnbindingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding unbinding forbidden response a status code equal to that given
func (o *ServiceBindingUnbindingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service binding unbinding forbidden response
func (o *ServiceBindingUnbindingForbidden) Code() int {
	return 403
}

func (o *ServiceBindingUnbindingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBindingUnbindingForbidden) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBindingUnbindingForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingUnbindingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingNotFound creates a ServiceBindingUnbindingNotFound with default headers values
func NewServiceBindingUnbindingNotFound() *ServiceBindingUnbindingNotFound {
	return &ServiceBindingUnbindingNotFound{}
}

/*
ServiceBindingUnbindingNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBindingUnbindingNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding unbinding not found response has a 2xx status code
func (o *ServiceBindingUnbindingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding unbinding not found response has a 3xx status code
func (o *ServiceBindingUnbindingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding unbinding not found response has a 4xx status code
func (o *ServiceBindingUnbindingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding unbinding not found response has a 5xx status code
func (o *ServiceBindingUnbindingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding unbinding not found response a status code equal to that given
func (o *ServiceBindingUnbindingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service binding unbinding not found response
func (o *ServiceBindingUnbindingNotFound) Code() int {
	return 404
}

func (o *ServiceBindingUnbindingNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBindingUnbindingNotFound) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBindingUnbindingNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingUnbindingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingGone creates a ServiceBindingUnbindingGone with default headers values
func NewServiceBindingUnbindingGone() *ServiceBindingUnbindingGone {
	return &ServiceBindingUnbindingGone{}
}

/*
ServiceBindingUnbindingGone describes a response with status code 410, with default header values.

Gone
*/
type ServiceBindingUnbindingGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding unbinding gone response has a 2xx status code
func (o *ServiceBindingUnbindingGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding unbinding gone response has a 3xx status code
func (o *ServiceBindingUnbindingGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding unbinding gone response has a 4xx status code
func (o *ServiceBindingUnbindingGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding unbinding gone response has a 5xx status code
func (o *ServiceBindingUnbindingGone) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding unbinding gone response a status code equal to that given
func (o *ServiceBindingUnbindingGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the service binding unbinding gone response
func (o *ServiceBindingUnbindingGone) Code() int {
	return 410
}

func (o *ServiceBindingUnbindingGone) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingGone  %+v", 410, o.Payload)
}

func (o *ServiceBindingUnbindingGone) String() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingGone  %+v", 410, o.Payload)
}

func (o *ServiceBindingUnbindingGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingUnbindingGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
