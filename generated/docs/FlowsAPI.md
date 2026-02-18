# \FlowsAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlow**](FlowsAPI.md#CreateFlow) | **Post** /v1/flows | Create a Flow
[**DeleteFlow**](FlowsAPI.md#DeleteFlow) | **Delete** /v1/flows/{flowId} | Delete a Flow
[**GetFlow**](FlowsAPI.md#GetFlow) | **Get** /v1/flows/{flowId} | Get Flow details
[**ListFlows**](FlowsAPI.md#ListFlows) | **Get** /v1/zones/{zoneId}/flows | List Flows in a Zone
[**UpdateFlow**](FlowsAPI.md#UpdateFlow) | **Put** /v1/flows/{flowId} | Update a Flow



## CreateFlow

> Flow CreateFlow(ctx).Flow(flow).Execute()

Create a Flow

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
	flow := *openapiclient.NewFlow("Id_example", "ZoneId_example", "Name_example") // Flow | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.CreateFlow(context.Background()).Flow(flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.CreateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlow`: Flow
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.CreateFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | [**Flow**](Flow.md) |  | 

### Return type

[**Flow**](Flow.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlow

> DeleteFlow(ctx, flowId).Execute()

Delete a Flow

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
	flowId := "flowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.DeleteFlow(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.DeleteFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlow

> Flow GetFlow(ctx, flowId).Execute()

Get Flow details

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
	flowId := "flowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.GetFlow(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlow`: Flow
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Flow**](Flow.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlows

> ListFlows200Response ListFlows(ctx, zoneId).Execute()

List Flows in a Zone

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
	zoneId := "zoneId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ListFlows(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ListFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlows`: ListFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ListFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListFlows200Response**](ListFlows200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlow

> Flow UpdateFlow(ctx, flowId).Flow(flow).Execute()

Update a Flow

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
	flowId := "flowId_example" // string | 
	flow := *openapiclient.NewFlow("Id_example", "ZoneId_example", "Name_example") // Flow | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.UpdateFlow(context.Background(), flowId).Flow(flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.UpdateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFlow`: Flow
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.UpdateFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flow** | [**Flow**](Flow.md) |  | 

### Return type

[**Flow**](Flow.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

