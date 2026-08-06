// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_async_jobs"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIAsyncJobClient
type IBMPIAsyncJobClient struct {
	IBMPIClient
}

// NewIBMPIAsyncJobClient
func NewIBMPIAsyncJobClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIAsyncJobClient {
	return &IBMPIAsyncJobClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get gets an asynchronous job of the cloud instance
func (f *IBMPIAsyncJobClient) Get(asyncJobID string) (*models.AsyncJob, error) {
	params := p_cloud_async_jobs.NewPcloudV1CloudinstancesAsyncjobsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).
		WithAsyncJobID(asyncJobID)
	resp, err := f.session.Power.PCloudAsyncJobs.PcloudV1CloudinstancesAsyncjobsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get PI Async Job %s: %w", asyncJobID, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Async Job %s in %s", asyncJobID, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// GetAll gets all asynchronous jobs of the cloud instance
func (f *IBMPIAsyncJobClient) GetAll() (*models.AsyncJobs, error) {
	params := p_cloud_async_jobs.NewPcloudV1CloudinstancesAsyncjobsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudAsyncJobs.PcloudV1CloudinstancesAsyncjobsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all PI Async Jobs: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Async Jobs in %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}
