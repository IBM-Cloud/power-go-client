package main

import (
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
	getSapResp, err := powerClient.GetSAPImages(piID, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n\n", *getSapResp)

	getStockResp, err := powerClient.GetStockImages(piID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[5]****************** %+v \n\n", *getStockResp)
}
