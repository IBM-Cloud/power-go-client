# IBM Power-Go-Client
The Power-Go-Client is a Go SDK for the IBM Power Cloud. This repository implements functions for creating, managing, and deleting IBM Power Cloud resources.

## Requirements
 - [Go](https://golang.org/doc/install) v1.13
 - [Go Swagger](https://goswagger.io/install.html) v0.19.0

## Using the SDK
 - Install the SDK
    ```
    go get github.com/IBM-Cloud/power-go-client
    ```
 - Update the SDK
    ```
    go get -u github.com/IBM-Cloud/power-go-client
    ```
## Development Setup
 - Fork and Clone this repository
   ```shell
   git clone https://github.com/<Your Fork>/power-go-client.git
   cd power-go-client
   git remote add upstream https://github.com/IBM-Cloud/power-go-client.git
   git remote -vv
      origin	https://github.com/<Your Fork>/power-go-client.git (fetch)
      origin	https://github.com/<Your Fork>/power-go-client.git (push)
      upstream	https://github.com/IBM-Cloud/power-go-client.git (fetch)
      upstream	https://github.com/IBM-Cloud/power-go-client.git (push)
   ```
 - Create an [IBM Cloud Account](https://cloud.ibm.com/registration)
 - Create an IBM Cloud API Key [production](https://cloud.ibm.com/iam/apikeys) | [staging](https://test.cloud.ibm.com/iam/apikeys)
 
## Running Tests
Runnning examples in the `test` directory will create real IBM Power Cloud resources. Test configurations are located in `.vscode/launch.json` and test environment variables are defined in `.vscode/.env` by default. You will need to rename `example.env` as `.env`.

You should define all environment variables in `.env`, or you can create additional `.env` files (`staging.env`, `production.env`, `dal12.env`, ...) for different environments, regions, resources, etc. If you create a different `.env` file, modify `launch.json` to include a testing configuration for that file.

Each individual test may require specific environment variables, but all tests require these variables:
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
The Power-Go-Client direclty uses the IBM Power Cloud API. API definitions are imported into the `power` directory. To use the most recent version of the IBM Power Cloud API, run the `./power-go.sh` script in the IBM Power Cloud API respository and copy the generated files into the `power` directory.

