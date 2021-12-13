# IBM Cloud SDK for Power Cloud
To be updated
power-go-client provides the Go implementation for operating the IBM Power Cloud platform to build LPARS .
Installing

Install the SDK using the following command

```shell
go get github.com/IBM-Cloud/power-go-client
```

Update the SDK to the latest version using the following command

```shell
go get -u github.com/IBM-Cloud/power-go-client
```

Using the SDK

You must have a working IBM Cloud account to use the APIs. Sign up if you don't have one.

The SDK has examples folder which cites few examples on how to use the SDK. First you need to create a session.

```golang
import "github.com/IBM-Cloud/power-go-client/ibmpisession"

func main(){

    s := ibmpowersession.New()
    .....
}
```

Creating session in this way creates a default configuration which reads the value from the environment variables. You must export the following environment variables.

    IBMID - This is the IBM ID
    IBMID_PASSWORD - This is the password for the above ID

OR

    IC_API_KEY/IBMCLOUD_API_KEY - This is the Bluemix API Key. Login to IBMCloud to create one if you don't already have one. See instructions below for creating an API Key.

The default region is us_south. You can override it in the Config struct. 
You can also provide the value via environment variables; either via IC_REGION or IBMCLOUD_REGION. Valid regions for Power Cloud are -

    us-south
    us-east
    eu-de
If working in a multi-zone region like eu-de, the Zone must also be passed in via IBMCLOUD_ZONE. 

    eu-de-1
    eu-de-2    

Multi-Zone support for us-south and us-east is in the IBM Cloud Roadmap. 
