/*
Package test implements unit tests for the Power-Go-Client:

Files
	- test_utils.go: 	 	Defines commonly used test functions
	- test_variables.go:	Declares all test variables
	- .example.launch.json: 	Defines test configuration
	- *_test.go				Defines tests for an individual resource type
	- example.env:
		- Defines all test variables.
		- Different .env files (staging.env, production.env, dal12.env, instance.env,  etc.) can be used
		  depending on the user's test configuration

To run a test:
	- Define environment variables in .env
	- Double click the test function name to select the text
		- Click "Run and Debug" on the VScode sidebar, and then click "Run a test"
		- Output will be visible in the VScode Debug Console
    - launch.json runs tests by using this selected text
*/
package test

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	core "github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/require"
)

// init(): Loads all environment variables
func init() {
	// =====================================
	//   Cloud / Account Variables
	// =====================================
	// First look for API_KEY. If no API_KEY, try BEARER_TOKEN
	if apiKey = os.Getenv("API_KEY"); apiKey == "" {
		if bearerToken = os.Getenv("BEARER_TOKEN"); bearerToken == "" {
			fmt.Printf(warningError, "API_KEY or BEARER_TOKEN (deprecated)", "all")
		}
	}

	// if CRN is set:
	//    set accountID, cloudInstanceID, region, zone, and powerEndpoint based on crn
	// else:
	//    set accountID, cloudInstanceID, region, zone, and powerEndpoint from env variables
	if crn = os.Getenv("CRN"); crn != "" {
		crnParts := strings.Split(crn, ":")
		accountID = crnParts[6][2:]
		cloudInstanceID = crnParts[7]
		region = REGION_ZONE_MAP[crnParts[5]][0]
		zone = REGION_ZONE_MAP[crnParts[5]][1]
		if crnParts[2] == "staging" {
			powerEndpoint = region + ".power-iaas.test.cloud.ibm.com"
			apiURL = "https://iam.test.cloud.ibm.com"
		} else {
			powerEndpoint = region + ".power-iaas.cloud.ibm.com"
			apiURL = "https://iam.cloud.ibm.com"
		}
		// set powerEndpoint for IBMPISession to use
		os.Setenv("IBMCLOUD_POWER_API_ENDPOINT", powerEndpoint)
	} else {
		emptyEnvVars := false

		if accountID = os.Getenv("ACCOUNT_ID"); accountID == "" {
			emptyEnvVars = true
		}
		if cloudInstanceID = os.Getenv("CLOUD_INSTANCE_ID"); cloudInstanceID == "" {
			emptyEnvVars = true
		}
		if region = os.Getenv("REGION"); region == "" {
			emptyEnvVars = true
		}
		if zone = os.Getenv("ZONE"); zone == "" {
			emptyEnvVars = true
		}
		if powerEndpoint = os.Getenv("IBMCLOUD_POWER_API_ENDPOINT"); powerEndpoint == "" {
			emptyEnvVars = true
		} else if powerEndpoint[0] == '.' {
			powerEndpoint = region + powerEndpoint
		}

		if emptyEnvVars {
			vars := "CRN or (ACCOUNT_ID and CLOUD_INSTANCE_ID and REGION and ZONE and IBMCLOUD_POWER_API_ENDPOINT)"
			fmt.Printf(warningError, vars, "all")
		}
	}

	// =====================================
	//   Testing Variables (optional)
	// =====================================
	if debugStr := os.Getenv("DEBUG"); debugStr != "" {
		debug, _ = strconv.ParseBool(debugStr)
	} else {
		fmt.Printf(infoError, "DEBUG", strconv.FormatBool(debug))
	}
	if timeoutStr := os.Getenv("TIMEOUT"); timeoutStr != "" {
		timeout, _ = strconv.Atoi(timeoutStr)
	} else {
		fmt.Printf(infoError, "TIMEOUT", strconv.Itoa(timeout))
	}
	if jobTimeoutStr := os.Getenv("JOB_TIMEOUT"); jobTimeoutStr != "" {
		jobTimeout, _ = strconv.Atoi(jobTimeoutStr)
	} else {
		fmt.Printf(infoError, "JOB_TIMEOUT", strconv.Itoa(jobTimeout))
	}
	if timeoutAttemptsStr := os.Getenv("TIMEOUT_ATTEMPTS"); timeoutAttemptsStr != "" {
		timeoutAttempts, _ = strconv.Atoi(timeoutAttemptsStr)
	} else {
		fmt.Printf(infoError, "TIMEOUT_ATTEMPTS", strconv.Itoa(jobTimeout))
	}

	// =====================================
	//   Resource Variables
	//      - variables used when /tests was /examples
	// =====================================
	// Cloud Connection
	if cloudConnectionName = os.Getenv("CLOUD_CONNECTION_NAME"); cloudConnectionName == "" {
		fmt.Printf(warningError, "CLOUD_CONNECTION_NAME", "cloud connection")
	}
	if cloudConnectionSpeedStr := os.Getenv("CLOUD_CONNECTION_SPEED"); cloudConnectionSpeedStr != "" {
		cloudConnectionSpeed, _ = strconv.ParseInt(cloudConnectionSpeedStr, 10, 64)
	} else {
		fmt.Printf(warningError, "CLOUD_CONNECTION_SPEED", "cloud connection")
	}

	// DHCP
	//  - current tests dont use any env variables

	// Image
	if imageID = os.Getenv("IMAGE_ID"); imageID == "" {
		fmt.Printf(warningError, "IMAGE_ID", "image")
	}
	if imageName = os.Getenv("IMAGE_NAME"); imageName == "" {
		fmt.Printf(warningError, "IMAGE_NAME", "image")
	}
	imageSource = os.Getenv("IMAGE_SOURCE") // can be ""

	// Instance
	if instanceName = os.Getenv("INSTANCE_NAME"); instanceName == "" {
		fmt.Printf(warningError, "INSTANCE_NAME", "instance")
	}
	if instanceImageID = os.Getenv("INSTANCE_IMAGE_ID"); instanceImageID == "" {
		fmt.Printf(warningError, "INSTANCE_IMAGE_ID", "instance")
	}
	if instanceMemoryStr := os.Getenv("INSTANCE_MEMORY"); instanceMemoryStr != "" {
		instanceMemory, _ = strconv.ParseFloat(instanceMemoryStr, 64)
	} else {
		fmt.Printf(warningError, "INSTANCE_MEMORY", "instance")
	}
	if instanceNetworkID = os.Getenv("INSTANCE_NETWORK_ID"); instanceNetworkID == "" {
		fmt.Printf(warningError, "INSTANCE_NETWORK_ID", "instance")
	}
	if instanceProcessorsStr := os.Getenv("INSTANCE_PROCESSORS"); instanceProcessorsStr != "" {
		instanceProcessors, _ = strconv.ParseFloat(instanceProcessorsStr, 64)
	} else {
		fmt.Printf(warningError, "INSTANCE_PROCESSORS", "instance")
	}
	if instanceProcType = os.Getenv("INSTANCE_PROC_TYPE"); instanceProcType == "" {
		fmt.Printf(warningError, "INSTANCE_PROC_TYPE", "instance")
	}
	if instanceSSHKey = os.Getenv("INSTANCE_SSH_KEY"); instanceSSHKey == "" {
		fmt.Printf(warningError, "INSTANCE_SSH_KEY", "instance")
	}
	if instanceSysType = os.Getenv("INSTANCE_SYS_TYPE"); instanceSysType == "" {
		fmt.Printf(warningError, "INSTANCE_SYS_TYPE", "instance")
	}
	if instanceVolumeID = os.Getenv("INSTANCE_VOLUME_ID"); instanceVolumeID == "" {
		fmt.Printf(warningError, "INSTANCE_VOLUME_ID", "instance")
	}

	// Network
	if networkName = os.Getenv("NETWORK_NAME"); networkName == "" {
		fmt.Printf(warningError, "NETWORK_NAME", "network")
	}
	if networkCidr = os.Getenv("NETWORK_CIDR"); networkCidr == "" {
		fmt.Printf(warningError, "NETWORK_CIDR", "network")
	}
	if networkDnsServer = os.Getenv("NETWORK_DNS_SERVER"); networkDnsServer == "" {
		fmt.Printf(warningError, "NETWORK_DNS_SERVER", "network")
	}
	if networkJumboStr := os.Getenv("NETWORK_JUMBO"); networkJumboStr != "" {
		networkJumbo, _ = strconv.ParseBool(networkJumboStr)
	} else {
		fmt.Printf(warningError, "NETWORK_JUMBO", "network")
	}
	if networkType = os.Getenv("NETWORK_TYPE"); networkType == "" {
		fmt.Printf(warningError, "NETWORK_TYPE", "network")
	}

	// Snapshot
	if snapshotName = os.Getenv("SNAPSHOT_NAME"); snapshotName == "" {
		fmt.Printf(warningError, "SNAPSHOT_NAME", "snapshot")
	}
	if snapshotDescription = os.Getenv("SNAPSHOT_DESCRIPTION"); snapshotDescription == "" {
		fmt.Printf(warningError, "SNAPSHOT_DESCRIPTION", "snapshot")
	}
	if snapshotInstanceID = os.Getenv("SNAPSHOT_INSTAANCE_ID"); snapshotInstanceID == "" {
		fmt.Printf(warningError, "SNAPSHOT_INSTAANCE_ID", "snapshot")
	}

	// SSH Key
	if sshKeyName = os.Getenv("SSH_KEY_NAME"); sshKeyName == "" {
		fmt.Printf(warningError, "SSH_KEY_NAME", "ssh key")
	}
	if sshKeyRSA = os.Getenv("SSH_KEY_RSA"); sshKeyRSA == "" {
		fmt.Printf(warningError, "SSH_KEY_RSA", "ssh key")
	}

	// Storage Capacity
	if storagePool = os.Getenv("STORAGE_POOL"); storagePool == "" {
		fmt.Printf(warningError, "STORAGE_POOL", "storage capacity")
	}
	if storageType = os.Getenv("STORAGE_TYPE"); storageType == "" {
		fmt.Printf(warningError, "STORAGE_TYPE", "storage capacity")
	}

	// Volume
	if volumeName = os.Getenv("VOLUME_NAME"); volumeName == "" {
		fmt.Printf(warningError, "VOLUME_NAME", "volume")
	}
	if volumeShareableStr := os.Getenv("VOLUME_SHAREABLE"); volumeShareableStr != "" {
		volumeShareable, _ = strconv.ParseBool(volumeShareableStr)
	} else {
		fmt.Printf(warningError, "VOLUME_SHAREABLE", "volume")
	}
	if volumeSizeStr := os.Getenv("VOLUME_SIZE"); volumeSizeStr != "" {
		volumeSize, _ = strconv.ParseFloat(volumeSizeStr, 64)
	} else {
		fmt.Printf(warningError, "VOLUME_SIZE", "volume")
	}
	if volumeType = os.Getenv("VOLUME_TYPE"); volumeType == "" {
		fmt.Printf(warningError, "VOLUME_TYPE", "volume")
	}

	// VPN
	if vpnName = os.Getenv("VPN_NAME"); vpnName == "" {
		fmt.Printf(warningError, "VPN_NAME", "vpn")
	}
	if vpnMode = os.Getenv("VPN_MODE"); vpnMode == "" {
		fmt.Printf(warningError, "VPN_MODE", "vpn")
	}
	if vpnNetworkID = os.Getenv("VPN_NETWORK_ID"); vpnNetworkID == "" {
		fmt.Printf(warningError, "VPN_NETWORK_ID", "vpn")
	}
	if vpnPeerGatewayAddress = os.Getenv("VPN_PEER_GATEWAY_ADDRESS"); vpnPeerGatewayAddress == "" {
		fmt.Printf(warningError, "VPN_PEER_GATEWAY_ADDRESS", "vpn")
	}
	if vpnPeerSubnet = os.Getenv("VPN_PEER_SUBNET"); vpnPeerSubnet == "" {
		fmt.Printf(warningError, "VPN_PEER_SUBNET", "vpn")
	}

	// VPN IKE Policy
	if ikePolicyName = os.Getenv("IKE_POLICY_NAME"); ikePolicyName == "" {
		fmt.Printf(warningError, "IKE_POLICY_NAME", "vpn")
	}
	if ikePolicyDhGroupStr := os.Getenv("IKE_POLICY_DH_GROUP"); ikePolicyDhGroupStr != "" {
		ikePolicyDhGroup, _ = strconv.ParseInt(ikePolicyDhGroupStr, 10, 64)
	} else {
		fmt.Printf(warningError, "IKE_POLICY_DH_GROUP", "vpn")
	}
	if ikePolicyEncryption = os.Getenv("IKE_POLICY_ENCRYPTION"); ikePolicyEncryption == "" {
		fmt.Printf(warningError, "IKE_POLICY_ENCRYPTION", "vpn")
	}
	if ikePolicyKeyLifetimeStr := os.Getenv("IKE_POLICY_KEY_LIFETIME"); ikePolicyKeyLifetimeStr != "" {
		ikePolicyKeyLifetime, _ = strconv.Atoi(ikePolicyKeyLifetimeStr)
	} else {
		fmt.Printf(warningError, "IKE_POLICY_KEY_LIFETIME", "vpn")
	}
	if ikePolicyPresharedKey = os.Getenv("IKE_POLICY_PRESHARED_KEY"); ikePolicyPresharedKey == "" {
		fmt.Printf(warningError, "IKE_POLICY_PRESHARED_KEY", "vpn")
	}
	if ikePolicyVersionStr := os.Getenv("IKE_POLICY_VERSION"); ikePolicyVersionStr != "" {
		ikePolicyVersion, _ = strconv.ParseInt(ikePolicyVersionStr, 10, 64)
	} else {
		fmt.Printf(warningError, "IKE_POLICY_VERSION", "vpn")
	}

	// VPN IPSec Policy
	if ipSecPolicyName = os.Getenv("IPSEC_POLICY_NAME"); ipSecPolicyName == "" {
		fmt.Printf(warningError, "IPSEC_POLICY_NAME", "vpn")
	}
	if ipSecPolicyAuthentication = os.Getenv("IPSEC_POLICY_AUTHENTICATION"); ipSecPolicyAuthentication == "" {
		fmt.Printf(warningError, "IPSEC_POLICY_AUTHENTICATION", "vpn")
	}
	if ipSecPolicyDhGroupStr := os.Getenv("IPSEC_POLICY_DH_GROUP"); ipSecPolicyDhGroupStr != "" {
		ipSecPolicyDhGroup, _ = strconv.ParseInt(ipSecPolicyDhGroupStr, 10, 64)
	} else {
		fmt.Printf(warningError, "IPSEC_POLICY_DH_GROUP", "vpn")
	}
	if ipSecPolicyEncryption = os.Getenv("IPSEC_POLICY_ENCRYPTION"); ipSecPolicyEncryption == "" {
		fmt.Printf(warningError, "IPSEC_POLICY_ENCRYPTION", "vpn")
	}
	if ipSecPolicyKeyLifetimeStr := os.Getenv("IPSEC_POLICY_KEY_LIFETIME"); ipSecPolicyKeyLifetimeStr != "" {
		ipSecPolicyKeyLifetime, _ = strconv.Atoi(ipSecPolicyKeyLifetimeStr)
	} else {
		fmt.Printf(warningError, "IPSEC_POLICY_KEY_LIFETIME", "vpn")
	}
	if ipSecPolicyPfsStr := os.Getenv("IPSEC_POLICY_PFS"); ipSecPolicyPfsStr != "" {
		ipSecPolicyPfs, _ = strconv.ParseBool(ipSecPolicyPfsStr)
	} else {
		fmt.Printf(warningError, "IPSEC_POLICY_PFS", "vpn")
	}
}

