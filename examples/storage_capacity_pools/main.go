package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	//os.Setenv("IBMCLOUD_POWER_API_ENDPOINT", region + ".power-iaas.test.cloud.ibm.com")

	// pool inputs
	piID := " < POWER INSTANCE ID > "
	storageType := "tier1"
	storagePool := "Tier1-Flash-2"

	session, err := ps.New(token, region, true, accountID, zone)
	if err != nil {
		log.Fatal(err)
	}

	storage_capacity_client := v.NewIBMPIStorageCapacityClient(context.Background(), session, piID)

	getAllRespSP, err := storage_capacity_client.GetAllStoragePoolsCapacity()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n", *getAllRespSP)

	getRespSP, err := storage_capacity_client.GetStoragePoolCapacity(storagePool)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getRespSP)

	getAllRespST, err := storage_capacity_client.GetAllStorageTypesCapacity()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getAllRespST)
	getRespST, err := storage_capacity_client.GetStorageTypeCapacity(storageType)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", *getRespST)
}
