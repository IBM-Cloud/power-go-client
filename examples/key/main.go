package main

import (
	"context"
	"fmt"
	"log"
	"os"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
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

	// ssh inputs
	name := " < NAME OF THE ssh > "
	ssh := " <ssh ID> "

	session, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIKeyClient(context.Background(), session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}

	getAllResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[0]****************** %+v \n", *getAllResp)

	body := &models.SSHKey{
		Name:   &name,
		SSHKey: &ssh,
	}
	createRespOk, err := powerClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createRespOk)

	sshID := *createRespOk.Name
	getResp, err := powerClient.Get(sshID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	err = powerClient.Delete(sshID)
	if err != nil {
		log.Fatal(err)
	}
}
