package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	params "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_a_p"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPISAPInstanceClient ...
type IBMPISAPInstanceClient struct {
	auth            runtime.ClientAuthInfoWriter
	cloudInstanceID string
	context         context.Context
	request         params.ClientService
}

// NewIBMPISAPInstanceClient ...
func NewIBMPISAPInstanceClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPISAPInstanceClient {
	return &IBMPISAPInstanceClient{
		auth:            sess.AuthInfo(cloudInstanceID),
		cloudInstanceID: cloudInstanceID,
		context:         ctx,
		request:         sess.Power.PCloudsap,
	}
}

// Create a SAP Instance
func (f *IBMPISAPInstanceClient) Create(createBody *models.SAPCreate) (*models.PVMInstanceList, error) {

	// Create params and send request
	params := &params.PcloudSapPostParams{
		Body:            createBody,
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PICreateTimeOut)
	respOk, respCreated, respAccepted, err := f.request.PcloudSapPost(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to Create SAP Instance: %w", err)
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
	return nil, fmt.Errorf("failed to Create SAP Instance")
}

// Get a SAP Profile
func (f *IBMPISAPInstanceClient) GetSAPProfile(sapProfileID string) (*models.SAPProfile, error) {

	// Create params and send request
	params := &params.PcloudSapGetParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
		SapProfileID:    sapProfileID,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudSapGet(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to get sap profile %s : %w", sapProfileID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get sap profile %s", sapProfileID)
	}
	return resp.Payload, nil
}

// Get All SAP Profiles
func (f *IBMPISAPInstanceClient) GetAllSAPProfiles(cloudInstanceID string) (*models.SAPProfiles, error) {

	// Create params and send request
	params := &params.PcloudSapGetallParams{
		CloudInstanceID: f.cloudInstanceID,
		Context:         f.context,
	}
	//params.SetTimeout(helpers.PIGetTimeOut)
	resp, err := f.request.PcloudSapGetall(params, f.auth)

	// Handle Errors
	if err != nil {
		return nil, fmt.Errorf("failed to get all sap profiles for power instance %s: %w", cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all sap profiles for power instance %s", cloudInstanceID)
	}
	return resp.Payload, nil
}
