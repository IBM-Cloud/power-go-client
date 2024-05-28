// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

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

// PcloudVpnconnectionsDeleteReader is a Reader for the PcloudVpnconnectionsDelete structure.
type PcloudVpnconnectionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudVpnconnectionsDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVpnconnectionsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}] pcloud.vpnconnections.delete", response, response.Code())
	}
}

// NewPcloudVpnconnectionsDeleteAccepted creates a PcloudVpnconnectionsDeleteAccepted with default headers values
func NewPcloudVpnconnectionsDeleteAccepted() *PcloudVpnconnectionsDeleteAccepted {
	return &PcloudVpnconnectionsDeleteAccepted{}
}

/*
PcloudVpnconnectionsDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudVpnconnectionsDeleteAccepted struct {
	Payload *models.JobReference
}

// IsSuccess returns true when this pcloud vpnconnections delete accepted response has a 2xx status code
func (o *PcloudVpnconnectionsDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud vpnconnections delete accepted response has a 3xx status code
func (o *PcloudVpnconnectionsDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections delete accepted response has a 4xx status code
func (o *PcloudVpnconnectionsDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections delete accepted response has a 5xx status code
func (o *PcloudVpnconnectionsDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections delete accepted response a status code equal to that given
func (o *PcloudVpnconnectionsDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud vpnconnections delete accepted response
func (o *PcloudVpnconnectionsDeleteAccepted) Code() int {
	return 202
}

func (o *PcloudVpnconnectionsDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteAccepted %s", 202, payload)
}

func (o *PcloudVpnconnectionsDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteAccepted %s", 202, payload)
}

func (o *PcloudVpnconnectionsDeleteAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudVpnconnectionsDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteBadRequest creates a PcloudVpnconnectionsDeleteBadRequest with default headers values
func NewPcloudVpnconnectionsDeleteBadRequest() *PcloudVpnconnectionsDeleteBadRequest {
	return &PcloudVpnconnectionsDeleteBadRequest{}
}

/*
PcloudVpnconnectionsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections delete bad request response has a 2xx status code
func (o *PcloudVpnconnectionsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections delete bad request response has a 3xx status code
func (o *PcloudVpnconnectionsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections delete bad request response has a 4xx status code
func (o *PcloudVpnconnectionsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections delete bad request response has a 5xx status code
func (o *PcloudVpnconnectionsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections delete bad request response a status code equal to that given
func (o *PcloudVpnconnectionsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud vpnconnections delete bad request response
func (o *PcloudVpnconnectionsDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudVpnconnectionsDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteBadRequest %s", 400, payload)
}

func (o *PcloudVpnconnectionsDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteBadRequest %s", 400, payload)
}

func (o *PcloudVpnconnectionsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteUnauthorized creates a PcloudVpnconnectionsDeleteUnauthorized with default headers values
func NewPcloudVpnconnectionsDeleteUnauthorized() *PcloudVpnconnectionsDeleteUnauthorized {
	return &PcloudVpnconnectionsDeleteUnauthorized{}
}

/*
PcloudVpnconnectionsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections delete unauthorized response has a 2xx status code
func (o *PcloudVpnconnectionsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections delete unauthorized response has a 3xx status code
func (o *PcloudVpnconnectionsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections delete unauthorized response has a 4xx status code
func (o *PcloudVpnconnectionsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections delete unauthorized response has a 5xx status code
func (o *PcloudVpnconnectionsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections delete unauthorized response a status code equal to that given
func (o *PcloudVpnconnectionsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud vpnconnections delete unauthorized response
func (o *PcloudVpnconnectionsDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudVpnconnectionsDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudVpnconnectionsDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudVpnconnectionsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteForbidden creates a PcloudVpnconnectionsDeleteForbidden with default headers values
func NewPcloudVpnconnectionsDeleteForbidden() *PcloudVpnconnectionsDeleteForbidden {
	return &PcloudVpnconnectionsDeleteForbidden{}
}

/*
PcloudVpnconnectionsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections delete forbidden response has a 2xx status code
func (o *PcloudVpnconnectionsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections delete forbidden response has a 3xx status code
func (o *PcloudVpnconnectionsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections delete forbidden response has a 4xx status code
func (o *PcloudVpnconnectionsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections delete forbidden response has a 5xx status code
func (o *PcloudVpnconnectionsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections delete forbidden response a status code equal to that given
func (o *PcloudVpnconnectionsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud vpnconnections delete forbidden response
func (o *PcloudVpnconnectionsDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudVpnconnectionsDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteForbidden %s", 403, payload)
}

func (o *PcloudVpnconnectionsDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteForbidden %s", 403, payload)
}

func (o *PcloudVpnconnectionsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteNotFound creates a PcloudVpnconnectionsDeleteNotFound with default headers values
func NewPcloudVpnconnectionsDeleteNotFound() *PcloudVpnconnectionsDeleteNotFound {
	return &PcloudVpnconnectionsDeleteNotFound{}
}

/*
PcloudVpnconnectionsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVpnconnectionsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections delete not found response has a 2xx status code
func (o *PcloudVpnconnectionsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections delete not found response has a 3xx status code
func (o *PcloudVpnconnectionsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections delete not found response has a 4xx status code
func (o *PcloudVpnconnectionsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections delete not found response has a 5xx status code
func (o *PcloudVpnconnectionsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections delete not found response a status code equal to that given
func (o *PcloudVpnconnectionsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud vpnconnections delete not found response
func (o *PcloudVpnconnectionsDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudVpnconnectionsDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteNotFound %s", 404, payload)
}

func (o *PcloudVpnconnectionsDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteNotFound %s", 404, payload)
}

func (o *PcloudVpnconnectionsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteInternalServerError creates a PcloudVpnconnectionsDeleteInternalServerError with default headers values
func NewPcloudVpnconnectionsDeleteInternalServerError() *PcloudVpnconnectionsDeleteInternalServerError {
	return &PcloudVpnconnectionsDeleteInternalServerError{}
}

/*
PcloudVpnconnectionsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections delete internal server error response has a 2xx status code
func (o *PcloudVpnconnectionsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections delete internal server error response has a 3xx status code
func (o *PcloudVpnconnectionsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections delete internal server error response has a 4xx status code
func (o *PcloudVpnconnectionsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections delete internal server error response has a 5xx status code
func (o *PcloudVpnconnectionsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud vpnconnections delete internal server error response a status code equal to that given
func (o *PcloudVpnconnectionsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud vpnconnections delete internal server error response
func (o *PcloudVpnconnectionsDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudVpnconnectionsDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudVpnconnectionsDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudVpnconnectionsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
