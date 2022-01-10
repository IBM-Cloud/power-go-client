# IBM Power-Go-Client
The power-go-client project provides the Go SDK for IBM® Power Systems™ Virtual Server.

## Requirements
 - An [IBM Cloud Account](https://cloud.ibm.com/registration)
 - An IBM Cloud [IAM API key](https://cloud.ibm.com/docs/account?topic=account-userapikey) or [Token](https://cloud.ibm.com/docs/account?topic=account-iamtoken_from_apikey) allow the SDK to access your account
 - [Go](https://golang.org/doc/install) v1.17 or above.
 - [Go Swagger](https://goswagger.io/install.html) v0.19.0

## Install
 - Install the SDK
    ```
    go get github.com/IBM-Cloud/power-go-client
    ```
 - Update the SDK
    ```
    go get -u github.com/IBM-Cloud/power-go-client
    ```
 - If your application is using the `dep` dependency management tool, you can add a dependency
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

Also you can refer to the test directory for some resources that shows how to use the SDK.

## Running Tests
Runnning tests will create real IBM Power Cloud resources. Example test configurations `example.launch.json` and `example.env` are located in the `test` directory. To use these configurations, you will need to create a `.vscode` folder in the root directory. Then rename/copy `launch.json` and `.env` to this folder.

You can define all environment variables in `.env`, or you can create additional `.env` files (`staging.env`, `production.env`, `dal12.env`, ...) for different environments, regions, resources, etc. If you create a different `.env` file, modify `launch.json` to include a testing configuration for that file.

Each individual test requires specific environment variables, but all tests require these variables:
 - `API_KEY` or `BEARER_TOKEN` (deprecated but currently working)
 - `CRN` or (`ACCOUNT_ID`, `CLOUD_INSTANCE_ID`, `REGION`, `ZONE`, `IBMCLOUD_POWER_API_ENDPOINT`)

All account, cloud resource, and testing variables are declared in `test_variables.go`. Each test starts with these steps. First, `init()` in `test_utils.go` is called. This imports all environment variable definitions into the `test_variables.go` declarations. Second, a preCheck function (ex. `testImagePreCheck()`) is called. This verifies that all required variables for the preCheck's test are defined and loaded properly. Finally, the test runs.

To run a test:
 - Double click the test function name to select the text
 - Click `Run and Debug` on the VScode sidebar, and then click `Run a test`. Output will be visible in the VScode Debug Console.
   - `launch.json` runs tests by using this selected text. Feel free to create your own `launch.json` configuration.

When updating or creating new tests:
 - Add environment and test variable definitions to `example.env` and `test_variables.go`
 - Update `init()` and relevant `PreCheck()` functions in `test_utils.go`
 - Make sure to add a way for newly provisioned resources to be deleted at the end of the test

## Updating the IBM Power API
The Power-Go-Client direclty uses the IBM Power Cloud API. API definitions are imported into the `power` directory. To use the most recent version of the IBM Power Cloud API, run the `./generate.sh` script in the service broker/...power-go-client directory. Copy the generated files into the `power` directory.

## Issues
If you encounter an issue with the project please [report here.](https://github.com/IBM-Cloud/power-go-client/issues).