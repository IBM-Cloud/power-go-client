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

	// Cloud connection inputs
	name := " < NAME OF THE CONNECTION > "
	piID := " < POWER INSTANCE ID > "
	var speed int64 = 5000

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
	ccClient := v.NewIBMPICloudConnectionClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}
	jobClient := v.NewIBMPIJobClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	body := &models.CloudConnectionCreate{
		Name:  &name,
		Speed: &speed,
	}
	createRespOk, createRespAccepted, err := ccClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createRespOk)
	log.Printf("***************[1]****************** %+v\n\n", createRespAccepted)

	var ccId, jobId string
	if createRespOk != nil {
		ccId = *createRespOk.CloudConnectionID
	} else {
		ccId = *createRespAccepted.CloudConnectionID
		jobId = *createRespAccepted.JobRef.ID
		waitForJobState(jobClient, jobId, piID, 2000)
		if err != nil {
			log.Fatal(err)
		}
	}

	getResp, err := ccClient.Get(ccId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n\n", *getResp)

	getAllResp, err := ccClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n\n", *getAllResp)

	delResp, err := ccClient.Delete(ccId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v\n", delResp)

	if delResp != nil {
		jobId = *delResp.ID
		waitForJobState(jobClient, jobId, piID, 2000)
		if err != nil {
			log.Fatal(err)
		}
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
