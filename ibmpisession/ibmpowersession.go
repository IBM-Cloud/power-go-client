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
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/client"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/power-go-client/utils"
	"github.com/IBM/go-sdk-core/v5/core"
	//"github.com/IBM-Cloud/bluemix-go/crn"
)

const (
	offering                 = "power-iaas"
	crnString                = "crn"
	version                  = "v1"
	serviceType              = "public"
	serviceInstanceSeparator = "/"
	separator                = ":"
	//The connection to the server timed out after exceeding timeout duration.
	timeout time.Duration = 9000000000000000000
)

// IBMPISession ...
type IBMPISession struct {
	IAMToken      string
	IMSToken      string
	Power         *client.PowerIaas
	Timeout       time.Duration
	UserAccount   string
	Region        string
	Zone          string
	Authenticator core.Authenticator
}

// PIOptions
/*
Debug : Enable/Disable http transport debugging log, By default false.
Timeout: This is the duration exceeding which the connection awaiting response from the server will be timed out.
UserAccount: This is the accountID to which the targeted Power Cloud Service instance belongs.
Region: This is the region in which the targeted Power Cloud Service Instance resides.
Zone: This is the zone in which the targeted Power Cloud Service Instance resides.
Authenticator: The authenticator used to configure the appropriate type of authentication for use with service instances.
*/
type PIOptions struct {
	Debug         bool
	Timeout       time.Duration
	UserAccount   string
	Region        string
	Zone          string
	Authenticator core.Authenticator
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

/*
The method takes in the following params
options : This is the PIOptions param passed by clients.
This method will create a IBMPISession based on the options passed by client.
It will pass the newly created IBMPISession to newSubProcess method which will evaluate and set rest of the required options from session.
*/
func NewSession(options *PIOptions) (*IBMPISession, error) {
	t := timeout
	if options.Timeout.Seconds() != 0 {
		t = options.Timeout
	}

	newSession := &IBMPISession{
		UserAccount:   options.UserAccount,
		Region:        options.Region,
		Zone:          options.Zone,
		Authenticator: options.Authenticator,
	}

	session, err := process(newSession, options.Debug, t)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// New ...
/*
The method takes in the following params
iamtoken : this is the token that is passed from the client
region : Obtained from the terraform template. Every template /resource will be required to have this information
timeout:
useraccount:
*/
// - deprecated: New function can be used, but is slated to become `obsolete`, Instead try using NewSession function.
func New(iamtoken, region string, debug bool, timeout time.Duration, useraccount string, zone string) (*IBMPISession, error) {
	fmt.Println("Warning - New function is slated to become `obsolete`, Instead try using NewSession function.")
	session := &IBMPISession{
		IAMToken:    iamtoken,
		UserAccount: useraccount,
		Region:      region,
		Zone:        zone,
	}

	session, err := process(session, debug, timeout)
	if err != nil {
		return nil, err
	}
	return session, nil
}

/*
The method takes in the following params
session : This is the session passed via New OR NewSession method, created using params or PIOptions passed by clients.
debug : Enable/Disable http transport debugging log, By default false.
timeout: This is the duration exceeding which the connection awaiting response from the server will be timed out.
The method will assign appropriate PowerIaas client to session.Power option after appropriate evaluation.
*/
func process(session *IBMPISession, debug bool, timeout time.Duration) (*IBMPISession, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	apiEndpointURL := utils.GetPowerEndPoint(session.Region)
	transport := httptransport.New(apiEndpointURL, "/", []string{"https"})
	if debug {
		transport.Debug = debug
	}
	transport.Consumers[runtime.JSONMime] = powerJSONConsumer()
	session.Power = client.New(transport, nil)
	session.Timeout = timeout
	return session, nil
}

/*
Check if authenticator option is set
If yes, authenticate using it and fetch authorization header data required for the request.
*/
func fetchAuthorizationData(s *IBMPISession) (string, error) {
	var authdata = s.IAMToken
	if s.Authenticator != nil {
		req := &http.Request{
			Header: make(http.Header),
		}
		if err := s.Authenticator.Validate(); err != nil {
			return "", err
		}
		if err := s.Authenticator.Authenticate(req); err != nil {
			return "", err
		}
		authdata = req.Header.Get("Authorization")
	}
	return authdata, nil
}

// NewAuth ...
func NewAuth(sess *IBMPISession, PowerInstanceID string) runtime.ClientAuthInfoWriter {
	var crndata = crnBuilder(PowerInstanceID, sess.UserAccount, sess.Region, sess.Zone)
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
func crnBuilder(powerinstance, useraccount, region string, zone string) string {
	var service string
	if strings.Contains(utils.GetPowerEndPoint(region), ".power-iaas.cloud.ibm.com") {
		service = "bluemix"
	} else {
		service = "staging"
	}
	var crnData string
	if zone == "" {
		crnData = crnString + separator + version + separator + service + separator + serviceType + separator + offering + separator + region + separator + "a" + serviceInstanceSeparator + useraccount + separator + powerinstance + separator + separator
	} else {
		crnData = crnString + separator + version + separator + service + separator + serviceType + separator + offering + separator + zone + separator + "a" + serviceInstanceSeparator + useraccount + separator + powerinstance + separator + separator
	}
	return crnData
}
