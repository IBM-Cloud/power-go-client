package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_service_d_h_c_p"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// NewIBMPIDhcpClient ...
type IBMPIDhcpClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPIDhcpClient ...
func NewIBMPIDhcpClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIDhcpClient {
	return &IBMPIDhcpClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudServicedhcp,
	}
}

// Create a DHCP
func (f *IBMPIDhcpClient) Create(createBody *models.DHCPServerCreate) (*models.DHCPServer, error) {

	// Create params and send request
	params := &params.PcloudDhcpPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudDhcpPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateDchpOperationFailed, f.cloudInstanceID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create DHCP")
}

// Get a DHCP
func (f *IBMPIDhcpClient) Get(dhcpID string) (*models.DHCPServerDetail, error) {

	// Create params and send request
	params := &params.PcloudDhcpGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		DhcpID:          dhcpID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudDhcpGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetDhcpOperationFailed, dhcpID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get DHCP %s", dhcpID)
	}
	return resp.Payload, nil
}

// Get All DHCP's
func (f *IBMPIDhcpClient) GetAll() (models.DHCPServers, error) {

	// Create params and send request
	params := &params.PcloudDhcpGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudDhcpGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all DHCP servers: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all DHCP servers")
	}
	return resp.Payload, nil
}

// Delete a DHCP
func (f *IBMPIDhcpClient) Delete(dhcpID string) error {

	// Create params and send request
	params := &params.PcloudDhcpDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		DhcpID:          dhcpID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudDhcpDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.DeleteDhcpOperationFailed, dhcpID, err)
	}
	return nil
}
