/*
Package testutils implements tests for the Power-Go-Client:

Files
	- integration_utils.go:			Defines commonly used test functions
	- integration_variables.go:		Lists default values for all test variables
	- example.env:					An example .env which is used when overriding default test variables
	- launch.json:					An example test configuration
	- client/instance/*_test.go:	Defines tests for an individual resource type

To run a test:
	- Set DisableTesting = False in integration_variables.go
	- Create a .vscode folder in the project
	- Copy and rename launch.json and .env into .vscode
	- Either
		- Use default variables in integration_variables.go
		- Define variables in .env
	- Double click the test function name to select the text. launch.json uses this text
	- Click "Run and Debug" on the VScode sidebar, and then click "Run a test"
		- Output will be visible in the VScode Debug Console
*/
package testutils

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	client "github.com/IBM-Cloud/power-go-client/clients"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	core "github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/require"
)

// init loads env variables into testing variables, otherwise default is used.
func init() {

	// Skip init if testing is disabled
	if DisableTesting {
		return
	}

	// Cloud / Account Variables
	loadString(&apiKey, "apiKey", "API_KEY", "all")
	if apiKey != " " {
		loadString(&bearerToken, "bearerToken", "BEARER_TOKEN", "all")
	}
	loadString(&crn, "crn", "CRN", "all")

	// Testing Variables
	loadBool(&debug, "debug", "DEBUG", "all")
	loadInt(&timeout, "timeout", "TIMEOUT", "all")
	loadInt(&JobTimeout, "JobTimeout", "JOB_TIMEOUT", "all")
	loadInt(&TimeoutAttempts, "TimeoutAttempts", "TIMEOUT_ATTEMPTS", "all")

	// =====================================
	//   Resource Variables
	// =====================================

	// Cloud Connection
	loadString(&CloudConnectionName, "CloudConnectionName", "CLOUD_CONNECTION_NAME", "cloud connection")
	loadInt64(&CloudConnectionSpeed, "CloudConnectionSpeed", "CLOUD_CONNECTION_SPEED", "cloud connection")

	// DHCP (current tests dont use any env variables)

	// Image
	loadString(&ImageID, "ImageID", "IMAGE_ID", "image")
	loadString(&ImageName, "ImageName", "IMAGE_NAME", "image")
	loadString(&ImageSource, "ImageSource", "IMAGE_SOURCE", "image")

	// Instance
	loadString(&InstanceName, "InstanceName", "INSTANCE_NAME", "instance")
	loadString(&InstanceImageID, "InstanceImageID", "INSTANCE_IMAGE_ID", "instance")
	loadFloat64(&InstanceMemory, "InstanceMemory", "INSTANCE_MEMORY", "instance")
	loadString(&InstanceNetworkID, "InstanceNetworkID", "INSTANCE_NETWORK_ID", "instance")
	loadFloat64(&InstanceProcessors, "InstanceProcessors", "INSTANCE_PROCESSORS", "instance")
	loadString(&InstanceProcType, "InstanceProcType", "INSTANCE_PROC_TYPE", "instance")
	loadString(&InstanceSSHKey, "InstanceSSHKey", "INSTANCE_SSH_KEY", "instance")
	loadString(&InstanceSysType, "InstanceSysType", "INSTANCE_SYS_TYPE", "instance")
	loadString(&InstanceVolumeID, "InstanceVolumeID", "INSTANCE_VOLUME_ID", "instance")

	// Network
	loadString(&NetworkName, "NetworkName", "NETWORK_NAME", "network")
	loadString(&NetworkCidr, "NetworkCidr", "NETWORK_CIDR", "network")
	loadString(&NetworkDNSServer, "NetworkDNSServer", "NETWORK_DNS_SERVER", "network")
	loadBool(&NetworkJumbo, "NetworkJumbo", "NetworkJumbo", "network")
	loadString(&NetworkType, "NetworkType", "NETWORK_TYPE", "network")

	// Snapshot
	loadString(&SnapshotName, "SnapshotName", "SNAPSHOT_NAME", "snapshot")
	loadString(&SnapshotDescription, "SnapshotDescription", "SNAPSHOT_DESCRIPTION", "snapshot")
	loadString(&SnapshotInstanceID, "SnapshotInstanceID", "SNAPSHOT_INSTANCE_ID", "snapshot")

	// SSH Key
	loadString(&SSHKeyName, "SSHKeyName", "SSH_KEY_NAME", "ssh key")
	loadString(&SSHKeyRSA, "SSHKeyRSA", "SSH_KEY_RSA", "ssh key")

	// Storage Capacity
	loadString(&StoragePool, "StoragePool", "STORAGE_POOL", "storage capacity")
	loadString(&StorageType, "StorageType", "STORAGE_TYPE", "storage capacity")

	// Volume
	loadString(&VolumeName, "VolumeName", "VOLUME_NAME", "volume")
	loadBool(&VolumeShareable, "VolumeShareable", "VOLUME_SHAREABLE", "volume")
	loadFloat64(&VolumeSize, "VolumeSize", "VOLUME_SIZE", "volume")
	loadString(&VolumeType, "VolumeType", "VOLUME_TYPE", "volume")

	// VPN
	loadString(&VpnName, "VpnName", "VPN_NAME", "vpn")
	loadString(&VpnMode, "VpnMode", "VPN_MODE", "vpn")
	loadString(&VpnNetworkID, "VpnNetworkID", "VPN_NETWORK_ID", "vpn")
	loadString(&VpnPeerGatewayAddress, "VpnPeerGatewayAddress", "VPN_PEER_GATEWAY_ADDRESS", "vpn")
	loadString(&VpnPeerSubnet, "VpnPeerSubnet", "VPN_PEER_SUBNET", "vpn")

	// VPN IKE Policy
	loadString(&IKEPolicyName, "IKEPolicyName", "IKE_POLICY_NAME", "vpn")
	loadInt64(&IKEPolicyDhGroup, "IKEPolicyDhGroup", "IKE_POLICY_DH_GROUP", "vpn")
	loadString(&IKEPolicyEncryption, "IKEPolicyEncryption", "IKE_POLICY_ENCRYPTION", "vpn")
	loadInt(&IKEPolicyKeyLifetime, "IKEPolicyKeyLifetime", "IKE_POLICY_KEY_LIFETIME", "vpn")
	loadString(&IKEPolicyPresharedKey, "IKEPolicyPresharedKey", "IKE_POLICY_PRESHARED_KEY", "vpn")
	loadInt64(&IKEPolicyVersion, "IKEPolicyVersion", "IKE_POLICY_VERSION", "vpn")

	// VPN IPSec Policy
	loadString(&IPSecPolicyName, "IPSecPolicyName", "IPSEC_POLICY_NAME", "vpn")
	loadString(&IPSecPolicyAuthentication, "IPSecPolicyAuthentication", "IPSEC_POLICY_AUTHENTICATION", "vpn")
	loadInt64(&IPSecPolicyDhGroup, "IPSecPolicyDhGroup", "IPSEC_POLICY_DH_GROUP", "vpn")
	loadString(&IPSecPolicyEncryption, "IPSecPolicyEncryption", "IPSEC_POLICY_ENCRYPTION", "vpn")
	loadInt(&IPSecPolicyKeyLifetime, "IPSecPolicyKeyLifetime", "IPSEC_POLICY_KEY_LIFETIME", "vpn")
	loadBool(&IPSecPolicyPfs, "IPSEC_POLICY_PFS", "IPSEC_POLICY_PFS", "vpn")
}

