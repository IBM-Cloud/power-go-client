package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestStorageCapacityGet(t *testing.T) {
	// check for testing variables
	testStorageCapacityPreCheck(t)

	// create session and client
	session := getSession(t)
	storageCapacityClient := client.NewIBMPIStorageCapacityClient(context.Background(), session, cloudInstanceID)

	getAllPools, err := storageCapacityClient.GetAllStoragePoolsCapacity()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET ALL Storage Capacity", "", getAllPools)

	getRespSP, err := storageCapacityClient.GetStoragePoolCapacity(storagePool)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET Storage Pool", storagePool, *getRespSP)

	getAllCap, err := storageCapacityClient.GetAllStorageTypesCapacity()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET All Storage Types", "", *getAllCap)

	getType, err := storageCapacityClient.GetStorageTypeCapacity(storageType)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET Storage Type", storageType, *getType)
}