// testPreCheck(): Checks for all required cloud / account variables
func testPreCheck(t *testing.T) {
	// First look for API_KEY. If no API_KEY, try BEARER_TOKEN
	if apiKey == "" {
		if bearerToken == "" {
			t.Fatal("[ERROR] API_KEY or BEARER_TOKEN (deprecated) must be set for all tests")
		}
	}
	if accountID == "" {
		t.Fatal("[ERROR] ACCOUNT_ID must be set for all tests")
	}
	if cloudInstanceID == "" {
		t.Fatal("[ERROR] CLOUD_INSTANCE_ID must be set for all tests")
	}
	if region == "" {
		t.Fatal("[ERROR] REGION must be set for all tests")
	}
	if zone == "" {
		t.Fatal("[ERROR] ZONE must be set for all tests")
	}
	if powerEndpoint == "" {
		t.Fatal("[ERROR] IBMCLOUD_POWER_API_ENDPOINT must be set for all tests")
	}
}

// testCloudConnectionPreCheck(): Checks for Cloud Connection variables
func testCloudConnectionPreCheck(t *testing.T) {
	testPreCheck(t)
	if cloudConnectionName == "" {
		t.Fatal("[ERROR] CLOUD_CONNECTION_NAME must be set for cloud connection tests")
	}
	if speed := os.Getenv("CLOUD_CONNECTION_SPEED"); speed == "" {
		t.Fatal("[ERROR] CLOUD_CONNECTION_SPEED must be set for cloud connection tests")
	}
}

