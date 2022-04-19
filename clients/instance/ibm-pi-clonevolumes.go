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

// IBMPICloneVolumeClient ...
type IBMPICloneVolumeClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPICloneVolumeClient ...
func NewIBMPICloneVolumeClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPICloneVolumeClient {
	return &IBMPICloneVolumeClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudVolumes,
	}
}

// Create a Clone Volume (V2) - This creates a clone
func (f *IBMPICloneVolumeClient) Create(createBody *models.VolumesCloneAsyncRequest) (*models.CloneTaskReference, error) {

	// Create params and send request
	params := &params.PcloudV2VolumesClonePostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudV2VolumesClonePost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateCloneOperationFailed, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform the create clone operation")
	}
	return resp.Payload, nil
}

// Get Status of a Clone Request
func (f *IBMPICloneVolumeClient) Get(cloneTaskID string) (*models.CloneTaskStatus, error) {

	// Create params and send request
	params := &params.PcloudV2VolumesClonetasksGetParams{
		CloneTaskID:     cloneTaskID,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudV2VolumesClonetasksGet(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to get the clone task %s status for the cloud instance %s with error %w", cloneTaskID, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the clone task %s status for the cloud instance %s", cloneTaskID, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a Clone Volume (V2) - This is the prepare operation
func (f *IBMPICloneVolumeClient) CreateV2Clone(createBody *models.VolumesCloneCreate) (*models.VolumesClone, error) {

	// Create params and send request
	params := &params.PcloudV2VolumesclonePostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudV2VolumesclonePost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf(errors.PrepareCloneOperationFailed, *createBody.Name, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to prepare the clone operation")
	}
	return resp.Payload, nil
}

// Get a list of Volume Clones
func (f *IBMPICloneVolumeClient) GetV2Clones(queryFilter string) (*models.VolumesClones, error) {

	// Create params and send request
	params := &params.PcloudV2VolumescloneGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		Filter:          &queryFilter,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudV2VolumescloneGetall(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to get the volumes-clones for the cloud instance %s with error %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the volumes-clones for the cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Delete a Volume Clone
func (f *IBMPICloneVolumeClient) DeleteClone(volumeCloneID string) error {

	// Create params and send request
	params := &params.PcloudV2VolumescloneDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VolumesCloneID:  volumeCloneID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudV2VolumescloneDelete(params, f.auth)

	// Handle Errors
	if err != nil {
		return fmt.Errorf(errors.DeleteCloneOperationFailed, err)
	}
	return nil
}

// Initiate a Start Clone Request
func (f *IBMPICloneVolumeClient) StartClone(volumeCloneID string) (*models.VolumesClone, error) {

	// Create params and send request
	params := &params.PcloudV2VolumescloneStartPostParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VolumesCloneID:  volumeCloneID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudV2VolumescloneStartPost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf(errors.StartCloneOperationFailed, volumeCloneID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to start the clone operation for volume-clone %s", volumeCloneID)
	}
	return resp.Payload, nil
}

// Initiate the Execute Action for a Clone
func (f *IBMPICloneVolumeClient) PrepareClone(volumeCloneID string) (*models.VolumesClone, error) {

	// Create params and send request
	params := &params.PcloudV2VolumescloneExecutePostParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VolumesCloneID:  volumeCloneID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudV2VolumescloneExecutePost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf(errors.PrepareCloneOperationFailed, volumeCloneID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to prepare the clone operation for %s", volumeCloneID)
	}
	return resp.Payload, nil
}

// Get a V2Clone Task Status
func (f *IBMPICloneVolumeClient) GetV2CloneStatus(volumeCloneID string) (*models.VolumesCloneDetail, error) {

	// Create params and send request
	params := &params.PcloudV2VolumescloneGetParams{
		CloudInstanceID: f.cloudInstanceID,
		VolumesCloneID:  volumeCloneID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudV2VolumescloneGet(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetCloneOperationFailed, volumeCloneID, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the volumes-clone %s for the cloud instance %s", volumeCloneID, f.cloudInstanceID)
	}
	return resp.Payload, nil
}
