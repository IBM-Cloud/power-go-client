// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudTenantsSshkeysPostReader is a Reader for the PcloudTenantsSshkeysPost structure.
type PcloudTenantsSshkeysPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsSshkeysPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTenantsSshkeysPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPcloudTenantsSshkeysPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTenantsSshkeysPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTenantsSshkeysPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudTenantsSshkeysPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudTenantsSshkeysPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTenantsSshkeysPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudTenantsSshkeysPostOK creates a PcloudTenantsSshkeysPostOK with default headers values
func NewPcloudTenantsSshkeysPostOK() *PcloudTenantsSshkeysPostOK {
	return &PcloudTenantsSshkeysPostOK{}
}

/* PcloudTenantsSshkeysPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTenantsSshkeysPostOK struct {
	Payload *models.SSHKey
}

func (o *PcloudTenantsSshkeysPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysPostOK  %+v", 200, o.Payload)
}
func (o *PcloudTenantsSshkeysPostOK) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPostCreated creates a PcloudTenantsSshkeysPostCreated with default headers values
func NewPcloudTenantsSshkeysPostCreated() *PcloudTenantsSshkeysPostCreated {
	return &PcloudTenantsSshkeysPostCreated{}
}

/* PcloudTenantsSshkeysPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudTenantsSshkeysPostCreated struct {
	Payload *models.SSHKey
}

func (o *PcloudTenantsSshkeysPostCreated) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysPostCreated  %+v", 201, o.Payload)
}
func (o *PcloudTenantsSshkeysPostCreated) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPostBadRequest creates a PcloudTenantsSshkeysPostBadRequest with default headers values
func NewPcloudTenantsSshkeysPostBadRequest() *PcloudTenantsSshkeysPostBadRequest {
	return &PcloudTenantsSshkeysPostBadRequest{}
}

/* PcloudTenantsSshkeysPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTenantsSshkeysPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudTenantsSshkeysPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPostUnauthorized creates a PcloudTenantsSshkeysPostUnauthorized with default headers values
func NewPcloudTenantsSshkeysPostUnauthorized() *PcloudTenantsSshkeysPostUnauthorized {
	return &PcloudTenantsSshkeysPostUnauthorized{}
}

/* PcloudTenantsSshkeysPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTenantsSshkeysPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysPostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudTenantsSshkeysPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPostConflict creates a PcloudTenantsSshkeysPostConflict with default headers values
func NewPcloudTenantsSshkeysPostConflict() *PcloudTenantsSshkeysPostConflict {
	return &PcloudTenantsSshkeysPostConflict{}
}

/* PcloudTenantsSshkeysPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudTenantsSshkeysPostConflict struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudTenantsSshkeysPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPostUnprocessableEntity creates a PcloudTenantsSshkeysPostUnprocessableEntity with default headers values
func NewPcloudTenantsSshkeysPostUnprocessableEntity() *PcloudTenantsSshkeysPostUnprocessableEntity {
	return &PcloudTenantsSshkeysPostUnprocessableEntity{}
}

/* PcloudTenantsSshkeysPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudTenantsSshkeysPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudTenantsSshkeysPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPostInternalServerError creates a PcloudTenantsSshkeysPostInternalServerError with default headers values
func NewPcloudTenantsSshkeysPostInternalServerError() *PcloudTenantsSshkeysPostInternalServerError {
	return &PcloudTenantsSshkeysPostInternalServerError{}
}

/* PcloudTenantsSshkeysPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTenantsSshkeysPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudTenantsSshkeysPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