// testDHCPPreCheck(): Checks for DHCP variables
func testDHCPPreCheck(t *testing.T) {
	testPreCheck(t)
}

// testImagePreCheck(): Checks for Image variables
func testImagePreCheck(t *testing.T) {
	testPreCheck(t)
	if imageID == "" {
		t.Fatal("[ERROR] IMAGE_ID must be set for image tests")
	}
	if imageName == "" {
		t.Fatal("[ERROR] IMAGE_NAME must be set for image tests")
	}
}

// testInstancePreCheck(): Checks for Instance variables
func testInstancePreCheck(t *testing.T) {
	testPreCheck(t)
	if instanceName == "" {
		t.Fatal("[ERROR] INSTANCE_NAME must be set for instance tests")
	}
	if instanceImageID == "" {
		t.Fatal("[ERROR] INSTANCE_IMAGE_ID must be set for instance tests")
	}
	if memory := os.Getenv("INSTANCE_MEMORY"); memory == "" {
		t.Fatal("[ERROR] INSTANCE_MEMORY must be set for instance tests")
	}
	if instanceNetworkID == "" {
		t.Fatal("[ERROR] INSTANCE_NETWORK_ID must be set for instance tests")
	}
	if processors := os.Getenv("INSTANCE_PROCESSORS"); processors == "" {
		t.Fatal("[ERROR] INSTANCE_PROCESSORS must be set for instance tests")
	}
	if instanceProcType == "" {
		t.Fatal("[ERROR] INSTANCE_PROC_TYPE must be set for instance tests")
	}
	if instanceSSHKey == "" {
		t.Fatal("[ERROR] INSTANCE_SSH_KEY must be set for instance tests")
	}
	if instanceSysType == "" {
		t.Fatal("[ERROR] INSTANCE_SYS_TYPE must be set for instance tests")
	}
	if instanceVolumeID == "" {
		t.Fatal("[ERROR] INSTANCE_VOLUME_ID must be set for instance tests")
	}
}

