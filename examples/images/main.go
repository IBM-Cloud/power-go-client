package main

import (
	"fmt"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	accountID := " < ACCOUNT ID > "

	// Image inputs
	name := " < NAME OF THE IMAGE > "
	piID := " < POWER INSTANCE ID > "
	image := " <IMAGE ID> "

	session, err := ps.New(token, region, true, 50000000, accountID, region)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIImageClient(session, piID)
	if err != nil {
		log.Fatal(err)
	}

	createResp, err := powerClient.Create(name, image, piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n\n", *createResp)

	imageID := *createResp.ImageID
	getResp, err := powerClient.Get(imageID, piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n\n", *getResp)

	getAllResp, err := powerClient.GetAll(piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n\n", *getAllResp)

	err = powerClient.Delete(imageID, piID)
	if err != nil {
		log.Fatal(err)
	}

	getStockResp, err := powerClient.GetAllStockImages(piID, true, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n\n", *getStockResp)

	getSapResp, err := powerClient.GetAllStockSAPImages(piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[5]****************** %+v \n\n", *getSapResp)

	getVtlImages, err := powerClient.GetAllStockVTLImages(piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[6]****************** %+v \n\n", *getVtlImages)

	if len(getVtlImages.Images) > 0 {
		testVtlId := *getVtlImages.Images[0].ImageID
		imageCheck, err := powerClient.IsVTLImage(testVtlId, piID)

		msg := fmt.Sprintf("IsVtlImage returned true for vtl image with ID %s", testVtlId)
		if imageCheck == false {
			msg = fmt.Sprintf("IsVtlImage returned false for image with ID %s. Error: %s", testVtlId, err)
		}
		log.Printf("***************[7]****************** %s \n\n", msg)
	}
}
