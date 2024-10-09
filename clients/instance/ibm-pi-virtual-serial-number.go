package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_virtual_serial_number"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIVSNClient

type IBMPIVSNClient struct {
	IBMPIClient
}

// NewIBMPIVSNClient
func NewIBMPIVSNClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVSNClient {
	return &IBMPIVSNClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get Virtual Serial Number
func (f *IBMPIVSNClient) Get(id string) (*models.VirtualSerialNumber, error) {
	params := p_cloud_virtual_serial_number.NewPcloudVirtualserialnumberGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithVirtualSerialNumber(id)
	resp, err := f.session.Power.PCloudVirtualSerialNumber.PcloudVirtualserialnumberGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get virtual serial number %s :%w", id, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get virtual serial number %s", id)
	}
	return resp.Payload, nil
}

// Get All Virtual Serial Numbers
func (f *IBMPIVSNClient) GetAll() (models.VirtualSerialNumberList, error) {
	params := p_cloud_virtual_serial_number.NewPcloudVirtualserialnumberGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.PCloudVirtualSerialNumber.PcloudVirtualserialnumberGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all virtual serial numbers in %s :%w", f.cloudInstanceID, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all virtual serial numbers in %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}
