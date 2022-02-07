# IBM Cloud SDK for Power Cloud

The power-go-client project provides the Go SDK for IBM® Power Systems™ Virtual Server.

## Prerequisites

- An [IBM Cloud Account](https://cloud.ibm.com/registration)
- An IBM Cloud [IAM API key](https://cloud.ibm.com/docs/account?topic=account-userapikey) or [Token](https://cloud.ibm.com/docs/account?topic=account-iamtoken_from_apikey) allow the SDK to access your account.
- Go version 1.17 or above.

## Install

Install the SDK using the following methods.

### `go get` command

Use this command to download and install the SDK to allow your Go application to use it:
```
go get -u github.com/IBM-Cloud/power-go-client
```

### Go modules

If your application is using Go modules, you can add a suitable import to your Go application, like this:
```go
import (
  "github.com/IBM-Cloud/power-go-client"
)
```
then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

### `dep` dependency manager

If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file. Here is an example:
```
[[constraint]]
  name = "github.com/IBM-Cloud/power-go-client"
  version = "1.1.0"

```
then run `dep ensure`.

## Usage

First you need to create a session and use it for creating the client.

```golang
import "github.com/IBM-Cloud/power-go-client/ibmpisession"

func main(){
    o := &ibmpisession.IBMPIOptions{
		Authenticator: authenticator,
		UserAccount:   accountID,
		Zone:          zone,
		URL:           url,
	}
    s, err := ibmpisession.NewIBMPISession(o)
    .....
}
```
- `authenticator`: Please check https://github.com/IBM/go-sdk-core/blob/main/Authentication.md for different options available for authenticating API calls.
- `accountID`: Account id of the Power Cloud Service Instance
- `zone`: Zone of the Power Cloud Service Instance
- `url`: Power Virtual Server host or URL endpoint. You can also set env variable `IBMCLOUD_POWER_API_ENDPOINT`.

Also you can refer to the [examples](examples) directory for some resources that shows how to use the SDK.


## Issues

If you encounter an issue with the project please [report here.](https://github.com/IBM-Cloud/power-go-client/issues).

## Development

This section contains some steps required by the dev for contributing changes to the SDK.

### Update the power directory

1. Install the latest [go-swagger](https://github.com/go-swagger/go-swagger/releases) version on your machine.
1. Download the last released version of the service broker API definitions (`swagger.yaml`).
1. Run `./bin/generate-swagger.sh` from the project root.
