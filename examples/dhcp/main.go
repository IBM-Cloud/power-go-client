package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
)

func main() {

	// session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	//os.Setenv("IBMCLOUD_POWER_API_ENDPOINT", region+".power-iaas.test.cloud.ibm.com")

	// dhcp inputs
	piID := " < POWER INSTANCE ID > "

	session, err := ps.New(token, region, true, accountID, zone)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIDhcpClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	dhcpServer, err := powerClient.Create()
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