// testNetworkPreCheck(): Checks for Network variables
func testNetworkPreCheck(t *testing.T) {
	testPreCheck(t)
	if networkName == "" {
		t.Fatal("[ERROR] NETWORK_NAME must be set for network tests")
	}
	if networkCidr == "" {
		t.Fatal("[ERROR] NETWORK_CIDR must be set for network tests")
	}
	if networkDnsServer == "" {
		t.Fatal("[ERROR] NETWORK_DNS_SERVER must be set for network tests")
	}
	if jumbo := os.Getenv("NETWORK_JUMBO"); jumbo == "" {
		t.Fatal("[ERROR] NETWORK_JUMBO must be set for network tests")
	}
	if networkType == "" {
		t.Fatal("[ERROR] NETWORK_TYPE must be set for network tests")
	}
}

// testSnapshotPreCheck(): Checks for Snapshot Variables
func testSnapshotPreCheck(t *testing.T) {
	testPreCheck(t)
	if snapshotName == "" {
		t.Fatal("[ERROR] SNAPSHOT_NAME must be set for snapshot tests")
	}
	if snapshotDescription == "" {
		t.Fatal("[ERROR] SNAPSHOT_DESCRIPTION must be set for snapshot tests")
	}
	if snapshotInstanceID == "" {
		t.Fatal("[ERROR] SNAPSHOT_INSTANCE_ID must be set for snapshottests")
	}
}

