package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_storage_capacity"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIStorageCapacityClient.
type IBMPIStorageCapacityClient struct {
	IBMPIClient
}

// NewIBMPIStorageCapacityClient.
func NewIBMPIStorageCapacityClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIStorageCapacityClient {
	return &IBMPIStorageCapacityClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get All Storage Capacities for all available Storage Pools in a Region.
func (f *IBMPIStorageCapacityClient) GetAllStoragePoolsCapacity() (*models.StoragePoolsCapacity, error) {
	params := p.NewPcloudStoragecapacityPoolsGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.storageCapacityRequest.PcloudStoragecapacityPoolsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage pools: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage pools")
	}
	return resp.Payload, nil
}

// Get Storage Capacity for a Storage Pool in a Region.
func (f *IBMPIStorageCapacityClient) GetStoragePoolCapacity(storagePool string) (*models.StoragePoolCapacity, error) {
	params := p.NewPcloudStoragecapacityPoolsGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithStoragePoolName(storagePool).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.storageCapacityRequest.PcloudStoragecapacityPoolsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for storage pool %s: %w", storagePool, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for storage pool %s", storagePool)
	}
	return resp.Payload, nil
}

// Get Storage capacity for a Storage Type in a Region.
func (f *IBMPIStorageCapacityClient) GetStorageTypeCapacity(storageType string) (*models.StorageTypeCapacity, error) {
	params := p.NewPcloudStoragecapacityTypesGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithStorageTypeName(storageType).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.storageCapacityRequest.PcloudStoragecapacityTypesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for storage type %s: %w", storageType, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for storage type %s", storageType)
	}
	return resp.Payload, nil
}

// Get Storage Capacity for all available Storage Types in a Region.
func (f *IBMPIStorageCapacityClient) GetAllStorageTypesCapacity() (*models.StorageTypesCapacity, error) {
	params := p.NewPcloudStoragecapacityTypesGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.storageCapacityRequest.PcloudStoragecapacityTypesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage types %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage types")
	}
	return resp.Payload, nil
}
