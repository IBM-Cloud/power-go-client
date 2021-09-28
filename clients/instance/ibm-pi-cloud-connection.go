package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_cloud_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPICloudConnectionClient ...
type IBMPICloudConnectionClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPICloudConnectionClient ...
func NewIBMPICloudConnectionClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPICloudConnectionClient {
	return &IBMPICloudConnectionClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

// Create a Cloud Connection
func (f *IBMPICloudConnectionClient) Create(body *models.CloudConnectionCreate, powerinstanceid string) (*models.CloudConnection, *models.CloudConnectionCreateResponse, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsPostParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithBody(body)
	postok, postcreated, postAccepted, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, nil, fmt.Errorf(errors.CreateCloudConnectionOperationFailed, powerinstanceid, err)
	}
	if postok != nil {
		return postok.Payload, nil, nil
	}
	if postcreated != nil {
		return postcreated.Payload, nil, nil
	}
	if postAccepted != nil {
		return nil, postAccepted.Payload, nil
	}
	return nil, nil, nil
}

// Create a Cloud Connection with context
func (f *IBMPICloudConnectionClient) CreateWithContext(ctx context.Context, body *models.CloudConnectionCreate, cloudInstanceID string) (cloudConnection *models.CloudConnection, cloudConnectionJob *models.CloudConnectionCreateResponse, err error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsPostParamsWithContext(ctx).
		WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithBody(body)
	postok, postcreated, postAccepted, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsPost(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if postok != nil {
		cloudConnection = postok.Payload
	} else if postcreated != nil {
		cloudConnection = postcreated.Payload
	} else if postAccepted != nil {
		cloudConnectionJob = postAccepted.Payload
	}
	return
}

// Get ...
func (f *IBMPICloudConnectionClient) Get(id, cloudInstanceID string) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsGetParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(cloudInstanceID).WithCloudConnectionID(id)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetCloudConnectionOperationFailed, cloudInstanceID, err)
	}
	return resp.Payload, nil
}

// Get with Context...
func (f *IBMPICloudConnectionClient) GetWithContext(ctx context.Context, id, cloudInstanceID string) (cloudConnection *models.CloudConnection, err error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsGetParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithCloudConnectionID(id)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsGet(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	cloudConnection = resp.Payload
	return
}

// GetAll ..
func (f *IBMPICloudConnectionClient) GetAll(powerinstanceid string) (*models.CloudConnections, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsGetallParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf("failed to get all cloud connection %s", err)
	}
	return resp.Payload, nil
}

// GetAll with Context...
func (f *IBMPICloudConnectionClient) GetAllWithContext(ctx context.Context, cloudInstanceID string) (cloudConnections *models.CloudConnections, err error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsGetallParamsWithContext(ctx).
		WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(cloudInstanceID)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsGetall(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	cloudConnections = resp.Payload
	return
}

// Update a cloud Connection
func (f *IBMPICloudConnectionClient) Update(updateparams *p_cloud_cloud_connections.PcloudCloudconnectionsPutParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsPutParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(updateparams.CloudInstanceID).WithCloudConnectionID(updateparams.CloudConnectionID).WithBody(updateparams.Body)
	_, _, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsPut(params, ibmpisession.NewAuth(f.session, updateparams.CloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateCloudConnectionOperationFailed, updateparams.CloudConnectionID, err)
	}

	return nil, nil
}

// Update a cloud Connection with context
func (f *IBMPICloudConnectionClient) UpdateWithContext(ctx context.Context, id, cloudInstanceID string, body *models.CloudConnectionUpdate) (cloudConnection *models.CloudConnection, cloudConnectionJob *models.JobReference, err error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsPutParamsWithContext(ctx).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithCloudConnectionID(id).
		WithBody(body)
	resp, respAccepted, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsPut(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if resp != nil {
		cloudConnection = resp.Payload
	} else if respAccepted != nil {
		cloudConnectionJob = respAccepted.Payload
	}
	return
}

// Delete a Cloud Connection
func (f *IBMPICloudConnectionClient) Delete(id, cloudInstanceID string) (models.Object, *models.JobReference, error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsDeleteParamsWithTimeout(helpers.PIDeleteTimeOut).WithCloudInstanceID(cloudInstanceID).WithCloudConnectionID(id)
	respOk, respAccepted, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))

	if err != nil {
		return nil, nil, fmt.Errorf(errors.DeleteCloudConnectionOperationFailed, id, err)
	}
	if respOk != nil {
		return respOk.Payload, nil, nil
	}
	if respAccepted != nil {
		return nil, respAccepted.Payload, nil
	}
	return nil, nil, nil
}

