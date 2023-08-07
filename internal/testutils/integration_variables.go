/*
Package testutils implements tests for the Power-Go-Client:

Files
  - integration_utils.go:			Defines commonly used test functions
  - integration_variables.go:		Lists default values for all test variables
  - example.env:					An example .env which is used when overriding default test variables
  - launch.json:					An example test configuration
  - client/instance/*_test.go:	Defines tests for an individual resource type

To run a test:
  - Set DisableTesting = False in this file
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

import "github.com/IBM-Cloud/power-go-client/power/models"

// set to "true" / "false" to run or not run testing.
// set to false when pushing changes.
var DisableTesting bool = true

// =====================================
//
//	Cloud / Account Variables
//
// =====================================
var apiKey string = ""      // API_KEY
var bearerToken string = "" // BEARER_TOKEN (deprecated)
// CRN
var crn string = "crn:v1:staging:public:power-iaas:dal12:a/efe5e8b9d3f04b948790fe5499bd18bc:6021a723-bcab-4d3f-9985-d0cb2f864f35::"

// derived from crn, can be left empty
var accountID string = ""
var apiURL string = ""
var CloudInstanceID string = ""
var powerEndpoint string = ""
var zone string = ""

// =====================================
//
//	Testing Variables
//
// =====================================
var debug bool = false                    // DEBUG
var timeout int = 9000000000000000000     // TIMEOUT
var TimeoutAttempts int = 500000000000000 // TIMEOUT_ATTEMPTS
var JobTimeout int = 2000                 // JOB_TIMEOUT

const TEST_MESSAGE_HEADER_LENGTH = 100
const WARNING_ERROR = "[WARN] ENV variable %s is not set. Default value: %s = %s will be used. Required for %s tests\n"
const NIL_ERROR = "[ERROR] %s must be set for %s integration tests"
const JOB_ERROR = "Cannot find job status for job id %s with cloud instance %s\n"

// =====================================
//
//	Resource Variables
//
// =====================================
// Cloud Connection
var CloudConnectionID string = ""                            // CLOUD_CONNECTION_ID
var CloudConnectionName string = "Power-Go-Integration-Test" // CLOUD_CONNECTION_NAME
var CloudConnectionSpeed int64 = 50                          // CLOUD_CONNECTION_SPEED

// Image
var ImageID string = "7fb4ac53-0928-4218-af9c-367e33d7c50e" // IMAGE_ID
var ImageName string = "Power-Go-Integration-Test"          // IMAGE_NAME
var ImageSource string = ""                                 // IMAGE_SOURCE

// Instance
var InstanceID string = ""                                            // INSTANCE_ID
var InstanceName string = "Power-Go-Integration-Test"                 // INSTANCE_NAME
var InstanceImageID string = "5a18b1a4-47f6-4fc5-a921-5da487fc5c71"   // INSTANCE_IMAGE_ID
var InstanceLicenseRepositoryCapacity int64 = 0                       // INSTANCE_LRC (need to make VTL test)
var InstanceMemory float64 = 2                                        // INSTANCE_MEMORY
var InstanceNetworkID string = "074b8b59-cad7-4978-83e8-8490cc558e5b" // INSTANCE_NETWORK_ID
var InstanceProcessors float64 = 4                                    // INSTANCE_PROCESSORS
var InstanceProcType string = "shared"                                // INSTANCE_PROC_TYPE
var InstanceSSHKey string = "POWER-GO-INTEGRATION-TEST"               // INSTANCE_SSH_KEY
var InstanceSysType string = "e980"                                   // INSTANCE_SYS_TYPE
var InstanceStorageType string = "tier1"                              // INSTANCE_STORAGE_TYPE
var InstanceVolumeID string = ""                                      // INSTANCE_VOLUME_ID

// Network
var NetworkID string = ""                                     // NETWORK_ID
var NetworkName string = "Power-Go-Integration-Test"          // NETWORK_NAME
var NetworkCidr string = "192.168.0.0/24"                     // NETWORK_CIDR
var NetworkDNSServer string = "127.0.0.1"                     // NETWORK_DNS_SERVER
var NetworkJumbo bool = true                                  // NETWORK_JUMBO
var NetworkMtu int64 = 1450                                   // NETWORK_MTU
var NetworkAccessConfig models.AccessConfig = "internal-only" // NETWORK_ACCESS_CONFIG
var NetworkType string = "vlan"                               // NETWORK_TYPE

// Snapshot
var SnapshotName string = "Power-Go-Integration-Test"                  // SNAPSHOT_NAME
var SnapshotDescription string = "power go integration test"           //SNAPSHOT_DESCRIPTION
var SnapshotInstanceID string = "c24e784c-3d4f-4344-9681-e6c6c1ddc181" // SNAPSHOT_INSTANCE_ID

// SSH Key
var SSHKeyName string = "Power-Go-Integration-Test" // SSH_KEY_NAME
// SSH_KEY_RSA
var SSHKeyRSA string = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCmW5b+shyeSVQrFFshL8YPMGxZdZZT+Ukvp1fMgQCTUS0DSH3GhoJeSvcZSjvDdBDKfy98TFQf3JcZyVLQsiEtymtCPmkZV5kzZRfOdG/Ush034S+RaCXZ7pINoy5VDmYRtgBk8h2urB/jAM/wMR/83DCfpY0BLvvNDJ8jQoBmIaZM5KmdpBqRd9bhmhavWZy/P4KJZctn/45cTkAtW1jd9+oKdIzpoNykcsWXChY8C6SzUXcYBgXluOwnY1fPPdZUyPUghoGqDbHi95E8ee07fap7iS4E+yCAMbW5zIVpOo8SLbLxqWF6JD7pc8oMtgqrxsT01O+3S65QSYZazR1FG7lgyTOGBCuvPxkTbtBk8ytmYuRnFs2XZOR2vD9BSpc7XgDt79SLZS3oIQVwqJBpVZWIJIDbnc8yHRY9X6yDgojBAm5S9KC4vZKMmQ9LBUUOy85XOueDeICGJdt3POefffWNv0kQ/Fscsw/yJS82lOdWo8eHmlLb1ybpr1pp+Is= chase@Chases-MBP.gha.chartermi.net"

// Storage Capacity
var StoragePool string = "Tier1-Flash-2" // STORAGE_POOL
var StorageType string = "tier1"         // STORAGE_TYPE

// Volume
var VolumeID string = ""                            // VOLUME_ID
var VolumeName string = "Power-Go-Integration-Test" // VOLUME_NAME
var VolumeShareable bool = true                     // VOLUME_SHAREABLE
var VolumeSize float64 = 20                         // VOLUME_SIZE
var VolumeType string = "tier1"                     // VOLUME_TYPE

// VPN
var VpnID string = ""                                            // VPN_ID
var VpnName string = "Power-Go-Integration-Test"                 // VPN_NAME
var VpnMode string = "policy"                                    // VPN_MODE
var VpnNetworkID string = "84ff660d-1d21-40ec-b862-d44b535aea72" // VPN_NETWORK_ID
var VpnPeerGatewayAddress string = "1.22.124.1"                  // VPN_PEER_GATEWAY_ADDRESS
var VpnPeerSubnet string = "107.0.0.0/24"                        // VPN_PEER_SUBNET

// VPN IKE Policy
var IKEPolicyID string = ""                            // IKE_POLICY_ID
var IKEPolicyName string = "Power-Go-Integration-Test" // IKE_POLICY_NAME
var IKEPolicyDhGroup int64 = 1                         // IKE_POLICY_DH_GROUP
var IKEPolicyEncryption string = "aes-256-cbc"         // IKE_POLICY_ENCRYPTION
var IKEPolicyKeyLifetime int = 180                     // IKE_POLICY_KEY_LIFETIME
var IKEPolicyPresharedKey string = "test"              // IKE_POLICY_PRESHARED_KEY
var IKEPolicyVersion int64 = 1                         // IKE_POLICY_VERSION

// VPN IPSEC Policy
var IPSecPolicyID string = ""                             // IPSEC_POLICY_ID
var IPSecPolicyName string = "Power-Go-Integration-Test"  // IPSEC_POLICY_NAME
var IPSecPolicyAuthentication string = "hmac-sha-256-128" // IPSEC_POLICY_AUTHENTICATION
var IPSecPolicyDhGroup int64 = 1                          // IPSEC_POLICY_DH_GROUP
var IPSecPolicyEncryption string = "aes-256-cbc"          // IPSEC_POLICY_ENCRYPTION
var IPSecPolicyKeyLifetime int = 180                      // IPSEC_POLICY_KEY_LIFETIME
var IPSecPolicyPfs bool = true                            // IPSEC_POLICY_PFS
