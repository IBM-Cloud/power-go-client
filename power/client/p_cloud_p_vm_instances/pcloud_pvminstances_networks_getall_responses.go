// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesNetworksGetallReader is a Reader for the PcloudPvminstancesNetworksGetall structure.
type PcloudPvminstancesNetworksGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesNetworksGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesNetworksGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesNetworksGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesNetworksGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesNetworksGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesNetworksGetallOK creates a PcloudPvminstancesNetworksGetallOK with default headers values
func NewPcloudPvminstancesNetworksGetallOK() *PcloudPvminstancesNetworksGetallOK {
	return &PcloudPvminstancesNetworksGetallOK{}
}

/*
PcloudPvminstancesNetworksGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesNetworksGetallOK struct {
	Payload *models.PVMInstanceNetworks
}

// IsSuccess returns true when this pcloud pvminstances networks getall o k response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances networks getall o k response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks getall o k response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances networks getall o k response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks getall o k response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances networks getall o k response
func (o *PcloudPvminstancesNetworksGetallOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesNetworksGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallOK) GetPayload() *models.PVMInstanceNetworks {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstanceNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetallBadRequest creates a PcloudPvminstancesNetworksGetallBadRequest with default headers values
func NewPcloudPvminstancesNetworksGetallBadRequest() *PcloudPvminstancesNetworksGetallBadRequest {
	return &PcloudPvminstancesNetworksGetallBadRequest{}
}

/*
PcloudPvminstancesNetworksGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesNetworksGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks getall bad request response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks getall bad request response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks getall bad request response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks getall bad request response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks getall bad request response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances networks getall bad request response
func (o *PcloudPvminstancesNetworksGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesNetworksGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetallUnauthorized creates a PcloudPvminstancesNetworksGetallUnauthorized with default headers values
func NewPcloudPvminstancesNetworksGetallUnauthorized() *PcloudPvminstancesNetworksGetallUnauthorized {
	return &PcloudPvminstancesNetworksGetallUnauthorized{}
}

/*
PcloudPvminstancesNetworksGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesNetworksGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks getall unauthorized response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks getall unauthorized response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks getall unauthorized response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks getall unauthorized response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks getall unauthorized response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances networks getall unauthorized response
func (o *PcloudPvminstancesNetworksGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesNetworksGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetallInternalServerError creates a PcloudPvminstancesNetworksGetallInternalServerError with default headers values
func NewPcloudPvminstancesNetworksGetallInternalServerError() *PcloudPvminstancesNetworksGetallInternalServerError {
	return &PcloudPvminstancesNetworksGetallInternalServerError{}
}

/*
PcloudPvminstancesNetworksGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesNetworksGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks getall internal server error response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks getall internal server error response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks getall internal server error response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances networks getall internal server error response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances networks getall internal server error response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances networks getall internal server error response
func (o *PcloudPvminstancesNetworksGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesNetworksGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
