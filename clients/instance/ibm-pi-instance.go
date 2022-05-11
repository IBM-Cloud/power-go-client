package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIInstanceClient.
type IBMPIInstanceClient struct {
	IBMPIClient
}

// NewIBMPIInstanceClient.
func NewIBMPIInstanceClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIInstanceClient {
	return &IBMPIInstanceClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get an Instance.
func (f *IBMPIInstanceClient) Get(id string) (*models.PVMInstance, error) {
	params := p.NewPcloudPvminstancesGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get PVM Instance %s :%w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PVM Instance %s", id)
	}
	return resp.Payload, nil
}

// Get All Instances.
func (f *IBMPIInstanceClient) GetAll() (*models.PVMInstances, error) {
	params := p.NewPcloudPvminstancesGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PVM Instances of Power Instance %s :%w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PVM Instances of Power Instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create an Instance.
func (f *IBMPIInstanceClient) Create(body *models.PVMInstanceCreate) (*models.PVMInstanceList, error) {
	params := p.NewPcloudPvminstancesPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	postok, postcreated, postAccepted, err := f.instanceRequest.PcloudPvminstancesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Create PVM Instance :%w", err)
	}
	if postok != nil && len(postok.Payload) > 0 {
		return &postok.Payload, nil
	}
	if postcreated != nil && len(postcreated.Payload) > 0 {
		return &postcreated.Payload, nil
	}
	if postAccepted != nil && len(postAccepted.Payload) > 0 {
		return &postAccepted.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create PVM Instance")
}

// Delete an Instance.
func (f *IBMPIInstanceClient) Delete(id string) error {
	params := p.NewPcloudPvminstancesDeleteParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.instanceRequest.PcloudPvminstancesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PVM Instance %s :%w", id, err)
	}
	return nil
}

// Update an Instance.
func (f *IBMPIInstanceClient) Update(id string, body *models.PVMInstanceUpdate) (*models.PVMInstanceUpdateResponse, error) {
	params := p.NewPcloudPvminstancesPutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Update PVM Instance %s :%w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update PVM Instance %s", id)
	}
	return resp.Payload, nil
}

// Action Operation for an Instance.
func (f *IBMPIInstanceClient) Action(id string, body *models.PVMInstanceAction) error {
	params := p.NewPcloudPvminstancesActionPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut)
	_, err := f.instanceRequest.PcloudPvminstancesActionPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to perform Action on PVM Instance %s :%w", id, err)
	}
	return nil

}

// Generate the Console URL for an Instance.
func (f *IBMPIInstanceClient) PostConsoleURL(id string) (*models.PVMInstanceConsole, error) {
	params := p.NewPcloudPvminstancesConsolePostParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut)
	postok, err := f.instanceRequest.PcloudPvminstancesConsolePost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Generate the Console URL PVM Instance %s :%w", id, err)
	}
	if postok == nil || postok.Payload == nil {
		return nil, fmt.Errorf("failed to Generate the Console URL PVM Instance %s", id)
	}
	return postok.Payload, nil
}

// List the available Console Languages for an Instance.
func (f *IBMPIInstanceClient) GetConsoleLanguages(id string) (*models.ConsoleLanguages, error) {
	params := p.NewPcloudPvminstancesConsoleGetParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesConsoleGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get Console Languages for PVM Instance %s :%w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Console Languages for PVM Instance %s", id)
	}
	return resp.Payload, nil
}

// Update the available Console Languages for an Instance.
func (f *IBMPIInstanceClient) UpdateConsoleLanguage(id string, body *models.ConsoleLanguage) (*models.ConsoleLanguage, error) {
	params := p.NewPcloudPvminstancesConsolePutParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesConsolePut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Update Console Language for PVM Instance %s :%w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update Console Language for PVM Instance %s", id)
	}
	return resp.Payload, nil
}

// Capture an Instance.
func (f *IBMPIInstanceClient) CaptureInstanceToImageCatalog(id string, body *models.PVMInstanceCapture) error {
	params := p.NewPcloudPvminstancesCapturePostParams().
		WithBody(body).
		WithContext(f.ctx).
		WithCloudInstanceID(f.cloudInstanceID).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIGetTimeOut)
	_, _, err := f.instanceRequest.PcloudPvminstancesCapturePost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Capture the PVM Instance %s: %w", id, err)
	}
	return nil

}

