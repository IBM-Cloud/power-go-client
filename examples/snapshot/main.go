package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	//os.Setenv("IBMCLOUD_POWER_API_ENDPOINT", region+".power-iaas.test.cloud.ibm.com")

	piID := " < POWER INSTANCE ID > "
	instance_id := " < INSTANCE ID > "
	snap_name := " < SNAPSHOT NAME > "
	description := " < DESCRIPTION > "

	snapshotBody := &models.SnapshotCreate{Name: &snap_name, Description: description}

	session, err := ps.New(token, region, true, accountID, zone)
	if err != nil {
		log.Fatal(err)
	}

	powerClient := v.NewIBMPIInstanceClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

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
