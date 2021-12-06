package test

import (
	"context"
	"log"
	"net"
	"strconv"
	"testing"

	"github.com/apparentlymart/go-cidr/cidr"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestNetworkCreateGetDelete(t *testing.T) {
	// check for testing variables
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
	if err != nil {
		t.Fatal(err)
	}
	networkID := *createNetResp.NetworkID
	testMessage("CREATE Network", networkID, nil)

	// CREATE Port
	portBody := &models.NetworkPortCreate{
		Description: "Network Port",
	}
	createPortResp, err := networkClient.CreatePort(networkID, portBody)
	if err != nil {
		t.Fatal(err)
	}
	portID := *createPortResp.PortID
	testMessage("CREATE Network Port", portID, *createPortResp)

	// GET Network
	getNetResp, err := networkClient.Get(networkID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET Network", networkID, *getNetResp)

	// GET Public
	getPublicResp, err := networkClient.GetAllPublic()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET Public", "", *getPublicResp)

	// GET Port
	getPortResp, err := networkClient.GetPort(networkID, portID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET Network ", networkID+"Port "+portID, *getPortResp)

	// GET ALL Ports
	getAllPortResp, err := networkClient.GetAllPorts(networkID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET ALL Network", networkID+" Ports", *getAllPortResp)

	// DELETE Port
	err = networkClient.DeletePort(networkID, portID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("DELETE Network Port", portID, nil)

	// DELETE Network
	err = networkClient.Delete(networkID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("DELETE Network", networkID, nil)
}

func generateIPData(t *testing.T, cdir string) (gway, firstip, lastip string) {
	_, ipv4Net, err := net.ParseCIDR(cdir)

	if err != nil {
		t.Fatal(err)
	}
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
	if err != nil {
		log.Printf("Failed to get the gateway for this cdir passed in %s", cdir)
		t.Fatal(err)
	}
	ad := cidr.AddressCount(ipv4Net)

	convertedad := strconv.FormatUint(ad, 10)

	// Powervc in wdc04 has to reserve 3 ip address hence we start from the 4th. This will be the default behaviour
	firstusable, err := cidr.Host(ipv4Net, 4)
	if err != nil {
		t.Fatal(err)
	}
	lastusable, err := cidr.Host(ipv4Net, subnetToSize[convertedad]-2)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("The gateway value is %s and the count is %s and first ip is %s last one is  %s", gateway, convertedad, firstusable, lastusable)
	return gateway.String(), firstusable.String(), lastusable.String()
}
