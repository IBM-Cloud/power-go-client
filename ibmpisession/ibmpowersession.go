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
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/client"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/power-go-client/utils"
	//"github.com/IBM-Cloud/bluemix-go/crn"
)

const (
	offering                 = "power-iaas"
	crnString                = "crn"
	version                  = "v1"
	serviceType              = "public"
	serviceInstanceSeparator = "/"
	separator                = ":"
)

// IBMPISession ...
type IBMPISession struct {
	IAMToken    string
	IMSToken    string
	Power       *client.PowerIaas
	Timeout     time.Duration
	UserAccount string
	Region      string
	Zone        string
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

// New ...
/*
The method takes in the following params
iamtoken : this is the token that is passed from the client
region : Obtained from the terraform template. Every template /resource will be required to have this information
timeout:
useraccount:
*/
//func New(bearerToken, region string, debug bool, timeout time.Duration, accountID string, zone string, endpoint string) (*IBMPISession, error) {
func New() (*IBMPISession, error) {
	// get variables
	powerEndpoint, bearerToken, accountID, region, zone, debug, timeout, err := getEnvVariables()
	if err != nil {
		return nil, err
	}

	// create session
	session := &IBMPISession{
		IAMToken:    bearerToken,
		UserAccount: accountID,
		Region:      region,
		Zone:        zone,
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	transport := httptransport.New(powerEndpoint, "/", []string{"https"})
	transport.Debug = debug
	transport.Consumers[runtime.JSONMime] = powerJSONConsumer()
	session.Power = client.New(transport, nil)
	session.Timeout = time.Duration(timeout)
	return session, nil
}

// NewAuth ...
func NewAuth(sess *IBMPISession, PowerInstanceID string) runtime.ClientAuthInfoWriter {
	var crndata = crnBuilder(PowerInstanceID, sess.UserAccount, sess.Region, sess.Zone)
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		if err := r.SetHeaderParam("Authorization", sess.IAMToken); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", crndata)
	})

}

// BearerTokenAndCRN ...
func BearerTokenAndCRN(session *IBMPISession, crn string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		if err := r.SetHeaderParam("Authorization", session.IAMToken); err != nil {
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

func getEnvVariables() (string, string, string, string, string, bool, time.Duration, error) {

	// session variables
	powerEndpoint := ""
	bearerToken := ""
	accountID := ""
	region := ""
	zone := ""
	// default values
	debug := true
	timeout := 9000000000000000000

	// error vars
	err := ""
	envFile := ""

	// get environment
	if env := os.Getenv("ENVIRONMENT"); env == "" {
		envFile = ".env." + env
	}

	// required variables
	if bearerToken = os.Getenv("BEARER_TOKEN"); bearerToken == "" {
		err += fmt.Sprintf("BEARER_TOKEN is empty: define in %v\n", envFile)
	}
	if accountID = os.Getenv("ACCOUNT_ID"); accountID == "" {
		err += fmt.Sprintf("ACCOUNT_ID is empty: define in %v\n", envFile)
	}
	if region = os.Getenv("REGION"); region == "" {
		err += fmt.Sprintf("REGION is empty: define in %v\n", envFile)
	}
	if zone = os.Getenv("ZONE"); zone == "" {
		err += fmt.Sprintf("ZONE is empty: define in %v\n", envFile)
	}
	if powerEndpoint = os.Getenv("ENDPOINT"); powerEndpoint == "" {
		err += fmt.Sprintf("ENDPOINT is empty: define in %v\n", envFile)
	}
	powerEndpoint = region + powerEndpoint

	// optional variables
	if debugStr := os.Getenv("DEBUG"); debugStr != "" {
		debug, _ = strconv.ParseBool(debugStr)
	}
	if timeoutStr := os.Getenv("TIMEOUT"); timeoutStr == "" {
		timeout, _ = strconv.Atoi(timeoutStr)
	}

	if err != "" {
		return "", "", "", "", "", false, 0, fmt.Errorf(err)
	}
	return powerEndpoint, bearerToken, accountID, region, zone, debug, time.Duration(timeout), nil
}
