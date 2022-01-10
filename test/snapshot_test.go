package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestSnapshotCreateGetDelete(t *testing.T) {
	testSnapshotPreCheck(t)

	// create session and client
	session := getSession(t)
	instanceClient := client.NewIBMPIInstanceClient(context.Background(), session, cloudInstanceID)
	snapshotClient := client.NewIBMPISnapshotClient(context.Background(), session, cloudInstanceID)

	// CREATE
	body := &models.SnapshotCreate{
		Name:        &snapshotName,
		Description: snapshotDescription,
	}
	createResp, err := instanceClient.CreatePvmSnapShot(snapshotInstanceID, body)
	require.Nil(t, err)
	snapshotID := createResp.SnapshotID
	testMessage("CREATE Snapshot", *snapshotID, *createResp)

	// GET
	getResp, err := snapshotClient.Get(*snapshotID)
	require.Nil(t, err)
	testMessage("GET Snapshot", *snapshotID, *getResp)

	// DELETE
	err = snapshotClient.Delete(sshKeyName)
	require.Nil(t, err)
	testMessage("DELETE Snapshot", *snapshotID, nil)
}
func TestSnapshotGetAll(t *testing.T) {
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	snapshotClient := client.NewIBMPISnapshotClient(context.Background(), session, cloudInstanceID)

	// GET ALL
	getAllResp, err := snapshotClient.GetAll()
	require.Nil(t, err)
	testMessage("GET All Snapshots", "", *getAllResp)
}
