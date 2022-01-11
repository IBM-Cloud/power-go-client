// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsPostReader is a Reader for the PcloudCloudconnectionsPost structure.
type PcloudCloudconnectionsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPcloudCloudconnectionsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudCloudconnectionsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudCloudconnectionsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudconnectionsPostOK creates a PcloudCloudconnectionsPostOK with default headers values
func NewPcloudCloudconnectionsPostOK() *PcloudCloudconnectionsPostOK {
	return &PcloudCloudconnectionsPostOK{}
}

/* PcloudCloudconnectionsPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsPostOK struct {
	Payload *models.CloudConnection
}

func (o *PcloudCloudconnectionsPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudconnectionsPostOK) GetPayload() *models.CloudConnection {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostCreated creates a PcloudCloudconnectionsPostCreated with default headers values
func NewPcloudCloudconnectionsPostCreated() *PcloudCloudconnectionsPostCreated {
	return &PcloudCloudconnectionsPostCreated{}
}

/* PcloudCloudconnectionsPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudCloudconnectionsPostCreated struct {
	Payload *models.CloudConnection
}

func (o *PcloudCloudconnectionsPostCreated) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostCreated  %+v", 201, o.Payload)
}
func (o *PcloudCloudconnectionsPostCreated) GetPayload() *models.CloudConnection {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostAccepted creates a PcloudCloudconnectionsPostAccepted with default headers values
func NewPcloudCloudconnectionsPostAccepted() *PcloudCloudconnectionsPostAccepted {
	return &PcloudCloudconnectionsPostAccepted{}
}

/* PcloudCloudconnectionsPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsPostAccepted struct {
	Payload *models.CloudConnectionCreateResponse
}

func (o *PcloudCloudconnectionsPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudCloudconnectionsPostAccepted) GetPayload() *models.CloudConnectionCreateResponse {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnectionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostBadRequest creates a PcloudCloudconnectionsPostBadRequest with default headers values
func NewPcloudCloudconnectionsPostBadRequest() *PcloudCloudconnectionsPostBadRequest {
	return &PcloudCloudconnectionsPostBadRequest{}
}

/* PcloudCloudconnectionsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudconnectionsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostUnauthorized creates a PcloudCloudconnectionsPostUnauthorized with default headers values
func NewPcloudCloudconnectionsPostUnauthorized() *PcloudCloudconnectionsPostUnauthorized {
	return &PcloudCloudconnectionsPostUnauthorized{}
}

/* PcloudCloudconnectionsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudconnectionsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostConflict creates a PcloudCloudconnectionsPostConflict with default headers values
func NewPcloudCloudconnectionsPostConflict() *PcloudCloudconnectionsPostConflict {
	return &PcloudCloudconnectionsPostConflict{}
}

/* PcloudCloudconnectionsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudCloudconnectionsPostConflict struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudCloudconnectionsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostUnprocessableEntity creates a PcloudCloudconnectionsPostUnprocessableEntity with default headers values
func NewPcloudCloudconnectionsPostUnprocessableEntity() *PcloudCloudconnectionsPostUnprocessableEntity {
	return &PcloudCloudconnectionsPostUnprocessableEntity{}
}

/* PcloudCloudconnectionsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudconnectionsPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudCloudconnectionsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostInternalServerError creates a PcloudCloudconnectionsPostInternalServerError with default headers values
func NewPcloudCloudconnectionsPostInternalServerError() *PcloudCloudconnectionsPostInternalServerError {
	return &PcloudCloudconnectionsPostInternalServerError{}
}

/* PcloudCloudconnectionsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudconnectionsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
