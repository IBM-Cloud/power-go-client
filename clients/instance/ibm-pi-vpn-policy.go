package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_policies"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPIVpnPolicyClient ...
type IBMPIVpnPolicyClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// IBMPIVpnPolicyClient ...
func NewIBMPIVpnPolicyClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVpnPolicyClient {
	return &IBMPIVpnPolicyClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudvpnPolicies,
	}
}

// IKE Policies
// Get an IKE Policy
func (f *IBMPIVpnPolicyClient) GetIKEPolicy(ikePolicyID string) (*models.IKEPolicy, error) {

	// Create params and send request
	params := &params.PcloudIkepoliciesGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		IkePolicyID:     ikePolicyID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudIkepoliciesGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetVPNPolicyOperationFailed, ikePolicyID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get ike policy for policy id %s", ikePolicyID)
	}
	return resp.Payload, nil
}

// Create a IKE Policy
func (f *IBMPIVpnPolicyClient) CreateIKEPolicy(createBody *models.IKEPolicyCreate) (*models.IKEPolicy, error) {

	// Create params and send request
	params := &params.PcloudIkepoliciesPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudIkepoliciesPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVPNPolicyOperationFailed, f.cloudInstanceID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create VPN Policy")
}

// Update an IKE Policy
func (f *IBMPIVpnPolicyClient) UpdateIKEPolicy(ikePolicyID string, updateBody *models.IKEPolicyUpdate) (*models.IKEPolicy, error) {

	// Create params and send request
	params := &params.PcloudIkepoliciesPutParams{
		Body:            updateBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		IkePolicyID:     ikePolicyID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudIkepoliciesPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVPNPolicyOperationFailed, ikePolicyID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, fmt.Errorf("failed to Update VPN Policy")
}

// Get All IKE Policies
func (f *IBMPIVpnPolicyClient) GetAllIKEPolicies() (*models.IKEPolicies, error) {

	// Create params and send request
	params := &params.PcloudIkepoliciesGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudIkepoliciesGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get all ike policies: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all ike policies")
	}
	return resp.Payload, nil
}

// Delete an IKE Policy
func (f *IBMPIVpnPolicyClient) DeleteIKEPolicy(ikePolicyID string) error {

	// Create params and send request
	params := &params.PcloudIkepoliciesDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		IkePolicyID:     ikePolicyID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudIkepoliciesDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.DeleteVPNPolicyOperationFailed, ikePolicyID, err)
	}
	return nil
}

// IPSec Policies
// Get an IPSec Policy
func (f *IBMPIVpnPolicyClient) GetIPSecPolicy(ipsecPolicyID string) (*models.IPSecPolicy, error) {

	// Create params and send request
	params := &params.PcloudIpsecpoliciesGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		IpsecPolicyID:   ipsecPolicyID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudIpsecpoliciesGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetVPNPolicyOperationFailed, ipsecPolicyID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get ipsec policy for policy id %s", ipsecPolicyID)
	}
	return resp.Payload, nil
}

// Create an IPSec Policy
func (f *IBMPIVpnPolicyClient) CreateIPSecPolicy(createBody *models.IPSecPolicyCreate) (*models.IPSecPolicy, error) {

	// Create params and send request
	params := &params.PcloudIpsecpoliciesPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudIpsecpoliciesPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVPNPolicyOperationFailed, f.cloudInstanceID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create VPN Policy")
}

// Update an IPSec Policy
func (f *IBMPIVpnPolicyClient) UpdateIPSecPolicy(ipsecPolicyID string, updateBody *models.IPSecPolicyUpdate) (*models.IPSecPolicy, error) {

	// Create params and send request
	params := &params.PcloudIpsecpoliciesPutParams{
		Body:            updateBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		IpsecPolicyID:   ipsecPolicyID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudIpsecpoliciesPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVPNPolicyOperationFailed, ipsecPolicyID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, fmt.Errorf("failed to Update VPN Policy")
}

// Get all IPSec Policies
func (f *IBMPIVpnPolicyClient) GetAllIPSecPolicies() (*models.IPSecPolicies, error) {

	// Create params and send request
	params := &params.PcloudIpsecpoliciesGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudIpsecpoliciesGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get all ipsec policies: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all ipsec policies")
	}
	return resp.Payload, nil
}

// Delete an IPSec Policy
func (f *IBMPIVpnPolicyClient) DeleteIPSecPolicy(ipsecPolicyID string) error {

	// Create params and send request
	params := &params.PcloudIpsecpoliciesDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		IpsecPolicyID:   ipsecPolicyID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudIpsecpoliciesDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.DeleteVPNPolicyOperationFailed, ipsecPolicyID, err)
	}
	return nil
}
