/*
Package test implements unit tests for the Power-Go-Client:

Files
	- test_utils.go: 	 	Defines commonly used test functions
	- test_variables.go:	Declares all test variables
	- .vscode/launch.json: 	Defines test configuration
	- *_test.go				Defines tests for an individual resource type
	- .vscode/.env:
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

// =====================================
//   Cloud / Account Variables
// =====================================
// converts crn region to (region, zone)
var REGION_ZONE_MAP = map[string][]string{
	"dal12":   {"dal", "12"},
	"dal13":   {"dal", "13"},
	"eu-de-1": {"fra", "04"},
	"eu-de-2": {"fra", "05"},
	"fra04":   {"fra", "04"},
	"fra05":   {"fra", "05"},
	"lon04":   {"lon", "4"},
	"lon06":   {"lon", "6"},
	"mon01":   {"mon", "1"},
	"syd05":   {"syd", "5"},
	"tor01":   {"tor", "1"},
	"osa21":   {"osa", "21"},
	"tok04":   {"tok", "4"},
	"sao01":   {"sao", "1"},
	"syd01":   {"syd", "1"},
	"syd04":   {"syd", "4"},
	"us-east": {"wdc", "04"},
	//"us-east": {"wdc", "wdc06"},
	"wdc04": {"wdc", "04"},
	"wdc06": {"wdc", "wdc06"},
}

var accountID string       // ACCOUNT_ID
var apiKey string          // API_KEY
var bearerToken string     // BEARER_TOKEN (deprecated)
var crn string             // CRN
var cloudInstanceID string // CLOUD_INSTANCE_ID
var powerEndpoint string   // IBMCLOUD_POWER_API_ENDPOINT
var region string          // REGION
var zone string            // ZONE

// =====================================
//   Testing Variables (default values)
// =====================================
var debug bool = false                // DEBUG
var timeout int = 9000000000000000000 // TIMEOUT
var jobTimeout int = 2000             // JOB_TIMEOUT

const TEST_MESSAGE_HEADER_LENGTH = 100
const infoError = "[INFO] %s is not set. Default value: %s.\n"
const warningError = "[WARN] Set the environment variable %s, otherwise %s tests will not run.\n"

// =====================================
//   Resource Variables
// =====================================
// Cloud Connection
var cloudConnectionID string   // CLOUD_CONNECTION_ID
var cloudConnectionName string // CLOUD_CONNECTION_NAME
var cloudConnectionSpeed int64 // CLOUD_CONNECTION_SPEED

// Cloud Instance
//  - currently no tests

// DHCP
//  - current tests dont use any env variables

// Image
var imageID string     // IMAGE_ID
var imageName string   // IMAGE_NAME
var imageSource string // IMAGE_SOURCE

// Instance
var instanceID string                       // INSTANCE_ID
var instanceName string                     // INSTANCE_NAME
var instanceImageID string                  // INSTANCE_IMAGE_ID
var instanceLicenseRepositoryCapacity int64 // INSTANCE_LRC (need to make VTL test)
var instanceMemory float64                  // INSTANCE_MEMORY
var instanceNetworkID string                // INSTANCE_NETWORK_ID
var instanceProcessors float64              // INSTANCE_PROCESSORS
var instanceProcType string                 // INSTANCE_PROC_TYPE
var instanceSSHKey string                   // INSTANCE_SSH_KEY
var instanceSysType string                  // INSTANCE_SYS_TYPE
var instanceVolumeID string                 // INSTANCE_VOLUME_ID

// Job
//  - currently no tests

// Network
var networkID string        // NETWORK_ID
var networkName string      // NETWORK_NAME
var networkCidr string      // NETWORK_CIDR
var networkDnsServer string // NETWORK_DNS_SERVER
var networkJumbo bool       // NETWORK_JUMBO
var networkType string      // NETWORK_TYPE

// Placement Group
//  - currently no tests

// SAP Profile
//  - currently no tests

// Snapshot
//  - currently no tests

// SSH Key
var sshKeyName string // SSH_KEY_NAME
var sshKeyRSA string  // SSH_KEY_RSA

// Storage Capacity
var storagePool string // STORAGE_POOL
var storageType string // STORAGE_TYPE

// System Pool
//  - currently no tests

// Task
//  - currently no tests

// Tenant
//  - currently no tests

// Volume
var volumeID string      // VOLUME_ID
var volumeName string    // VOLUME_NAME
var volumeShareable bool // VOLUME_SHAREABLE
var volumeSize float64   // VOLUME_SIZE
var volumeType string    // VOLUME_TYPE

// VPN
var vpnID string                 // VPN_ID
var vpnName string               // VPN_NAME
var vpnMode string               // VPN_MODE
var vpnNetworkID string          // VPN_NETWORK_ID
var vpnPeerGatewayAddress string // VPN_PEER_GATEWAY_ADDRESS
var vpnPeerSubnet string         // VPN_PEER_SUBNET

// VPN IKE Policy
var ikePolicyID string           // IKE_POLICY_ID
var ikePolicyName string         // IKE_POLICY_NAME
var ikePolicyDhGroup int64       // IKE_POLICY_DH_GROUP
var ikePolicyEncryption string   // IKE_POLICY_ENCRYPTION
var ikePolicyKeyLifetime int     // IKE_POLICY_KEY_LIFETIME
var ikePolicyPresharedKey string // IKE_POLICY_PRESHARED_KEY
var ikePolicyVersion int64       // IKE_POLICY_VERSION

// VPN IPSEC Policy
var ipSecPolicyID string             // IPSEC_POLICY_ID
var ipSecPolicyName string           // IPSEC_POLICY_NAME
var ipSecPolicyAuthentication string // IPSEC_POLICY_AUTHENTICATION
var ipSecPolicyDhGroup int64         // IPSEC_POLICY_DH_GROUP
var ipSecPolicyEncryption string     // IPSEC_POLICY_ENCRYPTION
var ipSecPolicyKeyLifetime int       // IPSEC_POLICY_KEY_LIFETIME
var ipSecPolicyPfs bool              // IPSEC_POLICY_PFS
