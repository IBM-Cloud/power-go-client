// package main

// import (
// 	"log"
// 	"os"
// 	"time"

// 	v "github.com/IBM-Cloud/power-go-client/clients/instance"
// 	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
// 	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
// 	"github.com/IBM-Cloud/power-go-client/power/models"
// )

// func main() {

// 	//session Inputs
// 	token := "Bearer eyJraWQiOiIyMDIxMTAxODA4MTkiLCJhbGciOiJSUzI1NiJ9.eyJpYW1faWQiOiJJQk1pZC02NjYwMDBEOEZSIiwiaWQiOiJJQk1pZC02NjYwMDBEOEZSIiwicmVhbG1pZCI6IklCTWlkIiwic2Vzc2lvbl9pZCI6IkMtZDQ0MDhjZTgtYzRlZS00ZGJiLTkyMTctMGU1NTQxODU1OWM3IiwianRpIjoiYWU5YmVjYTYtZGE4ZC00ZjhiLTlkYzktYTlhNTBhODM0ZDliIiwiaWRlbnRpZmllciI6IjY2NjAwMEQ4RlIiLCJnaXZlbl9uYW1lIjoiQ2hhc2UiLCJmYW1pbHlfbmFtZSI6IkF1c3RpbiIsIm5hbWUiOiJDaGFzZSBBdXN0aW4iLCJlbWFpbCI6ImNoYXNlQGlibS5jb20iLCJzdWIiOiJjaGFzZUBpYm0uY29tIiwiYXV0aG4iOnsic3ViIjoiY2hhc2VAaWJtLmNvbSIsImlhbV9pZCI6IklCTWlkLTY2NjAwMEQ4RlIiLCJuYW1lIjoiQ2hhc2UgQXVzdGluIiwiZ2l2ZW5fbmFtZSI6IkNoYXNlIiwiZmFtaWx5X25hbWUiOiJBdXN0aW4iLCJlbWFpbCI6ImNoYXNlQGlibS5jb20ifSwiYWNjb3VudCI6eyJib3VuZGFyeSI6Imdsb2JhbCIsInZhbGlkIjp0cnVlLCJic3MiOiJkOWNlYzgwZDBhZGM0MDBlYWQ4ZTIwNzZhZmUyNjY5OCIsImltc191c2VyX2lkIjoiOTAzNTgyMiIsImltcyI6IjE3NDg1MjMifSwiaWF0IjoxNjM0NjcyNzgwLCJleHAiOjE2MzQ2NzM5ODAsImlzcyI6Imh0dHBzOi8vaWFtLmNsb3VkLmlibS5jb20vaWRlbnRpdHkiLCJncmFudF90eXBlIjoidXJuOmlibTpwYXJhbXM6b2F1dGg6Z3JhbnQtdHlwZTpwYXNzY29kZSIsInNjb3BlIjoiaWJtIG9wZW5pZCIsImNsaWVudF9pZCI6ImJ4IiwiYWNyIjoxLCJhbXIiOlsicHdkIl19.LITkeH0p-6roll-RtelR6xQNYDuTqDcwp9UBoeZ8yi_lnaya8DEdRj5vUpAhyhHtRPYLqawPq0tupqH01MbZ54Ojgsg-IeFuR8UuhSQFJQtjlULda_lU9-mCyXeTqoohVtqgo5s-VkJsqK0rANY16vUuU1pNxTRWbt6F4lg6S80UlIWr3Zl2_AV_QmtZpBafU2pLpetL9iGOvMptw5zItqjXoBBmzzz-auYi1axTb0ysbpEHjQQ1OvytGCBI3eLhUTteBGF9sWDUqbXSjMo_9Hd7hhQyyh2hBIwW_qj4qp6CZsmV37h-K3QMkR364U3zS4hlTnRxvtdJ4bqDYiLTdg"
// 	region := "dal"
// 	accountID := "efe5e8b9d3f04b948790fe5499bd18bc"

// 	// volume inputs
// 	name := "Chase-terraform-bug-test-1"
// 	piID := "6021a723-bcab-4d3f-9985-d0cb2f864f35"
// 	//timeout := time.Duration(9000000000000000000)
// 	imageID := "dc14d908-9836-42a2-848d-a3af6d309f6b"
// 	volumes := make([]string, 1)
// 	volumes[0] = "e2823fba-5480-422d-8786-77982ebc14ef"

// 	networks := make([]string, 1)
// 	networks[0] = "1bfad140-588f-45d3-aaaa-a3f0abbd441f"
// 	memory := 4.0
// 	processors := 2.0
// 	procType := "shared"
// 	key := "CHASE-TEST-4"
// 	//sysType := "s922"

// 	os.Setenv("IBMCLOUD_POWER_API_ENDPOINT", region+".power-iaas.test.cloud.ibm.com")

// 	session, err := ps.New(token, region, true, 9000000000000000000, accountID, region)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	powerClient := v.NewIBMPIInstanceClient(session, piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	params := p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
// 		Body: &models.PVMInstanceCreate{
// 			ImageID:     &imageID,
// 			KeyPairName: key,
// 			NetworkIds:  networks,
// 			ServerName:  &name,
// 			VolumeIds:   volumes,
// 			Memory:      &memory,
// 			Processors:  &processors,
// 			ProcType:    &procType,
// 			//SysType:     sysType,
// 		},
// 		CloudInstanceID: piID,
// 	}
// 	createRespOk, err := powerClient.Create(&params, piID, 360*time.Second)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[1]****************** %+v\n\n\n\n", *createRespOk)

// 	insIDs := make([]string, 0)
// 	for _, in := range *createRespOk {
// 		insID := in.PvmInstanceID
// 		insIDs = append(insIDs, *insID)
// 	}

// 	getResp, err := powerClient.Get(insIDs[0], piID, timeout)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[2]****************** %+v \n", *getResp)

// 	getallResp, err := powerClient.GetAll(piID, timeout)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[3]****************** %+v \n", *getallResp)

// 	err = powerClient.Delete(insIDs[0], piID, timeout)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
