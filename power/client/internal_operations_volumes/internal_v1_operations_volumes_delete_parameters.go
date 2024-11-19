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
)

// NewInternalV1OperationsVolumesDeleteParams creates a new InternalV1OperationsVolumesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1OperationsVolumesDeleteParams() *InternalV1OperationsVolumesDeleteParams {
	return &InternalV1OperationsVolumesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1OperationsVolumesDeleteParamsWithTimeout creates a new InternalV1OperationsVolumesDeleteParams object
// with the ability to set a timeout on a request.
func NewInternalV1OperationsVolumesDeleteParamsWithTimeout(timeout time.Duration) *InternalV1OperationsVolumesDeleteParams {
	return &InternalV1OperationsVolumesDeleteParams{
		timeout: timeout,
	}
}

// NewInternalV1OperationsVolumesDeleteParamsWithContext creates a new InternalV1OperationsVolumesDeleteParams object
// with the ability to set a context for a request.
func NewInternalV1OperationsVolumesDeleteParamsWithContext(ctx context.Context) *InternalV1OperationsVolumesDeleteParams {
	return &InternalV1OperationsVolumesDeleteParams{
		Context: ctx,
	}
}

// NewInternalV1OperationsVolumesDeleteParamsWithHTTPClient creates a new InternalV1OperationsVolumesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1OperationsVolumesDeleteParamsWithHTTPClient(client *http.Client) *InternalV1OperationsVolumesDeleteParams {
	return &InternalV1OperationsVolumesDeleteParams{
		HTTPClient: client,
	}
}

/*
InternalV1OperationsVolumesDeleteParams contains all the parameters to send to the API endpoint

	for the internal v1 operations volumes delete operation.

	Typically these are written to a http.Request.
*/
type InternalV1OperationsVolumesDeleteParams struct {

	/* ResourceCrn.

	   Encoded resource CRN, "/" to be encoded into "%2F", example 'crn:v1:staging:public:power-iaas:satloc_dal_clp2joc20ppo19876n50:a%2Fc7e6bd2517ad44eabbd61fcc25cf68d5:79bffc73-0035-4e7b-b34a-15da38927424:network:d8d51d44-053b-4df3-90b6-31fbe72ba600'
	*/
	ResourceCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 operations volumes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsVolumesDeleteParams) WithDefaults() *InternalV1OperationsVolumesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 operations volumes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsVolumesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 operations volumes delete params
func (o *InternalV1OperationsVolumesDeleteParams) WithTimeout(timeout time.Duration) *InternalV1OperationsVolumesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 operations volumes delete params
func (o *InternalV1OperationsVolumesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 operations volumes delete params
func (o *InternalV1OperationsVolumesDeleteParams) WithContext(ctx context.Context) *InternalV1OperationsVolumesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 operations volumes delete params
func (o *InternalV1OperationsVolumesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 operations volumes delete params
func (o *InternalV1OperationsVolumesDeleteParams) WithHTTPClient(client *http.Client) *InternalV1OperationsVolumesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 operations volumes delete params
func (o *InternalV1OperationsVolumesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceCrn adds the resourceCrn to the internal v1 operations volumes delete params
func (o *InternalV1OperationsVolumesDeleteParams) WithResourceCrn(resourceCrn string) *InternalV1OperationsVolumesDeleteParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the internal v1 operations volumes delete params
func (o *InternalV1OperationsVolumesDeleteParams) SetResourceCrn(resourceCrn string) {
	o.ResourceCrn = resourceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1OperationsVolumesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_crn
	if err := r.SetPathParam("resource_crn", o.ResourceCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
