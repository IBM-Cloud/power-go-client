package instance

import (
	"fmt"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"time"

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
func (f *IBMPICloudConnectionClient) Create(pclouddef *p_cloud_cloud_connections.PcloudCloudconnectionsPostParams, powerinstanceid string) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsPostParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithBody(pclouddef.Body)
	postok, postcreated, err, _ := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateCloudConnectionOperationFailed, powerinstanceid, err)
	}
	if postok != nil {
		return postok.Payload, nil
	}
	if postcreated != nil {
		return postcreated.Payload, nil
	}
	return nil, nil
}

// Get ...
func (f *IBMPICloudConnectionClient) Get(cloudinstanceid, cloudconnectionid string) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsGetParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(cloudinstanceid).WithCloudConnectionID(cloudconnectionid)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsGet(params, ibmpisession.NewAuth(f.session, cloudinstanceid))
	if err != nil {
		return nil, fmt.Errorf(errors.GetCloudConnectionOperationFailed, err)
	}
	return resp.Payload, nil
}

// GetAll ..
func (f *IBMPICloudConnectionClient) GetAll(powerinstanceid string, timeout time.Duration) (*models.CloudConnections, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsGetallParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf("Failed to get all cloud connection %s", err)
	}
	return resp.Payload, nil
}

// Update a cloud Connection
func (f *IBMPICloudConnectionClient) Update(updateparams *p_cloud_cloud_connections.PcloudCloudconnectionsPutParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsPutParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(updateparams.CloudInstanceID).WithCloudConnectionID(updateparams.CloudConnectionID).WithBody(updateparams.Body)
	resp, err, _ := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsPut(params, ibmpisession.NewAuth(f.session, updateparams.CloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateCloudConnectionOperationFailed, updateparams.CloudConnectionID, err)
	}
	return resp.Payload, nil
}

// Delete a Cloud Connection
func (f *IBMPICloudConnectionClient) Delete(pclouddef *p_cloud_cloud_connections.PcloudCloudconnectionsDeleteParams) (models.Object, error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsDeleteParamsWithTimeout(helpers.PIDeleteTimeOut).WithCloudInstanceID(pclouddef.CloudInstanceID).WithCloudConnectionID(pclouddef.CloudConnectionID)
	respok, _, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsDelete(params, ibmpisession.NewAuth(f.session, pclouddef.CloudInstanceID))

	if err != nil || respok.Payload == nil {
		return nil, fmt.Errorf(errors.DeleteCloudConnectionOperationFailed, pclouddef.CloudConnectionID, err)
	}
	return respok.Payload, nil
}

// AddNetwork to a cloud connection
func (f *IBMPICloudConnectionClient) AddNetwork(pcloudnetworkdef *p_cloud_cloud_connections.PcloudCloudconnectionsNetworksPutParams) (*models.CloudConnection, error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksPutParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(pcloudnetworkdef.CloudInstanceID).WithCloudConnectionID(pcloudnetworkdef.CloudConnectionID).WithNetworkID(pcloudnetworkdef.NetworkID)
	resp, err, _ := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksPut(params, ibmpisession.NewAuth(f.session, pcloudnetworkdef.CloudInstanceID))
	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to add the network to the cloudconnection %s", err)
	}
	return resp.Payload, nil
}

// DeleteNetwork Deletes a network from a cloud connection
func (f *IBMPICloudConnectionClient) DeleteNetwork(pcloudnetworkdef *p_cloud_cloud_connections.PcloudCloudconnectionsNetworksDeleteParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksDeleteParamsWithTimeout(helpers.PIDeleteTimeOut).WithCloudInstanceID(pcloudnetworkdef.CloudInstanceID).WithCloudConnectionID(pcloudnetworkdef.CloudConnectionID).WithNetworkID(pcloudnetworkdef.NetworkID)
	resp, err, _ := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksDelete(params, ibmpisession.NewAuth(f.session, pcloudnetworkdef.CloudInstanceID))

	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to perform the delete operation... %s", err)
	}
	return resp.Payload, nil
}

// UpdateNetwork Update a network from a cloud connection
func (f *IBMPICloudConnectionClient) UpdateNetwork(pcloudnetworkdef *p_cloud_cloud_connections.PcloudCloudconnectionsNetworksPutParams) (*models.CloudConnection, error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksPutParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(pcloudnetworkdef.CloudInstanceID).WithCloudConnectionID(pcloudnetworkdef.CloudConnectionID).WithNetworkID(pcloudnetworkdef.NetworkID)
	resp, err, _ := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksPut(params, ibmpisession.NewAuth(f.session, pcloudnetworkdef.CloudInstanceID))

	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform the update operation...%v", err)
	}
	return resp.Payload, nil
}

// get VPCs

func (f *IBMPICloudConnectionClient) GetVPC(cloudinstanceid string) (*models.CloudConnectionVirtualPrivateClouds, error) {
	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsVirtualprivatecloudsGetallParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(cloudinstanceid)

	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsVirtualprivatecloudsGetall(params, ibmpisession.NewAuth(f.session, cloudinstanceid))
	if err != nil || resp.Payload == nil {

		return nil, fmt.Errorf("failed to perform the getvpc operation...%v", err)
	}
	return resp.Payload, nil
}
