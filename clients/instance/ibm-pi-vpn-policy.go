package instance

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_policies"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIVpnPolicyClient ...
type IBMPIVpnPolicyClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// IBMPIVpnPolicyClient ...
func NewIBMPIVpnPolicyClient(sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVpnPolicyClient {
	return &IBMPIVpnPolicyClient{
		session:         sess,
		powerinstanceid: cloudInstanceID,
	}
}

// IKE Policies

// Get information about a IKE Policy
func (f *IBMPIVpnPolicyClient) GetIKEPolicy(id, cloudInstanceID string) (*models.IKEPolicy, error) {
	return f.GetIKEPolicyWithContext(context.Background(), id, cloudInstanceID)
}

// Get information about a IKE Policy with Context
func (f *IBMPIVpnPolicyClient) GetIKEPolicyWithContext(ctx context.Context, id, cloudInstanceID string) (ikePolicy *models.IKEPolicy, err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithIkePolicyID(id)
	resp, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	ikePolicy = resp.Payload
	return
}

// Create a IKE Policy
func (f *IBMPIVpnPolicyClient) CreateIKEPolicy(body *models.IKEPolicyCreate, cloudInstanceID string) (ikePolicy *models.IKEPolicy, err error) {
	return f.CreateIKEPolicyWithContext(context.Background(), body, cloudInstanceID)
}

// Create a IKE Policy with context
func (f *IBMPIVpnPolicyClient) CreateIKEPolicyWithContext(ctx context.Context, body *models.IKEPolicyCreate, cloudInstanceID string) (ikePolicy *models.IKEPolicy, err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesPostParamsWithContext(ctx).
		WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithBody(body)
	postok, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesPost(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if postok != nil {
		ikePolicy = postok.Payload
	}
	return
}

// Update a IKE Policy
func (f *IBMPIVpnPolicyClient) UpdateIKEPolicy(body *models.IKEPolicyUpdate, id, cloudInstanceID string) (ikePolicy *models.IKEPolicy, err error) {
	return f.UpdateIKEPolicyWithContext(context.Background(), body, id, cloudInstanceID)
}

// Update a IKE Policy with context
func (f *IBMPIVpnPolicyClient) UpdateIKEPolicyWithContext(ctx context.Context, body *models.IKEPolicyUpdate, id, cloudInstanceID string) (ikePolicy *models.IKEPolicy, err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesPutParamsWithContext(ctx).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithBody(body).
		WithIkePolicyID(id)
	putok, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesPut(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if putok != nil {
		ikePolicy = putok.Payload
	}
	return
}

// Get all IKE Policies
func (f *IBMPIVpnPolicyClient) GetAllIKEPolicies(cloudInstanceID string) (ikePolicies *models.IKEPolicies, err error) {
	return f.GetAllIKEPoliciesWithContext(context.Background(), cloudInstanceID)
}

// Get all IKE Policies with context
func (f *IBMPIVpnPolicyClient) GetAllIKEPoliciesWithContext(ctx context.Context, cloudInstanceID string) (ikePolicies *models.IKEPolicies, err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesGetallParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID)
	resp, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesGetall(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	ikePolicies = resp.Payload
	return
}

// Delete a IKE Policy
func (f *IBMPIVpnPolicyClient) DeleteIKEPolicy(id, cloudInstanceID string) (err error) {
	return f.DeleteIKEPolicyWithContext(context.Background(), id, cloudInstanceID)
}

// Delete a IKE Policy with context
func (f *IBMPIVpnPolicyClient) DeleteIKEPolicyWithContext(ctx context.Context, id, cloudInstanceID string) (err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithIkePolicyID(id)
	_, err = f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	return
}

// IPSec Policies

// Get information about a IPSec Policy
func (f *IBMPIVpnPolicyClient) GetIPSecPolicy(id, cloudInstanceID string) (*models.IPSecPolicy, error) {
	return f.GetIPSecPolicyWithContext(context.Background(), id, cloudInstanceID)
}

// Get information about a IPSec Policy with Context
func (f *IBMPIVpnPolicyClient) GetIPSecPolicyWithContext(ctx context.Context, id, cloudInstanceID string) (ipsecPolicy *models.IPSecPolicy, err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithIpsecPolicyID(id)
	resp, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	ipsecPolicy = resp.Payload
	return
}

// Create a IPSec Policy
func (f *IBMPIVpnPolicyClient) CreateIPSecPolicy(body *models.IPSecPolicyCreate, cloudInstanceID string) (ipsecPolicy *models.IPSecPolicy, err error) {
	return f.CreateIPSecPolicyWithContext(context.Background(), body, cloudInstanceID)
}

// Create a IPSec Policy with context
func (f *IBMPIVpnPolicyClient) CreateIPSecPolicyWithContext(ctx context.Context, body *models.IPSecPolicyCreate, cloudInstanceID string) (ipsecPolicy *models.IPSecPolicy, err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesPostParamsWithContext(ctx).
		WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithBody(body)
	postok, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesPost(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if postok != nil {
		ipsecPolicy = postok.Payload
	}
	return
}

// Update a IPSec Policy
func (f *IBMPIVpnPolicyClient) UpdateIPSecPolicy(body *models.IPSecPolicyUpdate, id, cloudInstanceID string) (ipsecPolicy *models.IPSecPolicy, err error) {
	return f.UpdateIPSecPolicyWithContext(context.Background(), body, id, cloudInstanceID)
}

// Update a IPSec Policy with context
func (f *IBMPIVpnPolicyClient) UpdateIPSecPolicyWithContext(ctx context.Context, body *models.IPSecPolicyUpdate, id, cloudInstanceID string) (ipsecPolicy *models.IPSecPolicy, err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesPutParamsWithContext(ctx).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithBody(body).
		WithIpsecPolicyID(id)
	putok, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesPut(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if putok != nil {
		ipsecPolicy = putok.Payload
	}
	return
}

// Get all IPSec Policies
func (f *IBMPIVpnPolicyClient) GetAllIPSecPolicies(cloudInstanceID string) (IPSecPolicies *models.IPSecPolicies, err error) {
	return f.GetAllIPSecPoliciesWithContext(context.Background(), cloudInstanceID)
}

// Get all IPSec Policies with context
func (f *IBMPIVpnPolicyClient) GetAllIPSecPoliciesWithContext(ctx context.Context, cloudInstanceID string) (ipsecPolicies *models.IPSecPolicies, err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesGetallParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID)
	resp, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesGetall(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	ipsecPolicies = resp.Payload
	return
}

// Delete a IPSec Policy
func (f *IBMPIVpnPolicyClient) DeleteIPSecPolicy(id, cloudInstanceID string) (err error) {
	return f.DeleteIPSecPolicyWithContext(context.Background(), id, cloudInstanceID)
}

// Delete a IPSec Policy with context
func (f *IBMPIVpnPolicyClient) DeleteIPSecPolicyWithContext(ctx context.Context, id, cloudInstanceID string) (err error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithIpsecPolicyID(id)
	_, err = f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	return
}
