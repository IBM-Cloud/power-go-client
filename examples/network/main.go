package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/apparentlymart/go-cidr/cidr"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	accountID := " < ACCOUNT ID > "

	// network public vlan inputs
	// name := " < NAME OF THE network > "
	// piID := " < POWER INSTANCE ID > "
	// netType := "pub-vlan"
	// dnsServers := make([]string, 0)
	// cidr, gateway, startIP, endIP := "", "", "", ""
	// timeout := time.Duration(9000000000000000000)

	// network private VLANinputs
	name := " < NAME OF THE network > "
	piID := " < POWER INSTANCE ID > "
	netType := "vlan"
	cidr := "10.243.65.0/24"
	dnsServers := make([]string, 1)
	dnsServers[0] = "127.0.0.1"
	gateway, startIP, endIP := generateIPData(cidr)
	jumbo := false
	timeout := time.Duration(9000000000000000000)

	session, err := ps.New(token, region, true, 9000000000000000000, accountID, region)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPINetworkClient(session, piID)
	if err != nil {
		log.Fatal(err)
	}
	createRespOk, err := powerClient.Create(name, netType, cidr, dnsServers, gateway, startIP, endIP, jumbo, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createRespOk)

	networkID := *createRespOk.NetworkID
	getResp, err := powerClient.Get(networkID, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getpubResp, err := powerClient.GetPublic(piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getpubResp)

	params := p_cloud_networks.PcloudNetworksPortsPostParams{
		CloudInstanceID: piID,
	}
	body := models.NetworkPortCreate{
		Description: "Network Port",
	}
	params.Body = &body
	log.Println(params)
	createPortResp, err := powerClient.CreatePort(networkID, piID, &params, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", *createPortResp)

	getPortResp, err := powerClient.GetPort(networkID, piID, *createPortResp.PortID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[5]****************** %+v \n", *getPortResp)

	getallPortResp, err := powerClient.GetAllPort(networkID, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[6]****************** %+v \n", *getallPortResp)

	_, err = powerClient.DeletePort(networkID, piID, *createPortResp.PortID, timeout)
	if err != nil {
		log.Fatal(err)
	}

	err = powerClient.Delete(networkID, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
}

func generateIPData(cdir string) (gway, firstip, lastip string) {
	_, ipv4Net, err := net.ParseCIDR(cdir)

	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	ad := cidr.AddressCount(ipv4Net)

	convertedad := strconv.FormatUint(ad, 10)
	// Powervc in wdc04 has to reserve 3 ip address hence we start from the 4th. This will be the default behaviour
	firstusable, err := cidr.Host(ipv4Net, 4)
	lastusable, err := cidr.Host(ipv4Net, subnetToSize[convertedad]-2)
	log.Printf("The gateway  value is %s and the count is %s and first ip is %s last one is  %s", gateway, convertedad, firstusable, lastusable)

	return gateway.String(), firstusable.String(), lastusable.String()
}
