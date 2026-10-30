package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_snapshots"
	"github.com/IBM-Cloud/power-go-client/power/client/snapshots"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPISnapshotClient
type IBMPISnapshotClient struct {
	IBMPIClient
}

// NewIBMPISnapshotClient
func NewIBMPISnapshotClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPISnapshotClient {
	return &IBMPISnapshotClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Snapshot
func (f *IBMPISnapshotClient) Get(id string) (*models.Snapshot, error) {
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsGetParams().
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSnapshotID(id)
	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsGetContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get PI Snapshot %s: %w", id, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Snapshot %s", id)
	}
	return resp.Payload, nil
}

// Delete a Snapshot
func (f *IBMPISnapshotClient) Delete(id string) error {
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsDeleteParams().
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSnapshotID(id)
	_, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsDeleteContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PI Snapshot %s: %w", id, err)
	}
	return nil
}

// Update a Snapshot
func (f *IBMPISnapshotClient) Update(id string, body *models.SnapshotUpdate) (models.Object, error) {
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsPutParams().
		WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSnapshotID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsPutContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Update PI Snapshot %s: %w", id, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update PI Snapshot %s", id)
	}
	return resp.Payload, nil
}

// Get All Snapshots
func (f *IBMPISnapshotClient) GetAll() (*models.Snapshots, error) {
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsGetallParams().
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsGetallContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all PI Snapshots: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Snapshots")
	}
	return resp.Payload, nil
}

// Create or Restore a Snapshot
func (f *IBMPISnapshotClient) Create(instanceID, snapshotID, restoreFailAction string) (*models.Snapshot, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsRestorePostParams().
		WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(instanceID).
		WithSnapshotID(snapshotID).WithRestoreFailAction(&restoreFailAction)
	resp, err := f.session.Power.PCloudpVMInstances.PcloudPvminstancesSnapshotsRestorePostContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to restore PI Snapshot %s of the instance %s: %w", snapshotID, instanceID, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to restore PI Snapshot %s of the instance %s", snapshotID, instanceID)
	}
	return resp.Payload, nil
}

// Get a Volume SnapshotV1
func (f *IBMPISnapshotClient) V1VolumeSnapshotsGet(id string) (*models.SnapshotV1, error) {
	params := snapshots.NewV1VolumeSnapshotsGetParams().WithTimeout(helpers.PIGetTimeOut).WithVolumeSnapshotUUID(id)
	resp, err := f.session.Power.Snapshots.V1VolumeSnapshotsGetContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))

	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get volume snapshot %s: %w", id, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Snapshot %s", id)
	}
	return resp.Payload, nil
}

// Get Volume All SnapshotsV1
func (f *IBMPISnapshotClient) V1VolumeSnapshotsGetall() (*models.VolumeSnapshotList, error) {
	params := snapshots.NewV1VolumeSnapshotsGetallParams().WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.Snapshots.V1VolumeSnapshotsGetallContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))

	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all volume snapshots: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all volume Snapshots")
	}
	return resp.Payload, nil
}

// Deprecated Get a SnapshotV1
func (f *IBMPISnapshotClient) V1SnapshotsGet(id string) (*models.SnapshotV1, error) {
	params := snapshots.NewV1SnapshotsGetParams().WithTimeout(helpers.PIGetTimeOut).WithSnapshotID(id)
	resp, err := f.session.Power.Snapshots.V1SnapshotsGetContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))

	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get PI Snapshot %s: %w", id, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Snapshot %s", id)
	}
	return resp.Payload, nil
}

// Deprecated Get All SnapshotsV1
func (f *IBMPISnapshotClient) V1SnapshotsGetall() (*models.SnapshotList, error) {
	params := snapshots.NewV1SnapshotsGetallParams().WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.Snapshots.V1SnapshotsGetallContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))

	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all PI Snapshots: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Snapshots")
	}
	return resp.Payload, nil
}

// V2Create creates a V2 PVM instance snapshot
func (f *IBMPISnapshotClient) V2Create(pvmInstanceID string, body *models.SnapshotCreateV2) (*models.SnapshotCreateResponse, error) {
	params := p_cloud_p_vm_instances.NewPcloudV2PvminstancesSnapshotsPostParams().
		WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(pvmInstanceID).
		WithBody(body)
	resp, err := f.session.Power.PCloudpVMInstances.PcloudV2PvminstancesSnapshotsPostContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Create V2 PI Snapshot for instance %s: %w", pvmInstanceID, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create V2 PI Snapshot for instance %s", pvmInstanceID)
	}
	return resp.Payload, nil
}

// V2Get gets a V2 PVM instance snapshot by snapshot ID
func (f *IBMPISnapshotClient) V2Get(snapshotID string) (*models.SnapshotV2, error) {
	params := p_cloud_snapshots.NewPcloudV2CloudinstancesSnapshotsGetParams().
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSnapshotID(snapshotID)
	resp, err := f.session.Power.PCloudSnapshots.PcloudV2CloudinstancesSnapshotsGetContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get V2 PI Snapshot %s: %w", snapshotID, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get V2 PI Snapshot %s", snapshotID)
	}
	return resp.Payload, nil
}

// V2GetAll gets all V2 PVM instance snapshots for a cloud instance
func (f *IBMPISnapshotClient) V2GetAll() (*models.SnapshotsV2, error) {
	params := p_cloud_snapshots.NewPcloudV2CloudinstancesSnapshotsGetallParams().
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudSnapshots.PcloudV2CloudinstancesSnapshotsGetallContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all V2 PI Snapshots: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all V2 PI Snapshots")
	}
	return resp.Payload, nil
}

// V2PVMInstanceGetAll gets all V2 snapshots scoped to a specific PVM instance
func (f *IBMPISnapshotClient) V2PVMInstanceGetAll(pvmInstanceID string) (*models.SnapshotsV2, error) {
	params := p_cloud_p_vm_instances.NewPcloudV2PvminstancesSnapshotsGetallParams().
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(pvmInstanceID)
	resp, err := f.session.Power.PCloudpVMInstances.PcloudV2PvminstancesSnapshotsGetallContext(f.ctx, params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all V2 PI Snapshots for instance %s: %w", pvmInstanceID, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all V2 PI Snapshots for instance %s", pvmInstanceID)
	}
	return resp.Payload, nil
}
