// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewInternalV1OperationsVolumesPostParams creates a new InternalV1OperationsVolumesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1OperationsVolumesPostParams() *InternalV1OperationsVolumesPostParams {
	return &InternalV1OperationsVolumesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1OperationsVolumesPostParamsWithTimeout creates a new InternalV1OperationsVolumesPostParams object
// with the ability to set a timeout on a request.
func NewInternalV1OperationsVolumesPostParamsWithTimeout(timeout time.Duration) *InternalV1OperationsVolumesPostParams {
	return &InternalV1OperationsVolumesPostParams{
		timeout: timeout,
	}
}

// NewInternalV1OperationsVolumesPostParamsWithContext creates a new InternalV1OperationsVolumesPostParams object
// with the ability to set a context for a request.
func NewInternalV1OperationsVolumesPostParamsWithContext(ctx context.Context) *InternalV1OperationsVolumesPostParams {
	return &InternalV1OperationsVolumesPostParams{
		Context: ctx,
	}
}

// NewInternalV1OperationsVolumesPostParamsWithHTTPClient creates a new InternalV1OperationsVolumesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1OperationsVolumesPostParamsWithHTTPClient(client *http.Client) *InternalV1OperationsVolumesPostParams {
	return &InternalV1OperationsVolumesPostParams{
		HTTPClient: client,
	}
}

/*
InternalV1OperationsVolumesPostParams contains all the parameters to send to the API endpoint

	for the internal v1 operations volumes post operation.

	Typically these are written to a http.Request.
*/
type InternalV1OperationsVolumesPostParams struct {

	/* Authorization.

	   Authentication of the service token
	*/
	Authorization string

	/* CRN.

	   the CRN of the workspace
	*/
	CRN string

	/* IBMUserAuthorization.

	   Authentication of the operation account user
	*/
	IBMUserAuthorization string

	/* Body.

	   Parameters for creating Volume CRN
	*/
	Body *models.InternalOperationsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 operations volumes post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsVolumesPostParams) WithDefaults() *InternalV1OperationsVolumesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 operations volumes post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsVolumesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) WithTimeout(timeout time.Duration) *InternalV1OperationsVolumesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) WithContext(ctx context.Context) *InternalV1OperationsVolumesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) WithHTTPClient(client *http.Client) *InternalV1OperationsVolumesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) WithAuthorization(authorization string) *InternalV1OperationsVolumesPostParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithCRN adds the cRN to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) WithCRN(cRN string) *InternalV1OperationsVolumesPostParams {
	o.SetCRN(cRN)
	return o
}

// SetCRN adds the cRN to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) SetCRN(cRN string) {
	o.CRN = cRN
}

// WithIBMUserAuthorization adds the iBMUserAuthorization to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) WithIBMUserAuthorization(iBMUserAuthorization string) *InternalV1OperationsVolumesPostParams {
	o.SetIBMUserAuthorization(iBMUserAuthorization)
	return o
}

// SetIBMUserAuthorization adds the iBMUserAuthorization to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) SetIBMUserAuthorization(iBMUserAuthorization string) {
	o.IBMUserAuthorization = iBMUserAuthorization
}

// WithBody adds the body to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) WithBody(body *models.InternalOperationsRequest) *InternalV1OperationsVolumesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the internal v1 operations volumes post params
func (o *InternalV1OperationsVolumesPostParams) SetBody(body *models.InternalOperationsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1OperationsVolumesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// header param CRN
	if err := r.SetHeaderParam("CRN", o.CRN); err != nil {
		return err
	}

	// header param IBM-UserAuthorization
	if err := r.SetHeaderParam("IBM-UserAuthorization", o.IBMUserAuthorization); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}