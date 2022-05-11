package instance

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"

	vpn "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_connections"
	vpnPolicy "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_policies"
	volume "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
)

// Helper methods that will be used by the client classes

// IBMPIHelperClient
type IBMPIClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	ctx             context.Context

	volumeRequest    volume.ClientService
	vpnRequest       vpn.ClientService
	vpnPolicyRequest vpnPolicy.ClientService
}

// NewIBMPIClient
func NewIBMPIClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIClient {
	return &IBMPIClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		ctx:             ctx,

		volumeRequest:    sess.Power.PCloudVolumes,
		vpnRequest:       sess.Power.PCloudvpnConnections,
		vpnPolicyRequest: sess.Power.PCloudvpnPolicies,
	}
}
