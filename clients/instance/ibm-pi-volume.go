package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIVolumeClient.
type IBMPIVolumeClient struct {
	IBMPIClient
}

// NewIBMPIVolumeClient.
func NewIBMPIVolumeClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVolumeClient {
	return &IBMPIVolumeClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Volume.
func (f *IBMPIVolumeClient) Get(id string) (*models.Volume, error) {
	params := p.NewPcloudCloudinstancesVolumesGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithVolumeID(id)
	resp, err := f.volumeRequest.PcloudCloudinstancesVolumesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVolumeOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Volume %s", id)
	}
	return resp.Payload, nil
}

// Get All Volumes.
func (f *IBMPIVolumeClient) GetAll() (*models.Volumes, error) {
	params := p.NewPcloudCloudinstancesVolumesGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.volumeRequest.PcloudCloudinstancesVolumesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes for Cloud Instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes for Cloud Instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Affinity Volumes.
func (f *IBMPIVolumeClient) GetAllAffinityVolumes(affinity string) (*models.Volumes, error) {
	params := p.NewPcloudCloudinstancesVolumesGetallParams().
		WithAffinity(&affinity).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.volumeRequest.PcloudCloudinstancesVolumesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes with affinity %s for Cloud Instance %s: %w", affinity, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes with affinity %s for Cloud Instance %s", affinity, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a VolumeV2.
func (f *IBMPIVolumeClient) CreateVolumeV2(body *models.MultiVolumesCreate) (*models.Volumes, error) {
	params := p.NewPcloudV2VolumesPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.volumeRequest.PcloudV2VolumesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVolumeV2OperationFailed, *body.Name, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create Volume v2")
	}
	return resp.Payload, nil
}

// Create a Volume.
func (f *IBMPIVolumeClient) CreateVolume(body *models.CreateDataVolume) (*models.Volume, error) {
	params := p.NewPcloudCloudinstancesVolumesPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.volumeRequest.PcloudCloudinstancesVolumesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVolumeOperationFailed, *body.Name, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create Volume")
	}
	return resp.Payload, nil
}

// Update a Volume.
func (f *IBMPIVolumeClient) UpdateVolume(id string, body *models.UpdateVolume) (*models.Volume, error) {
	params := p.NewPcloudCloudinstancesVolumesPutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithVolumeID(id)
	resp, err := f.volumeRequest.PcloudCloudinstancesVolumesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVolumeOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update Volume %s", id)
	}
	return resp.Payload, nil
}

// Delete a Volume.
func (f *IBMPIVolumeClient) DeleteVolume(id string) error {
	params := p.NewPcloudCloudinstancesVolumesDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithVolumeID(id)
	_, err := f.volumeRequest.PcloudCloudinstancesVolumesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteVolumeOperationFailed, id, err)
	}
	return nil
}

// Attach a Volume to an Instance.
func (f *IBMPIVolumeClient) Attach(id, volumename string) error {
	params := p.NewPcloudPvminstancesVolumesPostParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut).
		WithVolumeID(volumename)
	_, err := f.volumeRequest.PcloudPvminstancesVolumesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.AttachVolumeOperationFailed, volumename, err)
	}
	return nil
}

// Detach a Volume from an Instance.
func (f *IBMPIVolumeClient) Detach(id, volumename string) error {
	params := p.NewPcloudPvminstancesVolumesDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut).
		WithVolumeID(volumename)
	_, err := f.volumeRequest.PcloudPvminstancesVolumesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DetachVolumeOperationFailed, volumename, err)
	}
	return nil
}

// Get All Volumes attached to an Instance.
func (f *IBMPIVolumeClient) GetAllInstanceVolumes(id string) (*models.Volumes, error) {
	params := p.NewPcloudPvminstancesVolumesGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.volumeRequest.PcloudPvminstancesVolumesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes for PI Instance %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes for PI Instance %s", id)
	}
	return resp.Payload, nil
}

// Set a Volume as the Boot Volume for an Instance.
func (f *IBMPIVolumeClient) SetBootVolume(id, volumename string) error {
	params := p.NewPcloudPvminstancesVolumesSetbootPutParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut).
		WithVolumeID(volumename)
	_, err := f.volumeRequest.PcloudPvminstancesVolumesSetbootPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to set the boot volume %s for instance %s", volumename, id)
	}
	return nil
}

// Check if a Volume is attached to an Instance.
func (f *IBMPIVolumeClient) CheckVolumeAttach(id, volumeID string) (*models.Volume, error) {
	params := p.NewPcloudPvminstancesVolumesGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIGetTimeOut).
		WithVolumeID(volumeID)
	resp, err := f.volumeRequest.PcloudPvminstancesVolumesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s: %w", volumeID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s", volumeID, id)
	}
	return resp.Payload, nil
}

// Update a Volume attached to an Instance.
func (f *IBMPIVolumeClient) UpdateVolumeAttach(id, volumeID string, body *models.PVMInstanceVolumeUpdate) error {
	params := p.NewPcloudPvminstancesVolumesPutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithVolumeID(volumeID)
	resp, err := f.volumeRequest.PcloudPvminstancesVolumesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s: %w", volumeID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s", volumeID, id)
	}
	return nil
}
