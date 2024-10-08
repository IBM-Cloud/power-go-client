// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_virtual_serial_number

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

// NewPcloudPvminstancesVirtualserialnumberGetParams creates a new PcloudPvminstancesVirtualserialnumberGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesVirtualserialnumberGetParams() *PcloudPvminstancesVirtualserialnumberGetParams {
	return &PcloudPvminstancesVirtualserialnumberGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesVirtualserialnumberGetParamsWithTimeout creates a new PcloudPvminstancesVirtualserialnumberGetParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesVirtualserialnumberGetParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesVirtualserialnumberGetParams {
	return &PcloudPvminstancesVirtualserialnumberGetParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesVirtualserialnumberGetParamsWithContext creates a new PcloudPvminstancesVirtualserialnumberGetParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesVirtualserialnumberGetParamsWithContext(ctx context.Context) *PcloudPvminstancesVirtualserialnumberGetParams {
	return &PcloudPvminstancesVirtualserialnumberGetParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesVirtualserialnumberGetParamsWithHTTPClient creates a new PcloudPvminstancesVirtualserialnumberGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesVirtualserialnumberGetParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesVirtualserialnumberGetParams {
	return &PcloudPvminstancesVirtualserialnumberGetParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesVirtualserialnumberGetParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances virtualserialnumber get operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesVirtualserialnumberGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances virtualserialnumber get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesVirtualserialnumberGetParams) WithDefaults() *PcloudPvminstancesVirtualserialnumberGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances virtualserialnumber get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesVirtualserialnumberGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesVirtualserialnumberGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) WithContext(ctx context.Context) *PcloudPvminstancesVirtualserialnumberGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesVirtualserialnumberGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesVirtualserialnumberGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesVirtualserialnumberGetParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances virtualserialnumber get params
func (o *PcloudPvminstancesVirtualserialnumberGetParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesVirtualserialnumberGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
