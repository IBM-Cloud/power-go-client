package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	s "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_snapshots"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPISnapshotClient.
type IBMPISnapshotClient struct {
	IBMPIClient
}

// NewIBMPISnapshotClient.
func NewIBMPISnapshotClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPISnapshotClient {
	return &IBMPISnapshotClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Snapshot.
func (f *IBMPISnapshotClient) Get(id string) (*models.Snapshot, error) {
	params := s.NewPcloudCloudinstancesSnapshotsGetParams().
		WithContext(f.ctx).
		WithCloudInstanceID(f.cloudInstanceID).
		WithSnapshotID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.snapshotRequest.PcloudCloudinstancesSnapshotsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get PI Snapshot %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Snapshot %s", id)
	}
	return resp.Payload, nil
}

// Delete a Snapshot.
func (f *IBMPISnapshotClient) Delete(id string) error {
	params := s.NewPcloudCloudinstancesSnapshotsDeleteParams().
		WithContext(f.ctx).
		WithCloudInstanceID(f.cloudInstanceID).
		WithSnapshotID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.snapshotRequest.PcloudCloudinstancesSnapshotsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PI Snapshot %s: %w", id, err)
	}
	return nil
}

// Update a Snapshot.
func (f *IBMPISnapshotClient) Update(id string, body *models.SnapshotUpdate) (models.Object, error) {
	params := s.NewPcloudCloudinstancesSnapshotsPutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithSnapshotID(id).
		WithTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.snapshotRequest.PcloudCloudinstancesSnapshotsPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Update PI Snapshot %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update PI Snapshot %s", id)
	}
	return resp.Payload, nil
}

// Get All Snapshots.
func (f *IBMPISnapshotClient) GetAll() (*models.Snapshots, error) {
	params := s.NewPcloudCloudinstancesSnapshotsGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.snapshotRequest.PcloudCloudinstancesSnapshotsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PI Snapshots: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Snapshots")
	}
	return resp.Payload, nil
}

// Create or Restore a Snapshot.
func (f *IBMPISnapshotClient) Create(instanceID, snapshotID, restoreFailAction string) (*models.Snapshot, error) {
	params := p.NewPcloudPvminstancesSnapshotsRestorePostParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(instanceID).
		WithRestoreFailAction(&restoreFailAction).
		WithSnapshotID(snapshotID).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesSnapshotsRestorePost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to restore PI Snapshot %s of the instance %s: %w", snapshotID, instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to restore PI Snapshot %s of the instance %s", snapshotID, instanceID)
	}
	return resp.Payload, nil
}
