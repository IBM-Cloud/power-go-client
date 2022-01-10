package test

import "testing"

func TestAll(t *testing.T) {
	// Cloud Connection
	TestCloudConnectionCreateGetDelete(t)
	TestCloudConnectionGetAll(t)

	// DHCP
	TestDHCPCreateGetDelete(t)
	TestDHCPGetAll(t)

	// Image
	TestImageCreateGetDelete(t)
	TestImageGetAll(t)

	// Instance
	TestInstanceCreateGetDelete(t)
	TestInstanceGetAll(t)

	// Network
	TestNetworkCreateGetDelete(t)

	// Snapshot
	TestSnapshotCreateGetDelete(t)
	TestSnapshotGetAll(t)

	// SSH
	TestSSHKeyCreateGetDelete(t)
	TestSSHKeyGetAll(t)

	// Storage Capactiy
	TestStorageCapacityGet(t)

	// Volume
	TestVolumeCreateGetDelete(t)
	TestVolumeGetAll(t)

	// VPN
	TestVPNCreateGetDelete(t)
	TestVPNGetAll(t)
}
