# \WalletServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletServiceCreateWallet**](WalletServiceAPI.md#WalletServiceCreateWallet) | **Post** /v1/wallets | 
[**WalletServiceGetWallet**](WalletServiceAPI.md#WalletServiceGetWallet) | **Get** /v1/wallets/{userId} | 
[**WalletServiceTopUp**](WalletServiceAPI.md#WalletServiceTopUp) | **Post** /v1/wallets/top-up | 
[**WalletServiceTransfer**](WalletServiceAPI.md#WalletServiceTransfer) | **Post** /v1/wallets/transfer | 



## WalletServiceCreateWallet

> WalletWallet WalletServiceCreateWallet(ctx).Body(body).Execute()



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
	body := *openapiclient.NewWalletCreateWalletRequest() // WalletCreateWalletRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletServiceAPI.WalletServiceCreateWallet(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletServiceAPI.WalletServiceCreateWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletServiceCreateWallet`: WalletWallet
	fmt.Fprintf(os.Stdout, "Response from `WalletServiceAPI.WalletServiceCreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletServiceCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WalletCreateWalletRequest**](WalletCreateWalletRequest.md) |  | 

### Return type

[**WalletWallet**](WalletWallet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletServiceGetWallet

> WalletWallet WalletServiceGetWallet(ctx, userId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletServiceAPI.WalletServiceGetWallet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletServiceAPI.WalletServiceGetWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletServiceGetWallet`: WalletWallet
	fmt.Fprintf(os.Stdout, "Response from `WalletServiceAPI.WalletServiceGetWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWalletServiceGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WalletWallet**](WalletWallet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletServiceTopUp

> WalletTransactionResponse WalletServiceTopUp(ctx).Body(body).Execute()



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
	body := *openapiclient.NewWalletTopUpRequest() // WalletTopUpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletServiceAPI.WalletServiceTopUp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletServiceAPI.WalletServiceTopUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletServiceTopUp`: WalletTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletServiceAPI.WalletServiceTopUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletServiceTopUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WalletTopUpRequest**](WalletTopUpRequest.md) |  | 

### Return type

[**WalletTransactionResponse**](WalletTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletServiceTransfer

> WalletTransactionResponse WalletServiceTransfer(ctx).Body(body).Execute()



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
	body := *openapiclient.NewWalletTransferRequest() // WalletTransferRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletServiceAPI.WalletServiceTransfer(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletServiceAPI.WalletServiceTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletServiceTransfer`: WalletTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletServiceAPI.WalletServiceTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletServiceTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WalletTransferRequest**](WalletTransferRequest.md) |  | 

### Return type

[**WalletTransactionResponse**](WalletTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

