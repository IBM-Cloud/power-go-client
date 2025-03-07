// Code generated by go-swagger; DO NOT EDIT.

package routes

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

// V1RoutesReportGetReader is a Reader for the V1RoutesReportGet structure.
type V1RoutesReportGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1RoutesReportGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1RoutesReportGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1RoutesReportGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1RoutesReportGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1RoutesReportGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1RoutesReportGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1RoutesReportGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/routes/report] v1.routes.report.get", response, response.Code())
	}
}

// NewV1RoutesReportGetOK creates a V1RoutesReportGetOK with default headers values
func NewV1RoutesReportGetOK() *V1RoutesReportGetOK {
	return &V1RoutesReportGetOK{}
}

/*
V1RoutesReportGetOK describes a response with status code 200, with default header values.

OK
*/
type V1RoutesReportGetOK struct {
	Payload *models.RouteReport
}

// IsSuccess returns true when this v1 routes report get o k response has a 2xx status code
func (o *V1RoutesReportGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 routes report get o k response has a 3xx status code
func (o *V1RoutesReportGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 routes report get o k response has a 4xx status code
func (o *V1RoutesReportGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 routes report get o k response has a 5xx status code
func (o *V1RoutesReportGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 routes report get o k response a status code equal to that given
func (o *V1RoutesReportGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 routes report get o k response
func (o *V1RoutesReportGetOK) Code() int {
	return 200
}

func (o *V1RoutesReportGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetOK %s", 200, payload)
}

func (o *V1RoutesReportGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetOK %s", 200, payload)
}

func (o *V1RoutesReportGetOK) GetPayload() *models.RouteReport {
	return o.Payload
}

func (o *V1RoutesReportGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1RoutesReportGetBadRequest creates a V1RoutesReportGetBadRequest with default headers values
func NewV1RoutesReportGetBadRequest() *V1RoutesReportGetBadRequest {
	return &V1RoutesReportGetBadRequest{}
}

/*
V1RoutesReportGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1RoutesReportGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 routes report get bad request response has a 2xx status code
func (o *V1RoutesReportGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 routes report get bad request response has a 3xx status code
func (o *V1RoutesReportGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 routes report get bad request response has a 4xx status code
func (o *V1RoutesReportGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 routes report get bad request response has a 5xx status code
func (o *V1RoutesReportGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 routes report get bad request response a status code equal to that given
func (o *V1RoutesReportGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 routes report get bad request response
func (o *V1RoutesReportGetBadRequest) Code() int {
	return 400
}

func (o *V1RoutesReportGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetBadRequest %s", 400, payload)
}

func (o *V1RoutesReportGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetBadRequest %s", 400, payload)
}

func (o *V1RoutesReportGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1RoutesReportGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1RoutesReportGetUnauthorized creates a V1RoutesReportGetUnauthorized with default headers values
func NewV1RoutesReportGetUnauthorized() *V1RoutesReportGetUnauthorized {
	return &V1RoutesReportGetUnauthorized{}
}

/*
V1RoutesReportGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1RoutesReportGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 routes report get unauthorized response has a 2xx status code
func (o *V1RoutesReportGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 routes report get unauthorized response has a 3xx status code
func (o *V1RoutesReportGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 routes report get unauthorized response has a 4xx status code
func (o *V1RoutesReportGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 routes report get unauthorized response has a 5xx status code
func (o *V1RoutesReportGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 routes report get unauthorized response a status code equal to that given
func (o *V1RoutesReportGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 routes report get unauthorized response
func (o *V1RoutesReportGetUnauthorized) Code() int {
	return 401
}

func (o *V1RoutesReportGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetUnauthorized %s", 401, payload)
}

func (o *V1RoutesReportGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetUnauthorized %s", 401, payload)
}

func (o *V1RoutesReportGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1RoutesReportGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1RoutesReportGetForbidden creates a V1RoutesReportGetForbidden with default headers values
func NewV1RoutesReportGetForbidden() *V1RoutesReportGetForbidden {
	return &V1RoutesReportGetForbidden{}
}

/*
V1RoutesReportGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1RoutesReportGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 routes report get forbidden response has a 2xx status code
func (o *V1RoutesReportGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 routes report get forbidden response has a 3xx status code
func (o *V1RoutesReportGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 routes report get forbidden response has a 4xx status code
func (o *V1RoutesReportGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 routes report get forbidden response has a 5xx status code
func (o *V1RoutesReportGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 routes report get forbidden response a status code equal to that given
func (o *V1RoutesReportGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 routes report get forbidden response
func (o *V1RoutesReportGetForbidden) Code() int {
	return 403
}

func (o *V1RoutesReportGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetForbidden %s", 403, payload)
}

func (o *V1RoutesReportGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetForbidden %s", 403, payload)
}

func (o *V1RoutesReportGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1RoutesReportGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1RoutesReportGetNotFound creates a V1RoutesReportGetNotFound with default headers values
func NewV1RoutesReportGetNotFound() *V1RoutesReportGetNotFound {
	return &V1RoutesReportGetNotFound{}
}

/*
V1RoutesReportGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1RoutesReportGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 routes report get not found response has a 2xx status code
func (o *V1RoutesReportGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 routes report get not found response has a 3xx status code
func (o *V1RoutesReportGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 routes report get not found response has a 4xx status code
func (o *V1RoutesReportGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 routes report get not found response has a 5xx status code
func (o *V1RoutesReportGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 routes report get not found response a status code equal to that given
func (o *V1RoutesReportGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 routes report get not found response
func (o *V1RoutesReportGetNotFound) Code() int {
	return 404
}

func (o *V1RoutesReportGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetNotFound %s", 404, payload)
}

func (o *V1RoutesReportGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetNotFound %s", 404, payload)
}

func (o *V1RoutesReportGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1RoutesReportGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1RoutesReportGetInternalServerError creates a V1RoutesReportGetInternalServerError with default headers values
func NewV1RoutesReportGetInternalServerError() *V1RoutesReportGetInternalServerError {
	return &V1RoutesReportGetInternalServerError{}
}

/*
V1RoutesReportGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1RoutesReportGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 routes report get internal server error response has a 2xx status code
func (o *V1RoutesReportGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 routes report get internal server error response has a 3xx status code
func (o *V1RoutesReportGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 routes report get internal server error response has a 4xx status code
func (o *V1RoutesReportGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 routes report get internal server error response has a 5xx status code
func (o *V1RoutesReportGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 routes report get internal server error response a status code equal to that given
func (o *V1RoutesReportGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 routes report get internal server error response
func (o *V1RoutesReportGetInternalServerError) Code() int {
	return 500
}

func (o *V1RoutesReportGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetInternalServerError %s", 500, payload)
}

func (o *V1RoutesReportGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routes/report][%d] v1RoutesReportGetInternalServerError %s", 500, payload)
}

func (o *V1RoutesReportGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1RoutesReportGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
