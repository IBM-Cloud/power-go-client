// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

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
	case 403:
		result := NewPcloudCloudinstancesImagesExportPostForbidden()
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
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export] pcloud.cloudinstances.images.export.post", response, response.Code())
	}
}

// NewPcloudCloudinstancesImagesExportPostAccepted creates a PcloudCloudinstancesImagesExportPostAccepted with default headers values
func NewPcloudCloudinstancesImagesExportPostAccepted() *PcloudCloudinstancesImagesExportPostAccepted {
	return &PcloudCloudinstancesImagesExportPostAccepted{}
}

/*
PcloudCloudinstancesImagesExportPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudinstancesImagesExportPostAccepted struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudinstances images export post accepted response has a 2xx status code
func (o *PcloudCloudinstancesImagesExportPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances images export post accepted response has a 3xx status code
func (o *PcloudCloudinstancesImagesExportPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images export post accepted response has a 4xx status code
func (o *PcloudCloudinstancesImagesExportPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances images export post accepted response has a 5xx status code
func (o *PcloudCloudinstancesImagesExportPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images export post accepted response a status code equal to that given
func (o *PcloudCloudinstancesImagesExportPostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud cloudinstances images export post accepted response
func (o *PcloudCloudinstancesImagesExportPostAccepted) Code() int {
	return 202
}

func (o *PcloudCloudinstancesImagesExportPostAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostAccepted %s", 202, payload)
}

func (o *PcloudCloudinstancesImagesExportPostAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostAccepted %s", 202, payload)
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

/*
PcloudCloudinstancesImagesExportPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesImagesExportPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images export post bad request response has a 2xx status code
func (o *PcloudCloudinstancesImagesExportPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images export post bad request response has a 3xx status code
func (o *PcloudCloudinstancesImagesExportPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images export post bad request response has a 4xx status code
func (o *PcloudCloudinstancesImagesExportPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images export post bad request response has a 5xx status code
func (o *PcloudCloudinstancesImagesExportPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images export post bad request response a status code equal to that given
func (o *PcloudCloudinstancesImagesExportPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances images export post bad request response
func (o *PcloudCloudinstancesImagesExportPostBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesImagesExportPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostBadRequest %s", 400, payload)
}

func (o *PcloudCloudinstancesImagesExportPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostBadRequest %s", 400, payload)
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

/*
PcloudCloudinstancesImagesExportPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesImagesExportPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images export post unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesImagesExportPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images export post unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesImagesExportPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images export post unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesImagesExportPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images export post unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesImagesExportPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images export post unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesImagesExportPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances images export post unauthorized response
func (o *PcloudCloudinstancesImagesExportPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesImagesExportPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostUnauthorized %s", 401, payload)
}

func (o *PcloudCloudinstancesImagesExportPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostUnauthorized %s", 401, payload)
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

// NewPcloudCloudinstancesImagesExportPostForbidden creates a PcloudCloudinstancesImagesExportPostForbidden with default headers values
func NewPcloudCloudinstancesImagesExportPostForbidden() *PcloudCloudinstancesImagesExportPostForbidden {
	return &PcloudCloudinstancesImagesExportPostForbidden{}
}

/*
PcloudCloudinstancesImagesExportPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesImagesExportPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images export post forbidden response has a 2xx status code
func (o *PcloudCloudinstancesImagesExportPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images export post forbidden response has a 3xx status code
func (o *PcloudCloudinstancesImagesExportPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images export post forbidden response has a 4xx status code
func (o *PcloudCloudinstancesImagesExportPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images export post forbidden response has a 5xx status code
func (o *PcloudCloudinstancesImagesExportPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images export post forbidden response a status code equal to that given
func (o *PcloudCloudinstancesImagesExportPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances images export post forbidden response
func (o *PcloudCloudinstancesImagesExportPostForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesImagesExportPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostForbidden %s", 403, payload)
}

func (o *PcloudCloudinstancesImagesExportPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostForbidden %s", 403, payload)
}

func (o *PcloudCloudinstancesImagesExportPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesExportPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
PcloudCloudinstancesImagesExportPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesImagesExportPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images export post not found response has a 2xx status code
func (o *PcloudCloudinstancesImagesExportPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images export post not found response has a 3xx status code
func (o *PcloudCloudinstancesImagesExportPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images export post not found response has a 4xx status code
func (o *PcloudCloudinstancesImagesExportPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images export post not found response has a 5xx status code
func (o *PcloudCloudinstancesImagesExportPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images export post not found response a status code equal to that given
func (o *PcloudCloudinstancesImagesExportPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances images export post not found response
func (o *PcloudCloudinstancesImagesExportPostNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesImagesExportPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostNotFound %s", 404, payload)
}

func (o *PcloudCloudinstancesImagesExportPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostNotFound %s", 404, payload)
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

/*
PcloudCloudinstancesImagesExportPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudinstancesImagesExportPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images export post unprocessable entity response has a 2xx status code
func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images export post unprocessable entity response has a 3xx status code
func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images export post unprocessable entity response has a 4xx status code
func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images export post unprocessable entity response has a 5xx status code
func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images export post unprocessable entity response a status code equal to that given
func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud cloudinstances images export post unprocessable entity response
func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudCloudinstancesImagesExportPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostUnprocessableEntity %s", 422, payload)
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

/*
PcloudCloudinstancesImagesExportPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesImagesExportPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images export post internal server error response has a 2xx status code
func (o *PcloudCloudinstancesImagesExportPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images export post internal server error response has a 3xx status code
func (o *PcloudCloudinstancesImagesExportPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images export post internal server error response has a 4xx status code
func (o *PcloudCloudinstancesImagesExportPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances images export post internal server error response has a 5xx status code
func (o *PcloudCloudinstancesImagesExportPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances images export post internal server error response a status code equal to that given
func (o *PcloudCloudinstancesImagesExportPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances images export post internal server error response
func (o *PcloudCloudinstancesImagesExportPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesImagesExportPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostInternalServerError %s", 500, payload)
}

func (o *PcloudCloudinstancesImagesExportPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudCloudinstancesImagesExportPostInternalServerError %s", 500, payload)
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
