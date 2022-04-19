package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPINetworkClient ...
type IBMPINetworkClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPINetworkClient ...
func NewIBMPINetworkClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPINetworkClient {
	return &IBMPINetworkClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudNetworks,
	}
}

// Get a Network
func (f *IBMPINetworkClient) Get(networkID string) (*models.Network, error) {

	// Create params and send request
	params := &params.PcloudNetworksGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		NetworkID:       networkID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudNetworksGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetNetworkOperationFailed, networkID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Network %s", networkID)
	}
	return resp.Payload, nil
}

// Get All Networks
func (f *IBMPINetworkClient) GetAll() (*models.Networks, error) {

	// Create params and send request
	params := &params.PcloudNetworksGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudNetworksGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get Network for cloud instance %s with error %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Network for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a Network
func (f *IBMPINetworkClient) Create(createBody *models.NetworkCreate) (*models.Network, error) {

	// Create params and send request
	params := &params.PcloudNetworksPostParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		Body:            createBody,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	respOk, respCreated, err := f.request.PcloudNetworksPost(params, f.auth)

	// Handle errors
	if err != nil {
	}
	if respOk != nil && respOk.Payload != nil {
		return respOk.Payload, nil
	}
	if respCreated != nil && respCreated.Payload != nil {
		return respCreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to perform Create Network Operation for Network %s", createBody.Name)
}

// Update a Network
func (f *IBMPINetworkClient) Update(networkID string, updateBody *models.NetworkUpdate) (*models.Network, error) {

	// Create params and send request
	params := &params.PcloudNetworksPutParams{
		CloudInstanceID: f.cloudInstanceID,
		Body:            updateBody,
		Context:         f.context,
		NetworkID:       networkID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudNetworksPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to perform Update Network Operation for Network %s with error %w", networkID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Update Network Operation for Network %s", networkID)
	}
	return resp.Payload, nil
}

// Get All Public Networks
func (f *IBMPINetworkClient) GetAllPublic() (*models.Networks, error) {

	// Create params and send request
	filterQuery := "type=\"pub-vlan\""
	params := &params.PcloudNetworksGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		Filter:          &filterQuery,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudNetworksGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Public Networks for cloud instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Public Networks for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Delete a Network
func (f *IBMPINetworkClient) Delete(networkID string) error {

	// Create params and send request
	params := &params.PcloudNetworksDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		NetworkID:       networkID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudNetworksDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf("failed to Delete PI Network %s: %w", networkID, err)
	}
	return nil
}

// Get All Ports on a Network
func (f *IBMPINetworkClient) GetAllPorts(networkID string) (*models.NetworkPorts, error) {

	// Create params and send request
	params := &params.PcloudNetworksPortsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		NetworkID:       networkID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudNetworksPortsGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Network Ports for Network %s: %w", networkID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Network Ports for Network %s", networkID)
	}
	return resp.Payload, nil
}

// Get a Port on a Network
func (f *IBMPINetworkClient) GetPort(networkID string, portID string) (*models.NetworkPort, error) {

	// Create params and send request
	params := &params.PcloudNetworksPortsGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		NetworkID:       networkID,
		PortID:          portID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudNetworksPortsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get PI Network Port %s for Network %s: %w", portID, networkID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Network Port %s for Network %s", portID, networkID)
	}
	return resp.Payload, nil
}

// Create a Port on a Network
func (f *IBMPINetworkClient) CreatePort(networkID string, portBody *models.NetworkPortCreate) (*models.NetworkPort, error) {

	// Create params and send request
	params := &params.PcloudNetworksPortsPostParams{
		Body:            portBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		NetworkID:       networkID,
	}
	// params.SetTimeout(helpers.PICreateTimeOu)
	resp, err := f.request.PcloudNetworksPortsPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateNetworkPortOperationFailed, networkID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Create Network Port Operation for Network %s", networkID)
	}
	return resp.Payload, nil
}

// Delete a Port on a Network
func (f *IBMPINetworkClient) DeletePort(networkID string, portID string) error {

	// Create params and send request
	params := &params.PcloudNetworksPortsDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		NetworkID:       networkID,
		PortID:          portID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudNetworksPortsDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf("failed to delete the network port %s for network %s with error %w", portID, networkID, err)
	}
	return nil
}

// Update a Port on a Network
func (f *IBMPINetworkClient) UpdatePort(networkID, portID string, body *models.NetworkPortUpdate) (*models.NetworkPort, error) {

	// Create params and send request
	params := &params.PcloudNetworksPortsPutParams{
		Body:            body,
		CloudInstanceID: f.cloudInstanceID,
		NetworkID:       networkID,
		PortID:          portID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudNetworksPortsPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to update the port %s and Network %s with error %w", portID, networkID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to update the port %s and Network %s", portID, networkID)
	}
	return resp.Payload, nil
}
