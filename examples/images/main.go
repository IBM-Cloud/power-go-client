package main

import (
	"fmt"
	"log"
	"os"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/joho/godotenv"
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

	// Image inputs
	name := " < NAME OF THE IMAGE > "
	image := " <IMAGE ID> "

	session, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIImageClient(session, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}

	createResp, err := powerClient.Create(name, image, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n\n", *createResp)

	imageID := *createResp.ImageID
	getResp, err := powerClient.Get(imageID, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n\n", *getResp)

	getAllResp, err := powerClient.GetAll(cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n\n", *getAllResp)

	err = powerClient.Delete(imageID, cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}

	getStockResp, err := powerClient.GetAllStockImages(cloudInstanceId, true, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n\n", *getStockResp)

	getSapResp, err := powerClient.GetAllStockSAPImages(cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[5]****************** %+v \n\n", *getSapResp)

	getVtlImages, err := powerClient.GetAllStockVTLImages(cloudInstanceId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[6]****************** %+v \n\n", *getVtlImages)

	if len(getVtlImages.Images) > 0 {
		testVtlId := *getVtlImages.Images[0].ImageID
		imageCheck, err := powerClient.IsVTLImage(testVtlId, cloudInstanceId)

		msg := fmt.Sprintf("IsVtlImage returned true for vtl image with ID %s", testVtlId)
		if imageCheck == false {
			msg = fmt.Sprintf("IsVtlImage returned false for image with ID %s. Error: %s", testVtlId, err)
		}
		log.Printf("***************[7]****************** %s \n\n", msg)
	}
}
