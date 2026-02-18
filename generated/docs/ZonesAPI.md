# \ZonesAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateZone**](ZonesAPI.md#CreateZone) | **Post** /v1/zones | Create a Zone
[**ListZones**](ZonesAPI.md#ListZones) | **Get** /v1/zones | List or Get Zones



## CreateZone

> ListZones200ResponseInner CreateZone(ctx).CreateZoneRequest(createZoneRequest).Execute()

Create a Zone

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
	createZoneRequest := *openapiclient.NewCreateZoneRequest("OrgId_example", "Name_example", "Mode_example") // CreateZoneRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.CreateZone(context.Background()).CreateZoneRequest(createZoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.CreateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateZone`: ListZones200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.CreateZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createZoneRequest** | [**CreateZoneRequest**](CreateZoneRequest.md) |  | 

### Return type

[**ListZones200ResponseInner**](ListZones200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZones

> []ListZones200ResponseInner ListZones(ctx).OrgId(orgId).Id(id).Execute()

List or Get Zones

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
	orgId := "orgId_example" // string |  (optional)
	id := "id_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.ListZones(context.Background()).OrgId(orgId).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.ListZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListZones`: []ListZones200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.ListZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** |  | 
 **id** | **string** |  | 

### Return type

[**[]ListZones200ResponseInner**](ListZones200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

