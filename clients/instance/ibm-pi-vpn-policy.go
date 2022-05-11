package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_policies"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIVpnPolicyClient.
type IBMPIVpnPolicyClient struct {
	IBMPIClient
}

// IBMPIVpnPolicyClient.
func NewIBMPIVpnPolicyClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVpnPolicyClient {
	return &IBMPIVpnPolicyClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// IKE Policies.
// Get an IKE Policy.
func (f *IBMPIVpnPolicyClient) GetIKEPolicy(id string) (*models.IKEPolicy, error) {
	params := p.NewPcloudIkepoliciesGetParams().
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).
		WithIkePolicyID(id)
	resp, err := f.vpnPolicyRequest.PcloudIkepoliciesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVPNPolicyOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get ike policy for policy id %s", id)
	}
	return resp.Payload, nil
}

// Create an IKE Policy.
func (f *IBMPIVpnPolicyClient) CreateIKEPolicy(body *models.IKEPolicyCreate) (*models.IKEPolicy, error) {
	params := p.NewPcloudIkepoliciesPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	postok, err := f.vpnPolicyRequest.PcloudIkepoliciesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVPNPolicyOperationFailed, f.cloudInstanceID, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create VPN Policy")
}

// Update an IKE Policy.
func (f *IBMPIVpnPolicyClient) UpdateIKEPolicy(id string, body *models.IKEPolicyUpdate) (*models.IKEPolicy, error) {
	params := p.NewPcloudIkepoliciesPutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithIkePolicyID(id).
		WithTimeout(helpers.PIUpdateTimeOut)
	putok, err := f.vpnPolicyRequest.PcloudIkepoliciesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVPNPolicyOperationFailed, id, err)
	}
	if putok != nil && putok.Payload != nil {
		return putok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Update VPN Policy")
}

// Get All IKE Policies.
func (f *IBMPIVpnPolicyClient) GetAllIKEPolicies() (*models.IKEPolicies, error) {
	params := p.NewPcloudIkepoliciesGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.vpnPolicyRequest.PcloudIkepoliciesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get all ike policies: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all ike policies")
	}
	return resp.Payload, nil
}

// Delete an IKE Policy.
func (f *IBMPIVpnPolicyClient) DeleteIKEPolicy(id string) error {
	params := p.NewPcloudIkepoliciesDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithIkePolicyID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.vpnPolicyRequest.PcloudIkepoliciesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteVPNPolicyOperationFailed, id, err)
	}
	return nil
}

// IPSec Policies.
// Get an IPSec Policy.
func (f *IBMPIVpnPolicyClient) GetIPSecPolicy(id string) (*models.IPSecPolicy, error) {
	params := p.NewPcloudIpsecpoliciesGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithIpsecPolicyID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.vpnPolicyRequest.PcloudIpsecpoliciesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVPNPolicyOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get ipsec policy for policy id %s", id)
	}
	return resp.Payload, nil
}

// Create an IPSec Policy.
func (f *IBMPIVpnPolicyClient) CreateIPSecPolicy(body *models.IPSecPolicyCreate) (*models.IPSecPolicy, error) {
	params := p.NewPcloudIpsecpoliciesPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	postok, err := f.vpnPolicyRequest.PcloudIpsecpoliciesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVPNPolicyOperationFailed, f.cloudInstanceID, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create VPN Policy")
}

// Update an IPSec Policy.
func (f *IBMPIVpnPolicyClient) UpdateIPSecPolicy(id string, body *models.IPSecPolicyUpdate) (*models.IPSecPolicy, error) {
	params := p.NewPcloudIpsecpoliciesPutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithIpsecPolicyID(id).
		WithTimeout(helpers.PIUpdateTimeOut)
	putok, err := f.vpnPolicyRequest.PcloudIpsecpoliciesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVPNPolicyOperationFailed, id, err)
	}
	if putok != nil && putok.Payload != nil {
		return putok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Update VPN Policy")
}

// Get All IPSec Policies.
func (f *IBMPIVpnPolicyClient) GetAllIPSecPolicies() (*models.IPSecPolicies, error) {
	params := p.NewPcloudIpsecpoliciesGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.vpnPolicyRequest.PcloudIpsecpoliciesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get all ipsec policies: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all ipsec policies")
	}
	return resp.Payload, nil
}

// Delete an IPSec Policy.
func (f *IBMPIVpnPolicyClient) DeleteIPSecPolicy(id string) error {
	params := p.NewPcloudIpsecpoliciesDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithIpsecPolicyID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.vpnPolicyRequest.PcloudIpsecpoliciesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteVPNPolicyOperationFailed, id, err)
	}
	return nil
}
