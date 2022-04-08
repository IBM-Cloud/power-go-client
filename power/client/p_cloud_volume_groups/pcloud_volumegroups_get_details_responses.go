// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVolumegroupsGetDetailsReader is a Reader for the PcloudVolumegroupsGetDetails structure.
type PcloudVolumegroupsGetDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumegroupsGetDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVolumegroupsGetDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumegroupsGetDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVolumegroupsGetDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumegroupsGetDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVolumegroupsGetDetailsOK creates a PcloudVolumegroupsGetDetailsOK with default headers values
func NewPcloudVolumegroupsGetDetailsOK() *PcloudVolumegroupsGetDetailsOK {
	return &PcloudVolumegroupsGetDetailsOK{}
}

/* PcloudVolumegroupsGetDetailsOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVolumegroupsGetDetailsOK struct {
	Payload *models.VolumeGroupCreateResponse
}

func (o *PcloudVolumegroupsGetDetailsOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsOK  %+v", 200, o.Payload)
}
func (o *PcloudVolumegroupsGetDetailsOK) GetPayload() *models.VolumeGroupCreateResponse {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetDetailsBadRequest creates a PcloudVolumegroupsGetDetailsBadRequest with default headers values
func NewPcloudVolumegroupsGetDetailsBadRequest() *PcloudVolumegroupsGetDetailsBadRequest {
	return &PcloudVolumegroupsGetDetailsBadRequest{}
}

/* PcloudVolumegroupsGetDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumegroupsGetDetailsBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVolumegroupsGetDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudVolumegroupsGetDetailsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetDetailsNotFound creates a PcloudVolumegroupsGetDetailsNotFound with default headers values
func NewPcloudVolumegroupsGetDetailsNotFound() *PcloudVolumegroupsGetDetailsNotFound {
	return &PcloudVolumegroupsGetDetailsNotFound{}
}

/* PcloudVolumegroupsGetDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVolumegroupsGetDetailsNotFound struct {
	Payload *models.Error
}

func (o *PcloudVolumegroupsGetDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsNotFound  %+v", 404, o.Payload)
}
func (o *PcloudVolumegroupsGetDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetDetailsInternalServerError creates a PcloudVolumegroupsGetDetailsInternalServerError with default headers values
func NewPcloudVolumegroupsGetDetailsInternalServerError() *PcloudVolumegroupsGetDetailsInternalServerError {
	return &PcloudVolumegroupsGetDetailsInternalServerError{}
}

/* PcloudVolumegroupsGetDetailsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumegroupsGetDetailsInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVolumegroupsGetDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudVolumegroupsGetDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
