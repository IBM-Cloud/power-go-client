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

// NewV1RoutesGetParams creates a new V1RoutesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1RoutesGetParams() *V1RoutesGetParams {
	return &V1RoutesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1RoutesGetParamsWithTimeout creates a new V1RoutesGetParams object
// with the ability to set a timeout on a request.
func NewV1RoutesGetParamsWithTimeout(timeout time.Duration) *V1RoutesGetParams {
	return &V1RoutesGetParams{
		timeout: timeout,
	}
}

// NewV1RoutesGetParamsWithContext creates a new V1RoutesGetParams object
// with the ability to set a context for a request.
func NewV1RoutesGetParamsWithContext(ctx context.Context) *V1RoutesGetParams {
	return &V1RoutesGetParams{
		Context: ctx,
	}
}

// NewV1RoutesGetParamsWithHTTPClient creates a new V1RoutesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1RoutesGetParamsWithHTTPClient(client *http.Client) *V1RoutesGetParams {
	return &V1RoutesGetParams{
		HTTPClient: client,
	}
}

/*
V1RoutesGetParams contains all the parameters to send to the API endpoint

	for the v1 routes get operation.

	Typically these are written to a http.Request.
*/
type V1RoutesGetParams struct {

	/* RouteID.

	   Route ID
	*/
	RouteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 routes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1RoutesGetParams) WithDefaults() *V1RoutesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 routes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1RoutesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 routes get params
func (o *V1RoutesGetParams) WithTimeout(timeout time.Duration) *V1RoutesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 routes get params
func (o *V1RoutesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 routes get params
func (o *V1RoutesGetParams) WithContext(ctx context.Context) *V1RoutesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 routes get params
func (o *V1RoutesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 routes get params
func (o *V1RoutesGetParams) WithHTTPClient(client *http.Client) *V1RoutesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 routes get params
func (o *V1RoutesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRouteID adds the routeID to the v1 routes get params
func (o *V1RoutesGetParams) WithRouteID(routeID string) *V1RoutesGetParams {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the v1 routes get params
func (o *V1RoutesGetParams) SetRouteID(routeID string) {
	o.RouteID = routeID
}

// WriteToRequest writes these params to a swagger request
func (o *V1RoutesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param route_id
	if err := r.SetPathParam("route_id", o.RouteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
