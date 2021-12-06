package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/joho/godotenv"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func main() {

	// set to staging or production
	environment := "staging"
	if environment == "staging" {
		godotenv.Load("../../.env.staging")
	} else {
		godotenv.Load("../../.env.production")
	}

	// load cloud instance id
	cloudInstanceId := os.Getenv("CLOUD_INSTANCE_ID")
	if cloudInstanceId == "" {
		log.Fatal(fmt.Errorf("CLOUD_INSTANCE_ID is empty: define in .env.%v", environment))
	}

	// network private VLANinputs
	name := " < NAME OF THE network > "
	netType := "vlan"
	cidr := "10.243.65.0/24"
	dnsServers := make([]string, 1)
	dnsServers[0] = "127.0.0.1"
	gateway, startIP, endIP := generateIPData(cidr)
	jumbo := false
	timeout := time.Duration(9000000000000000000)

	session, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPINetworkClient(session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	createRespOk, err := powerClient.Create(name, netType, cidr, dnsServers, gateway, startIP, endIP, jumbo, cloudInstanceId, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", createRespOk)

	networkID := *createRespOk.NetworkID
	getResp, err := powerClient.Get(networkID, cloudInstanceId, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getpubResp, err := powerClient.GetPublic(cloudInstanceId, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getpubResp)

	params := p_cloud_networks.PcloudNetworksPortsPostParams{
		CloudInstanceID: cloudInstanceId,
	}
	body := models.NetworkPortCreate{
		Description: "Network Port",
	}
	params.Body = &body
	log.Println(params)
	createPortResp, err := powerClient.CreatePort(networkID, cloudInstanceId, &params, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", *createPortResp)

	getPortResp, err := powerClient.GetPort(networkID, cloudInstanceId, *createPortResp.PortID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[5]****************** %+v \n", *getPortResp)

	getallPortResp, err := powerClient.GetAllPort(networkID, cloudInstanceId, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[6]****************** %+v \n", *getallPortResp)

	_, err = powerClient.DeletePort(networkID, cloudInstanceId, *createPortResp.PortID, timeout)
	if err != nil {
		log.Fatal(err)
	}

	err = powerClient.Delete(networkID, cloudInstanceId, timeout)
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
