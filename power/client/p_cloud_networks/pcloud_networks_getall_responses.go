// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudNetworksGetallReader is a Reader for the PcloudNetworksGetall structure.
type PcloudNetworksGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudNetworksGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudNetworksGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudNetworksGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudNetworksGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudNetworksGetallOK creates a PcloudNetworksGetallOK with default headers values
func NewPcloudNetworksGetallOK() *PcloudNetworksGetallOK {
	return &PcloudNetworksGetallOK{}
}

/*
PcloudNetworksGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudNetworksGetallOK struct {
	Payload *models.Networks
}

// IsSuccess returns true when this pcloud networks getall o k response has a 2xx status code
func (o *PcloudNetworksGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks getall o k response has a 3xx status code
func (o *PcloudNetworksGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks getall o k response has a 4xx status code
func (o *PcloudNetworksGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks getall o k response has a 5xx status code
func (o *PcloudNetworksGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks getall o k response a status code equal to that given
func (o *PcloudNetworksGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud networks getall o k response
func (o *PcloudNetworksGetallOK) Code() int {
	return 200
}

func (o *PcloudNetworksGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksGetallOK) GetPayload() *models.Networks {
	return o.Payload
}

func (o *PcloudNetworksGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Networks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetallBadRequest creates a PcloudNetworksGetallBadRequest with default headers values
func NewPcloudNetworksGetallBadRequest() *PcloudNetworksGetallBadRequest {
	return &PcloudNetworksGetallBadRequest{}
}

/*
PcloudNetworksGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudNetworksGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks getall bad request response has a 2xx status code
func (o *PcloudNetworksGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks getall bad request response has a 3xx status code
func (o *PcloudNetworksGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks getall bad request response has a 4xx status code
func (o *PcloudNetworksGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks getall bad request response has a 5xx status code
func (o *PcloudNetworksGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks getall bad request response a status code equal to that given
func (o *PcloudNetworksGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud networks getall bad request response
func (o *PcloudNetworksGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudNetworksGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetallUnauthorized creates a PcloudNetworksGetallUnauthorized with default headers values
func NewPcloudNetworksGetallUnauthorized() *PcloudNetworksGetallUnauthorized {
	return &PcloudNetworksGetallUnauthorized{}
}

/*
PcloudNetworksGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks getall unauthorized response has a 2xx status code
func (o *PcloudNetworksGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks getall unauthorized response has a 3xx status code
func (o *PcloudNetworksGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks getall unauthorized response has a 4xx status code
func (o *PcloudNetworksGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks getall unauthorized response has a 5xx status code
func (o *PcloudNetworksGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks getall unauthorized response a status code equal to that given
func (o *PcloudNetworksGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud networks getall unauthorized response
func (o *PcloudNetworksGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudNetworksGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetallForbidden creates a PcloudNetworksGetallForbidden with default headers values
func NewPcloudNetworksGetallForbidden() *PcloudNetworksGetallForbidden {
	return &PcloudNetworksGetallForbidden{}
}

/*
PcloudNetworksGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudNetworksGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks getall forbidden response has a 2xx status code
func (o *PcloudNetworksGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks getall forbidden response has a 3xx status code
func (o *PcloudNetworksGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks getall forbidden response has a 4xx status code
func (o *PcloudNetworksGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks getall forbidden response has a 5xx status code
func (o *PcloudNetworksGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks getall forbidden response a status code equal to that given
func (o *PcloudNetworksGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud networks getall forbidden response
func (o *PcloudNetworksGetallForbidden) Code() int {
	return 403
}

func (o *PcloudNetworksGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetallInternalServerError creates a PcloudNetworksGetallInternalServerError with default headers values
func NewPcloudNetworksGetallInternalServerError() *PcloudNetworksGetallInternalServerError {
	return &PcloudNetworksGetallInternalServerError{}
}

/*
PcloudNetworksGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks getall internal server error response has a 2xx status code
func (o *PcloudNetworksGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks getall internal server error response has a 3xx status code
func (o *PcloudNetworksGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks getall internal server error response has a 4xx status code
func (o *PcloudNetworksGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks getall internal server error response has a 5xx status code
func (o *PcloudNetworksGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks getall internal server error response a status code equal to that given
func (o *PcloudNetworksGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud networks getall internal server error response
func (o *PcloudNetworksGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudNetworksGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
