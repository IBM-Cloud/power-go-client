package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	instanceParams "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	snapshotParams "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_snapshots"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPISnapshotClient ...
type IBMPISnapshotClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	instanceRequest instanceParams.ClientService
	snapshotRequest snapshotParams.ClientService
}

// NewIBMPISnapshotClient ...
func NewIBMPISnapshotClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPISnapshotClient {
	return &IBMPISnapshotClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		instanceRequest: sess.Power.PCloudpVMInstances,
		snapshotRequest: sess.Power.PCloudSnapshots,
	}
}

// Get a Snapshot
func (f *IBMPISnapshotClient) Get(snapshotID string) (*models.Snapshot, error) {

	// Create params and send request
	params := &snapshotParams.PcloudCloudinstancesSnapshotsGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		SnapshotID:      snapshotID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.snapshotRequest.PcloudCloudinstancesSnapshotsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get PI Snapshot %s: %w", snapshotID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Snapshot %s", snapshotID)
	}
	return resp.Payload, nil
}

// Delete a Snapshot
func (f *IBMPISnapshotClient) Delete(snapshotID string) error {

	// Create params and send request
	params := &snapshotParams.PcloudCloudinstancesSnapshotsDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		SnapshotID:      snapshotID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.snapshotRequest.PcloudCloudinstancesSnapshotsDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf("failed to Delete PI Snapshot %s: %w", snapshotID, err)
	}
	return nil
}

// Update a Snapshot
func (f *IBMPISnapshotClient) Update(snapshotID string, updateBody *models.SnapshotUpdate) (models.Object, error) {

	// Create params and send request
	params := &snapshotParams.PcloudCloudinstancesSnapshotsPutParams{
		Body:            updateBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		SnapshotID:      snapshotID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.snapshotRequest.PcloudCloudinstancesSnapshotsPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Update PI Snapshot %s: %w", snapshotID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update PI Snapshot %s", snapshotID)
	}
	return resp.Payload, nil
}

// Get All Snapshots
func (f *IBMPISnapshotClient) GetAll() (*models.Snapshots, error) {

	// Create params and send request
	params := &snapshotParams.PcloudCloudinstancesSnapshotsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.snapshotRequest.PcloudCloudinstancesSnapshotsGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PI Snapshots: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Snapshots")
	}
	return resp.Payload, nil
}

// Create or Restore a Snapshot
func (f *IBMPISnapshotClient) Create(instanceID, snapshotID, restoreFailAction string) (*models.Snapshot, error) {

	// Create params and send request
	params := &instanceParams.PcloudPvminstancesSnapshotsRestorePostParams{
		CloudInstanceID:   f.cloudInstanceID,
		Context:           f.context,
		PvmInstanceID:     instanceID,
		RestoreFailAction: &restoreFailAction,
		SnapshotID:        snapshotID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesSnapshotsRestorePost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to restore PI Snapshot %s of the instance %s: %w", snapshotID, instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to restore PI Snapshot %s of the instance %s", snapshotID, instanceID)
	}
	return resp.Payload, nil
}
