package main

import (
	"context"
	"fmt"
	"log"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"
)

const (
	JOBCOMPLETED = "completed"
	JOBFAILED    = "failed"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".power-iaas.test.cloud.ibm.com"

	// VPN inputs
	name := " < NAME OF THE VPN CONNECTION > "
	piID := " < POWER INSTANCE ID > "
	ikePolicyName := " < NAME OF THE IKE POLICY > "
	ipsecPolicyName := " < NAME OF THE IPSEC POLICY > "
	networks := []string{" < NAME OF THE NETWORK > "}

	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: token,
	}
	// authenticator := &core.IamAuthenticator{
	// 	ApiKey: "< API KEY >",
	// 	// Uncomment for test environment
	// 	URL: "https://iam.test.cloud.ibm.com",
	// }
	// Create the session
	options := &ps.IBMPIOptions{
		Authenticator: authenticator,
		UserAccount:   accountID,
		Zone:          zone,
		URL:           url,
		Debug:         true,
	}
	session, err := ps.NewIBMPISession(options)
	if err != nil {
		log.Fatal(err)
	}
	vpnClient := v.NewIBMPIVpnConnectionClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}
	vpnPolicyClient := v.NewIBMPIVpnPolicyClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}
	jobClient := v.NewIBMPIJobClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	// Create and Get IKE Policy
	var dhGroup, version int64 = 1, 1
	encryption := "3des-cbc"
	presharedKey := "sample"
	keyLifeTime := models.KeyLifetime(180)
	bodyike := &models.IKEPolicyCreate{
		DhGroup:      &dhGroup,
		Encryption:   &encryption,
		KeyLifetime:  &keyLifeTime,
		Name:         &ikePolicyName,
		PresharedKey: &presharedKey,
		Version:      &version,
	}
	ikePolicy, err := vpnPolicyClient.CreateIKEPolicy(bodyike)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", ikePolicy)

	ikePolicyRead, err := vpnPolicyClient.GetIKEPolicy(*ikePolicy.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v\n", ikePolicyRead)

	// Create and Get IPSec Policy
	pfs := false
	bodyipsec := &models.IPSecPolicyCreate{
		Authentication: models.IPSECPolicyAuthenticationHmacDashShaDash256Dash128,
		DhGroup:        &dhGroup,
		Encryption:     &encryption,
		KeyLifetime:    &keyLifeTime,
		Name:           &ipsecPolicyName,
		Pfs:            &pfs,
	}
	ipsecPolicy, err := vpnPolicyClient.CreateIPSecPolicy(bodyipsec)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", ipsecPolicy)

	ipsecPolicyRead, err := vpnPolicyClient.GetIPSecPolicy(*ipsecPolicy.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v\n", ipsecPolicyRead)

	// Create and Get VPN Connection
	mode := "policy"
	pga := models.PeerGatewayAddress("1.1.1.1")
	body := &models.VPNConnectionCreate{
		IkePolicy:          ikePolicy.ID,
		IPSecPolicy:        ipsecPolicy.ID,
		Mode:               &mode,
		Name:               &name,
		Networks:           networks,
		PeerGatewayAddress: &pga,
		PeerSubnets:        []string{"128.0.111.0/24"},
	}
	createResp, err := vpnClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createResp)

	waitForJobState(jobClient, *createResp.JobRef.ID, piID, 2000)
	if err != nil {
		log.Fatal(err)
	}
	getResp, err := vpnClient.Get(*createResp.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := vpnClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getallResp)
	getallIKEResp, err := vpnPolicyClient.GetAllIKEPolicies()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getallIKEResp)
	getallIPSecResp, err := vpnPolicyClient.GetAllIPSecPolicies()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getallIPSecResp)

	resp, err := vpnClient.Delete(*createResp.ID)
	if err != nil {
		log.Fatal(err)
	}
	waitForJobState(jobClient, *resp.ID, piID, 2000)
	if err != nil {
		log.Fatal(err)
	}

	err = vpnPolicyClient.DeleteIKEPolicy(*ikePolicy.ID)
	if err != nil {
		log.Fatal(err)
	}
	err = vpnPolicyClient.DeleteIPSecPolicy(*ipsecPolicy.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func waitForJobState(jobClient *v.IBMPIJobClient, jobId, cloudinstanceid string, interval time.Duration) error {
	var status string

	for status != JOBCOMPLETED && status != JOBFAILED {
		job, err := jobClient.Get(jobId)
		if err != nil {
			return err
		}
		if job == nil || job.Status == nil {
			return fmt.Errorf("cannot find job status for job id %s with cloud instance %s", jobId, cloudinstanceid)
		}
		time.Sleep(interval)
		status = *job.Status.State
	}
	return nil
}
