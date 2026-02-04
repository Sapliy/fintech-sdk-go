# \ZoneServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZoneServiceBulkUpdateMetadata**](ZoneServiceAPI.md#ZoneServiceBulkUpdateMetadata) | **Post** /v1/zones/bulk-metadata | 
[**ZoneServiceCreateZone**](ZoneServiceAPI.md#ZoneServiceCreateZone) | **Post** /v1/zones | 
[**ZoneServiceGetZone**](ZoneServiceAPI.md#ZoneServiceGetZone) | **Get** /v1/zones/{id} | 
[**ZoneServiceListZones**](ZoneServiceAPI.md#ZoneServiceListZones) | **Get** /v1/zones | 
[**ZoneServiceUpdateZone**](ZoneServiceAPI.md#ZoneServiceUpdateZone) | **Put** /v1/zones/{id} | 



## ZoneServiceBulkUpdateMetadata

> ZoneBulkUpdateMetadataResponse ZoneServiceBulkUpdateMetadata(ctx).Body(body).Execute()



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
	body := *openapiclient.NewZoneBulkUpdateMetadataRequest() // ZoneBulkUpdateMetadataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneServiceAPI.ZoneServiceBulkUpdateMetadata(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneServiceAPI.ZoneServiceBulkUpdateMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneServiceBulkUpdateMetadata`: ZoneBulkUpdateMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `ZoneServiceAPI.ZoneServiceBulkUpdateMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZoneServiceBulkUpdateMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ZoneBulkUpdateMetadataRequest**](ZoneBulkUpdateMetadataRequest.md) |  | 

### Return type

[**ZoneBulkUpdateMetadataResponse**](ZoneBulkUpdateMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneServiceCreateZone

> ZoneZone ZoneServiceCreateZone(ctx).Body(body).Execute()



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
	body := *openapiclient.NewZoneCreateZoneRequest() // ZoneCreateZoneRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneServiceAPI.ZoneServiceCreateZone(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneServiceAPI.ZoneServiceCreateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneServiceCreateZone`: ZoneZone
	fmt.Fprintf(os.Stdout, "Response from `ZoneServiceAPI.ZoneServiceCreateZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZoneServiceCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ZoneCreateZoneRequest**](ZoneCreateZoneRequest.md) |  | 

### Return type

[**ZoneZone**](ZoneZone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneServiceGetZone

> ZoneZone ZoneServiceGetZone(ctx, id).Execute()



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
	resp, r, err := apiClient.ZoneServiceAPI.ZoneServiceGetZone(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneServiceAPI.ZoneServiceGetZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneServiceGetZone`: ZoneZone
	fmt.Fprintf(os.Stdout, "Response from `ZoneServiceAPI.ZoneServiceGetZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneServiceGetZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ZoneZone**](ZoneZone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneServiceListZones

> ZoneListZonesResponse ZoneServiceListZones(ctx).OrgId(orgId).Execute()



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
	orgId := "orgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneServiceAPI.ZoneServiceListZones(context.Background()).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneServiceAPI.ZoneServiceListZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneServiceListZones`: ZoneListZonesResponse
	fmt.Fprintf(os.Stdout, "Response from `ZoneServiceAPI.ZoneServiceListZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZoneServiceListZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** |  | 

### Return type

[**ZoneListZonesResponse**](ZoneListZonesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneServiceUpdateZone

> ZoneZone ZoneServiceUpdateZone(ctx, id).Body(body).Execute()



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
	body := *openapiclient.NewZoneServiceUpdateZoneBody() // ZoneServiceUpdateZoneBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneServiceAPI.ZoneServiceUpdateZone(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneServiceAPI.ZoneServiceUpdateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZoneServiceUpdateZone`: ZoneZone
	fmt.Fprintf(os.Stdout, "Response from `ZoneServiceAPI.ZoneServiceUpdateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneServiceUpdateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ZoneServiceUpdateZoneBody**](ZoneServiceUpdateZoneBody.md) |  | 

### Return type

[**ZoneZone**](ZoneZone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

