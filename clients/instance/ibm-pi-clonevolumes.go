package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_clone_volumes"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
)

type IBMPICloneVolumeClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewSnapShotClient ...
func NewIBMPICloneVolumeClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPICloneVolumeClient {
	return &IBMPICloneVolumeClient{
		sess, powerinstanceid,
	}
}

//Create a clone volume
func (f *IBMPICloneVolumeClient) Create(id, powerinstanceid string) (*models.CloneVolumesResponse, error) {

	log.Printf("Calling the CloneVolumeget Method..")
	log.Printf("The input snapshot name is %s and  to the cloudinstance id %s", id, powerinstanceid)

	params := p_cloud_clone_volumes.NewPcloudClonevolumesPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudCloneVolumes.PcloudClonevolumesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
