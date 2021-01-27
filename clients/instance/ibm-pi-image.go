package instance

import (
	"fmt"

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

	params := p_cloud_images.NewPcloudCloudinstancesImagesGetParamsWithTimeout(postTimeOut).WithCloudInstanceID(powerinstanceid).WithImageID(id)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get PI Image %s :%s", id, err)
	}
	return resp.Payload, nil
}

//GetAll Images that are imported into Power Instance
func (f *IBMPIImageClient) GetAll(powerinstanceid string) (*models.Images, error) {

	params := p_cloud_images.NewPcloudCloudinstancesImagesGetallParamsWithTimeout(getTimeOut).WithCloudInstanceID(powerinstanceid)
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
	params := p_cloud_images.NewPcloudCloudinstancesImagesPostParamsWithTimeout(postTimeOut).WithCloudInstanceID(powerinstanceid).WithBody(&body)
	_, result, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || result == nil || result.Payload == nil {
		return nil, fmt.Errorf("Failed to Create Image of the PVM instance %s : %s", powerinstanceid, err)
	}
	return result.Payload, nil

}

// Delete ...
func (f *IBMPIImageClient) Delete(id string, powerinstanceid string) error {
	params := p_cloud_images.NewPcloudCloudinstancesImagesDeleteParamsWithTimeout(deleteTimeOut).WithCloudInstanceID(powerinstanceid).WithImageID(id)
	_, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return fmt.Errorf("Failed to Delete PI Image %s :%s", id, err)
	}
	return nil
}

// GetStockImages ...
func (f *IBMPIImageClient) GetStockImages(powerinstanceid string) (*models.Images, error) {

	params := p_cloud_images.NewPcloudImagesGetallParams()
	resp, err := f.session.Power.PCloudImages.PcloudImagesGetall(params, ibmpisession.NewAuth(f.session, f.powerinstanceid))

	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get all PI Stock Images of the PVM instance %s : %s", powerinstanceid, err)
	}
	return resp.Payload, nil
}

//GetSAPImages ...
func (f *IBMPIImageClient) GetSAPImages(powerinstanceid string, sapimage bool) (*models.Images, error) {

	params := p_cloud_images.NewPcloudImagesGetallParams()
	params.Sap = &sapimage

	resp, err := f.session.Power.PCloudImages.PcloudImagesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get all PI Sap Images of the PVM instance %s : %s", powerinstanceid, err)
	}
	return resp.Payload, nil
}

// Get a single SAP Image
