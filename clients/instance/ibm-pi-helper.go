package instance

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	storageCapacity "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_storage_capacity"
	systemPool "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_system_pools"
	task "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tasks"
	tenant "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants"
)

// Helper methods that will be used by the client classes

// IBMPIHelperClient
type IBMPIClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	ctx             context.Context

	storageCapacityRequest storageCapacity.ClientService
	systemPoolRequest      systemPool.ClientService
	taskRequest            task.ClientService
	tenantRequest          tenant.ClientService
}

// NewIBMPIClient
func NewIBMPIClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIClient {
	return &IBMPIClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		ctx:             ctx,

		storageCapacityRequest: sess.Power.PCloudStorageCapacity,
		systemPoolRequest:      sess.Power.PCloudSystemPools,
		taskRequest:            sess.Power.PCloudTasks,
		tenantRequest:          sess.Power.PCloudTenants,
	}
}
