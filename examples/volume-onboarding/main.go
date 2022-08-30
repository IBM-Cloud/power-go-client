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
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".power-iaas.test.cloud.ibm.com"

	// volume-onboarding inputs
	piID := " < POWER INSTANCE ID > "
	auxVolName := " < NAME OF THE auxiliary volume > "
	sourceCRN := " < SOURCE CRN > "
	onboardingName := " < ONBOARDING NAME > "

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
	powerClient := v.NewIBMPIVolumeOnboardingClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	auxVols := []*models.AuxiliaryVolumeForOnboarding{}
	auxVols = append(auxVols, &models.AuxiliaryVolumeForOnboarding{
		AuxVolumeName: &auxVolName,
		Name:          onboardingName,
	})

	vols := []*models.AuxiliaryVolumesForOnboarding{}
	vols = append(vols, &models.AuxiliaryVolumesForOnboarding{
		AuxiliaryVolumes: auxVols,
		SourceCRN:        &sourceCRN,
	})
	body := &models.VolumeOnboardingCreate{
		Volumes: vols,
	}

	createRespOk, err := powerClient.CreateVolumeOnboarding(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	onboardingID := createRespOk.ID
	getResp, err := powerClient.Get(onboardingID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getallResp)
}
