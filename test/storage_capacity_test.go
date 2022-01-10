package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/stretchr/testify/require"
)

func TestStorageCapacityGet(t *testing.T) {
	testStorageCapacityPreCheck(t)

	// create session and client
	session := getSession(t)
	storageCapacityClient := client.NewIBMPIStorageCapacityClient(context.Background(), session, cloudInstanceID)

	getAllPools, err := storageCapacityClient.GetAllStoragePoolsCapacity()
	require.Nil(t, err)
	testMessage("GET ALL Storage Capacity", "", getAllPools)

	getRespSP, err := storageCapacityClient.GetStoragePoolCapacity(storagePool)
	require.Nil(t, err)
	testMessage("GET Storage Pool", storagePool, *getRespSP)

	getAllCap, err := storageCapacityClient.GetAllStorageTypesCapacity()
	require.Nil(t, err)
	testMessage("GET All Storage Types", "", *getAllCap)

	getType, err := storageCapacityClient.GetStorageTypeCapacity(storageType)
	require.Nil(t, err)
	testMessage("GET Storage Type", storageType, *getType)
}
