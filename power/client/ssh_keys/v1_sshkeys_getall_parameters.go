// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

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

// NewV1SshkeysGetallParams creates a new V1SshkeysGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1SshkeysGetallParams() *V1SshkeysGetallParams {
	return &V1SshkeysGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1SshkeysGetallParamsWithTimeout creates a new V1SshkeysGetallParams object
// with the ability to set a timeout on a request.
func NewV1SshkeysGetallParamsWithTimeout(timeout time.Duration) *V1SshkeysGetallParams {
	return &V1SshkeysGetallParams{
		timeout: timeout,
	}
}

// NewV1SshkeysGetallParamsWithContext creates a new V1SshkeysGetallParams object
// with the ability to set a context for a request.
func NewV1SshkeysGetallParamsWithContext(ctx context.Context) *V1SshkeysGetallParams {
	return &V1SshkeysGetallParams{
		Context: ctx,
	}
}

// NewV1SshkeysGetallParamsWithHTTPClient creates a new V1SshkeysGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1SshkeysGetallParamsWithHTTPClient(client *http.Client) *V1SshkeysGetallParams {
	return &V1SshkeysGetallParams{
		HTTPClient: client,
	}
}

/*
V1SshkeysGetallParams contains all the parameters to send to the API endpoint

	for the v1 sshkeys getall operation.

	Typically these are written to a http.Request.
*/
type V1SshkeysGetallParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 sshkeys getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1SshkeysGetallParams) WithDefaults() *V1SshkeysGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 sshkeys getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1SshkeysGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 sshkeys getall params
func (o *V1SshkeysGetallParams) WithTimeout(timeout time.Duration) *V1SshkeysGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 sshkeys getall params
func (o *V1SshkeysGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 sshkeys getall params
func (o *V1SshkeysGetallParams) WithContext(ctx context.Context) *V1SshkeysGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 sshkeys getall params
func (o *V1SshkeysGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 sshkeys getall params
func (o *V1SshkeysGetallParams) WithHTTPClient(client *http.Client) *V1SshkeysGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 sshkeys getall params
func (o *V1SshkeysGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1SshkeysGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
