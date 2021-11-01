package instance

import (
	"fmt"
	"strings"
	"time"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_a_p"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// temporary fix to OCP 504 error
// https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3178
const CreateRetries = 6

/*  ChangeLog

2020-June-05 : Added the timeout variable to the clients since a lot of the SB / Powervc calls are timing out.

*/

// IBMPIInstanceClient ...
type IBMPIInstanceClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIInstanceClient ...
func NewIBMPIInstanceClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIInstanceClient {
	return &IBMPIInstanceClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

//Get information about a single pvm only
func (f *IBMPIInstanceClient) Get(id, powerinstanceid string, timeout time.Duration) (*models.PVMInstance, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesGetParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get PVM Instance %s :%v", id, err)
	}
	return resp.Payload, nil
}

// GetAll Information about all the PVM Instances for a Client
func (f *IBMPIInstanceClient) GetAll(powerinstanceid string, timeout time.Duration) (*models.PVMInstances, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesGetallParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get all PVM Instances of Power Instance %s :%v", powerinstanceid, err)
	}
	return resp.Payload, nil
}

// GetServerFromName
// temporary fix to OCP 504 error
// https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3178
func (f *IBMPIInstanceClient) GetServerFromName(serverName, powerinstanceid string, timeout time.Duration) (*models.PVMInstanceList, error) {

	// get all
	servers, err := f.GetAll(powerinstanceid, timeout)
	if err != nil || servers == nil {
		return nil, fmt.Errorf("Failed to Get all PVM Instances of Power Instance %s. Searching for server name:%v", powerinstanceid, err)
	}

	// search all servers for server with name
	for _, server := range servers.PvmInstances {

		if *server.ServerName == serverName {

			// Need to run a Get call to return PVMInstance not PVMInstanceReference
			//
			// Create -> models.PVMInstanceList -> PVMInstance
			// Get ----> models.PVMInstance
			// GetAll -> models.PVMInstances ----> PVMInstanceReference

			serverInfo, err := f.Get(*server.PvmInstanceID, powerinstanceid, helpers.PIGetTimeOut)
			if err != nil {
				return nil, fmt.Errorf("Failed to Get PVM Instance with name %s :%v", serverName, err)
			}
			instanceList := new(models.PVMInstanceList)
			*instanceList = append(*instanceList, serverInfo)
			return instanceList, nil
		}
	}
	return nil, fmt.Errorf("Failed to Get PVM Instance with name %s :%v", serverName, err)
}

//Create ...
func (f *IBMPIInstanceClient) Create(powerdef *p_cloud_p_vm_instances.PcloudPvminstancesPostParams, powerinstanceid string, timeout time.Duration) (*models.PVMInstanceList, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(powerdef.Body)
	postok, postcreated, postAccepted, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		// temporary fix to OCP 504 error
		// https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3178
		if strings.Contains(err.Error(), "504") {
			// 504 Error
			// Retry Create Call
			for retry := 0; retry <= CreateRetries; retry++ {
				result, err := f.CreateRetry(powerdef, powerinstanceid, timeout)
				if result != nil {
					return result, err
				}
			}
			// didn't work after retries
			return nil, nil
		} else if strings.Contains(err.Error(), "server name already exists") || strings.Contains(err.Error(), "context deadline exceeded") {
			// 409 error
			// Return actual instance. Get instance info using name
			server, err := f.GetServerFromName(*powerdef.Body.ServerName, powerinstanceid, timeout)
			if err != nil {
				return nil, err
			}
			return server, nil
		} else {
			// normal error handling
			return nil, fmt.Errorf("Failed to Create PVM Instance: %v", err)
		}
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
	return nil, nil
}

// Create Retry
// temporary fix to OCP 504 error
// https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3178
func (f *IBMPIInstanceClient) CreateRetry(powerdef *p_cloud_p_vm_instances.PcloudPvminstancesPostParams, powerinstanceid string, timeout time.Duration) (*models.PVMInstanceList, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(powerdef.Body)
	postok, postcreated, postAccepted, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		// temporary fix to OCP 504 error
		// https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3178
		if strings.Contains(err.Error(), "504") {
			// 504 Error
			// Will retry in main loop
			return nil, nil
		} else if strings.Contains(err.Error(), "server name already exists") || strings.Contains(err.Error(), "context deadline exceeded") {
			// 409 error
			// Return actual instance. Get instance info using name
			server, err := f.GetServerFromName(*powerdef.Body.ServerName, powerinstanceid, timeout)
			if err != nil {
				return nil, err
			}
			return server, nil
		} else {
			// normal error handling
			return nil, fmt.Errorf("Failed to Create PVM Instance: %v", err)
		}
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
	return nil, nil
}

// Delete PVM Instances
func (f *IBMPIInstanceClient) Delete(id, powerinstanceid string, timeout time.Duration) error {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesDeleteParamsWithTimeout(helpers.PIDeleteTimeOut).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)
	_, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return fmt.Errorf("Failed to Delete PVM Instance %s :%s", id, err)
	}

	return nil
}

// Update PVM Instances
func (f *IBMPIInstanceClient) Update(id, powerinstanceid string, powerupdateparams *p_cloud_p_vm_instances.PcloudPvminstancesPutParams, timeout time.Duration) (*models.PVMInstanceUpdateResponse, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPutParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(powerupdateparams.Body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Update PVM Instance %s :%s", id, err)
	}
	return resp.Payload, nil
}

// Action PVM Instances Operations
func (f *IBMPIInstanceClient) Action(poweractionparams *p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams, id, powerinstanceid string, timeout time.Duration) (models.Object, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesActionPostParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(poweractionparams.Body)
	postok, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesActionPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf("Failed to Action PVM Instance :%s", err)
	}

	return postok.Payload, nil

}

