# \LedgerAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLedgerAccount**](LedgerAPI.md#GetLedgerAccount) | **Get** /v1/ledger/accounts/{id} | Get Ledger Account details
[**GetLedgerTransaction**](LedgerAPI.md#GetLedgerTransaction) | **Get** /v1/ledger/transactions/{id} | Get Ledger Transaction details
[**V1LedgerAccountsPost**](LedgerAPI.md#V1LedgerAccountsPost) | **Post** /v1/ledger/accounts | Create Ledger Account
[**V1LedgerTransactionsPost**](LedgerAPI.md#V1LedgerTransactionsPost) | **Post** /v1/ledger/transactions | Record Transaction



## GetLedgerAccount

> LedgerAccount GetLedgerAccount(ctx, id).XZoneID(xZoneID).Execute()

Get Ledger Account details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/sapliy-sdk-go"
)

func main() {
	id := "id_example" // string | 
	xZoneID := "xZoneID_example" // string | The ID of the zone for this request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerAPI.GetLedgerAccount(context.Background(), id).XZoneID(xZoneID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerAPI.GetLedgerAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerAccount`: LedgerAccount
	fmt.Fprintf(os.Stdout, "Response from `LedgerAPI.GetLedgerAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xZoneID** | **string** | The ID of the zone for this request. | 

### Return type

[**LedgerAccount**](LedgerAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLedgerTransaction

> LedgerTransaction GetLedgerTransaction(ctx, id).XZoneID(xZoneID).Execute()

Get Ledger Transaction details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/sapliy-sdk-go"
)

func main() {
	id := "id_example" // string | 
	xZoneID := "xZoneID_example" // string | The ID of the zone for this request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerAPI.GetLedgerTransaction(context.Background(), id).XZoneID(xZoneID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerAPI.GetLedgerTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerTransaction`: LedgerTransaction
	fmt.Fprintf(os.Stdout, "Response from `LedgerAPI.GetLedgerTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xZoneID** | **string** | The ID of the zone for this request. | 

### Return type

[**LedgerTransaction**](LedgerTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LedgerAccountsPost

> LedgerAccount V1LedgerAccountsPost(ctx).XZoneID(xZoneID).V1LedgerAccountsPostRequest(v1LedgerAccountsPostRequest).Execute()

Create Ledger Account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/sapliy-sdk-go"
)

func main() {
	xZoneID := "xZoneID_example" // string | The ID of the zone for this request.
	v1LedgerAccountsPostRequest := *openapiclient.NewV1LedgerAccountsPostRequest("Name_example", "Type_example") // V1LedgerAccountsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerAPI.V1LedgerAccountsPost(context.Background()).XZoneID(xZoneID).V1LedgerAccountsPostRequest(v1LedgerAccountsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerAPI.V1LedgerAccountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LedgerAccountsPost`: LedgerAccount
	fmt.Fprintf(os.Stdout, "Response from `LedgerAPI.V1LedgerAccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LedgerAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xZoneID** | **string** | The ID of the zone for this request. | 
 **v1LedgerAccountsPostRequest** | [**V1LedgerAccountsPostRequest**](V1LedgerAccountsPostRequest.md) |  | 

### Return type

[**LedgerAccount**](LedgerAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LedgerTransactionsPost

> V1LedgerTransactionsPost201Response V1LedgerTransactionsPost(ctx).XZoneID(xZoneID).V1LedgerTransactionsPostRequest(v1LedgerTransactionsPostRequest).Execute()

Record Transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/sapliy-sdk-go"
)

func main() {
	xZoneID := "xZoneID_example" // string | The ID of the zone for this request.
	v1LedgerTransactionsPostRequest := *openapiclient.NewV1LedgerTransactionsPostRequest("ReferenceId_example", []openapiclient.LedgerEntry{*openapiclient.NewLedgerEntry("AccountId_example", int64(123))}) // V1LedgerTransactionsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerAPI.V1LedgerTransactionsPost(context.Background()).XZoneID(xZoneID).V1LedgerTransactionsPostRequest(v1LedgerTransactionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerAPI.V1LedgerTransactionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LedgerTransactionsPost`: V1LedgerTransactionsPost201Response
	fmt.Fprintf(os.Stdout, "Response from `LedgerAPI.V1LedgerTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LedgerTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xZoneID** | **string** | The ID of the zone for this request. | 
 **v1LedgerTransactionsPostRequest** | [**V1LedgerTransactionsPostRequest**](V1LedgerTransactionsPostRequest.md) |  | 

### Return type

[**V1LedgerTransactionsPost201Response**](V1LedgerTransactionsPost201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

