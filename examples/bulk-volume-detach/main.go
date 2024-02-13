package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
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
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}
	randomNumber2, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}
	piID := " < POWER INSTANCE ID > "
	name := fmt.Sprintf("power-go-test-volume-%d", randomNumber)
	name2 := fmt.Sprintf("power-go-test-volume-%d", randomNumber2)
	size := 20.0
	vtype := "tier3"
	sharable := true
	replicationEnabled := false

	// VM inputs
	serverName := fmt.Sprintf("power-go-test-%d", randomNumber)
	imageID := "cd447165-f45d-48f4-b07f-7c9a11ec84c7"
	networkID := "f324429e-5716-40c1-a01b-4d36ec8b82e3"
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
	powerClientVolume := v.NewIBMPIVolumeClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}
	bodyVolume := &models.CreateDataVolume{
		Name:               &name,
		Size:               &size,
		DiskType:           vtype,
		Shareable:          &sharable,
		ReplicationEnabled: &replicationEnabled,
	}
	bodyVolume2 := &models.CreateDataVolume{
		Name:               &name2,
		Size:               &size,
		DiskType:           vtype,
		Shareable:          &sharable,
		ReplicationEnabled: &replicationEnabled,
	}
	log.Print("Creating first Volume\n")
	createRespOk, err := powerClientVolume.CreateVolume(bodyVolume)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	volumeID := *createRespOk.VolumeID
	getResp, err := powerClientVolume.Get(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)
	log.Print("Creating Second Volume\n")
	createRespOk2, err := powerClientVolume.CreateVolume(bodyVolume2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v\n", *createRespOk2)

	volumeID2 := *createRespOk2.VolumeID
	getResp2, err := powerClientVolume.Get(volumeID2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", *getResp2)

	powerClientVM := v.NewIBMPIInstanceClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}
	bodyVM := &models.PVMInstanceCreate{
		ImageID:    &imageID,
		NetworkIDs: []string{networkID},
		ServerName: &serverName,
		Memory:     &memory,
		Processors: &processors,
		ProcType:   &procType,
		SysType:    sysType,
	}

	log.Print("Creating VM\n")
	createRespOkVM, err := powerClientVM.Create(bodyVM)
	if err != nil {
		log.Fatal(err)
	}
	if len(*createRespOkVM) == 0 {
		log.Fatal("create response is empty")
	}
	log.Printf("***************[5]****************** %+v\n", *createRespOk)

	insID := ""
	for _, in := range *createRespOkVM {
		insID = *in.PvmInstanceID
	}
	if insID == "" {
		log.Fatal("instance ID is empty")
	}

	getRespVM, err := powerClientVM.Get(insID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[6]****************** %+v \n", *getRespVM)
	waitForInstanceAvailable(insID, powerClientVM)

	log.Print("Attaching Volumes\n")
	attachBody := &models.VolumesAttach{
		VolumeIDs: []string{volumeID, volumeID2},
	}
	attachVolsResp, err := powerClientVolume.BulkVolumeAttach(insID, attachBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[7]****************** %+v\n", *attachVolsResp)
	time.Sleep(2 * time.Minute)

	log.Print("Detaching Volumes\n")
	detachBody := &models.VolumesDetach{
		VolumeIDs: []string{volumeID, volumeID2},
	}
	createRespOkDetach, err := powerClientVolume.BulkVolumeDetach(insID, detachBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[8]****************** %+v\n", *createRespOkDetach)
	log.Print("Deleting VM\n")
	err = powerClientVM.Delete(insID)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Deleting First Volume\n")
	err = powerClientVolume.DeleteVolume(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Deleting Second Volume\n")
	err = powerClientVolume.DeleteVolume(volumeID2)
	if err != nil {
		log.Fatal(err)
	}

}

func waitForInstanceAvailable(id string, powerClientVM *v.IBMPIInstanceClient) {
	for start := time.Now(); time.Since(start) < time.Minute*30; {
		time.Sleep(60 * time.Second)
		pvm, err := powerClientVM.Get(id)
		if err != nil {
			log.Fatal(err)
		}
		if *pvm.Status == helpers.PIInstanceAvailable && pvm.Health.Status == helpers.PIInstanceHealthOk {
			break
		}
	}
}
