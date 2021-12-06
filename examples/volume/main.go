package main

import (
	"fmt"
	"log"
	"os"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/joho/godotenv"
)

func main() {

	// set to staging or production
	environment := "staging"
	if environment == "staging" {
		godotenv.Load("../../.env.staging")
	} else {
		godotenv.Load("../../.env.production")
	}

	// load cloud instance id
	cloudInstanceId := os.Getenv("CLOUD_INSTANCE_ID")
	if cloudInstanceId == "" {
		log.Fatal(fmt.Errorf("CLOUD_INSTANCE_ID is empty: define in .env.%v", environment))
	}

	// volume inputs
	name := " < NAME OF THE volume > "
	size := 20.0
	vtype := "tier1"
	sharable := true
	timeout := time.Duration(9000000000000000000)

	session, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIVolumeClient(session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}

	createRespOk, err := powerClient.Create(name, size, vtype, sharable, cloudInstanceId, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createRespOk)

	volumeID := *createRespOk.VolumeID
	getResp, err := powerClient.Get(volumeID, cloudInstanceId, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := powerClient.GetAll(volumeID, cloudInstanceId, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getallResp)

	err = powerClient.Delete(volumeID, cloudInstanceId, timeout)
	if err != nil {
		log.Fatal(err)
	}
}
