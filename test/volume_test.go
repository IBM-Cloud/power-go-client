package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestVolumeCreateGetDelete(t *testing.T) {
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
	require.Nil(t, err)
	volumeID := *createRespOk.VolumeID
	testMessage("CREATE Volume", volumeID, *createRespOk)

	// GET
	getResp, err := volumeClient.Get(volumeID)
	require.Nil(t, err)
	testMessage("GET Volume", volumeID, *getResp)
	// verify variables match
	require.Equal(t, *getResp.Name, volumeName)
	require.Equal(t, *getResp.Size, volumeSize)
	require.Equal(t, getResp.DiskType, volumeType)
	require.Equal(t, *getResp.Shareable, volumeShareable)

	// DELETE
	err = volumeClient.DeleteVolume(volumeID)
	require.Nil(t, err)
	testMessage("DELETE Volume", volumeID, nil)
}
func TestVolumeGetAll(t *testing.T) {
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	volumeClient := client.NewIBMPIVolumeClient(context.Background(), session, cloudInstanceID)

	getallResp, err := volumeClient.GetAll()
	require.Nil(t, err)
	testMessage("GET ALL Volumes", "", *getallResp)
}
