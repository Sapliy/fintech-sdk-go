# \FlowServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowServiceBulkUpdateFlows**](FlowServiceAPI.md#FlowServiceBulkUpdateFlows) | **Post** /v1/flows/bulk-update | 
[**FlowServiceCreateFlow**](FlowServiceAPI.md#FlowServiceCreateFlow) | **Post** /v1/flows | 
[**FlowServiceGetExecution**](FlowServiceAPI.md#FlowServiceGetExecution) | **Get** /v1/flows/executions/{id} | 
[**FlowServiceGetFlow**](FlowServiceAPI.md#FlowServiceGetFlow) | **Get** /v1/flows/{id} | 
[**FlowServiceListFlows**](FlowServiceAPI.md#FlowServiceListFlows) | **Get** /v1/flows | 
[**FlowServiceResumeExecution**](FlowServiceAPI.md#FlowServiceResumeExecution) | **Post** /v1/flows/executions/{executionId}/resume | 
[**FlowServiceUpdateFlow**](FlowServiceAPI.md#FlowServiceUpdateFlow) | **Put** /v1/flows/{id} | 



## FlowServiceBulkUpdateFlows

> FlowBulkUpdateFlowsResponse FlowServiceBulkUpdateFlows(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := *openapiclient.NewFlowBulkUpdateFlowsRequest() // FlowBulkUpdateFlowsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowServiceAPI.FlowServiceBulkUpdateFlows(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowServiceAPI.FlowServiceBulkUpdateFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowServiceBulkUpdateFlows`: FlowBulkUpdateFlowsResponse
	fmt.Fprintf(os.Stdout, "Response from `FlowServiceAPI.FlowServiceBulkUpdateFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowServiceBulkUpdateFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FlowBulkUpdateFlowsRequest**](FlowBulkUpdateFlowsRequest.md) |  | 

### Return type

[**FlowBulkUpdateFlowsResponse**](FlowBulkUpdateFlowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowServiceCreateFlow

> FlowFlow FlowServiceCreateFlow(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := *openapiclient.NewFlowCreateFlowRequest() // FlowCreateFlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowServiceAPI.FlowServiceCreateFlow(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowServiceAPI.FlowServiceCreateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowServiceCreateFlow`: FlowFlow
	fmt.Fprintf(os.Stdout, "Response from `FlowServiceAPI.FlowServiceCreateFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowServiceCreateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FlowCreateFlowRequest**](FlowCreateFlowRequest.md) |  | 

### Return type

[**FlowFlow**](FlowFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowServiceGetExecution

> FlowFlowExecution FlowServiceGetExecution(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowServiceAPI.FlowServiceGetExecution(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowServiceAPI.FlowServiceGetExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowServiceGetExecution`: FlowFlowExecution
	fmt.Fprintf(os.Stdout, "Response from `FlowServiceAPI.FlowServiceGetExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowServiceGetExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowFlowExecution**](FlowFlowExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowServiceGetFlow

> FlowFlow FlowServiceGetFlow(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowServiceAPI.FlowServiceGetFlow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowServiceAPI.FlowServiceGetFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowServiceGetFlow`: FlowFlow
	fmt.Fprintf(os.Stdout, "Response from `FlowServiceAPI.FlowServiceGetFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowServiceGetFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowFlow**](FlowFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowServiceListFlows

> FlowListFlowsResponse FlowServiceListFlows(ctx).ZoneId(zoneId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	zoneId := "zoneId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowServiceAPI.FlowServiceListFlows(context.Background()).ZoneId(zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowServiceAPI.FlowServiceListFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowServiceListFlows`: FlowListFlowsResponse
	fmt.Fprintf(os.Stdout, "Response from `FlowServiceAPI.FlowServiceListFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowServiceListFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **string** |  | 

### Return type

[**FlowListFlowsResponse**](FlowListFlowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowServiceResumeExecution

> FlowResumeExecutionResponse FlowServiceResumeExecution(ctx, executionId).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	executionId := "executionId_example" // string | 
	body := *openapiclient.NewFlowServiceResumeExecutionBody() // FlowServiceResumeExecutionBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowServiceAPI.FlowServiceResumeExecution(context.Background(), executionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowServiceAPI.FlowServiceResumeExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowServiceResumeExecution`: FlowResumeExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `FlowServiceAPI.FlowServiceResumeExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowServiceResumeExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FlowServiceResumeExecutionBody**](FlowServiceResumeExecutionBody.md) |  | 

### Return type

[**FlowResumeExecutionResponse**](FlowResumeExecutionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowServiceUpdateFlow

> FlowFlow FlowServiceUpdateFlow(ctx, id).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	body := *openapiclient.NewFlowServiceUpdateFlowBody() // FlowServiceUpdateFlowBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowServiceAPI.FlowServiceUpdateFlow(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowServiceAPI.FlowServiceUpdateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowServiceUpdateFlow`: FlowFlow
	fmt.Fprintf(os.Stdout, "Response from `FlowServiceAPI.FlowServiceUpdateFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowServiceUpdateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FlowServiceUpdateFlowBody**](FlowServiceUpdateFlowBody.md) |  | 

### Return type

[**FlowFlow**](FlowFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

