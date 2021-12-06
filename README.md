# IBM Cloud SDK for Power Cloud
Power-Go-Client provides the Go implementation for operating the IBM Power Cloud platform to build LPARS .

## Installing
    go get github.com/IBM-Cloud/power-go-client

Update the SDK to the latest version using the following command

    go get -u github.com/IBM-Cloud/power-go-client

Install [go-swagger v0.19.0](https://goswagger.io/install.html)

## Using the SDK

You must have a working IBM Cloud account to use the APIs. Sign up if you don't have one.
The SDK has an /examples folder which cites few examples on how to use the SDK. 

First you need to create a session

    import ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
    ...
    session, err := ps.New()

Creating session in this way uses environment variables defined in .env.staging and .env.production. You must define the following environment variables.

    ACCOUNT_ID = "abc123..."
    BEARER_TOKEN = "Bearer abc123..."
    REGION = "dal"
    ZONE = "12"
    CLOUD_INSTANCE_ID = "abc123..."

    // Optional
    DEBUG = "true"
    TIMEOUT = "9000000000000000000"