// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPvminstancesConsolePutParams creates a new PcloudPvminstancesConsolePutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesConsolePutParams() *PcloudPvminstancesConsolePutParams {
	return &PcloudPvminstancesConsolePutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesConsolePutParamsWithTimeout creates a new PcloudPvminstancesConsolePutParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesConsolePutParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesConsolePutParams {
	return &PcloudPvminstancesConsolePutParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesConsolePutParamsWithContext creates a new PcloudPvminstancesConsolePutParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesConsolePutParamsWithContext(ctx context.Context) *PcloudPvminstancesConsolePutParams {
	return &PcloudPvminstancesConsolePutParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesConsolePutParamsWithHTTPClient creates a new PcloudPvminstancesConsolePutParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesConsolePutParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesConsolePutParams {
	return &PcloudPvminstancesConsolePutParams{
		HTTPClient: client,
	}
}

/* PcloudPvminstancesConsolePutParams contains all the parameters to send to the API endpoint
   for the pcloud pvminstances console put operation.

   Typically these are written to a http.Request.
*/
type PcloudPvminstancesConsolePutParams struct {

	/* Body.

	   Parameters to update a PVMInstance console required codepage
	*/
	Body *models.ConsoleLanguage

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

// WithDefaults hydrates default values in the pcloud pvminstances console put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesConsolePutParams) WithDefaults() *PcloudPvminstancesConsolePutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances console put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesConsolePutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesConsolePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) WithContext(ctx context.Context) *PcloudPvminstancesConsolePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesConsolePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) WithBody(body *models.ConsoleLanguage) *PcloudPvminstancesConsolePutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) SetBody(body *models.ConsoleLanguage) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesConsolePutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesConsolePutParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances console put params
func (o *PcloudPvminstancesConsolePutParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesConsolePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
