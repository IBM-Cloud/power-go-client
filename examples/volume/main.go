package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {

	//session Inputs
	//  < IAM TOKEN >
	token := ""
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".power-iaas.test.cloud.ibm.com"

	// volume inputs
	piID := " < POWER INSTANCE ID > "
	name := " < NAME OF THE volume > "
	size := 20.0
	vtype := "tier1"
	sharable := true
	replicationEnabled := false

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
	powerClient := v.NewIBMPIVolumeClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	body := &models.CreateDataVolume{
		Name:               &name,
		Size:               &size,
		DiskType:           vtype,
		Shareable:          &sharable,
		ReplicationEnabled: &replicationEnabled,
	}
	createRespOk, err := powerClient.CreateVolume(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	volumeID := *createRespOk.VolumeID
	getResp, err := powerClient.Get(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getallResp)

	getVolRCRelationships, err := powerClient.GetVolumeRemoteCopyRelationships(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getVolRCRelationships)

	getVolFCMappings, err := powerClient.GetVolumeFlashCopyMappings(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", getVolFCMappings)

	err = powerClient.DeleteVolume(volumeID)
	if err != nil {
		log.Fatal(err)
	}
}
