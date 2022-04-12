package instance_test

import (
	"context"
	"testing"

	utl "internal/testutils"

	client "github.com/IBM-Cloud/power-go-client/clients"
	"github.com/stretchr/testify/require"
)

func TestImage(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and client
	utl.ImagePreCheck(t)
	session := utl.TestSession(t)
	imageClient := client.NewIBMPIImageClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE
	// imageBody := &models.CreateImage{
	// 	ImageID: utl.ImageID,
	// 	Source:  &utl.ImageSource,
	// }
	// createResp, err := imageClient.Create(imageBody)
	// require.Nil(t, err)
	// createdImageID := *createResp.ImageID
	// utl.TestMessage("CREATE Image", createdImageID, *createResp)

	// // DELETE
	// defer func() {
	// 	err = imageClient.Delete(createdImageID)
	// 	require.Nil(t, err)
	// 	utl.TestMessage("DELETE Image", createdImageID, nil)
	// }()

	// GET
	// getResp, err := imageClient.Get(createdImageID)
	// require.Nil(t, err)
	// utl.TestMessage("GET Image", createdImageID, *getResp)

	// GET ALL
	getAllResp, err := imageClient.GetAll()
	require.Nil(t, err)
	utl.TestMessage("GET All Images", "", *getAllResp)

	// GET All SAP Images
	getSapResp, err := imageClient.GetAllStockImages(true, false)
	require.Nil(t, err)
	utl.TestMessage("GET All SAP Images", "", *getSapResp)

	// GET ALL Stock Images
	getStockResp, err := imageClient.GetAllStockImages(true, true)
	require.Nil(t, err)
	utl.TestMessage("GET All Stock Images", "", *getStockResp)

	// GET ALL VTL Images
	getVtlImages, err := imageClient.GetAllStockImages(false, true)
	require.Nil(t, err)
	utl.TestMessage("GET All VTL Images", "", *getVtlImages)
}
