# \EventsAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmitEvent**](EventsAPI.md#EmitEvent) | **Post** /v1/events/emit | Emit an Event
[**GetPastEvents**](EventsAPI.md#GetPastEvents) | **Get** /v1/zones/{zoneId}/events/past | Get Past Events (Webhook Replay)
[**ReplayEvent**](EventsAPI.md#ReplayEvent) | **Post** /v1/events/{eventId}/replay | Replay an Event



## EmitEvent

> EmitEvent202Response EmitEvent(ctx).EmitEventRequest(emitEventRequest).Execute()

Emit an Event

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
	emitEventRequest := *openapiclient.NewEmitEventRequest("Type_example") // EmitEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EmitEvent(context.Background()).EmitEventRequest(emitEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EmitEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmitEvent`: EmitEvent202Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EmitEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmitEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emitEventRequest** | [**EmitEventRequest**](EmitEventRequest.md) |  | 

### Return type

[**EmitEvent202Response**](EmitEvent202Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPastEvents

> GetPastEvents200Response GetPastEvents(ctx, zoneId).Limit(limit).Offset(offset).Execute()

Get Past Events (Webhook Replay)

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
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetPastEvents(context.Background(), zoneId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetPastEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPastEvents`: GetPastEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetPastEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPastEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**GetPastEvents200Response**](GetPastEvents200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplayEvent

> ReplayEvent200Response ReplayEvent(ctx, eventId).ReplayEventRequest(replayEventRequest).Execute()

Replay an Event

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
	eventId := "eventId_example" // string | 
	replayEventRequest := *openapiclient.NewReplayEventRequest("ZoneId_example") // ReplayEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ReplayEvent(context.Background(), eventId).ReplayEventRequest(replayEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ReplayEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayEvent`: ReplayEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ReplayEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replayEventRequest** | [**ReplayEventRequest**](ReplayEventRequest.md) |  | 

### Return type

[**ReplayEvent200Response**](ReplayEvent200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

