package instance

import (
	"fmt"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
	"time"
)

type IBMPIVolumeClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewPowerVolumeClient ...
func NewIBMPIVolumeClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIVolumeClient {
	return &IBMPIVolumeClient{
		sess, powerinstanceid,
	}
}

//Get information about a single volume only
func (f *IBMPIVolumeClient) Get(id, powerinstanceid string, timeout time.Duration) (*models.Volume, error) {

	log.Printf("Calling the VolumeGet Method..")
	log.Printf("The input volume name is %s and  to the cloudinstance id %s", id, powerinstanceid)

	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesGetParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithVolumeID(id)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//V2 Create Volume

func (f *IBMPIVolumeClient) CreateVolumeV2(create_vol_defs *p_cloud_volumes.PcloudV2VolumesPostParams, powerinstanceid string, timeout time.Duration) (*models.Volumes, error) {

	log.Printf("Calling the Create Volume v2 method")
	params := p_cloud_volumes.NewPcloudV2VolumesPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(create_vol_defs.Body)
	resp, err := f.session.Power.PCloudVolumes.PcloudV2VolumesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// CreateVolume
func (f *IBMPIVolumeClient) CreateVolume(create_vol_defs *p_cloud_volumes.PcloudCloudinstancesVolumesPostParams, powerinstanceid string, timeout time.Duration) (*models.Volume, error) {

	log.Printf("Calling the Create Volume method ")
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(create_vol_defs.Body)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// UpdateVolume

func (f *IBMPIVolumeClient) UpdateVolume(update_vol_defs *p_cloud_volumes.PcloudCloudinstancesVolumesPutParams, volumeid, powerinstanceid string, timeout time.Duration) (*models.Volume, error) {

	log.Printf("Calling the Create Volume method ")
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPutParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(update_vol_defs.Body).WithVolumeID(volumeid)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteVolume
func (f *IBMPIVolumeClient) DeleteVolume(id string, powerinstanceid string, timeout time.Duration) error {

	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesDeleteParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithVolumeID(id)
	_, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

//Create
// TO be Deprecated

func (f *IBMPIVolumeClient) Create(volumename string, volumesize float64, volumetype string, volumeshareable bool, powerinstanceid string, timeout time.Duration) (*models.Volume, error) {

	log.Printf("calling the PowerVolume Create Method")

	var body = models.CreateDataVolume{
		Name:      &volumename,
		Size:      &volumesize,
		DiskType:  volumetype,
		Shareable: &volumeshareable,
	}

	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(&body)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *IBMPIVolumeClient) Delete(id string, powerinstanceid string, timeout time.Duration) error {
	//var cloudinstanceid = f.session.PowerServiceInstance
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesDeleteParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithVolumeID(id)
	_, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// Update..
func (f *IBMPIVolumeClient) Update(id, volumename string, volumesize float64, volumeshare bool, powerinstanceid string, timeout time.Duration) (*models.Volume, error) {

	var patchbody = models.UpdateVolume{}
	patchbody.Name = &volumename
	patchbody.Size = volumesize
	patchbody.Shareable = &volumeshare
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPutParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithVolumeID(id).WithBody(&patchbody)

	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Attach a volume

func (f *IBMPIVolumeClient) Attach(id, volumename string, powerinstanceid string, timeout time.Duration) (models.Object, error) {

	log.Printf("Calling the Power Volume Attach method")

	params := p_cloud_volumes.NewPcloudPvminstancesVolumesPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithVolumeID(volumename)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, errors.ToError(err)
	}
	log.Printf("Successfully attached the volume to the instance")

	return resp.Payload, nil

}

//Detach a volume

func (f *IBMPIVolumeClient) Detach(id, volumename string, powerinstanceid string, timeout time.Duration) (models.Object, error) {
	log.Printf("Calling the Power Volume Detach method")

	params := p_cloud_volumes.NewPcloudPvminstancesVolumesDeleteParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithVolumeID(volumename)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		//return nil, errors.ToError(err)
		return nil, fmt.Errorf("Failed to detach the volume [%s ] for pvm instance with id [%s] ", volumename, id)
	}
	return resp.Payload, nil

}

// All volumes part of an instance

func (f *IBMPIVolumeClient) GetAll(id, cloud_instance_id string, timeout time.Duration) (*models.Volumes, error) {

	log.Printf("Calling the Power Volumes GetAll Method")
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesGetallParamsWithTimeout(timeout).WithPvmInstanceID(id).WithCloudInstanceID(cloud_instance_id)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesGetall(params, ibmpisession.NewAuth(f.session, cloud_instance_id))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}

// Set a volume as the boot volume - PUT Operation

func (f *IBMPIVolumeClient) SetBootVolume(id, volumename, cloud_instance_id string, timeout time.Duration) (models.Object, error) {
	log.Printf("Setting the Boot Volume for this %s instance as ", cloud_instance_id)
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesSetbootPutParamsWithTimeout(timeout).WithCloudInstanceID(cloud_instance_id).WithPvmInstanceID(id).WithVolumeID(volumename)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesSetbootPut(params, ibmpisession.NewAuth(f.session, cloud_instance_id))
	if err != nil {
		//return nil, errors.ToError(err)
		return nil, fmt.Errorf("Failed to set the boot volume for cloud instance id [%s] ", cloud_instance_id)
	}
	return resp.Payload, nil
}

// Check if the volume is attached to the instance
func (f *IBMPIVolumeClient) CheckVolumeAttach(cloud_instance_id, pvm_instance_id, volume_id string, timeout time.Duration) (*models.Volume, error) {

	log.Printf("Checking if the volume [%s] has been attached to the pvm_instance [%s] for cloud instance id [%s]", volume_id, pvm_instance_id, cloud_instance_id)
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesGetParamsWithTimeout(timeout).WithCloudInstanceID(cloud_instance_id).WithPvmInstanceID(pvm_instance_id).WithVolumeID(volume_id)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesGet(params, ibmpisession.NewAuth(f.session, cloud_instance_id))

	if err != nil {
		return nil, fmt.Errorf("Failed to validate that the volume [%s] is attached to the pvminstance [%s]", volume_id, pvm_instance_id)
	}

	return resp.Payload, nil
}
