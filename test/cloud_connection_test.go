package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestCloudConnectionCreateGetDelete(t *testing.T) {
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
	require.Nil(t, err)

	// get cloud connection id from recently created cloud connection
	var cloudConnectionID string
	if createRespOk != nil {
		cloudConnectionID = *createRespOk.CloudConnectionID
	} else {
		cloudConnectionID = *createRespAccepted.CloudConnectionID
		waitForJobState(t, jobClient, *createRespAccepted.JobRef.ID)
	}
	testMessage("CREATE Cloud Connection", cloudConnectionID, *createRespAccepted)

	// GET
	getResp, err := cloudConnectionClient.Get(cloudConnectionID)
	require.Nil(t, err)
	testMessage("GET Cloud Connection", cloudConnectionID, *getResp)
	// verify variables match
	require.Equal(t, *getResp.Name, cloudConnectionName)
	require.Equal(t, *getResp.Speed, cloudConnectionSpeed)

	// DELETE
	delResp, err := cloudConnectionClient.Delete(cloudConnectionID)
	require.Nil(t, err)
	waitForJobState(t, jobClient, *delResp.ID)
	testMessage("DELETE Cloud Connection", cloudConnectionID, *delResp)
}
func TestCloudConnectionGetAll(t *testing.T) {
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	cloudConnectionClient := client.NewIBMPICloudConnectionClient(context.Background(), session, cloudInstanceID)

	// GET ALL Cloud Connections
	getAllResp, err := cloudConnectionClient.GetAll()
	require.Nil(t, err)
	testMessage("GET All Cloud Connections", "", *getAllResp)
}
