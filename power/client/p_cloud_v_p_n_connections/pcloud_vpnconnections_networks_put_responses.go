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

// PcloudVpnconnectionsNetworksPutReader is a Reader for the PcloudVpnconnectionsNetworksPut structure.
type PcloudVpnconnectionsNetworksPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsNetworksPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudVpnconnectionsNetworksPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsNetworksPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsNetworksPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsNetworksPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVpnconnectionsNetworksPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsNetworksPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsNetworksPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVpnconnectionsNetworksPutAccepted creates a PcloudVpnconnectionsNetworksPutAccepted with default headers values
func NewPcloudVpnconnectionsNetworksPutAccepted() *PcloudVpnconnectionsNetworksPutAccepted {
	return &PcloudVpnconnectionsNetworksPutAccepted{}
}

/* PcloudVpnconnectionsNetworksPutAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudVpnconnectionsNetworksPutAccepted struct {
	Payload *models.JobReference
}

func (o *PcloudVpnconnectionsNetworksPutAccepted) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksPutAccepted  %+v", 202, o.Payload)
}
func (o *PcloudVpnconnectionsNetworksPutAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksPutBadRequest creates a PcloudVpnconnectionsNetworksPutBadRequest with default headers values
func NewPcloudVpnconnectionsNetworksPutBadRequest() *PcloudVpnconnectionsNetworksPutBadRequest {
	return &PcloudVpnconnectionsNetworksPutBadRequest{}
}

/* PcloudVpnconnectionsNetworksPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsNetworksPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsNetworksPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksPutBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudVpnconnectionsNetworksPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksPutUnauthorized creates a PcloudVpnconnectionsNetworksPutUnauthorized with default headers values
func NewPcloudVpnconnectionsNetworksPutUnauthorized() *PcloudVpnconnectionsNetworksPutUnauthorized {
	return &PcloudVpnconnectionsNetworksPutUnauthorized{}
}

/* PcloudVpnconnectionsNetworksPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsNetworksPutUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsNetworksPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksPutUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudVpnconnectionsNetworksPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksPutForbidden creates a PcloudVpnconnectionsNetworksPutForbidden with default headers values
func NewPcloudVpnconnectionsNetworksPutForbidden() *PcloudVpnconnectionsNetworksPutForbidden {
	return &PcloudVpnconnectionsNetworksPutForbidden{}
}

/* PcloudVpnconnectionsNetworksPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsNetworksPutForbidden struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsNetworksPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksPutForbidden  %+v", 403, o.Payload)
}
func (o *PcloudVpnconnectionsNetworksPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksPutNotFound creates a PcloudVpnconnectionsNetworksPutNotFound with default headers values
func NewPcloudVpnconnectionsNetworksPutNotFound() *PcloudVpnconnectionsNetworksPutNotFound {
	return &PcloudVpnconnectionsNetworksPutNotFound{}
}

/* PcloudVpnconnectionsNetworksPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVpnconnectionsNetworksPutNotFound struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsNetworksPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksPutNotFound  %+v", 404, o.Payload)
}
func (o *PcloudVpnconnectionsNetworksPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksPutUnprocessableEntity creates a PcloudVpnconnectionsNetworksPutUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsNetworksPutUnprocessableEntity() *PcloudVpnconnectionsNetworksPutUnprocessableEntity {
	return &PcloudVpnconnectionsNetworksPutUnprocessableEntity{}
}

/* PcloudVpnconnectionsNetworksPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsNetworksPutUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsNetworksPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksPutUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudVpnconnectionsNetworksPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksPutInternalServerError creates a PcloudVpnconnectionsNetworksPutInternalServerError with default headers values
func NewPcloudVpnconnectionsNetworksPutInternalServerError() *PcloudVpnconnectionsNetworksPutInternalServerError {
	return &PcloudVpnconnectionsNetworksPutInternalServerError{}
}

/* PcloudVpnconnectionsNetworksPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsNetworksPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsNetworksPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksPutInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudVpnconnectionsNetworksPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
