# \NotificationServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationServiceCreateWebhookEndpoint**](NotificationServiceAPI.md#NotificationServiceCreateWebhookEndpoint) | **Post** /v1/webhooks/endpoints | 
[**NotificationServiceDeleteWebhookEndpoint**](NotificationServiceAPI.md#NotificationServiceDeleteWebhookEndpoint) | **Delete** /v1/webhooks/endpoints/{id} | 
[**NotificationServiceGetNotificationHistory**](NotificationServiceAPI.md#NotificationServiceGetNotificationHistory) | **Get** /v1/notifications | 
[**NotificationServiceListWebhookEndpoints**](NotificationServiceAPI.md#NotificationServiceListWebhookEndpoints) | **Get** /v1/webhooks/endpoints | 



## NotificationServiceCreateWebhookEndpoint

> NotificationsWebhookEndpoint NotificationServiceCreateWebhookEndpoint(ctx).Body(body).Execute()



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
	body := *openapiclient.NewNotificationsCreateWebhookEndpointRequest() // NotificationsCreateWebhookEndpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationServiceAPI.NotificationServiceCreateWebhookEndpoint(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceAPI.NotificationServiceCreateWebhookEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationServiceCreateWebhookEndpoint`: NotificationsWebhookEndpoint
	fmt.Fprintf(os.Stdout, "Response from `NotificationServiceAPI.NotificationServiceCreateWebhookEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceCreateWebhookEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationsCreateWebhookEndpointRequest**](NotificationsCreateWebhookEndpointRequest.md) |  | 

### Return type

[**NotificationsWebhookEndpoint**](NotificationsWebhookEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationServiceDeleteWebhookEndpoint

> NotificationsDeleteWebhookEndpointResponse NotificationServiceDeleteWebhookEndpoint(ctx, id).Execute()



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
	resp, r, err := apiClient.NotificationServiceAPI.NotificationServiceDeleteWebhookEndpoint(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceAPI.NotificationServiceDeleteWebhookEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationServiceDeleteWebhookEndpoint`: NotificationsDeleteWebhookEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationServiceAPI.NotificationServiceDeleteWebhookEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceDeleteWebhookEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsDeleteWebhookEndpointResponse**](NotificationsDeleteWebhookEndpointResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationServiceGetNotificationHistory

> NotificationsGetNotificationHistoryResponse NotificationServiceGetNotificationHistory(ctx).UserId(userId).Limit(limit).Offset(offset).Execute()



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
	userId := "userId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationServiceAPI.NotificationServiceGetNotificationHistory(context.Background()).UserId(userId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceAPI.NotificationServiceGetNotificationHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationServiceGetNotificationHistory`: NotificationsGetNotificationHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationServiceAPI.NotificationServiceGetNotificationHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceGetNotificationHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**NotificationsGetNotificationHistoryResponse**](NotificationsGetNotificationHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationServiceListWebhookEndpoints

> NotificationsListWebhookEndpointsResponse NotificationServiceListWebhookEndpoints(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationServiceAPI.NotificationServiceListWebhookEndpoints(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceAPI.NotificationServiceListWebhookEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationServiceListWebhookEndpoints`: NotificationsListWebhookEndpointsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationServiceAPI.NotificationServiceListWebhookEndpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceListWebhookEndpointsRequest struct via the builder pattern


### Return type

[**NotificationsListWebhookEndpointsResponse**](NotificationsListWebhookEndpointsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

