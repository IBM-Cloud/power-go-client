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

	// Cloud connection inputs
	name := " < NAME OF THE CONNECTION > "
	var speed int64 = 5000

	session, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}
	ccClient := v.NewIBMPICloudConnectionClient(context.Background(), session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	jobClient := v.NewIBMPIJobClient(context.Background(), session, cloudInstanceId)
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
		waitForJobState(jobClient, jobId, cloudInstanceId, 2000)
		if err != nil {
			log.Fatal(err)
		}
	}

	getResp, err := ccClient.Get(ccId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n\n", *getResp)

	getAllResp, err := ccClient.GetAll(cloudInstanceId)
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
		waitForJobState(jobClient, jobId, cloudInstanceId, 2000)
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