// Capture an Instance (V2).
func (f *IBMPIInstanceClient) CaptureInstanceToImageCatalogV2(id string, body *models.PVMInstanceCapture) (*models.JobReference, error) {
	params := p.NewPcloudV2PvminstancesCapturePostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.instanceRequest.PcloudV2PvminstancesCapturePost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Capture the PVM Instance %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Capture the PVM Instance %s", id)
	}
	return resp.Payload, nil
}

// Create a snapshot of an Instance.
func (f *IBMPIInstanceClient) CreatePvmSnapShot(id string, body *models.SnapshotCreate) (*models.SnapshotCreateResponse, error) {
	params := p.NewPcloudPvminstancesSnapshotsPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut).
		WithPvmInstanceID(id)
	snapshotpostaccepted, err := f.instanceRequest.PcloudPvminstancesSnapshotsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Create the snapshot for the pvminstance %s: %w", id, err)
	}
	if snapshotpostaccepted == nil || snapshotpostaccepted.Payload == nil {
		return nil, fmt.Errorf("failed to Create the snapshot for the pvminstance %s", id)
	}
	return snapshotpostaccepted.Payload, nil
}

// Create a Clone of an Instance.
func (f *IBMPIInstanceClient) CreateClone(id string, body *models.PVMInstanceClone) (*models.PVMInstance, error) {
	params := p.NewPcloudPvminstancesClonePostParams().
		WithBody(body).
		WithContext(f.ctx).
		WithCloudInstanceID(f.cloudInstanceID).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut)
	clonePost, err := f.instanceRequest.PcloudPvminstancesClonePost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to create the clone of the pvm instance %s: %w", id, err)
	}
	if clonePost == nil || clonePost.Payload == nil {
		return nil, fmt.Errorf("failed to create the clone of the pvm instance %s", id)
	}
	return clonePost.Payload, nil
}

// Get an Instance's Snapshots.
func (f *IBMPIInstanceClient) GetSnapShotVM(id string) (*models.Snapshots, error) {
	params := p.NewPcloudPvminstancesSnapshotsGetallParams().
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesSnapshotsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get the snapshot for the pvminstance %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get the snapshot for the pvminstance %s", id)
	}
	return resp.Payload, nil

}

// Restore a Snapshot of an Instance.
func (f *IBMPIInstanceClient) RestoreSnapShotVM(id, snapshotid, restoreAction string, body *models.SnapshotRestore) (*models.Snapshot, error) {
	params := p.NewPcloudPvminstancesSnapshotsRestorePostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithRestoreFailAction(&restoreAction).
		WithSnapshotID(snapshotid).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesSnapshotsRestorePost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to restrore the snapshot for the pvminstance %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to restrore the snapshot for the pvminstance %s", id)
	}
	return resp.Payload, nil
}

// Add a Network to an Instance.
func (f *IBMPIInstanceClient) AddNetwork(id string, body *models.PVMInstanceAddNetwork) (*models.PVMInstanceNetwork, error) {
	params := p.NewPcloudPvminstancesNetworksPostParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithPvmInstanceID(id).
		WithContext(f.ctx).
		WithTimeout(helpers.PICreateTimeOut)
	resp, err := f.instanceRequest.PcloudPvminstancesNetworksPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to attach the network to the pvminstanceid %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to attach the network to the pvminstanceid %s", id)
	}
	return resp.Payload, nil
}

// Delete a Network from an Instance.
func (f *IBMPIInstanceClient) DeleteNetwork(id string, body *models.PVMInstanceRemoveNetwork) error {
	params := p.NewPcloudPvminstancesNetworksDeleteParams().
		WithBody(body).
		WithCloudInstanceID(f.cloudInstanceID).
		WithContext(f.ctx).
		WithPvmInstanceID(id).
		WithTimeout(helpers.PICreateTimeOut)
	_, err := f.instanceRequest.PcloudPvminstancesNetworksDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to delete the network to the pvminstanceid %s: %w", id, err)
	}
	return nil
}
