package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/routes"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIPlacementGroupClient
type IBMPIRouteClient struct {
	IBMPIClient
}

// NewIBMPIPlacementGroupClient
func NewIBMPIPRouteClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIRouteClient {
	return &IBMPIRouteClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get a routing rule
func (f *IBMPIRouteClient) Get(id string) (*models.Route, error) {
	params := routes.NewV1RoutesGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).WithRouteID(id)
	resp, err := f.session.Power.Routes.V1RoutesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to get route %s: %w", id, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get route %s", id)
	}
	return resp.Payload, nil
}

// Get all routes
func (f *IBMPIRouteClient) GetAll() (*models.Routes, error) {
	params := routes.NewV1RoutesGetallParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.Routes.V1RoutesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to get all routes: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all routes", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get route report
func (f *IBMPIRouteClient) GetRouteReport() (*models.RouteReport, error) {
	params := routes.NewV1RoutesReportGetParams().WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut)
	resp, err := f.session.Power.Routes.V1RoutesReportGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to get route report: %w", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get route report", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a routing rule
func (f *IBMPIRouteClient) Create(body *models.RouteCreate) (*models.Route, error) {
	params := routes.NewV1RoutesPostParams().WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).WithBody(body)
	resp, err := f.session.Power.Routes.V1RoutesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to create route: %s", err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to create route")
	}
	return resp.Payload, nil
}

// Update routing rule
func (f *IBMPIRouteClient) Update(id string, body *models.RouteUpdate) (*models.Route, error) {
	params := routes.NewV1RoutesPutParams().WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).WithRouteID(id)
	resp, err := f.session.Power.Routes.V1RoutesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, ibmpisession.SDKFailWithAPIError(err, fmt.Errorf("failed to update route %s: %w", id, err))
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to update route %s", id)
	}
	return resp.Payload, nil
}

// Delete routing rule
func (f *IBMPIRouteClient) Delete(id string) error {
	params := routes.NewV1RoutesDeleteParams().WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).WithRouteID(id)
	_, err := f.session.Power.Routes.V1RoutesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to delete route %s: %w", id, err)
	}
	return nil
}
