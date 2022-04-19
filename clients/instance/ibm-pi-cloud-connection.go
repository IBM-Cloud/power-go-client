package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_cloud_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPICloudConnectionClient
type IBMPICloudConnectionClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPICloudConnectionClient
func NewIBMPICloudConnectionClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPICloudConnectionClient {
	return &IBMPICloudConnectionClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudCloudConnections,
	}
}

// Create a Cloud Connection
func (f *IBMPICloudConnectionClient) Create(createBody *models.CloudConnectionCreate) (*models.CloudConnection, *models.CloudConnectionCreateResponse, error) {

	// Create params and send request
	params := &params.PcloudCloudconnectionsPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	respOk, respCreated, respAccepted, err := f.request.PcloudCloudconnectionsPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, nil, fmt.Errorf(errors.CreateCloudConnectionOperationFailed, f.cloudInstanceID, err)
	}
	if respOk != nil && respOk.Payload != nil {
		return respOk.Payload, nil, nil
	}
	if respCreated != nil && respCreated.Payload != nil {
		return respCreated.Payload, nil, nil
	}
	if respAccepted != nil && respAccepted.Payload != nil {
		return nil, respAccepted.Payload, nil
	}
	return nil, nil, fmt.Errorf("failed to Create Cloud Connection")
}

// Get a Cloud Connection
func (f *IBMPICloudConnectionClient) Get(cloudConnectionID string) (*models.CloudConnection, error) {

	// Create params and send request
	params := &params.PcloudCloudconnectionsGetParams{
		CloudConnectionID: cloudConnectionID,
		CloudInstanceID:   f.cloudInstanceID,
		Context:           f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudconnectionsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetCloudConnectionOperationFailed, cloudConnectionID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Get Cloud Connections Operation for cloudconnectionid %s", cloudConnectionID)
	}
	return resp.Payload, nil
}

// GetAll Cloud Connections
func (f *IBMPICloudConnectionClient) GetAll() (*models.CloudConnections, error) {

	// Create params and send request
	params := &params.PcloudCloudconnectionsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudconnectionsGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Cloud Connections: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Cloud Connections")
	}
	return resp.Payload, nil
}

// Update a Cloud Connection
func (f *IBMPICloudConnectionClient) Update(cloudConnectionID string, updateBody *models.CloudConnectionUpdate) (*models.CloudConnection, *models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudCloudconnectionsPutParams{
		Body:              updateBody,
		CloudConnectionID: cloudConnectionID,
		CloudInstanceID:   f.cloudInstanceID,
		Context:           f.context,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	respOk, respAccepted, err := f.request.PcloudCloudconnectionsPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, nil, fmt.Errorf(errors.UpdateCloudConnectionOperationFailed, cloudConnectionID, err)
	}
	if respOk != nil && respOk.Payload != nil {
		return respOk.Payload, nil, nil
	}
	if respAccepted != nil && respAccepted.Payload != nil {
		return nil, respAccepted.Payload, nil
	}
	return nil, nil, fmt.Errorf("failed to Update Cloud Connection %s", cloudConnectionID)
}

// Delete a Cloud Connection
func (f *IBMPICloudConnectionClient) Delete(cloudConnectionID string) (*models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudCloudconnectionsDeleteParams{
		CloudConnectionID: cloudConnectionID,
		CloudInstanceID:   f.cloudInstanceID,
		Context:           f.context,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, respAccepted, err := f.request.PcloudCloudconnectionsDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.DeleteCloudConnectionOperationFailed, cloudConnectionID, err)
	}
	if respAccepted != nil && respAccepted.Payload != nil {
		return respAccepted.Payload, nil
	}
	return nil, nil
}

// Add a Network to a Cloud Connection
func (f *IBMPICloudConnectionClient) AddNetwork(cloudConnectionID, networkID string) (*models.CloudConnection, *models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudCloudconnectionsNetworksPutParams{
		CloudConnectionID: cloudConnectionID,
		CloudInstanceID:   f.cloudInstanceID,
		Context:           f.context,
		NetworkID:         networkID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	respOk, respAccepted, err := f.request.PcloudCloudconnectionsNetworksPut(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, nil, fmt.Errorf("failed to Add Network %s to Cloud Connection %s: %w", networkID, cloudConnectionID, err)
	}
	if respOk != nil && respOk.Payload != nil {
		return respOk.Payload, nil, nil
	}
	if respAccepted != nil && respAccepted.Payload != nil {
		return nil, respAccepted.Payload, nil
	}
	return nil, nil, nil
}

// Remove a Network from a Vloud Connection
func (f *IBMPICloudConnectionClient) DeleteNetwork(cloudConnectionID, networkID string) (*models.CloudConnection, *models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudCloudconnectionsNetworksDeleteParams{
		CloudConnectionID: cloudConnectionID,
		CloudInstanceID:   f.cloudInstanceID,
		Context:           f.context,
		NetworkID:         networkID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	respOk, respAccepted, err := f.request.PcloudCloudconnectionsNetworksDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, nil, fmt.Errorf("failed to Delete Network %s from Cloud Connection %s: %w", networkID, cloudConnectionID, err)
	}
	if respOk != nil && respOk.Payload != nil {
		return respOk.Payload, nil, nil
	}
	if respAccepted != nil && respAccepted.Payload != nil {
		return nil, respAccepted.Payload, nil
	}
	return nil, nil, nil
}

// Get all VPCs for a Cloud Instance
func (f *IBMPICloudConnectionClient) GetVPC() (*models.CloudConnectionVirtualPrivateClouds, error) {

	// Create params and send request
	params := &params.PcloudCloudconnectionsVirtualprivatecloudsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudCloudconnectionsVirtualprivatecloudsGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to perform the get vpc operation: %w", err)
	}
	if resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform the get vpc operation")
	}
	return resp.Payload, nil
}
