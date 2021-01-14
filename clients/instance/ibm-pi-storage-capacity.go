package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_storage_capacity"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
	"time"
)

type IBMPIStorageCapacityClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIStorageCapacityClient ...
func NewIBMPIStorageCapacityClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIStorageCapacityClient {
	return &IBMPIStorageCapacityClient{
		sess, powerinstanceid,
	}
}

//Get information about all the storage pools
func (f *IBMPIStorageCapacityClient) GetAll(powerinstanceid string, timeout time.Duration) (*models.StoragePoolsCapacity, error) {
	log.Printf("Calling the Get AllStorage Pools Method..")
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityPoolsGetallParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityPoolsGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
