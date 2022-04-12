package instance_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	utl "internal/testutils"

	client "github.com/IBM-Cloud/power-go-client/clients"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestInstance(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and client
	utl.InstancePreCheck(t)
	session := utl.TestSession(t)
	instanceClient := client.NewIBMPIInstanceClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE
	body := &models.PVMInstanceCreate{
		ImageID:     &utl.InstanceImageID,
		KeyPairName: utl.InstanceSSHKey,
		NetworkIDs:  []string{utl.InstanceNetworkID},
		ServerName:  &utl.InstanceName,
		//VolumeIds:   []string{InstanceVolumeID},
		Memory:      &utl.InstanceMemory,
		Processors:  &utl.InstanceProcessors,
		ProcType:    &utl.InstanceProcType,
		SysType:     utl.InstanceSysType,
		StorageType: utl.InstanceStorageType,
	}
	createResp, err := instanceClient.Create(body)
	require.Nil(t, err)
	InstanceID := *(*createResp)[0].PvmInstanceID

	// DELETE
	defer func() {
		err = instanceClient.Delete(InstanceID)
		require.Nil(t, err)
		utl.TestMessage("DELETE Instance", InstanceID, nil)
	}()

	// check that instance is successfully created
	createErr := fmt.Sprintf("Create Instance %s has failed. Still in Build state after %d minutes", InstanceID, ((utl.TimeoutAttempts * utl.JobTimeout) / 60000000000))
	checkAttempts := 0
	getResp := &models.PVMInstance{}
	for checkAttempts < utl.TimeoutAttempts {
		getResp, err = instanceClient.Get(InstanceID)
		require.Nil(t, err)
		if *getResp.Status == "ERROR" {
			t.Fatalf("Newly Created Instance: %s has Status: ERROR", InstanceID)
		}
		if *getResp.Status != "BUILD" {
			break
		}
		time.Sleep(time.Duration(utl.JobTimeout))
		checkAttempts += 1
	}
	require.Less(t, checkAttempts, utl.TimeoutAttempts, createErr)
	utl.TestMessage("CREATE Instance", InstanceID, *createResp)

	// GET
	getResp, err = instanceClient.Get(InstanceID)
	require.Nil(t, err)
	utl.TestMessage("GET Instance", InstanceID, *getResp)
	// verify variables match
	//require.Equal(t, *getResp.ImageID, utl.InstanceImageID)
	require.Equal(t, getResp.NetworkIDs, []string{utl.InstanceNetworkID})
	require.Equal(t, *getResp.ServerName, utl.InstanceName)
	require.Equal(t, getResp.SysType, utl.InstanceSysType)

	// these are not "set" when calling get immediately after  reation for some reason
	//require.Equal(t, *getResp.Memory, InstanceMemory)
	//require.Equal(t, *getResp.Processors, InstanceProcessors)
	//require.Equal(t, *getResp.ProcType, InstanceProcType)

	// GET ALL
	getAllResp, err := instanceClient.GetAll()
	require.Nil(t, err)
	utl.TestMessage("GET All Instances", "", *getAllResp)
}
