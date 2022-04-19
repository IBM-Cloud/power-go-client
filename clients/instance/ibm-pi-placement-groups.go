package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_placement_groups"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

//IBMPIPlacementGroupClient ...
type IBMPIPlacementGroupClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPIPlacementGroupClient ...
func NewIBMPIPlacementGroupClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIPlacementGroupClient {
	return &IBMPIPlacementGroupClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudPlacementGroups,
	}
}

// Get a Placement Group
func (f *IBMPIPlacementGroupClient) Get(placementGroupID string) (*models.PlacementGroup, error) {

	// Create params and send request
	params := &params.PcloudPlacementgroupsGetParams{
		CloudInstanceID:  f.cloudInstanceID,
		Context:          f.context,
		PlacementGroupID: placementGroupID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudPlacementgroupsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetPlacementGroupOperationFailed, placementGroupID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Placement Group %s", placementGroupID)
	}
	return resp.Payload, nil
}

// Get All Placement Groups
func (f *IBMPIPlacementGroupClient) GetAll() (*models.PlacementGroups, error) {

	// Create params and send request
	params := &params.PcloudPlacementgroupsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudPlacementgroupsGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get All Placement Groups: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Placement Groups")
	}
	return resp.Payload, nil
}

// Create a Placement Group
func (f *IBMPIPlacementGroupClient) Create(createBody *models.PlacementGroupCreate) (*models.PlacementGroup, error) {

	// Create params and send request
	params := &params.PcloudPlacementgroupsPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudPlacementgroupsPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreatePlacementGroupOperationFailed, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create Placement Group")
	}
	return resp.Payload, nil
}

// Delete a Placement Group
func (f *IBMPIPlacementGroupClient) Delete(placementGroupID string) error {

	// Create params and send request
	params := &params.PcloudPlacementgroupsDeleteParams{
		CloudInstanceID:  f.cloudInstanceID,
		Context:          f.context,
		PlacementGroupID: placementGroupID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudPlacementgroupsDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.DeletePlacementGroupOperationFailed, placementGroupID, err)
	}
	return nil
}

// Add a Member to a Placement Group
func (f *IBMPIPlacementGroupClient) AddMember(placementGroupID string, udpateBody *models.PlacementGroupServer) (*models.PlacementGroup, error) {

	// Create params and send request
	params := &params.PcloudPlacementgroupsMembersPostParams{
		Body:             udpateBody,
		CloudInstanceID:  f.cloudInstanceID,
		Context:          f.context,
		PlacementGroupID: placementGroupID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudPlacementgroupsMembersPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.AddMemberPlacementGroupOperationFailed, *udpateBody.ID, placementGroupID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Add Member for instance %s and placement group %s", *udpateBody.ID, placementGroupID)
	}
	return resp.Payload, nil
}

// Delete a Member from Placement Group
func (f *IBMPIPlacementGroupClient) DeleteMember(placementGroupID string, udpateBody *models.PlacementGroupServer) (*models.PlacementGroup, error) {

	// Create params and send request
	params := &params.PcloudPlacementgroupsMembersDeleteParams{
		Body:             udpateBody,
		CloudInstanceID:  f.cloudInstanceID,
		Context:          f.context,
		PlacementGroupID: placementGroupID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	resp, err := f.request.PcloudPlacementgroupsMembersDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.DeleteMemberPlacementGroupOperationFailed, *udpateBody.ID, placementGroupID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Delete Member for instance %s and placement group %s", *udpateBody.ID, placementGroupID)
	}
	return resp.Payload, nil
}
