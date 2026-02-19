# \WalletsAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWallet**](WalletsAPI.md#GetWallet) | **Get** /v1/wallets/{user_id} | Get Wallet Balance
[**TopupWallet**](WalletsAPI.md#TopupWallet) | **Post** /v1/wallets/topup | Top up a wallet
[**TransferWallet**](WalletsAPI.md#TransferWallet) | **Post** /v1/wallets/transfer | Transfer between wallets



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


## TopupWallet

> TopupWallet200Response TopupWallet(ctx).XZoneID(xZoneID).TopupWalletRequest(topupWalletRequest).XZoneMode(xZoneMode).Execute()

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
	xZoneID := "xZoneID_example" // string | The ID of the zone for this request.
	topupWalletRequest := *openapiclient.NewTopupWalletRequest(int64(123), "Currency_example", "ReferenceId_example") // TopupWalletRequest | 
	xZoneMode := "xZoneMode_example" // string | The mode of the zone (live or test). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.TopupWallet(context.Background()).XZoneID(xZoneID).TopupWalletRequest(topupWalletRequest).XZoneMode(xZoneMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.TopupWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TopupWallet`: TopupWallet200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.TopupWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTopupWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xZoneID** | **string** | The ID of the zone for this request. | 
 **topupWalletRequest** | [**TopupWalletRequest**](TopupWalletRequest.md) |  | 
 **xZoneMode** | **string** | The mode of the zone (live or test). | 

### Return type

[**TopupWallet200Response**](TopupWallet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferWallet

> TopupWallet200Response TransferWallet(ctx).XZoneID(xZoneID).TransferWalletRequest(transferWalletRequest).XZoneMode(xZoneMode).Execute()

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
	xZoneID := "xZoneID_example" // string | The ID of the zone for this request.
	transferWalletRequest := *openapiclient.NewTransferWalletRequest("ToUserId_example", int64(123), "Currency_example", "ReferenceId_example") // TransferWalletRequest | 
	xZoneMode := "xZoneMode_example" // string | The mode of the zone (live or test). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.TransferWallet(context.Background()).XZoneID(xZoneID).TransferWalletRequest(transferWalletRequest).XZoneMode(xZoneMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.TransferWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferWallet`: TopupWallet200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.TransferWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xZoneID** | **string** | The ID of the zone for this request. | 
 **transferWalletRequest** | [**TransferWalletRequest**](TransferWalletRequest.md) |  | 
 **xZoneMode** | **string** | The mode of the zone (live or test). | 

### Return type

[**TopupWallet200Response**](TopupWallet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

