package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".power-iaas.test.cloud.ibm.com"

	// pool inputs
	piID := " < POWER INSTANCE ID > "
	storageType := "tier1"
	storagePool := "Tier1-Flash-2"

	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: token,
	}
	// authenticator := &core.IamAuthenticator{
	// 	ApiKey: "< API KEY >",
	// 	// Uncomment for test environment
	// 	URL: "https://iam.test.cloud.ibm.com",
	// }
	// Create the session
	options := &ps.IBMPIOptions{
		Authenticator: authenticator,
		UserAccount:   accountID,
		Zone:          zone,
		URL:           url,
		Debug:         true,
	}
	session, err := ps.NewIBMPISession(options)
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
