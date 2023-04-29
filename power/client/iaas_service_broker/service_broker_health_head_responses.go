// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerHealthHeadReader is a Reader for the ServiceBrokerHealthHead structure.
type ServiceBrokerHealthHeadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerHealthHeadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerHealthHeadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerHealthHeadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerHealthHeadOK creates a ServiceBrokerHealthHeadOK with default headers values
func NewServiceBrokerHealthHeadOK() *ServiceBrokerHealthHeadOK {
	return &ServiceBrokerHealthHeadOK{}
}

/*
ServiceBrokerHealthHeadOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerHealthHeadOK struct {
	Payload *models.Health
}

// IsSuccess returns true when this service broker health head o k response has a 2xx status code
func (o *ServiceBrokerHealthHeadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker health head o k response has a 3xx status code
func (o *ServiceBrokerHealthHeadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker health head o k response has a 4xx status code
func (o *ServiceBrokerHealthHeadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker health head o k response has a 5xx status code
func (o *ServiceBrokerHealthHeadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker health head o k response a status code equal to that given
func (o *ServiceBrokerHealthHeadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker health head o k response
func (o *ServiceBrokerHealthHeadOK) Code() int {
	return 200
}

func (o *ServiceBrokerHealthHeadOK) Error() string {
	return fmt.Sprintf("[HEAD /broker/v1/health][%d] serviceBrokerHealthHeadOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerHealthHeadOK) String() string {
	return fmt.Sprintf("[HEAD /broker/v1/health][%d] serviceBrokerHealthHeadOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerHealthHeadOK) GetPayload() *models.Health {
	return o.Payload
}

func (o *ServiceBrokerHealthHeadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Health)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerHealthHeadBadRequest creates a ServiceBrokerHealthHeadBadRequest with default headers values
func NewServiceBrokerHealthHeadBadRequest() *ServiceBrokerHealthHeadBadRequest {
	return &ServiceBrokerHealthHeadBadRequest{}
}

/*
ServiceBrokerHealthHeadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerHealthHeadBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker health head bad request response has a 2xx status code
func (o *ServiceBrokerHealthHeadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker health head bad request response has a 3xx status code
func (o *ServiceBrokerHealthHeadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker health head bad request response has a 4xx status code
func (o *ServiceBrokerHealthHeadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker health head bad request response has a 5xx status code
func (o *ServiceBrokerHealthHeadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker health head bad request response a status code equal to that given
func (o *ServiceBrokerHealthHeadBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker health head bad request response
func (o *ServiceBrokerHealthHeadBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerHealthHeadBadRequest) Error() string {
	return fmt.Sprintf("[HEAD /broker/v1/health][%d] serviceBrokerHealthHeadBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerHealthHeadBadRequest) String() string {
	return fmt.Sprintf("[HEAD /broker/v1/health][%d] serviceBrokerHealthHeadBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerHealthHeadBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerHealthHeadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
