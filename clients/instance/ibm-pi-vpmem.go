package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/v_p_mem_volumes"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIVPMEMClient
type IBMPIVPMEMClient struct {
	IBMPIClient
}

// NewIBMPIVPMEMClient
func NewIBMPIVPMEMClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVPMEMClient {
	return &IBMPIVPMEMClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get all VPMEM volumes
func (f *IBMPIVPMEMClient) GetAll() (*models.VPMemVolumes, error) {
	params := v_p_mem_volumes.NewV1VpmemVolumesGetallParams().WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.VpMemVolumes.V1VpmemVolumesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to get all VPMEM volumes  with %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all VPMEM volumes")
	}
	return resp.Payload, nil
}
