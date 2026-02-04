# \LedgerServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LedgerServiceBulkRecordTransactions**](LedgerServiceAPI.md#LedgerServiceBulkRecordTransactions) | **Post** /v1/ledger/bulk-transactions | 
[**LedgerServiceCreateAccount**](LedgerServiceAPI.md#LedgerServiceCreateAccount) | **Post** /v1/ledger/accounts | 
[**LedgerServiceGetAccount**](LedgerServiceAPI.md#LedgerServiceGetAccount) | **Get** /v1/ledger/accounts/{accountId} | 
[**LedgerServiceRecordTransaction**](LedgerServiceAPI.md#LedgerServiceRecordTransaction) | **Post** /v1/ledger/transactions | 



## LedgerServiceBulkRecordTransactions

> LedgerBulkRecordResponse LedgerServiceBulkRecordTransactions(ctx).Body(body).Execute()



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
	body := *openapiclient.NewLedgerBulkRecordRequest() // LedgerBulkRecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerServiceAPI.LedgerServiceBulkRecordTransactions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerServiceAPI.LedgerServiceBulkRecordTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LedgerServiceBulkRecordTransactions`: LedgerBulkRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `LedgerServiceAPI.LedgerServiceBulkRecordTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLedgerServiceBulkRecordTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LedgerBulkRecordRequest**](LedgerBulkRecordRequest.md) |  | 

### Return type

[**LedgerBulkRecordResponse**](LedgerBulkRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LedgerServiceCreateAccount

> LedgerCreateAccountResponse LedgerServiceCreateAccount(ctx).Body(body).Execute()



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
	body := *openapiclient.NewLedgerCreateAccountRequest() // LedgerCreateAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerServiceAPI.LedgerServiceCreateAccount(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerServiceAPI.LedgerServiceCreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LedgerServiceCreateAccount`: LedgerCreateAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `LedgerServiceAPI.LedgerServiceCreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLedgerServiceCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LedgerCreateAccountRequest**](LedgerCreateAccountRequest.md) |  | 

### Return type

[**LedgerCreateAccountResponse**](LedgerCreateAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LedgerServiceGetAccount

> LedgerGetAccountResponse LedgerServiceGetAccount(ctx, accountId).Execute()



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
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerServiceAPI.LedgerServiceGetAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerServiceAPI.LedgerServiceGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LedgerServiceGetAccount`: LedgerGetAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `LedgerServiceAPI.LedgerServiceGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLedgerServiceGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LedgerGetAccountResponse**](LedgerGetAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LedgerServiceRecordTransaction

> LedgerRecordTransactionResponse LedgerServiceRecordTransaction(ctx).Body(body).Execute()



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
	body := *openapiclient.NewLedgerRecordTransactionRequest() // LedgerRecordTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerServiceAPI.LedgerServiceRecordTransaction(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerServiceAPI.LedgerServiceRecordTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LedgerServiceRecordTransaction`: LedgerRecordTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `LedgerServiceAPI.LedgerServiceRecordTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLedgerServiceRecordTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LedgerRecordTransactionRequest**](LedgerRecordTransactionRequest.md) |  | 

### Return type

[**LedgerRecordTransactionResponse**](LedgerRecordTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

