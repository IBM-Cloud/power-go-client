package test

import (
	"context"
	"log"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestVPNCreateGetDelete(t *testing.T) {
	// check for testing variables
	testVPNPreCheck(t)

	// create session and clients
	session := getSession(t)
	vpnClient := client.NewIBMPIVpnConnectionClient(context.Background(), session, cloudInstanceID)
	vpnPolicyClient := client.NewIBMPIVpnPolicyClient(context.Background(), session, cloudInstanceID)
	jobClient := client.NewIBMPIJobClient(context.Background(), session, cloudInstanceID)

	// CREATE
	// IKE Policy
	bodyIKE := &models.IKEPolicyCreate{
		DhGroup:      &ikePolicyDhGroup,
		Encryption:   &ikePolicyEncryption,
		KeyLifetime:  models.KeyLifetime(ikePolicyKeyLifetime),
		Name:         &ikePolicyName,
		PresharedKey: &ikePolicyPresharedKey,
		Version:      &ikePolicyVersion,
	}

	createIkeResp, err := vpnPolicyClient.CreateIKEPolicy(bodyIKE)
	if err != nil {
		log.Fatal(err)
	}
	ikePolicyID := *createIkeResp.ID
	testMessage("CREATE IKE Policy", ikePolicyID, *createIkeResp)

	// IPSec Policy
	bodyIPSEC := &models.IPSecPolicyCreate{
		Authentication: models.IPSECPolicyAuthenticationHmacMd596,
		DhGroup:        &ipSecPolicyDhGroup,
		Encryption:     &ipSecPolicyEncryption,
		KeyLifetime:    models.KeyLifetime(ipSecPolicyKeyLifetime),
		Name:           &ipSecPolicyName,
		Pfs:            &ipSecPolicyPfs,
	}
	createIpsecResp, err := vpnPolicyClient.CreateIPSecPolicy(bodyIPSEC)
	if err != nil {
		log.Fatal(err)
	}
	ipsecPolicyID := *createIpsecResp.ID
	testMessage("CREATE IPSec Policy", ipsecPolicyID, *createIpsecResp)

	// VPN
	body := &models.VPNConnectionCreate{
		IkePolicy:          &ikePolicyID,
		IPSecPolicy:        &ipsecPolicyID,
		Mode:               &vpnMode,
		Name:               &vpnName,
		Networks:           []string{vpnNetworkID},
		PeerGatewayAddress: models.PeerGatewayAddress(vpnPeerGatewayAddress),
		PeerSubnets:        []string{vpnPeerSubnet},
	}
	createVpnResp, err := vpnClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	waitForJobState(jobClient, *createVpnResp.JobRef.ID, jobTimeout)
	if err != nil {
		log.Fatal(err)
	}
	vpnID := *createVpnResp.ID
	testMessage("CREATE VPN", vpnID, *createVpnResp)

	// GET
	// IKE Policy
	ikePolicyRead, err := vpnPolicyClient.GetIKEPolicy(ikePolicyID)
	if err != nil {
		log.Fatal(err)
	}
	testMessage("GET IKE Policy", ikePolicyID, *ikePolicyRead)

	// IPSec Policy
	ipsecPolicyRead, err := vpnPolicyClient.GetIPSecPolicy(ipsecPolicyID)
	if err != nil {
		log.Fatal(err)
	}
	testMessage("GET IPSec Policy", ipsecPolicyID, *ipsecPolicyRead)

	// VPN
	getResp, err := vpnClient.Get(vpnID)
	if err != nil {
		log.Fatal(err)
	}
	testMessage("GET VPN", vpnID, *getResp)

	// DELETE
	// VPN
	resp, err := vpnClient.Delete(vpnID)
	if err != nil {
		log.Fatal(err)
	}
	waitForJobState(jobClient, *resp.ID, jobTimeout)
	if err != nil {
		log.Fatal(err)
	}
	testMessage("DELETE VPN", vpnID, *resp)

	// IKE Policy
	err = vpnPolicyClient.DeleteIKEPolicy(ikePolicyID)
	if err != nil {
		log.Fatal(err)
	}
	testMessage("DELETE IKE Policy", ikePolicyID, nil)

	// IPSec Policy
	err = vpnPolicyClient.DeleteIPSecPolicy(ipsecPolicyID)
	if err != nil {
		log.Fatal(err)
	}
	testMessage("DELETE IPSec Policy", ipsecPolicyID, nil)
}

func TestVPNGetAll(t *testing.T) {
	// check for testing variables
	testPreCheck(t)

	// create session and clients
	session := getSession(t)
	vpnClient := client.NewIBMPIVpnConnectionClient(context.Background(), session, cloudInstanceID)
	vpnPolicyClient := client.NewIBMPIVpnPolicyClient(context.Background(), session, cloudInstanceID)

	// VPN
	getAllResp, err := vpnClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	testMessage("GET ALL VPNs", "", *getAllResp)

	// IKE Policy
	getAllIKEResp, err := vpnPolicyClient.GetAllIKEPolicies()
	if err != nil {
		log.Fatal(err)
	}
	testMessage("GET ALL IKE Policies", "", *getAllIKEResp)

	// IPSec Policy
	getAllIPSecResp, err := vpnPolicyClient.GetAllIPSecPolicies()
	if err != nil {
		log.Fatal(err)
	}
	testMessage("GET ALL IPSec Policies", "", *getAllIPSecResp)
}
