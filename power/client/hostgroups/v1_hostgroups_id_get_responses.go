// Code generated by go-swagger; DO NOT EDIT.

package hostgroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1HostgroupsIDGetReader is a Reader for the V1HostgroupsIDGet structure.
type V1HostgroupsIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1HostgroupsIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1HostgroupsIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1HostgroupsIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1HostgroupsIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1HostgroupsIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1HostgroupsIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1HostgroupsIDGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewV1HostgroupsIDGetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/hostgroups/{hostgroup_id}] v1.hostgroups.id.get", response, response.Code())
	}
}

// NewV1HostgroupsIDGetOK creates a V1HostgroupsIDGetOK with default headers values
func NewV1HostgroupsIDGetOK() *V1HostgroupsIDGetOK {
	return &V1HostgroupsIDGetOK{}
}

/*
V1HostgroupsIDGetOK describes a response with status code 200, with default header values.

OK
*/
type V1HostgroupsIDGetOK struct {
	Payload *models.HostgroupWithSharingInfo
}

// IsSuccess returns true when this v1 hostgroups Id get o k response has a 2xx status code
func (o *V1HostgroupsIDGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 hostgroups Id get o k response has a 3xx status code
func (o *V1HostgroupsIDGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hostgroups Id get o k response has a 4xx status code
func (o *V1HostgroupsIDGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hostgroups Id get o k response has a 5xx status code
func (o *V1HostgroupsIDGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hostgroups Id get o k response a status code equal to that given
func (o *V1HostgroupsIDGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 hostgroups Id get o k response
func (o *V1HostgroupsIDGetOK) Code() int {
	return 200
}

func (o *V1HostgroupsIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetOK  %+v", 200, o.Payload)
}

func (o *V1HostgroupsIDGetOK) String() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetOK  %+v", 200, o.Payload)
}

func (o *V1HostgroupsIDGetOK) GetPayload() *models.HostgroupWithSharingInfo {
	return o.Payload
}

func (o *V1HostgroupsIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostgroupWithSharingInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostgroupsIDGetBadRequest creates a V1HostgroupsIDGetBadRequest with default headers values
func NewV1HostgroupsIDGetBadRequest() *V1HostgroupsIDGetBadRequest {
	return &V1HostgroupsIDGetBadRequest{}
}

/*
V1HostgroupsIDGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1HostgroupsIDGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hostgroups Id get bad request response has a 2xx status code
func (o *V1HostgroupsIDGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hostgroups Id get bad request response has a 3xx status code
func (o *V1HostgroupsIDGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hostgroups Id get bad request response has a 4xx status code
func (o *V1HostgroupsIDGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hostgroups Id get bad request response has a 5xx status code
func (o *V1HostgroupsIDGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hostgroups Id get bad request response a status code equal to that given
func (o *V1HostgroupsIDGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 hostgroups Id get bad request response
func (o *V1HostgroupsIDGetBadRequest) Code() int {
	return 400
}

func (o *V1HostgroupsIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *V1HostgroupsIDGetBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *V1HostgroupsIDGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostgroupsIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostgroupsIDGetUnauthorized creates a V1HostgroupsIDGetUnauthorized with default headers values
func NewV1HostgroupsIDGetUnauthorized() *V1HostgroupsIDGetUnauthorized {
	return &V1HostgroupsIDGetUnauthorized{}
}

/*
V1HostgroupsIDGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1HostgroupsIDGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hostgroups Id get unauthorized response has a 2xx status code
func (o *V1HostgroupsIDGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hostgroups Id get unauthorized response has a 3xx status code
func (o *V1HostgroupsIDGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hostgroups Id get unauthorized response has a 4xx status code
func (o *V1HostgroupsIDGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hostgroups Id get unauthorized response has a 5xx status code
func (o *V1HostgroupsIDGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hostgroups Id get unauthorized response a status code equal to that given
func (o *V1HostgroupsIDGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 hostgroups Id get unauthorized response
func (o *V1HostgroupsIDGetUnauthorized) Code() int {
	return 401
}

func (o *V1HostgroupsIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V1HostgroupsIDGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V1HostgroupsIDGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostgroupsIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostgroupsIDGetForbidden creates a V1HostgroupsIDGetForbidden with default headers values
func NewV1HostgroupsIDGetForbidden() *V1HostgroupsIDGetForbidden {
	return &V1HostgroupsIDGetForbidden{}
}

/*
V1HostgroupsIDGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1HostgroupsIDGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hostgroups Id get forbidden response has a 2xx status code
func (o *V1HostgroupsIDGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hostgroups Id get forbidden response has a 3xx status code
func (o *V1HostgroupsIDGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hostgroups Id get forbidden response has a 4xx status code
func (o *V1HostgroupsIDGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hostgroups Id get forbidden response has a 5xx status code
func (o *V1HostgroupsIDGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hostgroups Id get forbidden response a status code equal to that given
func (o *V1HostgroupsIDGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 hostgroups Id get forbidden response
func (o *V1HostgroupsIDGetForbidden) Code() int {
	return 403
}

func (o *V1HostgroupsIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetForbidden  %+v", 403, o.Payload)
}

func (o *V1HostgroupsIDGetForbidden) String() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetForbidden  %+v", 403, o.Payload)
}

func (o *V1HostgroupsIDGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostgroupsIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostgroupsIDGetNotFound creates a V1HostgroupsIDGetNotFound with default headers values
func NewV1HostgroupsIDGetNotFound() *V1HostgroupsIDGetNotFound {
	return &V1HostgroupsIDGetNotFound{}
}

/*
V1HostgroupsIDGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1HostgroupsIDGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hostgroups Id get not found response has a 2xx status code
func (o *V1HostgroupsIDGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hostgroups Id get not found response has a 3xx status code
func (o *V1HostgroupsIDGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hostgroups Id get not found response has a 4xx status code
func (o *V1HostgroupsIDGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hostgroups Id get not found response has a 5xx status code
func (o *V1HostgroupsIDGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hostgroups Id get not found response a status code equal to that given
func (o *V1HostgroupsIDGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 hostgroups Id get not found response
func (o *V1HostgroupsIDGetNotFound) Code() int {
	return 404
}

func (o *V1HostgroupsIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetNotFound  %+v", 404, o.Payload)
}

func (o *V1HostgroupsIDGetNotFound) String() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetNotFound  %+v", 404, o.Payload)
}

func (o *V1HostgroupsIDGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostgroupsIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostgroupsIDGetInternalServerError creates a V1HostgroupsIDGetInternalServerError with default headers values
func NewV1HostgroupsIDGetInternalServerError() *V1HostgroupsIDGetInternalServerError {
	return &V1HostgroupsIDGetInternalServerError{}
}

/*
V1HostgroupsIDGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1HostgroupsIDGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hostgroups Id get internal server error response has a 2xx status code
func (o *V1HostgroupsIDGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hostgroups Id get internal server error response has a 3xx status code
func (o *V1HostgroupsIDGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hostgroups Id get internal server error response has a 4xx status code
func (o *V1HostgroupsIDGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hostgroups Id get internal server error response has a 5xx status code
func (o *V1HostgroupsIDGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 hostgroups Id get internal server error response a status code equal to that given
func (o *V1HostgroupsIDGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 hostgroups Id get internal server error response
func (o *V1HostgroupsIDGetInternalServerError) Code() int {
	return 500
}

func (o *V1HostgroupsIDGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetInternalServerError  %+v", 500, o.Payload)
}

func (o *V1HostgroupsIDGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetInternalServerError  %+v", 500, o.Payload)
}

func (o *V1HostgroupsIDGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostgroupsIDGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostgroupsIDGetGatewayTimeout creates a V1HostgroupsIDGetGatewayTimeout with default headers values
func NewV1HostgroupsIDGetGatewayTimeout() *V1HostgroupsIDGetGatewayTimeout {
	return &V1HostgroupsIDGetGatewayTimeout{}
}

/*
V1HostgroupsIDGetGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. Request is still processing and taking longer than expected.
*/
type V1HostgroupsIDGetGatewayTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hostgroups Id get gateway timeout response has a 2xx status code
func (o *V1HostgroupsIDGetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hostgroups Id get gateway timeout response has a 3xx status code
func (o *V1HostgroupsIDGetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hostgroups Id get gateway timeout response has a 4xx status code
func (o *V1HostgroupsIDGetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hostgroups Id get gateway timeout response has a 5xx status code
func (o *V1HostgroupsIDGetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 hostgroups Id get gateway timeout response a status code equal to that given
func (o *V1HostgroupsIDGetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the v1 hostgroups Id get gateway timeout response
func (o *V1HostgroupsIDGetGatewayTimeout) Code() int {
	return 504
}

func (o *V1HostgroupsIDGetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *V1HostgroupsIDGetGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/hostgroups/{hostgroup_id}][%d] v1HostgroupsIdGetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *V1HostgroupsIDGetGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostgroupsIDGetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}