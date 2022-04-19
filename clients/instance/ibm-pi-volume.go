package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPIVolumeClient ..
type IBMPIVolumeClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPIVolumeClient ...
func NewIBMPIVolumeClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVolumeClient {
	return &IBMPIVolumeClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudVolumes,
	}
}

// Get a Volume
func (f *IBMPIVolumeClient) Get(volumeID string) (*models.Volume, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesVolumesGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VolumeID:        volumeID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudinstancesVolumesGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetVolumeOperationFailed, volumeID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Volume %s", volumeID)
	}
	return resp.Payload, nil
}

// Get All Volumes
func (f *IBMPIVolumeClient) GetAll() (*models.Volumes, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesVolumesGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudinstancesVolumesGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes for Cloud Instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes for Cloud Instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Affinity Volumes
func (f *IBMPIVolumeClient) GetAllAffinityVolumes(affinity string) (*models.Volumes, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesVolumesGetallParams{
		Affinity:        &affinity,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudinstancesVolumesGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes with affinity %s for Cloud Instance %s: %w", affinity, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes with affinity %s for Cloud Instance %s", affinity, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a VolumeV2
func (f *IBMPIVolumeClient) CreateVolumeV2(createBody *models.MultiVolumesCreate) (*models.Volumes, error) {

	// Create params and send request
	params := &params.PcloudV2VolumesPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudV2VolumesPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVolumeV2OperationFailed, *createBody.Name, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create Volume v2")
	}
	return resp.Payload, nil
}

// Create a Volume
func (f *IBMPIVolumeClient) CreateVolume(createBody *models.CreateDataVolume) (*models.Volume, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesVolumesPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudCloudinstancesVolumesPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVolumeOperationFailed, *createBody.Name, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create Volume")
	}
	return resp.Payload, nil
}

// Update a Volume
func (f *IBMPIVolumeClient) UpdateVolume(volumeID string, updateBody *models.UpdateVolume) (*models.Volume, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesVolumesPutParams{
		Body:            updateBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VolumeID:        volumeID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudCloudinstancesVolumesPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVolumeOperationFailed, volumeID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update Volume %s", volumeID)
	}
	return resp.Payload, nil
}

// Delete a Volume
func (f *IBMPIVolumeClient) DeleteVolume(volumeID string) error {

	// Create params and send request
	params := &params.PcloudCloudinstancesVolumesDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VolumeID:        volumeID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudCloudinstancesVolumesDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.DeleteVolumeOperationFailed, volumeID, err)
	}
	return nil
}

// Attach a Volume to an Instance
func (f *IBMPIVolumeClient) Attach(instanceID, volumeID string) error {

	// Create params and send request
	params := &params.PcloudPvminstancesVolumesPostParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
		VolumeID:        volumeID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	_, err := f.request.PcloudPvminstancesVolumesPost(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.AttachVolumeOperationFailed, volumeID, err)
	}
	return nil
}

// Detach a Volume from an Instance
func (f *IBMPIVolumeClient) Detach(instanceID, volumeID string) error {

	// Create params and send request
	params := &params.PcloudPvminstancesVolumesDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
		VolumeID:        volumeID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	_, err := f.request.PcloudPvminstancesVolumesDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.DetachVolumeOperationFailed, volumeID, err)
	}
	return nil
}

// Get All Volumes attached to an Instance
func (f *IBMPIVolumeClient) GetAllInstanceVolumes(instanceID string) (*models.Volumes, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesVolumesGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudPvminstancesVolumesGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes for PI Instance %s: %w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes for PI Instance %s", instanceID)
	}
	return resp.Payload, nil
}

// Set a Volume as the Boot Volume for an Instance
func (f *IBMPIVolumeClient) SetBootVolume(instanceID, volumeID string) error {

	// Create params and send request
	params := &params.PcloudPvminstancesVolumesSetbootPutParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
		VolumeID:        volumeID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	_, err := f.request.PcloudPvminstancesVolumesSetbootPut(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf("failed to set the boot volume %s for instance %s", volumeID, instanceID)
	}
	return nil
}

// Check if a Volume is attached to an Instance
func (f *IBMPIVolumeClient) CheckVolumeAttach(instanceID, volumeID string) (*models.Volume, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesVolumesGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
		VolumeID:        volumeID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudPvminstancesVolumesGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s: %w", volumeID, instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s", volumeID, instanceID)
	}
	return resp.Payload, nil
}

// Update a Volume attached to an Instance
func (f *IBMPIVolumeClient) UpdateVolumeAttach(instanceID, volumeID string, body *models.PVMInstanceVolumeUpdate) error {

	// Create params and send request
	params := &params.PcloudPvminstancesVolumesPutParams{
		Body:            body,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
		VolumeID:        volumeID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudPvminstancesVolumesPut(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s: %w", volumeID, instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s", volumeID, instanceID)
	}
	return nil
}
