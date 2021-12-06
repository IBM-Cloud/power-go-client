package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestImageCreateGetDelete(t *testing.T) {
	// check for testing variables
	testImagePreCheck(t)

	// create session and client
	session := getSession(t)
	imageClient := client.NewIBMPIImageClient(context.Background(), session, cloudInstanceID)

	// CREATE
	imageBody := &models.CreateImage{
		ImageName: imageName,
		ImageID:   imageID,
		Source:    &imageSource,
	}
	createResp, err := imageClient.Create(imageBody)
	if err != nil {
		t.Fatal(err)
	}
	imageID := *createResp.ImageID
	testMessage("CREATE Image", imageID, *createResp)

	// GET
	getResp, err := imageClient.Get(imageID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET Image", imageID, *getResp)

	// DELETE
	err = imageClient.Delete(imageID)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("DELETE Image", imageID, nil)
}
func TestImageGetAll(t *testing.T) {
	// check for testing variables
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	imageClient := client.NewIBMPIImageClient(context.Background(), session, cloudInstanceID)

	getAllResp, err := imageClient.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET All Images", "", *getAllResp)

	getStockResp, err := imageClient.GetAllStockImages(true, true)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET All Stock Images", "", *getStockResp)

	getSapResp, err := imageClient.GetAllStockImages(true, false)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET All SAP Images", "", *getSapResp)

	getVtlImages, err := imageClient.GetAllStockImages(false, true)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET All VTL Images", "", *getVtlImages)
}
