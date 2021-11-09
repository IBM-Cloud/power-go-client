package instance

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_service_d_h_c_p"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewIBMPIDhcpClient ...
type IBMPIDhcpClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIDhcpClient ...
func NewIBMPIDhcpClient(sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIDhcpClient {
	return &IBMPIDhcpClient{
		session:         sess,
		powerinstanceid: cloudInstanceID,
	}
}

// Create
func (f *IBMPIDhcpClient) Create(cloudInstanceID string) (dhcpServer *models.DHCPServer, err error) {
	return f.CreateWithContext(context.Background(), cloudInstanceID)
}

// Create a DHCP with context
func (f *IBMPIDhcpClient) CreateWithContext(ctx context.Context, cloudInstanceID string) (dhcpServer *models.DHCPServer, err error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpPostParamsWithContext(ctx).
		WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(cloudInstanceID)
	postAccepted, err := f.session.Power.PCloudServiceDHCP.PcloudDhcpPost(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if postAccepted != nil {
		dhcpServer = postAccepted.Payload
	}
	return
}

// Get
func (f *IBMPIDhcpClient) Get(id, cloudInstanceID string) (dhcpServer *models.DHCPServerDetail, err error) {
	return f.GetWithContext(context.Background(), id, cloudInstanceID)
}

// Get with Context...
func (f *IBMPIDhcpClient) GetWithContext(ctx context.Context, id, cloudInstanceID string) (dhcpServer *models.DHCPServerDetail, err error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithDhcpID(id)
	resp, err := f.session.Power.PCloudServiceDHCP.PcloudDhcpGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		dhcpServer = resp.Payload
	}
	return
}

// GetAll
func (f *IBMPIDhcpClient) GetAll(cloudInstanceID string) (dhcpServers models.DHCPServers, err error) {
	return f.GetAllWithContext(context.Background(), cloudInstanceID)
}

// GetAll with Context...
func (f *IBMPIDhcpClient) GetAllWithContext(ctx context.Context, cloudInstanceID string) (dhcpServers models.DHCPServers, err error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpGetallParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID)
	resp, err := f.session.Power.PCloudServiceDHCP.PcloudDhcpGetall(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		dhcpServers = resp.Payload
	}
	return
}

// Delete
func (f *IBMPIDhcpClient) Delete(id, cloudInstanceID string) (obj models.Object, err error) {
	return f.DeleteWithContext(context.Background(), id, cloudInstanceID)
}

// Delete with Context...
func (f *IBMPIDhcpClient) DeleteWithContext(ctx context.Context, id, cloudInstanceID string) (obj models.Object, err error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithDhcpID(id)
	respAccepted, err := f.session.Power.PCloudServiceDHCP.PcloudDhcpDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if respAccepted != nil {
		obj = respAccepted.Payload
	}
	return
}
