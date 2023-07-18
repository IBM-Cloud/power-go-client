// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesVolumesRemoteCopyRelationshipGetReader is a Reader for the PcloudCloudinstancesVolumesRemoteCopyRelationshipGet structure.
type PcloudCloudinstancesVolumesRemoteCopyRelationshipGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy] pcloud.cloudinstances.volumes.remoteCopyRelationship.get", response, response.Code())
	}
}

// NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK creates a PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK with default headers values
func NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK() *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK {
	return &PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK{}
}

/*
PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK struct {
	Payload *models.VolumeRemoteCopyRelationship
}

// IsSuccess returns true when this pcloud cloudinstances volumes remote copy relationship get o k response has a 2xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances volumes remote copy relationship get o k response has a 3xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes remote copy relationship get o k response has a 4xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances volumes remote copy relationship get o k response has a 5xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes remote copy relationship get o k response a status code equal to that given
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances volumes remote copy relationship get o k response
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) GetPayload() *models.VolumeRemoteCopyRelationship {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeRemoteCopyRelationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest creates a PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest with default headers values
func NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest() *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest {
	return &PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest{}
}

/*
PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes remote copy relationship get bad request response has a 2xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes remote copy relationship get bad request response has a 3xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes remote copy relationship get bad request response has a 4xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes remote copy relationship get bad request response has a 5xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes remote copy relationship get bad request response a status code equal to that given
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances volumes remote copy relationship get bad request response
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized creates a PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized with default headers values
func NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized() *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized {
	return &PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized{}
}

/*
PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes remote copy relationship get unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes remote copy relationship get unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes remote copy relationship get unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes remote copy relationship get unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes remote copy relationship get unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances volumes remote copy relationship get unauthorized response
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden creates a PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden with default headers values
func NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden() *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden {
	return &PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden{}
}

/*
PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes remote copy relationship get forbidden response has a 2xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes remote copy relationship get forbidden response has a 3xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes remote copy relationship get forbidden response has a 4xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes remote copy relationship get forbidden response has a 5xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes remote copy relationship get forbidden response a status code equal to that given
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances volumes remote copy relationship get forbidden response
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound creates a PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound with default headers values
func NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound() *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound {
	return &PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound{}
}

/*
PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes remote copy relationship get not found response has a 2xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes remote copy relationship get not found response has a 3xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes remote copy relationship get not found response has a 4xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes remote copy relationship get not found response has a 5xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes remote copy relationship get not found response a status code equal to that given
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances volumes remote copy relationship get not found response
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests creates a PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests with default headers values
func NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests() *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests {
	return &PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests{}
}

/*
PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes remote copy relationship get too many requests response has a 2xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes remote copy relationship get too many requests response has a 3xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes remote copy relationship get too many requests response has a 4xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes remote copy relationship get too many requests response has a 5xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes remote copy relationship get too many requests response a status code equal to that given
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the pcloud cloudinstances volumes remote copy relationship get too many requests response
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) Code() int {
	return 429
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError creates a PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError with default headers values
func NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError() *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError {
	return &PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError{}
}

/*
PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes remote copy relationship get internal server error response has a 2xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes remote copy relationship get internal server error response has a 3xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes remote copy relationship get internal server error response has a 4xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances volumes remote copy relationship get internal server error response has a 5xx status code
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances volumes remote copy relationship get internal server error response a status code equal to that given
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances volumes remote copy relationship get internal server error response
func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/remote-copy][%d] pcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
