package instance

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"

	instance "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	sap "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_a_p"
	snapshot "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_snapshots"
	key "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants_ssh_keys"
)

// Helper methods that will be used by the client classes

// IBMPIHelperClient
type IBMPIClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	ctx             context.Context

	instanceRequest instance.ClientService
	keyRequest      key.ClientService
	sapRequest      sap.ClientService
	snapshotRequest snapshot.ClientService
}

// NewIBMPIClient
func NewIBMPIClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIClient {
	return &IBMPIClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		ctx:             ctx,

		instanceRequest: sess.Power.PCloudpVMInstances,
		keyRequest:      sess.Power.PCloudTenantsSSHKeys,
		sapRequest:      sess.Power.PCloudsap,
		snapshotRequest: sess.Power.PCloudSnapshots,
	}
}