// AccountPreCheck verifies that all cloud / account variables are defined.
func AccountPreCheck(t *testing.T) {
	// First look for API_KEY. If no API_KEY, try BEARER_TOKEN
	if &apiKey == nil {
		ifNil(t, bearerToken, "apiKey or bearerToken", "all")
	}
	ifNil(t, accountID, "accountID", "all")
	ifNil(t, CloudInstanceID, "CloudInstanceID", "all")
	ifNil(t, zone, "zone", "all")
	ifNil(t, powerEndpoint, "powerEndpoint", "all")
}

// CloudConnectionPreCheck verifies that all Cloud Connection variables are defined.
func CloudConnectionPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, CloudConnectionName, "CloudConnectionName", "cloud connection")
	ifNil(t, CloudConnectionSpeed, "CloudConnectionSpeed", "cloud connection")
}

// DHCPPreCheck verifies that all DHCP variables are defined.
func DHCPPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
}

// ImagePreCheck verifies that all Image variables are defined.
func ImagePreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, ImageID, "ImageID", "image")
	ifNil(t, ImageName, "ImageName", "image")
	ifNil(t, ImageSource, "ImageSource", "image")
}

// InstancePreCheck verifies that all Instance variables are defined.
func InstancePreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	// not checking for VTL currently
	ifNil(t, InstanceName, "InstanceName", "instance")
	ifNil(t, InstanceImageID, "InstanceImageID", "instance")
	ifNil(t, InstanceMemory, "InstanceMemory", "instance")
	ifNil(t, InstanceNetworkID, "InstanceNetworkID", "instance")
	ifNil(t, InstanceProcessors, "InstanceProcessors", "instance")
	ifNil(t, InstanceProcType, "InstanceProcType", "instance")
	ifNil(t, InstanceSSHKey, "InstanceSSHKey", "instance")
	ifNil(t, InstanceSysType, "InstanceSysType", "instance")
	ifNil(t, InstanceVolumeID, "InstanceVolumeID", "instance")
}

