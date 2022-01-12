// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerHealthReader is a Reader for the ServiceBrokerHealth structure.
type ServiceBrokerHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerHealthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerHealthOK creates a ServiceBrokerHealthOK with default headers values
func NewServiceBrokerHealthOK() *ServiceBrokerHealthOK {
	return &ServiceBrokerHealthOK{}
}

/* ServiceBrokerHealthOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerHealthOK struct {
	Payload *models.Health
}

func (o *ServiceBrokerHealthOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/health][%d] serviceBrokerHealthOK  %+v", 200, o.Payload)
}
func (o *ServiceBrokerHealthOK) GetPayload() *models.Health {
	return o.Payload
}

func (o *ServiceBrokerHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Health)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerHealthBadRequest creates a ServiceBrokerHealthBadRequest with default headers values
func NewServiceBrokerHealthBadRequest() *ServiceBrokerHealthBadRequest {
	return &ServiceBrokerHealthBadRequest{}
}

/* ServiceBrokerHealthBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerHealthBadRequest struct {
	Payload *models.Error
}

func (o *ServiceBrokerHealthBadRequest) Error() string {
	return fmt.Sprintf("[GET /broker/v1/health][%d] serviceBrokerHealthBadRequest  %+v", 400, o.Payload)
}
func (o *ServiceBrokerHealthBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerHealthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
