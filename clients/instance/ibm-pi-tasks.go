package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tasks"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPITaskClient ...
type IBMPITaskClient struct {
	auth    runtime.ClientAuthInfoWriter
	context context.Context
	request params.ClientService
}

// NewIBMPITaskClient ...
func NewIBMPITaskClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPITaskClient {
	return &IBMPITaskClient{
		auth:    sess.AuthInfo(cloudInstanceID),
		context: ctx,
		request: sess.Power.PCloudTasks,
	}
}

// Get a Task
func (f *IBMPITaskClient) Get(taskID string) (*models.Task, error) {

	// Create params and send request
	params := &params.PcloudTasksGetParams{
		Context: f.context,
		TaskID:  taskID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudTasksGet(params, f.auth)

	// Handle errors
	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the task %s: %w", taskID, err)
	}
	return resp.Payload, nil
}

// Delete a Task
func (f *IBMPITaskClient) Delete(taskID string) error {

	// Create params and send request
	params := &params.PcloudTasksDeleteParams{
		Context: f.context,
		TaskID:  taskID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudTasksDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf("failed to delete the task id ... %w", err)
	}
	return nil
}
