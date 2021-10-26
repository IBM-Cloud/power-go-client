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

// GetAll Images that are imported into Power Instance
func (f *IBMPIImageClient) GetAll(powerinstanceid string) (*models.Images, error) {

	params := p_cloud_images.NewPcloudCloudinstancesImagesGetallParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get all PI Images of the PVM instance %s : %s", powerinstanceid, err)
	}
	return resp.Payload, nil
}

// Create the stock image
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
		return nil, fmt.Errorf("Failed to Get PI Stock Image with id: %s of the cloud instance  %s : %s", id, powerinstanceid, err)
	}
	return resp.Payload, nil
}

// Get StockImage
func (f *IBMPIImageClient) GetAllStockImages(cloudInstanceID string, includeSAP bool, includeVTl bool) (*models.Images, error) {

	params := p_cloud_images.NewPcloudCloudinstancesStockimagesGetallParams().
		WithCloudInstanceID(cloudInstanceID).
		WithTimeout(helpers.PICreateTimeOut).
		WithSap(&includeSAP).
		WithVtl(&includeVTl)

	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesStockimagesGetall(params, ibmpisession.NewAuth(f.session, f.powerinstanceid))
	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to get all PI Stock Images with (SAP=%t, VTL=%t) of the cloud instance %s : %s", includeSAP, includeVTl, cloudInstanceID, err)
	}
	return resp.Payload, nil
}

// GetAllStockSAPImages returns all stock SAP images. No Other images are included
func (f *IBMPIImageClient) GetAllStockSAPImages(cloudInstanceID string) (*models.Images, error) {

	// get stock images. include all available SAP images
	images, err := f.GetAllStockImages(cloudInstanceID, true, false)
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

// GetAllStockVTLImages returns all VTL images. No Other images are included
func (f *IBMPIImageClient) GetAllStockVTLImages(cloudInstanceID string) (*models.Images, error) {

	// get stock images. include all available stock VTL images
	images, err := f.GetAllStockImages(cloudInstanceID, false, true)
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

// IsVtlImage returns true if image is a VTL images
func (f *IBMPIImageClient) IsVTLImage(imageId string, cloudInstanceID string) (bool, error) {

	images := new(models.Images)

	// get all stock vtl images
	stockVTLImages, err := f.GetAllStockVTLImages(cloudInstanceID)
	if err != nil {
		return false, err
	}
	images.Images = append(images.Images, stockVTLImages.Images...)

	// get al images
	cloudInstanceImages, err := f.GetAll(cloudInstanceID)
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
				return false, fmt.Errorf("Image with id: %s is not a VTL image", imageId)
			}
		}
	}
	return false, fmt.Errorf("Image with id: %s is not a VTL image", imageId)
}
