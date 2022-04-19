package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPICloudInstanceClient ...
type IBMPICloudInstanceClient struct {
	auth    runtime.ClientAuthInfoWriter
	context context.Context
	request params.ClientService
}

// NewIBMPICloudInstanceClient ...
func NewIBMPICloudInstanceClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPICloudInstanceClient {
	return &IBMPICloudInstanceClient{
		auth:    sess.AuthInfo(cloudInstanceID),
		context: ctx,
		request: sess.Power.PCloudInstances,
	}
}

// Get a Cloud Instance
func (f *IBMPICloudInstanceClient) Get(cloudInstanceID string) (*models.CloudInstance, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesGetParams{
		CloudInstanceID: cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudinstancesGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetCloudInstanceOperationFailed, cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Cloud Instance %s", cloudInstanceID)
	}
	return resp.Payload, nil
}

// Update a Cloud Instance
func (f *IBMPICloudInstanceClient) Update(cloudInstanceID string, updateBody *models.CloudInstanceUpdate) (*models.CloudInstance, error) {

	// Create params and send request
	params := &params.PcloudCloudinstancesPutParams{
		Body:            updateBody,
		CloudInstanceID: cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudCloudinstancesPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateCloudInstanceOperationFailed, cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to update the Cloud instance %s", cloudInstanceID)
	}
	return resp.Payload, nil
}

// Delete a Cloud Instance
func (f *IBMPICloudInstanceClient) Delete(cloudInstanceID string) error {

	// Create params and send request
	params := &params.PcloudCloudinstancesDeleteParams{
		CloudInstanceID: cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)

	// Handle errors
	_, err := f.request.PcloudCloudinstancesDelete(params, f.auth)
	if err != nil {
		return fmt.Errorf(errors.DeleteCloudInstanceOperationFailed, cloudInstanceID, err)
	}
	return nil
}
