// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

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

// V1SshkeysPutReader is a Reader for the V1SshkeysPut structure.
type V1SshkeysPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SshkeysPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SshkeysPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1SshkeysPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1SshkeysPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1SshkeysPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1SshkeysPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV1SshkeysPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewV1SshkeysPutGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1SshkeysPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/ssh-keys/{sshkey_id}] v1.sshkeys.put", response, response.Code())
	}
}

// NewV1SshkeysPutOK creates a V1SshkeysPutOK with default headers values
func NewV1SshkeysPutOK() *V1SshkeysPutOK {
	return &V1SshkeysPutOK{}
}

/*
V1SshkeysPutOK describes a response with status code 200, with default header values.

OK
*/
type V1SshkeysPutOK struct {
	Payload *models.WorkspaceSSHKey
}

// IsSuccess returns true when this v1 sshkeys put o k response has a 2xx status code
func (o *V1SshkeysPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 sshkeys put o k response has a 3xx status code
func (o *V1SshkeysPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys put o k response has a 4xx status code
func (o *V1SshkeysPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 sshkeys put o k response has a 5xx status code
func (o *V1SshkeysPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys put o k response a status code equal to that given
func (o *V1SshkeysPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 sshkeys put o k response
func (o *V1SshkeysPutOK) Code() int {
	return 200
}

func (o *V1SshkeysPutOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutOK %s", 200, payload)
}

func (o *V1SshkeysPutOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutOK %s", 200, payload)
}

func (o *V1SshkeysPutOK) GetPayload() *models.WorkspaceSSHKey {
	return o.Payload
}

func (o *V1SshkeysPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceSSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPutBadRequest creates a V1SshkeysPutBadRequest with default headers values
func NewV1SshkeysPutBadRequest() *V1SshkeysPutBadRequest {
	return &V1SshkeysPutBadRequest{}
}

/*
V1SshkeysPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1SshkeysPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys put bad request response has a 2xx status code
func (o *V1SshkeysPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys put bad request response has a 3xx status code
func (o *V1SshkeysPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys put bad request response has a 4xx status code
func (o *V1SshkeysPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys put bad request response has a 5xx status code
func (o *V1SshkeysPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys put bad request response a status code equal to that given
func (o *V1SshkeysPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 sshkeys put bad request response
func (o *V1SshkeysPutBadRequest) Code() int {
	return 400
}

func (o *V1SshkeysPutBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutBadRequest %s", 400, payload)
}

func (o *V1SshkeysPutBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutBadRequest %s", 400, payload)
}

func (o *V1SshkeysPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPutUnauthorized creates a V1SshkeysPutUnauthorized with default headers values
func NewV1SshkeysPutUnauthorized() *V1SshkeysPutUnauthorized {
	return &V1SshkeysPutUnauthorized{}
}

/*
V1SshkeysPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1SshkeysPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys put unauthorized response has a 2xx status code
func (o *V1SshkeysPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys put unauthorized response has a 3xx status code
func (o *V1SshkeysPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys put unauthorized response has a 4xx status code
func (o *V1SshkeysPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys put unauthorized response has a 5xx status code
func (o *V1SshkeysPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys put unauthorized response a status code equal to that given
func (o *V1SshkeysPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 sshkeys put unauthorized response
func (o *V1SshkeysPutUnauthorized) Code() int {
	return 401
}

func (o *V1SshkeysPutUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutUnauthorized %s", 401, payload)
}

func (o *V1SshkeysPutUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutUnauthorized %s", 401, payload)
}

func (o *V1SshkeysPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPutForbidden creates a V1SshkeysPutForbidden with default headers values
func NewV1SshkeysPutForbidden() *V1SshkeysPutForbidden {
	return &V1SshkeysPutForbidden{}
}

/*
V1SshkeysPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1SshkeysPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys put forbidden response has a 2xx status code
func (o *V1SshkeysPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys put forbidden response has a 3xx status code
func (o *V1SshkeysPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys put forbidden response has a 4xx status code
func (o *V1SshkeysPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys put forbidden response has a 5xx status code
func (o *V1SshkeysPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys put forbidden response a status code equal to that given
func (o *V1SshkeysPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 sshkeys put forbidden response
func (o *V1SshkeysPutForbidden) Code() int {
	return 403
}

func (o *V1SshkeysPutForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutForbidden %s", 403, payload)
}

func (o *V1SshkeysPutForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutForbidden %s", 403, payload)
}

func (o *V1SshkeysPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPutNotFound creates a V1SshkeysPutNotFound with default headers values
func NewV1SshkeysPutNotFound() *V1SshkeysPutNotFound {
	return &V1SshkeysPutNotFound{}
}

/*
V1SshkeysPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1SshkeysPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys put not found response has a 2xx status code
func (o *V1SshkeysPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys put not found response has a 3xx status code
func (o *V1SshkeysPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys put not found response has a 4xx status code
func (o *V1SshkeysPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys put not found response has a 5xx status code
func (o *V1SshkeysPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys put not found response a status code equal to that given
func (o *V1SshkeysPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 sshkeys put not found response
func (o *V1SshkeysPutNotFound) Code() int {
	return 404
}

func (o *V1SshkeysPutNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutNotFound %s", 404, payload)
}

func (o *V1SshkeysPutNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutNotFound %s", 404, payload)
}

func (o *V1SshkeysPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPutConflict creates a V1SshkeysPutConflict with default headers values
func NewV1SshkeysPutConflict() *V1SshkeysPutConflict {
	return &V1SshkeysPutConflict{}
}

/*
V1SshkeysPutConflict describes a response with status code 409, with default header values.

Conflict
*/
type V1SshkeysPutConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys put conflict response has a 2xx status code
func (o *V1SshkeysPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys put conflict response has a 3xx status code
func (o *V1SshkeysPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys put conflict response has a 4xx status code
func (o *V1SshkeysPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys put conflict response has a 5xx status code
func (o *V1SshkeysPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys put conflict response a status code equal to that given
func (o *V1SshkeysPutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the v1 sshkeys put conflict response
func (o *V1SshkeysPutConflict) Code() int {
	return 409
}

func (o *V1SshkeysPutConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutConflict %s", 409, payload)
}

func (o *V1SshkeysPutConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutConflict %s", 409, payload)
}

func (o *V1SshkeysPutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPutGone creates a V1SshkeysPutGone with default headers values
func NewV1SshkeysPutGone() *V1SshkeysPutGone {
	return &V1SshkeysPutGone{}
}

/*
V1SshkeysPutGone describes a response with status code 410, with default header values.

Gone
*/
type V1SshkeysPutGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys put gone response has a 2xx status code
func (o *V1SshkeysPutGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys put gone response has a 3xx status code
func (o *V1SshkeysPutGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys put gone response has a 4xx status code
func (o *V1SshkeysPutGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys put gone response has a 5xx status code
func (o *V1SshkeysPutGone) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys put gone response a status code equal to that given
func (o *V1SshkeysPutGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the v1 sshkeys put gone response
func (o *V1SshkeysPutGone) Code() int {
	return 410
}

func (o *V1SshkeysPutGone) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutGone %s", 410, payload)
}

func (o *V1SshkeysPutGone) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutGone %s", 410, payload)
}

func (o *V1SshkeysPutGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPutGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPutInternalServerError creates a V1SshkeysPutInternalServerError with default headers values
func NewV1SshkeysPutInternalServerError() *V1SshkeysPutInternalServerError {
	return &V1SshkeysPutInternalServerError{}
}

/*
V1SshkeysPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1SshkeysPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys put internal server error response has a 2xx status code
func (o *V1SshkeysPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys put internal server error response has a 3xx status code
func (o *V1SshkeysPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys put internal server error response has a 4xx status code
func (o *V1SshkeysPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 sshkeys put internal server error response has a 5xx status code
func (o *V1SshkeysPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 sshkeys put internal server error response a status code equal to that given
func (o *V1SshkeysPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 sshkeys put internal server error response
func (o *V1SshkeysPutInternalServerError) Code() int {
	return 500
}

func (o *V1SshkeysPutInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutInternalServerError %s", 500, payload)
}

func (o *V1SshkeysPutInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/ssh-keys/{sshkey_id}][%d] v1SshkeysPutInternalServerError %s", 500, payload)
}

func (o *V1SshkeysPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
