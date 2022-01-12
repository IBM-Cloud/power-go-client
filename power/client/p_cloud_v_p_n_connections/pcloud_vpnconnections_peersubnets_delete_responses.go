// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVpnconnectionsPeersubnetsDeleteReader is a Reader for the PcloudVpnconnectionsPeersubnetsDelete structure.
type PcloudVpnconnectionsPeersubnetsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsPeersubnetsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVpnconnectionsPeersubnetsDeleteOK creates a PcloudVpnconnectionsPeersubnetsDeleteOK with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteOK() *PcloudVpnconnectionsPeersubnetsDeleteOK {
	return &PcloudVpnconnectionsPeersubnetsDeleteOK{}
}

/* PcloudVpnconnectionsPeersubnetsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVpnconnectionsPeersubnetsDeleteOK struct {
	Payload *models.PeerSubnets
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteOK  %+v", 200, o.Payload)
}
func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) GetPayload() *models.PeerSubnets {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerSubnets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteBadRequest creates a PcloudVpnconnectionsPeersubnetsDeleteBadRequest with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteBadRequest() *PcloudVpnconnectionsPeersubnetsDeleteBadRequest {
	return &PcloudVpnconnectionsPeersubnetsDeleteBadRequest{}
}

/* PcloudVpnconnectionsPeersubnetsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsPeersubnetsDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteUnauthorized creates a PcloudVpnconnectionsPeersubnetsDeleteUnauthorized with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteUnauthorized() *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized {
	return &PcloudVpnconnectionsPeersubnetsDeleteUnauthorized{}
}

/* PcloudVpnconnectionsPeersubnetsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsPeersubnetsDeleteUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteForbidden creates a PcloudVpnconnectionsPeersubnetsDeleteForbidden with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteForbidden() *PcloudVpnconnectionsPeersubnetsDeleteForbidden {
	return &PcloudVpnconnectionsPeersubnetsDeleteForbidden{}
}

/* PcloudVpnconnectionsPeersubnetsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsPeersubnetsDeleteForbidden struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteForbidden  %+v", 403, o.Payload)
}
func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity creates a PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity() *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity {
	return &PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity{}
}

/* PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteInternalServerError creates a PcloudVpnconnectionsPeersubnetsDeleteInternalServerError with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteInternalServerError() *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError {
	return &PcloudVpnconnectionsPeersubnetsDeleteInternalServerError{}
}

/* PcloudVpnconnectionsPeersubnetsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsPeersubnetsDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
