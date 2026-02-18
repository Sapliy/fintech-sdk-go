# \WalletsAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWallet**](WalletsAPI.md#GetWallet) | **Get** /v1/wallets/{user_id} | Get Wallet Balance
[**V1WalletsTopupPost**](WalletsAPI.md#V1WalletsTopupPost) | **Post** /v1/wallets/topup | Top up a wallet
[**V1WalletsTransferPost**](WalletsAPI.md#V1WalletsTransferPost) | **Post** /v1/wallets/transfer | Transfer between wallets



## GetWallet

> Wallet GetWallet(ctx, userId).XZoneID(xZoneID).XZoneMode(xZoneMode).Execute()

Get Wallet Balance

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
	userId := "userId_example" // string | 
	xZoneID := "xZoneID_example" // string | The ID of the zone for this request.
	xZoneMode := "xZoneMode_example" // string | The mode of the zone (live or test). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWallet(context.Background(), userId).XZoneID(xZoneID).XZoneMode(xZoneMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallet`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xZoneID** | **string** | The ID of the zone for this request. | 
 **xZoneMode** | **string** | The mode of the zone (live or test). | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WalletsTopupPost

> V1WalletsTopupPost200Response V1WalletsTopupPost(ctx).V1WalletsTopupPostRequest(v1WalletsTopupPostRequest).Execute()

Top up a wallet

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
	v1WalletsTopupPostRequest := *openapiclient.NewV1WalletsTopupPostRequest(int64(123), "Currency_example", "ReferenceId_example") // V1WalletsTopupPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.V1WalletsTopupPost(context.Background()).V1WalletsTopupPostRequest(v1WalletsTopupPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.V1WalletsTopupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WalletsTopupPost`: V1WalletsTopupPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.V1WalletsTopupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1WalletsTopupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1WalletsTopupPostRequest** | [**V1WalletsTopupPostRequest**](V1WalletsTopupPostRequest.md) |  | 

### Return type

[**V1WalletsTopupPost200Response**](V1WalletsTopupPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WalletsTransferPost

> V1WalletsTopupPost200Response V1WalletsTransferPost(ctx).V1WalletsTransferPostRequest(v1WalletsTransferPostRequest).Execute()

Transfer between wallets

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
	v1WalletsTransferPostRequest := *openapiclient.NewV1WalletsTransferPostRequest("ToUserId_example", int64(123), "Currency_example", "ReferenceId_example") // V1WalletsTransferPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.V1WalletsTransferPost(context.Background()).V1WalletsTransferPostRequest(v1WalletsTransferPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.V1WalletsTransferPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WalletsTransferPost`: V1WalletsTopupPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.V1WalletsTransferPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1WalletsTransferPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1WalletsTransferPostRequest** | [**V1WalletsTransferPostRequest**](V1WalletsTransferPostRequest.md) |  | 

### Return type

[**V1WalletsTopupPost200Response**](V1WalletsTopupPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

