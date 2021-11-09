package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_jobs"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIJobClient ...
type IBMPIJobClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIJobClient ...
func NewIBMPIJobClient(sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIJobClient {
	return &IBMPIJobClient{
		session:         sess,
		powerinstanceid: cloudInstanceID,
	}
}

// Get information about a job
func (f *IBMPIJobClient) Get(id, cloudInstanceID string) (*models.Job, error) {
	params := p_cloud_jobs.NewPcloudCloudinstancesJobsGetParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(cloudInstanceID).WithJobID(id)
	resp, err := f.session.Power.PCloudJobs.PcloudCloudinstancesJobsGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetJobOperationFailed, id, err)
	}
	return resp.Payload, nil
}

// Get information about a job with Context
func (f *IBMPIJobClient) GetWithContext(ctx context.Context, id, cloudInstanceID string) (*models.Job, error) {
	params := p_cloud_jobs.NewPcloudCloudinstancesJobsGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithJobID(id)
	resp, err := f.session.Power.PCloudJobs.PcloudCloudinstancesJobsGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetJobOperationFailed, id, err)
	}
	return resp.Payload, nil
}

// Update a job
func (f *IBMPIJobClient) GetAll(cloudInstanceID string) (*models.Jobs, error) {
	params := p_cloud_jobs.NewPcloudCloudinstancesJobsGetallParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(cloudInstanceID)
	resp, err := f.session.Power.PCloudJobs.PcloudCloudinstancesJobsGetall(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetAllJobsOperationFailed, err)

	}
	return resp.Payload, nil
}

// Delete a job
func (f *IBMPIJobClient) Delete(id, cloudInstanceID string) error {
	params := p_cloud_jobs.NewPcloudCloudinstancesJobsDeleteParamsWithTimeout(helpers.PIDeleteTimeOut).WithCloudInstanceID(cloudInstanceID).WithJobID(id)
	_, err := f.session.Power.PCloudJobs.PcloudCloudinstancesJobsDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteJobsOperationFailed, id, err)
	}
	return nil
}
