package main

import (
	"log"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	accountID := " < ACCOUNT ID > "

	// volume inputs
	name := " < NAME OF THE volume > "
	piID := " < POWER INSTANCE ID > "
	size := 20.0
	vtype := "tier1"
	sharable := true
	timeout := time.Duration(9000000000000000000)

	session, err := ps.New(token, region, true, 9000000000000000000, accountID, region)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIVolumeClient(session, piID)
	if err != nil {
		log.Fatal(err)
	}

	createRespOk, err := powerClient.Create(name, size, vtype, sharable, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createRespOk)

	volumeID := *createRespOk.VolumeID
	getResp, err := powerClient.Get(volumeID, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := powerClient.GetAll(volumeID, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getallResp)

	err = powerClient.Delete(volumeID, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
}
