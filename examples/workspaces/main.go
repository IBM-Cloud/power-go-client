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
	region := "dal"
	zone := "satloc_dal_cjkfnk0205fov31ir19g"
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
	powerClient := v.NewIBMPIWorkspacesClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}
	// getWorkspace, err := powerClient.Get(piID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("***************[0]****************** %+v \n", *getWorkspace)
	// getWorkspaces, err := powerClient.GetAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("***************[1]****************** %+v \n", *getWorkspaces)
	createWorkspace := powerClient.Create("Michael_ws_test", "satloc_dal_cjkfnk0205fov31ir19g", "f47bd99f28544e609886585f99c3743a", "1112d6a9-71d6-4968-956b-eb3edbf0225b")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Printf("***************[2]****************** %+v \n", createWorkspace)
}
