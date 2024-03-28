package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/host_groups"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIHostgroupsClient
type IBMPIHostgroupsClient struct {
	IBMPIClient
}

// NewIBMPIHostgroupsClient
func NewIBMPHostgroupsClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIHostgroupsClient {
	return &IBMPIHostgroupsClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get All available hosts
func (f *IBMPIHostgroupsClient) GetAvailableHosts() (models.AvailableHostList, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1AvailableHostsParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.HostGroups.V1AvailableHosts(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get available hosts for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Available hosts")
	}
	return resp.Payload, nil
}

// Get all HostGroups
func (f *IBMPIHostgroupsClient) GetHostGroups() (models.HostGroupList, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostGroupsGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.HostGroups.V1HostGroupsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get Hostgroups for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Hostgroups")
	}
	return resp.Payload, nil
}

// Create a HostGroup
func (f *IBMPIHostgroupsClient) CreateHostGroup(body *models.HostGroupCreate) (*models.HostGroup, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostGroupsPostParams().WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).WithBody(body)
	resp, err := f.session.Power.HostGroups.V1HostGroupsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Create Hostgroup for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Hostgroups")
	}
	return resp.Payload, nil
}

// Update a HostGroup
func (f *IBMPIHostgroupsClient) UpdateHostGroup(body *models.HostGroupShareOp, id string) (*models.HostGroup, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostGroupsIDPutParams().WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).WithBody(body).WithHostGroupID(id)
	resp, err := f.session.Power.HostGroups.V1HostGroupsIDPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Update Hostgroup for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update Hostgroups")
	}
	return resp.Payload, nil
}

// Get a Hostgroup
func (f *IBMPIHostgroupsClient) GetHostGroup(id string) (*models.HostGroup, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostGroupsIDGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithHostGroupID(id)
	resp, err := f.session.Power.HostGroups.V1HostGroupsIDGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get Hostgroup %s for %s: %w", id, f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Hostgroup %s", id)
	}
	return resp.Payload, nil
}

// Get  All Hosts
func (f *IBMPIHostgroupsClient) GetHosts() (models.HostList, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostsGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.HostGroups.V1HostsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get Hosts for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Hosts")
	}
	return resp.Payload, nil
}

// Create a Host
func (f *IBMPIHostgroupsClient) CreateHost(body *models.HostCreate) (models.HostList, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostsPostParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithBody(body)
	resp, err := f.session.Power.HostGroups.V1HostsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Create Hosts for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create Hosts")
	}
	return resp.Payload, nil
}

// Get a Host
func (f *IBMPIHostgroupsClient) GetHost(id string) (*models.Host, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostsIDGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithHostID(id)
	resp, err := f.session.Power.HostGroups.V1HostsIDGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get Host %s for %s: %w", id, f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Host %s", id)
	}
	return resp.Payload, nil
}

// Update a Host
func (f *IBMPIHostgroupsClient) UpdateHost(body *models.HostPut, id string) (*models.Host, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostsIDPutParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithHostID(id).WithBody(body)
	resp, err := f.session.Power.HostGroups.V1HostsIDPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Update Host %s for %s: %w", id, f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update Host %s", id)
	}
	return resp.Payload, nil
}

// Delete a Host
func (f *IBMPIHostgroupsClient) DeleteHost(id string) error {
	if f.session.IsOnPrem() {
		return fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := host_groups.NewV1HostsIDDeleteParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithHostID(id)
	resp, err := f.session.Power.HostGroups.V1HostsIDDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Delete Host %s for %s: %w", id, f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return fmt.Errorf("failed to Delete Host %s", id)
	}
	return nil
}
