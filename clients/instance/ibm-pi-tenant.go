package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPITenantClient.
type IBMPITenantClient struct {
	IBMPIClient
}

// NewIBMPITenantClient.
func NewIBMPITenantClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPITenantClient {
	return &IBMPITenantClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Tenant.
func (f *IBMPITenantClient) Get(tenantid string) (*models.Tenant, error) {
	params := p.NewPcloudTenantsGetParams().
		WithContext(f.ctx).
		WithTenantID(tenantid).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.tenantRequest.PcloudTenantsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant %s with error %w", tenantid, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get tenant %s", tenantid)
	}
	return resp.Payload, nil
}

// Get own Tenant.
func (f *IBMPITenantClient) GetSelfTenant() (*models.Tenant, error) {
	params := p.NewPcloudTenantsGetParams().
		WithContext(f.ctx).
		WithTenantID(f.session.Options.UserAccount).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.tenantRequest.PcloudTenantsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get self tenant with error %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get self tenant")
	}
	return resp.Payload, nil
}
