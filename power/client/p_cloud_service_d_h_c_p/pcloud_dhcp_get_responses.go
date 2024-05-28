// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_service_d_h_c_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudDhcpGetReader is a Reader for the PcloudDhcpGet structure.
type PcloudDhcpGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudDhcpGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudDhcpGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudDhcpGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudDhcpGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudDhcpGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudDhcpGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudDhcpGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}] pcloud.dhcp.get", response, response.Code())
	}
}

// NewPcloudDhcpGetOK creates a PcloudDhcpGetOK with default headers values
func NewPcloudDhcpGetOK() *PcloudDhcpGetOK {
	return &PcloudDhcpGetOK{}
}

/*
PcloudDhcpGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudDhcpGetOK struct {
	Payload *models.DHCPServerDetail
}

// IsSuccess returns true when this pcloud dhcp get o k response has a 2xx status code
func (o *PcloudDhcpGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud dhcp get o k response has a 3xx status code
func (o *PcloudDhcpGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp get o k response has a 4xx status code
func (o *PcloudDhcpGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud dhcp get o k response has a 5xx status code
func (o *PcloudDhcpGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp get o k response a status code equal to that given
func (o *PcloudDhcpGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud dhcp get o k response
func (o *PcloudDhcpGetOK) Code() int {
	return 200
}

func (o *PcloudDhcpGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetOK %s", 200, payload)
}

func (o *PcloudDhcpGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetOK %s", 200, payload)
}

func (o *PcloudDhcpGetOK) GetPayload() *models.DHCPServerDetail {
	return o.Payload
}

func (o *PcloudDhcpGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DHCPServerDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpGetBadRequest creates a PcloudDhcpGetBadRequest with default headers values
func NewPcloudDhcpGetBadRequest() *PcloudDhcpGetBadRequest {
	return &PcloudDhcpGetBadRequest{}
}

/*
PcloudDhcpGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudDhcpGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp get bad request response has a 2xx status code
func (o *PcloudDhcpGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp get bad request response has a 3xx status code
func (o *PcloudDhcpGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp get bad request response has a 4xx status code
func (o *PcloudDhcpGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp get bad request response has a 5xx status code
func (o *PcloudDhcpGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp get bad request response a status code equal to that given
func (o *PcloudDhcpGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud dhcp get bad request response
func (o *PcloudDhcpGetBadRequest) Code() int {
	return 400
}

func (o *PcloudDhcpGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetBadRequest %s", 400, payload)
}

func (o *PcloudDhcpGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetBadRequest %s", 400, payload)
}

func (o *PcloudDhcpGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpGetUnauthorized creates a PcloudDhcpGetUnauthorized with default headers values
func NewPcloudDhcpGetUnauthorized() *PcloudDhcpGetUnauthorized {
	return &PcloudDhcpGetUnauthorized{}
}

/*
PcloudDhcpGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudDhcpGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp get unauthorized response has a 2xx status code
func (o *PcloudDhcpGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp get unauthorized response has a 3xx status code
func (o *PcloudDhcpGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp get unauthorized response has a 4xx status code
func (o *PcloudDhcpGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp get unauthorized response has a 5xx status code
func (o *PcloudDhcpGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp get unauthorized response a status code equal to that given
func (o *PcloudDhcpGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud dhcp get unauthorized response
func (o *PcloudDhcpGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudDhcpGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetUnauthorized %s", 401, payload)
}

func (o *PcloudDhcpGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetUnauthorized %s", 401, payload)
}

func (o *PcloudDhcpGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpGetForbidden creates a PcloudDhcpGetForbidden with default headers values
func NewPcloudDhcpGetForbidden() *PcloudDhcpGetForbidden {
	return &PcloudDhcpGetForbidden{}
}

/*
PcloudDhcpGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudDhcpGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp get forbidden response has a 2xx status code
func (o *PcloudDhcpGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp get forbidden response has a 3xx status code
func (o *PcloudDhcpGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp get forbidden response has a 4xx status code
func (o *PcloudDhcpGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp get forbidden response has a 5xx status code
func (o *PcloudDhcpGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp get forbidden response a status code equal to that given
func (o *PcloudDhcpGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud dhcp get forbidden response
func (o *PcloudDhcpGetForbidden) Code() int {
	return 403
}

func (o *PcloudDhcpGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetForbidden %s", 403, payload)
}

func (o *PcloudDhcpGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetForbidden %s", 403, payload)
}

func (o *PcloudDhcpGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpGetNotFound creates a PcloudDhcpGetNotFound with default headers values
func NewPcloudDhcpGetNotFound() *PcloudDhcpGetNotFound {
	return &PcloudDhcpGetNotFound{}
}

/*
PcloudDhcpGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudDhcpGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp get not found response has a 2xx status code
func (o *PcloudDhcpGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp get not found response has a 3xx status code
func (o *PcloudDhcpGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp get not found response has a 4xx status code
func (o *PcloudDhcpGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp get not found response has a 5xx status code
func (o *PcloudDhcpGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp get not found response a status code equal to that given
func (o *PcloudDhcpGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud dhcp get not found response
func (o *PcloudDhcpGetNotFound) Code() int {
	return 404
}

func (o *PcloudDhcpGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetNotFound %s", 404, payload)
}

func (o *PcloudDhcpGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetNotFound %s", 404, payload)
}

func (o *PcloudDhcpGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpGetInternalServerError creates a PcloudDhcpGetInternalServerError with default headers values
func NewPcloudDhcpGetInternalServerError() *PcloudDhcpGetInternalServerError {
	return &PcloudDhcpGetInternalServerError{}
}

/*
PcloudDhcpGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudDhcpGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp get internal server error response has a 2xx status code
func (o *PcloudDhcpGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp get internal server error response has a 3xx status code
func (o *PcloudDhcpGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp get internal server error response has a 4xx status code
func (o *PcloudDhcpGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud dhcp get internal server error response has a 5xx status code
func (o *PcloudDhcpGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud dhcp get internal server error response a status code equal to that given
func (o *PcloudDhcpGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud dhcp get internal server error response
func (o *PcloudDhcpGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudDhcpGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetInternalServerError %s", 500, payload)
}

func (o *PcloudDhcpGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpGetInternalServerError %s", 500, payload)
}

func (o *PcloudDhcpGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
