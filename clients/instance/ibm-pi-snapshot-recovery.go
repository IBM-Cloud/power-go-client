// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_snapshot_recovery"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPISnapshotRecoveryClient
type IBMPISnapshotRecoveryClient struct {
	IBMPIClient
}

// NewIBMPISnapshotRecoveryClient
func NewIBMPISnapshotRecoveryClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPISnapshotRecoveryClient {
	return &IBMPISnapshotRecoveryClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get gets the snapshot recovery site details for the current workspace location
func (f *IBMPISnapshotRecoveryClient) Get() (*models.SnapshotRecoveryLocation, error) {
	params := p_cloud_snapshot_recovery.NewPcloudLocationsSnapshotrecoveryGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudSnapshotRecovery.PcloudLocationsSnapshotrecoveryGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get PI Snapshot Recovery location: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Snapshot Recovery location")
	}
	return resp.Payload, nil
}

// GetAll gets all snapshot recovery locations supported by Power Virtual Server
func (f *IBMPISnapshotRecoveryClient) GetAll() (*models.SnapshotRecoveryLocations, error) {
	params := p_cloud_snapshot_recovery.NewPcloudLocationsSnapshotrecoveryGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.PCloudSnapshotRecovery.PcloudLocationsSnapshotrecoveryGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all PI Snapshot Recovery locations: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Snapshot Recovery locations")
	}
	return resp.Payload, nil
}
