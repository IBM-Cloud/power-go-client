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

// ServiceInstanceProvisionReader is a Reader for the ServiceInstanceProvision structure.
type ServiceInstanceProvisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceInstanceProvisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceInstanceProvisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewServiceInstanceProvisionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewServiceInstanceProvisionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceInstanceProvisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewServiceInstanceProvisionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceInstanceProvisionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceInstanceProvisionOK creates a ServiceInstanceProvisionOK with default headers values
func NewServiceInstanceProvisionOK() *ServiceInstanceProvisionOK {
	return &ServiceInstanceProvisionOK{}
}

/*
ServiceInstanceProvisionOK describes a response with status code 200, with default header values.

OK
*/
type ServiceInstanceProvisionOK struct {
	Payload *models.ServiceInstanceProvision
}

// IsSuccess returns true when this service instance provision o k response has a 2xx status code
func (o *ServiceInstanceProvisionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service instance provision o k response has a 3xx status code
func (o *ServiceInstanceProvisionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance provision o k response has a 4xx status code
func (o *ServiceInstanceProvisionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service instance provision o k response has a 5xx status code
func (o *ServiceInstanceProvisionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance provision o k response a status code equal to that given
func (o *ServiceInstanceProvisionOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServiceInstanceProvisionOK) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionOK  %+v", 200, o.Payload)
}

func (o *ServiceInstanceProvisionOK) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionOK  %+v", 200, o.Payload)
}

func (o *ServiceInstanceProvisionOK) GetPayload() *models.ServiceInstanceProvision {
	return o.Payload
}

func (o *ServiceInstanceProvisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstanceProvision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionCreated creates a ServiceInstanceProvisionCreated with default headers values
func NewServiceInstanceProvisionCreated() *ServiceInstanceProvisionCreated {
	return &ServiceInstanceProvisionCreated{}
}

/*
ServiceInstanceProvisionCreated describes a response with status code 201, with default header values.

Created
*/
type ServiceInstanceProvisionCreated struct {
	Payload *models.ServiceInstanceProvision
}

// IsSuccess returns true when this service instance provision created response has a 2xx status code
func (o *ServiceInstanceProvisionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service instance provision created response has a 3xx status code
func (o *ServiceInstanceProvisionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance provision created response has a 4xx status code
func (o *ServiceInstanceProvisionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this service instance provision created response has a 5xx status code
func (o *ServiceInstanceProvisionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance provision created response a status code equal to that given
func (o *ServiceInstanceProvisionCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ServiceInstanceProvisionCreated) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionCreated  %+v", 201, o.Payload)
}

func (o *ServiceInstanceProvisionCreated) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionCreated  %+v", 201, o.Payload)
}

func (o *ServiceInstanceProvisionCreated) GetPayload() *models.ServiceInstanceProvision {
	return o.Payload
}

func (o *ServiceInstanceProvisionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstanceProvision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionAccepted creates a ServiceInstanceProvisionAccepted with default headers values
func NewServiceInstanceProvisionAccepted() *ServiceInstanceProvisionAccepted {
	return &ServiceInstanceProvisionAccepted{}
}

/*
ServiceInstanceProvisionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ServiceInstanceProvisionAccepted struct {
	Payload *models.ServiceInstanceAsyncOperation
}

// IsSuccess returns true when this service instance provision accepted response has a 2xx status code
func (o *ServiceInstanceProvisionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service instance provision accepted response has a 3xx status code
func (o *ServiceInstanceProvisionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance provision accepted response has a 4xx status code
func (o *ServiceInstanceProvisionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this service instance provision accepted response has a 5xx status code
func (o *ServiceInstanceProvisionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance provision accepted response a status code equal to that given
func (o *ServiceInstanceProvisionAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ServiceInstanceProvisionAccepted) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionAccepted  %+v", 202, o.Payload)
}

func (o *ServiceInstanceProvisionAccepted) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionAccepted  %+v", 202, o.Payload)
}

func (o *ServiceInstanceProvisionAccepted) GetPayload() *models.ServiceInstanceAsyncOperation {
	return o.Payload
}

func (o *ServiceInstanceProvisionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstanceAsyncOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionBadRequest creates a ServiceInstanceProvisionBadRequest with default headers values
func NewServiceInstanceProvisionBadRequest() *ServiceInstanceProvisionBadRequest {
	return &ServiceInstanceProvisionBadRequest{}
}

/*
ServiceInstanceProvisionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceInstanceProvisionBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance provision bad request response has a 2xx status code
func (o *ServiceInstanceProvisionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance provision bad request response has a 3xx status code
func (o *ServiceInstanceProvisionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance provision bad request response has a 4xx status code
func (o *ServiceInstanceProvisionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance provision bad request response has a 5xx status code
func (o *ServiceInstanceProvisionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance provision bad request response a status code equal to that given
func (o *ServiceInstanceProvisionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ServiceInstanceProvisionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceInstanceProvisionBadRequest) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceInstanceProvisionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceProvisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionConflict creates a ServiceInstanceProvisionConflict with default headers values
func NewServiceInstanceProvisionConflict() *ServiceInstanceProvisionConflict {
	return &ServiceInstanceProvisionConflict{}
}

/*
ServiceInstanceProvisionConflict describes a response with status code 409, with default header values.

Conflict
*/
type ServiceInstanceProvisionConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance provision conflict response has a 2xx status code
func (o *ServiceInstanceProvisionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance provision conflict response has a 3xx status code
func (o *ServiceInstanceProvisionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance provision conflict response has a 4xx status code
func (o *ServiceInstanceProvisionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance provision conflict response has a 5xx status code
func (o *ServiceInstanceProvisionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance provision conflict response a status code equal to that given
func (o *ServiceInstanceProvisionConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ServiceInstanceProvisionConflict) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionConflict  %+v", 409, o.Payload)
}

func (o *ServiceInstanceProvisionConflict) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionConflict  %+v", 409, o.Payload)
}

func (o *ServiceInstanceProvisionConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceProvisionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionUnprocessableEntity creates a ServiceInstanceProvisionUnprocessableEntity with default headers values
func NewServiceInstanceProvisionUnprocessableEntity() *ServiceInstanceProvisionUnprocessableEntity {
	return &ServiceInstanceProvisionUnprocessableEntity{}
}

/*
ServiceInstanceProvisionUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ServiceInstanceProvisionUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this service instance provision unprocessable entity response has a 2xx status code
func (o *ServiceInstanceProvisionUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service instance provision unprocessable entity response has a 3xx status code
func (o *ServiceInstanceProvisionUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service instance provision unprocessable entity response has a 4xx status code
func (o *ServiceInstanceProvisionUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this service instance provision unprocessable entity response has a 5xx status code
func (o *ServiceInstanceProvisionUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this service instance provision unprocessable entity response a status code equal to that given
func (o *ServiceInstanceProvisionUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *ServiceInstanceProvisionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceInstanceProvisionUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceInstanceProvisionUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceProvisionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
