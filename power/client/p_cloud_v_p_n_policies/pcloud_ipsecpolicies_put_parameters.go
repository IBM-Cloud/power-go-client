// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

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

// NewPcloudIpsecpoliciesPutParams creates a new PcloudIpsecpoliciesPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudIpsecpoliciesPutParams() *PcloudIpsecpoliciesPutParams {
	return &PcloudIpsecpoliciesPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudIpsecpoliciesPutParamsWithTimeout creates a new PcloudIpsecpoliciesPutParams object
// with the ability to set a timeout on a request.
func NewPcloudIpsecpoliciesPutParamsWithTimeout(timeout time.Duration) *PcloudIpsecpoliciesPutParams {
	return &PcloudIpsecpoliciesPutParams{
		timeout: timeout,
	}
}

// NewPcloudIpsecpoliciesPutParamsWithContext creates a new PcloudIpsecpoliciesPutParams object
// with the ability to set a context for a request.
func NewPcloudIpsecpoliciesPutParamsWithContext(ctx context.Context) *PcloudIpsecpoliciesPutParams {
	return &PcloudIpsecpoliciesPutParams{
		Context: ctx,
	}
}

// NewPcloudIpsecpoliciesPutParamsWithHTTPClient creates a new PcloudIpsecpoliciesPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudIpsecpoliciesPutParamsWithHTTPClient(client *http.Client) *PcloudIpsecpoliciesPutParams {
	return &PcloudIpsecpoliciesPutParams{
		HTTPClient: client,
	}
}

/* PcloudIpsecpoliciesPutParams contains all the parameters to send to the API endpoint
   for the pcloud ipsecpolicies put operation.

   Typically these are written to a http.Request.
*/
type PcloudIpsecpoliciesPutParams struct {

	/* Body.

	   Parameters for the update of an IPSec Policy
	*/
	Body *models.IPSecPolicyUpdate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* IpsecPolicyID.

	   ID of a IPSec Policy
	*/
	IpsecPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud ipsecpolicies put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIpsecpoliciesPutParams) WithDefaults() *PcloudIpsecpoliciesPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud ipsecpolicies put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIpsecpoliciesPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) WithTimeout(timeout time.Duration) *PcloudIpsecpoliciesPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) WithContext(ctx context.Context) *PcloudIpsecpoliciesPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) WithHTTPClient(client *http.Client) *PcloudIpsecpoliciesPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) WithBody(body *models.IPSecPolicyUpdate) *PcloudIpsecpoliciesPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) SetBody(body *models.IPSecPolicyUpdate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudIpsecpoliciesPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithIpsecPolicyID adds the ipsecPolicyID to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) WithIpsecPolicyID(ipsecPolicyID string) *PcloudIpsecpoliciesPutParams {
	o.SetIpsecPolicyID(ipsecPolicyID)
	return o
}

// SetIpsecPolicyID adds the ipsecPolicyId to the pcloud ipsecpolicies put params
func (o *PcloudIpsecpoliciesPutParams) SetIpsecPolicyID(ipsecPolicyID string) {
	o.IpsecPolicyID = ipsecPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudIpsecpoliciesPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param ipsec_policy_id
	if err := r.SetPathParam("ipsec_policy_id", o.IpsecPolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
