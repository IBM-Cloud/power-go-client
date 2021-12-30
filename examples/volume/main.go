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
	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	//os.Setenv("IBMCLOUD_POWER_API_ENDPOINT", region + ".power-iaas.test.cloud.ibm.com")

	// volume inputs
	piID := " < POWER INSTANCE ID > "
	name := " < NAME OF THE volume > "
	size := 20.0
	vtype := "tier1"
	sharable := true

	session, err := ps.New(token, region, true, accountID, zone)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIVolumeClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	body := &models.CreateDataVolume{
		Name:      &name,
		Size:      &size,
		DiskType:  vtype,
		Shareable: &sharable,
	}
	createRespOk, err := powerClient.CreateVolume(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	volumeID := *createRespOk.VolumeID
	getResp, err := powerClient.Get(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getallResp)

	err = powerClient.DeleteVolume(volumeID)
	if err != nil {
		log.Fatal(err)
	}
}
