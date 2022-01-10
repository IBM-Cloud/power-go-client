package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestDHCPCreateGetDelete(t *testing.T) {
	testDHCPPreCheck(t)

	// create session and clients
	session := getSession(t)
	dhcpClient := client.NewIBMPIDhcpClient(context.Background(), session, cloudInstanceID)

	// CREATE
	createResp, err := dhcpClient.Create()
	require.Nil(t, err)
	dhcpID := *createResp.ID

	// check that DHCP is successfully created
	createErr := fmt.Sprintf("Create DHCP %s has failed. Still in Build state after %d minutes", dhcpID, ((timeoutAttempts * jobTimeout) / 60000000000))
	checkAttempts := 0
	getResp := &models.DHCPServerDetail{}
	for checkAttempts < timeoutAttempts {
		getResp, err = dhcpClient.Get(dhcpID)
		require.Nil(t, err)
		if *getResp.Status == "ERROR" {
			t.Fatalf("Newly Created DHCP: %s has Status: ERROR", dhcpID)
		}
		if *getResp.Status != "BUILD" {
			break
		}
		time.Sleep(time.Duration(jobTimeout))
		checkAttempts += 1
	}
	require.Less(t, checkAttempts, timeoutAttempts, createErr)
	testMessage("CREATE DHCP", dhcpID, *createResp)

	// GET
	getResp, err = dhcpClient.Get(dhcpID)
	require.Nil(t, err)
	testMessage("GET DHCP", dhcpID, *getResp)

	// DELETE
	err = dhcpClient.Delete(dhcpID)
	require.Nil(t, err)
	testMessage("DELETE DHCP", dhcpID, nil)
}
func TestDHCPGetAll(t *testing.T) {
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	dhcpClient := client.NewIBMPIDhcpClient(context.Background(), session, cloudInstanceID)

	// GET ALL
	getAllResp, err := dhcpClient.GetAll()
	require.Nil(t, err)
	testMessage("GET ALL DHCPs", "", getAllResp)
}
