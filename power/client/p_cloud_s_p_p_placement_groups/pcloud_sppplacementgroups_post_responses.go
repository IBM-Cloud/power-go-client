// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_p_p_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudSppplacementgroupsPostReader is a Reader for the PcloudSppplacementgroupsPost structure.
type PcloudSppplacementgroupsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSppplacementgroupsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSppplacementgroupsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSppplacementgroupsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSppplacementgroupsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudSppplacementgroupsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudSppplacementgroupsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudSppplacementgroupsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSppplacementgroupsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups] pcloud.sppplacementgroups.post", response, response.Code())
	}
}

// NewPcloudSppplacementgroupsPostOK creates a PcloudSppplacementgroupsPostOK with default headers values
func NewPcloudSppplacementgroupsPostOK() *PcloudSppplacementgroupsPostOK {
	return &PcloudSppplacementgroupsPostOK{}
}

/*
PcloudSppplacementgroupsPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSppplacementgroupsPostOK struct {
	Payload *models.SPPPlacementGroup
}

// IsSuccess returns true when this pcloud sppplacementgroups post o k response has a 2xx status code
func (o *PcloudSppplacementgroupsPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sppplacementgroups post o k response has a 3xx status code
func (o *PcloudSppplacementgroupsPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups post o k response has a 4xx status code
func (o *PcloudSppplacementgroupsPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sppplacementgroups post o k response has a 5xx status code
func (o *PcloudSppplacementgroupsPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups post o k response a status code equal to that given
func (o *PcloudSppplacementgroupsPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud sppplacementgroups post o k response
func (o *PcloudSppplacementgroupsPostOK) Code() int {
	return 200
}

func (o *PcloudSppplacementgroupsPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostOK  %+v", 200, o.Payload)
}

func (o *PcloudSppplacementgroupsPostOK) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostOK  %+v", 200, o.Payload)
}

func (o *PcloudSppplacementgroupsPostOK) GetPayload() *models.SPPPlacementGroup {
	return o.Payload
}

func (o *PcloudSppplacementgroupsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SPPPlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsPostBadRequest creates a PcloudSppplacementgroupsPostBadRequest with default headers values
func NewPcloudSppplacementgroupsPostBadRequest() *PcloudSppplacementgroupsPostBadRequest {
	return &PcloudSppplacementgroupsPostBadRequest{}
}

/*
PcloudSppplacementgroupsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSppplacementgroupsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups post bad request response has a 2xx status code
func (o *PcloudSppplacementgroupsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups post bad request response has a 3xx status code
func (o *PcloudSppplacementgroupsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups post bad request response has a 4xx status code
func (o *PcloudSppplacementgroupsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups post bad request response has a 5xx status code
func (o *PcloudSppplacementgroupsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups post bad request response a status code equal to that given
func (o *PcloudSppplacementgroupsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud sppplacementgroups post bad request response
func (o *PcloudSppplacementgroupsPostBadRequest) Code() int {
	return 400
}

func (o *PcloudSppplacementgroupsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSppplacementgroupsPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSppplacementgroupsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsPostUnauthorized creates a PcloudSppplacementgroupsPostUnauthorized with default headers values
func NewPcloudSppplacementgroupsPostUnauthorized() *PcloudSppplacementgroupsPostUnauthorized {
	return &PcloudSppplacementgroupsPostUnauthorized{}
}

/*
PcloudSppplacementgroupsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSppplacementgroupsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups post unauthorized response has a 2xx status code
func (o *PcloudSppplacementgroupsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups post unauthorized response has a 3xx status code
func (o *PcloudSppplacementgroupsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups post unauthorized response has a 4xx status code
func (o *PcloudSppplacementgroupsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups post unauthorized response has a 5xx status code
func (o *PcloudSppplacementgroupsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups post unauthorized response a status code equal to that given
func (o *PcloudSppplacementgroupsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud sppplacementgroups post unauthorized response
func (o *PcloudSppplacementgroupsPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudSppplacementgroupsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSppplacementgroupsPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSppplacementgroupsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsPostForbidden creates a PcloudSppplacementgroupsPostForbidden with default headers values
func NewPcloudSppplacementgroupsPostForbidden() *PcloudSppplacementgroupsPostForbidden {
	return &PcloudSppplacementgroupsPostForbidden{}
}

/*
PcloudSppplacementgroupsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudSppplacementgroupsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups post forbidden response has a 2xx status code
func (o *PcloudSppplacementgroupsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups post forbidden response has a 3xx status code
func (o *PcloudSppplacementgroupsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups post forbidden response has a 4xx status code
func (o *PcloudSppplacementgroupsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups post forbidden response has a 5xx status code
func (o *PcloudSppplacementgroupsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups post forbidden response a status code equal to that given
func (o *PcloudSppplacementgroupsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud sppplacementgroups post forbidden response
func (o *PcloudSppplacementgroupsPostForbidden) Code() int {
	return 403
}

func (o *PcloudSppplacementgroupsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudSppplacementgroupsPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudSppplacementgroupsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsPostConflict creates a PcloudSppplacementgroupsPostConflict with default headers values
func NewPcloudSppplacementgroupsPostConflict() *PcloudSppplacementgroupsPostConflict {
	return &PcloudSppplacementgroupsPostConflict{}
}

/*
PcloudSppplacementgroupsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudSppplacementgroupsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups post conflict response has a 2xx status code
func (o *PcloudSppplacementgroupsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups post conflict response has a 3xx status code
func (o *PcloudSppplacementgroupsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups post conflict response has a 4xx status code
func (o *PcloudSppplacementgroupsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups post conflict response has a 5xx status code
func (o *PcloudSppplacementgroupsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups post conflict response a status code equal to that given
func (o *PcloudSppplacementgroupsPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud sppplacementgroups post conflict response
func (o *PcloudSppplacementgroupsPostConflict) Code() int {
	return 409
}

func (o *PcloudSppplacementgroupsPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudSppplacementgroupsPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudSppplacementgroupsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsPostUnprocessableEntity creates a PcloudSppplacementgroupsPostUnprocessableEntity with default headers values
func NewPcloudSppplacementgroupsPostUnprocessableEntity() *PcloudSppplacementgroupsPostUnprocessableEntity {
	return &PcloudSppplacementgroupsPostUnprocessableEntity{}
}

/*
PcloudSppplacementgroupsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudSppplacementgroupsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups post unprocessable entity response has a 2xx status code
func (o *PcloudSppplacementgroupsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups post unprocessable entity response has a 3xx status code
func (o *PcloudSppplacementgroupsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups post unprocessable entity response has a 4xx status code
func (o *PcloudSppplacementgroupsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups post unprocessable entity response has a 5xx status code
func (o *PcloudSppplacementgroupsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups post unprocessable entity response a status code equal to that given
func (o *PcloudSppplacementgroupsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud sppplacementgroups post unprocessable entity response
func (o *PcloudSppplacementgroupsPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudSppplacementgroupsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudSppplacementgroupsPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudSppplacementgroupsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsPostInternalServerError creates a PcloudSppplacementgroupsPostInternalServerError with default headers values
func NewPcloudSppplacementgroupsPostInternalServerError() *PcloudSppplacementgroupsPostInternalServerError {
	return &PcloudSppplacementgroupsPostInternalServerError{}
}

/*
PcloudSppplacementgroupsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSppplacementgroupsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups post internal server error response has a 2xx status code
func (o *PcloudSppplacementgroupsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups post internal server error response has a 3xx status code
func (o *PcloudSppplacementgroupsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups post internal server error response has a 4xx status code
func (o *PcloudSppplacementgroupsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sppplacementgroups post internal server error response has a 5xx status code
func (o *PcloudSppplacementgroupsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud sppplacementgroups post internal server error response a status code equal to that given
func (o *PcloudSppplacementgroupsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud sppplacementgroups post internal server error response
func (o *PcloudSppplacementgroupsPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudSppplacementgroupsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSppplacementgroupsPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups][%d] pcloudSppplacementgroupsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSppplacementgroupsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