// testSSHKeyPreCheck(): Checks for SSH Key variables
func testSSHKeyPreCheck(t *testing.T) {
	testPreCheck(t)
	if sshKeyName == "" {
		t.Fatal("[ERROR] SSH_KEY_NAME must be set for ssh key tests")
	}
	if sshKeyRSA == "" {
		t.Fatal("[ERROR] SSH_KEY_RSA must be set for ssh key tests")
	}
}

// testStorageCapacityPreCheck(): Checks for Storage Capacity variables
func testStorageCapacityPreCheck(t *testing.T) {
	testPreCheck(t)
	if storagePool == "" {
		t.Fatal("[ERROR] STORAGE_POOL must be set for storage capacity tests")
	}
	if storageType == "" {
		t.Fatal("[ERROR] STORAGE_TYPE must be set for storage capacity tests")
	}
}

// testVolumePreCheck(): Checks for Volume variables
func testVolumePreCheck(t *testing.T) {
	testPreCheck(t)
	if volumeName == "" {
		t.Fatal("[ERROR] VOLUME_NAME must be set for volume tests")
	}
	if shareable := os.Getenv("VOLUME_SHAREABLE"); shareable == "" {
		t.Fatal("[ERROR] VOLUME_SHAREABLE must be set for volume tests")
	}
	if size := os.Getenv("VOLUME_SIZE"); size == "" {
		t.Fatal("[ERROR] VOLUME_SIZE must be set for volume tests")
	}
	if volumeType == "" {
		t.Fatal("[ERROR] VOLUME_TYPE must be set for volume tests")
	}
}

