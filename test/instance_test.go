package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestInstanceCreateGetDelete(t *testing.T) {
	testInstancePreCheck(t)

	// create session and client
	session := getSession(t)
	instanceClient := client.NewIBMPIInstanceClient(context.Background(), session, cloudInstanceID)

	// CREATE
	body := &models.PVMInstanceCreate{
		ImageID:     &instanceImageID,
		KeyPairName: instanceSSHKey,
		NetworkIDs:  []string{instanceNetworkID},
		ServerName:  &instanceName,
		//VolumeIds:   []string{instanceVolumeID},
		Memory:     &instanceMemory,
		Processors: &instanceProcessors,
		ProcType:   &instanceProcType,
		SysType:    instanceSysType,
	}
	createResp, err := instanceClient.Create(body)
	require.Nil(t, err)
	instanceID := *(*createResp)[0].PvmInstanceID

	// check that instance is successfully created
	createErr := fmt.Sprintf("Create Instance %s has failed. Still in Build state after %d minutes", instanceID, ((timeoutAttempts * jobTimeout) / 60000000000))
	checkAttempts := 0
	getResp := &models.PVMInstance{}
	for checkAttempts < timeoutAttempts {
		getResp, err = instanceClient.Get(instanceID)
		require.Nil(t, err)
		if *getResp.Status == "ERROR" {
			t.Fatalf("Newly Created Instance: %s has Status: ERROR", instanceID)
		}
		if *getResp.Status != "BUILD" {
			break
		}
		time.Sleep(time.Duration(jobTimeout))
		checkAttempts += 1
	}
	require.Less(t, checkAttempts, timeoutAttempts, createErr)
	testMessage("CREATE Instance", instanceID, *createResp)

	// GET
	getResp, err = instanceClient.Get(instanceID)
	require.Nil(t, err)
	testMessage("GET Instance", instanceID, *getResp)
	// verify variables match
	require.Equal(t, *getResp.ImageID, instanceImageID)
	require.Equal(t, getResp.NetworkIDs, []string{instanceNetworkID})
	require.Equal(t, *getResp.ServerName, instanceName)
	require.Equal(t, getResp.SysType, instanceSysType)

	// these are not "set" when calling get immediately after  reation for some reason
	//require.Equal(t, *getResp.Memory, instanceMemory)
	//require.Equal(t, *getResp.Processors, instanceProcessors)
	//require.Equal(t, *getResp.ProcType, instanceProcType)

	// DELETE
	err = instanceClient.Delete(instanceID)
	require.Nil(t, err)
	testMessage("DELETE Instance", instanceID, nil)
}
func TestInstanceGetAll(t *testing.T) {
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	instanceClient := client.NewIBMPIInstanceClient(context.Background(), session, cloudInstanceID)

	// GET ALL
	getAllResp, err := instanceClient.GetAll()
	require.Nil(t, err)
	testMessage("GET All Instances", "", *getAllResp)
}
