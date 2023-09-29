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

// PcloudVpnconnectionsPeersubnetsPutReader is a Reader for the PcloudVpnconnectionsPeersubnetsPut structure.
type PcloudVpnconnectionsPeersubnetsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsPeersubnetsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVpnconnectionsPeersubnetsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsPeersubnetsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsPeersubnetsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsPeersubnetsPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVpnconnectionsPeersubnetsPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsPeersubnetsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsPeersubnetsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets] pcloud.vpnconnections.peersubnets.put", response, response.Code())
	}
}

// NewPcloudVpnconnectionsPeersubnetsPutOK creates a PcloudVpnconnectionsPeersubnetsPutOK with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutOK() *PcloudVpnconnectionsPeersubnetsPutOK {
	return &PcloudVpnconnectionsPeersubnetsPutOK{}
}

/*
PcloudVpnconnectionsPeersubnetsPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVpnconnectionsPeersubnetsPutOK struct {
	Payload *models.PeerSubnets
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets put o k response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets put o k response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets put o k response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections peersubnets put o k response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets put o k response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud vpnconnections peersubnets put o k response
func (o *PcloudVpnconnectionsPeersubnetsPutOK) Code() int {
	return 200
}

func (o *PcloudVpnconnectionsPeersubnetsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutOK) GetPayload() *models.PeerSubnets {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerSubnets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutBadRequest creates a PcloudVpnconnectionsPeersubnetsPutBadRequest with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutBadRequest() *PcloudVpnconnectionsPeersubnetsPutBadRequest {
	return &PcloudVpnconnectionsPeersubnetsPutBadRequest{}
}

/*
PcloudVpnconnectionsPeersubnetsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsPeersubnetsPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets put bad request response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets put bad request response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets put bad request response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets put bad request response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets put bad request response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud vpnconnections peersubnets put bad request response
func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) Code() int {
	return 400
}

func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutUnauthorized creates a PcloudVpnconnectionsPeersubnetsPutUnauthorized with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutUnauthorized() *PcloudVpnconnectionsPeersubnetsPutUnauthorized {
	return &PcloudVpnconnectionsPeersubnetsPutUnauthorized{}
}

/*
PcloudVpnconnectionsPeersubnetsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsPeersubnetsPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets put unauthorized response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets put unauthorized response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets put unauthorized response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets put unauthorized response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets put unauthorized response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud vpnconnections peersubnets put unauthorized response
func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) Code() int {
	return 401
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutForbidden creates a PcloudVpnconnectionsPeersubnetsPutForbidden with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutForbidden() *PcloudVpnconnectionsPeersubnetsPutForbidden {
	return &PcloudVpnconnectionsPeersubnetsPutForbidden{}
}

/*
PcloudVpnconnectionsPeersubnetsPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsPeersubnetsPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets put forbidden response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets put forbidden response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets put forbidden response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets put forbidden response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets put forbidden response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud vpnconnections peersubnets put forbidden response
func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) Code() int {
	return 403
}

func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutNotFound creates a PcloudVpnconnectionsPeersubnetsPutNotFound with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutNotFound() *PcloudVpnconnectionsPeersubnetsPutNotFound {
	return &PcloudVpnconnectionsPeersubnetsPutNotFound{}
}

/*
PcloudVpnconnectionsPeersubnetsPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVpnconnectionsPeersubnetsPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets put not found response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets put not found response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets put not found response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets put not found response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets put not found response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud vpnconnections peersubnets put not found response
func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) Code() int {
	return 404
}

func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutUnprocessableEntity creates a PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutUnprocessableEntity() *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity {
	return &PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity{}
}

/*
PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets put unprocessable entity response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets put unprocessable entity response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets put unprocessable entity response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets put unprocessable entity response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets put unprocessable entity response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud vpnconnections peersubnets put unprocessable entity response
func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutInternalServerError creates a PcloudVpnconnectionsPeersubnetsPutInternalServerError with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutInternalServerError() *PcloudVpnconnectionsPeersubnetsPutInternalServerError {
	return &PcloudVpnconnectionsPeersubnetsPutInternalServerError{}
}

/*
PcloudVpnconnectionsPeersubnetsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsPeersubnetsPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets put internal server error response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets put internal server error response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets put internal server error response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections peersubnets put internal server error response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud vpnconnections peersubnets put internal server error response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud vpnconnections peersubnets put internal server error response
func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) Code() int {
	return 500
}

func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
