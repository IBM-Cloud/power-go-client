package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestCloudConnectionCreateGetDelete(t *testing.T) {
	// check for testing variables
	testCloudConnectionPreCheck(t)

	// create session and clients
	session := getSession(t)
	cloudConnectionClient := client.NewIBMPICloudConnectionClient(context.Background(), session, cloudInstanceID)
	jobClient := client.NewIBMPIJobClient(context.Background(), session, cloudInstanceID)

	// CREATE
	body := &models.CloudConnectionCreate{
		Name:  &cloudConnectionName,
		Speed: &cloudConnectionSpeed,
	}
	createRespOk, createRespAccepted, err := cloudConnectionClient.Create(body)
	if err != nil {
		t.Fatal(err)
	}

	// get cloud connection id from recently created cloud connection
	var cloudConnectionID, jobID string
	if createRespOk != nil {
		cloudConnectionID = *createRespOk.CloudConnectionID
	} else {
		cloudConnectionID = *createRespAccepted.CloudConnectionID
		jobID = *createRespAccepted.JobRef.ID
		waitForJobState(jobClient, jobID, jobTimeout)
		if err != nil {
			t.Fatal(err)
		}
	}
	testMessage("CREATE Cloud Connection", cloudConnectionID, *createRespAccepted)

	// GET
	getResp, err := cloudConnectionClient.Get(cloudConnectionID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET Cloud Connection", cloudConnectionID, *getResp)

	// DELETE
	delResp, err := cloudConnectionClient.Delete(cloudConnectionID)
	if err != nil {
		t.Fatal(err)
	}
	if delResp != nil {
		jobID = *delResp.ID
		waitForJobState(jobClient, jobID, jobTimeout)
		if err != nil {
			t.Fatal(err)
		}
	}
	testMessage("DELETE Cloud Connection", cloudConnectionID, *delResp)
}
func TestCloudConnectionGetAll(t *testing.T) {
	// check for testing variables
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	cloudConnectionClient := client.NewIBMPICloudConnectionClient(context.Background(), session, cloudInstanceID)

	// GET ALL Cloud Connections
	getAllResp, err := cloudConnectionClient.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET All Cloud Connections", "", *getAllResp)
}
