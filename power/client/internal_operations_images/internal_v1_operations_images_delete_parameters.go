// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_images

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
)

// NewInternalV1OperationsImagesDeleteParams creates a new InternalV1OperationsImagesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1OperationsImagesDeleteParams() *InternalV1OperationsImagesDeleteParams {
	return &InternalV1OperationsImagesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1OperationsImagesDeleteParamsWithTimeout creates a new InternalV1OperationsImagesDeleteParams object
// with the ability to set a timeout on a request.
func NewInternalV1OperationsImagesDeleteParamsWithTimeout(timeout time.Duration) *InternalV1OperationsImagesDeleteParams {
	return &InternalV1OperationsImagesDeleteParams{
		timeout: timeout,
	}
}

// NewInternalV1OperationsImagesDeleteParamsWithContext creates a new InternalV1OperationsImagesDeleteParams object
// with the ability to set a context for a request.
func NewInternalV1OperationsImagesDeleteParamsWithContext(ctx context.Context) *InternalV1OperationsImagesDeleteParams {
	return &InternalV1OperationsImagesDeleteParams{
		Context: ctx,
	}
}

// NewInternalV1OperationsImagesDeleteParamsWithHTTPClient creates a new InternalV1OperationsImagesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1OperationsImagesDeleteParamsWithHTTPClient(client *http.Client) *InternalV1OperationsImagesDeleteParams {
	return &InternalV1OperationsImagesDeleteParams{
		HTTPClient: client,
	}
}

/*
InternalV1OperationsImagesDeleteParams contains all the parameters to send to the API endpoint

	for the internal v1 operations images delete operation.

	Typically these are written to a http.Request.
*/
type InternalV1OperationsImagesDeleteParams struct {

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

	/* ResourceCrn.

	   Encoded resource CRN, "/" to be encoded into "%2F", example 'crn:v1:staging:public:power-iaas:satloc_dal_clp2joc20ppo19876n50:a%2Fc7e6bd2517ad44eabbd61fcc25cf68d5:79bffc73-0035-4e7b-b34a-15da38927424:network:d8d51d44-053b-4df3-90b6-31fbe72ba600'
	*/
	ResourceCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 operations images delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsImagesDeleteParams) WithDefaults() *InternalV1OperationsImagesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 operations images delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsImagesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) WithTimeout(timeout time.Duration) *InternalV1OperationsImagesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) WithContext(ctx context.Context) *InternalV1OperationsImagesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) WithHTTPClient(client *http.Client) *InternalV1OperationsImagesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) WithAuthorization(authorization string) *InternalV1OperationsImagesDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithCRN adds the cRN to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) WithCRN(cRN string) *InternalV1OperationsImagesDeleteParams {
	o.SetCRN(cRN)
	return o
}

// SetCRN adds the cRN to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) SetCRN(cRN string) {
	o.CRN = cRN
}

// WithIBMUserAuthorization adds the iBMUserAuthorization to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) WithIBMUserAuthorization(iBMUserAuthorization string) *InternalV1OperationsImagesDeleteParams {
	o.SetIBMUserAuthorization(iBMUserAuthorization)
	return o
}

// SetIBMUserAuthorization adds the iBMUserAuthorization to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) SetIBMUserAuthorization(iBMUserAuthorization string) {
	o.IBMUserAuthorization = iBMUserAuthorization
}

// WithResourceCrn adds the resourceCrn to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) WithResourceCrn(resourceCrn string) *InternalV1OperationsImagesDeleteParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the internal v1 operations images delete params
func (o *InternalV1OperationsImagesDeleteParams) SetResourceCrn(resourceCrn string) {
	o.ResourceCrn = resourceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1OperationsImagesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resource_crn
	if err := r.SetPathParam("resource_crn", o.ResourceCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}