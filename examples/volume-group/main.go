package main

import (
	"context"
	"log"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
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
	piID := " < POWER INSTANCE ID > "
	vgname := "< NAME OF THE volume > "
	volIDs := []string{" < VOLUME ID > "}
	rmVolIDs := []string{" < VOLUME ID > "}
	addVolIDs := []string{" < VOLUME ID > "}

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
	powerClient := v.NewIBMPIVolumeGroupClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	body := &models.VolumeGroupCreate{
		Name:      vgname,
		VolumeIDs: volIDs,
	}

	createRespOk, err := powerClient.CreateVolumeGroup(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	vgID := *createRespOk.ID
	getResp, err := powerClient.Get(vgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v\n", *getResp)

	getDetailsResp, err := powerClient.GetDetails(vgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v\n", *getDetailsResp)

	vgStatusCheck(vgID, powerClient)

	updateBody := &models.VolumeGroupUpdate{
		AddVolumes:    addVolIDs,
		RemoveVolumes: rmVolIDs,
	}

	err = powerClient.UpdateVolumeGroup(vgID, updateBody)
	if err != nil {
		log.Fatal(err)
	}

	vgStatusCheck(vgID, powerClient)

	getAllVG, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v\n", *getAllVG)

	getAllVGDetails, err := powerClient.GetAllDetails()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[5]****************** %+v\n", *getAllVGDetails)

	getDetailsResp, err = powerClient.GetDetails(vgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[6]****************** %+v\n", *getDetailsResp)

	getLiveDetails, err := powerClient.GetVolumeGroupLiveDetails(vgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[7]****************** %+v\n", *getLiveDetails)

	getRCRelationships, err := powerClient.GetVolumeGroupRemoteCopyRelationships(vgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[8]****************** %+v\n", *getRCRelationships)

	// removing volumes before deleting
	updateBody = &models.VolumeGroupUpdate{
		RemoveVolumes: getDetailsResp.VolumeIDs,
	}
	err = powerClient.UpdateVolumeGroup(vgID, updateBody)
	if err != nil {
		log.Fatal(err)
	}

	vgStatusCheck(vgID, powerClient)

	err = powerClient.DeleteVolumeGroup(vgID)
	if err != nil {
		log.Fatal(err)
	}
}

// vgStatusCheck waits for volume-group status to available
func vgStatusCheck(id string, powerClient *v.IBMPIVolumeGroupClient) {
	for start := time.Now(); time.Since(start) < time.Second*30; {
		time.Sleep(10 * time.Second)
		vg, err := powerClient.GetDetails(id)
		if err != nil {
			log.Fatal(err)
		}
		if vg.Status == "available" {
			break
		}
	}
}
