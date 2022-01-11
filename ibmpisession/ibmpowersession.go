/*
Code to call the IBM IAM Services and get a session object that will be used by the Power Colo Code.
*/

package ibmpisession

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/client"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"
)

// IBMPISession ...
type IBMPISession struct {
	IAMToken      string
	IMSToken      string
	Power         *client.PowerIaasAPI
	UserAccount   string
	Region        string
	Zone          string
	Authenticator core.Authenticator
}

// PIOptions
type PIOptions struct {
	// The authenticator used to configure the appropriate type of authentication
	// Required
	Authenticator core.Authenticator

	//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
	// Enable/Disable http transport debugging log
	Debug bool

	// Account id of the Power Cloud Service Instance
	// Required
	UserAccount string

	// Region of the Power Cloud Service Instance
	// Required
	Region string

	// Zone of the Power Cloud Service Instance; Use Region if not set
	Zone string
}

func powerJSONConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(reader)
		if err != nil {
			return err
		}
		b := buf.Bytes()
		if b != nil {
			dec := json.NewDecoder(bytes.NewReader(b))
			dec.UseNumber() // preserve number formats
			err = dec.Decode(data)
		}
		if string(b) == "null" || err != nil {
			errorRecord, _ := data.(*models.Error)
			log.Printf("The errorrecord is %s ", errorRecord.Error)
			return nil
		}
		return err
	})
}

// Create a IBMPISession
func NewSession(options *PIOptions) (*IBMPISession, error) {
	newSession := &IBMPISession{
		UserAccount:   options.UserAccount,
		Region:        options.Region,
		Zone:          options.Zone,
		Authenticator: options.Authenticator,
	}
	session, err := process(newSession, options.Debug)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// Create a IBMPISession
// - deprecated: use NewSession
func New(iamtoken, region string, debug bool, useraccount string, zone string) (*IBMPISession, error) {
	session := &IBMPISession{
		IAMToken:    iamtoken,
		UserAccount: useraccount,
		Region:      region,
		Zone:        zone,
	}
	session, err := process(session, debug)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// Assign appropriate PowerIaas client to the session after appropriate evaluation
func process(session *IBMPISession, debug bool) (*IBMPISession, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	apiEndpointURL := GetPowerEndPoint(session.Region)
	transport := httptransport.New(apiEndpointURL, "/", []string{"https"})
	transport.Debug = debug
	transport.SetLogger(IBMPILogger{})
	transport.Consumers[runtime.JSONMime] = powerJSONConsumer()
	session.Power = client.New(transport, nil)
	return session, nil
}

// Fetch Authorization token
// Authenticate using Authenticator if set in the session
// Otherwise return IAMToken set in the session
func fetchAuthorizationData(s *IBMPISession) (string, error) {
	var authdata = s.IAMToken
	if s.Authenticator != nil {
		req := &http.Request{
			Header: make(http.Header),
		}
		if err := s.Authenticator.Authenticate(req); err != nil {
			return "", err
		}
		authdata = req.Header.Get("Authorization")
	}
	return authdata, nil
}

// NewAuth ...
func NewAuth(sess *IBMPISession, cloudInstanceID string) runtime.ClientAuthInfoWriter {
	var crndata = crnBuilder(cloudInstanceID, sess.UserAccount, sess.Region, sess.Zone)
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		authdata, err := fetchAuthorizationData(sess)
		if err != nil {
			return err
		}
		if err := r.SetHeaderParam("Authorization", authdata); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", crndata)
	})
}

// BearerTokenAndCRN ...
func BearerTokenAndCRN(session *IBMPISession, crn string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		authdata, err := fetchAuthorizationData(session)
		if err != nil {
			return err
		}
		if err := r.SetHeaderParam("Authorization", authdata); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", crn)
	})
}

// crnBuilder ...
func crnBuilder(cloudInstanceID, useraccount, region, zone string) string {
	var service string
	if strings.Contains(GetPowerEndPoint(region), ".power-iaas.cloud.ibm.com") {
		service = "bluemix"
	} else {
		service = "staging"
	}
	if zone == "" {
		zone = region
	}
	return fmt.Sprintf("crn:v1:%s:public:power-iaas:%s:a/%s:%s::", service, zone, useraccount, cloudInstanceID)
}

func GetPowerEndPoint(regionName string) string {
	if url := os.Getenv("IBMCLOUD_POWER_API_ENDPOINT"); url != "" {
		if url[0] == '.' {
			return regionName + url
		}
		return url
	}
	return regionName + ".power-iaas.cloud.ibm.com"
}
