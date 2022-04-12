package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".power-iaas.test.cloud.ibm.com"

	piID := " < POWER INSTANCE ID > "
	instance_id := " < INSTANCE ID > "
	snap_name := " < SNAPSHOT NAME > "
	description := " < DESCRIPTION > "

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

	powerClient := v.NewIBMPIInstanceClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	snapshotBody := &models.SnapshotCreate{Name: &snap_name, Description: description}
	createRespSnapOk, err := powerClient.CreatePvmSnapShot(instance_id, snapshotBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n", *createRespSnapOk)

	powerSnapClient := v.NewIBMPISnapshotClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	getAllSnapResp, err := powerSnapClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getAllSnapResp)

	getSnapResp, err := powerSnapClient.Get(*createRespSnapOk.SnapshotID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getSnapResp)

	err = powerSnapClient.Delete(*createRespSnapOk.SnapshotID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", err)

}
