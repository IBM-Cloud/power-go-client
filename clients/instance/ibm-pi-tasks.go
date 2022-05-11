package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tasks"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPITaskClient.
type IBMPITaskClient struct {
	IBMPIClient
}

// NewIBMPITaskClient.
func NewIBMPITaskClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPITaskClient {
	return &IBMPITaskClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Task.
func (f *IBMPITaskClient) Get(id string) (*models.Task, error) {
	params := p.NewPcloudTasksGetParams().
		WithContext(f.ctx).
		WithTaskID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.taskRequest.PcloudTasksGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the task %s: %w", id, err)
	}
	return resp.Payload, nil
}

// Delete a Task.
func (f *IBMPITaskClient) Delete(id string) error {
	params := p.NewPcloudTasksDeleteParams().
		WithContext(f.ctx).
		WithTaskID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.taskRequest.PcloudTasksDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to delete the task id ... %w", err)
	}
	return nil
}
