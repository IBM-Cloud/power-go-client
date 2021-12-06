package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/joho/godotenv"
)

const (
	JOBCOMPLETED = "completed"
	JOBFAILED    = "failed"
)

func main() {

	// set to staging or production
	environment := "staging"
	if environment == "staging" {
		godotenv.Load("../../.env.staging")
	} else {
		godotenv.Load("../../.env.production")
	}

	// load cloud instance id
	cloudInstanceId := os.Getenv("CLOUD_INSTANCE_ID")
	if cloudInstanceId == "" {
		log.Fatal(fmt.Errorf("CLOUD_INSTANCE_ID is empty: define in .env.%v", environment))
	}

	// VPN inputs
	name := " < NAME OF THE VPN CONNECTION > "
	ikePolicyName := " < NAME OF THE IKE POLICY > "
	ipsecPolicyName := " < NAME OF THE IPSEC POLICY > "
	networks := []string{" < NAME OF THE NETWORK > "}

	session, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}
	vpnClient := v.NewIBMPIVpnConnectionClient(context.Background(), session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	vpnPolicyClient := v.NewIBMPIVpnPolicyClient(context.Background(), session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	jobClient := v.NewIBMPIJobClient(context.Background(), session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}

	// Create and Get IKE Policy
	var dhGroup, version int64 = 1, 1
	encryption := "3des-cbc"
	presharedKey := "sample"
	bodyike := &models.IKEPolicyCreate{
		DhGroup:      &dhGroup,
		Encryption:   &encryption,
		KeyLifetime:  models.KeyLifetime(180),
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
		Authentication: models.IPSECPolicyAuthenticationHmacMd596,
		DhGroup:        &dhGroup,
		Encryption:     &encryption,
		KeyLifetime:    models.KeyLifetime(180),
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
	body := &models.VPNConnectionCreate{
		IkePolicy:          ikePolicy.ID,
		IPSecPolicy:        ipsecPolicy.ID,
		Mode:               &mode,
		Name:               &name,
		Networks:           networks,
		PeerGatewayAddress: "1.1.1.1",
		PeerSubnets:        []string{"128.0.111.0/24"},
	}
	createResp, err := vpnClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createResp)

	waitForJobState(jobClient, *createResp.JobRef.ID, cloudInstanceId, 2000)
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
	waitForJobState(jobClient, *resp.ID, cloudInstanceId, 2000)
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
