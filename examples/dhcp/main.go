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

	// session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".power-iaas.test.cloud.ibm.com"

	// dhcp inputs
	piID := " < POWER INSTANCE ID > "

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
	powerClient := v.NewIBMPIDhcpClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	dhcpServer, err := powerClient.Create(&models.DHCPServerCreate{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *dhcpServer)

	dhcpId := *dhcpServer.ID
	getResp, err := powerClient.Get(dhcpId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getAllResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", getAllResp)

	err = powerClient.Delete(dhcpId)
	if err != nil {
		log.Fatal(err)
	}
}
