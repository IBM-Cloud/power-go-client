package instance

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"

	cloudConnection "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_cloud_connections"
	image "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	cloudInstance "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_instances"
)

// Helper methods that will be used by the client classes

// IBMPIHelperClient
type IBMPIClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	ctx             context.Context

	cloudConnectionRequest cloudConnection.ClientService
	cloudInstanceRequest   cloudInstance.ClientService
	imageRequest           image.ClientService
}

// NewIBMPIClient
func NewIBMPIClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIClient {
	return &IBMPIClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		ctx:             ctx,

		cloudConnectionRequest: sess.Power.PCloudCloudConnections,
		cloudInstanceRequest:   sess.Power.PCloudInstances,
		imageRequest:           sess.Power.PCloudImages,
	}
}
