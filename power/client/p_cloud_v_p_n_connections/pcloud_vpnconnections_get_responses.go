// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVpnconnectionsGetReader is a Reader for the PcloudVpnconnectionsGet structure.
type PcloudVpnconnectionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVpnconnectionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVpnconnectionsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVpnconnectionsGetOK creates a PcloudVpnconnectionsGetOK with default headers values
func NewPcloudVpnconnectionsGetOK() *PcloudVpnconnectionsGetOK {
	return &PcloudVpnconnectionsGetOK{}
}

/*
PcloudVpnconnectionsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVpnconnectionsGetOK struct {
	Payload *models.VPNConnection
}

// IsSuccess returns true when this pcloud vpnconnections get o k response has a 2xx status code
func (o *PcloudVpnconnectionsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud vpnconnections get o k response has a 3xx status code
func (o *PcloudVpnconnectionsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections get o k response has a 4xx status code
func (o *PcloudVpnconnectionsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections get o k response has a 5xx status code
func (o *PcloudVpnconnectionsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections get o k response a status code equal to that given
func (o *PcloudVpnconnectionsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud vpnconnections get o k response
func (o *PcloudVpnconnectionsGetOK) Code() int {
	return 200
}

func (o *PcloudVpnconnectionsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudVpnconnectionsGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudVpnconnectionsGetOK) GetPayload() *models.VPNConnection {
	return o.Payload
}

func (o *PcloudVpnconnectionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VPNConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsGetBadRequest creates a PcloudVpnconnectionsGetBadRequest with default headers values
func NewPcloudVpnconnectionsGetBadRequest() *PcloudVpnconnectionsGetBadRequest {
	return &PcloudVpnconnectionsGetBadRequest{}
}

/*
PcloudVpnconnectionsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections get bad request response has a 2xx status code
func (o *PcloudVpnconnectionsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections get bad request response has a 3xx status code
func (o *PcloudVpnconnectionsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections get bad request response has a 4xx status code
func (o *PcloudVpnconnectionsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections get bad request response has a 5xx status code
func (o *PcloudVpnconnectionsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections get bad request response a status code equal to that given
func (o *PcloudVpnconnectionsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud vpnconnections get bad request response
func (o *PcloudVpnconnectionsGetBadRequest) Code() int {
	return 400
}

func (o *PcloudVpnconnectionsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsGetUnauthorized creates a PcloudVpnconnectionsGetUnauthorized with default headers values
func NewPcloudVpnconnectionsGetUnauthorized() *PcloudVpnconnectionsGetUnauthorized {
	return &PcloudVpnconnectionsGetUnauthorized{}
}

/*
PcloudVpnconnectionsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections get unauthorized response has a 2xx status code
func (o *PcloudVpnconnectionsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections get unauthorized response has a 3xx status code
func (o *PcloudVpnconnectionsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections get unauthorized response has a 4xx status code
func (o *PcloudVpnconnectionsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections get unauthorized response has a 5xx status code
func (o *PcloudVpnconnectionsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections get unauthorized response a status code equal to that given
func (o *PcloudVpnconnectionsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud vpnconnections get unauthorized response
func (o *PcloudVpnconnectionsGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudVpnconnectionsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsGetForbidden creates a PcloudVpnconnectionsGetForbidden with default headers values
func NewPcloudVpnconnectionsGetForbidden() *PcloudVpnconnectionsGetForbidden {
	return &PcloudVpnconnectionsGetForbidden{}
}

/*
PcloudVpnconnectionsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections get forbidden response has a 2xx status code
func (o *PcloudVpnconnectionsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections get forbidden response has a 3xx status code
func (o *PcloudVpnconnectionsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections get forbidden response has a 4xx status code
func (o *PcloudVpnconnectionsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections get forbidden response has a 5xx status code
func (o *PcloudVpnconnectionsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections get forbidden response a status code equal to that given
func (o *PcloudVpnconnectionsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud vpnconnections get forbidden response
func (o *PcloudVpnconnectionsGetForbidden) Code() int {
	return 403
}

func (o *PcloudVpnconnectionsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsGetNotFound creates a PcloudVpnconnectionsGetNotFound with default headers values
func NewPcloudVpnconnectionsGetNotFound() *PcloudVpnconnectionsGetNotFound {
	return &PcloudVpnconnectionsGetNotFound{}
}

/*
PcloudVpnconnectionsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVpnconnectionsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections get not found response has a 2xx status code
func (o *PcloudVpnconnectionsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections get not found response has a 3xx status code
func (o *PcloudVpnconnectionsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections get not found response has a 4xx status code
func (o *PcloudVpnconnectionsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections get not found response has a 5xx status code
func (o *PcloudVpnconnectionsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections get not found response a status code equal to that given
func (o *PcloudVpnconnectionsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud vpnconnections get not found response
func (o *PcloudVpnconnectionsGetNotFound) Code() int {
	return 404
}

func (o *PcloudVpnconnectionsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsGetUnprocessableEntity creates a PcloudVpnconnectionsGetUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsGetUnprocessableEntity() *PcloudVpnconnectionsGetUnprocessableEntity {
	return &PcloudVpnconnectionsGetUnprocessableEntity{}
}

/*
PcloudVpnconnectionsGetUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsGetUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections get unprocessable entity response has a 2xx status code
func (o *PcloudVpnconnectionsGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections get unprocessable entity response has a 3xx status code
func (o *PcloudVpnconnectionsGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections get unprocessable entity response has a 4xx status code
func (o *PcloudVpnconnectionsGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections get unprocessable entity response has a 5xx status code
func (o *PcloudVpnconnectionsGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections get unprocessable entity response a status code equal to that given
func (o *PcloudVpnconnectionsGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud vpnconnections get unprocessable entity response
func (o *PcloudVpnconnectionsGetUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudVpnconnectionsGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsGetUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsGetInternalServerError creates a PcloudVpnconnectionsGetInternalServerError with default headers values
func NewPcloudVpnconnectionsGetInternalServerError() *PcloudVpnconnectionsGetInternalServerError {
	return &PcloudVpnconnectionsGetInternalServerError{}
}

/*
PcloudVpnconnectionsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections get internal server error response has a 2xx status code
func (o *PcloudVpnconnectionsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections get internal server error response has a 3xx status code
func (o *PcloudVpnconnectionsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections get internal server error response has a 4xx status code
func (o *PcloudVpnconnectionsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections get internal server error response has a 5xx status code
func (o *PcloudVpnconnectionsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud vpnconnections get internal server error response a status code equal to that given
func (o *PcloudVpnconnectionsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud vpnconnections get internal server error response
func (o *PcloudVpnconnectionsGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudVpnconnectionsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
