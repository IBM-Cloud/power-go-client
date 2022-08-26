package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_disaster_recovery"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIDisasterRecoveryLocationClient
type IBMPIDisasterRecoveryLocationClient struct {
	IBMPIClient
}

// NewIBMPIDisasterRecoveryLocationClient
func NewIBMPIDisasterRecoveryLocationClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIDisasterRecoveryLocationClient {
	return &IBMPIDisasterRecoveryLocationClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

func (f *IBMPIDisasterRecoveryLocationClient) Get() (*models.DisasterRecoveryLocation, error) {
	params := p_cloud_disaster_recovery.NewPcloudLocationsDisasterrecoveryGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudDisasterRecovery.PcloudLocationsDisasterrecoveryGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetDisasterRecoveryLocationOperationFailed, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform the Get Disaster Recovery Location for the cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

func (f *IBMPIDisasterRecoveryLocationClient) GetAll() (*models.DisasterRecoveryLocations, error) {
	params := p_cloud_disaster_recovery.NewPcloudLocationsDisasterrecoveryGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.PCloudDisasterRecovery.PcloudLocationsDisasterrecoveryGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetAllDisasterRecoveryLocationOperationFailed, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get All Disaster Recovery Location of Power Virtual Server")
	}
	return resp.Payload, nil
}
