// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesImagesPostReader is a Reader for the PcloudCloudinstancesImagesPost structure.
type PcloudCloudinstancesImagesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesImagesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudCloudinstancesImagesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewPcloudCloudinstancesImagesPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudCloudinstancesImagesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudCloudinstancesImagesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPcloudCloudinstancesImagesPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudCloudinstancesImagesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudCloudinstancesImagesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudCloudinstancesImagesPostOK creates a PcloudCloudinstancesImagesPostOK with default headers values
func NewPcloudCloudinstancesImagesPostOK() *PcloudCloudinstancesImagesPostOK {
	return &PcloudCloudinstancesImagesPostOK{}
}

/*PcloudCloudinstancesImagesPostOK handles this case with default header values.

OK
*/
type PcloudCloudinstancesImagesPostOK struct {
	Payload *models.Image
}

func (o *PcloudCloudinstancesImagesPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images][%d] pcloudCloudinstancesImagesPostOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesImagesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesPostCreated creates a PcloudCloudinstancesImagesPostCreated with default headers values
func NewPcloudCloudinstancesImagesPostCreated() *PcloudCloudinstancesImagesPostCreated {
	return &PcloudCloudinstancesImagesPostCreated{}
}

/*PcloudCloudinstancesImagesPostCreated handles this case with default header values.

Created
*/
type PcloudCloudinstancesImagesPostCreated struct {
	Payload *models.Image
}

func (o *PcloudCloudinstancesImagesPostCreated) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images][%d] pcloudCloudinstancesImagesPostCreated  %+v", 201, o.Payload)
}

func (o *PcloudCloudinstancesImagesPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesPostBadRequest creates a PcloudCloudinstancesImagesPostBadRequest with default headers values
func NewPcloudCloudinstancesImagesPostBadRequest() *PcloudCloudinstancesImagesPostBadRequest {
	return &PcloudCloudinstancesImagesPostBadRequest{}
}

/*PcloudCloudinstancesImagesPostBadRequest handles this case with default header values.

Bad Request
*/
type PcloudCloudinstancesImagesPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images][%d] pcloudCloudinstancesImagesPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesImagesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesPostUnauthorized creates a PcloudCloudinstancesImagesPostUnauthorized with default headers values
func NewPcloudCloudinstancesImagesPostUnauthorized() *PcloudCloudinstancesImagesPostUnauthorized {
	return &PcloudCloudinstancesImagesPostUnauthorized{}
}

/*PcloudCloudinstancesImagesPostUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudCloudinstancesImagesPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images][%d] pcloudCloudinstancesImagesPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesImagesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesPostConflict creates a PcloudCloudinstancesImagesPostConflict with default headers values
func NewPcloudCloudinstancesImagesPostConflict() *PcloudCloudinstancesImagesPostConflict {
	return &PcloudCloudinstancesImagesPostConflict{}
}

/*PcloudCloudinstancesImagesPostConflict handles this case with default header values.

Conflict
*/
type PcloudCloudinstancesImagesPostConflict struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images][%d] pcloudCloudinstancesImagesPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudCloudinstancesImagesPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesPostUnprocessableEntity creates a PcloudCloudinstancesImagesPostUnprocessableEntity with default headers values
func NewPcloudCloudinstancesImagesPostUnprocessableEntity() *PcloudCloudinstancesImagesPostUnprocessableEntity {
	return &PcloudCloudinstancesImagesPostUnprocessableEntity{}
}

/*PcloudCloudinstancesImagesPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudCloudinstancesImagesPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images][%d] pcloudCloudinstancesImagesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudinstancesImagesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesPostInternalServerError creates a PcloudCloudinstancesImagesPostInternalServerError with default headers values
func NewPcloudCloudinstancesImagesPostInternalServerError() *PcloudCloudinstancesImagesPostInternalServerError {
	return &PcloudCloudinstancesImagesPostInternalServerError{}
}

/*PcloudCloudinstancesImagesPostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudCloudinstancesImagesPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images][%d] pcloudCloudinstancesImagesPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesImagesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
