package test

import (
	"context"
	"testing"

	client "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func TestSSHKeyCreateGetDelete(t *testing.T) {
	// check for testing variables
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
	if err != nil {
		t.Fatal(err)
	}
	testMessage("CREATE SSH Key", sshKeyName, *createResp)

	// GET
	getResp, err := sshKeyClient.Get(sshKeyName)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET SSH Key", sshKeyName, *getResp)

	// DELETE
	err = sshKeyClient.Delete(sshKeyName)
	if err != nil {
		t.Fatal(err)
	}
	testMessage("DELETE SSH Key", sshKeyName, nil)
}
func TestSSHKeyGetAll(t *testing.T) {
	// check for testing variables
	testPreCheck(t)

	// create session and client
	session := getSession(t)
	sshKeyClient := client.NewIBMPIKeyClient(context.Background(), session, cloudInstanceID)

	// GET ALL
	getAllResp, err := sshKeyClient.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	testMessage("GET All SSH Keys", "", *getAllResp)
}
