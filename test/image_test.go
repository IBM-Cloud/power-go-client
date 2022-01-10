package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestImageCreateGetDelete(t *testing.T) {
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
	require.Nil(t, err)
	createdImageID := *createResp.ImageID
	testMessage("CREATE Image", createdImageID, *createResp)

	// GET
	getResp, err := imageClient.Get(createdImageID)
	require.Nil(t, err)
	testMessage("GET Image", createdImageID, *getResp)
	// verify variables match
	//require.Equal(t, *getResp.Name, imageName)

	//DELETE
	err = imageClient.Delete(createdImageID)
	require.Nil(t, err)
	testMessage("DELETE Image", createdImageID, nil)
}
func TestImageGetAll(t *testing.T) {
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	imageClient := client.NewIBMPIImageClient(context.Background(), session, cloudInstanceID)

	getAllResp, err := imageClient.GetAll()
	require.Nil(t, err)
	testMessage("GET All Images", "", *getAllResp)

	getStockResp, err := imageClient.GetAllStockImages(true, true)
	require.Nil(t, err)
	testMessage("GET All Stock Images", "", *getStockResp)

	getSapResp, err := imageClient.GetAllStockImages(true, false)
	require.Nil(t, err)
	testMessage("GET All SAP Images", "", *getSapResp)

	getVtlImages, err := imageClient.GetAllStockImages(false, true)
	require.Nil(t, err)
	testMessage("GET All VTL Images", "", *getVtlImages)
}
