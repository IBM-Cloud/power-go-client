package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/examples/clientutils"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
)

func main() {

	crn, ok := os.LookupEnv("IBMCLOUD_POWERVS_CRN")
	if !ok {
		fmt.Println("crn  is not set. Please set the variable IBMCLOUD_POWERVS_CRN")
		return
	} else {
		fmt.Printf("crn value is : %s\n", crn)
	}

	// IBMCLOUD_STORAGE_TIER
	storage_type, ok := os.LookupEnv("IBMCLOUD_STORAGE_TIER")
	if !ok {
		fmt.Println("storage Tier is not set. Please set the variable IBMCLOUD_STORAGE_TIER")
		return
	} else {
		fmt.Printf("storage tier value is : %s\n", storage_type)
	}
	// IBMCLOUD_STORAGE_POOL_NAME
	poolName, ok := os.LookupEnv("IBMCLOUD_STORAGE_POOL_NAME")
	if !ok {
		fmt.Println("storage pool name is not set.Please set the variable IBMCLOUD_STORAGE_POOL_NAME" +
			"Sample value is Tier3-Flash-1 etc")
		return
	} else {
		fmt.Printf("storage pool  value is : %s\n", poolName)
	}

	// IBMCLOUD_IAM_TOKEN
	token, ok := os.LookupEnv("IBMCLOUD_IAM_TOKEN")
	if !ok {
		fmt.Println("token is not set.Please set the variable IBMCLOUD_IAM_TOKEN" +
			"You can get this by running ibmcloud iam oauth-tokens\n")
		return
	} else {
		fmt.Printf("Token is value is : %s\n", token)
	}

	cloud_crn := strings.Split(crn, ":")
	str := strings.Join(cloud_crn, "")
	fmt.Println(str)

	cloud_account_id := strings.Split(cloud_crn[6], "/")
	accountid := cloud_account_id[1]

	timeout := time.Duration(9000000000000000000)

	session, err := ps.New(token, clientutils.Zonemaps(cloud_crn[5]), true, 9000000000000000000, accountid, cloud_crn[5])
	if err != nil {
		log.Fatal(err)
	}
	storage_capacity_client := v.NewIBMPIStorageCapacityClient(session, cloud_crn[7])
	if err != nil {
		log.Fatal(err)
	}

	getResp, err := storage_capacity_client.GetAllStoragePools(cloud_crn[7], timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n", *getResp)

	getallResp, err := storage_capacity_client.GetAvailableStorageCapacity(cloud_crn[7], storage_type, timeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getallResp)

	//getpoolResp, err := storage_capacity_client.GetAvailableStoragePool(cloud_crn[7],poolName,timeout)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("***************[3]****************** %+v \n", *getpoolResp.MaxAllocationSize)

}