// PostConsoleURL Generate the Console URL
func (f *IBMPIInstanceClient) PostConsoleURL(id, powerinstanceid string, timeout time.Duration) (models.Object, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesConsolePostParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)
	postok, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesConsolePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf("Failed to Generate the Console URL PVM Instance:%s", err)
	}
	return postok.Payload, nil
}

// CaptureInstanceToImageCatalog Captures an instance
func (f *IBMPIInstanceClient) CaptureInstanceToImageCatalog(id, powerinstanceid string, picaptureparams *p_cloud_p_vm_instances.PcloudPvminstancesCapturePostParams, timeout time.Duration) (models.Object, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesCapturePostParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(picaptureparams.Body)
	postok, _, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesCapturePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf("Failed to Generate the Console URL PVM Instance:%s", err)
	}
	return postok.Payload, nil

}

// CreatePvmSnapShot Create a snapshot of the instance
func (f *IBMPIInstanceClient) CreatePvmSnapShot(snapshotdef *p_cloud_p_vm_instances.PcloudPvminstancesSnapshotsPostParams, pvminstanceid, powerinstanceid string, timeout time.Duration) (*models.SnapshotCreateResponse, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsPostParamsWithTimeout(helpers.PICreateTimeOut).WithPvmInstanceID(pvminstanceid).WithCloudInstanceID(powerinstanceid).WithBody(snapshotdef.Body)
	snapshotpostaccepted, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || snapshotpostaccepted == nil {
		return nil, fmt.Errorf("Failed to Create the snapshot %s for the pvminstance : %s", pvminstanceid, err)
	}
	return snapshotpostaccepted.Payload, nil
}

// CreateClone ...
func (f *IBMPIInstanceClient) CreateClone(clonedef *p_cloud_p_vm_instances.PcloudPvminstancesClonePostParams, pvminstanceid, powerinstanceid string) (*models.PVMInstance, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesClonePostParamsWithTimeout(helpers.PICreateTimeOut).WithPvmInstanceID(pvminstanceid).WithCloudInstanceID(powerinstanceid).WithBody(clonedef.Body)
	clonePost, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesClonePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf("Failed to create the clone of the pvm instance %s", err)
	}
	return clonePost.Payload, nil
}

// GetSnapShotVM Get information about the snapshots for a vm
func (f *IBMPIInstanceClient) GetSnapShotVM(powerinstanceid, pvminstanceid string, timeout time.Duration) (*models.Snapshots, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsGetallParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(pvminstanceid)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to Get the snapshot for the pvminstance [%s]: %v", pvminstanceid, err)
	}
	return resp.Payload, nil

}

// RestoreSnapShotVM Restore a snapshot
func (f *IBMPIInstanceClient) RestoreSnapShotVM(powerinstanceid, pvminstanceid, snapshotid, restoreAction string, restoreparams *p_cloud_p_vm_instances.PcloudPvminstancesSnapshotsRestorePostParams, timeout time.Duration) (*models.Snapshot, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsRestorePostParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(pvminstanceid).WithSnapshotID(snapshotid).WithRestoreFailAction(&restoreAction).WithBody(restoreparams.Body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsRestorePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to restrore the snapshot for the pvminstance [%s]: %v", pvminstanceid, err)
	}
	return resp.Payload, nil
}

// AddNetwork Add a network to the instance
func (f *IBMPIInstanceClient) AddNetwork(powerinstanceid, pvminstanceid string, networkdef *p_cloud_p_vm_instances.PcloudPvminstancesNetworksPostParams, timeout time.Duration) (*models.PVMInstanceNetwork, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesNetworksPostParamsWithTimeout(helpers.PICreateTimeOut).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(pvminstanceid).WithBody(networkdef.Body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesNetworksPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload.NetworkID == "" {
		return nil, fmt.Errorf("Failed to attach the network to the pvminstanceid %s : %v", pvminstanceid, err)
	}
	return resp.Payload, nil
}

// Delete a network from an instance

// CreateSAP Create SAP Systems
func (f *IBMPIInstanceClient) CreateSAP(powerdef *p_cloud_s_a_p.PcloudSapPostParams, powerinstanceid string, timeout time.Duration) (*models.PVMInstanceList, error) {

	params := p_cloud_s_a_p.NewPcloudSapPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(powerdef.Body)
	postok, postcreated, postAccepted, err := f.session.Power.PCloudSAP.PcloudSapPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, fmt.Errorf("Failed to create sap to the power instance %s : %s", powerinstanceid, err)
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

	//return &postok.Payload, nil
	return nil, nil
}

// GetSAPProfiles Get All SAP Profiles
func (f *IBMPIInstanceClient) GetSAPProfiles(powerinstanceid string) (*models.SAPProfiles, error) {

	params := p_cloud_s_a_p.NewPcloudSapGetallParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudSAP.PcloudSapGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("Failed to get sap profiles to the power instance %s : %s", powerinstanceid, err)
	}
	return resp.Payload, nil
}

// GetSap Get an SAP profile
func (f *IBMPIInstanceClient) GetSap(powerinstanceid, sapprofileID string) (*models.SAPProfile, error) {
	params := p_cloud_s_a_p.NewPcloudSapGetParamsWithTimeout(helpers.PIGetTimeOut).WithCloudInstanceID(powerinstanceid).WithSapProfileID(sapprofileID)
	resp, err := f.session.Power.PCloudSAP.PcloudSapGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf("Failed to get sap profile %s to the power instance %s : %s", sapprofileID, powerinstanceid, err)
	}
	return resp.Payload, nil

}
