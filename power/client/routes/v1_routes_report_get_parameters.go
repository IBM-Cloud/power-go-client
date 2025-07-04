// Code generated by go-swagger; DO NOT EDIT.

package routes

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

// NewV1RoutesReportGetParams creates a new V1RoutesReportGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1RoutesReportGetParams() *V1RoutesReportGetParams {
	return &V1RoutesReportGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1RoutesReportGetParamsWithTimeout creates a new V1RoutesReportGetParams object
// with the ability to set a timeout on a request.
func NewV1RoutesReportGetParamsWithTimeout(timeout time.Duration) *V1RoutesReportGetParams {
	return &V1RoutesReportGetParams{
		timeout: timeout,
	}
}

// NewV1RoutesReportGetParamsWithContext creates a new V1RoutesReportGetParams object
// with the ability to set a context for a request.
func NewV1RoutesReportGetParamsWithContext(ctx context.Context) *V1RoutesReportGetParams {
	return &V1RoutesReportGetParams{
		Context: ctx,
	}
}

// NewV1RoutesReportGetParamsWithHTTPClient creates a new V1RoutesReportGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1RoutesReportGetParamsWithHTTPClient(client *http.Client) *V1RoutesReportGetParams {
	return &V1RoutesReportGetParams{
		HTTPClient: client,
	}
}

/*
V1RoutesReportGetParams contains all the parameters to send to the API endpoint

	for the v1 routes report get operation.

	Typically these are written to a http.Request.
*/
type V1RoutesReportGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 routes report get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1RoutesReportGetParams) WithDefaults() *V1RoutesReportGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 routes report get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1RoutesReportGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 routes report get params
func (o *V1RoutesReportGetParams) WithTimeout(timeout time.Duration) *V1RoutesReportGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 routes report get params
func (o *V1RoutesReportGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 routes report get params
func (o *V1RoutesReportGetParams) WithContext(ctx context.Context) *V1RoutesReportGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 routes report get params
func (o *V1RoutesReportGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 routes report get params
func (o *V1RoutesReportGetParams) WithHTTPClient(client *http.Client) *V1RoutesReportGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 routes report get params
func (o *V1RoutesReportGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1RoutesReportGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
