package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants_ssh_keys"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPIKeyClient ...
type IBMPIKeyClient struct {
	auth     runtime.ClientAuthInfoWriter
	context  context.Context
	request  params.ClientService
	tenantID string
}

// NewIBMPIKeyClient ...
func NewIBMPIKeyClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIKeyClient {
	return &IBMPIKeyClient{
		auth:     sess.AuthInfo(cloudInstanceID),
		context:  ctx,
		request:  sess.Power.PCloudTenantsSSHKeys,
		tenantID: sess.Options.UserAccount,
	}
}

// Get a SSH Key
func (f *IBMPIKeyClient) Get(keyName string) (*models.SSHKey, error) {

	// Create params and send request
	params := &params.PcloudTenantsSshkeysGetParams{
		Context:    f.context,
		SshkeyName: keyName,
		TenantID:   f.tenantID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudTenantsSshkeysGet(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.GetPIKeyOperationFailed, keyName, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Key %s", keyName)
	}
	return resp.Payload, nil
}

// Get All SSH Keys
func (f *IBMPIKeyClient) GetAll() (*models.SSHKeys, error) {

	// Create params and send request
	params := &params.PcloudTenantsSshkeysGetallParams{
		Context:  f.context,
		TenantID: f.tenantID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudTenantsSshkeysGetall(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PI Keys: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Keys")
	}
	return resp.Payload, nil
}

// Create a SSH Key
func (f *IBMPIKeyClient) Create(keyBody *models.SSHKey) (*models.SSHKey, error) {

	// Create params and send request
	params := &params.PcloudTenantsSshkeysPostParams{
		Body:     keyBody,
		Context:  f.context,
		TenantID: f.tenantID,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	respOk, respCreated, err := f.request.PcloudTenantsSshkeysPost(params, f.auth)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf(errors.CreatePIKeyOperationFailed, err)
	}
	if respOk != nil && respOk.Payload != nil {
		return respOk.Payload, nil
	}
	if respCreated != nil && respCreated.Payload != nil {
		return respCreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create PI Key")
}

// Delete a SSH Key
func (f *IBMPIKeyClient) Delete(keyName string) error {

	// Create params and send request
	params := &params.PcloudTenantsSshkeysDeleteParams{
		Context:    f.context,
		SshkeyName: keyName,
		TenantID:   f.tenantID,
	}
	//params.SetTimeout(helpers.PIDeleteTimeOut)
	_, err := f.request.PcloudTenantsSshkeysDelete(params, f.auth)

	// Handle errors
	if err != nil {
		return fmt.Errorf(errors.DeletePIKeyOperationFailed, keyName, err)
	}
	return nil
}
