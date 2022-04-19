package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_system_pools"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPISystemPoolClient ...
type IBMPISystemPoolClient struct {
	auth    runtime.ClientAuthInfoWriter
	context context.Context
	request params.ClientService
}

// NewIBMPISystemPoolClient ...
func NewIBMPISystemPoolClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPISystemPoolClient {
	return &IBMPISystemPoolClient{
		auth:    sess.AuthInfo(cloudInstanceID),
		context: ctx,
		request: sess.Power.PCloudSystemPools,
	}
}

// Get a System Pool
func (f *IBMPISystemPoolClient) Get(cloudInstanceID string) (models.SystemPools, error) {

	// Create params and send request
	params := &params.PcloudSystempoolsGetParams{
		CloudInstanceID: cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudSystempoolsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetSystemPoolsOperationFailed, cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Get System Pools Operation for cloud instance id %s", cloudInstanceID)
	}
	return resp.Payload, nil
}
