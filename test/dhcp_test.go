package test

import (
	"context"
	"testing"
	"time"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestDHCPCreateGetDelete(t *testing.T) {
	// check for testing variables
	testDHCPPreCheck(t)

	// create session and clients
	session := getSession(t)
	dhcpClient := client.NewIBMPIDhcpClient(context.Background(), session, cloudInstanceID)

	// CREATE
	createResp, err := dhcpClient.Create()
	if err != nil {
		t.Fatal(err)
	}
	dhcpID := *createResp.ID
	testMessage("CREATE DHCP", dhcpID, *createResp)

	// GET
	// wait until DHCP is built to before deletion
	dhcpBuildComplete := false
	for dhcpBuildComplete == false {
		getResp, err := dhcpClient.Get(dhcpID)
		if err != nil {
			t.Fatal(err)
		}
		if *getResp.Status == "BUILD" {
			time.Sleep(time.Duration(jobTimeout))
		} else {
			dhcpBuildComplete = true
			testMessage("GET DHCP", dhcpID, *getResp)
		}
	}

	// DELETE
	err = dhcpClient.Delete(dhcpID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("DELETE DHCP", dhcpID, nil)
}
func TestDHCPGetAll(t *testing.T) {
	// check for testing variables
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	dhcpClient := client.NewIBMPIDhcpClient(context.Background(), session, cloudInstanceID)

	// GET ALL
	getAllResp, err := dhcpClient.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET ALL DHCPs", "", getAllResp)
}
