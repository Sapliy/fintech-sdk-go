# \AuthAPI

All URIs are relative to *https://api.sapliy.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AuthLoginPost**](AuthAPI.md#V1AuthLoginPost) | **Post** /v1/auth/login | Login
[**V1AuthRegisterPost**](AuthAPI.md#V1AuthRegisterPost) | **Post** /v1/auth/register | Register a new user



## V1AuthLoginPost

> V1AuthRegisterPost201Response V1AuthLoginPost(ctx).V1AuthRegisterPostRequest(v1AuthRegisterPostRequest).Execute()

Login

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
	v1AuthRegisterPostRequest := *openapiclient.NewV1AuthRegisterPostRequest("Email_example", "Password_example") // V1AuthRegisterPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.V1AuthLoginPost(context.Background()).V1AuthRegisterPostRequest(v1AuthRegisterPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.V1AuthLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AuthLoginPost`: V1AuthRegisterPost201Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.V1AuthLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AuthRegisterPostRequest** | [**V1AuthRegisterPostRequest**](V1AuthRegisterPostRequest.md) |  | 

### Return type

[**V1AuthRegisterPost201Response**](V1AuthRegisterPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthRegisterPost

> V1AuthRegisterPost201Response V1AuthRegisterPost(ctx).V1AuthRegisterPostRequest(v1AuthRegisterPostRequest).Execute()

Register a new user

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
	v1AuthRegisterPostRequest := *openapiclient.NewV1AuthRegisterPostRequest("Email_example", "Password_example") // V1AuthRegisterPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.V1AuthRegisterPost(context.Background()).V1AuthRegisterPostRequest(v1AuthRegisterPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.V1AuthRegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AuthRegisterPost`: V1AuthRegisterPost201Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.V1AuthRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AuthRegisterPostRequest** | [**V1AuthRegisterPostRequest**](V1AuthRegisterPostRequest.md) |  | 

### Return type

[**V1AuthRegisterPost201Response**](V1AuthRegisterPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

