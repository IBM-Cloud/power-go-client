// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudIpsecpoliciesDeleteReader is a Reader for the PcloudIpsecpoliciesDelete structure.
type PcloudIpsecpoliciesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudIpsecpoliciesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudIpsecpoliciesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudIpsecpoliciesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudIpsecpoliciesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudIpsecpoliciesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudIpsecpoliciesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudIpsecpoliciesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudIpsecpoliciesDeleteOK creates a PcloudIpsecpoliciesDeleteOK with default headers values
func NewPcloudIpsecpoliciesDeleteOK() *PcloudIpsecpoliciesDeleteOK {
	return &PcloudIpsecpoliciesDeleteOK{}
}

/* PcloudIpsecpoliciesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudIpsecpoliciesDeleteOK struct {
	Payload models.Object
}

func (o *PcloudIpsecpoliciesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesDeleteOK  %+v", 200, o.Payload)
}
func (o *PcloudIpsecpoliciesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudIpsecpoliciesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesDeleteBadRequest creates a PcloudIpsecpoliciesDeleteBadRequest with default headers values
func NewPcloudIpsecpoliciesDeleteBadRequest() *PcloudIpsecpoliciesDeleteBadRequest {
	return &PcloudIpsecpoliciesDeleteBadRequest{}
}

/* PcloudIpsecpoliciesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudIpsecpoliciesDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudIpsecpoliciesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesDeleteUnauthorized creates a PcloudIpsecpoliciesDeleteUnauthorized with default headers values
func NewPcloudIpsecpoliciesDeleteUnauthorized() *PcloudIpsecpoliciesDeleteUnauthorized {
	return &PcloudIpsecpoliciesDeleteUnauthorized{}
}

/* PcloudIpsecpoliciesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudIpsecpoliciesDeleteUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudIpsecpoliciesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesDeleteForbidden creates a PcloudIpsecpoliciesDeleteForbidden with default headers values
func NewPcloudIpsecpoliciesDeleteForbidden() *PcloudIpsecpoliciesDeleteForbidden {
	return &PcloudIpsecpoliciesDeleteForbidden{}
}

/* PcloudIpsecpoliciesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudIpsecpoliciesDeleteForbidden struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesDeleteForbidden  %+v", 403, o.Payload)
}
func (o *PcloudIpsecpoliciesDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesDeleteNotFound creates a PcloudIpsecpoliciesDeleteNotFound with default headers values
func NewPcloudIpsecpoliciesDeleteNotFound() *PcloudIpsecpoliciesDeleteNotFound {
	return &PcloudIpsecpoliciesDeleteNotFound{}
}

/* PcloudIpsecpoliciesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudIpsecpoliciesDeleteNotFound struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesDeleteNotFound  %+v", 404, o.Payload)
}
func (o *PcloudIpsecpoliciesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesDeleteInternalServerError creates a PcloudIpsecpoliciesDeleteInternalServerError with default headers values
func NewPcloudIpsecpoliciesDeleteInternalServerError() *PcloudIpsecpoliciesDeleteInternalServerError {
	return &PcloudIpsecpoliciesDeleteInternalServerError{}
}

/* PcloudIpsecpoliciesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudIpsecpoliciesDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudIpsecpoliciesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
