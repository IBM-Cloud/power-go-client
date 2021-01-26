package main

import (
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	accountID := " < ACCOUNT ID > "

	// ssh inputs
	name := " < NAME OF THE ssh > "
	piID := " < POWER INSTANCE ID > "
	ssh := " <ssh ID> "
	session, err := ps.New(token, region, true, 50000000000, accountID, region)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIKeyClient(session, piID)
	if err != nil {
		log.Fatal(err)
	}
	_, createRespOk, err := powerClient.Create(name, ssh, piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createRespOk)

	sshID := *createRespOk.Name
	getResp, err := powerClient.Get(sshID, piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	err = powerClient.Delete(sshID, piID)
	if err != nil {
		log.Fatal(err)
	}
}
