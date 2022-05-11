package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_placement_groups"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

//IBMPIPlacementGroupClient.
type IBMPIPlacementGroupClient struct {
	IBMPIClient
}

// NewIBMPIPlacementGroupClient.
func NewIBMPIPlacementGroupClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIPlacementGroupClient {
	return &IBMPIPlacementGroupClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a PI Placement Group.
func (f *IBMPIPlacementGroupClient) Get(id string) (*models.PlacementGroup, error) {
	params := p.NewPcloudPlacementgroupsGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPlacementGroupID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.placementGroupRequest.PcloudPlacementgroupsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetPlacementGroupOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Placement Group %s", id)
	}
	return resp.Payload, nil
}

// Get All Placement Groups.
func (f *IBMPIPlacementGroupClient) GetAll() (*models.PlacementGroups, error) {
	params := p.NewPcloudPlacementgroupsGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.placementGroupRequest.PcloudPlacementgroupsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get All Placement Groups: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Placement Groups")
	}
	return resp.Payload, nil
}

// Create a Placement Group.
func (f *IBMPIPlacementGroupClient) Create(body *models.PlacementGroupCreate) (*models.PlacementGroup, error) {
	params := p.NewPcloudPlacementgroupsPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	postok, err := f.placementGroupRequest.PcloudPlacementgroupsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreatePlacementGroupOperationFailed, f.cloudInstanceID, err)
	}
	if postok == nil || postok.Payload == nil {
		return nil, fmt.Errorf("failed to Create Placement Group")
	}
	return postok.Payload, nil
}

// Delete a Placement Group.
func (f *IBMPIPlacementGroupClient) Delete(id string) error {
	params := p.NewPcloudPlacementgroupsDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPlacementGroupID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.placementGroupRequest.PcloudPlacementgroupsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeletePlacementGroupOperationFailed, id, err)
	}
	return nil
}

// Add an Instance to a Placement Group.
func (f *IBMPIPlacementGroupClient) AddMember(id string, body *models.PlacementGroupServer) (*models.PlacementGroup, error) {
	params := p.NewPcloudPlacementgroupsMembersPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPlacementGroupID(id).
		WithTimeout(helpers.PICreateTimeOut)
	postok, err := f.placementGroupRequest.PcloudPlacementgroupsMembersPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.AddMemberPlacementGroupOperationFailed, *body.ID, id, err)
	}
	if postok == nil || postok.Payload == nil {
		return nil, fmt.Errorf("failed to Add Member for instance %s and placement group %s", *body.ID, id)
	}
	return postok.Payload, nil
}

// Remove an Instance to a Placement Group.
func (f *IBMPIPlacementGroupClient) DeleteMember(id string, body *models.PlacementGroupServer) (*models.PlacementGroup, error) {
	params := p.NewPcloudPlacementgroupsMembersDeleteParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPlacementGroupID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	delok, err := f.placementGroupRequest.PcloudPlacementgroupsMembersDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.DeleteMemberPlacementGroupOperationFailed, *body.ID, id, err)
	}
	if delok == nil || delok.Payload == nil {
		return nil, fmt.Errorf("failed to Delete Member for instance %s and placement group %s", *body.ID, id)
	}
	return delok.Payload, nil
}
