package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"github.com/apparentlymart/go-cidr/cidr"

	v "github.com/IBM-Cloud/power-go-client/clients"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".power-iaas.test.cloud.ibm.com"

	// network public vlan inputs
	// name := " < NAME OF THE network > "
	// piID := " < POWER INSTANCE ID > "
	// netType := "pub-vlan"
	// dnsServers := make([]string, 0)
	// cidr, gateway, startIP, endIP := "", "", "", ""

	// network private VLANinputs
	piID := " < POWER INSTANCE ID > "
	name := " < NAME OF THE network > "
	netType := "vlan"
	cidr := "10.243.65.0/24"
	dnsServers := make([]string, 1)
	dnsServers[0] = "127.0.0.1"
	gateway, startIP, endIP := generateIPData(cidr)
	jumbo := false

	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: token,
	}
	// authenticator := &core.IamAuthenticator{
	// 	ApiKey: "< API KEY >",
	// 	// Uncomment for test environment
	// 	URL: "https://iam.test.cloud.ibm.com",
	// }

	// Create the session
	options := &ps.IBMPIOptions{
		Authenticator: authenticator,
		UserAccount:   accountID,
		Zone:          zone,
		URL:           url,
		Debug:         true,
	}
	session, err := ps.NewIBMPISession(options)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPINetworkClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}
	body := &models.NetworkCreate{
		Type:  &netType,
		Name:  name,
		Jumbo: jumbo,
	}
	if netType == "vlan" {
		ipbody := []*models.IPAddressRange{
			{EndingIPAddress: &endIP, StartingIPAddress: &startIP}}
		body.IPAddressRanges = ipbody
		body.Gateway = gateway
		body.Cidr = cidr
	}
	if dnsServers != nil {
		body.DNSServers = dnsServers
	}
	createRespOk, err := powerClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	networkID := *createRespOk.NetworkID

	newName := "updated_name"
	updateBody := &models.NetworkUpdate{
		Name: &newName,
	}
	updateRespOk, err := powerClient.Update(networkID, updateBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v\n", *updateRespOk)

	getResp, err := powerClient.Get(networkID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getResp)

	getpubResp, err := powerClient.GetAllPublic()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", *getpubResp)

	portBody := &models.NetworkPortCreate{
		Description: "Network Port",
	}
	createPortResp, err := powerClient.CreatePort(networkID, portBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[5]****************** %+v \n", *createPortResp)

	getPortResp, err := powerClient.GetPort(networkID, *createPortResp.PortID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[6]****************** %+v \n", *getPortResp)

	getallPortResp, err := powerClient.GetAllPorts(networkID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[7]****************** %+v \n", *getallPortResp)

	err = powerClient.DeletePort(networkID, *createPortResp.PortID)
	if err != nil {
		log.Fatal(err)
	}

	err = powerClient.Delete(networkID)
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
	if err != nil {
		log.Fatal(err)
	}
	lastusable, err := cidr.Host(ipv4Net, subnetToSize[convertedad]-2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The gateway  value is %s and the count is %s and first ip is %s last one is  %s", gateway, convertedad, firstusable, lastusable)

	return gateway.String(), firstusable.String(), lastusable.String()
}
