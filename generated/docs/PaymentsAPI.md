# \PaymentsAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmPaymentIntent**](PaymentsAPI.md#ConfirmPaymentIntent) | **Post** /v1/payments/intents/{id}/confirm | Confirm a Payment Intent
[**CreatePaymentIntent**](PaymentsAPI.md#CreatePaymentIntent) | **Post** /v1/payments | Create a Payment Intent
[**GetPaymentIntent**](PaymentsAPI.md#GetPaymentIntent) | **Get** /v1/payments/{id} | Get Payment Intent details



## ConfirmPaymentIntent

> PaymentIntent ConfirmPaymentIntent(ctx, id).XZoneID(xZoneID).XZoneMode(xZoneMode).ConfirmPaymentIntentRequest(confirmPaymentIntentRequest).Execute()

Confirm a Payment Intent

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
	xZoneMode := "xZoneMode_example" // string | The mode of the zone (live or test). (optional)
	confirmPaymentIntentRequest := *openapiclient.NewConfirmPaymentIntentRequest() // ConfirmPaymentIntentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ConfirmPaymentIntent(context.Background(), id).XZoneID(xZoneID).XZoneMode(xZoneMode).ConfirmPaymentIntentRequest(confirmPaymentIntentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ConfirmPaymentIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmPaymentIntent`: PaymentIntent
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ConfirmPaymentIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmPaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xZoneID** | **string** | The ID of the zone for this request. | 
 **xZoneMode** | **string** | The mode of the zone (live or test). | 
 **confirmPaymentIntentRequest** | [**ConfirmPaymentIntentRequest**](ConfirmPaymentIntentRequest.md) |  | 

### Return type

[**PaymentIntent**](PaymentIntent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentIntent

> PaymentIntent CreatePaymentIntent(ctx).XZoneID(xZoneID).CreatePaymentIntentRequest(createPaymentIntentRequest).XZoneMode(xZoneMode).Execute()

Create a Payment Intent

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
	createPaymentIntentRequest := *openapiclient.NewCreatePaymentIntentRequest(int64(123), "Currency_example") // CreatePaymentIntentRequest | 
	xZoneMode := "xZoneMode_example" // string | The mode of the zone (live or test). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.CreatePaymentIntent(context.Background()).XZoneID(xZoneID).CreatePaymentIntentRequest(createPaymentIntentRequest).XZoneMode(xZoneMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.CreatePaymentIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentIntent`: PaymentIntent
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.CreatePaymentIntent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xZoneID** | **string** | The ID of the zone for this request. | 
 **createPaymentIntentRequest** | [**CreatePaymentIntentRequest**](CreatePaymentIntentRequest.md) |  | 
 **xZoneMode** | **string** | The mode of the zone (live or test). | 

### Return type

[**PaymentIntent**](PaymentIntent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentIntent

> PaymentIntent GetPaymentIntent(ctx, id).XZoneID(xZoneID).XZoneMode(xZoneMode).Execute()

Get Payment Intent details

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
	xZoneMode := "xZoneMode_example" // string | The mode of the zone (live or test). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.GetPaymentIntent(context.Background(), id).XZoneID(xZoneID).XZoneMode(xZoneMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPaymentIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentIntent`: PaymentIntent
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPaymentIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xZoneID** | **string** | The ID of the zone for this request. | 
 **xZoneMode** | **string** | The mode of the zone (live or test). | 

### Return type

[**PaymentIntent**](PaymentIntent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

