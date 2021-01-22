package main

import (
	"log"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	accountID := " < ACCOUNT ID > "

	// volume inputs
	name := " < NAME OF THE volume > "
	piID := " < POWER INSTANCE ID > "
	timeout := time.Duration(9000000000000000000)
	imageID := "c6b32fda-9979-4ce7-abee-ecb54df5237a"
	volumes := make([]string, 1)
	volumes[0] = "ef82d430-4cb0-46a5-be11-e61fb129fe18"

	networks := make([]string, 1)
	networks[0] = "69fd9be7-2b68-410f-ab39-05d41e619dca"
	memory := 4.0
	processors := 2.0
	procType := "shared"
	sysType := "s922"

	session, err := ps.New(token, region, true, 9000000000000000000, accountID, region)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIInstanceClient(session, piID)
	if err != nil {
		log.Fatal(err)
	}

	params := p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
		Body: &models.PVMInstanceCreate{
			ImageID:     &imageID,
			KeyPairName: "test-key",
			NetworkIds:  networks,
			ServerName:  &name,
			VolumeIds:   volumes,
			Memory:      &memory,
			Processors:  &processors,
			ProcType:    &procType,
			SysType:     sysType,
		},
		CloudInstanceID: piID,
	}
	createRespOk, err := powerClient.Create(&params, piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	insIDs := make([]string, 0)
	for _, in := range *createRespOk {
		insID := in.PvmInstanceID
		insIDs = append(insIDs, *insID)
	}

	getResp, err := powerClient.Get(insIDs[0], piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := powerClient.GetAll(piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getallResp)

	err = powerClient.Delete(insIDs[0], piID, timeout)
	if err != nil {
		log.Fatal(err)
	}
}
