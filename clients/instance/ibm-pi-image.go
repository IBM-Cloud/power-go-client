package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

//IBMPIImageClient ...
type IBMPIImageClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIImageClient ...
func NewIBMPIImageClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIImageClient {
	return &IBMPIImageClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

// Get PI Image
func (f *IBMPIImageClient) Get(id, powerinstanceid string) (*models.Image, error) {

	params := p_cloud_images.NewPcloudCloudinstancesImagesGetParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(powerinstanceid).WithImageID(id)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf(errors.GetImageOperationFailed, id, err)
	}
	return resp.Payload, nil
}

// Get with context
func (f *IBMPIImageClient) GetWithContext(ctx context.Context, id, cloudInstanceID string) (image *models.Image, err error) {
	params := p_cloud_images.NewPcloudCloudinstancesImagesGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithImageID(id)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	image = resp.Payload
	return
}

//GetAll Images that are imported into Power Instance
func (f *IBMPIImageClient) GetAll(powerinstanceid string) (*models.Images, error) {

	params := p_cloud_images.NewPcloudCloudinstancesImagesGetallParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get all PI Images of the PVM instance %s : %s", powerinstanceid, err)
	}
	return resp.Payload, nil

}

//Create the stock image
func (f *IBMPIImageClient) Create(name, imageid string, powerinstanceid string) (*models.Image, error) {

	var source = "root-project"
	var body = models.CreateImage{
		ImageName: name,
		ImageID:   imageid,
		Source:    &source,
	}
	params := p_cloud_images.NewPcloudCloudinstancesImagesPostParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithBody(&body)
	_, result, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || result == nil || result.Payload == nil {
		return nil, fmt.Errorf(errors.CreateImageOperationFailed, powerinstanceid, err)
	}
	return result.Payload, nil

}

// Import the image
func (f *IBMPIImageClient) CreateCosImage(body *models.CreateCosImageImportJob, cloudInstanceID string) (imageJob *models.JobReference, err error) {
	return f.CreateCosImageWithContext(context.Background(), body, cloudInstanceID)
}

func (f *IBMPIImageClient) CreateCosImageWithContext(ctx context.Context, body *models.CreateCosImageImportJob, cloudInstanceID string) (imageJob *models.JobReference, err error) {
	params := p_cloud_images.NewPcloudV1CloudinstancesCosimagesPostParamsWithContext(ctx).
		WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithBody(body)
	resp, err := f.session.Power.PCloudImages.PcloudV1CloudinstancesCosimagesPost(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		imageJob = resp.Payload
	}
	return
}

// Delete ...
func (f *IBMPIImageClient) Delete(id string, powerinstanceid string) error {
	params := p_cloud_images.NewPcloudCloudinstancesImagesDeleteParamsWithTimeout(helpers.PIDeleteTimeOut).WithCloudInstanceID(powerinstanceid).WithImageID(id)
	_, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return fmt.Errorf("Failed to Delete PI Image %s :%s", id, err)
	}
	return nil
}

// Delete with context...
func (f *IBMPIImageClient) DeleteWithContext(ctx context.Context, id string, cloudInstanceID string) (obj models.Object, err error) {
	params := p_cloud_images.NewPcloudCloudinstancesImagesDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithImageID(id)
	respOk, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if respOk != nil {
		obj = respOk.Payload
	}
	return
}

// GetStockImages ...
func (f *IBMPIImageClient) GetStockImage(id, powerinstanceid string) (*models.Image, error) {

	params := p_cloud_images.NewPcloudCloudinstancesStockimagesGetParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithImageID(id)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesStockimagesGet(params, ibmpisession.NewAuth(f.session, f.powerinstanceid))

	if err != nil || resp == nil {
		return nil, fmt.Errorf("Failed to Get PI Stock Imageid  %s : %s", powerinstanceid, err)
	}
	return resp.Payload, nil
}

// Get StockImage
func (f *IBMPIImageClient) GetStockImages(powerinstanceid string) (*models.Images, error) {

	params := p_cloud_images.NewPcloudCloudinstancesStockimagesGetallParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesStockimagesGetall(params, ibmpisession.NewAuth(f.session, f.powerinstanceid))

	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get all PI Stock Images of the PVM instance %s : %s", powerinstanceid, err)
	}
	return resp.Payload, nil
}

// GetSAPImages ...
func (f *IBMPIImageClient) GetSAPImages(powerinstanceid string, sapimage bool) (*models.Images, error) {

	params := p_cloud_images.NewPcloudImagesGetallParams().WithTimeout(helpers.PIGetTimeOut).WithSap(&sapimage)
	resp, err := f.session.Power.PCloudImages.PcloudImagesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get all PI Sap Images of the PVM instance %s : %s", powerinstanceid, err)
	}
	return resp.Payload, nil
}

// Get a single SAP Image

// Get single VTL Image ID
func (f *IBMPIImageClient) GetVTLImageID(cloudInstanceID string) (string, error) {

	yesVTL := true
	params := p_cloud_images.NewPcloudImagesGetallParams().WithTimeout(helpers.PIGetTimeOut).WithVtl(&yesVTL)

	resp, err := f.session.Power.PCloudImages.PcloudImagesGetall(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil || resp == nil || resp.Payload == nil {
		return "", fmt.Errorf("Failed to Get VTL ImageID for cloud instance %s : %s", cloudInstanceID, err)
	}

	// Find VTL image out of all images
	for _, image := range resp.Payload.Images {
		if image.Specifications.OperatingSystem == "vtl" {
			return *image.ImageID, nil
		}
	}
	return "", fmt.Errorf("Failed to Get VTL ImageID for cloud instance %s", cloudInstanceID)
}
