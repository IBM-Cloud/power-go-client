package main

import (
	"context"
	"log"

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

	// volume inputs
	name := " < NAME OF THE volume > "
	piID := " < POWER INSTANCE ID > "
	imageID := "c6b32fda-9979-4ce7-abee-ecb54df5237a"
	volumes := make([]string, 1)
	volumes[0] = "ef82d430-4cb0-46a5-be11-e61fb129fe18"

	networks := make([]string, 1)
	networks[0] = "69fd9be7-2b68-410f-ab39-05d41e619dca"
	memory := 4.0
	processors := 2.0
	procType := "shared"
	sysType := "s922"

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
	powerClient := v.NewIBMPIInstanceClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	body := &models.PVMInstanceCreate{
		ImageID:     &imageID,
		KeyPairName: "test-key",
		NetworkIDs:  networks,
		ServerName:  &name,
		VolumeIDs:   volumes,
		Memory:      &memory,
		Processors:  &processors,
		ProcType:    &procType,
		SysType:     sysType,
	}
	createRespOk, err := powerClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	insIDs := make([]string, 0)
	for _, in := range *createRespOk {
		insID := in.PvmInstanceID
		insIDs = append(insIDs, *insID)
	}

	getResp, err := powerClient.Get(insIDs[0])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getallResp)

	err = powerClient.Delete(insIDs[0])
	if err != nil {
		log.Fatal(err)
	}
}
