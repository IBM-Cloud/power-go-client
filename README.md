![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM-Cloud/power-go-client)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/IBM-Cloud/power-go-client)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/IBM-Cloud/power-go-client/Go)
![GitHub contributors](https://img.shields.io/github/contributors/IBM-Cloud/power-go-client?color=blueviolet)
![GitHub](https://img.shields.io/github/license/IBM-Cloud/power-go-client)

# IBM Cloud SDK for Power Cloud

The power-go-client project provides the Go SDK for IBM® Power Systems™ Virtual Server.

## Prerequisites

- An [IBM Cloud Account](https://cloud.ibm.com/registration)
- An IBM Cloud [IAM API key](https://cloud.ibm.com/docs/account?topic=account-userapikey) or [Token](https://cloud.ibm.com/docs/account?topic=account-iamtoken_from_apikey) allow the SDK to access your account
- Go 1.17+
  - [Homebrew](https://formulae.brew.sh/formula/go)
  - [Source](https://go.dev/doc/install)

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
	}
    s, err := ibmpisession.NewIBMPISession(o)
    .....
}
```
Type `IBMPIOptions` required fields:
- `Authenticator`: Please check https://github.com/IBM/go-sdk-core/blob/main/Authentication.md for different options available for authenticating API calls.
- `UserAccount`: Account ID of the Power Cloud Service Instance.
- `Zone`: Location of the Power Cloud Service Instance.

Other optional fields:
-	`Debug`: Enable/Disable http transport debugging log.
- `Region`: Region of the Power Cloud Service Instance. This is used for generating the default service URL. *Deprecated*: The region and endpoint is auto generated based on `Zone`.
- `URL`: Power Virtual Server host or URL endpoint. By default it uses `power-iaas.cloud.ibm.com`. In case of test environment you can set the value to `power-iaas.test.cloud.ibm.com`. You can also use env variable `IBMCLOUD_POWER_API_ENDPOINT`. Note that the region value is prepended to the host string if not present eg: `<region>.power-iaas.cloud.ibm.com`.

Also you can refer to the [examples](examples) directory for some resources that shows how to use the SDK.


## Issues

If you encounter an issue with the project please [report here.](https://github.com/IBM-Cloud/power-go-client/issues).

## Development

This section contains some steps required by the dev for contributing changes to the SDK.

### Update the power directory

1. Install the latest [go-swagger](https://github.com/go-swagger/go-swagger/releases) version on your machine.
1. Download the last released version of the service broker API definitions (`swagger.yaml`).
1. Run `./bin/update-swagger.sh` from the project root.