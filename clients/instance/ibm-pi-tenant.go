package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPITenantClient ...
type IBMPITenantClient struct {
	auth     runtime.ClientAuthInfoWriter
	context  context.Context
	request  params.ClientService
	tenantID string
}

// NewIBMPITenantClient ...
func NewIBMPITenantClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPITenantClient {
	return &IBMPITenantClient{
		auth:     sess.AuthInfo(cloudInstanceID),
		context:  ctx,
		request:  sess.Power.PCloudTenants,
		tenantID: sess.Options.UserAccount,
	}
}

// Get a Tenant
func (f *IBMPITenantClient) Get(tenantID string) (*models.Tenant, error) {

	// Create params and send request
	params := &params.PcloudTenantsGetParams{
		Context:  f.context,
		TenantID: tenantID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudTenantsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant %s with error %w", tenantID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get tenant %s", tenantID)
	}
	return resp.Payload, nil
}

// Get own Tenant
func (f *IBMPITenantClient) GetSelfTenant() (*models.Tenant, error) {

	// Create params and send request
	params := &params.PcloudTenantsGetParams{
		TenantID: f.tenantID,
		Context:  f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudTenantsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get self tenant with error %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get self tenant")
	}
	return resp.Payload, nil
}
