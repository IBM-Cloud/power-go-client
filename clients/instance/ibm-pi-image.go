package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/runtime"
)

//IBMPIImageClient
type IBMPIImageClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPIImageClient
func NewIBMPIImageClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIImageClient {
	return &IBMPIImageClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudImages,
	}
}

// Get an Image
func (f *IBMPIImageClient) Get(imageID string) (*models.Image, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesImagesGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		ImageID:         imageID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudinstancesImagesGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetImageOperationFailed, imageID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Get Image Operation for image %s", imageID)
	}
	return resp.Payload, nil
}

// Get All Images that are imported into Power Instance
func (f *IBMPIImageClient) GetAll() (*models.Images, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesImagesGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudinstancesImagesGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PI Images of the PVM instance %s : %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Images of the PVM instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a Stock Image
func (f *IBMPIImageClient) Create(imageBody *models.CreateImage) (*models.Image, error) {

	// Create params and send request
	if imageBody.Source == nil || len(*imageBody.Source) == 0 {
		imageBody.Source = core.StringPtr("root-project")
	}
	params := &params.PcloudCloudinstancesImagesPostParams{
		Body:            imageBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	respOk, respCreated, err := f.request.PcloudCloudinstancesImagesPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateImageOperationFailed, f.cloudInstanceID, err)
	}
	if respOk != nil && respOk.Payload != nil {
		return respOk.Payload, nil
	}
	if respCreated != nil && respCreated.Payload != nil {
		return respCreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to perform Create Image Operation for cloud instance %s", f.cloudInstanceID)
}

// Import an Image
func (f *IBMPIImageClient) CreateCosImage(imageBody *models.CreateCosImageImportJob) (imageJob *models.JobReference, err error) {

	// Create params and send request
	params := &params.PcloudV1CloudinstancesCosimagesPostParams{
		Body:            imageBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudV1CloudinstancesCosimagesPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to perform Create COS Image Operation for cloud instance %s with error %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Create COS Image Operation for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Export an Image
func (f *IBMPIImageClient) ExportImage(imageID string, imageBody *models.ExportImage) (*models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudV2ImagesExportPostParams{
		CloudInstanceID: f.cloudInstanceID,
		Body:            imageBody,
		Context:         f.context,
		ImageID:         imageID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudV2ImagesExportPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Export COS Image for image id %s with error %w", imageID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Export COS Image for image id %s", imageID)
	}
	return resp.Payload, nil
}

// Delete an Image
func (f *IBMPIImageClient) Delete(imageId string) error {

	// Create params and send request
	params := &params.PcloudCloudinstancesImagesDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		ImageID:         imageId,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudCloudinstancesImagesDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf("failed to Delete PI Image %s: %w", imageId, err)
	}
	return nil
}

// Get a Stock Image
func (f *IBMPIImageClient) GetStockImage(imageId string) (*models.Image, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesStockimagesGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		ImageID:         imageId,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	resp, err := f.request.PcloudCloudinstancesStockimagesGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get PI Stock Image %s for cloud instance %s: %w", imageId, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Stock Image %s for cloud instance %s", imageId, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Stock Images
func (f *IBMPIImageClient) GetAllStockImages(includeSAP bool, includeVTl bool) (*models.Images, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesStockimagesGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		Sap:             &includeSAP,
		Vtl:             &includeVTl,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	resp, err := f.request.PcloudCloudinstancesStockimagesGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get Stock Images with (SAP=%t, VTL=%t) for cloud instance %s: %w", includeSAP, includeVTl, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get Stock Images with (SAP=%t, VTL=%t) for cloud instance %s", includeSAP, includeVTl, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Stock SAP Images
func (f *IBMPIImageClient) GetAllStockSAPImages() (*models.Images, error) {
	// get stock images. include all available SAP images
	images, err := f.GetAllStockImages(true, false)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// select SAP images
	sapImages := new(models.Images)
	for _, image := range images.Images {
		if image.Specifications.ImageType == "stock-sap" {
			sapImages.Images = append(sapImages.Images, image)
		}
	}
	return sapImages, nil
}

// Get All Stock VTL Images
func (f *IBMPIImageClient) GetAllStockVTLImages() (*models.Images, error) {
	// get stock images. include all available stock VTL images
	images, err := f.GetAllStockImages(false, true)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// select VTL images
	vtlImages := new(models.Images)
	for _, image := range images.Images {
		if image.Specifications.ImageType == "stock-vtl" {
			vtlImages.Images = append(vtlImages.Images, image)
		}
	}
	return vtlImages, nil
}

// Return true if Image is a VTL Image
func (f *IBMPIImageClient) IsVTLImage(imageId string) (bool, error) {
	images := new(models.Images)

	// get all stock vtl images
	stockVTLImages, err := f.GetAllStockVTLImages()

	// Handle errors
	if err != nil {
		return false, err
	}
	images.Images = append(images.Images, stockVTLImages.Images...)

	// get all images
	cloudInstanceImages, err := f.GetAll()

	// Handle errors
	if err != nil {
		return false, err
	}
	images.Images = append(images.Images, cloudInstanceImages.Images...)

	// search for image in list of all images
	for _, image := range images.Images {
		if imageId == *image.ImageID {
			if image.Specifications.ImageType == "stock-vtl" {
				return true, nil
			} else {
				return false, fmt.Errorf("image with id: %s is not a VTL image", imageId)
			}
		}
	}
	return false, fmt.Errorf("image with id: %s is not a VTL image", imageId)
}
