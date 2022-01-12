// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesImagesExportPostReader is a Reader for the PcloudCloudinstancesImagesExportPost structure.
type PcloudCloudinstancesImagesExportPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesImagesExportPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudCloudinstancesImagesExportPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesImagesExportPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesImagesExportPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesImagesExportPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudCloudinstancesImagesExportPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesImagesExportPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesImagesExportPostAccepted creates a PcloudCloudinstancesImagesExportPostAccepted with default headers values
func NewPcloudCloudinstancesImagesExportPostAccepted() *PcloudCloudinstancesImagesExportPostAccepted {
	return &PcloudCloudinstancesImagesExportPostAccepted{}
}

/* PcloudCloudinstancesImagesExportPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudinstancesImagesExportPostAccepted struct {
	Payload models.Object
}

func (o *PcloudCloudinstancesImagesExportPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudCloudinstancesImagesExportPostAccepted) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesExportPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesExportPostBadRequest creates a PcloudCloudinstancesImagesExportPostBadRequest with default headers values
func NewPcloudCloudinstancesImagesExportPostBadRequest() *PcloudCloudinstancesImagesExportPostBadRequest {
	return &PcloudCloudinstancesImagesExportPostBadRequest{}
}

/* PcloudCloudinstancesImagesExportPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesImagesExportPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesExportPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudinstancesImagesExportPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesExportPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesExportPostUnauthorized creates a PcloudCloudinstancesImagesExportPostUnauthorized with default headers values
func NewPcloudCloudinstancesImagesExportPostUnauthorized() *PcloudCloudinstancesImagesExportPostUnauthorized {
	return &PcloudCloudinstancesImagesExportPostUnauthorized{}
}

/* PcloudCloudinstancesImagesExportPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesImagesExportPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesExportPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudinstancesImagesExportPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesExportPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesExportPostNotFound creates a PcloudCloudinstancesImagesExportPostNotFound with default headers values
func NewPcloudCloudinstancesImagesExportPostNotFound() *PcloudCloudinstancesImagesExportPostNotFound {
	return &PcloudCloudinstancesImagesExportPostNotFound{}
}

/* PcloudCloudinstancesImagesExportPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesImagesExportPostNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesExportPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostNotFound  %+v", 404, o.Payload)
}
func (o *PcloudCloudinstancesImagesExportPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesExportPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesExportPostUnprocessableEntity creates a PcloudCloudinstancesImagesExportPostUnprocessableEntity with default headers values
func NewPcloudCloudinstancesImagesExportPostUnprocessableEntity() *PcloudCloudinstancesImagesExportPostUnprocessableEntity {
	return &PcloudCloudinstancesImagesExportPostUnprocessableEntity{}
}

/* PcloudCloudinstancesImagesExportPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudinstancesImagesExportPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesExportPostInternalServerError creates a PcloudCloudinstancesImagesExportPostInternalServerError with default headers values
func NewPcloudCloudinstancesImagesExportPostInternalServerError() *PcloudCloudinstancesImagesExportPostInternalServerError {
	return &PcloudCloudinstancesImagesExportPostInternalServerError{}
}

/* PcloudCloudinstancesImagesExportPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesImagesExportPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesExportPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudinstancesImagesExportPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesExportPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