// NetworkPreCheck verifies that all Network variables are defined.
func NetworkPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, NetworkName, "NetworkName", "network")
	ifNil(t, NetworkCidr, "NetworkCidr", "network")
	ifNil(t, NetworkDNSServer, "NetworkDNSServer", "network")
	ifNil(t, NetworkJumbo, "NetworkJumbo", "network")
	ifNil(t, NetworkType, "NetworkType", "network")
}

// SnapshotPreCheck verifies that all Snapshot Variables are defined.
func SnapshotPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, SnapshotName, "SnapshotName", "snapshot")
	ifNil(t, SnapshotDescription, "SnapshotDescription", "snapshot")
	ifNil(t, SnapshotInstanceID, "SnapshotInstanceID", "snapshot")
}

// SSHKeyPreCheck verifies that all SSH Key variables are defined.
func SSHKeyPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, SSHKeyName, "SnapshotInstanceID", "ssh key")
	ifNil(t, SSHKeyRSA, "SnapshotInstanceID", "ssh key")
}

// StorageCapacityPreCheck verifies that all Storage Capacity variables are defined.
func StorageCapacityPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, StoragePool, "StoragePool", "storage capacity")
	ifNil(t, StorageType, "StorageType", "storage capacity")
}

// VolumePreCheck verifies that all Volume variables are defined.
func VolumePreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, VolumeName, "VolumeName", "volume")
	ifNil(t, VolumeShareable, "VolumeShareable", "volume")
	ifNil(t, VolumeSize, "VolumeSize", "volume")
	ifNil(t, VolumeType, "VolumeType", "volume")
}

// VPNPreCheck verifies that all VPN variables are defined.
func VPNPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, VpnName, "VpnName", "vpn")
	ifNil(t, VpnMode, "VpnMode", "vpn")
	ifNil(t, VpnNetworkID, "VpnNetworkID", "vpn")
	ifNil(t, VpnPeerGatewayAddress, "VpnPeerGatewayAddress", "vpn")
	ifNil(t, VpnPeerSubnet, "VpnPeerSubnet", "vpn")
}

// IKEPolicyPreCheck verifies that all VPN IKE Policy variables are defined.
func IKEPolicyPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, IKEPolicyName, "IKEPolicyName", "vpn")
	ifNil(t, IKEPolicyDhGroup, "IKEPolicyDhGroup", "vpn")
	ifNil(t, IKEPolicyEncryption, "IKEPolicyEncryption", "vpn")
	ifNil(t, IKEPolicyKeyLifetime, "IKEPolicyKeyLifetime", "vpn")
	ifNil(t, IKEPolicyPresharedKey, "IKEPolicyPresharedKey", "vpn")
	ifNil(t, IKEPolicyVersion, "IKEPolicyVersion", "vpn")
}

// IPSecPreCheck verifies that all VPN IPSec Policy variables are defined.
func IPSecPreCheck(t *testing.T) {

	// Skip if testing is disabled
	if DisableTesting {
		return
	}

	AccountPreCheck(t)
	ifNil(t, IPSecPolicyName, "IPSecPolicyName", "vpn")
	ifNil(t, IPSecPolicyAuthentication, "IPSecPolicyAuthentication", "vpn")
	ifNil(t, IPSecPolicyDhGroup, "IPSecPolicyDhGroup", "vpn")
	ifNil(t, IPSecPolicyEncryption, "IPSecPolicyEncryption", "vpn")
	ifNil(t, IPSecPolicyKeyLifetime, "IPSecPolicyKeyLifetime", "vpn")
	ifNil(t, IPSecPolicyPfs, "IPSecPolicyPfs", "vpn")
}

// TestSession returns an *IBMPISession for testing.
func TestSession(t *testing.T) *ps.IBMPISession {

	extractCrn()

	// create authenticator
	var auth core.Authenticator
	if apiKey != "" {
		// API Key
		apiKeyAuth, err := core.NewIamAuthenticatorBuilder().
			SetApiKey(apiKey).
			SetURL(apiURL).
			Build()
		require.Nilf(t, err, "Could not build Authenticator with apiKey: %s", apiKey)
		auth = apiKeyAuth
	} else {
		// Bearer Token
		bearerAuth := &core.BearerTokenAuthenticator{BearerToken: bearerToken}
		auth = bearerAuth
	}

	// create session options
	sessionOptions := &ps.IBMPIOptions{
		Debug:         debug,
		UserAccount:   accountID,
		Zone:          zone,
		Authenticator: auth,
	}

	// create session
	session, err := ps.NewIBMPISession(sessionOptions)
	require.Nil(t, err)
	return session
}

