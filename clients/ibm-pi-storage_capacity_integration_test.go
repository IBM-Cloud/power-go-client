package instance_test

import (
	"context"
	"testing"

	utl "internal/testutils"

	client "github.com/IBM-Cloud/power-go-client/clients"
	"github.com/stretchr/testify/require"
)

func TestStorageCapacity(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and client
	utl.StorageCapacityPreCheck(t)
	session := utl.TestSession(t)
	storageCapacityClient := client.NewIBMPIStorageCapacityClient(context.Background(), session, utl.CloudInstanceID)

	getAllPools, err := storageCapacityClient.GetAllStoragePoolsCapacity()
	require.Nil(t, err)
	utl.TestMessage("GET ALL Storage Capacity", "", getAllPools)

	getRespSP, err := storageCapacityClient.GetStoragePoolCapacity(utl.StoragePool)
	require.Nil(t, err)
	utl.TestMessage("GET Storage Pool", utl.StoragePool, *getRespSP)

	getAllCap, err := storageCapacityClient.GetAllStorageTypesCapacity()
	require.Nil(t, err)
	utl.TestMessage("GET All Storage Types", "", *getAllCap)

	getType, err := storageCapacityClient.GetStorageTypeCapacity(utl.StorageType)
	require.Nil(t, err)
	utl.TestMessage("GET Storage Type", utl.StorageType, *getType)
}
