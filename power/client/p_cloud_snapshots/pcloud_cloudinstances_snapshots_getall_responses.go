// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesSnapshotsGetallReader is a Reader for the PcloudCloudinstancesSnapshotsGetall structure.
type PcloudCloudinstancesSnapshotsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesSnapshotsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesSnapshotsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesSnapshotsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesSnapshotsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesSnapshotsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesSnapshotsGetallOK creates a PcloudCloudinstancesSnapshotsGetallOK with default headers values
func NewPcloudCloudinstancesSnapshotsGetallOK() *PcloudCloudinstancesSnapshotsGetallOK {
	return &PcloudCloudinstancesSnapshotsGetallOK{}
}

/* PcloudCloudinstancesSnapshotsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesSnapshotsGetallOK struct {
	Payload *models.Snapshots
}

func (o *PcloudCloudinstancesSnapshotsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudinstancesSnapshotsGetallOK) GetPayload() *models.Snapshots {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snapshots)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetallBadRequest creates a PcloudCloudinstancesSnapshotsGetallBadRequest with default headers values
func NewPcloudCloudinstancesSnapshotsGetallBadRequest() *PcloudCloudinstancesSnapshotsGetallBadRequest {
	return &PcloudCloudinstancesSnapshotsGetallBadRequest{}
}

/* PcloudCloudinstancesSnapshotsGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesSnapshotsGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetallUnauthorized creates a PcloudCloudinstancesSnapshotsGetallUnauthorized with default headers values
func NewPcloudCloudinstancesSnapshotsGetallUnauthorized() *PcloudCloudinstancesSnapshotsGetallUnauthorized {
	return &PcloudCloudinstancesSnapshotsGetallUnauthorized{}
}

/* PcloudCloudinstancesSnapshotsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesSnapshotsGetallUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetallInternalServerError creates a PcloudCloudinstancesSnapshotsGetallInternalServerError with default headers values
func NewPcloudCloudinstancesSnapshotsGetallInternalServerError() *PcloudCloudinstancesSnapshotsGetallInternalServerError {
	return &PcloudCloudinstancesSnapshotsGetallInternalServerError{}
}

/* PcloudCloudinstancesSnapshotsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesSnapshotsGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
