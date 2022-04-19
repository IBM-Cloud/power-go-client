package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPIInstanceClient ...
type IBMPIInstanceClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPIInstanceClient ...
func NewIBMPIInstanceClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIInstanceClient {
	return &IBMPIInstanceClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudpVMInstances,
	}
}

// Get an Instance
func (f *IBMPIInstanceClient) Get(instanceID string) (*models.PVMInstance, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudPvminstancesGet(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get PVM Instance %s :%w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PVM Instance %s", instanceID)
	}
	return resp.Payload, nil
}

// Get All Instances
func (f *IBMPIInstanceClient) GetAll() (*models.PVMInstances, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudPvminstancesGetall(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PVM Instances of Power Instance %s :%w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PVM Instances of Power Instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create an Instance
func (f *IBMPIInstanceClient) Create(createBody *models.PVMInstanceCreate) (*models.PVMInstanceList, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	respOk, respCreated, respAccepted, err := f.request.PcloudPvminstancesPost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Create PVM Instance :%w", err)
	}
	if respOk != nil && len(respOk.Payload) > 0 {
		return &respOk.Payload, nil
	}
	if respCreated != nil && len(respCreated.Payload) > 0 {
		return &respCreated.Payload, nil
	}
	if respAccepted != nil && len(respAccepted.Payload) > 0 {
		return &respAccepted.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create PVM Instance")
}

// Delete an Instance
func (f *IBMPIInstanceClient) Delete(instanceID string) error {

	// Create params and send request
	params := &params.PcloudPvminstancesDeleteParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudPvminstancesDelete(params, f.auth)

	// Handle Errors
	if err != nil {
		return fmt.Errorf("failed to Delete PVM Instance %s :%w", instanceID, err)
	}
	return nil
}

// Update an Instance
func (f *IBMPIInstanceClient) Update(instanceID string, updateBody *models.PVMInstanceUpdate) (*models.PVMInstanceUpdateResponse, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesPutParams{
		Body:            updateBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudPvminstancesPut(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Update PVM Instance %s :%w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update PVM Instance %s", instanceID)
	}
	return resp.Payload, nil
}

// Action Operation for an Instance
func (f *IBMPIInstanceClient) Action(instanceID string, actionBody *models.PVMInstanceAction) error {

	// Create params and send request
	params := &params.PcloudPvminstancesActionPostParams{
		Body:            actionBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	_, err := f.request.PcloudPvminstancesActionPost(params, f.auth)

	// Handle Errors
	if err != nil {
		return fmt.Errorf("failed to perform Action on PVM Instance %s :%w", instanceID, err)
	}
	return nil

}

// Generate the Console URL for an Instance
func (f *IBMPIInstanceClient) PostConsoleURL(instanceID string) (*models.PVMInstanceConsole, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesConsolePostParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	respOk, err := f.request.PcloudPvminstancesConsolePost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Generate the Console URL PVM Instance %s :%w", instanceID, err)
	}
	if respOk == nil || respOk.Payload == nil {
		return nil, fmt.Errorf("failed to Generate the Console URL PVM Instance %s", instanceID)
	}
	return respOk.Payload, nil
}

// List the available Console Languages for an Instance
func (f *IBMPIInstanceClient) GetConsoleLanguages(instanceID string) (*models.ConsoleLanguages, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesConsoleGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudPvminstancesConsoleGet(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get Console Languages for PVM Instance %s :%w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Console Languages for PVM Instance %s", instanceID)
	}
	return resp.Payload, nil
}

// Update the available Console Languages for an Instance
func (f *IBMPIInstanceClient) UpdateConsoleLanguage(instanceID string, updateBody *models.ConsoleLanguage) (*models.ConsoleLanguage, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesConsolePutParams{
		Body:            updateBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIUpdateTimeOut)
	resp, err := f.request.PcloudPvminstancesConsolePut(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Update Console Language for PVM Instance %s :%w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update Console Language for PVM Instance %s", instanceID)
	}
	return resp.Payload, nil
}

// Capture an Instance
func (f *IBMPIInstanceClient) CaptureInstanceToImageCatalog(instanceID string, captureBody *models.PVMInstanceCapture) error {

	// Create params and send request
	params := &params.PcloudPvminstancesCapturePostParams{
		Body:            captureBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	_, _, err := f.request.PcloudPvminstancesCapturePost(params, f.auth)

	// Handle Errors
	if err != nil {
		return fmt.Errorf("failed to Capture the PVM Instance %s: %w", instanceID, err)
	}
	return nil

}

// Capture an Instance (V2)
func (f *IBMPIInstanceClient) CaptureInstanceToImageCatalogV2(instanceID string, captureBody *models.PVMInstanceCapture) (*models.JobReference, error) {

	// Create params and send request
	params := &params.PcloudV2PvminstancesCapturePostParams{
		Body:            captureBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudV2PvminstancesCapturePost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Capture the PVM Instance %s: %w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Capture the PVM Instance %s", instanceID)
	}
	return resp.Payload, nil
}

// Create a snapshot of an Instance
func (f *IBMPIInstanceClient) CreatePvmSnapShot(instanceID string, createBody *models.SnapshotCreate) (*models.SnapshotCreateResponse, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesSnapshotsPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudPvminstancesSnapshotsPost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Create the snapshot for the pvminstance %s: %w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create the snapshot for the pvminstance %s", instanceID)
	}
	return resp.Payload, nil
}

// Creatte a CLone of an Instance
func (f *IBMPIInstanceClient) CreateClone(instanceID string, createBody *models.PVMInstanceClone) (*models.PVMInstance, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesClonePostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudPvminstancesClonePost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to create the clone of the pvm instance %s: %w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to create the clone of the pvm instance %s", instanceID)
	}
	return resp.Payload, nil
}

// Get a Snapshot of an Instance
func (f *IBMPIInstanceClient) GetSnapShotVM(instanceID string) (*models.Snapshots, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesSnapshotsGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudPvminstancesSnapshotsGetall(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get the snapshot for the pvminstance %s: %w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get the snapshot for the pvminstance %s", instanceID)
	}
	return resp.Payload, nil

}

// Restore a Snapshot of an Instance
func (f *IBMPIInstanceClient) RestoreSnapShotVM(instanceID, snapshotID, restoreAction string, updateBody *models.SnapshotRestore) (*models.Snapshot, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesSnapshotsRestorePostParams{
		Body:              updateBody,
		CloudInstanceID:   f.cloudInstanceID,
		Context:           f.context,
		PvmInstanceID:     instanceID,
		RestoreFailAction: &restoreAction,
		SnapshotID:        snapshotID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudPvminstancesSnapshotsRestorePost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to restrore the snapshot for the pvminstance %s: %w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to restrore the snapshot for the pvminstance %s", instanceID)
	}
	return resp.Payload, nil
}

// Add a Network to an Instance
func (f *IBMPIInstanceClient) AddNetwork(instanceID string, updateBody *models.PVMInstanceAddNetwork) (*models.PVMInstanceNetwork, error) {

	// Create params and send request
	params := &params.PcloudPvminstancesNetworksPostParams{
		Body:            updateBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	resp, err := f.request.PcloudPvminstancesNetworksPost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to attach the network to the pvminstanceid %s: %w", instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to attach the network to the pvminstanceid %s", instanceID)
	}
	return resp.Payload, nil
}

// Delete a Network from an Instance
func (f *IBMPIInstanceClient) DeleteNetwork(instanceID string, deleteBody *models.PVMInstanceRemoveNetwork) error {

	// Create params and send request
	params := &params.PcloudPvminstancesNetworksDeleteParams{
		Body:            deleteBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		PvmInstanceID:   instanceID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudPvminstancesNetworksDelete(params, f.auth)

	// Handle Errors
	if err != nil {
		return fmt.Errorf("failed to delete the network to the pvminstanceid %s: %w", instanceID, err)
	}
	return nil
}
