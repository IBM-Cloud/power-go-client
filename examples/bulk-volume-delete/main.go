package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

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
	time.Sleep(2 * time.Minute)
	// log.Print("Deleting Volumes\n")
	// body := &models.VolumesDelete{VolumeIDs: []string{volumeID, volumeID2}}
	// delResp, err := powerClientVolume.BulkVolumeDelete(body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if delResp == nil {
	// 	log.Print("***************[5]****************** Deleted\n")

	// } else {

	// 	log.Printf("***************[5]****************** %+v\n", *delResp)
	// }

}
