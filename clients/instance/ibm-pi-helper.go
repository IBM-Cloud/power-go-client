package instance

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	job "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_jobs"
	network "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	placementGroup "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_placement_groups"
	dhcp "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_service_d_h_c_p"
)

// Helper methods that will be used by the client classes

// IBMPIHelperClient
type IBMPIClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	ctx             context.Context

	dhcpRequest           dhcp.ClientService
	jobRequest            job.ClientService
	networkRequest        network.ClientService
	placementGroupRequest placementGroup.ClientService
}

// NewIBMPIClient
func NewIBMPIClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIClient {
	return &IBMPIClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		ctx:             ctx,

		dhcpRequest:           sess.Power.PCloudServicedhcp,
		jobRequest:            sess.Power.PCloudJobs,
		placementGroupRequest: sess.Power.PCloudPlacementGroups,
		networkRequest:        sess.Power.PCloudNetworks,
	}
}
