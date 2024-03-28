package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {
	//session Inputs
	//  < IAM TOKEN >
	// token := ""
	apiKey := ""
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".power-iaas.test.cloud.ibm.com"

	// dr location inputs
	piID := " < POWER INSTANCE ID > "

	// authenticator := &core.BearerTokenAuthenticator{
	// 	BearerToken: token,
	// }
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		// Uncomment for test environment
		URL: "https://iam.test.cloud.ibm.com",
	}

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
	powerClient := v.NewIBMPHostgroupsClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	getAllAvailableHost, err := powerClient.GetAvailableHosts()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[0]****************** %+v \n", getAllAvailableHost)

	getHostgroups, err := powerClient.GetHostGroups()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n", getHostgroups)
}
