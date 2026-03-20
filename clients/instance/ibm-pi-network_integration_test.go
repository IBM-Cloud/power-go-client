package instance_test

import (
	"context"
	"log"
	"net"
	"strconv"
	"testing"

	utl "github.com/IBM-Cloud/power-go-client/internal/testutils"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/stretchr/testify/require"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestNetwork(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and client
	utl.NetworkPreCheck(t)
	session := utl.TestSession(t)
	networkClient := client.NewIBMPINetworkClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE Network
	gateway, _, _ := generateIPData(t, utl.NetworkCidr)
	networkBody := &models.NetworkCreate{
		Name:         utl.NetworkName,
		Type:         &utl.NetworkType,
		Cidr:         utl.NetworkCidr,
		DNSServers:   []string{utl.NetworkDNSServer},
		Gateway:      gateway,
		Jumbo:        utl.NetworkJumbo,
		Mtu:          &utl.NetworkMtu,
		AccessConfig: utl.NetworkAccessConfig,
	}
	createNetResp, err := networkClient.Create(networkBody)
	require.Nil(t, err)
	NetworkID := *createNetResp.NetworkID
	utl.TestMessage("CREATE Network", NetworkID, nil)

	// DELETE Network
	defer func() {
		err = networkClient.Delete(NetworkID)
		require.Nil(t, err)
		utl.TestMessage("DELETE Network", NetworkID, nil)
	}()

	// GET Network
	getNetResp, err := networkClient.Get(NetworkID)
	require.Nil(t, err)
	utl.TestMessage("GET Network", NetworkID, *getNetResp)
	// verify variables match
	require.Equal(t, *createNetResp.Name, utl.NetworkName)
	require.Equal(t, *createNetResp.Type, utl.NetworkType)
	require.Equal(t, *createNetResp.Cidr, utl.NetworkCidr)
	require.Equal(t, createNetResp.DNSServers, []string{utl.NetworkDNSServer})
	require.Equal(t, createNetResp.Gateway, gateway)
	require.Equal(t, createNetResp.Jumbo, utl.NetworkJumbo)
	require.Equal(t, createNetResp.Mtu, utl.NetworkMtu)
	require.Equal(t, createNetResp.AccessConfig, utl.NetworkAccessConfig)

	// CREATE Network Interface (replacing deprecated CreatePort)
	netIntBody := &models.NetworkInterfaceCreate{
		Name: "Network Interface",
	}
	createNetIntResp, err := networkClient.CreateNetworkInterface(NetworkID, netIntBody)
	require.Nil(t, err)
	netIntID := *createNetIntResp.ID
	utl.TestMessage("CREATE Network Interface", netIntID, *createNetIntResp)

	// DELETE Network Interface (replacing deprecated DeletePort)
	defer func() {
		err = networkClient.DeleteNetworkInterface(NetworkID, netIntID)
		require.Nil(t, err)
		utl.TestMessage("DELETE Network: "+NetworkID+" Interface", netIntID, nil)
	}()

	// GET Network Interface (replacing deprecated GetPort)
	getNetIntResp, err := networkClient.GetNetworkInterface(NetworkID, netIntID)
	require.Nil(t, err)
	utl.TestMessage("GET Network ", NetworkID+" Interface "+netIntID, *getNetIntResp)
	// verify variables match
	require.Equal(t, *createNetIntResp.Name, "Network Interface")

	// GET ALL Network Interfaces (replacing deprecated GetAllPorts)
	getAllNetIntResp, err := networkClient.GetAllNetworkInterfaces(NetworkID)
	require.Nil(t, err)
	utl.TestMessage("GET ALL Network Interfaces", NetworkID, *getAllNetIntResp)

	// GET Public Networks
	getPublicResp, err := networkClient.GetAllPublic()
	require.Nil(t, err)
	utl.TestMessage("GET Public Networks", "", *getPublicResp)
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
