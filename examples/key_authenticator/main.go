package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {

	//session Inputs
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "

	// ssh inputs
	name := " < NAME OF THE ssh > "
	piID := " < POWER INSTANCE ID > "
	ssh := " <ssh ID> "

	// Create the authenticator
	authenticator := &core.IamAuthenticator{
		ApiKey: " < APIKEY > ",
	}

	// Create the session options struct
	options := &ps.PIOptions{
		Authenticator: authenticator,
		UserAccount:   accountID,
		Region:        region,
		Zone:          zone,
	}

	// Construct the session service instance
	session, err := ps.NewSession(options)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIKeyClient(context.Background(), session, piID)
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
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

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
