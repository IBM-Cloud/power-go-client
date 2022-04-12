package main

import (
	"context"
	"fmt"
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

	// Image inputs
	piID := " < POWER INSTANCE ID > "
	name := " < NAME OF THE IMAGE > "
	image := " <IMAGE ID> "

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
	powerClient := v.NewIBMPIImageClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	body := &models.CreateImage{
		ImageName: name,
		ImageID:   image,
	}
	createResp, err := powerClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n\n", *createResp)

	imageID := *createResp.ImageID
	getResp, err := powerClient.Get(imageID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n\n", *getResp)

	getAllResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n\n", *getAllResp)

	err = powerClient.Delete(imageID)
	if err != nil {
		log.Fatal(err)
	}

	getStockResp, err := powerClient.GetAllStockImages(true, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n\n", *getStockResp)

	getSapResp, err := powerClient.GetAllStockSAPImages()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[5]****************** %+v \n\n", *getSapResp)

	getVtlImages, err := powerClient.GetAllStockVTLImages()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[6]****************** %+v \n\n", *getVtlImages)

	if len(getVtlImages.Images) > 0 {
		testVtlId := *getVtlImages.Images[0].ImageID
		imageCheck, err := powerClient.IsVTLImage(testVtlId)

		msg := fmt.Sprintf("IsVtlImage returned true for vtl image with ID %s", testVtlId)
		if imageCheck == false {
			msg = fmt.Sprintf("IsVtlImage returned false for image with ID %s. Error: %s", testVtlId, err)
		}
		log.Printf("***************[7]****************** %s \n\n", msg)
	}
}
