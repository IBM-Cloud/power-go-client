package instance

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIVpnConnectionClient ...
type IBMPIVpnConnectionClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIVpnConnectionClient ...
func NewIBMPIVpnConnectionClient(sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVpnConnectionClient {
	return &IBMPIVpnConnectionClient{
		session:         sess,
		powerinstanceid: cloudInstanceID,
	}
}

// Get information about a VPN connection
func (f *IBMPIVpnConnectionClient) Get(id, cloudInstanceID string) (*models.VPNConnection, error) {
	return f.GetWithContext(context.Background(), id, cloudInstanceID)
}

// Get information about a VPN connection with Context
func (f *IBMPIVpnConnectionClient) GetWithContext(ctx context.Context, id, cloudInstanceID string) (vpnConnection *models.VPNConnection, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithVpnConnectionID(id)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	vpnConnection = resp.Payload
	return
}

// Create a VPN connection
func (f *IBMPIVpnConnectionClient) Create(body *models.VPNConnectionCreate, cloudInstanceID string) (vpnConnection *models.VPNConnectionCreateResponse, err error) {
	return f.CreateWithContext(context.Background(), body, cloudInstanceID)
}

// Create a VPN connection with context
func (f *IBMPIVpnConnectionClient) CreateWithContext(ctx context.Context, body *models.VPNConnectionCreate, cloudInstanceID string) (vpnConnection *models.VPNConnectionCreateResponse, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPostParamsWithContext(ctx).
		WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithBody(body)
	postok, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPost(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if postok != nil {
		vpnConnection = postok.Payload
	}
	return
}

// Update a VPN connection
func (f *IBMPIVpnConnectionClient) Update(body *models.VPNConnectionUpdate, id, cloudInstanceID string) (vpnConnection *models.VPNConnection, err error) {
	return f.UpdateWithContext(context.Background(), body, id, cloudInstanceID)
}

// Update a VPN connection with context
func (f *IBMPIVpnConnectionClient) UpdateWithContext(ctx context.Context, body *models.VPNConnectionUpdate, id, cloudInstanceID string) (vpnConnection *models.VPNConnection, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPutParamsWithContext(ctx).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithBody(body).
		WithVpnConnectionID(id)
	putok, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPut(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if putok != nil {
		vpnConnection = putok.Payload
	}
	return
}

// Get all VPN connections
func (f *IBMPIVpnConnectionClient) GetAll(cloudInstanceID string) (vpnConnections *models.VPNConnections, err error) {
	return f.GetAllWithContext(context.Background(), cloudInstanceID)
}

// Get all VPN connections with context
func (f *IBMPIVpnConnectionClient) GetAllWithContext(ctx context.Context, cloudInstanceID string) (vpnConnections *models.VPNConnections, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsGetallParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsGetall(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	vpnConnections = resp.Payload
	return
}

// Delete a VPN connection
func (f *IBMPIVpnConnectionClient) Delete(id, cloudInstanceID string) (job *models.JobReference, err error) {
	return f.DeleteWithContext(context.Background(), id, cloudInstanceID)
}

// Delete a VPN connection with context
func (f *IBMPIVpnConnectionClient) DeleteWithContext(ctx context.Context, id, cloudInstanceID string) (job *models.JobReference, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithVpnConnectionID(id)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	job = resp.Payload
	return
}

// Network get
func (f *IBMPIVpnConnectionClient) GetNetworkWithContext(ctx context.Context, id, cloudInstanceID string) (networkIDs *models.NetworkIds, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsNetworksGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithVpnConnectionID(id)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsNetworksGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		networkIDs = resp.Payload
	}
	return
}

// Network attach
func (f *IBMPIVpnConnectionClient) AddNetworkWithContext(ctx context.Context, id, networkID, cloudInstanceID string) (job *models.JobReference, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsNetworksPutParamsWithContext(ctx).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithVpnConnectionID(id).
		WithBody(&models.NetworkID{NetworkID: &networkID})
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsNetworksPut(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		job = resp.Payload
	}
	return
}

// Network detach
func (f *IBMPIVpnConnectionClient) DeleteNetworkWithContext(ctx context.Context, id, networkID, cloudInstanceID string) (job *models.JobReference, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsNetworksDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithVpnConnectionID(id).
		WithBody(&models.NetworkID{NetworkID: &networkID})
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsNetworksDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		job = resp.Payload
	}
	return
}

// Subnet get
func (f *IBMPIVpnConnectionClient) GetSubnetWithContext(ctx context.Context, id, subnetID, cloudInstanceID string) (subnetIDs *models.PeerSubnets, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPeersubnetsGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithVpnConnectionID(id)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPeersubnetsGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		subnetIDs = resp.Payload
	}
	return
}

// Subnet attach
func (f *IBMPIVpnConnectionClient) AddSubnetWithContext(ctx context.Context, id, subnet, cloudInstanceID string) (subnets *models.PeerSubnets, err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPeersubnetsPutParamsWithContext(ctx).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithVpnConnectionID(id).
		WithBody(&models.PeerSubnetUpdate{Cidr: &subnet})
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPeersubnetsPut(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		subnets = resp.Payload
	}
	return
}

// Subnet detach
func (f *IBMPIVpnConnectionClient) DeleteSubnetWithContext(ctx context.Context, id, subnet, cloudInstanceID string) (err error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPeersubnetsDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithVpnConnectionID(id).
		WithBody(&models.PeerSubnetUpdate{Cidr: &subnet})
	_, err = f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPeersubnetsDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	return
}