// testVPNPreCheck(): Checks for VPN variables
func testVPNPreCheck(t *testing.T) {
	testPreCheck(t)
	testIKEPolicyPreCheck(t)
	testIPSecPreCheck(t)
	if vpnName == "" {
		t.Fatal("[ERROR] VPN_NAME must be set for vpn tests")
	}
	if vpnMode == "" {
		t.Fatal("[ERROR] VPN_MODE must be set for vpn tests")
	}
	if vpnNetworkID == "" {
		t.Fatal("[ERROR] VPN_NETWORK_ID must be set for vpn tests")
	}
	if vpnPeerGatewayAddress == "" {
		t.Fatal("[ERROR] VPN_PEER_GATEWAY_ADDRESS must be set for vpn tests")
	}
	if vpnPeerSubnet == "" {
		t.Fatal("[ERROR] VPN_PEER_SUBNET must be set for vpn tests")
	}
}

// testIKEPolicyPreCheck(): Checks for VPN IKE Policy variables
func testIKEPolicyPreCheck(t *testing.T) {
	if ikePolicyName == "" {
		t.Fatal("[ERROR] IKE_POLICY_NAME must be set for vpn tests")
	}
	if dhGroup := os.Getenv("IKE_POLICY_DH_GROUP"); dhGroup == "" {
		t.Fatal("[ERROR] IKE_POLICY_DH_GROUP must be set for vpn tests")
	}
	if ikePolicyEncryption == "" {
		t.Fatal("[ERROR] IKE_POLICY_ENCRYPTION must be set for vpn tests")
	}
	if keyLifetime := os.Getenv("IKE_POLICY_KEY_LIFETIME"); keyLifetime == "" {
		t.Fatal("[ERROR] IKE_POLICY_KEY_LIFETIME must be set for vpn tests")
	}
	if ikePolicyPresharedKey == "" {
		t.Fatal("[ERROR] IKE_POLICY_PRESHARED_KEY must be set for vpn tests")
	}
	if policy := os.Getenv("IKE_POLICY_VERSION"); policy == "" {
		t.Fatal("[ERROR] IKE_POLICY_VERSION must be set for vpn tests")
	}
}

