package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"

	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {
	// token := ""
	apiKey := ""
	region := ""
	zone := ""
	accountID := ""
	url := region + ".power-iaas.test.cloud.ibm.com"

	piID := ""
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

	powerClient := v.NewIBMPIDatacenterClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}
	getDatacenter, err := powerClient.Get("dal12")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[0]****************** %+v \n", getDatacenter)
	getDatacenters, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n", getDatacenters)
}
