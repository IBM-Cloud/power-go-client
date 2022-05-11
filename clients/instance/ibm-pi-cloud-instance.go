package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPICloudInstanceClient.
type IBMPICloudInstanceClient struct {
	IBMPIClient
}

// NewIBMPICloudInstanceClient.
func NewIBMPICloudInstanceClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPICloudInstanceClient {
	return &IBMPICloudInstanceClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Cloud Instance.
func (f *IBMPICloudInstanceClient) Get(id string) (*models.CloudInstance, error) {
	params := p.NewPcloudCloudinstancesGetParams().
		WithCloudInstanceID(id).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.cloudInstanceRequest.PcloudCloudinstancesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetCloudInstanceOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Cloud Instance %s", id)
	}
	return resp.Payload, nil
}

// Update a Cloud Instance.
func (f *IBMPICloudInstanceClient) Update(id string, body *models.CloudInstanceUpdate) (*models.CloudInstance, error) {
	params := p.NewPcloudCloudinstancesPutParams().
		WithBody(body).
		WithCloudInstanceID(id).
		WithContext(f.ctx).
		WithTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.cloudInstanceRequest.PcloudCloudinstancesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateCloudInstanceOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to update the Cloud instance %s", id)
	}
	return resp.Payload, nil
}

// Delete a Cloud Instance.
func (f *IBMPICloudInstanceClient) Delete(id string) error {
	params := p.NewPcloudCloudinstancesDeleteParams().
		WithCloudInstanceID(id).
		WithContext(f.ctx).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.cloudInstanceRequest.PcloudCloudinstancesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteCloudInstanceOperationFailed, id, err)
	}
	return nil
}
