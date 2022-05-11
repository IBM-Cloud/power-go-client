package instance

import (
	"context"
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"

	"github.com/IBM-Cloud/power-go-client/errors"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

//IBMPIImageClient.
type IBMPIImageClient struct {
	IBMPIClient
}

// NewIBMPIImageClient.
func NewIBMPIImageClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIImageClient {
	return &IBMPIImageClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get an Image.
func (f *IBMPIImageClient) Get(id string) (*models.Image, error) {
	params := p.NewPcloudCloudinstancesImagesGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithImageID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.imageRequest.PcloudCloudinstancesImagesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetImageOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Get Image Operation for image %s", id)
	}
	return resp.Payload, nil
}

// Get All Images that are imported into Power Instance.
func (f *IBMPIImageClient) GetAll() (*models.Images, error) {
	params := p.NewPcloudCloudinstancesImagesGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.imageRequest.PcloudCloudinstancesImagesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PI Images of the PVM instance %s : %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Images of the PVM instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a Stock Image.
func (f *IBMPIImageClient) Create(body *models.CreateImage) (*models.Image, error) {
	if body.Source == nil || len(*body.Source) == 0 {
		body.Source = core.StringPtr("root-project")
	}
	params := p.NewPcloudCloudinstancesImagesPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	respok, respcreated, err := f.imageRequest.PcloudCloudinstancesImagesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateImageOperationFailed, f.cloudInstanceID, err)
	}
	if respok != nil && respok.Payload != nil {
		return respok.Payload, nil
	}
	if respcreated != nil && respcreated.Payload != nil {
		return respcreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to perform Create Image Operation for cloud instance %s", f.cloudInstanceID)
}

// Import an Image.
func (f *IBMPIImageClient) CreateCosImage(body *models.CreateCosImageImportJob) (imageJob *models.JobReference, err error) {
	params := p.NewPcloudV1CloudinstancesCosimagesPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.imageRequest.PcloudV1CloudinstancesCosimagesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to perform Create COS Image Operation for cloud instance %s with error %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Create COS Image Operation for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Export an Image.
func (f *IBMPIImageClient) ExportImage(id string, body *models.ExportImage) (*models.JobReference, error) {
	params := p.NewPcloudV2ImagesExportPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithImageID(id).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.imageRequest.PcloudV2ImagesExportPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Export COS Image for image id %s with error %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Export COS Image for image id %s", id)
	}
	return resp.Payload, nil
}

// Delete an Image.
func (f *IBMPIImageClient) Delete(id string) error {
	params := p.NewPcloudCloudinstancesImagesDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithImageID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.imageRequest.PcloudCloudinstancesImagesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PI Image %s: %w", id, err)
	}
	return nil
}

// Get a Stock Image.
func (f *IBMPIImageClient) GetStockImage(id string) (*models.Image, error) {
	params := p.NewPcloudCloudinstancesStockimagesGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithImageID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.imageRequest.PcloudCloudinstancesStockimagesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get PI Stock Image %s for cloud instance %s: %w", id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Stock Image %s for cloud instance %s", id, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Stock Images.
func (f *IBMPIImageClient) GetAllStockImages(includeSAP bool, includeVTl bool) (*models.Images, error) {
	params := p.NewPcloudCloudinstancesStockimagesGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithSap(&includeSAP).
		WithTimeout(helpers.PIGetTimeOut).
		WithVtl(&includeVTl)
	resp, err := f.imageRequest.PcloudCloudinstancesStockimagesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get Stock Images with (SAP=%t, VTL=%t) for cloud instance %s: %w", includeSAP, includeVTl, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get Stock Images with (SAP=%t, VTL=%t) for cloud instance %s", includeSAP, includeVTl, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Stock SAP Images.
func (f *IBMPIImageClient) GetAllStockSAPImages() (*models.Images, error) {
	// get stock images. include all available SAP images
	images, err := f.GetAllStockImages(true, false)
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

// Get All Stock VTL Images.
func (f *IBMPIImageClient) GetAllStockVTLImages() (*models.Images, error) {
	// get stock images. include all available stock VTL images
	images, err := f.GetAllStockImages(false, true)
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

// Return true if Image is a VTL Image.
func (f *IBMPIImageClient) IsVTLImage(imageId string) (bool, error) {
	images := new(models.Images)

	// get all stock vtl images
	stockVTLImages, err := f.GetAllStockVTLImages()
	if err != nil {
		return false, err
	}
	images.Images = append(images.Images, stockVTLImages.Images...)

	// get all images
	cloudInstanceImages, err := f.GetAll()
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