// Delete a Cloud Connection
func (f *IBMPICloudConnectionClient) DeleteWithContext(ctx context.Context, id, cloudInstanceID string) (obj models.Object, deleteJob *models.JobReference, err error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithCloudConnectionID(id)
	respOk, respAccepted, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if respOk != nil {
		obj = respOk.Payload
	}
	if respAccepted != nil {
		deleteJob = respAccepted.Payload
	}
	return
}

// AddNetwork to a cloud connection
func (f *IBMPICloudConnectionClient) AddNetwork(pcloudnetworkdef *p_cloud_cloud_connections.PcloudCloudconnectionsNetworksPutParams) (*models.CloudConnection, error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksPutParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(pcloudnetworkdef.CloudInstanceID).WithCloudConnectionID(pcloudnetworkdef.CloudConnectionID).WithNetworkID(pcloudnetworkdef.NetworkID)
	_, _, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksPut(params, ibmpisession.NewAuth(f.session, pcloudnetworkdef.CloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to add the network to the cloudconnection %s", err)
	}
	return nil, nil
}

// DeleteNetwork Deletes a network from a cloud connection
func (f *IBMPICloudConnectionClient) DeleteNetwork(pcloudnetworkdef *p_cloud_cloud_connections.PcloudCloudconnectionsNetworksDeleteParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksDeleteParamsWithTimeout(helpers.PIDeleteTimeOut).WithCloudInstanceID(pcloudnetworkdef.CloudInstanceID).WithCloudConnectionID(pcloudnetworkdef.CloudConnectionID).WithNetworkID(pcloudnetworkdef.NetworkID)
	_, _, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksDelete(params, ibmpisession.NewAuth(f.session, pcloudnetworkdef.CloudInstanceID))

	if err != nil {
		return nil, fmt.Errorf("failed to perform the delete operation... %s", err)
	}
	return nil, nil
}

// AddNetwork to a cloud connection
func (f *IBMPICloudConnectionClient) AddNetworkWithContext(ctx context.Context, networkID, cloudConnectionID, cloudInstanceID string) (jobReferece *models.JobReference, err error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksPutParamsWithContext(ctx).
		WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithCloudConnectionID(cloudConnectionID).
		WithNetworkID(networkID)
	_, respAccepted, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksPut(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if respAccepted != nil {
		jobReferece = respAccepted.Payload
	}
	return
}

// DeleteNetwork Deletes a network from a cloud connection
func (f *IBMPICloudConnectionClient) DeleteNetworkWithContext(ctx context.Context, networkID, cloudConnectionID, cloudInstanceID string) (jobReferece *models.JobReference, err error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksDeleteParamsWithContext(ctx).
		WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(cloudInstanceID).
		WithCloudConnectionID(cloudConnectionID).
		WithNetworkID(networkID)
	_, respAccepted, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksDelete(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil {
		return
	}
	if respAccepted != nil {
		jobReferece = respAccepted.Payload
	}
	return
}

// get VPCs

func (f *IBMPICloudConnectionClient) GetVPC(cloudInstanceID string) (*models.CloudConnectionVirtualPrivateClouds, error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsVirtualprivatecloudsGetallParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(cloudInstanceID)

	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsVirtualprivatecloudsGetall(params, ibmpisession.NewAuth(f.session, cloudInstanceID))
	if err != nil || resp.Payload == nil {

		return nil, fmt.Errorf("failed to perform the getvpc operation...%v", err)
	}
	return resp.Payload, nil
}
