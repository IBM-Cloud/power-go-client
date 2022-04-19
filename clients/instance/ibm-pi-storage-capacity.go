package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_storage_capacity"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPIStorageCapacityClient ..
type IBMPIStorageCapacityClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPIStorageCapacityClient ...
func NewIBMPIStorageCapacityClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIStorageCapacityClient {
	return &IBMPIStorageCapacityClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudStorageCapacity,
	}
}

// Get All Storage Capacities for all available Storage Pools in a Region
func (f *IBMPIStorageCapacityClient) GetAllStoragePoolsCapacity() (*models.StoragePoolsCapacity, error) {

	// Create params and send request
	params := &params.PcloudStoragecapacityPoolsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudStoragecapacityPoolsGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage pools: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage pools")
	}
	return resp.Payload, nil
}

// Get Storage Capacity for a Storage Pool in a Region
func (f *IBMPIStorageCapacityClient) GetStoragePoolCapacity(storagePool string) (*models.StoragePoolCapacity, error) {

	// Create params and send request
	params := &params.PcloudStoragecapacityPoolsGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		StoragePoolName: storagePool,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudStoragecapacityPoolsGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for storage pool %s: %w", storagePool, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for storage pool %s", storagePool)
	}
	return resp.Payload, nil
}

// Get Storage capacity for a Storage Type in a Region
func (f *IBMPIStorageCapacityClient) GetStorageTypeCapacity(storageType string) (*models.StorageTypeCapacity, error) {

	// Create params and send request
	params := &params.PcloudStoragecapacityTypesGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		StorageTypeName: storageType,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudStoragecapacityTypesGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for storage type %s: %w", storageType, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for storage type %s", storageType)
	}
	return resp.Payload, nil
}

// Get Storage Capacity for all available Storage Types in a Region
func (f *IBMPIStorageCapacityClient) GetAllStorageTypesCapacity() (*models.StorageTypesCapacity, error) {

	// Create params and send request
	params := &params.PcloudStoragecapacityTypesGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudStoragecapacityTypesGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage types %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage types")
	}
	return resp.Payload, nil
}
