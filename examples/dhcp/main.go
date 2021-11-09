package main

import (
	"log"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
)

func main() {

	// session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "

	// dhcp inputs
	piID := " < POWER INSTANCE ID > "
	timeout := time.Duration(9000000000000000000)

	session, err := ps.New(token, region, true, timeout, accountID, zone)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIDhcpClient(session, piID)
	if err != nil {
		log.Fatal(err)
	}

	dhcpServer, err := powerClient.Create(piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", dhcpServer)

	dhcpId := *dhcpServer.ID
	getResp, err := powerClient.Get(dhcpId, piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getAllResp, err := powerClient.GetAll(piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", getAllResp)

	_, err = powerClient.Delete(dhcpId, piID)
	if err != nil {
		log.Fatal(err)
	}
}
