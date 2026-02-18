# \ExecutionsAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecution**](ExecutionsAPI.md#GetExecution) | **Get** /v1/executions/{executionId} | Get Execution details
[**ResumeExecution**](ExecutionsAPI.md#ResumeExecution) | **Post** /v1/executions/{executionId}/resume | Resume a paused Execution



## GetExecution

> FlowExecution GetExecution(ctx, executionId).Execute()

Get Execution details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	executionId := "executionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionsAPI.GetExecution(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionsAPI.GetExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecution`: FlowExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionsAPI.GetExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowExecution**](FlowExecution.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeExecution

> ResumeExecution200Response ResumeExecution(ctx, executionId).RequestBody(requestBody).Execute()

Resume a paused Execution

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	executionId := "executionId_example" // string | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionsAPI.ResumeExecution(context.Background(), executionId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionsAPI.ResumeExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeExecution`: ResumeExecution200Response
	fmt.Fprintf(os.Stdout, "Response from `ExecutionsAPI.ResumeExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**ResumeExecution200Response**](ResumeExecution200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