// testIPSecPreCheck(): Checks for VPN IPSec Policy variables
func testIPSecPreCheck(t *testing.T) {
	if ipSecPolicyName == "" {
		t.Fatal("[ERROR] IPSEC_POLICY_NAME must be set for vpn tests")
	}
	if ipSecPolicyAuthentication == "" {
		t.Fatal("[ERROR] IPSEC_POLICY_AUTHENTICATION must be set for vpn tests")
	}
	if dhGroup := os.Getenv("IPSEC_POLICY_DH_GROUP"); dhGroup == "" {
		t.Fatal("[ERROR] IPSEC_POLICY_DH_GROUP must be set for vpn tests")
	}
	if ipSecPolicyEncryption == "" {
		t.Fatal("[ERROR] IPSEC_POLICY_ENCRYPTION must be set for vpn tests")
	}
	if keyLifetime := os.Getenv("IPSEC_POLICY_KEY_LIFETIME"); keyLifetime == "" {
		t.Fatal("[ERROR] IPSEC_POLICY_KEY_LIFETIME must be set for vpn tests")
	}
	if pfs := os.Getenv("IPSEC_POLICY_PFS"); pfs == "" {
		t.Fatal("[ERROR] IPSEC_POLICY_PFS must be set for vpn tests")
	}
}

// getSession() Returns *IBMPISession for creating resource clients
func getSession(t *testing.T) *ps.IBMPISession {

	sessionOptions := &ps.IBMPIOptions{
		Debug:       debug,
		UserAccount: accountID,
		Region:      region,
		Zone:        zone,
	}

	// use API Key
	if apiKey != "" {
		authenticator, err := core.NewIamAuthenticatorBuilder().
			SetApiKey(apiKey).
			SetURL(apiURL).
			Build()
		require.Nilf(t, err, "Could not build Authenticator with APIKey: %s", apiKey)
		sessionOptions.Authenticator = authenticator
	} else {
		// use bearer token
		authenticator := &core.BearerTokenAuthenticator{
			BearerToken: bearerToken,
		}
		sessionOptions.Authenticator = authenticator
	}

	session, err := ps.NewIBMPISession(sessionOptions)
	require.Nil(t, err)
	return session
}

func waitForJobState(t *testing.T, jobClient *client.IBMPIJobClient, jobId string) {
	var status string
	for status != "completed" && status != "failed" {
		job, err := jobClient.Get(jobId)
		require.Nil(t, err)
		require.NotNilf(t, job, jobError, jobId, cloudInstanceID)
		require.NotNilf(t, job.Status, jobError, jobId, cloudInstanceID)

		time.Sleep(time.Duration(jobTimeout))
		status = *job.Status.State
	}
}

/*
	testMessage(): Formats test messages
	**********************************[Example: ID]*********************************
	   Response{........}

	msg (Required): 	 Test case title | Required
	id (Optional):  	 ID associated with test case. If none, pass """
	response (Optional): Response associated with test case. If none, pass nil
*/
func testMessage(msg string, id string, response interface{}) {

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

/*
	- image   // imageID = "19215e3-70af-4c44-b048-8945cbc1577c"
	- instance // ask why mem, processors, and proc-type not immediately set
	- vpn // Issues with network
*/
