package instance_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	utl "internal/testutils"

	client "github.com/IBM-Cloud/power-go-client/clients"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestDHCP(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and client
	utl.DHCPPreCheck(t)
	session := utl.TestSession(t)
	dhcpClient := client.NewIBMPIDhcpClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE
	createResp, err := dhcpClient.Create(&models.DHCPServerCreate{})
	require.Nil(t, err)
	dhcpID := *createResp.ID

	// DELETE
	defer func() {
		err = dhcpClient.Delete(dhcpID)
		require.Nil(t, err)
		utl.TestMessage("DELETE DHCP", dhcpID, nil)
	}()

	// check that DHCP is successfully created
	createErr := fmt.Sprintf("Create DHCP %s has failed. Still in Build state after %d minutes", dhcpID, ((utl.TimeoutAttempts * utl.JobTimeout) / 60000000000))
	checkAttempts := 0
	getResp := &models.DHCPServerDetail{}
	for checkAttempts < utl.TimeoutAttempts {
		getResp, err = dhcpClient.Get(dhcpID)
		require.Nil(t, err)
		if *getResp.Status == "ERROR" {
			t.Fatalf("Newly Created DHCP: %s has Status: ERROR", dhcpID)
		}
		if *getResp.Status != "BUILD" {
			break
		}
		time.Sleep(time.Duration(utl.JobTimeout))
		checkAttempts += 1
	}
	require.Less(t, checkAttempts, utl.TimeoutAttempts, createErr)
	utl.TestMessage("CREATE DHCP", dhcpID, *createResp)

	// GET
	getResp, err = dhcpClient.Get(dhcpID)
	require.Nil(t, err)
	utl.TestMessage("GET DHCP", dhcpID, *getResp)

	// GET ALL
	getAllResp, err := dhcpClient.GetAll()
	require.Nil(t, err)
	utl.TestMessage("GET ALL DHCPs", "", getAllResp)
}