// WaitForJobState waits for a job to be completed before returning.
func WaitForJobState(t *testing.T, jobClient *client.IBMPIJobClient, jobId string) {
	var status string
	for status != "completed" && status != "failed" {
		job, err := jobClient.Get(jobId)
		require.Nil(t, err)
		require.NotNilf(t, job, JOB_ERROR, jobId, CloudInstanceID)
		require.NotNilf(t, job.Status, JOB_ERROR, jobId, CloudInstanceID)
		time.Sleep(time.Duration(JobTimeout))
		status = *job.Status.State
	}
}

// extract Crn extracts accountID, CloudInstanceID, powerEndpoint, zone, and apiURL
func extractCrn() {
	crnSlice := strings.Split(crn, ":")
	accountID = crnSlice[6][2:]
	CloudInstanceID = crnSlice[7]
	zone = crnSlice[5]
	if crnSlice[2] == "staging" {
		powerEndpoint = "power-iaas.test.cloud.ibm.com"
		apiURL = "https://iam.test.cloud.ibm.com"
	} else {
		powerEndpoint = "power-iaas.cloud.ibm.com"
		apiURL = "https://iam.cloud.ibm.com"
	}
	// set powerEndpoint for NewIBMPISession to use
	os.Setenv("IBMCLOUD_POWER_API_ENDPOINT", powerEndpoint)
}

// ifNil stops stops its execution if variable is not set.
func ifNil(t *testing.T, variable interface{}, varName, varTest string) {
	require.NotNilf(t, &variable, fmt.Sprintf(NIL_ERROR, varName, varTest))
}

// loadBool loads the env variable verson of testVar if it exists, otherwise the default value is used.
func loadBool(testVar *bool, testVarName, envVar, resource string) {
	if env := os.Getenv(envVar); env != "" {
		*testVar, _ = strconv.ParseBool(env)
	} else {
		boolVal := strconv.FormatBool(*testVar)
		fmt.Printf(WARNING_ERROR, envVar, testVarName, boolVal, resource)
	}
}

// loadFloat64 loads the env variable verson of testVar if it exists, otherwise the default value is used.
func loadFloat64(testVar *float64, testVarName, envVar, resource string) {
	if env := os.Getenv(envVar); env != "" {
		*testVar, _ = strconv.ParseFloat(env, 64)
	} else {
		floatVal := fmt.Sprintf("%.1f", *testVar)
		fmt.Printf(WARNING_ERROR, envVar, testVarName, floatVal, resource)
	}
}

// loadInt loads the env variable verson of testVar if it exists, otherwise the default value is used.
func loadInt(testVar *int, testVarName, envVar, resource string) {
	if env := os.Getenv(envVar); env != "" {
		*testVar, _ = strconv.Atoi(env)
	} else {
		intVal := strconv.Itoa(*testVar)
		fmt.Printf(WARNING_ERROR, envVar, testVarName, intVal, resource)
	}
}

// loadInt64 loads the env variable verson of testVar if it exists, otherwise the default value is used.
func loadInt64(testVar *int64, testVarName, envVar, resource string) {
	if env := os.Getenv(envVar); env != "" {
		*testVar, _ = strconv.ParseInt(env, 10, 64)
	} else {
		intVal := strconv.FormatInt(*testVar, 10)
		fmt.Printf(WARNING_ERROR, envVar, testVarName, intVal, resource)
	}
}

// loadString loads the env variable verson of testVar if it exists, otherwise the default value is used.
func loadString(testVar *string, testVarName, envVar, resource string) {
	if env := os.Getenv(envVar); env != "" {
		*testVar = env
	} else {
		strVal := *testVar
		if strVal == "" {
			strVal = "\"\""
		}
		fmt.Printf(WARNING_ERROR, envVar, testVarName, strVal, resource)
	}
}

// testMessage formats an prints a response.
func TestMessage(msg string, id string, response interface{}) {

	// format text in brackets
	if id != "" {
		msg += ": " + id
	}

	// add * around brackets
	messagePadding := float64(TEST_MESSAGE_HEADER_LENGTH - len(msg))
	leftPadding := int(math.Floor(messagePadding / 2))
	rightPadding := int(math.Ceil(messagePadding / 2))
	left := strings.Repeat("*", leftPadding) + "["
	right := "]" + strings.Repeat("*", rightPadding)
	msg = left + msg + right

	// add optional response
	if response != nil {
		msg += fmt.Sprintf("\n   %+v", response)
	}
	msg += "\n\n"
	log.Printf(msg)
}
