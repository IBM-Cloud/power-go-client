package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	p "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants_ssh_keys"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIKeyClient.
type IBMPIKeyClient struct {
	IBMPIClient
}

// NewIBMPIKeyClient.
func NewIBMPIKeyClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIKeyClient {
	return &IBMPIKeyClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a SSH Key.
func (f *IBMPIKeyClient) Get(id string) (*models.SSHKey, error) {
	var tenantid = f.session.Options.UserAccount
	params := p.NewPcloudTenantsSshkeysGetParams().
		WithContext(f.ctx).
		WithSshkeyName(id).
		WithTenantID(tenantid).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.keyRequest.PcloudTenantsSshkeysGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetPIKeyOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Key %s", id)
	}
	return resp.Payload, nil
}

// Get All SSH Keys.
func (f *IBMPIKeyClient) GetAll() (*models.SSHKeys, error) {
	var tenantid = f.session.Options.UserAccount
	params := p.NewPcloudTenantsSshkeysGetallParams().
		WithContext(f.ctx).
		WithTenantID(tenantid).
		WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.keyRequest.PcloudTenantsSshkeysGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PI Keys: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PI Keys")
	}
	return resp.Payload, nil
}

// Create a SSH Key.
func (f *IBMPIKeyClient) Create(body *models.SSHKey) (*models.SSHKey, error) {
	var tenantid = f.session.Options.UserAccount
	params := p.NewPcloudTenantsSshkeysPostParams().
		WithBody(body).
		WithContext(f.ctx).
		WithTenantID(tenantid).
		WithTimeout(helpers.PICreateTimeOut)
	postok, postcreated, err := f.keyRequest.PcloudTenantsSshkeysPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreatePIKeyOperationFailed, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	if postcreated != nil && postcreated.Payload != nil {
		return postcreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create PI Key")
}

// Delete a SSH Key.
func (f *IBMPIKeyClient) Delete(id string) error {
	var tenantid = f.session.Options.UserAccount
	params := p.NewPcloudTenantsSshkeysDeleteParams().
		WithContext(f.ctx).
		WithSshkeyName(id).
		WithTenantID(tenantid).
		WithTimeout(helpers.PIDeleteTimeOut)
	_, err := f.keyRequest.PcloudTenantsSshkeysDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeletePIKeyOperationFailed, id, err)
	}
	return nil
}
