package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_jobs"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPIJobClient ...
type IBMPIJobClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPIJobClient ...
func NewIBMPIJobClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIJobClient {
	return &IBMPIJobClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudJobs,
	}
}

// Get a Job
func (f *IBMPIJobClient) Get(jobId string) (*models.Job, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesJobsGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		JobID:           jobId,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudinstancesJobsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetJobOperationFailed, jobId, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform get Job operation for job id %s", jobId)
	}
	return resp.Payload, nil
}

// Get All Jobs
func (f *IBMPIJobClient) GetAll() (*models.Jobs, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesJobsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudinstancesJobsGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetAllJobsOperationFailed, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform get all jobs")
	}
	return resp.Payload, nil
}

// Delete a Job
func (f *IBMPIJobClient) Delete(jobId string) error {

	// Create params and send request
	params := &params.PcloudCloudinstancesJobsDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		JobID:           jobId,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	_, err := f.request.PcloudCloudinstancesJobsDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.DeleteJobsOperationFailed, jobId, err)
	}
	return nil
}
