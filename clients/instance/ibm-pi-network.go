package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPINetworkClient.
type IBMPINetworkClient struct {
	IBMPIClient
}

// NewIBMPINetworkClient.
func NewIBMPINetworkClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPINetworkClient {
	return &IBMPINetworkClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Network.
func (f *IBMPINetworkClient) Get(id string) (*models.Network, error) {
	params := p.NewPcloudNetworksGetParams().
		WithContext(f.ctx).
		WithCloudInstanceID(f.cloudInstanceID).
		WithNetworkID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.networkRequest.PcloudNetworksGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetNetworkOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Network %s", id)
	}
	return resp.Payload, nil
}

// Get All Networks.
func (f *IBMPINetworkClient) GetAll() (*models.Networks, error) {
	params := p.NewPcloudNetworksGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.networkRequest.PcloudNetworksGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get Network for cloud instance %s with error %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Network for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a Network.
func (f *IBMPINetworkClient) Create(body *models.NetworkCreate) (*models.Network, error) {
	params := p.NewPcloudNetworksPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	postok, postcreated, err := f.networkRequest.PcloudNetworksPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateNetworkOperationFailed, body.Name, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	if postcreated != nil && postcreated.Payload != nil {
		return postcreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to perform Create Network Operation for Network %s", body.Name)
}

// Update a Network.
func (f *IBMPINetworkClient) Update(id string, body *models.NetworkUpdate) (*models.Network, error) {
	params := p.NewPcloudNetworksPutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithNetworkID(id).
		WithTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.networkRequest.PcloudNetworksPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to perform Update Network Operation for Network %s with error %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Update Network Operation for Network %s", id)
	}
	return resp.Payload, nil
}

// Get All Public Networks.
func (f *IBMPINetworkClient) GetAllPublic() (*models.Networks, error) {
	filterQuery := "type=\"pub-vlan\""
	params := p.NewPcloudNetworksGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithFilter(&filterQuery).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.networkRequest.PcloudNetworksGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Public Networks for cloud instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Public Networks for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Delete a Network.
func (f *IBMPINetworkClient) Delete(id string) error {
	params := p.NewPcloudNetworksDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithNetworkID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.networkRequest.PcloudNetworksDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PI Network %s: %w", id, err)
	}
	return nil
}

// Get All Ports on a Network.
func (f *IBMPINetworkClient) GetAllPorts(id string) (*models.NetworkPorts, error) {
	params := p.NewPcloudNetworksPortsGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithNetworkID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.networkRequest.PcloudNetworksPortsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Network Ports for Network %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Network Ports for Network %s", id)
	}
	return resp.Payload, nil
}

// Get a Port on a Network.
func (f *IBMPINetworkClient) GetPort(id string, networkPortID string) (*models.NetworkPort, error) {
	params := p.NewPcloudNetworksPortsGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithNetworkID(id).
		WithPortID(networkPortID).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.networkRequest.PcloudNetworksPortsGet(params, f.session.AuthInfo(f.cloudInstanceID))

	if err != nil {
		return nil, fmt.Errorf("failed to Get PI Network Port %s for Network %s: %w", networkPortID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Network Port %s for Network %s", networkPortID, id)
	}
	return resp.Payload, nil
}

// Create a Port on a Network.
func (f *IBMPINetworkClient) CreatePort(id string, body *models.NetworkPortCreate) (*models.NetworkPort, error) {
	params := p.NewPcloudNetworksPortsPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithNetworkID(id).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.networkRequest.PcloudNetworksPortsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateNetworkPortOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Create Network Port Operation for Network %s", id)
	}
	return resp.Payload, nil
}

// Delete a Port on a Network.
func (f *IBMPINetworkClient) DeletePort(id string, networkPortID string) error {
	params := p.NewPcloudNetworksPortsDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithNetworkID(id).
		WithPortID(networkPortID).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.networkRequest.PcloudNetworksPortsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to delete the network port %s for network %s with error %w", networkPortID, id, err)
	}
	return nil
}

// Update a Port on a Network.
func (f *IBMPINetworkClient) UpdatePort(id, networkPortID string, body *models.NetworkPortUpdate) (*models.NetworkPort, error) {
	params := p.NewPcloudNetworksPortsPutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithNetworkID(id).
		WithPortID(networkPortID).
		WithTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.networkRequest.PcloudNetworksPortsPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to update the port %s and Network %s with error %w", networkPortID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to update the port %s and Network %s", networkPortID, id)
	}
	return resp.Payload, nil
}
