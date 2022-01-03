// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_snapshots

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

// NewPcloudCloudinstancesSnapshotsGetParams creates a new PcloudCloudinstancesSnapshotsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudinstancesSnapshotsGetParams() *PcloudCloudinstancesSnapshotsGetParams {
	return &PcloudCloudinstancesSnapshotsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesSnapshotsGetParamsWithTimeout creates a new PcloudCloudinstancesSnapshotsGetParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudinstancesSnapshotsGetParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesSnapshotsGetParams {
	return &PcloudCloudinstancesSnapshotsGetParams{
		timeout: timeout,
	}
}

// NewPcloudCloudinstancesSnapshotsGetParamsWithContext creates a new PcloudCloudinstancesSnapshotsGetParams object
// with the ability to set a context for a request.
func NewPcloudCloudinstancesSnapshotsGetParamsWithContext(ctx context.Context) *PcloudCloudinstancesSnapshotsGetParams {
	return &PcloudCloudinstancesSnapshotsGetParams{
		Context: ctx,
	}
}

// NewPcloudCloudinstancesSnapshotsGetParamsWithHTTPClient creates a new PcloudCloudinstancesSnapshotsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudinstancesSnapshotsGetParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesSnapshotsGetParams {
	return &PcloudCloudinstancesSnapshotsGetParams{
		HTTPClient: client,
	}
}

/* PcloudCloudinstancesSnapshotsGetParams contains all the parameters to send to the API endpoint
   for the pcloud cloudinstances snapshots get operation.

   Typically these are written to a http.Request.
*/
type PcloudCloudinstancesSnapshotsGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* SnapshotID.

	   PVM Instance snapshot id
	*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudinstances snapshots get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesSnapshotsGetParams) WithDefaults() *PcloudCloudinstancesSnapshotsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudinstances snapshots get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesSnapshotsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesSnapshotsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) WithContext(ctx context.Context) *PcloudCloudinstancesSnapshotsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesSnapshotsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesSnapshotsGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithSnapshotID adds the snapshotID to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) WithSnapshotID(snapshotID string) *PcloudCloudinstancesSnapshotsGetParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the pcloud cloudinstances snapshots get params
func (o *PcloudCloudinstancesSnapshotsGetParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesSnapshotsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param snapshot_id
	if err := r.SetPathParam("snapshot_id", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
