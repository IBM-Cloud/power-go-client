package ibmpisession

import (
	"fmt"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/client"
	"github.com/IBM/go-sdk-core/v5/core"
)

// IBMPISession ...
type IBMPISession struct {
	CRNFormat string
	Power     *client.PowerIaasAPI
	Options   *IBMPIOptions
}

// PIOptions
type IBMPIOptions struct {
	// The authenticator implementation to be used by the
	// service instance to authenticate outbound requests
	// Required
	Authenticator core.Authenticator

	// Enable/Disable http transport debugging log
	Debug bool

	// Region of the Power Cloud Service Instance
	// For generating the default host
	// Required when URL is not provided else ignored
	Region string

	// Power Virtual Server host or URL
	// This will be used instead of generating the default host
	// eg: dal.power-iaas.cloud.ibm.com
	// Required when Region is not provided
	URL string

	// Account id of the Power Cloud Service Instance
	// It will be part of the CRN string
	// Required
	UserAccount string

	// Zone of the Power Cloud Service Instance
	// It will be part of the CRN string
	// Required
	Zone string
}

// Create a IBMPISession
func NewIBMPISession(o *IBMPIOptions) (*IBMPISession, error) {
	if core.IsNil(o) {
		return nil, fmt.Errorf("options is required")
	}

	if core.IsNil(o.Authenticator) {
		return nil, fmt.Errorf("option Authenticator is required")
	}

	if o.UserAccount == "" {
		return nil, fmt.Errorf("option UserAccount is required")
	}

	if o.Zone == "" {
		return nil, fmt.Errorf("option Zone is required")
	}

	if o.URL == "" && o.Region == "" {
		return nil, fmt.Errorf("atleast one of Region or URL is required")
	}

	serviceURL := o.URL
	if serviceURL == "" {
		serviceURL = helpers.GetPowerEndPoint(o.Region)
	}

	// We need just the server host from the URL
	serviceURL = strings.TrimPrefix(serviceURL, "https://")
	serviceURL = strings.TrimPrefix(serviceURL, "http://")

	return &IBMPISession{
		CRNFormat: crnBuilder(o.UserAccount, o.Zone, serviceURL),
		Options:   o,
		Power:     getPIClient(o.Debug, serviceURL),
	}, nil
}

// authInfo ...
func (s *IBMPISession) AuthInfo(cloudInstanceID string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		auth, err := fetchAuthorizationData(s.Options.Authenticator)
		if err != nil {
			return err
		}
		if err := r.SetHeaderParam("Authorization", auth); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", fmt.Sprintf(s.CRNFormat, cloudInstanceID))
	})
}
