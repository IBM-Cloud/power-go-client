// Code generated by go-swagger; DO NOT EDIT.

package swagger_spec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerSwaggerspecReader is a Reader for the ServiceBrokerSwaggerspec structure.
type ServiceBrokerSwaggerspecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerSwaggerspecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerSwaggerspecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerSwaggerspecOK creates a ServiceBrokerSwaggerspecOK with default headers values
func NewServiceBrokerSwaggerspecOK() *ServiceBrokerSwaggerspecOK {
	return &ServiceBrokerSwaggerspecOK{}
}

/*
ServiceBrokerSwaggerspecOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerSwaggerspecOK struct {
	Payload models.Object
}

// IsSuccess returns true when this service broker swaggerspec o k response has a 2xx status code
func (o *ServiceBrokerSwaggerspecOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker swaggerspec o k response has a 3xx status code
func (o *ServiceBrokerSwaggerspecOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker swaggerspec o k response has a 4xx status code
func (o *ServiceBrokerSwaggerspecOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker swaggerspec o k response has a 5xx status code
func (o *ServiceBrokerSwaggerspecOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker swaggerspec o k response a status code equal to that given
func (o *ServiceBrokerSwaggerspecOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker swaggerspec o k response
func (o *ServiceBrokerSwaggerspecOK) Code() int {
	return 200
}

func (o *ServiceBrokerSwaggerspecOK) Error() string {
	return fmt.Sprintf("[GET /v1/swagger.json][%d] serviceBrokerSwaggerspecOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerSwaggerspecOK) String() string {
	return fmt.Sprintf("[GET /v1/swagger.json][%d] serviceBrokerSwaggerspecOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerSwaggerspecOK) GetPayload() models.Object {
	return o.Payload
}

func (o *ServiceBrokerSwaggerspecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
