package test

import (
	"context"
	"log"
	"net"
	"strconv"
	"testing"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/stretchr/testify/require"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestNetworkCreateGetDelete(t *testing.T) {
	testNetworkPreCheck(t)

	// create session and client
	session := getSession(t)
	networkClient := client.NewIBMPINetworkClient(context.Background(), session, cloudInstanceID)

	// CREATE Network
	gateway, _, _ := generateIPData(t, networkCidr)
	networkBody := &models.NetworkCreate{
		Name:       networkName,
		Type:       &networkType,
		Cidr:       networkCidr,
		DNSServers: []string{networkDnsServer},
		Gateway:    gateway,
		Jumbo:      networkJumbo,
	}
	createNetResp, err := networkClient.Create(networkBody)
	require.Nil(t, err)
	networkID := *createNetResp.NetworkID
	testMessage("CREATE Network", networkID, nil)

	// GET Network
	getNetResp, err := networkClient.Get(networkID)
	require.Nil(t, err)
	testMessage("GET Network", networkID, *getNetResp)
	// verify variables match
	require.Equal(t, *createNetResp.Name, networkName)
	require.Equal(t, *createNetResp.Type, networkType)
	require.Equal(t, *createNetResp.Cidr, networkCidr)
	require.Equal(t, createNetResp.DNSServers, []string{networkDnsServer})
	require.Equal(t, createNetResp.Gateway, gateway)
	require.Equal(t, *createNetResp.Jumbo, networkJumbo)

	// CREATE Port
	portBody := &models.NetworkPortCreate{
		Description: "Network Port",
	}
	createPortResp, err := networkClient.CreatePort(networkID, portBody)
	require.Nil(t, err)
	portID := *createPortResp.PortID
	testMessage("CREATE Network Port", portID, *createPortResp)

	// GET Port
	getPortResp, err := networkClient.GetPort(networkID, portID)
	require.Nil(t, err)
	testMessage("GET Network ", networkID+" Port "+portID, *getPortResp)
	// verify variables match
	require.Equal(t, *createPortResp.Description, "Network Port")

	// GET Public
	getPublicResp, err := networkClient.GetAllPublic()
	require.Nil(t, err)
	testMessage("GET Public Networks", "", *getPublicResp)

	// GET ALL Ports
	getAllPortResp, err := networkClient.GetAllPorts(networkID)
	require.Nil(t, err)
	testMessage("GET ALL Network Ports", networkID, *getAllPortResp)

	// DELETE Port
	err = networkClient.DeletePort(networkID, portID)
	require.Nil(t, err)
	testMessage("DELETE Network: "+networkID+" Port", portID, nil)

	// DELETE Network
	err = networkClient.Delete(networkID)
	require.Nil(t, err)
	testMessage("DELETE Network", networkID, nil)
}

func generateIPData(t *testing.T, cdir string) (gway, firstip, lastip string) {
	_, ipv4Net, err := net.ParseCIDR(cdir)

	require.Nil(t, err)
	var subnetToSize = map[string]int{
		"21": 2048,
		"22": 1024,
		"23": 512,
		"24": 256,
		"25": 128,
		"26": 64,
		"27": 32,
		"28": 16,
		"29": 8,
		"30": 4,
		"31": 2,
	}
	gateway, err := cidr.Host(ipv4Net, 1)
	require.Nilf(t, err, "Failed to get the gateway for this cdir passed in %s", cdir)
	ad := cidr.AddressCount(ipv4Net)

	convertedad := strconv.FormatUint(ad, 10)

	// Powervc in wdc04 has to reserve 3 ip address hence we start from the 4th. This will be the default behaviour
	firstusable, err := cidr.Host(ipv4Net, 4)
	require.Nil(t, err)
	lastusable, err := cidr.Host(ipv4Net, subnetToSize[convertedad]-2)
	require.Nil(t, err)

	log.Printf("The gateway value is %s and the count is %s and first ip is %s last one is  %s", gateway, convertedad, firstusable, lastusable)
	return gateway.String(), firstusable.String(), lastusable.String()
}
