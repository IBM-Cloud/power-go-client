package test

import (
	"context"
	"testing"
	"time"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestInstanceCreateGetDelete(t *testing.T) {
	// check for testing variables
	testInstancePreCheck(t)

	// create session and client
	session := getSession(t)
	instanceClient := client.NewIBMPIInstanceClient(context.Background(), session, cloudInstanceID)

	// CREATE
	body := &models.PVMInstanceCreate{
		ImageID:     &instanceImageID,
		KeyPairName: instanceSSHKey,
		NetworkIds:  []string{instanceNetworkID},
		ServerName:  &instanceName,
		//VolumeIds:   []string{instanceVolumeID},
		Memory:     &instanceMemory,
		Processors: &instanceProcessors,
		ProcType:   &instanceProcType,
		SysType:    instanceSysType,
	}
	createResp, err := instanceClient.Create(body)
	if err != nil {
		t.Fatal(err)
	}
	instanceID := *(*createResp)[0].PvmInstanceID
	testMessage("CREATE Instance", instanceID, *createResp)

	// GET
	// wait until Instance is built before deletion
	dhcpBuildComplete := false
	for dhcpBuildComplete == false {
		getResp, err := instanceClient.Get(instanceID)
		if err != nil {
			t.Fatal(err)
		}
		if *getResp.Status == "BUILD" {
			time.Sleep(time.Duration(jobTimeout))
		} else {
			dhcpBuildComplete = true
			testMessage("GET Instance", instanceID, *getResp)
		}
	}

	// DELETE
	err = instanceClient.Delete(instanceID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("DELETE Instance", instanceID, nil)
}
func TestInstanceGetAll(t *testing.T) {
	// check for testing variables
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	instanceClient := client.NewIBMPIInstanceClient(context.Background(), session, cloudInstanceID)

	// GET ALL
	getAllResp, err := instanceClient.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET All Instances", "", *getAllResp)
}
