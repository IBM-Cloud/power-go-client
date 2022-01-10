package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestVPNCreateGetDelete(t *testing.T) {
	testVPNPreCheck(t)

	// create session and clients
	session := getSession(t)
	vpnClient := client.NewIBMPIVpnConnectionClient(context.Background(), session, cloudInstanceID)
	vpnPolicyClient := client.NewIBMPIVpnPolicyClient(context.Background(), session, cloudInstanceID)
	jobClient := client.NewIBMPIJobClient(context.Background(), session, cloudInstanceID)

	// CREATE IKE Policy
	ikeKey := models.KeyLifetime(ipSecPolicyKeyLifetime)
	bodyIKE := &models.IKEPolicyCreate{
		DhGroup:      &ikePolicyDhGroup,
		Encryption:   &ikePolicyEncryption,
		KeyLifetime:  &ikeKey,
		Name:         &ikePolicyName,
		PresharedKey: &ikePolicyPresharedKey,
		Version:      &ikePolicyVersion,
	}
	createIkeResp, err := vpnPolicyClient.CreateIKEPolicy(bodyIKE)
	require.Nil(t, err)
	ikePolicyID := *createIkeResp.ID
	testMessage("CREATE IKE Policy", ikePolicyID, *createIkeResp)

	// GET IKE Policy
	getIkePolicy, err := vpnPolicyClient.GetIKEPolicy(ikePolicyID)
	require.Nil(t, err)
	testMessage("GET IKE Policy", ikePolicyID, *getIkePolicy)
	// verify variables match
	require.Equal(t, *getIkePolicy.DhGroup, ikePolicyDhGroup)
	require.Equal(t, *getIkePolicy.Encryption, ikePolicyEncryption)
	require.Equal(t, *getIkePolicy.KeyLifetime, models.KeyLifetime(ikePolicyKeyLifetime))
	require.Equal(t, *getIkePolicy.Name, ikePolicyName)
	require.Equal(t, *getIkePolicy.Version, ikePolicyVersion)

	// CREATE IPSec Policy
	ipSecKey := models.KeyLifetime(ipSecPolicyKeyLifetime)
	bodyIPSEC := &models.IPSecPolicyCreate{
		Authentication: models.IPSECPolicyAuthenticationHmacDashSha1Dash96,
		DhGroup:        &ipSecPolicyDhGroup,
		Encryption:     &ipSecPolicyEncryption,
		KeyLifetime:    &ipSecKey,
		Name:           &ipSecPolicyName,
		Pfs:            &ipSecPolicyPfs,
	}
	createIpsecResp, err := vpnPolicyClient.CreateIPSecPolicy(bodyIPSEC)
	require.Nil(t, err)
	ipsecPolicyID := *createIpsecResp.ID
	testMessage("CREATE IPSec Policy", ipsecPolicyID, *createIpsecResp)

	// GET IPSec Policy
	getIpsecPolicy, err := vpnPolicyClient.GetIPSecPolicy(ipsecPolicyID)
	require.Nil(t, err)
	testMessage("GET IPSec Policy", ipsecPolicyID, *getIpsecPolicy)
	// verify variables match
	require.Equal(t, *getIpsecPolicy.Authentication, models.IPSECPolicyAuthenticationHmacDashSha1Dash96)
	require.Equal(t, *getIpsecPolicy.DhGroup, ipSecPolicyDhGroup)
	require.Equal(t, *getIpsecPolicy.Encryption, ipSecPolicyEncryption)
	require.Equal(t, *getIpsecPolicy.KeyLifetime, models.KeyLifetime(ipSecPolicyKeyLifetime))
	require.Equal(t, *getIpsecPolicy.Name, ipSecPolicyName)
	require.Equal(t, *getIpsecPolicy.Pfs, ipSecPolicyPfs)

	// CREATE VPN
	peerGatewayAddress := models.PeerGatewayAddress(vpnPeerGatewayAddress)
	body := &models.VPNConnectionCreate{
		IkePolicy:          &ikePolicyID,
		IPSecPolicy:        &ipsecPolicyID,
		Mode:               &vpnMode,
		Name:               &vpnName,
		Networks:           []string{vpnNetworkID},
		PeerGatewayAddress: &peerGatewayAddress,
		PeerSubnets:        []string{vpnPeerSubnet},
	}
	createVpnResp, err := vpnClient.Create(body)
	require.Nil(t, err)
	waitForJobState(t, jobClient, *createVpnResp.JobRef.ID)
	vpnID := *createVpnResp.ID
	testMessage("CREATE VPN", vpnID, *createVpnResp)

	// GET VPN
	getVpn, err := vpnClient.Get(vpnID)
	require.Nil(t, err)
	testMessage("GET VPN", vpnID, *getVpn)
	// verify variables match
	require.Equal(t, *getVpn.IkePolicy.ID, *getIkePolicy.ID)
	require.Equal(t, *getVpn.IPSecPolicy.ID, *getIpsecPolicy.ID)
	require.Equal(t, *getVpn.Mode, vpnMode)
	require.Equal(t, *getVpn.Name, vpnName)
	require.Equal(t, getVpn.NetworkIDs, []string{vpnNetworkID})
	require.Equal(t, getVpn.PeerGatewayAddress, models.PeerGatewayAddress(vpnPeerGatewayAddress))
	require.Equal(t, getVpn.PeerSubnets, []string{vpnPeerSubnet})

	// DELETE VPN
	resp, err := vpnClient.Delete(vpnID)
	require.Nil(t, err)
	waitForJobState(t, jobClient, *resp.ID)
	testMessage("DELETE VPN", vpnID, *resp)

	// DELETE IKE Policy
	err = vpnPolicyClient.DeleteIKEPolicy(ikePolicyID)
	require.Nil(t, err)
	testMessage("DELETE IKE Policy", ikePolicyID, nil)

	// DELETE IPSec Policy
	err = vpnPolicyClient.DeleteIPSecPolicy(ipsecPolicyID)
	require.Nil(t, err)
	testMessage("DELETE IPSec Policy", ipsecPolicyID, nil)
}

func TestVPNGetAll(t *testing.T) {
	testPreCheck(t)

	// create session and clients
	session := getSession(t)
	vpnClient := client.NewIBMPIVpnConnectionClient(context.Background(), session, cloudInstanceID)
	vpnPolicyClient := client.NewIBMPIVpnPolicyClient(context.Background(), session, cloudInstanceID)

	// VPN
	getAllResp, err := vpnClient.GetAll()
	require.Nil(t, err)
	testMessage("GET ALL VPNs", "", *getAllResp)

	// IKE Policy
	getAllIKEResp, err := vpnPolicyClient.GetAllIKEPolicies()
	require.Nil(t, err)
	testMessage("GET ALL IKE Policies", "", *getAllIKEResp)

	// IPSec Policy
	getAllIPSecResp, err := vpnPolicyClient.GetAllIPSecPolicies()
	require.Nil(t, err)
	testMessage("GET ALL IPSec Policies", "", *getAllIPSecResp)
}
