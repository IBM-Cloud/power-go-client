// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceInstanceUpdateReader is a Reader for the ServiceInstanceUpdate structure.
type ServiceInstanceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceInstanceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceInstanceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewServiceInstanceUpdateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceInstanceUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceInstanceUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceInstanceUpdateOK creates a ServiceInstanceUpdateOK with default headers values
func NewServiceInstanceUpdateOK() *ServiceInstanceUpdateOK {
	return &ServiceInstanceUpdateOK{}
}

/*
ServiceInstanceUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ServiceInstanceUpdateOK struct {
	Payload models.Object
}

// IsSuccess returns true when this service instance update o k response has a 2xx status code
func (o *ServiceInstanceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service instance update o k response has a 3xx status code
func (o *ServiceInstanceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance update o k response has a 4xx status code
func (o *ServiceInstanceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service instance update o k response has a 5xx status code
func (o *ServiceInstanceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance update o k response a status code equal to that given
func (o *ServiceInstanceUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service instance update o k response
func (o *ServiceInstanceUpdateOK) Code() int {
	return 200
}

func (o *ServiceInstanceUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateOK  %+v", 200, o.Payload)
}

func (o *ServiceInstanceUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateOK  %+v", 200, o.Payload)
}

func (o *ServiceInstanceUpdateOK) GetPayload() models.Object {
	return o.Payload
}

func (o *ServiceInstanceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdateAccepted creates a ServiceInstanceUpdateAccepted with default headers values
func NewServiceInstanceUpdateAccepted() *ServiceInstanceUpdateAccepted {
	return &ServiceInstanceUpdateAccepted{}
}

/*
ServiceInstanceUpdateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ServiceInstanceUpdateAccepted struct {
	Payload *models.ServiceInstanceAsyncOperation
}

// IsSuccess returns true when this service instance update accepted response has a 2xx status code
func (o *ServiceInstanceUpdateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service instance update accepted response has a 3xx status code
func (o *ServiceInstanceUpdateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance update accepted response has a 4xx status code
func (o *ServiceInstanceUpdateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this service instance update accepted response has a 5xx status code
func (o *ServiceInstanceUpdateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance update accepted response a status code equal to that given
func (o *ServiceInstanceUpdateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the service instance update accepted response
func (o *ServiceInstanceUpdateAccepted) Code() int {
	return 202
}

func (o *ServiceInstanceUpdateAccepted) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateAccepted  %+v", 202, o.Payload)
}

func (o *ServiceInstanceUpdateAccepted) String() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateAccepted  %+v", 202, o.Payload)
}

func (o *ServiceInstanceUpdateAccepted) GetPayload() *models.ServiceInstanceAsyncOperation {
	return o.Payload
}

func (o *ServiceInstanceUpdateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstanceAsyncOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdateBadRequest creates a ServiceInstanceUpdateBadRequest with default headers values
func NewServiceInstanceUpdateBadRequest() *ServiceInstanceUpdateBadRequest {
	return &ServiceInstanceUpdateBadRequest{}
}

/*
ServiceInstanceUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceInstanceUpdateBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance update bad request response has a 2xx status code
func (o *ServiceInstanceUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance update bad request response has a 3xx status code
func (o *ServiceInstanceUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance update bad request response has a 4xx status code
func (o *ServiceInstanceUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance update bad request response has a 5xx status code
func (o *ServiceInstanceUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance update bad request response a status code equal to that given
func (o *ServiceInstanceUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service instance update bad request response
func (o *ServiceInstanceUpdateBadRequest) Code() int {
	return 400
}

func (o *ServiceInstanceUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceInstanceUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceInstanceUpdateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdateUnprocessableEntity creates a ServiceInstanceUpdateUnprocessableEntity with default headers values
func NewServiceInstanceUpdateUnprocessableEntity() *ServiceInstanceUpdateUnprocessableEntity {
	return &ServiceInstanceUpdateUnprocessableEntity{}
}

/*
ServiceInstanceUpdateUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable entity
*/
type ServiceInstanceUpdateUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance update unprocessable entity response has a 2xx status code
func (o *ServiceInstanceUpdateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance update unprocessable entity response has a 3xx status code
func (o *ServiceInstanceUpdateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance update unprocessable entity response has a 4xx status code
func (o *ServiceInstanceUpdateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance update unprocessable entity response has a 5xx status code
func (o *ServiceInstanceUpdateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance update unprocessable entity response a status code equal to that given
func (o *ServiceInstanceUpdateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the service instance update unprocessable entity response
func (o *ServiceInstanceUpdateUnprocessableEntity) Code() int {
	return 422
}

func (o *ServiceInstanceUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceInstanceUpdateUnprocessableEntity) String() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceInstanceUpdateUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
