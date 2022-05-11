package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_service_d_h_c_p"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewIBMPIDhcpClient.
type IBMPIDhcpClient struct {
	IBMPIClient
}

// NewIBMPIDhcpClient.
func NewIBMPIDhcpClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIDhcpClient {
	return &IBMPIDhcpClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Create a DHCP server.
func (f *IBMPIDhcpClient) Create(body *models.DHCPServerCreate) (*models.DHCPServer, error) {
	params := p.NewPcloudDhcpPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	postaccepted, err := f.dhcpRequest.PcloudDhcpPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateDchpOperationFailed, f.cloudInstanceID, err)
	}
	if postaccepted != nil && postaccepted.Payload != nil {
		return postaccepted.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create DHCP")
}

// Get a DHCP server.
func (f *IBMPIDhcpClient) Get(id string) (*models.DHCPServerDetail, error) {
	params := p.NewPcloudDhcpGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithDhcpID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.dhcpRequest.PcloudDhcpGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetDhcpOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get DHCP %s", id)
	}
	return resp.Payload, nil
}

// Get All DHCP servers.
func (f *IBMPIDhcpClient) GetAll() (models.DHCPServers, error) {
	params := p.NewPcloudDhcpGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.dhcpRequest.PcloudDhcpGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all DHCP servers: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all DHCP servers")
	}
	return resp.Payload, nil
}

// Delete a DHCP server.
func (f *IBMPIDhcpClient) Delete(id string) error {
	params := p.NewPcloudDhcpDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithDhcpID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.dhcpRequest.PcloudDhcpDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteDhcpOperationFailed, id, err)
	}
	return nil
}
