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

	// ssh inputs
	name := " < NAME OF THE ssh > "
	piID := " < POWER INSTANCE ID > "
	ssh := " <ssh ID> "

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
