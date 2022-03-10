package ibmpisession

import (
	"fmt"
	"regexp"
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
	// For generating the default endpoint
	// Deprecated: Region is deprecated, the URL is auto generated based on Zone when not provided.
	Region string

	// Power Virtual Server host or URL endpoint
	// This will be used instead of generating the default host
	// eg: dal.power-iaas.cloud.ibm.com
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

	var serviceURL string
	if o.URL != "" {
		serviceURL = o.URL
	} else {
		// If URL is not set check in env
		serviceURL = helpers.GetPowerEndPoint()
		if serviceURL == "" {
			// Generate default endpoint with Region
			if o.Region != "" {
				serviceURL = o.Region + ".power-iaas.cloud.ibm.com"
			} else {
				region := costructRegionFromZone(o.Zone)
				serviceURL = region + ".power-iaas.cloud.ibm.com"

			}
		}
	}

	// We need just the server host from the URL
	var host, scheme string
	if strings.HasPrefix(serviceURL, "https://") {
		scheme = SCHEME_HTTPS
		host = strings.TrimPrefix(serviceURL, "https://")
	} else if strings.HasPrefix(serviceURL, "http://") {
		scheme = SCHEME_HTTP
		host = strings.TrimPrefix(serviceURL, "http://")
	} else {
		// by default we use "https"
		scheme = SCHEME_HTTPS
		host = serviceURL
	}

	return &IBMPISession{
		CRNFormat: crnBuilder(o.UserAccount, o.Zone, host),
		Options:   o,
		Power:     getPIClient(o.Debug, host, scheme),
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

func costructRegionFromZone(zone string) string {
	var regex string
	if strings.Contains(zone, "-") {
		// it's a region or AZ
		regex = "-[0-9]+$"
	} else {
		// it's a datacenter
		regex = "[0-9]+$"
	}

	reg, _ := regexp.Compile(regex)
	return reg.ReplaceAllString(zone, "")
}
