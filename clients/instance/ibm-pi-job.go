package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_jobs"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIJobClient.
type IBMPIJobClient struct {
	IBMPIClient
}

// NewIBMPIJobClient.
func NewIBMPIJobClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIJobClient {
	return &IBMPIJobClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Job.
func (f *IBMPIJobClient) Get(id string) (*models.Job, error) {
	params := p.NewPcloudCloudinstancesJobsGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithJobID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.jobRequest.PcloudCloudinstancesJobsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetJobOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform get Job operation for job id %s", id)
	}
	return resp.Payload, nil
}

// Get All Jobs.
func (f *IBMPIJobClient) GetAll() (*models.Jobs, error) {
	params := p.NewPcloudCloudinstancesJobsGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.jobRequest.PcloudCloudinstancesJobsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetAllJobsOperationFailed, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform get all jobs")
	}
	return resp.Payload, nil
}

// Delete a Job.
func (f *IBMPIJobClient) Delete(id string) error {
	params := p.NewPcloudCloudinstancesJobsDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithJobID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.jobRequest.PcloudCloudinstancesJobsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteJobsOperationFailed, id, err)
	}
	return nil
}
