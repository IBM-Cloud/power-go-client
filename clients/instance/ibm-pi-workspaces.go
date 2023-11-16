package instance

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/workspaces"
	"github.com/IBM-Cloud/power-go-client/power/models"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

// IBMPIWorkspacesClient
type IBMPIWorkspacesClient struct {
	IBMPIClient
}

// NewIBMPIWorkspacesClient
func NewIBMPIWorkspacesClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIWorkspacesClient {
	return &IBMPIWorkspacesClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a workspace
func (f *IBMPIWorkspacesClient) Get(cloudInstanceID string) (*models.Workspace, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := workspaces.NewV1WorkspacesGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithWorkspaceID(cloudInstanceID)
	resp, err := f.session.Power.Workspaces.V1WorkspacesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf(errors.GetWorkspaceOperationFailed, f.cloudInstanceID, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Workspace %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get all workspaces
func (f *IBMPIWorkspacesClient) GetAll() (*models.Workspaces, error) {
	if f.session.IsOnPrem() {
		return nil, fmt.Errorf("operation not supported in satellite location, check documentation")
	}
	params := workspaces.NewV1WorkspacesGetallParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.Workspaces.V1WorkspacesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to Get all Workspaces: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Workspaces")
	}
	return resp.Payload, nil
}

// ResourceControllerV2 struct
// type ResourceControllerV2 struct {
// 	Name           string `json:"name"`
// 	RegionTarget   string `json:"target"`
// 	ResourceGroup  string `json:"resource_group"`
// 	ResourcePlanID string `json:"resource_plan_id"`
// }

// type resourceControllerV2Response struct {
// 	ID        string `json:"id"`
// 	GUID      string `json:"guid"`
// 	CreatedAt string `json:"created_at"`
// 	UpdatedAt string `json:"updated_at"`
// 	DeletedAt string `json:"deleted_at"`
// 	CreatedBy string `json:"created_by"`
// response struct need to be completed...
// }

// func (f *IBMPIWorkspacesClient) Create(name, location, groudID, plan string) (*ResourceControllerV2, error) {
func (f *IBMPIWorkspacesClient) Create(name, location, groudID, planID string) error {

	endpoint := f.session.Options.URL
	env := ""
	if strings.Contains(endpoint, "test") {
		env = ".test"
	}
	rcUrl := "https://resource-controller" + env + ".cloud.ibm.com/v2/resource_instances"

	resourceController, err := rc.
		NewResourceControllerV2(&rc.ResourceControllerV2Options{
			Authenticator: f.session.Options.Authenticator,
			URL:           rcUrl,
			// ServiceName:   "power-virtual-server-private-group",
		})

	if err != nil {
		return fmt.Errorf("Error creating Resource Controller client: %v", err)
	}

	params := resourceController.NewCreateResourceInstanceOptions(name, location, groudID, planID)
	fmt.Printf("Params \n %v", params)
	result, response, err := resourceController.CreateResourceInstance(params)
	if err != nil {
		return fmt.Errorf("Error creating resource instance: %v", err, result, response)
	}
	if response.StatusCode >= 400 {
		return fmt.Errorf("Error creating resource instance. Status code: %d", response.StatusCode)
	}
	fmt.Printf("Resource Instance created successfully: %v\n", result)

	return nil
}

// Using the api directly

// payload := ResourceControllerV2{name, location, groudID, plan}
// endpoint := f.session.Options.URL
// env := ""
// if strings.Contains(endpoint, "test") {
// 	env = ".test"
// }
// rcUrl := "https://resource-controller" + env + ".cloud.ibm.com/v2/resource_instances"
// bearerToken, err := ibmpisession.FetchAuthorizationData(f.session.Options.Authenticator)
// if err != nil {
// 	fmt.Print(err)
// 	return err
// }
// // headers := map[string]string{
// // 	"Accept":        "application/json",
// // 	"Authorization": auth,
// // }

// fmt.Printf("data before call: %v \n%v \n%v \n%v \n", "payload", endpoint, "headers", url)
// // var createdController ResourceControllerV2
// // var response *resty.Response
// response, err := resty.R().
// 	SetHeaders(headers).
// 	SetBody(&payload).
// 	SetResult(&ResourceControllerV2{}).
// 	Post(url)
// if err != nil {
// 	log.Fatal(err)
// 	return nil, err
// }
// // if response.IsError() {
// // 	log.Fatal(response.Result())
// // 	return nil, err
// // }
// fmt.Print(response.Result().(*ResourceControllerV2))
// return response, nil

// }
