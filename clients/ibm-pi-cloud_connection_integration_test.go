package instance_test

import (
	"context"
	"testing"

	utl "internal/testutils"

	client "github.com/IBM-Cloud/power-go-client/clients"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestCloudConnection(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and clients
	utl.CloudConnectionPreCheck(t)
	session := utl.TestSession(t)
	cloudConnectionClient := client.NewIBMPICloudConnectionClient(context.Background(), session, utl.CloudInstanceID)
	jobClient := client.NewIBMPIJobClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE
	body := &models.CloudConnectionCreate{
		Name:  &utl.CloudConnectionName,
		Speed: &utl.CloudConnectionSpeed,
	}
	createRespOk, createRespAccepted, err := cloudConnectionClient.Create(body)
	require.Nil(t, err)

	// get cloud connection id from recently created cloud connection
	var CloudConnectionID string
	if createRespOk != nil {
		CloudConnectionID = *createRespOk.CloudConnectionID
	} else {
		CloudConnectionID = *createRespAccepted.CloudConnectionID
		utl.WaitForJobState(t, jobClient, *createRespAccepted.JobRef.ID)
	}
	utl.TestMessage("CREATE Cloud Connection", CloudConnectionID, *createRespAccepted)

	// DELETE
	defer func() {
		delResp, err := cloudConnectionClient.Delete(CloudConnectionID)
		require.Nil(t, err)
		utl.WaitForJobState(t, jobClient, *delResp.ID)
		utl.TestMessage("DELETE Cloud Connection", CloudConnectionID, *delResp)
	}()

	// GET
	getResp, err := cloudConnectionClient.Get(CloudConnectionID)
	require.Nil(t, err)
	utl.TestMessage("GET Cloud Connection", CloudConnectionID, *getResp)
	// verify variables match
	require.Equal(t, *getResp.Name, utl.CloudConnectionName)
	require.Equal(t, *getResp.Speed, utl.CloudConnectionSpeed)

	// GET ALL
	getAllResp, err := cloudConnectionClient.GetAll()
	require.Nil(t, err)
	utl.TestMessage("GET All Cloud Connections", "", *getAllResp)
}
