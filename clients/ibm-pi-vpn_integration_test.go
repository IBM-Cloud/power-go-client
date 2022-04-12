package instance_test

import (
	"context"
	"testing"

	utl "internal/testutils"

	client "github.com/IBM-Cloud/power-go-client/clients"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestVPN(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and clients
	utl.VPNPreCheck(t)
	utl.IKEPolicyPreCheck(t)
	utl.IPSecPreCheck(t)
	session := utl.TestSession(t)
	vpnClient := client.NewIBMPIVpnConnectionClient(context.Background(), session, utl.CloudInstanceID)
	vpnPolicyClient := client.NewIBMPIVpnPolicyClient(context.Background(), session, utl.CloudInstanceID)
	jobClient := client.NewIBMPIJobClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE IKE Policy
	ikeKey := models.KeyLifetime(utl.IPSecPolicyKeyLifetime)
	bodyIKE := &models.IKEPolicyCreate{
		DhGroup:      &utl.IKEPolicyDhGroup,
		Encryption:   &utl.IKEPolicyEncryption,
		KeyLifetime:  &ikeKey,
		Name:         &utl.IKEPolicyName,
		PresharedKey: &utl.IKEPolicyPresharedKey,
		Version:      &utl.IKEPolicyVersion,
	}
	createIkeResp, err := vpnPolicyClient.CreateIKEPolicy(bodyIKE)
	require.Nil(t, err)
	ikePolicyID := *createIkeResp.ID
	utl.TestMessage("CREATE IKE Policy", ikePolicyID, *createIkeResp)

	// DELETE IKE Policy
	defer func() {
		err = vpnPolicyClient.DeleteIKEPolicy(ikePolicyID)
		require.Nil(t, err)
		utl.TestMessage("DELETE IKE Policy", ikePolicyID, nil)
	}()

	// CREATE IPSec Policy
	ipSecKey := models.KeyLifetime(utl.IPSecPolicyKeyLifetime)
	bodyIPSEC := &models.IPSecPolicyCreate{
		Authentication: models.IPSECPolicyAuthenticationHmacDashSha1Dash96,
		DhGroup:        &utl.IPSecPolicyDhGroup,
		Encryption:     &utl.IPSecPolicyEncryption,
		KeyLifetime:    &ipSecKey,
		Name:           &utl.IPSecPolicyName,
		Pfs:            &utl.IPSecPolicyPfs,
	}
	createIpsecResp, err := vpnPolicyClient.CreateIPSecPolicy(bodyIPSEC)
	require.Nil(t, err)
	ipsecPolicyID := *createIpsecResp.ID
	utl.TestMessage("CREATE IPSec Policy", ipsecPolicyID, *createIpsecResp)

	// DELETE IPSec Policy
	defer func() {
		err = vpnPolicyClient.DeleteIPSecPolicy(ipsecPolicyID)
		require.Nil(t, err)
		utl.TestMessage("DELETE IPSec Policy", ipsecPolicyID, nil)
	}()

	// CREATE VPN
	peerGatewayAddress := models.PeerGatewayAddress(utl.VpnPeerGatewayAddress)
	body := &models.VPNConnectionCreate{
		IkePolicy:          &ikePolicyID,
		IPSecPolicy:        &ipsecPolicyID,
		Mode:               &utl.VpnMode,
		Name:               &utl.VpnName,
		Networks:           []string{utl.VpnNetworkID},
		PeerGatewayAddress: &peerGatewayAddress,
		PeerSubnets:        []string{utl.VpnPeerSubnet},
	}
	createVpnResp, err := vpnClient.Create(body)
	require.Nil(t, err)
	utl.WaitForJobState(t, jobClient, *createVpnResp.JobRef.ID)
	vpnID := *createVpnResp.ID
	utl.TestMessage("CREATE VPN", vpnID, *createVpnResp)

	// DELETE VPN
	defer func() {
		resp, err := vpnClient.Delete(vpnID)
		require.Nil(t, err)
		utl.WaitForJobState(t, jobClient, *resp.ID)
		utl.TestMessage("DELETE VPN", vpnID, *resp)
	}()

	// GET IKE Policy
	getIkePolicy, err := vpnPolicyClient.GetIKEPolicy(ikePolicyID)
	require.Nil(t, err)
	utl.TestMessage("GET IKE Policy", ikePolicyID, *getIkePolicy)
	// verify variables match
	require.Equal(t, *getIkePolicy.DhGroup, utl.IKEPolicyDhGroup)
	require.Equal(t, *getIkePolicy.Encryption, utl.IKEPolicyEncryption)
	require.Equal(t, *getIkePolicy.KeyLifetime, models.KeyLifetime(utl.IKEPolicyKeyLifetime))
	require.Equal(t, *getIkePolicy.Name, utl.IKEPolicyName)
	require.Equal(t, *getIkePolicy.Version, utl.IKEPolicyVersion)

	// GET IPSec Policy
	getIpsecPolicy, err := vpnPolicyClient.GetIPSecPolicy(ipsecPolicyID)
	require.Nil(t, err)
	utl.TestMessage("GET IPSec Policy", ipsecPolicyID, *getIpsecPolicy)
	// verify variables match
	require.Equal(t, *getIpsecPolicy.Authentication, models.IPSECPolicyAuthenticationHmacDashSha1Dash96)
	require.Equal(t, *getIpsecPolicy.DhGroup, utl.IPSecPolicyDhGroup)
	require.Equal(t, *getIpsecPolicy.Encryption, utl.IPSecPolicyEncryption)
	require.Equal(t, *getIpsecPolicy.KeyLifetime, models.KeyLifetime(utl.IPSecPolicyKeyLifetime))
	require.Equal(t, *getIpsecPolicy.Name, utl.IPSecPolicyName)
	require.Equal(t, *getIpsecPolicy.Pfs, utl.IPSecPolicyPfs)

	// GET VPN
	getVpn, err := vpnClient.Get(vpnID)
	require.Nil(t, err)
	utl.TestMessage("GET VPN", vpnID, *getVpn)
	// verify variables match
	require.Equal(t, *getVpn.IkePolicy.ID, *getIkePolicy.ID)
	require.Equal(t, *getVpn.IPSecPolicy.ID, *getIpsecPolicy.ID)
	require.Equal(t, *getVpn.Mode, utl.VpnMode)
	require.Equal(t, *getVpn.Name, utl.VpnName)
	//require.Equal(t, getVpn.NetworkIDs, []string{utl.VpnNetworkID})
	require.Equal(t, *getVpn.PeerGatewayAddress, models.PeerGatewayAddress(utl.VpnPeerGatewayAddress))
	require.Equal(t, getVpn.PeerSubnets, []string{utl.VpnPeerSubnet})

	// GET ALL IKE Policy
	getAllIKEResp, err := vpnPolicyClient.GetAllIKEPolicies()
	require.Nil(t, err)
	utl.TestMessage("GET ALL IKE Policies", "", *getAllIKEResp)

	// GET ALL IPSec Policy
	getAllIPSecResp, err := vpnPolicyClient.GetAllIPSecPolicies()
	require.Nil(t, err)
	utl.TestMessage("GET ALL IPSec Policies", "", *getAllIPSecResp)

	// GET ALL VPN
	getAllResp, err := vpnClient.GetAll()
	require.Nil(t, err)
	utl.TestMessage("GET ALL VPNs", "", *getAllResp)
}
