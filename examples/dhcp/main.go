package main

import (
	"fmt"
	"log"
	"os"

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

	session, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIDhcpClient(session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}

	dhcpServer, err := powerClient.Create(cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", dhcpServer)

	dhcpId := *dhcpServer.ID
	getResp, err := powerClient.Get(dhcpId, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getAllResp, err := powerClient.GetAll(cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", getAllResp)

	_, err = powerClient.Delete(dhcpId, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
}
