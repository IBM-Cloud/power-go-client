package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/require"
)

func TestSSHKeyCreateGetDelete(t *testing.T) {
	testSSHKeyPreCheck(t)

	// create session and client
	session := getSession(t)
	sshKeyClient := client.NewIBMPIKeyClient(context.Background(), session, cloudInstanceID)

	// CREATE
	body := &models.SSHKey{
		Name:   &sshKeyName,
		SSHKey: &sshKeyRSA,
	}
	createResp, err := sshKeyClient.Create(body)
	require.Nil(t, err)
	testMessage("CREATE SSH Key", sshKeyName, *createResp)

	// GET
	getResp, err := sshKeyClient.Get(sshKeyName)
	require.Nil(t, err)
	testMessage("GET SSH Key", sshKeyName, *getResp)
	// verify variables match
	require.Equal(t, *getResp.Name, sshKeyName)
	require.Equal(t, *getResp.SSHKey, sshKeyRSA)

	// DELETE
	err = sshKeyClient.Delete(sshKeyName)
	require.Nil(t, err)
	testMessage("DELETE SSH Key", sshKeyName, nil)
}
func TestSSHKeyGetAll(t *testing.T) {
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	sshKeyClient := client.NewIBMPIKeyClient(context.Background(), session, cloudInstanceID)

	// GET ALL
	getAllResp, err := sshKeyClient.GetAll()
	require.Nil(t, err)
	testMessage("GET All SSH Keys", "", *getAllResp)
}
