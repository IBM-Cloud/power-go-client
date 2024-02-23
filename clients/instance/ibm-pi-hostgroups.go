package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/hostgroups"
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
	params := hostgroups.NewV1AvailableHostsParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.Hostgroups.V1AvailableHosts(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get available hosts for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Available hosts")
	}
	return resp.Payload, nil
}

// Get all Hostgroups
func (f *IBMPIHostgroupsClient) GetHostgroups() (models.HostgroupList, error) {
	params := hostgroups.NewV1HostgroupsGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.Hostgroups.V1HostgroupsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get Hostgroups for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Hostgroups")
	}
	return resp.Payload, nil
}

// Create a Hostgroup
func (f *IBMPIHostgroupsClient) CreateHostgroup(body *models.HostgroupCreate) (*models.Hostgroup, error) {
	params := hostgroups.NewV1HostgroupsPostParams().WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).WithBody(body)
	resp, err := f.session.Power.Hostgroups.V1HostgroupsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Create Hostgroup for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Hostgroups")
	}
	return resp.Payload, nil
}

// Update a Hostgroup
func (f *IBMPIHostgroupsClient) UpdateHostgroup(body *models.HostgroupShareOp, id string) (*models.Hostgroup, error) {
	params := hostgroups.NewV1HostgroupsIDPutParams().WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).WithBody(body).WithHostgroupID(id)
	resp, err := f.session.Power.Hostgroups.V1HostgroupsIDPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Update Hostgroup for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update Hostgroups")
	}
	return resp.Payload, nil
}

// Get a Hostgroup
func (f *IBMPIHostgroupsClient) GetHostgroup(id string) (*models.Hostgroup, error) {
	params := hostgroups.NewV1HostgroupsIDGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithHostgroupID(id)
	resp, err := f.session.Power.Hostgroups.V1HostgroupsIDGet(params, f.session.AuthInfo(f.cloudInstanceID))
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
	params := hostgroups.NewV1HostsGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.Hostgroups.V1HostsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get Hosts for %s: %w", f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Hosts")
	}
	return resp.Payload, nil
}

// Create a Host
func (f *IBMPIHostgroupsClient) CreateHost(body *models.HostCreate) (*models.Host, error) {
	params := hostgroups.NewV1HostsPostParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithBody(body)
	resp, err := f.session.Power.Hostgroups.V1HostsPost(params, f.session.AuthInfo(f.cloudInstanceID))
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
	params := hostgroups.NewV1HostsIDGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithHostID(id)
	resp, err := f.session.Power.Hostgroups.V1HostsIDGet(params, f.session.AuthInfo(f.cloudInstanceID))
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
	params := hostgroups.NewV1HostsIDPutParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithHostID(id).WithBody(body)
	resp, err := f.session.Power.Hostgroups.V1HostsIDPut(params, f.session.AuthInfo(f.cloudInstanceID))
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
	params := hostgroups.NewV1HostsIDDeleteParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithHostID(id)
	resp, err := f.session.Power.Hostgroups.V1HostsIDDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Delete Host %s for %s: %w", id, f.cloudInstanceID, err))
	}

	if resp == nil || resp.Payload == nil {
		return fmt.Errorf("failed to Delete Host %s", id)
	}
	return nil
}
