package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPIVpnConnectionClient ...
type IBMPIVpnConnectionClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPIVpnConnectionClient ...
func NewIBMPIVpnConnectionClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVpnConnectionClient {
	return &IBMPIVpnConnectionClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudvpnConnections,
	}
}

// Get a VPN Connection
func (f *IBMPIVpnConnectionClient) Get(vpnID string) (*models.VPNConnection, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudVpnconnectionsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetVPNConnectionOperationFailed, vpnID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get VPN Connection %s", vpnID)
	}
	return resp.Payload, nil
}

// Create a VPN Connection
func (f *IBMPIVpnConnectionClient) Create(createBody *models.VPNConnectionCreate) (*models.VPNConnectionCreateResponse, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudVpnconnectionsPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVPNConnectionOperationFailed, f.cloudInstanceID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create VPN Connection")
}

// Update a VPN Connection
func (f *IBMPIVpnConnectionClient) Update(vpnID string, updateBody *models.VPNConnectionUpdate) (*models.VPNConnection, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsPutParams{
		Body:            updateBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudVpnconnectionsPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVPNConnectionOperationFailed, vpnID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, fmt.Errorf("failed to Update VPN Connection %s", vpnID)
}

// Get All VPN Connections
func (f *IBMPIVpnConnectionClient) GetAll() (*models.VPNConnections, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudVpnconnectionsGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all VPN Connections: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all VPN Connections")
	}
	return resp.Payload, nil
}

// Delete a VPN Connection
func (f *IBMPIVpnConnectionClient) Delete(vpnID string) (*models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	resp, err := f.request.PcloudVpnconnectionsDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.DeleteVPNConnectionOperationFailed, vpnID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, nil
}

// Get a VPN Connection Network
func (f *IBMPIVpnConnectionClient) GetNetwork(vpnID string) (*models.NetworkIDs, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsNetworksGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudVpnconnectionsNetworksGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get Networks for VPN Connection %s: %w", vpnID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Networks for VPN Connection %s", vpnID)
	}
	return resp.Payload, nil
}

// Add a Network to a VPN Connection
func (f *IBMPIVpnConnectionClient) AddNetwork(vpnID, networkID string) (*models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsNetworksPutParams{
		Body:            &models.NetworkID{NetworkID: &networkID},
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudVpnconnectionsNetworksPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Add Network %s to VPN Connection %s: %w", networkID, vpnID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, nil
}

// Detach a Network from a VPN Connection
func (f *IBMPIVpnConnectionClient) DeleteNetwork(vpnID, networkID string) (*models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsNetworksDeleteParams{
		Body:            &models.NetworkID{NetworkID: &networkID},
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	resp, err := f.request.PcloudVpnconnectionsNetworksDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Delete Network %s from VPN Connection %s: %w", networkID, vpnID, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, nil
}

// Get a VPN Connection's Subnet
func (f *IBMPIVpnConnectionClient) GetSubnet(vpnID string) (*models.PeerSubnets, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsPeersubnetsGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudVpnconnectionsPeersubnetsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get Subnets from VPN Connection %s: %w", vpnID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Subnets from VPN Connection %s", vpnID)
	}
	return resp.Payload, nil
}

// Attach a Subnet to a VPN Connection
func (f *IBMPIVpnConnectionClient) AddSubnet(vpnID, subnet string) (*models.PeerSubnets, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsPeersubnetsPutParams{
		Body:            &models.PeerSubnetUpdate{Cidr: &subnet},
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudVpnconnectionsPeersubnetsPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Add Subnets to VPN Connection %s: %w", vpnID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Add Subnets to VPN Connection %s", vpnID)
	}
	return resp.Payload, nil
}

// Detach a Subnet from a VPN Connection
func (f *IBMPIVpnConnectionClient) DeleteSubnet(vpnID, subnet string) (*models.PeerSubnets, error) {

	// Create params and send request
	params := &params.PcloudVpnconnectionsPeersubnetsDeleteParams{
		Body:            &models.PeerSubnetUpdate{Cidr: &subnet},
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		VpnConnectionID: vpnID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	resp, err := f.request.PcloudVpnconnectionsPeersubnetsDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Delete Subnet from VPN Connection %s: %w", vpnID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Delete Subnet from VPN Connection %s", vpnID)
	}
	return resp.Payload, nil
}
