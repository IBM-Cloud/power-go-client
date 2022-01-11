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

// NewPcloudPvminstancesActionPostParams creates a new PcloudPvminstancesActionPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesActionPostParams() *PcloudPvminstancesActionPostParams {
	return &PcloudPvminstancesActionPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesActionPostParamsWithTimeout creates a new PcloudPvminstancesActionPostParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesActionPostParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesActionPostParams {
	return &PcloudPvminstancesActionPostParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesActionPostParamsWithContext creates a new PcloudPvminstancesActionPostParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesActionPostParamsWithContext(ctx context.Context) *PcloudPvminstancesActionPostParams {
	return &PcloudPvminstancesActionPostParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesActionPostParamsWithHTTPClient creates a new PcloudPvminstancesActionPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesActionPostParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesActionPostParams {
	return &PcloudPvminstancesActionPostParams{
		HTTPClient: client,
	}
}

/* PcloudPvminstancesActionPostParams contains all the parameters to send to the API endpoint
   for the pcloud pvminstances action post operation.

   Typically these are written to a http.Request.
*/
type PcloudPvminstancesActionPostParams struct {

	/* Body.

	   Parameters for the desired action
	*/
	Body *models.PVMInstanceAction

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

// WithDefaults hydrates default values in the pcloud pvminstances action post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesActionPostParams) WithDefaults() *PcloudPvminstancesActionPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances action post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesActionPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesActionPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) WithContext(ctx context.Context) *PcloudPvminstancesActionPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesActionPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) WithBody(body *models.PVMInstanceAction) *PcloudPvminstancesActionPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) SetBody(body *models.PVMInstanceAction) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesActionPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesActionPostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances action post params
func (o *PcloudPvminstancesActionPostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesActionPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
