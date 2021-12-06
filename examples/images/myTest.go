// package main

// import (
// 	"log"
// 	"os"

// 	v "github.com/IBM-Cloud/power-go-client/clients/instance"
// 	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
// )

// func main() {

// 	//session Inputs
// 	token := "Bearer eyJraWQiOiIyMDIxMTAxODA4MTkiLCJhbGciOiJSUzI1NiJ9.eyJpYW1faWQiOiJJQk1pZC02NjYwMDBEOEZSIiwiaWQiOiJJQk1pZC02NjYwMDBEOEZSIiwicmVhbG1pZCI6IklCTWlkIiwic2Vzc2lvbl9pZCI6IkMtZDQ0MDhjZTgtYzRlZS00ZGJiLTkyMTctMGU1NTQxODU1OWM3IiwianRpIjoiYWU5YmVjYTYtZGE4ZC00ZjhiLTlkYzktYTlhNTBhODM0ZDliIiwiaWRlbnRpZmllciI6IjY2NjAwMEQ4RlIiLCJnaXZlbl9uYW1lIjoiQ2hhc2UiLCJmYW1pbHlfbmFtZSI6IkF1c3RpbiIsIm5hbWUiOiJDaGFzZSBBdXN0aW4iLCJlbWFpbCI6ImNoYXNlQGlibS5jb20iLCJzdWIiOiJjaGFzZUBpYm0uY29tIiwiYXV0aG4iOnsic3ViIjoiY2hhc2VAaWJtLmNvbSIsImlhbV9pZCI6IklCTWlkLTY2NjAwMEQ4RlIiLCJuYW1lIjoiQ2hhc2UgQXVzdGluIiwiZ2l2ZW5fbmFtZSI6IkNoYXNlIiwiZmFtaWx5X25hbWUiOiJBdXN0aW4iLCJlbWFpbCI6ImNoYXNlQGlibS5jb20ifSwiYWNjb3VudCI6eyJib3VuZGFyeSI6Imdsb2JhbCIsInZhbGlkIjp0cnVlLCJic3MiOiJkOWNlYzgwZDBhZGM0MDBlYWQ4ZTIwNzZhZmUyNjY5OCIsImltc191c2VyX2lkIjoiOTAzNTgyMiIsImltcyI6IjE3NDg1MjMifSwiaWF0IjoxNjM0NjcyNzgwLCJleHAiOjE2MzQ2NzM5ODAsImlzcyI6Imh0dHBzOi8vaWFtLmNsb3VkLmlibS5jb20vaWRlbnRpdHkiLCJncmFudF90eXBlIjoidXJuOmlibTpwYXJhbXM6b2F1dGg6Z3JhbnQtdHlwZTpwYXNzY29kZSIsInNjb3BlIjoiaWJtIG9wZW5pZCIsImNsaWVudF9pZCI6ImJ4IiwiYWNyIjoxLCJhbXIiOlsicHdkIl19.LITkeH0p-6roll-RtelR6xQNYDuTqDcwp9UBoeZ8yi_lnaya8DEdRj5vUpAhyhHtRPYLqawPq0tupqH01MbZ54Ojgsg-IeFuR8UuhSQFJQtjlULda_lU9-mCyXeTqoohVtqgo5s-VkJsqK0rANY16vUuU1pNxTRWbt6F4lg6S80UlIWr3Zl2_AV_QmtZpBafU2pLpetL9iGOvMptw5zItqjXoBBmzzz-auYi1axTb0ysbpEHjQQ1OvytGCBI3eLhUTteBGF9sWDUqbXSjMo_9Hd7hhQyyh2hBIwW_qj4qp6CZsmV37h-K3QMkR364U3zS4hlTnRxvtdJ4bqDYiLTdg"
// 	region := "dal"
// 	accountID := "efe5e8b9d3f04b948790fe5499bd18bc"

// 	// Image inputs
// 	//name := "Chase-power-go"
// 	piID := "6021a723-bcab-4d3f-9985-d0cb2f864f35"
// 	//image := "b4966c83-5eaf-4247-a5ba-e9e80324b44b"

// 	// crn:v1:staging:public:power-iaas:dal12:a/efe5e8b9d3f04b948790fe5499bd18bc:6021a723-bcab-4d3f-9985-d0cb2f864f35::

// 	os.Setenv("IBMCLOUD_POWER_API_ENDPOINT", region+".power-iaas.test.cloud.ibm.com")

// 	session, err := ps.New(token, region, true, 50000000, accountID, "dal12")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	powerClient := v.NewIBMPIImageClient(session, piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	createResp, err := powerClient.Create(name, image, piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[1]****************** %+v \n\n", *createResp)

// 	imageID := *createResp.ImageID
// 	getResp, err := powerClient.Get(imageID, piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[2]****************** %+v \n\n", *getResp)

// 	getAllResp, err := powerClient.GetAll(piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[3]****************** %+v \n\n", *getAllResp)

// 	err = powerClient.Delete(imageID, piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	getSapResp, err := powerClient.GetSAPImages(piID, true)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[4]****************** %+v \n\n", *getSapResp)

// 	getStockResp, err := powerClient.GetStockImages(piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[5]****************** %+v \n\n", *getStockResp)

// 	getVtlImages, err := powerClient.GetAllVtlImages(piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[6]****************** %+v \n\n", *getVtlImages)

// 	if len(getVtlImages.Images) > 0 {
// 		testVtlId := *getVtlImages.Images[0].ImageID
// 		imageCheck, err := powerClient.IsVtlImage(testVtlId, piID)

// 		msg := fmt.Sprintf("IsVtlImage returned true for vtl image with ID %s", testVtlId)
// 		if imageCheck == false {
// 			msg = fmt.Sprintf("IsVtlImage returned false for image with ID %s. Error: %s", testVtlId, err)
// 		}
// 		log.Printf("***************[7]****************** %s \n\n", msg)
// 	}
// }

// getStockResp, err := powerClient.GetAllStockImages(piID, true, true)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[4]****************** %+v \n\n", *getStockResp)

// 	getSapResp, err := powerClient.GetAllStockSAPImages(piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[5]****************** %+v \n\n", *getSapResp)

// 	getVtlImages, err := powerClient.GetAllStockVTLImages(piID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("***************[6]****************** %+v \n\n", *getVtlImages)

// 	if len(getVtlImages.Images) > 0 {
// 		testVtlId := *getVtlImages.Images[0].ImageID
// 		imageCheck, err := powerClient.IsVtlImage(testVtlId, piID)

// 		msg := fmt.Sprintf("IsVtlImage returned true for vtl image with ID %s", testVtlId)
// 		if imageCheck == false {
// 			msg = fmt.Sprintf("IsVtlImage returned false for image with ID %s. Error: %s", testVtlId, err)
// 		}
// 		log.Printf("***************[7]****************** %s \n\n", msg)
// 	}