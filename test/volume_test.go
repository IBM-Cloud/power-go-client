package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestVolumeCreateGetDelete(t *testing.T) {
	// check for testing variables
	testVolumePreCheck(t)

	// create session and client
	session := getSession(t)
	volumeClient := client.NewIBMPIVolumeClient(context.Background(), session, cloudInstanceID)

	// CREATE
	volumeBody := &models.CreateDataVolume{
		Name:      &volumeName,
		Size:      &volumeSize,
		DiskType:  volumeType,
		Shareable: &volumeShareable,
	}
	createRespOk, err := volumeClient.CreateVolume(volumeBody)
	if err != nil {
		t.Fatal(err)
	}
	volumeID := *createRespOk.VolumeID
	testMessage("CREATE Volume", volumeID, *createRespOk)

	// GET
	getResp, err := volumeClient.Get(volumeID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET Volume", volumeID, *getResp)

	// DELETE
	err = volumeClient.DeleteVolume(volumeID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("DELETE Volume", volumeID, nil)
}
func TestVolumeGetAll(t *testing.T) {
	// check for testing variables
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	volumeClient := client.NewIBMPIVolumeClient(context.Background(), session, cloudInstanceID)

	getallResp, err := volumeClient.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET ALL Volumes", "", *getallResp)
}
